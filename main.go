package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aquilax/m3u"
)

func main() {
	var set = make(map[string]struct{})
	var err error
	var playlist m3u.Playlist
	var resulting m3u.Playlist
	var i int
	var line string
	var ok bool
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()
		if line[0] == '#' {
			continue
		}
		fmt.Fprintln(os.Stderr, scanner.Text())
		playlist, err = m3u.Parse(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		for i = range playlist.Tracks {
			if _, ok = set[playlist.Tracks[i].URI]; ok {
				continue
			}
			resulting.Tracks = append(resulting.Tracks, playlist.Tracks[i])
			set[playlist.Tracks[i].URI] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	reader, err := m3u.Marshall(resulting)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(data)
}
