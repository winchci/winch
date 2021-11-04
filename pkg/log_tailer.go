package pkg

import (
	"crypto/md5"
	"github.com/mitchellh/colorstring"
	"io"
	"strings"
)

type LogTailer struct {
	output     io.Writer
	prefix     string
	needPrefix bool
}

func NewLogTailer(output io.Writer, prefix string) io.Writer {
	return &LogTailer{
		output:     output,
		prefix:     prefix,
		needPrefix: true,
	}
}

func (l *LogTailer) Write(p []byte) (n int, err error) {
	c := &colorstring.Colorize{
		Colors: colorstring.DefaultColors,
	}
	prefix := c.Color("[reset]" + l.prefix + " ")

	var p2 []byte
	for _, pb := range p {
		if l.needPrefix {
			p2 = append(p2, []byte(prefix)...)
			l.needPrefix = false
		}

		p2 = append(p2, pb)
		if pb == '\n' {
			l.needPrefix = true
		}
	}

	_, err = l.output.Write(p2)

	n = len(p)
	return
}

func ColorName(s string) string {
	var x int
	for _, c := range md5.Sum([]byte(s)) {
		x = x + int(c)
	}
	colors := []string{
		"red",
		"green",
		"yellow",
		"blue",
		"magenta",
		"cyan",
		"light_gray",
		"dark_gray",
		"light_red",
		"light_green",
		"light_yellow",
		"light_blue",
		"light_magenta",
		"light_cyan",
		"white",
	}
	return "[" + colors[x%len(colors)] + "]" + s + "[reset]"
}

func PadName(s string, n int) string {
	if len(s) >= n {
		return s
	}

	return s + strings.Repeat(" ", n-len(s))
}
