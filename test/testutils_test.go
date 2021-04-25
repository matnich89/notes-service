package test

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"notes-service/cmd/server/app"
	"notes-service/internal/logger"
	"notes-service/internal/note"
	"testing"
)

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewServer(h)

	return &testServer{ts}
}

func newTestApplication() *app.App {
	return app.NewApp(logger.NewLogger(), mux.NewRouter())
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, []byte) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, rs.Header, body
}

func (ts *testServer) postNote(t *testing.T, urlath string, note *note.Note) (int, http.Header, []byte) {
	noteBytes, _ := json.Marshal(note)
	rs, err := ts.Client().Post(ts.URL+urlath, "application/json; charset=UTF-8", bytes.NewReader(noteBytes))
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(rs.Body)

	return rs.StatusCode, rs.Header, body
}
