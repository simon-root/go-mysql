package utils

import (
	"bytes"
	"sync"
)

var (
	bytesBufferPool = sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
	bytesBufferChan = make(chan *bytes.Buffer, 10)
)

func BytesBufferGet() (data *bytes.Buffer) {
	select {
	case data = <-bytesBufferChan:
	default:
		data = bytesBufferPool.Get().(*bytes.Buffer)
	}

	data.Reset()

	return data
}

func BytesBufferPut(data *bytes.Buffer) {
	if len(data.Bytes()) > 1024*1024*16 {
		return
	}

	select {
	case bytesBufferChan <- data:
	default:
		bytesBufferPool.Put(data)
	}
}
