package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilder provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
type ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilderGetQueryParameters list of channels shared with the team.
type ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilderGetQueryParameters
}
// NewItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilderInternal instantiates a new ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilder and sets the default values.
func NewItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilder) {
    m := &ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/incomingChannels/{channel%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilder instantiates a new ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilder and sets the default values.
func NewItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list of channels shared with the team.
// returns a Channelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChannelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable), nil
}
// ToGetRequestInformation list of channels shared with the team.
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilder when successful
func (m *ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilder) {
    return NewItemJoinedteamsItemIncomingchannelsChannelItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
