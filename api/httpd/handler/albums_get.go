package handler

import (
	"net/http"
	"spc/api/platform/src/profile"

	"github.com/gin-gonic/gin"
)

func GetAlbums(collection *profile.Collection) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		results := collection.GetAll()
		ctx.IndentedJSON(http.StatusOK, results)
	}
}
