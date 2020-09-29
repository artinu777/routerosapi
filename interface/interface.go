package inter


import (
"flag"
"fmt"
"log"
"strings"

"github.com/go-routeros/routeros"
)

type Inter struct {
	Client *routeros.Client
}

func (receiver *Inter) Print() {
	properties := flag.String("properties", "name,rx-byte,tx-byte,rx-packet,tx-packet", "Properties")
	reply, err := receiver.Client.Run("/interface/print", "?disabled=false", "?running=true", "=.proplist="+*properties)
	if err != nil {
		log.Fatal(err)
	}

	for _, re := range reply.Re {
		for _, p := range strings.Split(*properties, ",") {
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
