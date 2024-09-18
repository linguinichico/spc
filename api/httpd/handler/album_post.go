package handler

import (
	"fmt"
	"net/http"
	"spc/api/platform/src/profile"

	"github.com/gin-gonic/gin"
)

type AlbumPostRequest struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func PostAlbum(collection *profile.Collection) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestBody := AlbumPostRequest{}

		if err := ctx.BindJSON(&requestBody); err != nil {
			fmt.Println(err)
			return
		}

		album := profile.Album{
			ID:     requestBody.ID,
			Title:  requestBody.Title,
			Artist: requestBody.Artist,
			Price:  requestBody.Price,
		}

		collection.Add(album)
		ctx.IndentedJSON(http.StatusCreated, album)
	}
}
