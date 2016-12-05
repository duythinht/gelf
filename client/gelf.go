package client

import (
	"fmt"

	"github.com/duythinht/gelf/chunk"
	p "github.com/duythinht/gelf/pool"
)

type Config struct {
	GraylogHost  string
	GraylogPort  int
	Connection   string //wan/lan
	MaxChunkSize int
	WorkerNumber int
}

type Gelf struct {
	pool   p.Pool
	config Config
}

const (
	defaultHost         = "localhost"
	defaultPort         = 12201
	defaultConn         = "lan"
	defaultMaxLan       = 8196
	defaultMaxWan       = 1024
	defaultWorkerNumber = 4
)

func New(config Config) *Gelf {
	if config.GraylogHost == "" {
		config.GraylogHost = defaultHost
	}

	if config.GraylogPort == 0 {
		config.GraylogPort = defaultPort
	}

	if config.Connection == "" {
		config.Connection = defaultConn
	}

	if config.MaxChunkSize == 0 {
		if config.Connection == "lan" {
			config.MaxChunkSize = defaultMaxLan
		} else {
			config.MaxChunkSize = defaultMaxWan
		}
	}

	if config.WorkerNumber == 0 {
		config.WorkerNumber = defaultWorkerNumber
	}
	g := Gelf{config: config}
	g.connect()
	return &g
}

func (g *Gelf) connect() {
	address := fmt.Sprintf("%s:%d", g.config.GraylogHost, g.config.GraylogPort)
	g.pool = p.NewUDPPool(address, g.config.WorkerNumber)
}

func (g *Gelf) Send(message string) {
	fmt.Println("Send", message)
	buf := chunk.ZipMessage(message)
	if len(buf) > g.config.MaxChunkSize {
		bfs := chunk.GetGelfChunks(buf, g.config.MaxChunkSize)
		for _, bf := range bfs {
			g.pool.Fire(bf)
		}
	} else {
		g.pool.Fire(buf)
	}
}
