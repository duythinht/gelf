package chunk

import (
	"bytes"
	"crypto/rand"
	"math"
	"testing"
)

func TestSplitBySize(t *testing.T) {
	bs := make([]byte, 100)
	_, _ = rand.Read(bs)

	chunks := Split(bs, 24)
	if len(chunks) != 5 {
		t.Log("chunks len must be 5")
		t.Log(chunks)
		t.Fail()
	}
}

func TestIntToByte(t *testing.T) {
	i := 145
	var b byte = IntToByte(i)

	if b != 145 {
		t.Log(b)
		t.Fail()
	}
}

func TestMakeRandomID(t *testing.T) {
	ID := RandomID()
	if len(ID) != 8 {
		t.Log("Random ID should have length = 8")
		t.Fail()
	}
	ID1 := RandomID()
	if bytes.Equal(ID, ID1) {
		t.Log("New ID should be difference")
		t.Fail()
	}
}

var msg = `
{
  "version": "1.1",
  "host": "chuck...",
  "short_message": "A short message %d",
  "full_message": "Backtrace here\n\nmore stuff",
  "timestamp": %d,
  "level": 1,
  "_user_id": %d,
  "_some_info": "foo",
  "_some_env_var": "bar"
}
`

var magicGelf = []byte{30, 15}

func TestToChunkedMessage(t *testing.T) {
	buf := ZipMessage(msg)

	l := len(buf)
	chunksSize := int(math.Ceil(float64(l) / 10.0))
	chunks := GetGelfChunks(buf, 10)
	if len(chunks) != chunksSize {
		t.Log(chunksSize, len(chunks))
		t.Log(chunks)
		t.Fail()
	}

	for index, chunk := range chunks {
		if !bytes.Equal(magicGelf, chunk[:2]) {
			t.Log("Chunk must be start by magic gelf", magicGelf)
			t.Fail()
		}

		if !bytes.Equal([]byte{byte(index)}, chunk[10:11]) {
			t.Log(index, chunk[10:11])
			t.Log("Chunk must contain index")
			t.Fail()
		}

		if !bytes.Equal([]byte{byte(chunksSize)}, chunk[11:12]) {
			t.Log(chunksSize, chunk[10:11])
			t.Log("Chunk must contain length of chunks")
			t.Fail()
		}
	}
}
