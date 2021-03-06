package tinyrpc

import (
	"bytes"
	"encoding/gob"
)

// define data format transported from server and client
type Data struct {
	Name string        // service name
    Args []interface{} // request's or response's body except error
    Err  string        // remote server error
}

func encode(data Data) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func decode(b []byte) (Data, error) {
	buf := bytes.NewBuffer(b)
	decoder := gob.NewDecoder(buf)
	var data Data
	err := decoder.Decode(&data)
	if err != nil {
		return Data{}, err
	}
	return data, nil
}