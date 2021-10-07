package goshopify

import (
	"fmt"
	"time"
)

const inventoryLevelsBasePath = "inventory_levels"

// InventoryLevelService is an interface for interacting with the
// inventory levels endpoints of the Shopify API
// See https://help.shopify.com/en/api/reference/inventory/inventorylevel
type InventoryLevelService interface {
	List(interface{}) ([]InventoryLevel, error)
	ListWithPagination(interface{}) ([]InventoryLevel, *Pagination, error)
}

type InventoryLevelListOptions struct {
	ListOptions

	InventoryItemIDs []int64 `url:"inventory_item_ids,omitempty,comma"`
	LocationIDs      []int64 `url:"location_ids,omitempty,comma"`
}

// InventoryLevelServiceOp is the default implementation of the InventoryLevelService interface
type InventoryLevelServiceOp struct {
	client *Client
}

// InventoryLevel represents a Shopify inventory level
type InventoryLevel struct {
	InventoryItemID int64      `json:"inventory_item_id,omitempty"`
	LocationID      int64      `json:"location_id,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	Available       int64      `json:"available,omitempty"`
}

// InventoryLevelResource is used for handling single level requests and responses
type InventoryLevelResource struct {
	InventoryLevel *InventoryLevel `json:"inventory_level"`
}

// InventoryLevelsResource is used for handling multiple level responsees
type InventoryLevelsResource struct {
	InventoryLevels []InventoryLevel `json:"inventory_levels"`
}

// List inventory levels
func (s *InventoryLevelServiceOp) List(options interface{}) ([]InventoryLevel, error) {
	levels, _, err := s.ListWithPagination(options)
	if err != nil {
		return nil, err
	}
	return levels, nil
}

func (s *InventoryLevelServiceOp) ListWithPagination(options interface{}) ([]InventoryLevel, *Pagination, error) {
	path := fmt.Sprintf("%s.json", inventoryLevelsBasePath)
	resource := new(InventoryLevelsResource)

	headers, err := s.client.createAndDoGetHeaders("GET", path, nil, options, resource)
	if err != nil {
		return nil, nil, err
	}

	// Extract pagination info from header
	linkHeader := headers.Get("Link")

	pagination, err := extractPagination(linkHeader)
	if err != nil {
		return nil, nil, err
	}

	return resource.InventoryLevels, pagination, nil
}
