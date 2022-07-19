package echo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"remote/pkg/messages"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func NewWsServer(t *testing.T, h http.Handler) (*httptest.Server, *websocket.Conn) {
	t.Helper()

	s := httptest.NewServer(h)
	wsURL := httpToWs(t, s.URL)

	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	return s, ws
}

func httpToWs(t *testing.T, u string) string {
	t.Helper()

	wsURL, err := url.Parse(u)
	if err != nil {
		t.Fatal(err)
	}

	switch wsURL.Scheme {
	case "http":
		wsURL.Scheme = "ws"
	case "https":
		wsURL.Scheme = "wss"
	}

	return wsURL.String()
}

func sendMessage(t *testing.T, ws *websocket.Conn, msg messages.DispatcherMsg) {
	t.Helper()

	m, err := json.Marshal(msg)
	if err != nil {
		t.Fatal(err)
	}

	if err := ws.WriteMessage(websocket.BinaryMessage, m); err != nil {
		t.Fatal(err)
	}
}

func ReceiveMessage(t *testing.T, ws *websocket.Conn) messages.DispatcherResponse {
	t.Helper()

	_, m, err := ws.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}

	var reply messages.DispatcherResponse
	err = json.Unmarshal(m, &reply)
	if err != nil {
		t.Fatal(err)
	}

	return reply
}

func TestEcho(t *testing.T) {
	tables := []struct {
		name    string
		message messages.DispatcherMsg
		reply   messages.DispatcherResponse
	}{
		{
			name:    "Initial message",
			message: messages.DispatcherMsg{TableId: "10", Origin: "PROCESSING", MsgType: "INITIAL"},
			reply:   messages.DispatcherResponse{GameId: "10"},
		},
		{
			name:    "Goal message",
			message: messages.DispatcherMsg{Goal: 1},
			reply:   messages.DispatcherResponse{},
		},
	}

	for _, tt := range tables {
		t.Run(tt.name, func(t *testing.T) {
			h := http.HandlerFunc(Echo)
			s, ws := NewWsServer(t, h)
			defer s.Close()
			defer ws.Close()

			sendMessage(t, ws, tt.message)

			if (tt.reply != messages.DispatcherResponse{}) {

				reply := ReceiveMessage(t, ws)

				if reply != tt.reply {
					t.Fatalf("Expected '%+v', got '%+v'", tt.reply, reply)
				}
			}
		})
	}
}

func Test_handleGoal(t *testing.T) {
	type args struct {
		team int
	}
	tests := []struct {
		name             string
		args             args
		wantedBlueScore  int
		wantedWhiteScore int
	}{
		{"should increment team white score by one", args{1}, 0, 1},
		{"should increment team blue score by one", args{2}, 1, 0},
		{"should not cause an error", args{-1}, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameScore.WhiteScore = 0
			gameScore.BlueScore = 0
			handleGoal(tt.args.team)
			assert.Equal(t, gameScore.BlueScore, tt.wantedBlueScore, "blue team score changes incorrectly")
			assert.Equal(t, gameScore.WhiteScore, tt.wantedWhiteScore, "white team score changes incorrectly")
		})
	}
}
