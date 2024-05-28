package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilder provides operations to manage the team property of the microsoft.graph.teamInfo entity.
type ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilderGetQueryParameters get team from groups
type ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilderGetQueryParameters
}
// NewItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilderInternal instantiates a new ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilder and sets the default values.
func NewItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilder) {
    m := &ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/channels/{channel%2Did}/sharedWithTeams/{sharedWithChannelTeamInfo%2Did}/team{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilder instantiates a new ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilder and sets the default values.
func NewItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get team from groups
// returns a Teamable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable), nil
}
// ToGetRequestInformation get team from groups
// returns a *RequestInformation when successful
func (m *ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilder when successful
func (m *ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilder) WithUrl(rawUrl string)(*ItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilder) {
    return NewItemTeamChannelsItemSharedwithteamsItemTeamRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
