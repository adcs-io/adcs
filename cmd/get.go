package cmd

import "github.com/adcs-io/adcs/storage"

func (this *CMD) Get(k string) string {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if data := storage.DB[this.DB][k]; data == nil {
		return ""
	} else {
		return data.V.(string)
	}
}
