package util

import (
	"fmt"
	"testing"
)

func TestHashRingNormal(t *testing.T) {
	nodeWeight := make(map[string]int)
	nodeWeight["10.0.0.1"] = 2
	nodeWeight["10.0.0.2"] = 2
	nodeWeight["10.0.0.3"] = 3

	hash := NewHashRing(100)

	hash.AddNodes(nodeWeight)
	if hash.GetNode("1") != "10.0.0.1" {
		t.Fatalf("expect %v got %v", "10.0.0.1", hash.GetNode("1"))
	}
	if hash.GetNode("2") != "10.0.0.2" {
		t.Fatalf("expect %v got %v", "10.0.0.2", hash.GetNode("2"))
	}
	if hash.GetNode("3") != "10.0.0.3" {
		t.Fatalf("expect %v got %v", "10.0.0.3", hash.GetNode("3"))
	}

	hash.RemoveNode("10.0.0.3")
	if hash.GetNode("1") != "10.0.0.1" {
		t.Fatalf("expect %v got %v", "10.0.0.1", hash.GetNode("1"))
	}
	if hash.GetNode("2") != "10.0.0.2" {
		t.Fatalf("expect %v got %v", "10.0.0.1", hash.GetNode("2"))
	}
	if hash.GetNode("3") != "10.0.0.1" {
		t.Fatalf("expect %v got %v", "10.0.0.1", hash.GetNode("3"))
	}

	hash.AddNode("10.0.0.4", 4)
	if hash.GetNode("1") != "10.0.0.1" {
		t.Fatalf("expect %v got %v", "10.0.0.1", hash.GetNode("1"))
	}
	if hash.GetNode("2") != "10.0.0.2" {
		t.Fatalf("expect %v got %v", "10.0.0.2", hash.GetNode("2"))
	}
	if hash.GetNode("3") != "10.0.0.4" {
		t.Fatalf("expect %v got %v", "10.0.0.4", hash.GetNode("3"))
	}
	if hash.GetNode("4") != "10.0.0.1" {
		t.Fatalf("expect %v got %v", "10.0.0.1", hash.GetNode("4"))
	}
}

func TestHashRingGraph(t *testing.T) {

	nodeWeight := make(map[string]int)
	nodeWeight["10.0.0.1"] = 1
	nodeWeight["10.0.0.2"] = 2
	nodeWeight["10.0.0.3"] = 4

	hash := NewHashRing(100)
	hash.AddNodes(nodeWeight)
	var c1, c2, c3 int
	for i := 0; i < 1000; i++ {
		aa := hash.GetNode(fmt.Sprintf("%d", i))
		if "10.0.0.1" == aa {
			c1++
		} else if "10.0.0.2" == aa {
			c2++
		} else if "10.0.0.3" == aa {
			c3++
		}
	}

	t.Logf("10.0.0.1: %d", c1)
	t.Logf("10.0.0.2: %d", c2)
	t.Logf("10.0.0.3: %d", c3)
}
