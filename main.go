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

    router.GET("/puzzle/1", func(c *gin.Context) {
        c.HTML(http.StatusOK, "1.tmpl", gin.H{})
    })

    router.GET("/puzzle/2", func(c *gin.Context) {
        c.HTML(http.StatusOK, "2.tmpl", gin.H{})
    })

    router.GET("/puzzle/tastes", func(c *gin.Context) {
        c.HTML(http.StatusOK, "tastes.tmpl", gin.H{})
    })

    port := os.Getenv("PORT") // Define this on Heroku.
    router.Run(":" + port)
}

/* TABLE OF CONTENTS
   (so we remember the order that puzzles come in)

   PUZZLE #, FILENAME
   1        1
   2        2
   3        tastes
   tastes   ...
*/
