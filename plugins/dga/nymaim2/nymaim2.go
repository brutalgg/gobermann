package nymaim2

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

const (
	seed = "3138C81ED54AD5F8E905555A6623C9C9"
)

// nymaim2 a
type nymaim2 struct {
	hashstring string
}

func md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// SeedRNG n
func SeedRNG(date time.Time) *nymaim2 {
	m := md5Hash(seed)
	s := fmt.Sprintf("%s%d%d", m, date.Year()%100, date.YearDay()-1)
	return &nymaim2{
		hashstring: md5Hash(s),
	}
}
func (r *nymaim2) getInt64() int64 {
	v, _ := strconv.ParseInt(r.hashstring[:8], 16, 64)
	r.hashstring = md5Hash(r.hashstring[7:])
	return v
}

// GenerateDomain a
func (r *nymaim2) GenerateDomain() string {
	domain := ""
	domain += firstWords[r.getInt64()%int64(len(firstWords))]
	domain += seperators[r.getInt64()%int64(len(seperators))]
	domain += secondWords[r.getInt64()%int64(len(secondWords))]
	domain += tlds[r.getInt64()%int64(len(tlds))]
	return domain
}
