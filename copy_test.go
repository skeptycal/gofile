package gofile

import (
	"bufio"
	"math/rand"
	"os"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

const fakesize = 2 << 16

var fake *bufio.ReadWriter

func makebuf(size int) []byte {
	b := make([]byte, 0, size)

	for i := 0; i < size; i++ {
		b = append(b, byte(rand.Intn(254)))
	}

	return b
}

func makeFake(src, dest string) (*bufio.ReadWriter, error) {

	s, err := os.Create(src)
	if err != nil {
		return nil, err
	}

	d, err := os.Create(dest)
	if err != nil {
		return nil, err
	}

	_, err = s.Write(makebuf(fakesize))
	if err != nil {
		return nil, err
	}

	err = s.Sync()
	if err != nil {
		return nil, err
	}

	err = s.Close()
	if err != nil {
		return nil, err
	}

	s, err = os.Open(src)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(s)
	w := bufio.NewWriter(d)

	br := bufio.NewReadWriter(r, w)

	return br, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())

	f, err := makeFake("fakeSrc", "fakeDst")
	if err != nil {
		log.Error(err)
	}

	fake = f
}

func Test_copy(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := copy(tt.args.src, tt.args.dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("copy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("copy() = %v, want %v", got, tt.want)
			}
		})
	}
}
