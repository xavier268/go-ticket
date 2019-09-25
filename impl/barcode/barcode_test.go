package barcode

import (
	"os"
	"testing"
)

func TestQR(t *testing.T) {

	f, _ := os.Create("qrcode.png")
	defer f.Close()

	b := New()
	b.SetFormat(QR200x200H)
	err := b.Encode(f, "https://github.com/xavier268/go-ticket")
	if err != nil {
		t.Fatal(err)
	}

}

func TestDM(t *testing.T) {

	f, _ := os.Create("dmcode.png")
	defer f.Close()

	b := New()
	b.SetFormat(D)
	err := b.Encode(f, "https://github.com/xavier268/go-ticket")
	if err != nil {
		t.Fatal(err)
	}

}
