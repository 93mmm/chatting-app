package models

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

type RegisterUser struct {
    Email           string  `json:"email"`
    PasswordHash    string  `json:"pwdHash"`
    Username        *string  `json:"username"`
}

func (this *RegisterUser) GetData(c *gin.Context) error {
    data, err := c.GetRawData()
    if err != nil {
        return err 
    }
    if err = json.Unmarshal(data, this); err != nil {
        return err
    }
    if this.Username == nil || len(*this.Username) == 0 {
        this.Username = nil
    }
    if len(this.PasswordHash) == 0 {
        return errors.New("Invalid PasswordHash")
    }
    if len(this.Email) == 0 {
        return errors.New("Invalid Email")
    }
    return nil
}

type EditUser struct {
    Id              int     `json:"id"`
    Username        string  `json:"username"`
    Email           string  `json:"email"`
    PasswordHash    string  `json:"pwdHash"`
}

func (this *EditUser) GetData(c *gin.Context) error {
    data, err := c.GetRawData()
    if err != nil {
        return err 
    }
    if err = json.Unmarshal(data, this); err != nil {
        return err
    }
    return nil
}
