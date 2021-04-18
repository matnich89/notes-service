package test

import (
	"gorm.io/gorm"
	"net/http"
	"notes-service/internal/handler"
	"notes-service/internal/logger"
	"notes-service/internal/note"
	"testing"
)

func TestHealth(t *testing.T) {
	app := newTestApplication(t)
	h := handler.NewHandler(note.NewService(&gorm.DB{}), logger.NewLogger())
	app.DefineRouting(h)
	ts := newTestServer(t, app.GetRouter())
	defer ts.Close()
	code, _, _ := ts.get(t, "/api/health")

	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}
}
