package necurs

import (
	"fmt"
	"time"
)

// adapted from https://github.com/baderj/domain_generation_algorithms
type necurs struct {
	year  uint64
	month uint64
	day   uint64
	pos   uint64
	seed  uint64
}

func mod64(n1 uint64, n2 uint64) uint64 {
	return n1 % n2
}

func pseudoRandom(v uint64) uint64 {
	loops := int((v & 0x7F)) + 21
	for i := 0; i < loops; i++ {
		v += ((v * 7) ^ (v << 15)) + 8*uint64(i) - (v >> 5)
		v &= ((1 << 64) - 1)
	}
	return v
}

// SeedRNG Initalize the necurs algorithm
func SeedRNG(pos uint64, seed uint64, date time.Time) *necurs {
	return &necurs{
		year:  uint64(date.Year()),
		month: uint64(date.Month()),
		day:   uint64(date.Day()),
		pos:   pos,
		seed:  seed,
	}
}

// GenerateDomain Returns a Domain
func (r *necurs) GenerateDomain() string {
	var i uint64
	tlds := []string{"tj", "in", "jp", "tw", "ac", "cm", "la", "mn", "so", "sh", "sc", "nu", "nf", "mu", "ms", "mx", "ki", "im", "cx", "cc", "tv", "bz", "me", "eu", "de", "ru", "co", "su", "pw", "kz", "sx", "us", "ug", "ir", "to", "ga", "com", "net", "org", "biz", "xxx", "pro", "bit"}
	n := pseudoRandom(r.year)
	n = pseudoRandom(n + r.month + 43690)
	n = pseudoRandom(n + (r.day >> 2))
	n = pseudoRandom(n + r.pos)
	n = pseudoRandom(n + r.seed)

	domainLength := mod64(n, 15) + 7

	domain := ""

	for i = 0; i < domainLength; i++ {
		n = pseudoRandom(n + i)
		ch := mod64(n, 25) + 0x61
		domain += fmt.Sprintf("%c", ch)
		n += 0xABBEDF
		n = pseudoRandom(n)
	}
	tld := tlds[mod64(n, uint64(len(tlds)))]
	domain += "." + tld
	r.pos++
	return domain
}
