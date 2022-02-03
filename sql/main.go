package main

import (
	"fmt"
	"os"

	"github.com/Ensena/compiler"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func main() {

	ws := gin.Default()
	ws.Use(CORSMiddleware())
	ws.Use(apmgin.Middleware(ws))
	ws.POST("/compiler/sql", func(ctx *gin.Context) {
		c := compiler.Compiler{}
		c.Save = save
		c.Build = build
		c.Run(ctx)
	})
	ws.Run(":8001")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

const PATH = "/opt/sql"

func save(UserID, Name, Time, Content string) string {
	folder := fmt.Sprintf("%s/src/%s/%s/%s", PATH, UserID, Name, Time)
	err := os.MkdirAll(folder, os.ModePerm)

	fmt.Println("ERR", err)

	folderBuilder := fmt.Sprintf("%s/build/%s/%s", PATH, UserID, Name)

	os.MkdirAll(folderBuilder, os.ModePerm)

	f, err := os.Create(folder + "/run.sql")

	folder = fmt.Sprintf("%s/build/%s", PATH, UserID)
	os.MkdirAll(folder, os.ModePerm)

	fmt.Println("FOLDER", folder)

	folderBuilder = fmt.Sprintf("%s/%s/%s", UserID, Name, Time)

	if err != nil {
		fmt.Println(err)
		return folderBuilder
	}
	_, err = f.WriteString(Content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return folderBuilder
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return folderBuilder
	}
	return folderBuilder
}

func build(path string) (bool, string) {
	// cmd := exec.Command("sh", "-c", "cd "+PATH+"/src/"+path+"/ && GO111MODULE=off go build -o "+PATH+"/build/"+path+" main.go")
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	return true, fmt.Sprintf("%s %s", out, err)
	// }
	return false, ""
}
