package compiler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/Ensena/core/env-global"
	"github.com/elmalba/oauth2-server/jwt"
	"github.com/gin-gonic/gin"
)

var key string

func init() {
	key = env.Check("secretKey", "Missing Params secretKey")
}

type data struct {
	File string
	Name string
}

type Compiler struct {
	MSG   string
	Path  string
	Error bool
	Save  func(string, string, string, string) string `json:"-"`
	Build func(string) (bool, string)                 `json:"-"`
}

func (c *Compiler) Run(ctx *gin.Context) {
	var d data
	UserID := "0"
	UserID = ctx.Request.Header.Get("UserID")

	JWT := ctx.Request.Header.Get("Authorization")
	user, err := jwt.Decode(JWT, key)
	if err != nil {
		ctx.AbortWithStatus(403)
		return
	}
	UserID = user.ID
	reqBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return
	}
	json.Unmarshal(reqBody, &d)
	if d.Name == "" {
		d.Name = "default"
	}
	t := fmt.Sprintf("%d", (time.Now().Unix()))
	c.Path = c.Save(UserID, d.Name, t, d.File)
	c.Error, c.MSG = c.Build(c.Path)
	ctx.JSON(200, c)
}
