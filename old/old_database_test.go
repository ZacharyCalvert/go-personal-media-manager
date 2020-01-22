package old

import (
	"testing"
)

func TestReadOldDatabase(t *testing.T) {
	loaded := ReadOld("./test.yml")
	if loaded["label1"].Sha256 != "label1" {
		t.Errorf("Wanted label1, got %s", loaded["label1"].Sha256)
	}
	var label1 = loaded["label1"]
	if len(label1.FileExtensions) != 2 {
		t.Errorf("Missing extensions")
	}
}
