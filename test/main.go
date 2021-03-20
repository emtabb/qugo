package main

import "github.com/emtabb/qugo"
import "log"

func main() {
	var qu = qugo.Operator().Init(nil).Skip(10).Limit(5)
	log.Println(qu)
}