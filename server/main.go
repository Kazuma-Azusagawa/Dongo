package main

// #include <stdlib.h>
// #include <string.h>
import "C"
import (
	//	"fmt"
	"log"
	"net/http"
	//	"os"
	//	"github.com/achille-roussel/go-ffi"
	"github.com/gin-gonic/gin"
	"os/exec"
)

func errLog(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
func main() {
	exec.Command("mkdir", "file")
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.LoadHTMLGlob("index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.POST("/upload", func(c *gin.Context) {
		savedFile, err := c.FormFile("file")
		errLog(err)
		log.Println(savedFile.Filename)
		err = c.SaveUploadedFile(savedFile, "./file/"+savedFile.Filename)
		errLog(err)
		C.system(C.CString("cd file && go mod init main"))
		C.system(C.CString("cd file && go build -o out"))
	})
	router.GET("/file/main", func(c *gin.Context) {
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename=out.exe")
		c.File("./file/out")
	})
	router.Run(":3000")
}
