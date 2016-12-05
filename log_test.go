package gelf

import "testing"

func TestCreateMessage(t *testing.T) {
	m := Create("hello")
	if m.Version != "1.1" {
		t.Log("Wrong version")
		t.Fail()
	}
	if m.Host != "default" {
		t.Log("Wrong default host")
		t.Fail()
	}

	if m.ShortMessage != "hello" {
		t.Log("Wrong message")
		t.Fail()
	}
	if m.Timestamp == 0 {
		t.Log("Missing Timestamp")
		t.Fail()
	}
	if m.Level != 1 {
		t.Log("Wrong level")
		t.Fail()
	}
}

func TestWithTimestamp(t *testing.T) {
	m := Create("OK")
	m.SetTimestamp(123)
	if m.Timestamp != 123 {
		t.Log("Wrong timestamp after set")
		t.Fail()
	}
}

func TestWithHost(t *testing.T) {
	m := Create("OK")
	m.SetHost("khongbiet")
	if m.Host != "khongbiet" {
		t.Log("Wrong host after set")
		t.Fail()
	}
}

func TestWithFullMessage(t *testing.T) {
	m := Create("OK")
	m.SetFullMessage("khongbiet")
	if m.FullMessage != "khongbiet" {
		t.Log("Wrong full message after set")
		t.Fail()
	}
}

func TestWithLevel(t *testing.T) {
	m := Create("OK")
	m.SetLevel(3)
	if m.Level != 3 {
		t.Log("Wrong level after set")
		t.Fail()
	}
}

func TestToJSON(t *testing.T) {
	expect := `{"version":"1.1","host":"default","short_message":"OK","full_message":"","timestamp":145,"level":1}`
	actual := Create("OK").SetTimestamp(145).ToJSON()
	if expect != actual {
		t.Log(expect)
		t.Log(actual)
		t.Fail()
	}
}
