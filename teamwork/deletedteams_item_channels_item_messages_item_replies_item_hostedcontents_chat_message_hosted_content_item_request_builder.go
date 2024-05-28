package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder provides operations to manage the hostedContents property of the microsoft.graph.chatMessage entity.
type DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderGetQueryParameters content in a message hosted by Microsoft Teams - for example, images or code snippets.
type DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderGetQueryParameters
}
// DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderInternal instantiates a new DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder and sets the default values.
func NewDeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) {
    m := &DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/hostedContents/{chatMessageHostedContent%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder instantiates a new DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder and sets the default values.
func NewDeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the teamwork entity.
// returns a *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsItemValueContentRequestBuilder when successful
func (m *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) Content()(*DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsItemValueContentRequestBuilder) {
    return NewDeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsItemValueContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property hostedContents for teamwork
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, error) {
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
// Patch update the navigation property hostedContents in teamwork
// returns a ChatMessageHostedContentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, requestConfiguration *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, error) {
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
// ToDeleteRequestInformation delete navigation property hostedContents for teamwork
// returns a *RequestInformation when successful
func (m *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property hostedContents in teamwork
// returns a *RequestInformation when successful
func (m *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, requestConfiguration *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder when successful
func (m *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) WithUrl(rawUrl string)(*DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) {
    return NewDeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
