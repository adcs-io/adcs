package server

import (
	"github.com/adcs-io/adcs/cmd"
	"github.com/adcs-io/adcs/storage"
	"strconv"
	"testing"
)

func TestRun(t *testing.T) {
	s := &server{}
	s.Run()

}

func BenchmarkSet(b *testing.B) {
	storage.Init()
	c := cmd.CMD{
		DB: 0,
	}

	for i := 0; i < b.N; i++ {
		c.Set(strconv.Itoa(i), "a")
		//c.Get("a")
	}
}

func BenchmarkGet(b *testing.B) {
	storage.Init()
	c := cmd.CMD{
		DB: 0,
	}
	for i := 0; i < b.N; i++ {
		c.Get(strconv.Itoa(i))
	}
}
