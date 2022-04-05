package rest

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/apex/log"
)

type LogContent struct {
	Timestamp    int64  `json:"timestamp"`
	Level        string `json:"level"`
	Message      string `json:"content"`
	FleetId      string `json:"fleet.id"`
	TaxiId       string `json:"taxi.id"`
	CustomDevice string `json:"dt.entity.custom_device"`
}

// Handler implementation
type Handler struct {
	mu     sync.Mutex
	Client DTClient
}

// New handler with DT Client
func New(dtc DTClient) *Handler {
	return &Handler{
		Client: dtc,
	}
}

// HandleLog implements log.Handler
func (h *Handler) HandleLog(e *log.Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	// build content
	logContent := LogContent{
		Timestamp:    e.Timestamp.UTC().UnixMilli(),
		Level:        e.Level.String(),
		Message:      e.Message,
		FleetId:      fmt.Sprintf("%v", e.Fields.Get("fleet.id")),
		CustomDevice: fmt.Sprintf("%v", e.Fields.Get("custom.device")),
	}
	if e.Fields.Get("taxi.id") != "" {
		logContent.TaxiId = fmt.Sprintf("%v", e.Fields.Get("taxi.id"))
	}
	content, _ := json.Marshal(logContent)

	// send to client
	h.Client.PostLogEvent(content)

	return nil
}
