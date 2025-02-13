package user

import (
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/db"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/helpers"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
	"github.com/gin-gonic/gin"
)

func Reg(c *gin.Context) {
    var reqUsr models.RegisterUser
    
    if err := reqUsr.GetData(c); err != nil {
        helpers.WrapAndSendError(c, err)
        return 
    }

    id, err := db.RegisterUser(reqUsr)
    if err != nil {
        helpers.WrapAndSendError(c, err)
        return
    }

    helpers.WrapAndSendId(c, id)
}
