package ddsk

import (
	"bytes"
	"io"
	"math/rand"
	"strings"
	"time"
)

const (
	dd      = "ドド"
	sk      = "スコ"
	love    = "ラブ"
	inject1 = "注"
	inject2 = "入"
	heart   = "♡"

	ddsksksk = dd + sk + sk + sk
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

// TODO: refactor
func (d *DDSK) Run(w io.Writer) error {
	cur := new(bytes.Buffer)

	for {
		next := d.choose()
		if !d.config.Color && !d.config.Animate {
			if _, err := w.Write([]byte(next)); err != nil {
				return err
			}
		}

		cur.WriteString(next)
		if cur.Len() > len(ddsksksk)*3 {
			if d.config.Color || d.config.Animate {
				if _, err := io.Copy(w, cur); err != nil {
					return err
				}
			}
			cur.Reset()
			continue
		}

		if cur.String() == strings.Repeat(ddsksksk, 3) {
			if d.config.Color || d.config.Animate {
				if err := d.ddsk(w); err != nil {
					return err
				}
			}

			if err := d.injectLove(w); err != nil {
				return err
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

func (d *DDSK) injectLove(w io.Writer) error {
	if !d.config.Color && !d.config.Animate {
		if _, err := w.Write([]byte(love + inject1 + inject2 + heart)); err != nil {
			return err
		}
		return nil
	}

	if d.config.Color {
		if _, err := w.Write([]byte("\x1b[1m\x1b[91m")); err != nil {
			return err
		}
	}

	for _, s := range []string{love, inject1, inject2 + heart} {
		if _, err := w.Write([]byte(s)); err != nil {
			return err
		}
		if d.config.Animate {
			time.Sleep(400 * time.Millisecond)
		}
	}

	if d.config.Color {
		if _, err := w.Write([]byte("\x1b[0m")); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte("\n")); err != nil {
		return err
	}

	return nil
}

func (d *DDSK) ddsk(w io.Writer) error {
	colors := []string{"\x1b[31m", "\x1b[32m", "\x1b[33m", "\x1b[34m", "\x1b[35m", "\x1b[36m"}

	for i, c := range []rune(strings.Repeat(ddsksksk, 3)) {
		if d.config.Color {
			if _, err := w.Write([]byte(colors[i%len(colors)])); err != nil {
				return err
			}
		}

		if _, err := w.Write([]byte(string(c))); err != nil {
			return err
		}

		if d.config.Animate {
			time.Sleep(200 * time.Millisecond)
		}
	}

	if d.config.Color {
		if _, err := w.Write([]byte("\x1b[0m")); err != nil {
			return err
		}
	}

	return nil
}
