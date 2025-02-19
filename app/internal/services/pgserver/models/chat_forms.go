package models

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

type ChatAdd struct {
    CreatorId   int     `json:"creatorId"`
    Members     []int   `json:"members"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
}

func (this *ChatAdd) GetData(c *gin.Context) error {
    data, err := c.GetRawData()
    if err != nil {
        return err 
    }
    if err = json.Unmarshal(data, this); err != nil {
        return err
    }
    if len(this.Name) == 0 {
        return errors.New("Can't create chat with empty name")
    }
    return nil
}

type EditChat struct {
    ChatId      int     `json:"id"`
    Name        *string `json:"name"`
    Description *string `json:"description"`
}

func (this *EditChat) GetData(c *gin.Context) error {
    data, err := c.GetRawData()
    if err != nil {
        return err 
    }
    if err = json.Unmarshal(data, this); err != nil {
        return err
    }
    if this.Name == nil && this.Description == nil {
        return errors.New("Wrong JSON form sent")
    }
    return nil
}

type MoveUserInChat struct {
    ChatId      int     `json:"chatId"`
    UserId      int     `json:"userId"`
}

func (this *MoveUserInChat) GetData(c *gin.Context) error {
    data, err := c.GetRawData()
    if err != nil {
        return err 
    }
    if err = json.Unmarshal(data, this); err != nil {
        return err
    }
    if this.ChatId == 0 && this.UserId == 0 {
        return errors.New("Wrong JSON form sent")
    }
    return nil
}

type UpdateUserInChat struct {
    ChatId      int     `json:"chatId"`
    UserId      int     `json:"userId"`
    Role        rune    `json:"role"`
}

func (this *UpdateUserInChat) GetData(c *gin.Context) error {
    data, err := c.GetRawData()
    if err != nil {
        return err 
    }
    if err = json.Unmarshal(data, this); err != nil {
        return err
    }
    if this.ChatId == 0 && this.UserId == 0 {
        return errors.New("Wrong JSON form sent")
    }
    if this.Role != 'c' && this.Role != 'm' && this.Role != 'a' {
        return errors.New("Wrong User role (possible roles is c, m or a)")
    }
    return nil
}

type User struct {
    Id          int     `json:"id"`
    Role        rune    `json:"role"`
}

type ChatInfo struct {
    ChatId      int     `json:"chatId"`
    CreatorId   int     `json:"creatorId"`
    Members     []User  `json:"members"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
}

func (this *ChatInfo) GetData(id int) {
    this.ChatId = id
}
