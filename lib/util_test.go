package lib

import (
	"log"
	"testing"
)

// TestUuid ...
func TestUuid(t *testing.T) {
	uuid := Uuid()
	log.Println(uuid)
}
