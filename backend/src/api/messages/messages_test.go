package messages

import (
	"api"
	"api/middlewares"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	"interfaces"
	"io/ioutil"
	"log"
	meetingsAPIMock "mock/api"
	mock "mock/repositories"
	"models"
	"os"
	"plugins/config"
	"repositories"
	"services"
	"services/errors"
	"services/proxies/validation"
	"strings"
	"testing"
	"utils"
)

var (
	db             *sqlx.DB
	sessionService interfaces.SessionAccessorService
	router         = api.GetRouter()
)

func init() {
	utils.SkipInShortMode()

	var err error
	db, err = config.GetConfiguredConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	coderKey, err := config.GetCoderKey()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sessionService = services.Session(coderKey)
	InitRequestHandlers(
		services.Messages(repositories.Messages(db)),
		middlewares.AuthSession{Service: sessionService}.HasValidSession,
	)
}

func getWS(serverURL string) *websocket.Conn {
	path := "ws" + strings.TrimPrefix(serverURL, "http") + "/ws"
	ws, res, err := websocket.DefaultDialer.Dial(path, nil)
	if err != nil {
		if err == websocket.ErrBadHandshake && res != nil {
			fmt.Printf("ErrBadHandshake. Status: %s\n", res.Status)
		}

		panic(err)
	}

	return ws
}

