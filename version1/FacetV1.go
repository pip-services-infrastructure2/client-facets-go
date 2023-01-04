package version1

type FacetV1 struct {
	Group string `json:"group"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func NewFacetV1(group, name string, count int) *FacetV1 {
	return &FacetV1{
		Group: group,
		Name:  name,
		Count: count,
	}
}
