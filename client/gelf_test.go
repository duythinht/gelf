package client

import "testing"

func TestDefaultConfig(t *testing.T) {
	c := New(Config{})
	if c.config.GraylogHost != "localhost" {
		t.Log("Config must be default")
		t.Log(c.config)
		t.Fail()
	}
	if c.config.GraylogPort != 12201 {
		t.Log("Config must be default")
		t.Log(c.config)
		t.Fail()
	}

	if c.config.Connection != "lan" {
		t.Log("Config must be default")
		t.Log(c.config)
		t.Fail()
	}
	if c.config.MaxChunkSize != 8196 {
		t.Log("Config must be default")
		t.Log(c.config)
		t.Fail()
	}
}

// Should mock a udp server for test Send
func TestSendMessage(t *testing.T) {

}
