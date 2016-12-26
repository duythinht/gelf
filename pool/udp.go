package pool

import (
	"fmt"
	"net"
)

const (
	_net = "udp4" //
)

func createUDPConnection(address string) (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr(_net, address)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialUDP(_net, nil, addr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func worker(id int, address string, done chan bool, buffers <-chan []byte) {

	conn, err := createUDPConnection(address)

	if err != nil {
		return
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	for buffer := range buffers {
		_, err := conn.Write(buffer)
		if err != nil {
			break
		}
	}

	done <- true
}

type UDPPool struct {
	buffers chan []byte
	done    chan bool
}

//NewUDPPool create new UDP workers pool
func NewUDPPool(address string, workerNumber int) *UDPPool {
	buffers := make(chan []byte, workerNumber)
	done := make(chan bool)

	for wid := 1; wid < workerNumber; wid++ {
		go worker(wid, address, done, buffers)
	}
	return &UDPPool{buffers, done}
}

func (p *UDPPool) Fire(buffer []byte) {
	p.buffers <- buffer
}

func (p *UDPPool) Close() {
	close(p.buffers)
}
