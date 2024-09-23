package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linguinichico/spc/storage"
)

func GetUsers(s *storage.PostgresStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		group_users, err := s.GetUsers()
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.IndentedJSON(http.StatusOK, group_users.Users)
	}
}
