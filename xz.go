// Package xz implements simple .xz decompression using external xz program
//
// No shared library (liblzma) dependencies.
package xz

import (
	"io"
	"os/exec"
)

// Reader does decompression using xz utility
type Reader struct {
	io.ReadCloser
}

// NewReader creates .xz decompression reader
//
// Internally it starts xz program, sets up input and output pipes
func NewReader(src io.Reader) (*Reader, error) {
	cmd := exec.Command("xz", "--decompress", "--stdout", "-T0")
	cmd.Stdin = src
	out, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	return &Reader{ReadCloser: out}, nil
}
