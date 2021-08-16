package store

type Store interface {
	Open(string) error
	GetLeadRep() LeadRep
}
