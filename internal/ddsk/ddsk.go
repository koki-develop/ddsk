package ddsk

import (
	"io"
	"math/rand"
	"strings"
)

const (
	dd   = "ドド"
	sk   = "スコ"
	love = "ラブ注入♡\n"

	target       = dd + sk + sk + sk
	targetRepeat = 3
	targetLen    = len(target) * targetRepeat
)

type Config struct {
	Color bool
}

type DDSK struct {
	config *Config
}

func New(cfg *Config) *DDSK {
	return &DDSK{config: cfg}
}

func (d *DDSK) Run(w io.Writer) error {
	cur := new(strings.Builder)

	for {
		var next string
		if rand.Intn(2) == 0 {
			next = dd
		} else {
			next = sk
		}
		if _, err := w.Write([]byte(next)); err != nil {
			return err
		}

		cur.WriteString(next)
		if cur.Len() > targetLen {
			cur.Reset()
		}

		if cur.String() == strings.Repeat(target, targetRepeat) {
			if _, err := w.Write([]byte(love)); err != nil {
				return err
			}
			break
		}
	}

	return nil
}
