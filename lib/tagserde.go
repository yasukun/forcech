package lib

import (
	"bytes"
	"encoding/gob"
	"time"
)

type Tag struct {
	ID     string
	Title  string
	Main   bool
	Create time.Time
}

// GobEncode ...
func (t *Tag) GobEncode() ([]byte, error) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	if err := encoder.Encode(t.ID); err != nil {
		return nil, err
	}
	if err := encoder.Encode(t.Title); err != nil {
		return nil, err
	}
	if err := encoder.Encode(t.Main); err != nil {
		return nil, err
	}  
	if err := encoder.Encode(t.Create); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

// GobDecode ...
func (t *Tag) GobDecode(buf []byte) error {
	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)

	if err := decoder.Decode(&t.ID); err != nil {
		return err
	}
	if err := decoder.Decode(&t.Title); err != nil {
		return err
	}
	if err := decoder.Decode(&t.Main); err != nil {
		return err
	}  
	if err := decoder.Decode(&t.Create); err != nil {
		return err
	}
	return nil
}
