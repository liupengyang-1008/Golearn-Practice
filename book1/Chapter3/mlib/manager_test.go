// manager_test
package mlib

import (
	"testing"
)

func TestPos(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed.")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManager failed ,not empty.")
	}

	m0 := &MusicEntry{
		"1", "My Heart Will Go On", "Celion Dion",
		"http://qbox.me/123213", "MP3"}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("NewMusicManager.add() failed")

	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("NewMusicManager.find() failed.")
	}

	if m.Id != m0.Id || m.Artist != m0.Artist ||
		m.Name != m0.Name || m.Source != m0.Source ||
		m.Type != m0.Type {
		t.Error("NewMusicManager.find() failed . Found item misatch.")
	}
	m, err := mm.Get(0)
	if m == nil {
		t.Error("NewMusicManager.get() failed.", err)
	}

	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("NewMusicManager.Remove() failed .", err)
	}

	m = mm.RemoveByName("My Heart Will Go On")
	if m == nil || mm.Len() != 0 {
		t.Error("NewMusicManager.RemoveByName() failed .", err)
	}

}
