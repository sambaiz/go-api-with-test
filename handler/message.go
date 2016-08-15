package handler

import (
	"github.com/sambaiz/go-api-with-test/service"
	"github.com/labstack/echo"
	"net/http"
	"github.com/gocraft/dbr"
	"github.com/sambaiz/go-api-with-test/dao"
	"strconv"
	"github.com/sambaiz/go-api-with-test/response"
)

type Message interface {
	CreateMessage(echo.Context) error
}

type MessageImpl struct {
	messageService service.Message
}

// NewMessage returns Message instance.
func NewMessage(messageService service.Message) *MessageImpl {
	return &MessageImpl{messageService: messageService}
}

// NewMessageWithSession returns Message instance.
func NewMessageWithSession(sess *dbr.Session) *MessageImpl {
	return &MessageImpl{messageService: service.NewMessage(*dao.NewTx(sess), *dao.NewMessage(sess))}
}

// Get messages and responses it in json.
func (m MessageImpl) GetMessages(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil{
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "Limit is required and must be a number."})
	}

	offsetStr := c.QueryParam("offset")
	if offsetStr == ""{
		offsetStr = "0"
	}
	offset, err2 := strconv.Atoi(offsetStr)
	if err2 != nil{
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "Offset must be a number."})
	}

	message, err := m.messageService.GetMessages(uint64(limit), uint64(offset))
	if err != nil{
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Something bad happend."})
	}
	return c.JSON(http.StatusCreated, message)
}

type createMessageParams struct {
	Content  string `json:"content" xml:"content" form:"content"`
}

// Create a message and responses it in json.
func (m MessageImpl) CreateMessage(c echo.Context) error {
	var u createMessageParams
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "Required parameters: content."})
	}
	message, err := m.messageService.CreateMessage(u.Content)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Something bad happend."})
	}
	return c.JSON(http.StatusCreated, message)
}
