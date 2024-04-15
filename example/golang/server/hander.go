package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/juaz/example/golang/protos/biometry"
	"github.com/isaqueveras/juaz/example/golang/protos/user"
)

func getUser(ctx *gin.Context) {
	info := new(user.User)
	if err := ctx.BindJSON(&info); err != nil {
		panic(err)
	}

	v, _ := json.Marshal(&info)
	log.Println("body", string(v))

	out := &user.User{
		Id:    pointer(int64(213121)),
		Name:  pointer("Isaque Veras"),
		Level: pointer(user.LevelAdmin),
	}

	ctx.JSON(http.StatusOK, out)
}

func pointer[T any](value T) *T {
	return &value
}

func search_biometry(ctx *gin.Context) {
	body := new(biometry.SearchParams)
	if err := ctx.BindJSON(&body); err != nil {
		panic(err)
	}

	v, _ := json.Marshal(&body)
	log.Println("body", string(v))

	ctx.JSON(http.StatusOK, &biometry.Biometry{
		Id:       pointer(int64(000)),
		Document: pointer("0000000000"),
		Flow:     pointer(biometry.FlowChangeOwnership),
		State:    pointer(biometry.StateDone),
	})
}
