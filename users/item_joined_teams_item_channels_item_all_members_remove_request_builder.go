// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilder provides operations to call the remove method.
type ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilderInternal instantiates a new ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilder) {
    m := &ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/channels/{channel%2Did}/allMembers/remove", pathParameters),
    }
    return m
}
// NewItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilder instantiates a new ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove multiple members from a team in a single request. The response provides details about which memberships could and couldn't be removed.
// Deprecated: This method is obsolete. Use PostAsRemovePostResponse instead.
// returns a ItemJoinedTeamsItemChannelsItemAllMembersRemoveResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conversationmember-remove?view=graph-rest-1.0
func (m *ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilder) Post(ctx context.Context, body ItemJoinedTeamsItemChannelsItemAllMembersRemovePostRequestBodyable, requestConfiguration *ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilderPostRequestConfiguration)(ItemJoinedTeamsItemChannelsItemAllMembersRemoveResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemJoinedTeamsItemChannelsItemAllMembersRemoveResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemJoinedTeamsItemChannelsItemAllMembersRemoveResponseable), nil
}
// PostAsRemovePostResponse remove multiple members from a team in a single request. The response provides details about which memberships could and couldn't be removed.
// returns a ItemJoinedTeamsItemChannelsItemAllMembersRemovePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conversationmember-remove?view=graph-rest-1.0
func (m *ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilder) PostAsRemovePostResponse(ctx context.Context, body ItemJoinedTeamsItemChannelsItemAllMembersRemovePostRequestBodyable, requestConfiguration *ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilderPostRequestConfiguration)(ItemJoinedTeamsItemChannelsItemAllMembersRemovePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemJoinedTeamsItemChannelsItemAllMembersRemovePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemJoinedTeamsItemChannelsItemAllMembersRemovePostResponseable), nil
}
// ToPostRequestInformation remove multiple members from a team in a single request. The response provides details about which memberships could and couldn't be removed.
// returns a *RequestInformation when successful
func (m *ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemJoinedTeamsItemChannelsItemAllMembersRemovePostRequestBodyable, requestConfiguration *ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilder when successful
func (m *ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilder) {
    return NewItemJoinedTeamsItemChannelsItemAllMembersRemoveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
