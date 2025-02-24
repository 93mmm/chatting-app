package models

import (
	"encoding/json"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/helpers"
	se "github.com/93mmm/chatting-app/app/pkg/server_errors"
	"github.com/gin-gonic/gin"
)

type ChatCreate struct {
    CreatorId   int     `json:"creatorId"`
    Members     []int   `json:"members"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
}

func (this *ChatCreate) GetData(c *gin.Context) error {
    data, err := c.GetRawData()
    if err != nil {
        return se.UndefinedError(err)
    }
    if err = json.Unmarshal(data, &this); err != nil {
        return se.JsonCorruptedError()
    }
    this.Members = helpers.GetDistinct(this.Members)
    if this.CreatorId == 0 {
        return se.WrongKVError("creatorId", this.CreatorId)
    }
    if len(this.Name) == 0 {
        return se.WrongKVError("name", this.Name)
    }
    return nil
}

// type MoveUserInChat struct {
//     ChatId      int     `json:"chatId"`
//     UserId      int     `json:"userId"`
// }

// type UpdateUserInChat struct {
//     ChatId      int     `json:"chatId"`
//     UserId      int     `json:"userId"`
//     Role        rune    `json:"role"`
// }

// type User struct {
//     Id          int     `json:"id"`
//     Role        rune    `json:"role"`
// }

type ChatInfo struct {
    ChatId      int     `json:"chatId"`
    CreatorId   int     `json:"creatorId"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
}

// func (this ChatInfo) GetData(c *gin.Context) error {
//     data, err := c.GetRawData()
//     if err != nil {
//         return se.UndefinedError(err)
//     }
    // err = db.GetChatInfo(&this)
//
//     return nil
// }

// type ChatMembersInfo struct {
//     ChatId      int     `json:"chatId"`
//     Members     []int  `json:"members"`
// }
