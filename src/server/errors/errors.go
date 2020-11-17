package errors

type Error struct {
	error
	Code uint16
	Msg  string
}

const (
	NotImplement  = 0
	LoginRequired = 1
)
