package main

import "io"

type Generator struct {
	template     *[]byte
	insertOffset int
	insert       []byte
	pos          int
}

func (gen *Generator) Read(p []byte) (n int, err error) {
	insertLen := len(gen.insert)
	pLen := len(p)
	finalPos := gen.pos + pLen
	insertFinalPos := gen.insertOffset + insertLen
	temp := *gen.template
	var read = 0
	if finalPos < gen.insertOffset || gen.pos > insertFinalPos {
		read += copy(p, temp[gen.pos:])
	} else {
		if gen.pos < gen.insertOffset {
			read += copy(p, temp[gen.pos+read:gen.insertOffset])
		}
		read += copy(p[read:], gen.insert[gen.pos+read-gen.insertOffset:])
		if read < pLen {
			read += copy(p[read:], temp[gen.pos+read:])
		}
	}
	gen.pos += read
	if gen.pos == len(temp) {
		return read, io.EOF
	}
	return read, nil
}
