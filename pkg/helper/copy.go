package helper

import (
	"bytes"
	"encoding/gob"
)

// The struct must be Pointer and the fields must be identical
func CopyStruct(dst, src interface{}) error {
	buf := new(bytes.Buffer)
	err := gob.NewEncoder(buf).Encode(src)
	if err != nil {
		return err
	}

	if err := gob.NewDecoder(buf).Decode(dst); err != nil {
		return err
	}

	return nil
}
