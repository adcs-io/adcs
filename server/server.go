package server

import (
	"github.com/adcs-io/adcs/cmd"
	"github.com/adcs-io/adcs/storage"
)

type server struct {
}

func (this *server) Run() {

	storage.Init()
	c := cmd.CMD{
		DB: 0,
	}

	c.Set("a", "a")
	c.Get("a")
	//fmt.Println(get)
	//select {
	//
	//}
}
