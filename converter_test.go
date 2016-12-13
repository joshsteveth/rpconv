package rpconv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRp(t *testing.T) {
	input := float64(100550123) //100 mils 550 thousands 123
	rp := New(input)
	assert.Equal(t, 0, rp.billions.val)
	assert.Equal(t, 100, rp.millions.val)
	assert.Equal(t, 550, rp.thousands.val)
	assert.Equal(t, 123, rp.ones.val)
}

func TestPrintHundred(t *testing.T) {
	h := hundred{val: 123}
	assert.Equal(t, "SERATUS DUA PULUH TIGA", fmt.Sprintf("%s", h))
}

func TestPrintRupiah(t *testing.T) {
	input := float64(993111550123)
	rp := New(input)
	assert.Equal(t, "SEMBILAN RATUS SEMBILAN PULUH TIGA MILYAR SERATUS SEBELAS JUTA LIMA RATUS LIMA PULUH RIBU SERATUS DUA PULUH TIGA RUPIAH", fmt.Sprintf("%s", rp))

	input = float64(0)
	rp = New(input)
	assert.Equal(t, "NOL RUPIAH", fmt.Sprintf("%s", rp))

	input = float64(-1)
	rp = New(input)
	assert.Equal(t, "NOL RUPIAH", fmt.Sprintf("%s", rp))
}

func TestSpecialSingleTens(t *testing.T) {
	input := float64(513)
	assert.Equal(t, "LIMA RATUS TIGA BELAS RUPIAH", Convert(input))

	input = float64(510111)
	assert.Equal(t, "LIMA RATUS SEPULUH RIBU SERATUS SEBELAS RUPIAH", Convert(input))
}
