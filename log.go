package gelf

import (
	"encoding/json"
	"os"
	"time"
)

// Log struct
/*
{
	"version": "1.1",
	"host": "example.org",
	"short_message": "A short message that helps you identify what is going on",
	"full_message": "Backtrace here\n\nmore stuff",
	"timestamp": 1385053862.3072,
	"level": 1,
	"_user_id": 9001,
	"_some_info": "foo",
	"_some_env_var": "bar"
}
*/

type Log struct {
	Version      string `json:"version"`
	Host         string `json:"host"`
	ShortMessage string `json:"short_message"`
	FullMessage  string `json:"full_message"`
	Timestamp    int64  `json:"timestamp"`
	Level        int    `json:"level"`
}

func Create(message string) *Log {
	return &Log{
		Version:      "1.1",
		Host:         "default",
		ShortMessage: message,
		Timestamp:    time.Now().Unix(),
		Level:        1,
	}
}

func (self *Log) SetTimestamp(timestamp int64) *Log {
	self.Timestamp = timestamp
	return self
}

func (self *Log) SetHost(host string) *Log {
	self.Host = host
	return self
}

func (self *Log) SetFullMessage(fullMessage string) *Log {
	self.FullMessage = fullMessage
	return self
}

func (self *Log) SetLevel(level int) *Log {
	self.Level = level
	return self
}

func (self *Log) ToJSON() string {
	message, err := json.Marshal(self)
	if err != nil {
		os.Exit(1)
	}
	return string(message)
}
