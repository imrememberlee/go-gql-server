package server

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/cmelgarejo/go-gql-server/internal/handlers"
)

var HOST, PORT string

func init() {
    HOST = "localhost"
    PORT = "4000"
}

// Run web server
func Run() {
    r := gin.Default()
    // Setup routes
    r.GET("/ping", handlers.Ping())
    log.Println("Running @ http://" + HOST + ":" + PORT )
    log.Fatalln(r.Run(HOST + ":" + PORT))
}