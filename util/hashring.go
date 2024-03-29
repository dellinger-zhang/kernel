package util

import (
	"crypto/sha1"
	"math"
	"sort"
	"strconv"
)

const (
	defaultVirualNodes = 100
)

type node struct {
	nodeKey   string
	nodeValue uint32
}

type nodesArray []node

func (p nodesArray) Len() int {
	return len(p)
}

func (p nodesArray) Less(i, j int) bool {
	return p[i].nodeValue < p[j].nodeValue
}

func (p nodesArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p nodesArray) Sort() {
	sort.Sort(p)
}

// HashRing consistent hashing
type HashRing struct {
	virualSpots int
	nodes       nodesArray
	weights     map[string]int
}

// NewHashRing constructor
func NewHashRing(spots int) *HashRing {
	if spots <= 1 {
		spots = defaultVirualNodes
	}

	h := &HashRing{
		virualSpots: spots,
		weights:     make(map[string]int),
	}
	return h
}

// AddNodes append nodes
func (h *HashRing) AddNodes(nodeWeight map[string]int) {
	for nodeKey, w := range nodeWeight {
		h.weights[nodeKey] = w
	}
	h.generate()
}

// AddNode append node
func (h *HashRing) AddNode(nodeKey string, weight int) {
	h.weights[nodeKey] = weight
	h.generate()
}

// RemoveNode remove node
func (h *HashRing) RemoveNode(nodeKey string) {
	delete(h.weights, nodeKey)
	h.generate()
}

// UpdateNode update weight of node
func (h *HashRing) UpdateNode(nodeKey string, weight int) {
	h.weights[nodeKey] = weight
	h.generate()
}

func (h *HashRing) generate() {
	var totalW int
	for _, w := range h.weights {
		totalW += w
	}

	totalVirtualSpots := h.virualSpots * len(h.weights)
	h.nodes = nodesArray{}

	for nodeKey, w := range h.weights {
		spots := int(math.Floor(float64(w) / float64(totalW) * float64(totalVirtualSpots)))
		for i := 1; i <= spots; i++ {
			hash := sha1.New()
			hash.Write([]byte(nodeKey + ":" + strconv.Itoa(i)))
			hashBytes := hash.Sum(nil)
			n := node{
				nodeKey:   nodeKey,
				nodeValue: genValue(hashBytes[6:10]),
			}
			h.nodes = append(h.nodes, n)
			hash.Reset()
		}
	}
	h.nodes.Sort()
}

func genValue(bs []byte) uint32 {
	if len(bs) < 4 {
		return 0
	}
	v := (uint32(bs[3]) << 24) | (uint32(bs[2]) << 16) | (uint32(bs[1]) << 8) | (uint32(bs[0]))
	return v
}

// GetNode returns the node
func (h *HashRing) GetNode(s string) string {
	if len(h.nodes) == 0 {
		return ""
	}

	hash := sha1.New()
	hash.Write([]byte(s))
	hashBytes := hash.Sum(nil)
	v := genValue(hashBytes[6:10])
	i := sort.Search(len(h.nodes), func(i int) bool {
		return h.nodes[i].nodeValue >= v
	})

	if i == len(h.nodes) {
		i = 0
	}

	return h.nodes[i].nodeKey
}
