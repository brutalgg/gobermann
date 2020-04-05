package locky

import (
	"fmt"
	"time"
)

const (
	cf = 0xFFFFFFFF
	c1 = 0xB11924E1
	c2 = 0x27100001
)

func ror32(v uint64, s uint64) uint64 {
	v &= cf
	return ((v >> s) | (v << (32 - s))) & cf
}

func rol32(v uint64, s uint64) uint64 {
	v &= cf
	return ((v << s) | (v >> (32 - s))) & cf
}

type locky struct {
	pos    uint64
	config config
	year   uint64
	month  uint64
	day    uint64
}

// SeedRNG Initialize the locky algorithm
func SeedRNG(pos uint64, configNum int, date time.Time) *locky {
	return &locky{
		config: configs[configNum],
		pos:    pos,
		year:   uint64(date.Year()),
		month:  uint64(date.Month()),
		day:    uint64(date.Day()),
	}
}

// GenerateDomain Returns a Domain
func (r *locky) GenerateDomain() string {
	var k, i uint64

	seedShifted := rol32(r.config.seed, 17)
	posShifted := rol32(r.pos, 21)

	k = 0
	year := r.year
	for i := 0; i < 7; i++ {
		t0 := ror32(c1*(year+k+0x1BF5), r.config.shift) & cf
		t1 := ((t0 + c2) ^ k) & cf
		t2 := (ror32(c1*(t1+r.config.seed), r.config.shift)) & cf
		t3 := ((t2 + c2) ^ t1) & cf
		t4 := (ror32(c1*(r.day/2+t3), r.config.shift)) & cf
		t5 := (0xD8EFFFFF - t4 + t3) & cf
		t6 := (ror32(c1*(r.month+t5-0x65CAD), r.config.shift)) & cf
		t7 := (t5 + t6 + c2) & cf
		t8 := (ror32(c1*(t7+seedShifted+posShifted), r.config.shift)) & cf
		k = ((t8 + c2) ^ t7) & cf
		year++
	}
	length := (k % 11) + 7
	domain := ""
	for i = 0; i < length; i++ {
		k = (ror32(c1*rol32(k, i), r.config.shift) + c2) & cf
		domain += fmt.Sprintf("%c", k%25+0x61)
	}
	domain += "."
	k = ror32(k*c1, r.config.shift)
	tldPos := ((k + c2) & cf) % uint64(len(r.config.topLevelDomains))
	domain += r.config.topLevelDomains[tldPos]
	r.pos++
	return domain
}
