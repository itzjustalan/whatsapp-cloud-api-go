package whatsapp_cloud_api_go

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	w := New("9123", "abc", V13_0, HTTPS)
	log.Println(w)
}

func TestSend(t *testing.T) {
	w := New("9123", "abc", V13_0, HTTPS)
	err := w.SendTextMessage("9321", "jaba", true)
	log.Println("err:", err)
}

func TestAdd(t *testing.T) {
	s := Add(2, 3)
	a := 5

	if s != a {
		t.Errorf("got %q wanted %q", s, a)
	}
}
