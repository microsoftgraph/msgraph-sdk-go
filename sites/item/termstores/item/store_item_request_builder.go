package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/termstore"
    i74df54b361b035b3b1c15bccb664cccb4c6f3b3f9d08ea3db87bafc5850d9c5c "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets"
    i943058511484df328af302c39ee5cb2c343138e90a60de21e94bad79f3a4f322 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups"
    iab727bf89f62d932b7d66fef22c81af81d898a18ffb7aff5b451ac7d9f07b34e "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item"
    ib6b55dfc35d41e5306edb4c31882d97fa855e36e1f8d413bfb9bdb0f532f0d2c "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item"
)

// StoreItemRequestBuilder provides operations to manage the termStores property of the microsoft.graph.site entity.
type StoreItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// StoreItemRequestBuilderDeleteOptions options for Delete
type StoreItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// StoreItemRequestBuilderGetOptions options for Get
type StoreItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *StoreItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// StoreItemRequestBuilderGetQueryParameters the collection of termStores under this site.
type StoreItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// StoreItemRequestBuilderPatchOptions options for Patch
type StoreItemRequestBuilderPatchOptions struct {
    // 
    Body id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.Storeable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewStoreItemRequestBuilderInternal instantiates a new StoreItemRequestBuilder and sets the default values.
func NewStoreItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*StoreItemRequestBuilder) {
    m := &StoreItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/termStores/{store_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewStoreItemRequestBuilder instantiates a new StoreItemRequestBuilder and sets the default values.
func NewStoreItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*StoreItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStoreItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property termStores for sites
func (m *StoreItemRequestBuilder) CreateDeleteRequestInformation(options *StoreItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the collection of termStores under this site.
func (m *StoreItemRequestBuilder) CreateGetRequestInformation(options *StoreItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property termStores in sites
func (m *StoreItemRequestBuilder) CreatePatchRequestInformation(options *StoreItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property termStores for sites
func (m *StoreItemRequestBuilder) Delete(options *StoreItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of termStores under this site.
func (m *StoreItemRequestBuilder) Get(options *StoreItemRequestBuilderGetOptions)(id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.Storeable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.CreateStoreFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.Storeable), nil
}
func (m *StoreItemRequestBuilder) Groups()(*i943058511484df328af302c39ee5cb2c343138e90a60de21e94bad79f3a4f322.GroupsRequestBuilder) {
    return i943058511484df328af302c39ee5cb2c343138e90a60de21e94bad79f3a4f322.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStores.item.groups.item collection
func (m *StoreItemRequestBuilder) GroupsById(id string)(*iab727bf89f62d932b7d66fef22c81af81d898a18ffb7aff5b451ac7d9f07b34e.GroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group_id"] = id
    }
    return iab727bf89f62d932b7d66fef22c81af81d898a18ffb7aff5b451ac7d9f07b34e.NewGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property termStores in sites
func (m *StoreItemRequestBuilder) Patch(options *StoreItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *StoreItemRequestBuilder) Sets()(*i74df54b361b035b3b1c15bccb664cccb4c6f3b3f9d08ea3db87bafc5850d9c5c.SetsRequestBuilder) {
    return i74df54b361b035b3b1c15bccb664cccb4c6f3b3f9d08ea3db87bafc5850d9c5c.NewSetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStores.item.sets.item collection
func (m *StoreItemRequestBuilder) SetsById(id string)(*ib6b55dfc35d41e5306edb4c31882d97fa855e36e1f8d413bfb9bdb0f532f0d2c.SetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["set_id"] = id
    }
    return ib6b55dfc35d41e5306edb4c31882d97fa855e36e1f8d413bfb9bdb0f532f0d2c.NewSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
