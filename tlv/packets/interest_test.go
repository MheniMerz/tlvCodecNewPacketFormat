package packets

import (
	"ndn-router/nfd/tlv/name"
	"testing"
)

func TestNewInterest(t *testing.T) {
	name := name.NewNameFromString("/a/b/c")
	i := NewInterest(name)
	if i == nil {
		t.Errorf("new interest : nil")
	} else {
		t.Logf("%v", name)
		t.Logf("%v", i)
	}
}
