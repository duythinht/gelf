package chunk

import "testing"

func TestCompressByGzip(t *testing.T) {
	message := `{"message": "hello world", "level": 123, "levelString": "DEBUG", "hello world", "DEBUG"}`
	zipped := ZipMessage(message)
	if len(message) <= len(zipped) {
		t.Log(message)
		t.Log(zipped)
		t.Fail()
	}
}
