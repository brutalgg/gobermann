package domaingeneratingalgorithm

// DomainGenerator a
type DomainGenerator interface {
	GenerateDomain() string
}

// DefaultGenerator s
type DefaultGenerator struct{}

// GenerateDomain s
func (b DefaultGenerator) GenerateDomain() string {
	return ""
}
