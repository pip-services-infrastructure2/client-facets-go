package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type FacetsNullClientV1 struct {
}

func NewFacetsNullClientV1() *FacetsNullClientV1 {
	return &FacetsNullClientV1{}
}

func (c *FacetsNullClientV1) GetFacetsByGroup(ctx context.Context, correlationId string, group string, paging *data.PagingParams) (data.DataPage[*FacetV1], error) {
	return *data.NewEmptyDataPage[*FacetV1](), nil
}

func (c *FacetsNullClientV1) AddFacet(ctx context.Context, correlationId string, group string, name string) (*FacetV1, error) {
	return NewFacetV1(group, name, 0), nil
}

func (c *FacetsNullClientV1) RemoveFacet(ctx context.Context, correlationId string, group string, name string) (*FacetV1, error) {
	return NewFacetV1(group, name, 0), nil
}

func (c *FacetsNullClientV1) DeleteFacetsByGroup(ctx context.Context, correlationId string, group string) error {
	return nil
}

func (c *FacetsNullClientV1) Clear(ctx context.Context, correlationId string) error {
	return nil
}
