package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder provides operations to manage the allowedMembers property of the microsoft.graph.sharedWithChannelTeamInfo entity.
type DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetQueryParameters a collection of team members who have access to the shared channel.
type DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetQueryParameters
}
// NewDeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderInternal instantiates a new DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder and sets the default values.
func NewDeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) {
    m := &DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/sharedWithTeams/{sharedWithChannelTeamInfo%2Did}/allowedMembers/{conversationMember%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder instantiates a new DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder and sets the default values.
func NewDeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get a collection of team members who have access to the shared channel.
// returns a ConversationMemberable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConversationMemberable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConversationMemberFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConversationMemberable), nil
}
// ToGetRequestInformation a collection of team members who have access to the shared channel.
// returns a *RequestInformation when successful
func (m *DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder when successful
func (m *DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) WithUrl(rawUrl string)(*DeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) {
    return NewDeletedteamsItemChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
