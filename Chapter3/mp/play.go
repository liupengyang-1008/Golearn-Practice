// play
package mp

import (
	"fmt"
)

type Player interface {
	Play(Source string)
}

func Play(source, mtype string) {
	var p Player
	switch mtype {
	case "MP3":
		p = &MP3Player{}
	case "WAV":
		p = &MAVPlayer{}
	default:
		fmt.Println("Unsupported music type", mtype)
		return
	}
	p.Play(source)
}
