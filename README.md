# go-xz

Simple .xz decompression using external program (xz --decompress)

## Why?

All go implementation of .xz decompression rely on `liblzma` dependency providing
in-process decompression.

This package uses external `xz` utility, so no depdendencies for the compiled binary.

## License

MIT