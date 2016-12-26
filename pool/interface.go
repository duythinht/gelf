package pool

type Pool interface {
	Fire(buffer []byte)
	Close()
}
