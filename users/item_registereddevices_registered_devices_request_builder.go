package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemRegistereddevicesRegisteredDevicesRequestBuilder provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
type ItemRegistereddevicesRegisteredDevicesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRegistereddevicesRegisteredDevicesRequestBuilderGetQueryParameters devices that are registered for the user. Read-only. Nullable. Supports $expand and returns up to 100 objects.
type ItemRegistereddevicesRegisteredDevicesRequestBuilderGetQueryParameters struct {
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
// ItemRegistereddevicesRegisteredDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRegistereddevicesRegisteredDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRegistereddevicesRegisteredDevicesRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
// returns a *ItemRegistereddevicesDirectoryObjectItemRequestBuilder when successful
func (m *ItemRegistereddevicesRegisteredDevicesRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*ItemRegistereddevicesDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewItemRegistereddevicesDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemRegistereddevicesRegisteredDevicesRequestBuilderInternal instantiates a new ItemRegistereddevicesRegisteredDevicesRequestBuilder and sets the default values.
func NewItemRegistereddevicesRegisteredDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegistereddevicesRegisteredDevicesRequestBuilder) {
    m := &ItemRegistereddevicesRegisteredDevicesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/registeredDevices{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemRegistereddevicesRegisteredDevicesRequestBuilder instantiates a new ItemRegistereddevicesRegisteredDevicesRequestBuilder and sets the default values.
func NewItemRegistereddevicesRegisteredDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegistereddevicesRegisteredDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRegistereddevicesRegisteredDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemRegistereddevicesCountRequestBuilder when successful
func (m *ItemRegistereddevicesRegisteredDevicesRequestBuilder) Count()(*ItemRegistereddevicesCountRequestBuilder) {
    return NewItemRegistereddevicesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get devices that are registered for the user. Read-only. Nullable. Supports $expand and returns up to 100 objects.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRegistereddevicesRegisteredDevicesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRegistereddevicesRegisteredDevicesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
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
// returns a *ItemRegistereddevicesGraphapproleassignmentGraphAppRoleAssignmentRequestBuilder when successful
func (m *ItemRegistereddevicesRegisteredDevicesRequestBuilder) GraphAppRoleAssignment()(*ItemRegistereddevicesGraphapproleassignmentGraphAppRoleAssignmentRequestBuilder) {
    return NewItemRegistereddevicesGraphapproleassignmentGraphAppRoleAssignmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphDevice casts the previous resource to device.
// returns a *ItemRegistereddevicesGraphdeviceGraphDeviceRequestBuilder when successful
func (m *ItemRegistereddevicesRegisteredDevicesRequestBuilder) GraphDevice()(*ItemRegistereddevicesGraphdeviceGraphDeviceRequestBuilder) {
    return NewItemRegistereddevicesGraphdeviceGraphDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphEndpoint casts the previous resource to endpoint.
// returns a *ItemRegistereddevicesGraphendpointGraphEndpointRequestBuilder when successful
func (m *ItemRegistereddevicesRegisteredDevicesRequestBuilder) GraphEndpoint()(*ItemRegistereddevicesGraphendpointGraphEndpointRequestBuilder) {
    return NewItemRegistereddevicesGraphendpointGraphEndpointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation devices that are registered for the user. Read-only. Nullable. Supports $expand and returns up to 100 objects.
// returns a *RequestInformation when successful
func (m *ItemRegistereddevicesRegisteredDevicesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRegistereddevicesRegisteredDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemRegistereddevicesRegisteredDevicesRequestBuilder when successful
func (m *ItemRegistereddevicesRegisteredDevicesRequestBuilder) WithUrl(rawUrl string)(*ItemRegistereddevicesRegisteredDevicesRequestBuilder) {
    return NewItemRegistereddevicesRegisteredDevicesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
