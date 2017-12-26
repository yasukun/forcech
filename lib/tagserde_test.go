package lib

import (
	"log"
	"testing"
	"time"
)

// TestTagGob ...
func TestTagGob(t *testing.T) {
	tag1 := &Tag{ID: Uuid(), Title: "kenmo", Create: time.Now()}
	b, err := tag1.GobEncode()
	if err != nil {
		log.Fatal("Error occured. %v", err)
	}
	str_b := string(b)
	log.Println(str_b)
	tag2 := &Tag{}
	if err := tag2.GobDecode([]byte(str_b)); err != nil {
		log.Println("Error occured %v", err)
	}
	log.Println(tag2.ID)
	log.Println(tag2.Title)
	log.Println(tag2.Create)

}
