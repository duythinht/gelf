package chunk

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

// ZipMessage from string to compress bytes
func ZipMessage(message string) []byte {

	var buffer bytes.Buffer
	writer := gzip.NewWriter(&buffer)
	_, err := writer.Write([]byte(message))

	if err != nil {
		fmt.Println(err)
	}

	_ = writer.Close()
	return buffer.Bytes()
}
