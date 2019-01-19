package server

import (
	"io/ioutil"
	"net/http"
	"testing"
)

// endpoint serves an mp4 file, TODO(kdd): figure out what/how this works.
var mp4respInWild = "https://smartech.gatech.edu/bitstream/handle/1853/54082/fisher.mp4?sequence=1&isAllowed=y"

func TestMP4Resp(t *testing.T) {
	resp, err := http.Get(mp4respInWild)
	if err != nil {
		t.Log(err)
	}
	var header = resp.Header
	var contLen, status = resp.ContentLength, int64(resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
	}
}
