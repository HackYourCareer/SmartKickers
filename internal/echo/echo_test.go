package echo

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"net/http/httptest"
	"remote/pkg/messages"
	"strings"
	"testing"
)

func Test_sendScore(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(SendScoreHandler))
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()
	// Send message to server, read response and check to see if it's what we expect.
	gameScore := messages.GameScore{BlueScore: 1, WhiteScore: 2}
	gameScoreMsg, _ := json.Marshal(gameScore)
	for i := 0; i < 10; i++ {

		if err := ws.WriteMessage(websocket.TextMessage, gameScoreMsg); err != nil {
			t.Fatalf("%v", err)
		}
	}
}
