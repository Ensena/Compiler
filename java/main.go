package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Ensena/compiler"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func main() {

	ws := gin.Default()
	ws.Use(apmgin.Middleware(ws))
	ws.GET("/api/v1/compiler/java", func(ctx *gin.Context) {
		c := compiler.Compiler{}
		c.Save = save
		c.Build = build
		c.Run(ctx)
	})
	ws.Run(":8000")
}

const PATH = "/files/java"

func save(UserID, Name, File string) string {
	folder := fmt.Sprintf(""+PATH+"/src/%s", UserID)
	os.MkdirAll(folder, os.ModePerm)

	folder = fmt.Sprintf("%s/%s", folder, UserID)
	os.MkdirAll(folder, os.ModePerm)
	f, err := os.Create(folder + "/main.java")
	if err != nil {
		fmt.Println(err)
		return folder
	}

	_, err = f.WriteString(File)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return folder
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return folder
	}
	return folder

}
func build(path string) (bool, string) {

	cmd := exec.Command("sh", "-c", "javac -d /files/java/build/"+path+"/ /files/java/src/"+path+"/main.java")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return true, fmt.Sprintf("%s %s", out, err)
	}
	return false, string(out)
}
