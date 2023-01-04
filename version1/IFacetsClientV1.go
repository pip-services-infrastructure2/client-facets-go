package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IFacetsClientV1 interface {
	GetFacetsByGroup(ctx context.Context, correlationId string, group string, paging *data.PagingParams) (data.DataPage[*FacetV1], error)

	AddFacet(ctx context.Context, correlationId string, group string, name string) (*FacetV1, error)

	RemoveFacet(ctx context.Context, correlationId string, group string, name string) (*FacetV1, error)

	DeleteFacetsByGroup(ctx context.Context, correlationId string, group string) error

	Clear(ctx context.Context, correlationId string) error
}
