package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder provides operations to manage the allChannels property of the microsoft.graph.team entity.
type ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilderGetQueryParameters list of channels either hosted in or shared with the team (incoming channels).
type ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilderGetQueryParameters struct {
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
// ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilderGetQueryParameters
}
// ByChannelId provides operations to manage the allChannels property of the microsoft.graph.team entity.
// returns a *ItemJoinedteamsItemAllchannelsChannelItemRequestBuilder when successful
func (m *ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder) ByChannelId(channelId string)(*ItemJoinedteamsItemAllchannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if channelId != "" {
        urlTplParams["channel%2Did"] = channelId
    }
    return NewItemJoinedteamsItemAllchannelsChannelItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemJoinedteamsItemAllchannelsAllChannelsRequestBuilderInternal instantiates a new ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder and sets the default values.
func NewItemJoinedteamsItemAllchannelsAllChannelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder) {
    m := &ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/allChannels{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder instantiates a new ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder and sets the default values.
func NewItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsItemAllchannelsAllChannelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemJoinedteamsItemAllchannelsCountRequestBuilder when successful
func (m *ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder) Count()(*ItemJoinedteamsItemAllchannelsCountRequestBuilder) {
    return NewItemJoinedteamsItemAllchannelsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of channels either hosted in or shared with the team (incoming channels).
// returns a ChannelCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChannelCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChannelCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChannelCollectionResponseable), nil
}
// ToGetRequestInformation list of channels either hosted in or shared with the team (incoming channels).
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder when successful
func (m *ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder) {
    return NewItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
