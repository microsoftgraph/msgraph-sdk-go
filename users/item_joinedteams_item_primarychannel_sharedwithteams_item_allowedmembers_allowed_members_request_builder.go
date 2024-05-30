package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder provides operations to manage the allowedMembers property of the microsoft.graph.sharedWithChannelTeamInfo entity.
type ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetQueryParameters a collection of team members who have access to the shared channel.
type ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetQueryParameters struct {
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
// ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetQueryParameters
}
// ByConversationMemberId provides operations to manage the allowedMembers property of the microsoft.graph.sharedWithChannelTeamInfo entity.
// returns a *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) ByConversationMemberId(conversationMemberId string)(*ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if conversationMemberId != "" {
        urlTplParams["conversationMember%2Did"] = conversationMemberId
    }
    return NewItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderInternal instantiates a new ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder and sets the default values.
func NewItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) {
    m := &ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/primaryChannel/sharedWithTeams/{sharedWithChannelTeamInfo%2Did}/allowedMembers{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder instantiates a new ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder and sets the default values.
func NewItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) Count()(*ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of team members who have access to the shared channel.
// returns a ConversationMemberCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConversationMemberCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConversationMemberCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConversationMemberCollectionResponseable), nil
}
// ToGetRequestInformation a collection of team members who have access to the shared channel.
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
