package main

import (
	"context"

	"github.com/isaqueveras/juazeiro"

	"github.com/isaqueveras/juaz/example/golang/protos/biometry"
	"github.com/isaqueveras/juaz/example/golang/protos/user"
)

func main() {
	conn, err := juazeiro.NewClient("http://localhost:8181")
	if err != nil {
		handling(err)
		return
	}

	input := &user.User{
		Id:    pointer(int64(52)),
		Name:  pointer("Steve"),
		Level: pointer(user.LevelAdmin),
	}

	input.NewParams()
	input.WithParamLimit(pointer(89))
	input.WithParamOffset(pointer(0))
	input.WithParamTickets([]*int64{
		pointer(int64(123213)),
		pointer(int64(121235)),
		pointer(int64(768663)),
	})

	repoUser := user.NewUserClient(conn)
	data, err := repoUser.GetUser(context.Background(), input)
	if err != nil {
		handling(err)
		return
	}

	inputBio := &biometry.SearchParams{
		RequestId: data.Name,
	}

	repoBio := biometry.NewBiometryClient(conn)
	if _, err = repoBio.Search(context.Background(), inputBio); err != nil {
		handling(err)
		return
	}

	if _, err = repoUser.EditUser(context.Background(), input); err != nil {
		handling(err)
		return
	}
}
