package pmm

import (
	"testing"
)

func TestReadOldDatabase(t *testing.T) {
	loaded := ReadOld("./old.yml")
	if loaded["label1"].Sha256 != "label1" {
		t.Errorf("Wanted label1, got %s", loaded["label1"].Sha256)
	}
}