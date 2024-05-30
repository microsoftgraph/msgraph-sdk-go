package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder provides operations to manage the pinnedMessages property of the microsoft.graph.chat entity.
type ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilderGetQueryParameters a collection of all the pinned messages in the chat. Nullable.
type ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilderGetQueryParameters struct {
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
// ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilderGetQueryParameters
}
// ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPinnedChatMessageInfoId provides operations to manage the pinnedMessages property of the microsoft.graph.chat entity.
// returns a *ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder when successful
func (m *ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder) ByPinnedChatMessageInfoId(pinnedChatMessageInfoId string)(*ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if pinnedChatMessageInfoId != "" {
        urlTplParams["pinnedChatMessageInfo%2Did"] = pinnedChatMessageInfoId
    }
    return NewItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemChatsItemPinnedmessagesPinnedMessagesRequestBuilderInternal instantiates a new ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder and sets the default values.
func NewItemChatsItemPinnedmessagesPinnedMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder) {
    m := &ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}/pinnedMessages{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder instantiates a new ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder and sets the default values.
func NewItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChatsItemPinnedmessagesPinnedMessagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemChatsItemPinnedmessagesCountRequestBuilder when successful
func (m *ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder) Count()(*ItemChatsItemPinnedmessagesCountRequestBuilder) {
    return NewItemChatsItemPinnedmessagesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of all the pinned messages in the chat. Nullable.
// returns a PinnedChatMessageInfoCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PinnedChatMessageInfoCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePinnedChatMessageInfoCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PinnedChatMessageInfoCollectionResponseable), nil
}
// Post create new navigation property to pinnedMessages for users
// returns a PinnedChatMessageInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PinnedChatMessageInfoable, requestConfiguration *ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PinnedChatMessageInfoable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePinnedChatMessageInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PinnedChatMessageInfoable), nil
}
// ToGetRequestInformation a collection of all the pinned messages in the chat. Nullable.
// returns a *RequestInformation when successful
func (m *ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to pinnedMessages for users
// returns a *RequestInformation when successful
func (m *ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PinnedChatMessageInfoable, requestConfiguration *ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder when successful
func (m *ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder) WithUrl(rawUrl string)(*ItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder) {
    return NewItemChatsItemPinnedmessagesPinnedMessagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
