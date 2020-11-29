package main

// Don't go snooping around for answers here unless you want the puzzles spoilt.

import (
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)


func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "Page 1",
        })
    })

    port := os.Getenv("PORT") // Define this on Heroku.
    router.Run(":" + port)
}
