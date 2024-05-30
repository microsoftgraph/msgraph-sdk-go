package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder provides operations to manage the hostedContents property of the microsoft.graph.chatMessage entity.
type ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderGetQueryParameters content in a message hosted by Microsoft Teams - for example, images or code snippets.
type ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderGetQueryParameters
}
// ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderInternal instantiates a new ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder and sets the default values.
func NewItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) {
    m := &ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/primaryChannel/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/hostedContents/{chatMessageHostedContent%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder instantiates a new ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder and sets the default values.
func NewItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the team entity.
// returns a *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsItemValueContentRequestBuilder when successful
func (m *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) Content()(*ItemPrimarychannelMessagesItemRepliesItemHostedcontentsItemValueContentRequestBuilder) {
    return NewItemPrimarychannelMessagesItemRepliesItemHostedcontentsItemValueContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property hostedContents for teams
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get content in a message hosted by Microsoft Teams - for example, images or code snippets.
// returns a ChatMessageHostedContentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatMessageHostedContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable), nil
}
// Patch update the navigation property hostedContents in teams
// returns a ChatMessageHostedContentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, requestConfiguration *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatMessageHostedContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable), nil
}
// ToDeleteRequestInformation delete navigation property hostedContents for teams
// returns a *RequestInformation when successful
func (m *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation content in a message hosted by Microsoft Teams - for example, images or code snippets.
// returns a *RequestInformation when successful
func (m *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property hostedContents in teams
// returns a *RequestInformation when successful
func (m *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, requestConfiguration *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder when successful
func (m *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) WithUrl(rawUrl string)(*ItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) {
    return NewItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
