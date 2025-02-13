package models

import (
	"encoding/json"

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
    return nil
}
