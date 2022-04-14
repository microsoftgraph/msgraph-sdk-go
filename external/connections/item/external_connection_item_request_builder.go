package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc "github.com/microsoftgraph/msgraph-sdk-go/models/externalconnectors"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i274d667d8ec0c3d039333dd12b5455aea280c5cf9f73bb6338166f3c8da9dfaf "github.com/microsoftgraph/msgraph-sdk-go/external/connections/item/groups"
    i441e9592f00d41e5a711ec006f03afa66dec29cfee128954ccf7c3239af7d91e "github.com/microsoftgraph/msgraph-sdk-go/external/connections/item/items"
    i970b7049d0cc43d708fa2e7d007e8e10fa290eecc038a57584def6b75c0b7cb9 "github.com/microsoftgraph/msgraph-sdk-go/external/connections/item/schema"
    id2fc2789ff7ccaf3c3f408ae8fe33ddf43f1a9d079f3962d4b053c44e24b1121 "github.com/microsoftgraph/msgraph-sdk-go/external/connections/item/operations"
    i64f7fe1781849551f0747409711376a931e53b1189f67fcc2365601b3c393b50 "github.com/microsoftgraph/msgraph-sdk-go/external/connections/item/groups/item"
    ib280bd1867c7c5e1ca1ef4564e9994b09cb6e6d7a7ffd961eaa450655177eefe "github.com/microsoftgraph/msgraph-sdk-go/external/connections/item/operations/item"
    icaf66d026d8a409b8012b31d621993f6257dc2a500c4e095c199593a9b6dea77 "github.com/microsoftgraph/msgraph-sdk-go/external/connections/item/items/item"
)

// ExternalConnectionItemRequestBuilder provides operations to manage the connections property of the microsoft.graph.externalConnectors.external entity.
type ExternalConnectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ExternalConnectionItemRequestBuilderDeleteOptions options for Delete
type ExternalConnectionItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ExternalConnectionItemRequestBuilderGetOptions options for Get
type ExternalConnectionItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExternalConnectionItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ExternalConnectionItemRequestBuilderGetQueryParameters get connections from external
type ExternalConnectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExternalConnectionItemRequestBuilderPatchOptions options for Patch
type ExternalConnectionItemRequestBuilderPatchOptions struct {
    // 
    Body i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.ExternalConnectionable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewExternalConnectionItemRequestBuilderInternal instantiates a new ExternalConnectionItemRequestBuilder and sets the default values.
func NewExternalConnectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExternalConnectionItemRequestBuilder) {
    m := &ExternalConnectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/external/connections/{externalConnection%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExternalConnectionItemRequestBuilder instantiates a new ExternalConnectionItemRequestBuilder and sets the default values.
func NewExternalConnectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExternalConnectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExternalConnectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property connections for external
func (m *ExternalConnectionItemRequestBuilder) CreateDeleteRequestInformation(options *ExternalConnectionItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get connections from external
func (m *ExternalConnectionItemRequestBuilder) CreateGetRequestInformation(options *ExternalConnectionItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property connections in external
func (m *ExternalConnectionItemRequestBuilder) CreatePatchRequestInformation(options *ExternalConnectionItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property connections for external
func (m *ExternalConnectionItemRequestBuilder) Delete(options *ExternalConnectionItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get connections from external
func (m *ExternalConnectionItemRequestBuilder) Get(options *ExternalConnectionItemRequestBuilderGetOptions)(i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.ExternalConnectionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.CreateExternalConnectionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.ExternalConnectionable), nil
}
// Groups the groups property
func (m *ExternalConnectionItemRequestBuilder) Groups()(*i274d667d8ec0c3d039333dd12b5455aea280c5cf9f73bb6338166f3c8da9dfaf.GroupsRequestBuilder) {
    return i274d667d8ec0c3d039333dd12b5455aea280c5cf9f73bb6338166f3c8da9dfaf.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.external.connections.item.groups.item collection
func (m *ExternalConnectionItemRequestBuilder) GroupsById(id string)(*i64f7fe1781849551f0747409711376a931e53b1189f67fcc2365601b3c393b50.ExternalGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalGroup%2Did"] = id
    }
    return i64f7fe1781849551f0747409711376a931e53b1189f67fcc2365601b3c393b50.NewExternalGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Items the items property
func (m *ExternalConnectionItemRequestBuilder) Items()(*i441e9592f00d41e5a711ec006f03afa66dec29cfee128954ccf7c3239af7d91e.ItemsRequestBuilder) {
    return i441e9592f00d41e5a711ec006f03afa66dec29cfee128954ccf7c3239af7d91e.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.external.connections.item.items.item collection
func (m *ExternalConnectionItemRequestBuilder) ItemsById(id string)(*icaf66d026d8a409b8012b31d621993f6257dc2a500c4e095c199593a9b6dea77.ExternalItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalItem%2Did"] = id
    }
    return icaf66d026d8a409b8012b31d621993f6257dc2a500c4e095c199593a9b6dea77.NewExternalItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *ExternalConnectionItemRequestBuilder) Operations()(*id2fc2789ff7ccaf3c3f408ae8fe33ddf43f1a9d079f3962d4b053c44e24b1121.OperationsRequestBuilder) {
    return id2fc2789ff7ccaf3c3f408ae8fe33ddf43f1a9d079f3962d4b053c44e24b1121.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.external.connections.item.operations.item collection
func (m *ExternalConnectionItemRequestBuilder) OperationsById(id string)(*ib280bd1867c7c5e1ca1ef4564e9994b09cb6e6d7a7ffd961eaa450655177eefe.ConnectionOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connectionOperation%2Did"] = id
    }
    return ib280bd1867c7c5e1ca1ef4564e9994b09cb6e6d7a7ffd961eaa450655177eefe.NewConnectionOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property connections in external
func (m *ExternalConnectionItemRequestBuilder) Patch(options *ExternalConnectionItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Schema the schema property
func (m *ExternalConnectionItemRequestBuilder) Schema()(*i970b7049d0cc43d708fa2e7d007e8e10fa290eecc038a57584def6b75c0b7cb9.SchemaRequestBuilder) {
    return i970b7049d0cc43d708fa2e7d007e8e10fa290eecc038a57584def6b75c0b7cb9.NewSchemaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
