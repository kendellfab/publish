package domain

type Pagination struct {
	HasOlder   bool
	OlderIndex int
	HasNewer   bool
	NewerIndex int
	Offset     int
	Count      int
	Total      int
}

func (p Pagination) Both() bool {
	return p.HasNewer && p.HasOlder
}
