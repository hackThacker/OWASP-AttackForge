package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/hackThacker/OWASP-AttackForge/backend/delivery/websocket"
	"github.com/hackThacker/OWASP-AttackForge/backend/domain"
)

type httpHandler struct {
	toolUC       domain.ToolUsecase
	metricUC     domain.MetricUsecase
	suggestionUC domain.SuggestionUsecase
	hub          *websocket.Hub
}

func NewHandler(toolUC domain.ToolUsecase, metricUC domain.MetricUsecase, suggestionUC domain.SuggestionUsecase, hub *websocket.Hub) http.Handler {
	handler := &httpHandler{
		toolUC:       toolUC,
		metricUC:     metricUC,
		suggestionUC: suggestionUC,
		hub:          hub,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/metrics", handler.getMetrics)
	mux.HandleFunc("/api/tools", handler.getTools)
	mux.HandleFunc("/api/tools/", handler.handleToolAction)
	mux.HandleFunc("/api/suggestions", handler.getSuggestions)
	mux.HandleFunc("/ws", handler.serveWebsocket)

	// Combine router with CORS & Logging middlewares
	return LoggingMiddleware(CORSMiddleware(mux))
}

func (h *httpHandler) getMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	metrics, err := h.metricUC.GetMetrics(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(metrics)
}

func (h *httpHandler) getTools(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	tools, err := h.toolUC.GetTools(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(tools)
}

func (h *httpHandler) handleToolAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Path parsing: /api/tools/{subdomain}/{action}
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 4 || parts[1] != "tools" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	subdomain := parts[2]
	action := parts[3]

	var err error
	switch action {
	case "start":
		err = h.toolUC.StartTool(r.Context(), subdomain)
	case "stop":
		err = h.toolUC.StopTool(r.Context(), subdomain)
	case "restart":
		err = h.toolUC.RestartTool(r.Context(), subdomain)
	default:
		http.Error(w, "Invalid Action", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Broadcast updated state over WebSockets immediately
	go h.hub.BroadcastUpdate()

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"success"}`))
}

func (h *httpHandler) getSuggestions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	suggestions, err := h.suggestionUC.GetSuggestions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(suggestions)
}

func (h *httpHandler) serveWebsocket(w http.ResponseWriter, r *http.Request) {
	websocket.ServeWs(h.hub, w, r)
}
