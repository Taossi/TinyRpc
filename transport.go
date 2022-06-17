package tinyrpc

import (
	"encoding/binary"
	"io"
	"log"
	"net"
)

type Transport struct {
	conn net.Conn
}

func NewTransport(conn net.Conn) *Transport {
	return &Transport{conn}
}



func (t *Transport) Send(data Data) error {
	bytes, err := encode(data)
	if err != nil {
		log.Println("send error", err)
		return err
	}
	buf := make([]byte, len(bytes) + 4)
	binary.BigEndian.PutUint32(buf[:4], uint32(len(bytes)))  // Header: first 4 bytes, value is the length of data
	copy(buf[4:],bytes)                                      // Body: actual data

	_, err = t.conn.Write(buf)
	return err
}

func (t *Transport) Receive() (Data, error) {
	header := make([]byte, 4)    // first 4 bytes is header
	_, err := io.ReadFull(t.conn, header)
	if err != nil {
		return Data{}, err
	}

	len := binary.BigEndian.Uint32(header)    // read length of data
	bytes := make([]byte, len)
	_, err = io.ReadFull(t.conn, bytes)
	if err != nil {
		return Data{}, err
	}

	data, err := decode(bytes)
	return data, err
}