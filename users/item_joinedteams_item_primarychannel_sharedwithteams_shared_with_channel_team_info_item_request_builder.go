package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
type ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderGetQueryParameters a collection of teams with which a channel is shared.
type ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderGetQueryParameters
}
// ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllowedMembers provides operations to manage the allowedMembers property of the microsoft.graph.sharedWithChannelTeamInfo entity.
// returns a *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder) AllowedMembers()(*ItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderInternal instantiates a new ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder and sets the default values.
func NewItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder) {
    m := &ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/primaryChannel/sharedWithTeams/{sharedWithChannelTeamInfo%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder instantiates a new ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder and sets the default values.
func NewItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property sharedWithTeams for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a collection of teams with which a channel is shared.
// returns a SharedWithChannelTeamInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SharedWithChannelTeamInfoable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSharedWithChannelTeamInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SharedWithChannelTeamInfoable), nil
}
// Patch update the navigation property sharedWithTeams in users
// returns a SharedWithChannelTeamInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SharedWithChannelTeamInfoable, requestConfiguration *ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SharedWithChannelTeamInfoable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSharedWithChannelTeamInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SharedWithChannelTeamInfoable), nil
}
// Team provides operations to manage the team property of the microsoft.graph.teamInfo entity.
// returns a *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder) Team()(*ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property sharedWithTeams for users
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of teams with which a channel is shared.
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property sharedWithTeams in users
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SharedWithChannelTeamInfoable, requestConfiguration *ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelSharedwithteamsSharedWithChannelTeamInfoItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
