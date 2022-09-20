package util

import (
	"crypto/rand"
	"io"
	"sync"
	"time"

	"github.com/oklog/ulid"
)

var (
	mux     sync.Mutex
	once    sync.Once
	entropy io.Reader
)

func init() {
	initializeEntropy()
}

func NewULID() (ulid.ULID, error) {
	mux.Lock()
	defer mux.Unlock()
	now := time.Now()
	id, err := ulid.New(ulid.Timestamp(now), entropy)
	if err != nil {
		return ulid.ULID{}, err
	}
	return id, nil
}

func initializeEntropy() {
	mux.Lock()
	defer mux.Unlock()

	once.Do(func() {
		entropy = ulid.Monotonic(rand.Reader, 0)
	})
}
