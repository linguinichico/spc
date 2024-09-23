package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	t "github.com/linguinichico/spc/api/types"
	"github.com/linguinichico/spc/storage"
)

func PostUser(s *storage.PostgresStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestBody := &t.UserPostRequest{}
		if err := ctx.BindJSON(&requestBody); err != nil {
			fmt.Println(err)
			return
		}
		user := t.User{
			Username:  requestBody.Username,
			Password:  requestBody.Password,
			CreatedAt: time.Now().UTC(),
		}
		err := s.CreateUser(&user)
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.IndentedJSON(http.StatusCreated, user)
	}
}
