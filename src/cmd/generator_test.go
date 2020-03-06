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

func TestSimpleFullFilling(t *testing.T) {
	gen := createGen()
	result := makeFilledSlice(templateLen, 0x00)
	read, err := gen.Read(result)

	checkErr(io.EOF, err, t)
	checkRead(templateLen, read, t)
	checkValues(result, 0, t)
}

func checkValues(result []byte, read int, t *testing.T) {
	for i, v := range result {
		var expected byte
		switch {
		case i+read < offset:
			expected = templateValue
		case i+read >= offset+insertLen:
			expected = templateValue
		default:
			expected = insertValue
		}
		if v != expected {
			t.Errorf("Invalid value at pos [%d]. Expected: [%d]. Actual: [%d]", i, expected, v)
		}
	}
}

func checkRead(expected, actual int, t *testing.T) {
	if actual != expected {
		t.Errorf("Invalid read value. Expected: [%d]. Actual: [%d]", expected, actual)
	}
}

func checkErr(expected, actual error, t *testing.T) {
	if actual != expected {
		t.Errorf("Error is incorrect. Expected: [%v]. Actual: [%v]", expected, actual)
	}
}

func TestFillingByParts(t *testing.T) {
	gen := createGen()
	result := makeFilledSlice(50, 0x0)
	read, err := gen.Read(result)

	checkErr(nil, err, t)
	checkRead(50, read, t)
	checkValues(result, 0, t)

	read, err = gen.Read(result)

	checkErr(io.EOF, err, t)
	checkRead(50, read, t)
	checkValues(result, 50, t)
}

func createGen() Generator {
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
