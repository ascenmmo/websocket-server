package tests

import (
	"context"
	"encoding/json"
	"fmt"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/ascenmmo/websocket-server/env"
	"github.com/ascenmmo/websocket-server/pkg/api/types"
	"github.com/ascenmmo/websocket-server/pkg/clients/wsGameServer"
	"github.com/ascenmmo/websocket-server/pkg/start"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"
)

var data = ""

var countErr = 0

var (
	clients  = 20
	msgs     = 1000
	dataSize = 32 * 1024

	baseURl  = fmt.Sprintf("ws://%s:%s", env.ServerAddress, env.WebsocketPort)
	restAddr = fmt.Sprintf("http://%s:%s", env.ServerAddress, env.TCPPort)
	token    = env.TokenKey
)

var ctx, cancel = context.WithCancel(context.Background())
var min, max time.Duration
var maxMsgs int

type Message struct {
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

type Request struct {
	Token string  `json:"token,omitempty"`
	Data  Message `json:"data,omitempty"`
}

type Response struct {
	Data Message
}

func TestConnection(t *testing.T) {
	func() {
		for i := 0; i < dataSize/2; i++ {
			data = data + "1"
		}
	}()

	//logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger := zerolog.Logger{}
	go start.StartWebSocket(
		ctx,
		env.ServerAddress,
		env.TCPPort,
		env.WebsocketPort,
		env.TokenKey,
		clients*msgs,
		1,
		logger,
		false,
	)
	time.Sleep(time.Second * 5)

	for i := 0; i < clients; i++ {
		//createRoom(t, createToken(t, i))
		go Listener(t, i)
		go Publisher(t, i)
		time.Sleep(time.Millisecond * 1)
	}

	time.Sleep(time.Second * 5)
	<-ctx.Done()

	fmt.Println(max, min, maxMsgs)
}

func Publisher(t *testing.T, i int) {
	time.Sleep(time.Second)
	connection := newConnection(t, i)
	for j := 0; j < msgs; j++ {
		if ctx.Err() != nil {
			return
		}
		msg := buildMessage(t, i, j)
		err := connection.WriteMessage(websocket.BinaryMessage, msg)
		assert.NoError(t, err)
		time.Sleep(time.Millisecond * 1)
	}
}

func Listener(t *testing.T, i int) {
	defer cancel()
	connection := newConnection(t, i)
	response := listen(t, connection)
	fmt.Println("done pubSub", i, "with msgs", response)

}

func createToken(t *testing.T, i int) string {
	z := 0
	if i > clients/2 {
		z = 1
	}
	gameID := uuid.NewMD5(uuid.UUID{}, []byte(strconv.Itoa(i)))
	roomID := uuid.NewMD5(uuid.UUID{}, []byte(strconv.Itoa(i)+strconv.Itoa(z)))
	userID := uuid.New()

	tokenGen, err := tokengenerator.NewTokenGenerator(token)
	assert.Nil(t, err, "init gen token expected nil")

	token, err := tokenGen.GenerateToken(tokentype.Info{
		GameID: gameID,
		RoomID: roomID,
		UserID: userID,
		TTL:    time.Second * 100,
	}, tokengenerator.AESGCM)
	assert.Nil(t, err, "gen token expected nil")

	return token
}

func createRoom(t *testing.T, userToken string) {
	cli := wsGameServer.New(restAddr)

	err := cli.ServerSettings().CreateRoom(context.Background(), userToken, types.CreateRoomRequest{})
	assert.Nil(t, err, "client.do expected nil")
}

func buildMessage(t *testing.T, i, j int) (msg []byte) {
	data := Message{
		Text:      data,
		CreatedAt: time.Now(),
	}
	if j == msgs-1 {
		data = Message{
			Text:      "close",
			CreatedAt: time.Now(),
		}
	}

	req := Request{
		Token: createToken(t, i),
		Data:  data,
	}
	marshal, err := json.Marshal(req)
	assert.Nil(t, err, "client.do expected nil")
	return marshal
}

func newConnection(t *testing.T, i int) *websocket.Conn {
	url := baseURl + "/api/ws/connect"

	headers := http.Header{}
	headers.Add("token", createToken(t, i))

	conn, _, err := websocket.DefaultDialer.Dial(url, headers)
	assert.Nil(t, err, "Dial expected nil")
	if err != nil {
		fmt.Println("Ошибка подключения:", err, i)
		os.Exit(1)
	}
	return conn
}

func listen(t *testing.T, conn *websocket.Conn) int {
	counter := 0
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				fmt.Println(err)
			} else {
				fmt.Println(err)
			}
			return counter
		}
		if messageType != websocket.BinaryMessage {
			continue
		}
		assert.Nil(t, err, "ReadFromUDP expected nil")

		var res Response
		err = json.Unmarshal(message, &res)
		assert.Nil(t, err, "Unmarshal expected nil")
		counter++
		maxMsgs++
		if res.Data.Text == "close" {
			return counter
		}

		sub := time.Now().Sub(res.Data.CreatedAt)
		if min == 0 {
			min = sub
		}
		if min > sub {
			min = sub
		}

		if max < sub {
			max = sub
		}

	}
}
