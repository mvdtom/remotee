package main

import (
	"io"
	"testing"
)

const templateLen = 100
const insertLen = 6
const offset = 40
const templateValue = 0xFF
const insertValue = 0xAA

func TestFullFilling(t *testing.T) {
	gen := createGet()
	result := makeFilledSlice(templateLen, 0x00)
	read, err := gen.Read(result)

	if err != io.EOF {
		t.Error("Result error is incorrect:", err)
	}
	if read != templateLen {
		t.Errorf("Invalid read value. Expected: [%d]. Actual: [%d]", templateLen, read)
	}

	for i, v := range result {
		var expected byte
		switch {
		case i < offset:
			expected = templateValue
		case i >= offset+insertLen:
			expected = templateValue
		default:
			expected = insertValue
		}
		if v != expected {
			t.Errorf("Invalid value at pos [%d]. Expected: [%d]. Actual: [%d]", i, expected, v)
		}
	}
}

func createGet() Generator {
	template := makeFilledSlice(templateLen, templateValue)
	insert := makeFilledSlice(insertLen, insertValue)
	gen := Generator{template: &template, insertOffset: offset, insert: insert}
	return gen
}

func makeFilledSlice(len int, val byte) []byte {
	slice := make([]byte, len)
	for i := range slice {
		slice[i] = val
	}
	return slice
}
