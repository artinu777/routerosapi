package routeosapi

import (
	"errors"

	"github.com/go-routeros/routeros"

	inter "github.com/artinu777/routerosapi/interface"
	"github.com/artinu777/routerosapi/ip"
)

type Router struct {
	routeAuth *RouterAuth
	client    *routeros.Client
}

type RouterAuth struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func NewRouter() *Router {
	return &Router{}
}

func (receiver *Router) SetAuth(auth RouterAuth) (*Router, error) {
	if auth.User == "" && auth.Password == "" {
		return receiver, errors.New("auth empty")
	}

	receiver.routeAuth = &auth
	return receiver, nil
}

func (receiver *Router) Connect(address string) (*Router, error) {
	if receiver.routeAuth == nil {
		return nil, errors.New("nil auth")
	}

	if address == "" {
		return nil, errors.New("address empty")
	}

	c, err := routeros.Dial(address, receiver.routeAuth.User, receiver.routeAuth.Password)
	if err != nil {
		return nil, err
	}

	receiver.client = c
	return receiver, nil
}

func (receiver *Router) Ip() *ip.Ip {
	return &ip.Ip{client: receiver.client}
}

func (receiver *Router) Inter() *inter.Inter {
	return &inter.Inter{client: receiver.client}
}

func Example() {

}
