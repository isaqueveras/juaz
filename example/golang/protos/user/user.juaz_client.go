// versions: v1.0.1
// source: user.juaz

package user

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/isaqueveras/juazeiro"
)

// Level defines the level enum type
type Level string

const (
	LevelAdmin    Level = "ADMIN"
	LevelUser     Level = "USER"
	LevelEmployee Level = "EMPLOYEE"
)

// String convert level type to string
func (t Level) String() string {
	return string(t)
}

// Empty data model for the empty structure
type Empty struct {
}

// UserParams data model for the user_params structure
type UserParams struct {
	Limit     *int       `json:"limit,omitempty"`
	Offset    *int       `json:"offset,omitempty"`
	Total     *bool      `json:"total,omitempty"`
	Tickets   []*int64   `json:"tickets,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// User data model for the user structure
type User struct {
	Id         *int64  `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	Level      *Level  `json:"level,omitempty"`
	parameters *UserParams
}

func (u *User) NewParams() {
	u.parameters = &UserParams{}
}

// WithParamLimit ...
func (u *User) WithParamLimit(limit *int) {
	u.parameters.Limit = limit
}

// WithParamOffset ...
func (u *User) WithParamOffset(offset *int) {
	u.parameters.Offset = offset
}

// WithParamTotal ...
func (u *User) WithParamTotal(total *bool) {
	u.parameters.Total = total
}

// WithParamTickets ...
func (u *User) WithParamTickets(tickets []*int64) {
	u.parameters.Tickets = tickets
}

// WithParamCreatedAt ...
func (u *User) WithParamCreatedAt(createdAt *time.Time) {
	u.parameters.CreatedAt = createdAt
}

// IUserClient defines the interface of the provided methods
type IUserClient interface {
	GetUser(context.Context, *User) (*User, error)
	CreateUser(context.Context, *User) (*Empty, error)
	EditUser(context.Context, *User) (*Empty, error)
	DeleteUser(context.Context, *Empty) (*Empty, error)
}

type UserClient struct {
	cc juazeiro.ClientConnInterface
}

func NewUserClient(cc juazeiro.ClientConnInterface) IUserClient {
	return &UserClient{cc: cc}
}

// GetUser implements the GetUser method of the interface
func (c *UserClient) GetUser(ctx context.Context, in *User) (*User, error) {
	out := new(User)
	uri := fmt.Sprintf("/v1/user/%v", *in.Id)
	if in.parameters != nil {
		uri += _build_user_params_parameters(in.parameters)
		in.parameters = nil
	}
	err := c.cc.Invoke(ctx, http.MethodGet, uri, in, out)
	return out, err
}

// CreateUser implements the CreateUser method of the interface
func (c *UserClient) CreateUser(ctx context.Context, in *User) (*Empty, error) {
	out := new(Empty)
	uri := fmt.Sprintf("/v1/account/user/%v/create", *in.Id)
	if in.parameters != nil {
		uri += _build_user_params_parameters(in.parameters)
		in.parameters = nil
	}
	err := c.cc.Invoke(ctx, http.MethodPost, uri, in, out)
	return out, err
}

// EditUser implements the EditUser method of the interface
func (c *UserClient) EditUser(ctx context.Context, in *User) (*Empty, error) {
	out := new(Empty)
	uri := fmt.Sprintf("/v1/account/user/%v/edit", *in.Id)
	if in.parameters != nil {
		uri += _build_user_params_parameters(in.parameters)
		in.parameters = nil
	}
	err := c.cc.Invoke(ctx, http.MethodPost, uri, in, out)
	return out, err
}

// DeleteUser implements the DeleteUser method of the interface
func (c *UserClient) DeleteUser(ctx context.Context, in *Empty) (*Empty, error) {
	out := new(Empty)
	uri := "/v1/account/user/delete"
	err := c.cc.Invoke(ctx, http.MethodDelete, uri, in, out)
	return out, err
}

func _build_user_params_parameters(in *UserParams) string {
	val := &url.Values{}
	if in.Limit != nil {
		val.Add("limit", fmt.Sprintf("%v", *in.Limit))
	}
	if in.Offset != nil {
		val.Add("offset", fmt.Sprintf("%v", *in.Offset))
	}
	if in.Total != nil {
		val.Add("total", fmt.Sprintf("%v", *in.Total))
	}
	for _, value := range in.Tickets {
		if value == nil {
			continue
		}
		val.Add("tickets", fmt.Sprintf("%v", *value))
	}
	if in.CreatedAt != nil {
		val.Add("created_at", fmt.Sprintf("%v", *in.CreatedAt))
	}
	return "?" + val.Encode()
}
