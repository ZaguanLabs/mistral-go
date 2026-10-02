package sdk

import (
	"fmt"
	"net/http"
	"net/url"
)

type ListManagedIndexesParams struct {
	Name      *string
	Status    *ManagedIndexStatus
	CreatorID *string
	PageSize  *int
	PageToken *string
}

func managedIndexPath(name string) string { return "v1/rag/managed_indexes/" + url.PathEscape(name) }

func (c *MistralClient) CreateManagedIndex(req *CreateManagedIndexRequest) (*ManagedIndexResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	return requestTyped[ManagedIndexResponse](c, http.MethodPost, req, "v1/rag/managed_indexes")
}
func (c *MistralClient) ListManagedIndexes(params *ListManagedIndexesParams) (*ListManagedIndexesResponse, error) {
	if params == nil {
		params = &ListManagedIndexesParams{}
	}
	size := 20
	var token *string
	if params != nil {
		if params.PageSize != nil {
			size = *params.PageSize
		}
		token = params.PageToken
	}
	return requestTyped[ListManagedIndexesResponse](c, http.MethodGet, nil, appendQuery("v1/rag/managed_indexes", queryWithOptionalValues(map[string]any{"page_size": size, "page_token": token, "name": params.Name, "status": enumString(params.Status), "creator_id": params.CreatorID})))
}
func (c *MistralClient) GetManagedIndex(name string) (*ManagedIndexResponse, error) {
	return requestTyped[ManagedIndexResponse](c, http.MethodGet, nil, managedIndexPath(name))
}
func (c *MistralClient) UpdateManagedIndex(name string, req *UpdateManagedIndexRequest) (*ManagedIndexResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	return requestTyped[ManagedIndexResponse](c, http.MethodPut, req, managedIndexPath(name))
}
func (c *MistralClient) DeleteManagedIndex(name string) (*DeleteManagedIndexResponse, error) {
	return requestTyped[DeleteManagedIndexResponse](c, http.MethodDelete, nil, managedIndexPath(name))
}
func (c *MistralClient) IngestManagedIndexDocuments(name string, req *IngestDocumentsRequest) (*IngestDocumentsResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	return requestTyped[IngestDocumentsResponse](c, http.MethodPost, req, managedIndexPath(name)+"/documents")
}
func (c *MistralClient) DeleteManagedIndexDocuments(name string, req *DeleteDocumentsRequest) (*DeleteDocumentsResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	return requestTyped[DeleteDocumentsResponse](c, http.MethodDelete, req, managedIndexPath(name)+"/documents")
}
func (c *MistralClient) SearchManagedIndex(name string, req *SearchRequest) (*SearchResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	return requestTyped[SearchResponse](c, http.MethodPost, req, managedIndexPath(name)+"/search")
}
