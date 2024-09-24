package main

import (
	"log"
	"os"

	"github.com/linguinichico/spc/api/httpd/handler"
	"github.com/linguinichico/spc/api/platform/src/profile"
	"github.com/linguinichico/spc/storage"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// albums slice to seed record album data.
// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func main() {

	// Fetch .env variables to go
	godotenv.Load()

	API_IP := os.Getenv("API_IP")
	API_PORT := os.Getenv("API_PORT")

	// this will be the place to the initial call to the db
	collection := profile.New()

	//connect to a Postgres database
	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	defer store.ClosePostgresStore()

	//initialize users table
	store.CreateGroupUserTable()

	router := gin.Default()
	group_api := router.Group("/api")
	{
		group_api.GET("/albums", handler.GetAlbums(collection))
		group_api.POST("/albums", handler.PostAlbum(collection))
		group_api.GET("/users", handler.GetUsers(store))
		group_api.POST("/users", handler.PostUser(store))
	}

	router.Run(API_IP + ":" + API_PORT)
}
