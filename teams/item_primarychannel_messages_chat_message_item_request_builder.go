package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPrimarychannelMessagesChatMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.channel entity.
type ItemPrimarychannelMessagesChatMessageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPrimarychannelMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimarychannelMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPrimarychannelMessagesChatMessageItemRequestBuilderGetQueryParameters a collection of all the messages in the channel. A navigation property. Nullable.
type ItemPrimarychannelMessagesChatMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPrimarychannelMessagesChatMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimarychannelMessagesChatMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPrimarychannelMessagesChatMessageItemRequestBuilderGetQueryParameters
}
// ItemPrimarychannelMessagesChatMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimarychannelMessagesChatMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPrimarychannelMessagesChatMessageItemRequestBuilderInternal instantiates a new ItemPrimarychannelMessagesChatMessageItemRequestBuilder and sets the default values.
func NewItemPrimarychannelMessagesChatMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrimarychannelMessagesChatMessageItemRequestBuilder) {
    m := &ItemPrimarychannelMessagesChatMessageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/primaryChannel/messages/{chatMessage%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPrimarychannelMessagesChatMessageItemRequestBuilder instantiates a new ItemPrimarychannelMessagesChatMessageItemRequestBuilder and sets the default values.
func NewItemPrimarychannelMessagesChatMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrimarychannelMessagesChatMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPrimarychannelMessagesChatMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property messages for teams
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPrimarychannelMessagesChatMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPrimarychannelMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of all the messages in the channel. A navigation property. Nullable.
// returns a ChatMessageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPrimarychannelMessagesChatMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPrimarychannelMessagesChatMessageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageable, error) {
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
// returns a *ItemPrimarychannelMessagesItemHostedcontentsHostedContentsRequestBuilder when successful
func (m *ItemPrimarychannelMessagesChatMessageItemRequestBuilder) HostedContents()(*ItemPrimarychannelMessagesItemHostedcontentsHostedContentsRequestBuilder) {
    return NewItemPrimarychannelMessagesItemHostedcontentsHostedContentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property messages in teams
// returns a ChatMessageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPrimarychannelMessagesChatMessageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageable, requestConfiguration *ItemPrimarychannelMessagesChatMessageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageable, error) {
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
// Replies provides operations to manage the replies property of the microsoft.graph.chatMessage entity.
// returns a *ItemPrimarychannelMessagesItemRepliesRequestBuilder when successful
func (m *ItemPrimarychannelMessagesChatMessageItemRequestBuilder) Replies()(*ItemPrimarychannelMessagesItemRepliesRequestBuilder) {
    return NewItemPrimarychannelMessagesItemRepliesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetReaction provides operations to call the setReaction method.
// returns a *ItemPrimarychannelMessagesItemSetreactionSetReactionRequestBuilder when successful
func (m *ItemPrimarychannelMessagesChatMessageItemRequestBuilder) SetReaction()(*ItemPrimarychannelMessagesItemSetreactionSetReactionRequestBuilder) {
    return NewItemPrimarychannelMessagesItemSetreactionSetReactionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SoftDelete provides operations to call the softDelete method.
// returns a *ItemPrimarychannelMessagesItemSoftdeleteSoftDeleteRequestBuilder when successful
func (m *ItemPrimarychannelMessagesChatMessageItemRequestBuilder) SoftDelete()(*ItemPrimarychannelMessagesItemSoftdeleteSoftDeleteRequestBuilder) {
    return NewItemPrimarychannelMessagesItemSoftdeleteSoftDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property messages for teams
// returns a *RequestInformation when successful
func (m *ItemPrimarychannelMessagesChatMessageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPrimarychannelMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of all the messages in the channel. A navigation property. Nullable.
// returns a *RequestInformation when successful
func (m *ItemPrimarychannelMessagesChatMessageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPrimarychannelMessagesChatMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property messages in teams
// returns a *RequestInformation when successful
func (m *ItemPrimarychannelMessagesChatMessageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageable, requestConfiguration *ItemPrimarychannelMessagesChatMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder when successful
func (m *ItemPrimarychannelMessagesChatMessageItemRequestBuilder) UndoSoftDelete()(*ItemPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder) {
    return NewItemPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UnsetReaction provides operations to call the unsetReaction method.
// returns a *ItemPrimarychannelMessagesItemUnsetreactionUnsetReactionRequestBuilder when successful
func (m *ItemPrimarychannelMessagesChatMessageItemRequestBuilder) UnsetReaction()(*ItemPrimarychannelMessagesItemUnsetreactionUnsetReactionRequestBuilder) {
    return NewItemPrimarychannelMessagesItemUnsetreactionUnsetReactionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPrimarychannelMessagesChatMessageItemRequestBuilder when successful
func (m *ItemPrimarychannelMessagesChatMessageItemRequestBuilder) WithUrl(rawUrl string)(*ItemPrimarychannelMessagesChatMessageItemRequestBuilder) {
    return NewItemPrimarychannelMessagesChatMessageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
