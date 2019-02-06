// Copyright (C) 2018-2019 Kdevb0x Ltd.
// Thiss.software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package netsound

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/wav"
)

type file struct {
	filename     string
	filedir      string
	streamFormat beep.Format
	streamBuff   beep.StreamSeekCloser
	fhandle      *os.File
}

type StreamSource interface {
	// Type returns the source's encoding or file extention.
	Type() string

	// DecodeSource prepares a StreamSource to stream; returns an err if we don't
	// recognize, understand, or can't convert the type.
	DecodeSource() error
}

func OpenFileAsSource(path string) (*file, error) {
	var file = new(file)
	dir, fname := filepath.Split(path)
	file.filename = fname
	file.filedir = dir
	ofile, err := os.OpenFile(path, 0755, os.FileMode(os.O_RDONLY|os.O_SYNC))
	if err != nil {
		if e := err.(error); e == os.ErrNotExist {
			err = errors.New("unable to open file; file not found")

		}
		return nil, err

	}
	file.fhandle = ofile
	return file, nil

}

// Type returns the file extention.
func (s *file) Type() string {
	return filepath.Ext(s.filename)
}

// DecodeSource prepares a StreamSource to stream; returns an err if we don't
// recognize, understand, or can't convert the type.
func (s *file) DecodeSource() error {
	switch s.Type() {
	case "wav":
		streamer, sfmt, err := wav.Decode(s.fhandle)
		if err != nil {
			return fmt.Errorf("unable to decode wav file: %s\n", err.Error())

		}
		s.streamBuff = streamer
		s.streamFormat = sfmt
	case "mp3":
		streamer, sfmt, err := mp3.Decode(s.fhandle)
		if err != nil {
			return fmt.Errorf("unable to decode mp3 file: %s\n", err.Error())

		}
		s.streamBuff = streamer
		s.streamFormat = sfmt
	case "flac":
		streamer, sfmt, err := flac.Decode(s.fhandle)
		if err != nil {
			return fmt.Errorf("unable to decode flac file: %s\n", err.Error())
		}
		s.streamBuff = streamer
		s.streamFormat = sfmt
	default:
	}

	return fmt.Errorf("error: failed to initialize stream from %s: unrecognized source\n", s.Type())

}
