package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
	"github.com/golang/mock/gomock"
	"github.com/sambaiz/go-api-with-test/service"
	"github.com/sambaiz/go-api-with-test/model"
)

func TestGetMessage(t *testing.T) {
	expectResponseJSON := `[{"id":1,"content":"メッセージ"}]`

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/messages?limit=1&offset=1", nil)
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	messageServiceMoc := service.NewMockMessage(ctl)
	messageServiceMoc.EXPECT().GetMessages(uint64(1), uint64(1)).Return([]model.Message{{ID: 1, Content: "メッセージ"}}, nil)

	h := NewMessage(messageServiceMoc)
	err = h.GetMessages(c)
	if err != nil{
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, expectResponseJSON, rec.Body.String())
}

func TestGetMessageWithNoLimit(t *testing.T) {
	expectResponseJSON := `{"error":"Limit is required and must be a number."}`

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/messages?offset=1", nil)
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	messageServiceMoc := service.NewMockMessage(ctl)

	h := NewMessage(messageServiceMoc)
	err = h.GetMessages(c)
	if err != nil{
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, expectResponseJSON, rec.Body.String())
}

func TestGetMessageWithNoOffset(t *testing.T) {
	expectResponseJSON := `[{"id":1,"content":"メッセージ"}]`

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/messages?limit=1", nil)
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	messageServiceMoc := service.NewMockMessage(ctl)
	messageServiceMoc.EXPECT().GetMessages(uint64(1), uint64(0)).Return([]model.Message{{ID: 1, Content: "メッセージ"}}, nil)

	h := NewMessage(messageServiceMoc)
	err = h.GetMessages(c)
	if err != nil{
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, expectResponseJSON, rec.Body.String())
}

func TestGetMessageWithStringOffset(t *testing.T) {
	expectResponseJSON := `{"error":"Offset must be a number."}`

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/messages?limit=1&offset=a", nil)
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	messageServiceMoc := service.NewMockMessage(ctl)

	h := NewMessage(messageServiceMoc)
	err = h.GetMessages(c)
	if err != nil{
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, expectResponseJSON, rec.Body.String())
}

func TestCreateMessage(t *testing.T) {
	requestJSON := `{"content":"メッセージ"}`
	expectResponseJSON := `{"id":1,"content":"メッセージ"}`

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/messages", strings.NewReader(requestJSON))
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	messageServiceMoc := service.NewMockMessage(ctl)
	messageServiceMoc.EXPECT().CreateMessage("メッセージ").Return(&model.Message{ID: 1, Content: "メッセージ"}, nil)

	h := NewMessage(messageServiceMoc)
	err = h.CreateMessage(c)
	if err != nil{
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, expectResponseJSON, rec.Body.String())
}