package assignwork

import (
	"hash/crc64"
	"math/rand"
)

var isoTab = crc64.MakeTable(crc64.ISO)

// Pool represents a work pool of Size members.
// Its methods assign ownership of work items,
// which are represented by byte sequences.
type Pool struct {
	Size int
}

// Owners returns the list of owners for item, in priority order.
// The returned slice has length p.Size.
//
// If a single owner is needed, use element 0 in the returned slice;
// if two are needed, use 0 and 1; etc.
func (p Pool) Owners(item string) []int {
	c := int64(crc64.Checksum([]byte(item), isoTab))
	return rand.New(rand.NewSource(c)).Perm(p.Size)
}

// OwnersExcluding is like Owners, except it removes from its
// return value the elements of exclude. (So the length of the
// returned slice may be less than p.Size.)
func (p Pool) OwnersExcluding(item string, exclude ...int) []int {
	var a []int
	for _, k := range p.Owners(item) {
		if !contains(exclude, k) {
			a = append(a, k)
		}
	}
	return a
}

// Member represents a member of a work pool.
type Member struct {
	ID int // Must be >= 0 and < Pool.Size
	Pool
}

// NewMember returns a Member with ID congruent to k mod n,
// in a Pool of size n.
func NewMember(k, n int) Member {
	return Member{(k%n + n) % n, Pool{n}}
}

// Owns returns whether m is the first owner of item.
func (m Member) Owns(item string) bool {
	return m.In(m.Owners(item)[:1])
}

// Owns returns whether m is in owners.
func (m Member) In(owners []int) bool {
	return contains(owners, m.ID)
}

func contains(a []int, i int) bool {
	for _, n := range a {
		if n == i {
			return true
		}
	}
	return false
}
