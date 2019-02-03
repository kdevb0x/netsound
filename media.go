// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package netsound

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/wav"
)

type soundfile struct {
	filename     string
	filedir      string
	streamFormat *beep.Format
	streamBuff   beep.StreamSeekCloser
	handle       *os.File
}

func LoadSoundFile(path string) (*soundfile, error) {
	var file = new(soundfile)
	dir, fname := filepath.Split(path)
	file.filename = fname
	file.filedir = dir
	ofile, err := os.OpenFile(path, 0755, os.O_RDONLY|os.O_SYNC)
	if err != nil {
		if e := err.(os.ErrNotExist); e {
			oerr := errors.New("unable to open file; file not found!")
			return nil, oerr
		}

	}
	file.handle = ofile
	return file, nil

}

// Type returns the file extention.
func (s *soundfile) Type() string {
	return filepath.Ext(s.filename)
}

// InitStreamer prepares the audio file to stream, returns an err if we don't
// understand, or can't convert the type.
func (s *soundfile) InitStreamer() error {
	switch s.Type() {
	case "wav":
		streamer, sfmt, err := wav.Decode(s.handle)
		if err != nil {
			return errors.New("unable to decode wav file; wrong format")
		}
		s.streamBuff = streamer
		s.streamFormat = &sfmt
	case "mp3":
		streamer, sfmt, err := mp3.Decode(s.handle)
		if err != nil {
			return errors.New("unable to decode mp3 file; wrong format")
		}
		s.streamBuff = streamer
		s.streamFormat = &sfmt
	case "flac":
		streamer, sfmt, err := flac.Decode(s.handle)
		if err != nil {
			return errors.New("unable to decode flac file; wrong format")
		}
		s.streamBuff = streamer
		s.streamFormat = &sfmt
	default:
		return errors.New("unable to decode unknown file format; Must be one of: 'wav', 'mp3', or 'flac'")
	}
}
