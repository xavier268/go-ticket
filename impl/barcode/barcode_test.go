package barcode

import (
	"os"
	"testing"

	"github.com/xavier268/go-ticket/common"
)

func TestQR(t *testing.T) {

	f, _ := os.Create("qrcode.png")
	defer f.Close()

	b := New()
	b.SetFormat(common.QR200x200H)
	err := b.Encode(f, "https://github.com/xavier268/go-ticket")
	if err != nil {
		t.Fatal(err)
	}

}

func TestDM(t *testing.T) {

	f, _ := os.Create("dmcode.png")
	defer f.Close()

	b := New()
	b.SetFormat(common.DM200x200)
	err := b.Encode(f, "https://github.com/xavier268/go-ticket")
	if err != nil {
		t.Fatal(err)
	}

}
