package main

// #include <stdlib.h>
// #include <string.h>
import "C"
import (
	"fmt"
	"log"
	"net/http"
	//	"os"
	//	"github.com/achille-roussel/go-ffi"
	"github.com/gin-gonic/gin"
	"os/exec"
	"strings"
)

func errLog(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
func RunCMD(path string, args []string, debug bool) (out string, err error) {

	cmd := exec.Command(path, args...)

	var b []byte
	b, err = cmd.CombinedOutput()
	out = string(b)

	if debug {
		fmt.Println(strings.Join(cmd.Args[:], " "))

		if err != nil {
			fmt.Println("RunCMD ERROR")
			fmt.Println(out)
		}
	}

	return
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
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", savedFile.Filename))
		//		xfile, err := os.Open("./file/" + savedFile.Filename)
		//		errLog(err)
		C.system(C.CString("cd file && go mod init main"))
		C.system(C.CString("cd file && go build " + savedFile.Filename))

	})
	//router.PUT("/send", func(c *gin.Context) {})
	router.Run(":3000")

}
