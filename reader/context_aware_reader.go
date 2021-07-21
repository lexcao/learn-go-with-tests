package reader

import (
	"context"
	"io"
)

type ContextAwareReader struct {
	reader  io.Reader
	context context.Context
}

func (r *ContextAwareReader) Read(bytes []byte) (n int, err error) {
	if err = r.context.Err(); err != nil {
		return
	}

	return r.reader.Read(bytes)
}

func NewCancellableReader(ctx context.Context, reader io.Reader) io.Reader {
	return &ContextAwareReader{
		reader,
		ctx,
	}
}
