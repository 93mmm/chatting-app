package models

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Chat struct {
    CreatorId   int     `json:"creatorId"`
    Members     []int   `json:"members"`
    Admins      []int   `json:"admins"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
}

func (this *Chat) GetData(c *gin.Context) error {
    data, err := c.GetRawData()
    if err != nil {
        return err 
    }
    if err = json.Unmarshal(data, this); err != nil {
        return err
    }
    fmt.Printf("this: %v\n", this)
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
    fmt.Printf("this: %v\n", this)
    return nil
}
