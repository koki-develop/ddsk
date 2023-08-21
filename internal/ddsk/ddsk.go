package ddsk

import (
	"bytes"
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
	Color   bool
	Animate bool
}

type DDSK struct {
	config *Config
}

func New(cfg *Config) *DDSK {
	return &DDSK{config: cfg}
}

func (d *DDSK) Run(w io.Writer) error {
	cur := new(bytes.Buffer)

	for {
		next := d.choose()
		if !d.config.Color {
			if _, err := w.Write([]byte(next)); err != nil {
				return err
			}
		}

		cur.WriteString(next)
		if cur.Len() > targetLen {
			if d.config.Color {
				if _, err := io.Copy(w, cur); err != nil {
					return err
				}
			}
			cur.Reset()
			continue
		}

		if cur.String() == strings.Repeat(target, targetRepeat) {
			if d.config.Color {
				rainbowColors := []string{"\x1b[31m", "\x1b[32m", "\x1b[33m", "\x1b[34m", "\x1b[35m", "\x1b[36m"}

				for i, c := range []rune(cur.String()) {
					if _, err := w.Write([]byte(rainbowColors[i%len(rainbowColors)])); err != nil {
						return err
					}
					if _, err := w.Write([]byte(string(c))); err != nil {
						return err
					}
				}

				if _, err := w.Write([]byte("\x1b[1m\x1b[91m")); err != nil {
					return err
				}
				if _, err := w.Write([]byte(love)); err != nil {
					return err
				}

				if _, err := w.Write([]byte("\x1b[0m")); err != nil {
					return err
				}
			} else {
				if _, err := w.Write([]byte(love)); err != nil {
					return err
				}
			}
			break
		}
	}

	return nil
}

func (*DDSK) choose() string {
	if rand.Intn(2) == 0 {
		return dd
	}
	return sk
}
