package monerodownloader

import (
	"fmt"
	"time"

	"github.com/brutalgg/gobermann/pkg/hashing"
)

const (
	seed   = "jkhhksugrhtijys78g46"
	static = "31b4bd31fg1x2"
)

type monerodownloader struct {
	tldPos int
	pos    int
	days   int
}

func dateConv(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

// SeedRNG asdf
func SeedRNG(pos int, date time.Time) *monerodownloader {
	epoch := dateConv(1970, 1, 1)
	days := int(date.Sub(epoch).Hours()/24) - 1
	return &monerodownloader{
		tldPos: 0,
		pos:    0,
		days:   days,
	}
}

func (r *monerodownloader) GenerateDomain() string {
	var domain string
	tlds := []string{".org", ".tickets", ".blackfriday", ".hosting", ".feedback"}

	s := fmt.Sprintf("%v-%v-%v", seed, r.days, r.pos)
	m := hashing.Md5Hash(s)
	if r.pos == 0 {
		domain = fmt.Sprintf("%v%v", static, tlds[r.tldPos])
	} else {
		domain = fmt.Sprintf("%v%v", m[:13], tlds[r.tldPos])
	}

	if r.tldPos >= len(tlds)-1 {
		r.tldPos = 0
		r.pos++
		// r.days-- // could be used to get domains from previous days
	} else {
		r.tldPos++
	}

	return domain
}
