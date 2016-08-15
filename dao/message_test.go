package dao
import (
	"testing"
	"github.com/sambaiz/go-api-with-test/model"
	"github.com/stretchr/testify/assert"
	"errors"
	"strings"
)

func TestMessageCreate(t *testing.T) {
	truncateTables()
	message := model.Message{
		ID: 1,
		Content: strings.Repeat("あ", 2000),
	}

	id, err := messageDao.Create(message.Content)
	if err != nil {
		t.Fatal(err)
	}
	actualMessage, err2 := messageDao.FindById(id)
	if err2 != nil {
		t.Fatal(err2)
	}
	assert.Equal(t, message, *actualMessage, "")
}

func TestMessageCreateContentTooLong(t *testing.T) {
	truncateTables()
	message := model.Message{
		ID: 1,
		Content: strings.Repeat("あ", 2001),
	}
	_, err := messageDao.Create(message.Content)
	assert.Equal(t, err, errors.New("content is too long"), "")
}

func TestMessageFindWithLimitOffset(t *testing.T) {
	truncateTables()
	messages := []model.Message{
		model.Message{
			ID: 1,
			Content: "メッセージ",
		},
		model.Message{
			ID: 2,
			Content: "メッセージ2",
		},
		model.Message{
			ID: 3,
			Content: "メッセージ3",
		},
	}

	for _, message := range messages {
		_, err := messageDao.Create(message.Content)
		if err != nil {
			t.Fatal(err)
		}
	}

	actualMessages, err := messageDao.FindWithLimitOffset(1, 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(actualMessages), 1, "")
	assert.Equal(t, messages[1], actualMessages[0], "")
}

func TestMessageFindById(t *testing.T) {
	truncateTables()
	message := model.Message{
		ID: 1,
		Content: "メッセージ",
	}

	id, err := messageDao.Create(message.Content)
	if err != nil {
		t.Fatal(err)
	}
	actualMessage, err2 := messageDao.FindById(id)
	if err2 != nil {
		t.Fatal(err)
	}
	assert.Equal(t, message, *actualMessage, "")
}

func TestMessageFindByIdNotFound(t *testing.T) {
	truncateTables()
	_, err := messageDao.FindById(1)
	assert.Equal(t, err, errors.New("message is not found"), "")
}