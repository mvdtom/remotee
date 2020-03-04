package main

import (
	"errors"
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

func (gen *Generator) Seek(offset int64, whence int) (int64, error) {
	newPos, offs := 0, int(offset)
	switch whence {
	case io.SeekStart:
		newPos = offs
	case io.SeekCurrent:
		newPos = gen.pos + offs
	case io.SeekEnd:
		newPos = len(*gen.template) + offs
	}
	if newPos < 0 {
		return 0, errors.New("negative position result")
	}
	gen.pos = newPos
	return int64(newPos), nil
}
