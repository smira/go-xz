package xz

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestDecompress(T *testing.T) {
	orig, err := os.Open("testdata/spec")
	if err != nil {
		T.Fatal(err)
	}
	defer orig.Close()

	expected := &bytes.Buffer{}
	_, err = io.Copy(expected, orig)
	if err != nil {
		T.Fatal(err)
	}

	source, err := os.Open("testdata/spec.xz")
	if err != nil {
		T.Fatal(err)
	}
	defer source.Close()

	r, err := NewReader(source)
	if err != nil {
		T.Fatal(err)
	}

	buf := &bytes.Buffer{}
	n, err := io.Copy(buf, r)
	if n != int64(expected.Len()) {
		T.Fail()
	}

	if bytes.Compare(buf.Bytes(), expected.Bytes()) != 0 {
		T.Fail()
	}

	err = r.Close()
	if err != nil {
		T.Fail()
	}
}

func TestEarlyClose(T *testing.T) {
	source, err := os.Open("testdata/spec.xz")
	if err != nil {
		T.Fatal(err)
	}
	defer source.Close()

	r, err := NewReader(source)
	if err != nil {
		T.Fatal(err)
	}

	buf := make([]byte, 10)
	n, err := r.Read(buf)
	if n != len(buf) {
		T.Fail()
	}

	err = r.Close()
	if err != nil {
		T.Fail()
	}
}

func BenchmarkDecompress(b *testing.B) {
	data, err := ioutil.ReadFile("testdata/spec.xz")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, err := NewReader(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_, err = io.Copy(ioutil.Discard, r)
		if err != nil {
			b.Fatal(err)
		}
	}
}
