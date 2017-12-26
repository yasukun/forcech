package lib

import (
	"bytes"
	"encoding/gob"
	"time"
)

type Thread struct {
	ID       string
	Title    string
	TagId    string
	SubTagId string
	Create   time.Time
}

// GobEncode ...
func (t *Thread) GobEncode() ([]byte, error) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	if err := encoder.Encode(t.ID); err != nil {
		return nil, err
	}
	if err := encoder.Encode(t.Title); err != nil {
		return nil, err
	}
	if err := encoder.Encode(t.TagId); err != nil {
		return nil, err
	}
	if err := encoder.Encode(t.SubTagId); err != nil {
		return nil, err
	}  
	if err := encoder.Encode(t.Create); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

// (t *Thread) GobDecode ...
func (t *Thread) GobDecode(buf []byte) error {
	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)

	if err := decoder.Decode(&t.ID); err != nil {
		return err
	}
	if err := decoder.Decode(&t.Title); err != nil {
		return err
	}
	if err := decoder.Decode(&t.TagId); err != nil {
		return err
	}
	if err := decoder.Decode(&t.SubTagId); err != nil {
		return err
	}  
	if err := decoder.Decode(&t.Create); err != nil {
		return err
	}
	return nil
}