func TestMain(m *testing.M) {
	mock.DropTables(db)
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestGetMessages_Success(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response meetingsAPIMock.MessagesResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.GetMessagesRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
	utils.AssertTrue(len(response.Data) <= meetingsAPIMock.DefaultMessagesCount, t)
}

func TestGetMessages_InvalidData(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.GetMessagesInvalidDataRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	for _, errorMessage := range []string{validation.InvalidId, validation.InvalidCount} {
		utils.AssertTrue(strings.Contains(response.ErrorDetail, errorMessage), t)
	}
}

func TestGetMessages_InternalError(t *testing.T) {
	mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.GetMessagesRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestGetMessagesAfter_Success(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response meetingsAPIMock.MessagesResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.GetMessagesAfterMessageRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
	utils.AssertTrue(len(response.Data) <= meetingsAPIMock.DefaultMessagesCount, t)
	for _, message := range response.Data {
		utils.AssertTrue(message.SendingTime.After(mock.GetFirstMessageSendingTime()), t)
	}
}

func TestGetMessagesAfter_InvalidData(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.GetMessagesAfterMessageInvalidDataRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	for _, errorMessage := range []string{validation.InvalidId, validation.InvalidCount} {
		utils.AssertTrue(strings.Contains(response.ErrorDetail, errorMessage), t)
	}
}

func TestGetMessagesAfter_InternalError(t *testing.T) {
	mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.GetMessagesAfterMessageRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestSendMessage_OneConnection(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	testServer := utils.GetTestServer(router)
	defer testServer.Close()

	ws := getWS(testServer.URL)
	defer func() {
		_ = ws.Close()
	}()
	renewWS := func() {
		_ = ws.Close()
		ws = getWS(testServer.URL)
	}

	t.Run("Write simple message", func(t *testing.T) {
		err := ws.WriteJSON(meetingsAPIMock.GetSimpleMessage())
		utils.AssertNil(err, t)

		var message models.Message
		err = ws.ReadJSON(&message)
		utils.AssertNil(err, t)
		utils.AssertEqual(meetingsAPIMock.GetSimpleMessage(), message, t)
	})

	t.Run("Chat not found", func(t *testing.T) {
		defer renewWS()

		err := ws.WriteJSON(meetingsAPIMock.GetMessageWithNotExistsChatId())
		utils.AssertNil(err, t)

		var message models.ErrorResponse
		err = ws.ReadJSON(&message)
		utils.AssertNil(err, t)
		utils.AssertEqual(api.StatusError, message.Status, t)
		utils.AssertEqual(errors.ChatIdNotFound.Error(), message.ErrorDetail, t)
	})

	t.Run("User not found", func(t *testing.T) {
		defer renewWS()

		err := ws.WriteJSON(meetingsAPIMock.GetMessageWithNotExistsSenderId())
		utils.AssertNil(err, t)

		var message models.ErrorResponse
		err = ws.ReadJSON(&message)
		utils.AssertNil(err, t)
		utils.AssertEqual(api.StatusError, message.Status, t)
		utils.AssertEqual(errors.UserIdNotFound.Error(), message.ErrorDetail, t)
	})

	t.Run("Internal error", func(t *testing.T) {
		mock.DropTables(db)
		defer mock.InitTables(db)
		defer renewWS()

		err := ws.WriteJSON(meetingsAPIMock.GetSimpleMessage())
		utils.AssertNil(err, t)

		var message models.ErrorResponse
		err = ws.ReadJSON(&message)
		utils.AssertNil(err, t)
		utils.AssertEqual(api.StatusError, message.Status, t)
		utils.AssertEqual(errors.InternalError.Error(), message.ErrorDetail, t)
	})

	t.Run("Invalid data", func(t *testing.T) {
		defer renewWS()

		err := ws.WriteJSON(meetingsAPIMock.GetMessageWithInvalidData())
		utils.AssertNil(err, t)

		var message models.ErrorResponse
		err = ws.ReadJSON(&message)
		utils.AssertNil(err, t)
		utils.AssertEqual(api.StatusError, message.Status, t)
		for _, errorMessage := range []string{validation.InvalidId, validation.InvalidMessageText} {
			utils.AssertTrue(strings.Contains(message.ErrorDetail, errorMessage), t)
		}
	})

	t.Run("Bad message", func(t *testing.T) {
		defer renewWS()

		err := ws.WriteMessage(websocket.TextMessage, []byte(``))
		utils.AssertNil(err, t)

		var message models.ErrorResponse
		err = ws.ReadJSON(&message)
		utils.AssertNil(err, t)
		utils.AssertEqual(api.StatusError, message.Status, t)
		utils.AssertEqual(ReadJSONError.Error(), message.ErrorDetail, t)
	})
}

func TestSendMessage_TwoConnections(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	testServer := utils.GetTestServer(router)
	defer testServer.Close()

	ws1, ws2 := getWS(testServer.URL), getWS(testServer.URL)
	defer func() {
		_ = ws1.Close()
		_ = ws2.Close()
	}()
	renewWSs := func() {
		_ = ws1.Close()
		_ = ws2.Close()
		ws1, ws2 = getWS(testServer.URL), getWS(testServer.URL)
	}

	t.Run("Write simple messages", func(t *testing.T) {
		defer renewWSs()

		var tmpMessage models.Message
		message1 := meetingsAPIMock.GetSimpleMessage()
		_ = ws1.WriteJSON(message1)
		_ = ws1.ReadJSON(&tmpMessage)

		message2 := meetingsAPIMock.GetAnotherSimpleMessage()
		_ = ws2.WriteJSON(message2)
		_ = ws1.ReadJSON(&tmpMessage)

		utils.AssertEqual(message2, tmpMessage, t)
	})

	t.Run("Try send message to closed connection", func(t *testing.T) {
		defer renewWSs()

		var tmpMessage models.Message
		message1 := meetingsAPIMock.GetSimpleMessage()
		_ = ws1.WriteJSON(message1)
		_ = ws1.ReadJSON(&tmpMessage)

		message2 := meetingsAPIMock.GetAnotherSimpleMessage()
		_ = ws1.Close()
		_ = ws2.WriteJSON(message2)

		err := ws1.ReadJSON(&tmpMessage)
		utils.AssertNotNil(err, t)

		err = ws2.ReadJSON(&tmpMessage)
		utils.AssertNil(err, t)
		utils.AssertEqual(message2, tmpMessage, t)
	})
}

func TestSendMessage_BadProtocol(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.InvalidProtocolRequest(router))).Decode(&response)

	utils.AssertNotNil(err, t)
}
