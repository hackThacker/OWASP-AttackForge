package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hackThacker/OWASP-AttackForge/backend/config"
	deliveryGRPC "github.com/hackThacker/OWASP-AttackForge/backend/delivery/grpc"
	deliveryHTTP "github.com/hackThacker/OWASP-AttackForge/backend/delivery/http"
	deliveryWS "github.com/hackThacker/OWASP-AttackForge/backend/delivery/websocket"
	repoDocker "github.com/hackThacker/OWASP-AttackForge/backend/repository/docker"
	"github.com/hackThacker/OWASP-AttackForge/backend/usecase"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting OWASP AttackForge Backend Server...")

	// 1. Load Configurations
	cfg := config.LoadConfig()

	// 2. Initialize Repositories (Docker SDK)
	dockerRepo, err := repoDocker.NewDockerRepository()
	if err != nil {
		log.Fatalf("Failed to initialize Docker SDK repository: %v", err)
	}

	// 3. Initialize Usecases
	toolUC := usecase.NewToolUsecase(dockerRepo)
	metricUC := usecase.NewMetricUsecase(dockerRepo)
	suggestionUC := usecase.NewSuggestionUsecase()

	// 4. Initialize WebSocket Hub
	hub := deliveryWS.NewHub(toolUC, metricUC)
	ctx, hubCancel := context.WithCancel(context.Background())
	go hub.Run(ctx)

	// 5. Start gRPC Server
	grpcLis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen for gRPC on %s: %v", cfg.GRPCPort, err)
	}
	grpcServer := grpc.NewServer()
	deliveryGRPC.RegisterServer(grpcServer, toolUC, metricUC)

	go func() {
		log.Printf("gRPC server listening on %s", cfg.GRPCPort)
		if err := grpcServer.Serve(grpcLis); err != nil && err != grpc.ErrServerStopped {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()

	// 6. Start HTTP Server (REST + WebSocket)
	httpHandler := deliveryHTTP.NewHandler(toolUC, metricUC, suggestionUC, hub)
	httpServer := &http.Server{
		Addr:    cfg.Port,
		Handler: httpHandler,
	}

	go func() {
		log.Printf("REST/WebSocket server listening on %s", cfg.Port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// 7. Handle OS Signals for Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down servers gracefully...")

	// Shutdown REST/WebSocket server
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP shutdown error: %v", err)
	}

	// Stop gRPC server
	grpcServer.GracefulStop()

	// Stop WebSocket Hub connection pumps
	hubCancel()

	log.Println("OWASP AttackForge Backend stopped successfully.")
}
