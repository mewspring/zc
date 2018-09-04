// The zc tool plays FLAC files using zikichombo.org/sio for playback.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"zikichombo.org/ext/flac"
	"zikichombo.org/sio"
)

func usage() {
	const use = `
Usage: zc [OPTION]... FILE.flac...`
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	for _, flacPath := range flag.Args() {
		if err := play(flacPath); err != nil {
			log.Fatalf("%+v", err)
		}
	}
}

// play plays the given FLAC audio file.
func play(flacPath string) error {
	// Open FLAC decoder.
	r, err := os.Open(flacPath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer r.Close()
	dec, err := flac.NewDecoder(r)
	if err != nil {
		return errors.WithStack(err)
	}
	// Play sound.
	if err := sio.Play(dec); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
