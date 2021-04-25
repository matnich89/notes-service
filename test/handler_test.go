package test

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"notes-service/internal/handler"
	"notes-service/internal/logger"
	"notes-service/internal/note"
	"testing"
)

func TestHealth(t *testing.T) {
	app := newTestApplication()
	h := handler.NewHandler(note.NewService(&gorm.DB{}), logger.NewLogger())
	app.DefineRouting(h)
	ts := newTestServer(t, app.GetRouter())
	defer ts.Close()
	code, _, _ := ts.get(t, "/api/health")

	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}
}

func TestBlah(t *testing.T) {
	var db *sql.DB

	db, mock, _ = sqlmock.New()

	gdb, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))

	app := newTestApplication()

	h := handler.NewHandler(note.NewService(gdb), logger.NewLogger())
	app.DefineRouting(h)
	ts := newTestServer(t, app.GetRouter())
	defer ts.Close()
	note := &note.Note{
		Model: gorm.Model{},
		Name:  "Order 66",
		Theme: "Be the Emperor",
		Text:  "Execute Order 66!",
		Owner: 0,
	}
	code, headers, body := ts.postNote(t, "/api/note", note)

	fmt.Println(code)
	fmt.Println(headers)
	fmt.Println(body)

}
