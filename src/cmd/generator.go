package main

import (
	log "github.com/Sirupsen/logrus"
	"io"
)

type Generator struct {
	template     *[]byte
	insertOffset int
	insert       []byte
	pos          int
}

func (gen *Generator) Read(p []byte) (n int, err error) {
	temp := *gen.template
	if gen.pos >= len(temp) {
		return 0, io.EOF
	}
	insertLen := len(gen.insert)
	pLen := len(p)
	finalPos := gen.pos + pLen
	insertFinalPos := gen.insertOffset + insertLen
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
	if gen.pos > len(temp) {
		log.Warnf("Invalid final position: %d > %d", gen.pos, len(temp))
	}
	if gen.pos == len(temp) {
		return read, io.EOF
	}
	return read, nil
}
