package devices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemRegisteredownersRegisteredOwnersRequestBuilder provides operations to manage the registeredOwners property of the microsoft.graph.device entity.
type ItemRegisteredownersRegisteredOwnersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRegisteredownersRegisteredOwnersRequestBuilderGetQueryParameters retrieve a list of users that are registered owners of the device. A registered owner is the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner.
type ItemRegisteredownersRegisteredOwnersRequestBuilderGetQueryParameters struct {
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
// ItemRegisteredownersRegisteredOwnersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRegisteredownersRegisteredOwnersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRegisteredownersRegisteredOwnersRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.devices.item.registeredOwners.item collection
// returns a *ItemRegisteredownersDirectoryObjectItemRequestBuilder when successful
func (m *ItemRegisteredownersRegisteredOwnersRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*ItemRegisteredownersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewItemRegisteredownersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemRegisteredownersRegisteredOwnersRequestBuilderInternal instantiates a new ItemRegisteredownersRegisteredOwnersRequestBuilder and sets the default values.
func NewItemRegisteredownersRegisteredOwnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredownersRegisteredOwnersRequestBuilder) {
    m := &ItemRegisteredownersRegisteredOwnersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/devices/{device%2Did}/registeredOwners{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemRegisteredownersRegisteredOwnersRequestBuilder instantiates a new ItemRegisteredownersRegisteredOwnersRequestBuilder and sets the default values.
func NewItemRegisteredownersRegisteredOwnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredownersRegisteredOwnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRegisteredownersRegisteredOwnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemRegisteredownersCountRequestBuilder when successful
func (m *ItemRegisteredownersRegisteredOwnersRequestBuilder) Count()(*ItemRegisteredownersCountRequestBuilder) {
    return NewItemRegisteredownersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of users that are registered owners of the device. A registered owner is the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/device-list-registeredowners?view=graph-rest-1.0
func (m *ItemRegisteredownersRegisteredOwnersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRegisteredownersRegisteredOwnersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
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
// returns a *ItemRegisteredownersGraphapproleassignmentGraphAppRoleAssignmentRequestBuilder when successful
func (m *ItemRegisteredownersRegisteredOwnersRequestBuilder) GraphAppRoleAssignment()(*ItemRegisteredownersGraphapproleassignmentGraphAppRoleAssignmentRequestBuilder) {
    return NewItemRegisteredownersGraphapproleassignmentGraphAppRoleAssignmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphEndpoint casts the previous resource to endpoint.
// returns a *ItemRegisteredownersGraphendpointGraphEndpointRequestBuilder when successful
func (m *ItemRegisteredownersRegisteredOwnersRequestBuilder) GraphEndpoint()(*ItemRegisteredownersGraphendpointGraphEndpointRequestBuilder) {
    return NewItemRegisteredownersGraphendpointGraphEndpointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *ItemRegisteredownersGraphserviceprincipalGraphServicePrincipalRequestBuilder when successful
func (m *ItemRegisteredownersRegisteredOwnersRequestBuilder) GraphServicePrincipal()(*ItemRegisteredownersGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    return NewItemRegisteredownersGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *ItemRegisteredownersGraphuserGraphUserRequestBuilder when successful
func (m *ItemRegisteredownersRegisteredOwnersRequestBuilder) GraphUser()(*ItemRegisteredownersGraphuserGraphUserRequestBuilder) {
    return NewItemRegisteredownersGraphuserGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ref provides operations to manage the collection of device entities.
// returns a *ItemRegisteredownersRefRequestBuilder when successful
func (m *ItemRegisteredownersRegisteredOwnersRequestBuilder) Ref()(*ItemRegisteredownersRefRequestBuilder) {
    return NewItemRegisteredownersRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation retrieve a list of users that are registered owners of the device. A registered owner is the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner.
// returns a *RequestInformation when successful
func (m *ItemRegisteredownersRegisteredOwnersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRegisteredownersRegisteredOwnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemRegisteredownersRegisteredOwnersRequestBuilder when successful
func (m *ItemRegisteredownersRegisteredOwnersRequestBuilder) WithUrl(rawUrl string)(*ItemRegisteredownersRegisteredOwnersRequestBuilder) {
    return NewItemRegisteredownersRegisteredOwnersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
