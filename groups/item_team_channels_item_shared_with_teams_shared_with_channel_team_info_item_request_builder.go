package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
type ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderGetQueryParameters get a team that has been shared with a specified channel. This operation is allowed only for channels with a **membershipType** value of `shared`.
type ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderGetQueryParameters
}
// ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllowedMembers provides operations to manage the allowedMembers property of the microsoft.graph.sharedWithChannelTeamInfo entity.
func (m *ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder) AllowedMembers()(*ItemTeamChannelsItemSharedWithTeamsItemAllowedMembersRequestBuilder) {
    return NewItemTeamChannelsItemSharedWithTeamsItemAllowedMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderInternal instantiates a new SharedWithChannelTeamInfoItemRequestBuilder and sets the default values.
func NewItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder) {
    m := &ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/channels/{channel%2Did}/sharedWithTeams/{sharedWithChannelTeamInfo%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder instantiates a new SharedWithChannelTeamInfoItemRequestBuilder and sets the default values.
func NewItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete unshare a channel with a team by deleting the corresponding sharedWithChannelTeamInfo resource. This operation is allowed only for channels with a **membershipType** value of `shared`.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/sharedwithchannelteaminfo-delete?view=graph-rest-1.0
func (m *ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get a team that has been shared with a specified channel. This operation is allowed only for channels with a **membershipType** value of `shared`.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/sharedwithchannelteaminfo-get?view=graph-rest-1.0
func (m *ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SharedWithChannelTeamInfoable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// Patch update the navigation property sharedWithTeams in groups
func (m *ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SharedWithChannelTeamInfoable, requestConfiguration *ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SharedWithChannelTeamInfoable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder) Team()(*ItemTeamChannelsItemSharedWithTeamsItemTeamRequestBuilder) {
    return NewItemTeamChannelsItemSharedWithTeamsItemTeamRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation unshare a channel with a team by deleting the corresponding sharedWithChannelTeamInfo resource. This operation is allowed only for channels with a **membershipType** value of `shared`.
func (m *ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get a team that has been shared with a specified channel. This operation is allowed only for channels with a **membershipType** value of `shared`.
func (m *ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property sharedWithTeams in groups
func (m *ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SharedWithChannelTeamInfoable, requestConfiguration *ItemTeamChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
