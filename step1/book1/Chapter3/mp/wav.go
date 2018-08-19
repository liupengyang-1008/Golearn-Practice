// wav
package mp

import (
	"fmt"
	"time"
)

type MAVPlayer struct {
	stat     int
	progress int
}

func (p *MAVPlayer) Play(source string) {
	fmt.Println("Playing wav music ", source)

	p.progress = 0

	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		p.progress += 10
	}

	fmt.Println("\n Finished playing ", source)
}
