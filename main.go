package main

import (
	"fmt"
	"os"
	"time"

	"github.com/duythinht/gelf/client"
)

func errOrNot(err error) {
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
}

var msgTpl = `
{
  "version": "1.1",
  "host": "chuck...1",
  "short_message": "A short message %d",
  "full_message": "Backtrace here\n\nmore stuff",
  "timestamp": %d,
  "level": 1,
  "_user_id": %d,
  "_some_info": "foo",
  "_some_env_var": "bar"
}
`

func main() {

	count := 0

	c := client.New(client.Config{
		GraylogHost: "10.60.6.48",
		GraylogPort: 13000,
	})

	for {
		count++
		now := time.Now().Unix()
		message := fmt.Sprintf(msgTpl, count, now, count)

		c.Send(message)

		time.Sleep(time.Second)
	}
}
