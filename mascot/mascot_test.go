package mascot_test

import (
	"HD-27/DEMO_1_GOLANG.git/DEMO_1_GOLANG/DEMO_1_GOLANG/mascot"
	"testing"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "GO Gopher" {
		t.Fail("wrong mascot")
	}
}
