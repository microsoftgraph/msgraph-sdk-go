package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTransitivememberofTransitiveMemberOfRequestBuilder provides operations to manage the transitiveMemberOf property of the microsoft.graph.user entity.
type ItemTransitivememberofTransitiveMemberOfRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTransitivememberofTransitiveMemberOfRequestBuilderGetQueryParameters get groups, directory roles, and administrative units that the user is a member of through either direct or transitive membership.
type ItemTransitivememberofTransitiveMemberOfRequestBuilderGetQueryParameters struct {
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
// ItemTransitivememberofTransitiveMemberOfRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTransitivememberofTransitiveMemberOfRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTransitivememberofTransitiveMemberOfRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId provides operations to manage the transitiveMemberOf property of the microsoft.graph.user entity.
// returns a *ItemTransitivememberofDirectoryObjectItemRequestBuilder when successful
func (m *ItemTransitivememberofTransitiveMemberOfRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*ItemTransitivememberofDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewItemTransitivememberofDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTransitivememberofTransitiveMemberOfRequestBuilderInternal instantiates a new ItemTransitivememberofTransitiveMemberOfRequestBuilder and sets the default values.
func NewItemTransitivememberofTransitiveMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitivememberofTransitiveMemberOfRequestBuilder) {
    m := &ItemTransitivememberofTransitiveMemberOfRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/transitiveMemberOf{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTransitivememberofTransitiveMemberOfRequestBuilder instantiates a new ItemTransitivememberofTransitiveMemberOfRequestBuilder and sets the default values.
func NewItemTransitivememberofTransitiveMemberOfRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitivememberofTransitiveMemberOfRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTransitivememberofTransitiveMemberOfRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTransitivememberofCountRequestBuilder when successful
func (m *ItemTransitivememberofTransitiveMemberOfRequestBuilder) Count()(*ItemTransitivememberofCountRequestBuilder) {
    return NewItemTransitivememberofCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get groups, directory roles, and administrative units that the user is a member of through either direct or transitive membership.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-list-transitivememberof?view=graph-rest-1.0
func (m *ItemTransitivememberofTransitiveMemberOfRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTransitivememberofTransitiveMemberOfRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
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
// GraphAdministrativeUnit casts the previous resource to administrativeUnit.
// returns a *ItemTransitivememberofGraphadministrativeunitGraphAdministrativeUnitRequestBuilder when successful
func (m *ItemTransitivememberofTransitiveMemberOfRequestBuilder) GraphAdministrativeUnit()(*ItemTransitivememberofGraphadministrativeunitGraphAdministrativeUnitRequestBuilder) {
    return NewItemTransitivememberofGraphadministrativeunitGraphAdministrativeUnitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphDirectoryRole casts the previous resource to directoryRole.
// returns a *ItemTransitivememberofGraphdirectoryroleGraphDirectoryRoleRequestBuilder when successful
func (m *ItemTransitivememberofTransitiveMemberOfRequestBuilder) GraphDirectoryRole()(*ItemTransitivememberofGraphdirectoryroleGraphDirectoryRoleRequestBuilder) {
    return NewItemTransitivememberofGraphdirectoryroleGraphDirectoryRoleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphGroup casts the previous resource to group.
// returns a *ItemTransitivememberofGraphgroupGraphGroupRequestBuilder when successful
func (m *ItemTransitivememberofTransitiveMemberOfRequestBuilder) GraphGroup()(*ItemTransitivememberofGraphgroupGraphGroupRequestBuilder) {
    return NewItemTransitivememberofGraphgroupGraphGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get groups, directory roles, and administrative units that the user is a member of through either direct or transitive membership.
// returns a *RequestInformation when successful
func (m *ItemTransitivememberofTransitiveMemberOfRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTransitivememberofTransitiveMemberOfRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTransitivememberofTransitiveMemberOfRequestBuilder when successful
func (m *ItemTransitivememberofTransitiveMemberOfRequestBuilder) WithUrl(rawUrl string)(*ItemTransitivememberofTransitiveMemberOfRequestBuilder) {
    return NewItemTransitivememberofTransitiveMemberOfRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
