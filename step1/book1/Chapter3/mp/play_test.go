// play_test.go
package mp

import (
	"fmt"
	"testing"
)

func TestPlay(t *testing.T) {
	p := new(MP3Player)
	fmt.Println(p)
	if p == nil {
		t.Error("Type Player var failed.")
	}
	p.Play("sss")
}
