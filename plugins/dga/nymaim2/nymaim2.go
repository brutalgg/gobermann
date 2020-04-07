package nymaim2

import (
	"fmt"
	"strconv"
	"time"

	"github.com/brutalgg/gobermann/pkg/hashing"
)

// adapted from https://github.com/baderj/domain_generation_algorithms
const (
	seed = "3138C81ED54AD5F8E905555A6623C9C9"
)

type nymaim2 struct {
	hashstring string
}

// SeedRNG Initalize the nymaim2 algorithm
func SeedRNG(date time.Time) *nymaim2 {
	m := hashing.Md5Hash(seed)
	s := fmt.Sprintf("%s%d%d", m, date.Year()%100, date.YearDay()-1)
	return &nymaim2{
		hashstring: hashing.Md5Hash(s),
	}
}
func (r *nymaim2) getInt64() int64 {
	v, _ := strconv.ParseInt(r.hashstring[:8], 16, 64)
	r.hashstring = hashing.Md5Hash(r.hashstring[7:])
	return v
}

// GenerateDomain Returns a Domain
func (r *nymaim2) GenerateDomain() string {
	domain := ""
	domain += firstWords[r.getInt64()%int64(len(firstWords))]
	domain += seperators[r.getInt64()%int64(len(seperators))]
	domain += secondWords[r.getInt64()%int64(len(secondWords))]
	domain += tlds[r.getInt64()%int64(len(tlds))]
	return domain
}
