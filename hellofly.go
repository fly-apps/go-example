package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/*name", handleIndex)
	r.Run()
}

// Hello is page data for the template
type Hello struct {
	Name string
}

func handleIndex(c *gin.Context) {
	name := strings.TrimPrefix(c.Param("name"), "/")
	c.HTML(http.StatusOK, "hellofly.tmpl", gin.H{"Name": name})
}
