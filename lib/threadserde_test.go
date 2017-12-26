package lib

import (
	"log"
	"testing"
	"time"
)

// TestThreadGob ...
func TestThreadGob(t *testing.T)  {
  thread1 := &Thread{ID: Uuid(), Title: "kemonofriend", TagId: Uuid(), SubTagId:  Uuid(), Create: time.Now()}
  b, err := thread1.GobEncode()
	if err != nil {
		log.Fatal("Error occured. %v", err)
	}
  str_b := string(b)
  log.Println(str_b)
  thread2 := &Thread{}
	if err := thread2.GobDecode([]byte(str_b)); err != nil {
		log.Println("Error occured %v", err)
	}
  log.Println(thread2.ID)
  log.Println(thread2.TagId)
  log.Println(thread2.SubTagId)
  log.Println(thread2.Title)
  log.Println(thread2.Create)
}
