package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-infrastructure2/client-facets-go/version1"
	"github.com/stretchr/testify/assert"
)

type FacetsClientFixtureV1 struct {
	Client version1.IFacetsClientV1
}

func NewFacetsClientFixtureV1(client version1.IFacetsClientV1) *FacetsClientFixtureV1 {
	return &FacetsClientFixtureV1{
		Client: client,
	}
}

func (c *FacetsClientFixtureV1) TestAddAndRemoveFacets(t *testing.T) {
	// Add facet 1

	facet, err := c.Client.AddFacet(context.Background(), "123", "test", "group1")
	assert.Nil(t, err)

	assert.Equal(t, facet.Group, "test")
	assert.Equal(t, facet.Name, "group1")
	assert.Equal(t, facet.Count, 1)

	// Remove facet 1
	facet, err = c.Client.RemoveFacet(context.Background(), "123", "test", "group2")
	assert.Nil(t, err)

	assert.Equal(t, facet.Group, "test")
	assert.Equal(t, facet.Name, "group2")
	assert.Equal(t, facet.Count, 0)

	// Read facets
	page, err := c.Client.GetFacetsByGroup(context.Background(), "123", "test", nil)
	assert.Nil(t, err)

	assert.Len(t, page.Data, 1)

	// Delete facets
	err = c.Client.DeleteFacetsByGroup(context.Background(), "123", "test")
	assert.Nil(t, err)

	// Read facets
	page, err = c.Client.GetFacetsByGroup(context.Background(), "123", "test", nil)
	assert.Nil(t, err)

	assert.Len(t, page.Data, 0)

}
