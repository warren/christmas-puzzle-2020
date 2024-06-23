package main

// Don't go snooping around here for answers! Puzzle spoilers lie below.

import (
    "net/http"
    "os"
    "strings"

    "github.com/gin-gonic/gin"
)


func main() {
    router := gin.Default()
    router.Static("/assets", "./assets")
    router.LoadHTMLGlob("templates/*")

    router.NoRoute(func(c *gin.Context) {
        relativePath := string(c.Request.URL.String())
        pathElems := strings.Split(relativePath, "/")

        if pathElems[1] == "puzzle" {
            if len(pathElems) >= 3 {
                // Serve the "wrong answer" page with a hint (if applicable).
                c.HTML(http.StatusOK, "wrong-answer.tmpl", gin.H{
                    "wronganswer": pathElems[2],
                    "hint": getHint(pathElems[2]),
                })
            } else {
                // Serve the "wrong answer" page with the empty string.
                c.HTML(http.StatusOK, "wrong-answer.tmpl", gin.H{
                    "wronganswer": "",
                })
            }
        }

        // Serve a normal 404 page.
        // TODO: Make a not-ugly 404 page.
    })

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

    router.GET("/puzzle/aftertaste", func(c *gin.Context) {
        c.HTML(http.StatusOK, "aftertaste.tmpl", gin.H{})
    })

    router.GET("/puzzle/phone", func(c *gin.Context) {
        c.HTML(http.StatusOK, "phone.tmpl", gin.H{})
    })
    router.GET("/puzzle/smartphone", func(c *gin.Context) {
        c.HTML(http.StatusOK, "phone.tmpl", gin.H{})
    })

    router.GET("/puzzle/gifts", func(c *gin.Context) {
        c.HTML(http.StatusOK, "gifts.tmpl", gin.H{})
    })

    router.GET("/puzzle/rehired", func(c *gin.Context) {
        c.HTML(http.StatusOK, "rehired.tmpl", gin.H{})
    })

    router.GET("/puzzle/square", func(c *gin.Context) {
        c.HTML(http.StatusOK, "square.tmpl", gin.H{})
    })

    router.GET("/puzzle/spicy", func(c *gin.Context) {
        c.HTML(http.StatusOK, "spicy.tmpl", gin.H{})
    })

    port := os.Getenv("PORT") // Define this on Heroku... still necessary?
    router.Run(":" + port)
}


// Text to add to a 404 page when the URL is equal to /puzzle/${DICT_KEY}.
var hints = map[string]string{
    "spaghetti":    "Hello! I am a hint.",
    "3":            "It's not going to be that easy.",
    "after":        "That's the first half.",
    "taste":        "That's the second half.",
    "hot":          "You're so close. What's the flavor called?",
    "acleverscreeninyourpocket": "You're close. What is a clever screen in your pocket?",
}
func getHint(wronganswer string) string {
    if val, ok := hints[wronganswer]; ok {
        return "** " + val + " **"
    }
    return ""
}


/* TABLE OF CONTENTS
   (so we remember which order the puzzles come in)

   PUZZLE #, FILENAME
   1            1
   2            2
   3            aftertaste
   aftertaste   phone (OR: smartphone)
   phone        gifts
   gifts        rehired
   rehired      square
   square       spicy
*/
