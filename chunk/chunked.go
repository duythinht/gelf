package chunk

import "crypto/rand"

//Split create chunks array by split messages buffer
func Split(a []byte, size int) [][]byte {
	l := len(a)
	chunks := [][]byte{}
	for i := 0; i < l; i += size {
		if i+size < l {
			chunks = append(chunks, a[i:i+size])
		} else {
			chunks = append(chunks, a[i:])
		}
	}
	return chunks
}

//IntToByte convert int to byte
func IntToByte(i int) byte {
	return byte(i % 256)
}

func RandomID() []byte {
	ID := make([]byte, 8)
	_, _ = rand.Read(ID)
	return ID
}

func GetGelfChunks(buffer []byte, size int) [][]byte {
	chunks := make([][]byte, 0)
	bufs := Split(buffer, size)
	ID := RandomID()
	chunksLength := len(bufs)
	for index, buf := range bufs {
		chunk := []byte{0x1e, 0x0f}
		chunk = append(chunk, ID...)
		chunk = append(chunk, IntToByte(index))
		chunk = append(chunk, IntToByte(chunksLength))
		chunk = append(chunk, buf...)
		chunks = append(chunks, chunk)
	}
	return chunks
}
