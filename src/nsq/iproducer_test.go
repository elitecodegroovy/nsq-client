package nsq

import (
	"testing"

)


func TestSimpleProducer1(t *testing.T){
	//init default config
	config := NewConfig()

	w, _ := NewProducer("10.50.115.16:4150", config)
	w.SetLogger(nullLogger, LogLevelInfo)

	err := w.Publish("req-test", []byte("test"))
	if err != nil {
		panic("should lazily connect - %s")
	}

	w.Stop()

	err = w.Publish("req_test", []byte("fail test"))
}
