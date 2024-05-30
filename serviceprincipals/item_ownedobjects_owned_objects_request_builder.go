package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOwnedobjectsOwnedObjectsRequestBuilder provides operations to manage the ownedObjects property of the microsoft.graph.servicePrincipal entity.
type ItemOwnedobjectsOwnedObjectsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOwnedobjectsOwnedObjectsRequestBuilderGetQueryParameters directory objects that this service principal owns. Read-only. Nullable. Supports $expand, $select nested in $expand, and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
type ItemOwnedobjectsOwnedObjectsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemOwnedobjectsOwnedObjectsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOwnedobjectsOwnedObjectsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOwnedobjectsOwnedObjectsRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId provides operations to manage the ownedObjects property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemOwnedobjectsDirectoryObjectItemRequestBuilder when successful
func (m *ItemOwnedobjectsOwnedObjectsRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*ItemOwnedobjectsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewItemOwnedobjectsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOwnedobjectsOwnedObjectsRequestBuilderInternal instantiates a new ItemOwnedobjectsOwnedObjectsRequestBuilder and sets the default values.
func NewItemOwnedobjectsOwnedObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOwnedobjectsOwnedObjectsRequestBuilder) {
    m := &ItemOwnedobjectsOwnedObjectsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/ownedObjects{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemOwnedobjectsOwnedObjectsRequestBuilder instantiates a new ItemOwnedobjectsOwnedObjectsRequestBuilder and sets the default values.
func NewItemOwnedobjectsOwnedObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOwnedobjectsOwnedObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOwnedobjectsOwnedObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemOwnedobjectsCountRequestBuilder when successful
func (m *ItemOwnedobjectsOwnedObjectsRequestBuilder) Count()(*ItemOwnedobjectsCountRequestBuilder) {
    return NewItemOwnedobjectsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get directory objects that this service principal owns. Read-only. Nullable. Supports $expand, $select nested in $expand, and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOwnedobjectsOwnedObjectsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOwnedobjectsOwnedObjectsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable), nil
}
// GraphApplication casts the previous resource to application.
// returns a *ItemOwnedobjectsGraphapplicationGraphApplicationRequestBuilder when successful
func (m *ItemOwnedobjectsOwnedObjectsRequestBuilder) GraphApplication()(*ItemOwnedobjectsGraphapplicationGraphApplicationRequestBuilder) {
    return NewItemOwnedobjectsGraphapplicationGraphApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphAppRoleAssignment casts the previous resource to appRoleAssignment.
// returns a *ItemOwnedobjectsGraphapproleassignmentGraphAppRoleAssignmentRequestBuilder when successful
func (m *ItemOwnedobjectsOwnedObjectsRequestBuilder) GraphAppRoleAssignment()(*ItemOwnedobjectsGraphapproleassignmentGraphAppRoleAssignmentRequestBuilder) {
    return NewItemOwnedobjectsGraphapproleassignmentGraphAppRoleAssignmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphEndpoint casts the previous resource to endpoint.
// returns a *ItemOwnedobjectsGraphendpointGraphEndpointRequestBuilder when successful
func (m *ItemOwnedobjectsOwnedObjectsRequestBuilder) GraphEndpoint()(*ItemOwnedobjectsGraphendpointGraphEndpointRequestBuilder) {
    return NewItemOwnedobjectsGraphendpointGraphEndpointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphGroup casts the previous resource to group.
// returns a *ItemOwnedobjectsGraphgroupGraphGroupRequestBuilder when successful
func (m *ItemOwnedobjectsOwnedObjectsRequestBuilder) GraphGroup()(*ItemOwnedobjectsGraphgroupGraphGroupRequestBuilder) {
    return NewItemOwnedobjectsGraphgroupGraphGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *ItemOwnedobjectsGraphserviceprincipalGraphServicePrincipalRequestBuilder when successful
func (m *ItemOwnedobjectsOwnedObjectsRequestBuilder) GraphServicePrincipal()(*ItemOwnedobjectsGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    return NewItemOwnedobjectsGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation directory objects that this service principal owns. Read-only. Nullable. Supports $expand, $select nested in $expand, and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
// returns a *RequestInformation when successful
func (m *ItemOwnedobjectsOwnedObjectsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOwnedobjectsOwnedObjectsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemOwnedobjectsOwnedObjectsRequestBuilder when successful
func (m *ItemOwnedobjectsOwnedObjectsRequestBuilder) WithUrl(rawUrl string)(*ItemOwnedobjectsOwnedObjectsRequestBuilder) {
    return NewItemOwnedobjectsOwnedObjectsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
