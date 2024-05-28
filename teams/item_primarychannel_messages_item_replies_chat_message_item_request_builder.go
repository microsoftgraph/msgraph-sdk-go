package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder provides operations to manage the replies property of the microsoft.graph.chatMessage entity.
type ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderGetQueryParameters replies for a specified message. Supports $expand for channel messages.
type ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderGetQueryParameters
}
// ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderInternal instantiates a new ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder and sets the default values.
func NewItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) {
    m := &ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/primaryChannel/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder instantiates a new ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder and sets the default values.
func NewItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property replies for teams
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get replies for a specified message. Supports $expand for channel messages.
// returns a ChatMessageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageable, error) {
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
// HostedContents provides operations to manage the hostedContents property of the microsoft.graph.chatMessage entity.
// returns a *ItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder when successful
func (m *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) HostedContents()(*ItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder) {
    return NewItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property replies in teams
// returns a ChatMessageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageable, requestConfiguration *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// SetReaction provides operations to call the setReaction method.
// returns a *ItemPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder when successful
func (m *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) SetReaction()(*ItemPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder) {
    return NewItemPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SoftDelete provides operations to call the softDelete method.
// returns a *ItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder when successful
func (m *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) SoftDelete()(*ItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder) {
    return NewItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property replies for teams
// returns a *RequestInformation when successful
func (m *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation replies for a specified message. Supports $expand for channel messages.
// returns a *RequestInformation when successful
func (m *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property replies in teams
// returns a *RequestInformation when successful
func (m *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageable, requestConfiguration *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UndoSoftDelete provides operations to call the undoSoftDelete method.
// returns a *ItemPrimarychannelMessagesItemRepliesItemUndosoftdeleteUndoSoftDeleteRequestBuilder when successful
func (m *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) UndoSoftDelete()(*ItemPrimarychannelMessagesItemRepliesItemUndosoftdeleteUndoSoftDeleteRequestBuilder) {
    return NewItemPrimarychannelMessagesItemRepliesItemUndosoftdeleteUndoSoftDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UnsetReaction provides operations to call the unsetReaction method.
// returns a *ItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder when successful
func (m *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) UnsetReaction()(*ItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) {
    return NewItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder when successful
func (m *ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) WithUrl(rawUrl string)(*ItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) {
    return NewItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
