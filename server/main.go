package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func errLog(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	exec.Command("mkfir", "file")
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		errLog(err)
		log.Println(file.Filename)
		err = c.SaveUploadedFile(file, "./file/"+file.Filename)
		errLog(err)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	file, err := os.Open("./file/*.go")
	errLog(err)
	exec.Command("go", "mod", "init")
	exec.Command("go", "build", file.Name(), "-o", strings.TrimSuffix(file.Name(), ".go"))
	router.PUT("/send", func(c *gin.Context) {

	})
	router.Run(":3000")
	exec.Command("rm", "-rf", "file")
}
