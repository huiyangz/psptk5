package model

type TK5File struct {
	BaseAddr      int64
	Size          int64
	OriginBuffer  []byte
	DecryptBuffer []byte
}
