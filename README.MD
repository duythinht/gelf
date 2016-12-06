GO: GELF implements
Because of performance, I recommender to build your own TCP or UDP connection to send GELF data:

#### Quick use
    
    import "github.com/duythinht/gelf"
   
#### Create message JSON

    func main() {
      message := gelf.Create("Short message content")
                 .SetTimestamp(time.Now().Unix())
                 .SetFullMessage("This is full message")
                 .SetLevel(3)
                 .SetHost("chat Server")
                 .ToJSON()
    }


#### Zip message

    import "github.com/duythinht/gelf/chunk"

    func main() {
      // message := ... # JSON message
      buffer := chunk.ZipMessage(message)
    }


#### Chunking message

    import "github.com/duythinht/gelf/chunk"

    func main() {
      // ZippedMessage := ... # type []byte
      if len(ZippedMessage) > MaxChunkSize) {
        buffers := chunk.GetGelfChunks(ZippedMessage, MaxChunkSize)
      }
    }

#### BTW, I offer a UDP clients pool to send message to graylog (workers by goroutine):


    import "github.com/duythinht/gelf/client"

    c := client.New(client.Config{
  	  GraylogHost: "10.60.6.48",
      GraylogPort: 13000,
    })
    
    c.Send("This is test message, auto zipping, auto chunking") 
