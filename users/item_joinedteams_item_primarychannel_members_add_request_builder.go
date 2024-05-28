package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilder provides operations to call the add method.
type ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemJoinedteamsItemPrimarychannelMembersAddRequestBuilderInternal instantiates a new ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilder and sets the default values.
func NewItemJoinedteamsItemPrimarychannelMembersAddRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilder) {
    m := &ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/primaryChannel/members/add", pathParameters),
    }
    return m
}
// NewItemJoinedteamsItemPrimarychannelMembersAddRequestBuilder instantiates a new ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilder and sets the default values.
func NewItemJoinedteamsItemPrimarychannelMembersAddRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsItemPrimarychannelMembersAddRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
// Deprecated: This method is obsolete. Use PostAsAddPostResponse instead.
// returns a ItemJoinedteamsItemPrimarychannelMembersAddResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conversationmembers-add?view=graph-rest-1.0
func (m *ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilder) Post(ctx context.Context, body ItemJoinedteamsItemPrimarychannelMembersAddPostRequestBodyable, requestConfiguration *ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilderPostRequestConfiguration)(ItemJoinedteamsItemPrimarychannelMembersAddResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemJoinedteamsItemPrimarychannelMembersAddResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemJoinedteamsItemPrimarychannelMembersAddResponseable), nil
}
// PostAsAddPostResponse add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
// returns a ItemJoinedteamsItemPrimarychannelMembersAddPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conversationmembers-add?view=graph-rest-1.0
func (m *ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilder) PostAsAddPostResponse(ctx context.Context, body ItemJoinedteamsItemPrimarychannelMembersAddPostRequestBodyable, requestConfiguration *ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilderPostRequestConfiguration)(ItemJoinedteamsItemPrimarychannelMembersAddPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemJoinedteamsItemPrimarychannelMembersAddPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemJoinedteamsItemPrimarychannelMembersAddPostResponseable), nil
}
// ToPostRequestInformation add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemJoinedteamsItemPrimarychannelMembersAddPostRequestBodyable, requestConfiguration *ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsItemPrimarychannelMembersAddRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelMembersAddRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
