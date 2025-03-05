package utils

import (
	//"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	//"github.com/gorilla/sessions"
)

func SaveSession(ctx *gin.Context, userId uint64) error {
	//session := sessions.Default(ctx)
	//session.Set("user_id", userId)
	//options := sessions.Options{
	//	Path:   "/",
	//	MaxAge: 86400 * 7,
	//	Secure: true,
	//}
	//session.Options(options)
	//return session.Save()
	return nil
}
