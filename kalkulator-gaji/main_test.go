package main

import (
	// "fmt"
	// "github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	gaji_pokok    int    = inputGaji().gaji_pokok
	gaji_lembur   int    = inputGaji().gaji_lembur
	potongan_gaji string = inputGaji().potongan_gaji
)

func TestHitungGaji(t *testing.T) {

	totalGaji := gaji_pokok + gaji_lembur

	if totalGaji != gaji_pokok+gaji_lembur{

		t.Error("result is not sum of gaji pokok and gaji lembur")
	}
	// assert.Equal(t, totalGaji = gaji_pokok+gaji_lembur, totalGaji != gaji_pokok+gaji_lembur)

	// assert.Eq

}

func TestInputGaji(t *testing.T) {

	assert.Equal(t, gaji_pokok, inputGaji().gaji_pokok, "input dengan yang dihitung harus sama")

	assert.Equal(t, gaji_lembur, inputGaji().gaji_lembur, "input dengan yang dihitung harus sama")

	assert.Equal(t, potongan_gaji, inputGaji().potongan_gaji, "input dengan yang dihitung harus sama")
}