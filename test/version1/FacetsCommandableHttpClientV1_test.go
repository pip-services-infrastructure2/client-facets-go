package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-infrastructure2/client-facets-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type facetsCommandableHttpClientV1Test struct {
	client  *version1.FacetsCommandableHttpClientV1
	fixture *FacetsClientFixtureV1
}

func newFacetsCommandableHttpClientV1Test() *facetsCommandableHttpClientV1Test {
	return &facetsCommandableHttpClientV1Test{}
}

func (c *facetsCommandableHttpClientV1Test) setup(t *testing.T) {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewFacetsCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewFacetsClientFixtureV1(c.client)
}

func (c *facetsCommandableHttpClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpAddAndRemoveFacets(t *testing.T) {
	c := newFacetsCommandableHttpClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestAddAndRemoveFacets(t)
}
