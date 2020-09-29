package inter

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-routeros/routeros"
)

type Inter struct {
	Client *routeros.Client
}

func (receiver *Inter) Print() {
	reply, err := receiver.Client.Run("/interface/print", "?disabled=false", "?running=true", "=.proplist=name,rx-byte,tx-byte,rx-packet,tx-packet")
	if err != nil {
		log.Fatal(err)
	}

	for _, re := range reply.Re {
		for _, p := range strings.Split("name,rx-byte,tx-byte,rx-packet,tx-packet", ",") {
			fmt.Print(re.Map[p], "\t")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func (receiver *Inter) Set() {

}

func (receiver *Inter) Remove() {
}

func (receiver *Inter) Add() {

}

func (receiver *Inter) Find() {

}

func (receiver *Inter) Get() {

}

func (receiver *Inter) Export() {

}

func (receiver *Inter) Enable() {

}

func (receiver *Inter) Disable() {

}

func (receiver *Inter) Comment() {}

func (receiver *Inter) Move() {

}
