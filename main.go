package main

import (
	"fmt"
	"math/rand"
	"time"
	"net/http"
	"path/filepath"
	"github.com/gin-gonic/gin"
)
func main() {
	router := gin.Default()
	//router.Static("/", "./paste")
	router.POST("/", func(c *gin.Context) {

		// Source
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		filext := filepath.Ext(file.Filename)
		rand.Seed(time.Now().UnixNano())
		filename := RandomString(10)+filepath.Ext(file.Filename)
		fmt.Println(filename)
		fmt.Println(filext)

		if err := c.SaveUploadedFile(file, "paste/"+filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("http://wfnintr.net/paste/%s\ngopher://wfnintr.net/0/paste/%s\n", filename, filename))
	})
	router.Run(":1337")
}

func RandomString(n int) string {
    var letters = []rune("VHVybk9uVHVuZUluRHJvcE91dAo=")
    s := make([]rune, n)
    for i := range s {
        s[i] = letters[rand.Intn(len(letters))]
    }
    return string(s)
}
