package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilder provides operations to call the getAllMessages method.
type ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilderGetQueryParameters retrieve messages across all channels in a team, including text, audio, and video conversations. To learn more about how to use the Microsoft Teams export APIs to export content, see Export content with the Microsoft Teams export APIs.
type ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // The payment model for the API
    Model *string `uriparametername:"model"`
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
// ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilderGetQueryParameters
}
// NewItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilderInternal instantiates a new ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilder and sets the default values.
func NewItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilder) {
    m := &ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/channels/getAllMessages(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top,model*}", pathParameters),
    }
    return m
}
// NewItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilder instantiates a new ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilder and sets the default values.
func NewItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve messages across all channels in a team, including text, audio, and video conversations. To learn more about how to use the Microsoft Teams export APIs to export content, see Export content with the Microsoft Teams export APIs.
// Deprecated: This method is obsolete. Use GetAsGetAllMessagesGetResponse instead.
// returns a ItemTeamChannelsGetallmessagesGetAllMessagesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/channel-getallmessages?view=graph-rest-1.0
func (m *ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilderGetRequestConfiguration)(ItemTeamChannelsGetallmessagesGetAllMessagesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamChannelsGetallmessagesGetAllMessagesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamChannelsGetallmessagesGetAllMessagesResponseable), nil
}
// GetAsGetAllMessagesGetResponse retrieve messages across all channels in a team, including text, audio, and video conversations. To learn more about how to use the Microsoft Teams export APIs to export content, see Export content with the Microsoft Teams export APIs.
// returns a ItemTeamChannelsGetallmessagesGetAllMessagesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/channel-getallmessages?view=graph-rest-1.0
func (m *ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilder) GetAsGetAllMessagesGetResponse(ctx context.Context, requestConfiguration *ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilderGetRequestConfiguration)(ItemTeamChannelsGetallmessagesGetAllMessagesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamChannelsGetallmessagesGetAllMessagesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamChannelsGetallmessagesGetAllMessagesGetResponseable), nil
}
// ToGetRequestInformation retrieve messages across all channels in a team, including text, audio, and video conversations. To learn more about how to use the Microsoft Teams export APIs to export content, see Export content with the Microsoft Teams export APIs.
// returns a *RequestInformation when successful
func (m *ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilder when successful
func (m *ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilder) WithUrl(rawUrl string)(*ItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilder) {
    return NewItemTeamChannelsGetallmessagesGetAllMessagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
