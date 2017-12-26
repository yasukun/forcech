package lib

import (
	// "encoding/gob"
  "github.com/twinj/uuid"
	"fmt"
)

// type Tag struct {
  
// }

// Uuid ...
func Uuid( ) string {
  u4 := uuid.NewV4()
  return fmt.Sprintf("%s", u4)
}
