package chats

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPinnedmessagesItemMessageRequestBuilder provides operations to manage the message property of the microsoft.graph.pinnedChatMessageInfo entity.
type ItemPinnedmessagesItemMessageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPinnedmessagesItemMessageRequestBuilderGetQueryParameters represents details about the chat message that is pinned.
type ItemPinnedmessagesItemMessageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPinnedmessagesItemMessageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPinnedmessagesItemMessageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPinnedmessagesItemMessageRequestBuilderGetQueryParameters
}
// NewItemPinnedmessagesItemMessageRequestBuilderInternal instantiates a new ItemPinnedmessagesItemMessageRequestBuilder and sets the default values.
func NewItemPinnedmessagesItemMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPinnedmessagesItemMessageRequestBuilder) {
    m := &ItemPinnedmessagesItemMessageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/chats/{chat%2Did}/pinnedMessages/{pinnedChatMessageInfo%2Did}/message{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPinnedmessagesItemMessageRequestBuilder instantiates a new ItemPinnedmessagesItemMessageRequestBuilder and sets the default values.
func NewItemPinnedmessagesItemMessageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPinnedmessagesItemMessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPinnedmessagesItemMessageRequestBuilderInternal(urlParams, requestAdapter)
}
// Get represents details about the chat message that is pinned.
// returns a ChatMessageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPinnedmessagesItemMessageRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPinnedmessagesItemMessageRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageable), nil
}
// ToGetRequestInformation represents details about the chat message that is pinned.
// returns a *RequestInformation when successful
func (m *ItemPinnedmessagesItemMessageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPinnedmessagesItemMessageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPinnedmessagesItemMessageRequestBuilder when successful
func (m *ItemPinnedmessagesItemMessageRequestBuilder) WithUrl(rawUrl string)(*ItemPinnedmessagesItemMessageRequestBuilder) {
    return NewItemPinnedmessagesItemMessageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
