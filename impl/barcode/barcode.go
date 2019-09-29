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
	enc  common.EncodingFormat
	w, h int      // dimensions
	qr   struct { // Parameters for qr encoding
		level qr.ErrorCorrectionLevel
		mode  qr.Encoding
	}
}

// Compiler checks interface contract.
var _ common.BarCoder = new(BC)

// New BC constructor.
// SDefaults to QR200x200H
func New() *BC {
	b := new(BC)
	b.SetFormat(common.QR200x200H)
	return b
}

// SetFormat sets the EncodingFormat to use.
func (bc *BC) SetFormat(enc common.EncodingFormat) {
	switch enc {
	case common.QR200x200H:
		bc.enc = common.QR200x200H
		bc.w, bc.h = 200, 200
		bc.qr.level = qr.H
		bc.qr.mode = qr.Auto

	case common.QR300x300H:
		bc.enc = common.QR300x300H
		bc.w, bc.h = 300, 300
		bc.qr.level = qr.H
		bc.qr.mode = qr.Auto

	case common.DM200x200:
		bc.enc = common.DM200x200
		bc.w, bc.h = 200, 200

	case common.DM300x300:
		bc.enc = common.DM300x300
		bc.w, bc.h = 300, 300

	default:
		panic("Format is not recogized ?")
	}
}

// Encode into a 'barcode' according to the specified encoding format.
func (bc *BC) Encode(w io.Writer, txt string) error {

	switch bc.enc {
	case common.QR200x200H, common.QR300x300H:
		return bc.qrEncoder(w, txt)

	case common.DM200x200, common.DM300x300:
		return bc.dmEncoder(w, txt)

	default:
		panic("Format is not recogized ?")
	}
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
