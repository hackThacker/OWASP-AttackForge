package websocket

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/hackThacker/OWASP-AttackForge/backend/domain"
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	toolUC     domain.ToolUsecase
	metricUC   domain.MetricUsecase
	mu         sync.Mutex
}

type wsMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

func NewHub(toolUC domain.ToolUsecase, metricUC domain.MetricUsecase) *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		toolUC:     toolUC,
		metricUC:   metricUC,
	}
}

func (h *Hub) Run(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
			// Send initial state
			h.sendState(client)

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			h.mu.Unlock()

		case message := <-h.broadcast:
			h.mu.Lock()
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
			h.mu.Unlock()

		case <-ticker.C:
			h.BroadcastUpdate()
		}
	}
}

func (h *Hub) BroadcastUpdate() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	tools, err := h.toolUC.GetTools(ctx)
	if err != nil {
		log.Printf("Failed to get tools for WS broadcast: %v", err)
		return
	}

	metrics, err := h.metricUC.GetMetrics(ctx)
	if err != nil {
		log.Printf("Failed to get metrics for WS broadcast: %v", err)
		return
	}

	msg := wsMessage{
		Type: "state",
		Payload: map[string]interface{}{
			"tools":   tools,
			"metrics": metrics,
		},
	}

	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Failed to marshal WS broadcast payload: %v", err)
		return
	}

	h.mu.Lock()
	for client := range h.clients {
		select {
		case client.send <- data:
		default:
			close(client.send)
			delete(h.clients, client)
		}
	}
	h.mu.Unlock()
}

func (h *Hub) sendState(client *Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	tools, err := h.toolUC.GetTools(ctx)
	if err != nil {
		return
	}

	metrics, err := h.metricUC.GetMetrics(ctx)
	if err != nil {
		return
	}

	msg := wsMessage{
		Type: "state",
		Payload: map[string]interface{}{
			"tools":   tools,
			"metrics": metrics,
		},
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return
	}

	select {
	case client.send <- data:
	default:
	}
}
