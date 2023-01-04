package test_version1

import (
	"testing"

	"github.com/pip-services-infrastructure2/client-facets-go/version1"
)

type facetsMockClientV1Test struct {
	client  *version1.FacetsMockClientV1
	fixture *FacetsClientFixtureV1
}

func newFacetsMockClientV1Test() *facetsMockClientV1Test {
	return &facetsMockClientV1Test{}
}

func (c *facetsMockClientV1Test) setup(t *testing.T) {
	c.client = version1.NewFacetsMockClientV1()
	c.fixture = NewFacetsClientFixtureV1(c.client)
}

func (c *facetsMockClientV1Test) teardown(t *testing.T) {
	c.client = nil
}

func TestMockAddAndRemoveFacets(t *testing.T) {
	c := newFacetsMockClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestAddAndRemoveFacets(t)
}
