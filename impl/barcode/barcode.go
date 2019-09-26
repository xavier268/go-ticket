// Package barcode provide easy to use services to produce barcode images.
// Internally, it relies on https://github.com/boombuler/barcode
package barcode

import (
	"image/png"
	"io"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/datamatrix"
	"github.com/boombuler/barcode/qr"
	"github.com/xavier268/go-ticket/common"
)

// BC is the BarCoder inmplementation.
// It uses https://github.com/boombuler/barcode
type BC struct {
	enc EncodingFormat
	// Encodes in the requested format.
	Encode func(io.Writer, string) error
	w, h   int      // dimensions
	qr     struct { // Parameters for qr encoding
		level qr.ErrorCorrectionLevel
		mode  qr.Encoding
	}
}

// EncodingFormat  defines QR, DataMatrix, ...
type EncodingFormat int

// Encoding formats that are available.
const (
	QR200x200H EncodingFormat = iota
	QR300x300H
	DM200x200
	DM300x300
)

// Compiler checks interface contract.
var _ common.BarCoder = new(BC)

// New BC constructor.
// SDefaults to QR200x200H
func New() *BC {
	b := new(BC)
	b.SetFormat(QR200x200H)
	return b
}

// SetFormat sets the EncodingFormat to use.
func (bc *BC) SetFormat(enc EncodingFormat) *BC {
	switch enc {
	case QR200x200H:
		bc.enc = QR200x200H
		bc.w, bc.h = 200, 200
		bc.qr.level = qr.H
		bc.qr.mode = qr.Auto
		bc.Encode = bc.qrEncoder
	case QR300x300H:
		bc.enc = QR300x300H
		bc.w, bc.h = 300, 300
		bc.qr.level = qr.H
		bc.qr.mode = qr.Auto
		bc.Encode = bc.qrEncoder
	case DM200x200:
		bc.enc = DM200x200
		bc.w, bc.h = 200, 200
		bc.Encode = bc.dmEncoder
	case DM300x300:
		bc.enc = DM300x300
		bc.w, bc.h = 300, 300
		bc.Encode = bc.dmEncoder
	default:
		panic("Format is not recogized ?")
	}
	return bc
}

// qrEncoder encodes as qr-code png into an io.Writer
func (bc *BC) qrEncoder(file io.Writer, txt string) error {

	// Encode qr code
	qr, err := qr.Encode(txt, bc.qr.level, bc.qr.mode)
	if err != nil {
		return err
	}

	// Scale to size w x h
	qr, err = barcode.Scale(qr, bc.w, bc.h)
	if err != nil {
		return nil
	}

	// encode as png to io.Writer
	return png.Encode(file, qr)
}

// dmEncoder is a DataMatrix encoder
func (bc *BC) dmEncoder(file io.Writer, txt string) error {

	// Encode qr code
	qr, err := datamatrix.Encode(txt)
	if err != nil {
		return err
	}

	// Scale to size w x h
	qr, err = barcode.Scale(qr, bc.w, bc.h)
	if err != nil {
		return nil
	}

	// encode as png to io.Writer
	return png.Encode(file, qr)
}
