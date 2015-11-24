package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

/*
 * Low level functions to write and read raw bytes to and from our TCP connection
 */

// To convert Big Endian binary format of a 4 byte integer to int32
func fromBytes32(b []byte) (uint32, error) {
	buf := bytes.NewReader(b)
	var result uint32
	err := binary.Read(buf, binary.BigEndian, &result)
	return result, err
}

// To convert an int32 to a 4 byte Big Endian binary format
func toBytes32(i uint32) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, i)
	return buf.Bytes(), err
}

// To convert Big Endian binary format of a 1 byte integer to int8
func fromBytes8(b []byte) (uint8, error) {
	buf := bytes.NewReader(b)
	var result uint8
	err := binary.Read(buf, binary.BigEndian, &result)
	return result, err
}

// To convert an int8 to a 1 byte Big Endian binary format
func toBytes8(i uint8) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, i)
	return buf.Bytes(), err
}

func writeRawMsg(conn net.Conn, dataType uint8, payload []byte) error {
	// Send the size of the message to be sent
	lenPayload, err := toBytes32(uint32(len(payload)))
	if err != nil {
		return err
	}
	_, err = conn.Write(lenPayload)
	if err != nil {
		return err
	}

	// Send the type of payload
	bDataType, err := toBytes8(dataType)
	if err != nil {
		return err
	}
	_, err = conn.Write(bDataType)
	if err != nil {
		return err
	}

	// Send the message
	_, err = conn.Write(payload)
	if err != nil {
		return err
	}
	return nil
}

func readRawMsg(conn net.Conn) (uint8, []byte, error) {
	// Make a buffer to hold length of data
	lenBuf := make([]byte, 4)
	_, err := conn.Read(lenBuf)
	if err != nil {
		return 0, nil, err
	}
	lenData, err := fromBytes32(lenBuf)
	if err != nil {
		return 0, nil, err
	}

	// Make a buffer to hold type of data (1byte)
	typeBuf := make([]byte, 1)
	_, err = conn.Read(typeBuf)
	if err != nil {
		return 0, nil, err
	}
	typeData, err := fromBytes8(typeBuf)
	if err != nil {
		return 0, nil, err
	}

	// Make a buffer to hold incoming data.
	buf := make([]byte, lenData)
	reqLen := 0
	// Keep reading data from the incoming connection into the buffer until all the data promised is
	// received
	for reqLen < int(lenData) {
		tempreqLen, err := conn.Read(buf[reqLen:])
		reqLen += tempreqLen
		if err != nil {
			if err == io.EOF {
				return 0, nil, fmt.Errorf("Received EOF before receiving all promised data.")
			}
			return 0, nil, fmt.Errorf("Error reading: %s", err.Error())
		}
	}
	return typeData, buf, nil
}
