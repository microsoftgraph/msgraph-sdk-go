package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder provides operations to manage the replies property of the microsoft.graph.chatMessage entity.
type ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderGetQueryParameters replies for a specified message. Supports $expand for channel messages.
type ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderGetQueryParameters
}
// ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderInternal instantiates a new ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder and sets the default values.
func NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) {
    m := &ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/primaryChannel/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder instantiates a new ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder and sets the default values.
func NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property replies for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageable, error) {
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
// returns a *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) HostedContents()(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property replies in users
// returns a ChatMessageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageable, requestConfiguration *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageable, error) {
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
// returns a *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) SetReaction()(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SoftDelete provides operations to call the softDelete method.
// returns a *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) SoftDelete()(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property replies for users
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property replies in users
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageable, requestConfiguration *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUndosoftdeleteUndoSoftDeleteRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) UndoSoftDelete()(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUndosoftdeleteUndoSoftDeleteRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUndosoftdeleteUndoSoftDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UnsetReaction provides operations to call the unsetReaction method.
// returns a *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) UnsetReaction()(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesChatMessageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
