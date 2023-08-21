package ddsk

import (
	"bytes"
	"io"
	"math/rand"
	"strings"
	"time"
)

const (
	dd       = "ドド"
	sk       = "スコ"
	love     = "ラブ"
	inject1  = "注"
	inject2  = "入"
	heart    = "♡"
	ddsksksk = dd + sk + sk + sk
)

type Config struct {
	Writer io.Writer

	Color   bool
	Animate bool
}

type DDSK struct {
	config *Config
}

func New(cfg *Config) *DDSK {
	return &DDSK{config: cfg}
}

func (d *DDSK) Run() error {
	cur := new(bytes.Buffer)

	for {
		if err := d.generateAndWriteSequence(cur); err != nil {
			return err
		}

		if cur.String() == strings.Repeat(ddsksksk, 3) {
			if err := d.generateColoredOrAnimatedString(d.ddsk); err != nil {
				return err
			}

			if err := d.injectLove(); err != nil {
				return err
			}
			break
		}
	}

	if _, err := d.config.Writer.Write([]byte("\n")); err != nil {
		return err
	}

	return nil
}

func (d *DDSK) generateAndWriteSequence(cur *bytes.Buffer) error {
	next := d.choose()
	cur.WriteString(next)

	if cur.Len() > len(ddsksksk)*3 {
		return d.flushBuffer(cur)
	}
	return nil
}

func (d *DDSK) flushBuffer(cur *bytes.Buffer) error {
	if _, err := io.Copy(d.config.Writer, cur); err != nil {
		return err
	}
	cur.Reset()
	return nil
}

func (d *DDSK) choose() string {
	if rand.Intn(2) == 0 {
		return dd
	}
	return sk
}

func (d *DDSK) injectLove() error {
	return d.generateColoredOrAnimatedString(func(w io.Writer) error {
		if d.config.Color {
			if _, err := w.Write([]byte("\x1b[1m")); err != nil {
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

		return nil
	})
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

	return nil
}

func (d *DDSK) generateColoredOrAnimatedString(generator func(io.Writer) error) error {
	if d.config.Color {
		if _, err := d.config.Writer.Write([]byte("\x1b[91m")); err != nil {
			return err
		}
	}

	if err := generator(d.config.Writer); err != nil {
		return err
	}

	if d.config.Color {
		if _, err := d.config.Writer.Write([]byte("\x1b[0m")); err != nil {
			return err
		}
	}

	return nil
}
