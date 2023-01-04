package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type FacetsMockClientV1 struct {
	items       []*FacetV1
	maxPageSize int
}

func NewFacetsMockClientV1() *FacetsMockClientV1 {
	return &FacetsMockClientV1{
		items:       make([]*FacetV1, 0),
		maxPageSize: 100,
	}
}

func (c *FacetsMockClientV1) GetFacetsByGroup(ctx context.Context, correlationId string, group string, paging *data.PagingParams) (data.DataPage[*FacetV1], error) {
	items := make([]*FacetV1, 0)
	for _, item := range c.items {
		if item.Group == group && item.Count > 0 {
			buf := *item
			items = append(items, &buf)
		}
	}

	// Extract a page
	if paging == nil {
		paging = data.NewEmptyPagingParams()
	}

	skip := paging.GetSkip(-1)
	take := paging.GetTake(int64(c.maxPageSize))

	total := 0
	if paging.Total {
		total = len(items)
	}

	if skip > 0 {
		if int(skip) > len(items) {
			items = items[len(items):]
		}
	}

	if int(take) < len(items) {
		items = items[:take]
	}

	return *data.NewDataPage(items, total), nil

}

func (c *FacetsMockClientV1) AddFacet(ctx context.Context, correlationId string, group string, name string) (*FacetV1, error) {
	var item *FacetV1
	for _, el := range c.items {
		if item.Group == group && item.Name == name {
			buf := *el
			item = &buf
			break
		}
	}

	if item != nil {
		item.Count++
	} else {
		item = NewFacetV1(group, name, 1)
		c.items = append(c.items, item)
	}

	buf := *item
	return &buf, nil
}

func (c *FacetsMockClientV1) RemoveFacet(ctx context.Context, correlationId string, group string, name string) (*FacetV1, error) {
	var item *FacetV1
	for _, el := range c.items {
		if el.Group == group && el.Name == name {
			buf := *el
			item = &buf
			break
		}
	}

	if item != nil {
		if item.Count > 0 {
			item.Count--
		} else {
			item.Count = 0
		}
	} else {
		item = NewFacetV1(group, name, 0)
		c.items = append(c.items, item)
	}

	buf := *item
	return &buf, nil
}

func (c *FacetsMockClientV1) DeleteFacetsByGroup(ctx context.Context, correlationId string, group string) error {
	for index, item := range c.items {
		if item.Group == group {
			if index < len(c.items) {
				c.items = append(c.items[:index], c.items[index+1:]...)
			} else {
				c.items = c.items[:index]
			}

		}
	}

	return nil
}

func (c *FacetsMockClientV1) Clear(ctx context.Context, correlationId string) error {
	c.items = make([]*FacetV1, 0)
	return nil
}
