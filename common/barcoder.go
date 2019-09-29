package common

import "io"

// BarCoder can generate bar codes, qr-codes, data matrixes, ...
type BarCoder interface {
	SetFormat(enc EncodingFormat)
	Encode(w io.Writer, txt string) error
}

// EncodingFormat  defines QR, DataMatrix, ...
type EncodingFormat int

// SetFormat sets the EncodingFormat to use.

// Encoding formats that are available.
const (
	QR200x200H EncodingFormat = iota
	QR300x300H
	DM200x200
	DM300x300
)
