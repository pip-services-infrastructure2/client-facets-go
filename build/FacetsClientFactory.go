package build

import (
	clients1 "github.com/pip-services-infrastructure2/client-facets-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type EmailClientFactory struct {
	*cbuild.Factory
}

func NewEmailClientFactory() *EmailClientFactory {
	c := &EmailClientFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-facets", "client", "null", "*", "1.0")
	mockClientDescriptor := cref.NewDescriptor("service-facets", "client", "mock", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-facets", "client", "commandable-http", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewFacetsNullClientV1)
	c.RegisterType(mockClientDescriptor, clients1.NewFacetsMockClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewFacetsCommandableHttpClientV1)

	return c
}
