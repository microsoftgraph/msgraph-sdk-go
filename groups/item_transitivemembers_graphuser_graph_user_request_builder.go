package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTransitivemembersGraphuserGraphUserRequestBuilder casts the previous resource to user.
type ItemTransitivemembersGraphuserGraphUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTransitivemembersGraphuserGraphUserRequestBuilderGetQueryParameters get a list of the group's members. A group can have different object types as members. For more information about supported member types for different groups, see Group membership. This operation is transitive and returns a flat list of all nested members. An attempt to filter by an OData cast that represents an unsupported member type returns a 400 Bad Request error with the Request_UnsupportedQuery code.
type ItemTransitivemembersGraphuserGraphUserRequestBuilderGetQueryParameters struct {
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
// ItemTransitivemembersGraphuserGraphUserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTransitivemembersGraphuserGraphUserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTransitivemembersGraphuserGraphUserRequestBuilderGetQueryParameters
}
// NewItemTransitivemembersGraphuserGraphUserRequestBuilderInternal instantiates a new ItemTransitivemembersGraphuserGraphUserRequestBuilder and sets the default values.
func NewItemTransitivemembersGraphuserGraphUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitivemembersGraphuserGraphUserRequestBuilder) {
    m := &ItemTransitivemembersGraphuserGraphUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/transitiveMembers/graph.user{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTransitivemembersGraphuserGraphUserRequestBuilder instantiates a new ItemTransitivemembersGraphuserGraphUserRequestBuilder and sets the default values.
func NewItemTransitivemembersGraphuserGraphUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitivemembersGraphuserGraphUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTransitivemembersGraphuserGraphUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTransitivemembersGraphuserCountRequestBuilder when successful
func (m *ItemTransitivemembersGraphuserGraphUserRequestBuilder) Count()(*ItemTransitivemembersGraphuserCountRequestBuilder) {
    return NewItemTransitivemembersGraphuserCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the group's members. A group can have different object types as members. For more information about supported member types for different groups, see Group membership. This operation is transitive and returns a flat list of all nested members. An attempt to filter by an OData cast that represents an unsupported member type returns a 400 Bad Request error with the Request_UnsupportedQuery code.
// returns a UserCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/group-list-transitivemembers?view=graph-rest-1.0
func (m *ItemTransitivemembersGraphuserGraphUserRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTransitivemembersGraphuserGraphUserRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserCollectionResponseable), nil
}
// ToGetRequestInformation get a list of the group's members. A group can have different object types as members. For more information about supported member types for different groups, see Group membership. This operation is transitive and returns a flat list of all nested members. An attempt to filter by an OData cast that represents an unsupported member type returns a 400 Bad Request error with the Request_UnsupportedQuery code.
// returns a *RequestInformation when successful
func (m *ItemTransitivemembersGraphuserGraphUserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTransitivemembersGraphuserGraphUserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTransitivemembersGraphuserGraphUserRequestBuilder when successful
func (m *ItemTransitivemembersGraphuserGraphUserRequestBuilder) WithUrl(rawUrl string)(*ItemTransitivemembersGraphuserGraphUserRequestBuilder) {
    return NewItemTransitivemembersGraphuserGraphUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
