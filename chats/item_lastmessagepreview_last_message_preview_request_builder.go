package chats

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemLastmessagepreviewLastMessagePreviewRequestBuilder provides operations to manage the lastMessagePreview property of the microsoft.graph.chat entity.
type ItemLastmessagepreviewLastMessagePreviewRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemLastmessagepreviewLastMessagePreviewRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemLastmessagepreviewLastMessagePreviewRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemLastmessagepreviewLastMessagePreviewRequestBuilderGetQueryParameters preview of the last message sent in the chat. Null if no messages were sent in the chat. Currently, only the list chats operation supports this property.
type ItemLastmessagepreviewLastMessagePreviewRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemLastmessagepreviewLastMessagePreviewRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemLastmessagepreviewLastMessagePreviewRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemLastmessagepreviewLastMessagePreviewRequestBuilderGetQueryParameters
}
// ItemLastmessagepreviewLastMessagePreviewRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemLastmessagepreviewLastMessagePreviewRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemLastmessagepreviewLastMessagePreviewRequestBuilderInternal instantiates a new ItemLastmessagepreviewLastMessagePreviewRequestBuilder and sets the default values.
func NewItemLastmessagepreviewLastMessagePreviewRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemLastmessagepreviewLastMessagePreviewRequestBuilder) {
    m := &ItemLastmessagepreviewLastMessagePreviewRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/chats/{chat%2Did}/lastMessagePreview{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemLastmessagepreviewLastMessagePreviewRequestBuilder instantiates a new ItemLastmessagepreviewLastMessagePreviewRequestBuilder and sets the default values.
func NewItemLastmessagepreviewLastMessagePreviewRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemLastmessagepreviewLastMessagePreviewRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemLastmessagepreviewLastMessagePreviewRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property lastMessagePreview for chats
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemLastmessagepreviewLastMessagePreviewRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemLastmessagepreviewLastMessagePreviewRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get preview of the last message sent in the chat. Null if no messages were sent in the chat. Currently, only the list chats operation supports this property.
// returns a ChatMessageInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemLastmessagepreviewLastMessagePreviewRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemLastmessagepreviewLastMessagePreviewRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageInfoable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatMessageInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageInfoable), nil
}
// Patch update the navigation property lastMessagePreview in chats
// returns a ChatMessageInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemLastmessagepreviewLastMessagePreviewRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageInfoable, requestConfiguration *ItemLastmessagepreviewLastMessagePreviewRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageInfoable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatMessageInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageInfoable), nil
}
// ToDeleteRequestInformation delete navigation property lastMessagePreview for chats
// returns a *RequestInformation when successful
func (m *ItemLastmessagepreviewLastMessagePreviewRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemLastmessagepreviewLastMessagePreviewRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation preview of the last message sent in the chat. Null if no messages were sent in the chat. Currently, only the list chats operation supports this property.
// returns a *RequestInformation when successful
func (m *ItemLastmessagepreviewLastMessagePreviewRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemLastmessagepreviewLastMessagePreviewRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property lastMessagePreview in chats
// returns a *RequestInformation when successful
func (m *ItemLastmessagepreviewLastMessagePreviewRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageInfoable, requestConfiguration *ItemLastmessagepreviewLastMessagePreviewRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemLastmessagepreviewLastMessagePreviewRequestBuilder when successful
func (m *ItemLastmessagepreviewLastMessagePreviewRequestBuilder) WithUrl(rawUrl string)(*ItemLastmessagepreviewLastMessagePreviewRequestBuilder) {
    return NewItemLastmessagepreviewLastMessagePreviewRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
