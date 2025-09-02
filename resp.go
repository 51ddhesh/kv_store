// resp.go: Implements basic RESP (REdis Serialization Protocol) parsing utilities
package main

import (
	"bufio"
	// "fmt"
	"io"
	"strconv"
)

// RESP type constants (first byte of each message)
const (
	STRING  = '+' // Simple String
	ERROR   = '-' // Error
	INTEGER = ':' // Integer
	BULK    = '$' // Bulk String
	ARRAY   = '*' // Array
)

// Value represents a parsed RESP value
type Value struct {
	typ   string  // Type of RESP value
	str   string  // For simple strings
	num   int     // For integers
	bulk  string  // For bulk strings
	array []Value // For arrays
}

// Resp wraps a bufio.Reader for parsing RESP messages
type Resp struct {
	reader *bufio.Reader
}

// NewResp creates a new Resp parser from an io.Reader
func NewResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

// readLine reads a line ending with CRLF from the RESP stream
func (r *Resp) readLine() (line []byte, n int, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		n += 1
		line = append(line, b)
		// Check for CRLF ending
		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}
	// Return line without CRLF
	return line[:len(line)-2], n, nil
}

// readInteger reads an integer value from the RESP stream
func (r *Resp) readInteger() (x int, n int, err error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, 0, err
	}
	// Parse integer from line
	i64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, n, err
	}
	return int(i64), n, nil
}
