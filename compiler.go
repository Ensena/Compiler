package compiler

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type data struct {
	File string
}

type Compiler struct {
	UserID string
	Name   string
	MSG    string
	Path   string
	Error  bool
	Save   func(string, string, string) string
	Build  func(string) (bool, string)
}

func (c *Compiler) Init(ctx *gin.Context) {
	c.UserID = "0"
	if ctx.Request.Header.Get("UserID") != "" {
		c.UserID = ctx.Request.Header.Get("UserID")
	}
	c.Name = "0"
	if ctx.Request.URL.Query().Get("name") != "" {
		c.Name = ctx.Request.URL.Query().Get("name")
	}
}

func (c *Compiler) Run(ctx *gin.Context) {
	var d data
	c.Init(ctx)
	reqBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return
	}
	json.Unmarshal(reqBody, &d)
	c.Path = c.Save(c.UserID, c.Name, d.File)
	c.Error, c.MSG = c.Build(c.Path)
	ctx.JSON(200, c)
}
