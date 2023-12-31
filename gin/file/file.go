package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")
	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")
		// Source
		file, err := c.FormFile("file")

    if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
    }
    
    if name == "" {
			c.String(http.StatusBadRequest, "이름을 입력해주세요.")
			return
    }
    
    filename := filepath.Base(file.Filename)
    
    if err := c.SaveUploadedFile(file, filename);
	
	err != nil {
			c.String(http.StatusBadRequest, "파일 업로드에 실패했습니다.: %s", err.Error())
			return
    }

    c.String(http.StatusOK, "File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email)
  })
  
  router.Run(":8080")
}
