package devices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemRegisteredusersRegisteredUsersRequestBuilder provides operations to manage the registeredUsers property of the microsoft.graph.device entity.
type ItemRegisteredusersRegisteredUsersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRegisteredusersRegisteredUsersRequestBuilderGetQueryParameters retrieve a list of users that are registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration.
type ItemRegisteredusersRegisteredUsersRequestBuilderGetQueryParameters struct {
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
// ItemRegisteredusersRegisteredUsersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRegisteredusersRegisteredUsersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRegisteredusersRegisteredUsersRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.devices.item.registeredUsers.item collection
// returns a *ItemRegisteredusersDirectoryObjectItemRequestBuilder when successful
func (m *ItemRegisteredusersRegisteredUsersRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*ItemRegisteredusersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewItemRegisteredusersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemRegisteredusersRegisteredUsersRequestBuilderInternal instantiates a new ItemRegisteredusersRegisteredUsersRequestBuilder and sets the default values.
func NewItemRegisteredusersRegisteredUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredusersRegisteredUsersRequestBuilder) {
    m := &ItemRegisteredusersRegisteredUsersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/devices/{device%2Did}/registeredUsers{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemRegisteredusersRegisteredUsersRequestBuilder instantiates a new ItemRegisteredusersRegisteredUsersRequestBuilder and sets the default values.
func NewItemRegisteredusersRegisteredUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredusersRegisteredUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRegisteredusersRegisteredUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemRegisteredusersCountRequestBuilder when successful
func (m *ItemRegisteredusersRegisteredUsersRequestBuilder) Count()(*ItemRegisteredusersCountRequestBuilder) {
    return NewItemRegisteredusersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of users that are registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/device-list-registeredusers?view=graph-rest-1.0
func (m *ItemRegisteredusersRegisteredUsersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRegisteredusersRegisteredUsersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
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
// returns a *ItemRegisteredusersGraphapproleassignmentGraphAppRoleAssignmentRequestBuilder when successful
func (m *ItemRegisteredusersRegisteredUsersRequestBuilder) GraphAppRoleAssignment()(*ItemRegisteredusersGraphapproleassignmentGraphAppRoleAssignmentRequestBuilder) {
    return NewItemRegisteredusersGraphapproleassignmentGraphAppRoleAssignmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphEndpoint casts the previous resource to endpoint.
// returns a *ItemRegisteredusersGraphendpointGraphEndpointRequestBuilder when successful
func (m *ItemRegisteredusersRegisteredUsersRequestBuilder) GraphEndpoint()(*ItemRegisteredusersGraphendpointGraphEndpointRequestBuilder) {
    return NewItemRegisteredusersGraphendpointGraphEndpointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *ItemRegisteredusersGraphserviceprincipalGraphServicePrincipalRequestBuilder when successful
func (m *ItemRegisteredusersRegisteredUsersRequestBuilder) GraphServicePrincipal()(*ItemRegisteredusersGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    return NewItemRegisteredusersGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *ItemRegisteredusersGraphuserGraphUserRequestBuilder when successful
func (m *ItemRegisteredusersRegisteredUsersRequestBuilder) GraphUser()(*ItemRegisteredusersGraphuserGraphUserRequestBuilder) {
    return NewItemRegisteredusersGraphuserGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ref provides operations to manage the collection of device entities.
// returns a *ItemRegisteredusersRefRequestBuilder when successful
func (m *ItemRegisteredusersRegisteredUsersRequestBuilder) Ref()(*ItemRegisteredusersRefRequestBuilder) {
    return NewItemRegisteredusersRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation retrieve a list of users that are registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration.
// returns a *RequestInformation when successful
func (m *ItemRegisteredusersRegisteredUsersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRegisteredusersRegisteredUsersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemRegisteredusersRegisteredUsersRequestBuilder when successful
func (m *ItemRegisteredusersRegisteredUsersRequestBuilder) WithUrl(rawUrl string)(*ItemRegisteredusersRegisteredUsersRequestBuilder) {
    return NewItemRegisteredusersRegisteredUsersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
