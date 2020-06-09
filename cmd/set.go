package cmd

import (
	"github.com/adcs-io/adcs/storage"
	"time"
)

func (this *CMD) Set(k string, v string, expire ...time.Duration) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	var e time.Duration

	if expire != nil && len(expire) > 0 {
		e = expire[0]
	} else {
		e = time.Duration(0)
	}

	storage.DB[this.DB][k] = &storage.Data{
		DataType: 1,
		V:        v,
		Expire:   e,
	}

	return true
}
