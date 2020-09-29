package ip

import (
	"fmt"

	"github.com/go-routeros/routeros"
)

type Ip struct {
	client *routeros.Client
	Route
}

func (receiver *Ip) Print() {
	fmt.Println("print ip")
}

func (receiver *Ip) Set() {

}

func (receiver *Ip) Remove() {
}

func (receiver *Ip) Add() {

}

func (receiver *Ip) Find() {

}

func (receiver *Ip) Get() {

}

func (receiver *Ip) Export() {

}

func (receiver *Ip) Enable() {

}

func (receiver *Ip) Disable() {

}

func (receiver *Ip) Comment() {}

func (receiver *Ip) Move() {

}
