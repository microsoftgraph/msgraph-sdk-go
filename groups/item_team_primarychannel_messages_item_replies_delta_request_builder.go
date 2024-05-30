package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilder provides operations to call the delta method.
type ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilderGetQueryParameters retrieve the list of messages (without the replies) in a channel of a team. By using delta query, you can get new or updated messages in a channel. Delta query supports both full synchronization that retrieves all the messages in the specified channel, and incremental synchronization that retrieves those messages that have been added or changed in the channel since the last synchronization. Typically, you would do an initial full synchronization, and then get incremental changes to that messages view periodically. To get the replies for a message, use the list message replies or the get message reply operation. A GET request with the delta function returns either: State tokens are opaque to the client. To proceed with a round of change tracking, copy and apply the @odata.nextLink or @odata.deltaLink URL returned from the last GET request to the next delta function call for that same calendar view. A @odata.deltaLink returned in a response signifies that the current round of change tracking is complete. You can save and use the @odata.deltaLink URL when you begin to retrieve more changes (messages changed or posted after acquiring @odata.deltaLink). For more information, see the delta query documentation.
type ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilderGetQueryParameters struct {
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
// ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilderGetQueryParameters
}
// NewItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilderInternal instantiates a new ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilder and sets the default values.
func NewItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilder) {
    m := &ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/primaryChannel/messages/{chatMessage%2Did}/replies/delta(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilder instantiates a new ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilder and sets the default values.
func NewItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve the list of messages (without the replies) in a channel of a team. By using delta query, you can get new or updated messages in a channel. Delta query supports both full synchronization that retrieves all the messages in the specified channel, and incremental synchronization that retrieves those messages that have been added or changed in the channel since the last synchronization. Typically, you would do an initial full synchronization, and then get incremental changes to that messages view periodically. To get the replies for a message, use the list message replies or the get message reply operation. A GET request with the delta function returns either: State tokens are opaque to the client. To proceed with a round of change tracking, copy and apply the @odata.nextLink or @odata.deltaLink URL returned from the last GET request to the next delta function call for that same calendar view. A @odata.deltaLink returned in a response signifies that the current round of change tracking is complete. You can save and use the @odata.deltaLink URL when you begin to retrieve more changes (messages changed or posted after acquiring @odata.deltaLink). For more information, see the delta query documentation.
// Deprecated: This method is obsolete. Use GetAsDeltaGetResponse instead.
// returns a ItemTeamPrimarychannelMessagesItemRepliesDeltaResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chatmessage-delta?view=graph-rest-1.0
func (m *ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration)(ItemTeamPrimarychannelMessagesItemRepliesDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamPrimarychannelMessagesItemRepliesDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamPrimarychannelMessagesItemRepliesDeltaResponseable), nil
}
// GetAsDeltaGetResponse retrieve the list of messages (without the replies) in a channel of a team. By using delta query, you can get new or updated messages in a channel. Delta query supports both full synchronization that retrieves all the messages in the specified channel, and incremental synchronization that retrieves those messages that have been added or changed in the channel since the last synchronization. Typically, you would do an initial full synchronization, and then get incremental changes to that messages view periodically. To get the replies for a message, use the list message replies or the get message reply operation. A GET request with the delta function returns either: State tokens are opaque to the client. To proceed with a round of change tracking, copy and apply the @odata.nextLink or @odata.deltaLink URL returned from the last GET request to the next delta function call for that same calendar view. A @odata.deltaLink returned in a response signifies that the current round of change tracking is complete. You can save and use the @odata.deltaLink URL when you begin to retrieve more changes (messages changed or posted after acquiring @odata.deltaLink). For more information, see the delta query documentation.
// returns a ItemTeamPrimarychannelMessagesItemRepliesDeltaGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chatmessage-delta?view=graph-rest-1.0
func (m *ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilder) GetAsDeltaGetResponse(ctx context.Context, requestConfiguration *ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration)(ItemTeamPrimarychannelMessagesItemRepliesDeltaGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamPrimarychannelMessagesItemRepliesDeltaGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamPrimarychannelMessagesItemRepliesDeltaGetResponseable), nil
}
// ToGetRequestInformation retrieve the list of messages (without the replies) in a channel of a team. By using delta query, you can get new or updated messages in a channel. Delta query supports both full synchronization that retrieves all the messages in the specified channel, and incremental synchronization that retrieves those messages that have been added or changed in the channel since the last synchronization. Typically, you would do an initial full synchronization, and then get incremental changes to that messages view periodically. To get the replies for a message, use the list message replies or the get message reply operation. A GET request with the delta function returns either: State tokens are opaque to the client. To proceed with a round of change tracking, copy and apply the @odata.nextLink or @odata.deltaLink URL returned from the last GET request to the next delta function call for that same calendar view. A @odata.deltaLink returned in a response signifies that the current round of change tracking is complete. You can save and use the @odata.deltaLink URL when you begin to retrieve more changes (messages changed or posted after acquiring @odata.deltaLink). For more information, see the delta query documentation.
// returns a *RequestInformation when successful
func (m *ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilder when successful
func (m *ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilder) WithUrl(rawUrl string)(*ItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilder) {
    return NewItemTeamPrimarychannelMessagesItemRepliesDeltaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
