package storage

import "time"

var DB = make(map[int8]map[string]*Data)



type Data struct {
	DataType int8 //1:string, 2:map
	V        interface{}
	Expire   time.Duration //0:never
}

func Init()  {
	var i int8
	for i = 0; i < 16; i++ {
		DB[i] = make(map[string]*Data)
	}
}
