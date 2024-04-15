package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/isaqueveras/juazeiro"

	"github.com/isaqueveras/juaz/example/golang/protos/biometry"
	"github.com/isaqueveras/juaz/example/golang/protos/user"
)

func main() {
	conn, _ := juazeiro.NewClient("http://localhost:8181")

	input := &user.User{
		Id:    pointer(int64(100)),
		Name:  pointer("John"),
		Level: pointer(user.LevelEmployee),
	}

	input.NewParams()
	input.WithParamLimit(pointer(10))
	input.WithParamOffset(pointer(0))
	input.WithParamTickets([]*int64{
		pointer(int64(123213)),
		pointer(int64(121235)),
		pointer(int64(768663)),
	})

	repoUser := user.NewUserClient(conn)
	data, err := repoUser.GetUser(context.Background(), input)
	if err != nil {
		log.Println(err)
		return
	}

	v, _ := json.Marshal(&data)
	log.Println(string(v))

	repoBio := biometry.NewBiometryClient(conn)
	bio, _ := repoBio.Search(context.Background(), &biometry.SearchParams{
		RequestId: pointer("09328932yg32432"),
	})

	biovalue, _ := json.Marshal(&bio)
	log.Println(string(biovalue))
}

// pointer returns a pointer reference
func pointer[T any](value T) *T {
	return &value
}
