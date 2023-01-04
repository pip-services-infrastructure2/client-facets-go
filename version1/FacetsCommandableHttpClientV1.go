package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type FacetsCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewFacetsCommandableHttpClientV1() *FacetsCommandableHttpClientV1 {
	return &FacetsCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/facets"),
	}
}

func (c *FacetsCommandableHttpClientV1) GetFacetsByGroup(ctx context.Context, correlationId string, group string, paging *data.PagingParams) (data.DataPage[*FacetV1], error) {
	res, err := c.CallCommand(ctx, "get_facets_by_group", correlationId, data.NewAnyValueMapFromTuples(
		"group", group,
		"paging", paging,
	))

	if err != nil {
		return *data.NewEmptyDataPage[*FacetV1](), err
	}

	return clients.HandleHttpResponse[data.DataPage[*FacetV1]](res, correlationId)
}

func (c *FacetsCommandableHttpClientV1) AddFacet(ctx context.Context, correlationId string, group string, name string) (*FacetV1, error) {
	res, err := c.CallCommand(ctx, "add_facet", correlationId, data.NewAnyValueMapFromTuples(
		"group", group,
		"name", name,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*FacetV1](res, correlationId)
}

func (c *FacetsCommandableHttpClientV1) RemoveFacet(ctx context.Context, correlationId string, group string, name string) (*FacetV1, error) {
	res, err := c.CallCommand(ctx, "remove_facet", correlationId, data.NewAnyValueMapFromTuples(
		"group", group,
		"name", name,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*FacetV1](res, correlationId)
}

func (c *FacetsCommandableHttpClientV1) DeleteFacetsByGroup(ctx context.Context, correlationId string, group string) error {
	_, err := c.CallCommand(ctx, "delete_facets_by_group", correlationId, data.NewAnyValueMapFromTuples(
		"group", group,
	))

	return err
}

func (c *FacetsCommandableHttpClientV1) Clear(ctx context.Context, correlationId string) error {
	_, err := c.CallCommand(ctx, "clear", correlationId, nil)
	return err
}
