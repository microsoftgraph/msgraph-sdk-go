package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOwneddevicesOwnedDevicesRequestBuilder provides operations to manage the ownedDevices property of the microsoft.graph.user entity.
type ItemOwneddevicesOwnedDevicesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOwneddevicesOwnedDevicesRequestBuilderGetQueryParameters devices that are owned by the user. Read-only. Nullable. Supports $expand and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
type ItemOwneddevicesOwnedDevicesRequestBuilderGetQueryParameters struct {
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
// ItemOwneddevicesOwnedDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOwneddevicesOwnedDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOwneddevicesOwnedDevicesRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId provides operations to manage the ownedDevices property of the microsoft.graph.user entity.
// returns a *ItemOwneddevicesDirectoryObjectItemRequestBuilder when successful
func (m *ItemOwneddevicesOwnedDevicesRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*ItemOwneddevicesDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewItemOwneddevicesDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOwneddevicesOwnedDevicesRequestBuilderInternal instantiates a new ItemOwneddevicesOwnedDevicesRequestBuilder and sets the default values.
func NewItemOwneddevicesOwnedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOwneddevicesOwnedDevicesRequestBuilder) {
    m := &ItemOwneddevicesOwnedDevicesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/ownedDevices{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemOwneddevicesOwnedDevicesRequestBuilder instantiates a new ItemOwneddevicesOwnedDevicesRequestBuilder and sets the default values.
func NewItemOwneddevicesOwnedDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOwneddevicesOwnedDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOwneddevicesOwnedDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemOwneddevicesCountRequestBuilder when successful
func (m *ItemOwneddevicesOwnedDevicesRequestBuilder) Count()(*ItemOwneddevicesCountRequestBuilder) {
    return NewItemOwneddevicesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get devices that are owned by the user. Read-only. Nullable. Supports $expand and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOwneddevicesOwnedDevicesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOwneddevicesOwnedDevicesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
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
// GraphAppRoleAssignment casts the previous resource to appRoleAssignment.
// returns a *ItemOwneddevicesGraphapproleassignmentGraphAppRoleAssignmentRequestBuilder when successful
func (m *ItemOwneddevicesOwnedDevicesRequestBuilder) GraphAppRoleAssignment()(*ItemOwneddevicesGraphapproleassignmentGraphAppRoleAssignmentRequestBuilder) {
    return NewItemOwneddevicesGraphapproleassignmentGraphAppRoleAssignmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphDevice casts the previous resource to device.
// returns a *ItemOwneddevicesGraphdeviceGraphDeviceRequestBuilder when successful
func (m *ItemOwneddevicesOwnedDevicesRequestBuilder) GraphDevice()(*ItemOwneddevicesGraphdeviceGraphDeviceRequestBuilder) {
    return NewItemOwneddevicesGraphdeviceGraphDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphEndpoint casts the previous resource to endpoint.
// returns a *ItemOwneddevicesGraphendpointGraphEndpointRequestBuilder when successful
func (m *ItemOwneddevicesOwnedDevicesRequestBuilder) GraphEndpoint()(*ItemOwneddevicesGraphendpointGraphEndpointRequestBuilder) {
    return NewItemOwneddevicesGraphendpointGraphEndpointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation devices that are owned by the user. Read-only. Nullable. Supports $expand and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
// returns a *RequestInformation when successful
func (m *ItemOwneddevicesOwnedDevicesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOwneddevicesOwnedDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOwneddevicesOwnedDevicesRequestBuilder when successful
func (m *ItemOwneddevicesOwnedDevicesRequestBuilder) WithUrl(rawUrl string)(*ItemOwneddevicesOwnedDevicesRequestBuilder) {
    return NewItemOwneddevicesOwnedDevicesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
