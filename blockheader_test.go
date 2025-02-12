package intcomp_test

import (
	"testing"

	"github.com/ronanh/intcomp"
)

func TestAddGroup(t *testing.T) {
	var gh intcomp.BlockHeader
	// add a group with bitlen=10 and ntz=3
	gh = gh.AddGroup(10, 3)
	// check that the resulting header has one group with bitlen=10 and ntz=3
	if gh.GroupCount() != 1 {
		t.Errorf("expected 1 group, got %d", gh.GroupCount())
	}
	bitlen, ntz := gh.GetGroup(0)
	if bitlen != 10 {
		t.Errorf("expected bitlen=10, got %d", bitlen)
	}
	if ntz != 3 {
		t.Errorf("expected ntz=3, got %d", ntz)
	}
	// add a group with bitlen=20 and ntz=5
	gh = gh.AddGroup(20, 5)
	// check that the resulting header has two groups with bitlen=10 and ntz=3 and bitlen=20 and ntz=5
	if gh.GroupCount() != 2 {
		t.Errorf("expected 2 groups, got %d", gh.GroupCount())
	}
	bitlen, ntz = gh.GetGroup(0)
	if bitlen != 10 {
		t.Errorf("expected bitlen=10, got %d", bitlen)
	}
	if ntz != 3 {
		t.Errorf("expected ntz=3, got %d", ntz)
	}
	bitlen, ntz = gh.GetGroup(1)
	if bitlen != 20 {
		t.Errorf("expected bitlen=20, got %d", bitlen)
	}
	if ntz != 5 {
		t.Errorf("expected ntz=5, got %d", ntz)
	}
	// add a group with bitlen=30 and ntz=7
	gh = gh.AddGroup(30, 7)
	// check that the resulting header has three groups with bitlen=10 and ntz=3 and bitlen=20 and ntz=5 and bitlen=30 and ntz=7
	if gh.GroupCount() != 3 {
		t.Errorf("expected 3 groups, got %d", gh.GroupCount())
	}
	bitlen, ntz = gh.GetGroup(0)
	if bitlen != 10 {
		t.Errorf("expected bitlen=10, got %d", bitlen)
	}
	if ntz != 3 {
		t.Errorf("expected ntz=3, got %d", ntz)
	}
	bitlen, ntz = gh.GetGroup(1)
	if bitlen != 20 {
		t.Errorf("expected bitlen=20, got %d", bitlen)
	}
	if ntz != 5 {
		t.Errorf("expected ntz=5, got %d", ntz)
	}
	bitlen, ntz = gh.GetGroup(2)
	if bitlen != 30 {
		t.Errorf("expected bitlen=30, got %d", bitlen)
	}
	if ntz != 7 {
		t.Errorf("expected ntz=7, got %d", ntz)
	}
	// add a group with bitlen=40 and ntz=9
	gh = gh.AddGroup(40, 9)
	// check that the resulting header has four groups with bitlen=10 and ntz=3 and bitlen=20 and ntz=5 and bitlen=30 and ntz=7 and bitlen=40 and ntz=9
	if gh.GroupCount() != 4 {
		t.Errorf("expected 4 groups, got %d", gh.GroupCount())
	}
	bitlen, ntz = gh.GetGroup(0)
	if bitlen != 10 {
		t.Errorf("expected bitlen=10, got %d", bitlen)
	}
	if ntz != 3 {
		t.Errorf("expected ntz=3, got %d", ntz)
	}
	bitlen, ntz = gh.GetGroup(1)
	if bitlen != 20 {
		t.Errorf("expected bitlen=20, got %d", bitlen)
	}
	if ntz != 5 {
		t.Errorf("expected ntz=5, got %d", ntz)
	}
	bitlen, ntz = gh.GetGroup(2)
	if bitlen != 30 {
		t.Errorf("expected bitlen=30, got %d", bitlen)
	}
	if ntz != 7 {
		t.Errorf("expected ntz=7, got %d", ntz)
	}
	bitlen, ntz = gh.GetGroup(3)
	if bitlen != 40 {
		t.Errorf("expected bitlen=40, got %d", bitlen)
	}
	if ntz != 9 {
		t.Errorf("expected ntz=9, got %d", ntz)
	}
}
