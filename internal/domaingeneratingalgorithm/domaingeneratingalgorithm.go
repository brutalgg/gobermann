package domaingeneratingalgorithm

// DomainGenerator Domain Generator interface for all domain generating algorithms to implement
type DomainGenerator interface {
	GenerateDomain() string
}

// DefaultGenerator An empty domain generator interface
type DefaultGenerator struct{}

// GenerateDomain Returns empty string
func (b DefaultGenerator) GenerateDomain() string {
	return ""
}
