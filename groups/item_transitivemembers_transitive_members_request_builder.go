package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTransitivemembersTransitiveMembersRequestBuilder provides operations to manage the transitiveMembers property of the microsoft.graph.group entity.
type ItemTransitivemembersTransitiveMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTransitivemembersTransitiveMembersRequestBuilderGetQueryParameters get a list of the group's members. A group can have different object types as members. For more information about supported member types for different groups, see Group membership. This operation is transitive and returns a flat list of all nested members. An attempt to filter by an OData cast that represents an unsupported member type returns a 400 Bad Request error with the Request_UnsupportedQuery code.
type ItemTransitivemembersTransitiveMembersRequestBuilderGetQueryParameters struct {
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
// ItemTransitivemembersTransitiveMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTransitivemembersTransitiveMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTransitivemembersTransitiveMembersRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId provides operations to manage the transitiveMembers property of the microsoft.graph.group entity.
// returns a *ItemTransitivemembersDirectoryObjectItemRequestBuilder when successful
func (m *ItemTransitivemembersTransitiveMembersRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*ItemTransitivemembersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewItemTransitivemembersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTransitivemembersTransitiveMembersRequestBuilderInternal instantiates a new ItemTransitivemembersTransitiveMembersRequestBuilder and sets the default values.
func NewItemTransitivemembersTransitiveMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitivemembersTransitiveMembersRequestBuilder) {
    m := &ItemTransitivemembersTransitiveMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/transitiveMembers{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTransitivemembersTransitiveMembersRequestBuilder instantiates a new ItemTransitivemembersTransitiveMembersRequestBuilder and sets the default values.
func NewItemTransitivemembersTransitiveMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitivemembersTransitiveMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTransitivemembersTransitiveMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTransitivemembersCountRequestBuilder when successful
func (m *ItemTransitivemembersTransitiveMembersRequestBuilder) Count()(*ItemTransitivemembersCountRequestBuilder) {
    return NewItemTransitivemembersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the group's members. A group can have different object types as members. For more information about supported member types for different groups, see Group membership. This operation is transitive and returns a flat list of all nested members. An attempt to filter by an OData cast that represents an unsupported member type returns a 400 Bad Request error with the Request_UnsupportedQuery code.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/group-list-transitivemembers?view=graph-rest-1.0
func (m *ItemTransitivemembersTransitiveMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTransitivemembersTransitiveMembersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
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
// returns a *ItemTransitivemembersGraphapplicationGraphApplicationRequestBuilder when successful
func (m *ItemTransitivemembersTransitiveMembersRequestBuilder) GraphApplication()(*ItemTransitivemembersGraphapplicationGraphApplicationRequestBuilder) {
    return NewItemTransitivemembersGraphapplicationGraphApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphDevice casts the previous resource to device.
// returns a *ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder when successful
func (m *ItemTransitivemembersTransitiveMembersRequestBuilder) GraphDevice()(*ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder) {
    return NewItemTransitivemembersGraphdeviceGraphDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphGroup casts the previous resource to group.
// returns a *ItemTransitivemembersGraphgroupGraphGroupRequestBuilder when successful
func (m *ItemTransitivemembersTransitiveMembersRequestBuilder) GraphGroup()(*ItemTransitivemembersGraphgroupGraphGroupRequestBuilder) {
    return NewItemTransitivemembersGraphgroupGraphGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphOrgContact casts the previous resource to orgContact.
// returns a *ItemTransitivemembersGraphorgcontactGraphOrgContactRequestBuilder when successful
func (m *ItemTransitivemembersTransitiveMembersRequestBuilder) GraphOrgContact()(*ItemTransitivemembersGraphorgcontactGraphOrgContactRequestBuilder) {
    return NewItemTransitivemembersGraphorgcontactGraphOrgContactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *ItemTransitivemembersGraphserviceprincipalGraphServicePrincipalRequestBuilder when successful
func (m *ItemTransitivemembersTransitiveMembersRequestBuilder) GraphServicePrincipal()(*ItemTransitivemembersGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    return NewItemTransitivemembersGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *ItemTransitivemembersGraphuserGraphUserRequestBuilder when successful
func (m *ItemTransitivemembersTransitiveMembersRequestBuilder) GraphUser()(*ItemTransitivemembersGraphuserGraphUserRequestBuilder) {
    return NewItemTransitivemembersGraphuserGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get a list of the group's members. A group can have different object types as members. For more information about supported member types for different groups, see Group membership. This operation is transitive and returns a flat list of all nested members. An attempt to filter by an OData cast that represents an unsupported member type returns a 400 Bad Request error with the Request_UnsupportedQuery code.
// returns a *RequestInformation when successful
func (m *ItemTransitivemembersTransitiveMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTransitivemembersTransitiveMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTransitivemembersTransitiveMembersRequestBuilder when successful
func (m *ItemTransitivemembersTransitiveMembersRequestBuilder) WithUrl(rawUrl string)(*ItemTransitivemembersTransitiveMembersRequestBuilder) {
    return NewItemTransitivemembersTransitiveMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
