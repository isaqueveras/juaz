package main

import (
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

	ctx.JSON(http.StatusOK, []*biometry.Biometry{
		{
			Id:       pointer(int64(1)),
			Document: pointer("122.534.634-56"),
			Flow:     pointer(biometry.FlowChangeOwnership),
			State:    pointer(biometry.StateDone),
		},
		{
			Id:       pointer(int64(2)),
			Document: pointer("237.848.239-54"),
			Flow:     pointer(biometry.FlowChangePlan),
			State:    pointer(biometry.StateCanceled),
		},
	})
}

func obtain_biometry(ctx *gin.Context) {
	body := new(biometry.Biometry)
	if err := ctx.BindJSON(&body); err != nil {
		panic(err)
	}

	body.State = pointer(biometry.StateUnderAnalysis)
	body.Id = pointer(int64(238942347238947923))

	ctx.JSON(http.StatusOK, body)
}

func create_biometry(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "4f68accd-a754-4218-91f6-b9e6c0fe8da7")
}
