package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMessagesMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.user entity.
type ItemMessagesMessageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemMessagesMessageItemRequestBuilderGetQueryParameters the messages in a mailbox or folder. Read-only. Nullable.
type ItemMessagesMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Include Hidden Messages
    IncludeHiddenMessages *string `uriparametername:"includeHiddenMessages"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMessagesMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMessagesMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMessagesMessageItemRequestBuilderGetQueryParameters
}
// ItemMessagesMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMessagesMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.message entity.
func (m *ItemMessagesMessageItemRequestBuilder) Attachments()(*ItemMessagesItemAttachmentsRequestBuilder) {
    return NewItemMessagesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMessagesMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewItemMessagesMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMessagesMessageItemRequestBuilder) {
    m := &ItemMessagesMessageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/messages/{message%2Did}{?includeHiddenMessages*,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemMessagesMessageItemRequestBuilder instantiates a new MessageItemRequestBuilder and sets the default values.
func NewItemMessagesMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMessagesMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMessagesMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
func (m *ItemMessagesMessageItemRequestBuilder) Content()(*ItemMessagesItemValueContentRequestBuilder) {
    return NewItemMessagesItemValueContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Copy provides operations to call the copy method.
func (m *ItemMessagesMessageItemRequestBuilder) Copy()(*ItemMessagesItemCopyRequestBuilder) {
    return NewItemMessagesItemCopyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateForward provides operations to call the createForward method.
func (m *ItemMessagesMessageItemRequestBuilder) CreateForward()(*ItemMessagesItemCreateForwardRequestBuilder) {
    return NewItemMessagesItemCreateForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateReply provides operations to call the createReply method.
func (m *ItemMessagesMessageItemRequestBuilder) CreateReply()(*ItemMessagesItemCreateReplyRequestBuilder) {
    return NewItemMessagesItemCreateReplyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateReplyAll provides operations to call the createReplyAll method.
func (m *ItemMessagesMessageItemRequestBuilder) CreateReplyAll()(*ItemMessagesItemCreateReplyAllRequestBuilder) {
    return NewItemMessagesItemCreateReplyAllRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete a message in the specified user's mailbox, or delete a relationship of the message. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/message-delete?view=graph-rest-1.0
func (m *ItemMessagesMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.message entity.
func (m *ItemMessagesMessageItemRequestBuilder) Extensions()(*ItemMessagesItemExtensionsRequestBuilder) {
    return NewItemMessagesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
func (m *ItemMessagesMessageItemRequestBuilder) Forward()(*ItemMessagesItemForwardRequestBuilder) {
    return NewItemMessagesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the messages in a mailbox or folder. Read-only. Nullable.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/opentypeextension-get?view=graph-rest-1.0
func (m *ItemMessagesMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable), nil
}
// Move provides operations to call the move method.
func (m *ItemMessagesMessageItemRequestBuilder) Move()(*ItemMessagesItemMoveRequestBuilder) {
    return NewItemMessagesItemMoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a message object. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/message-update?view=graph-rest-1.0
func (m *ItemMessagesMessageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, requestConfiguration *ItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable), nil
}
// Reply provides operations to call the reply method.
func (m *ItemMessagesMessageItemRequestBuilder) Reply()(*ItemMessagesItemReplyRequestBuilder) {
    return NewItemMessagesItemReplyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReplyAll provides operations to call the replyAll method.
func (m *ItemMessagesMessageItemRequestBuilder) ReplyAll()(*ItemMessagesItemReplyAllRequestBuilder) {
    return NewItemMessagesItemReplyAllRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Send provides operations to call the send method.
func (m *ItemMessagesMessageItemRequestBuilder) Send()(*ItemMessagesItemSendRequestBuilder) {
    return NewItemMessagesItemSendRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a message in the specified user's mailbox, or delete a relationship of the message. This API is available in the following national cloud deployments.
func (m *ItemMessagesMessageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the messages in a mailbox or folder. Read-only. Nullable.
func (m *ItemMessagesMessageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a message object. This API is available in the following national cloud deployments.
func (m *ItemMessagesMessageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, requestConfiguration *ItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemMessagesMessageItemRequestBuilder) WithUrl(rawUrl string)(*ItemMessagesMessageItemRequestBuilder) {
    return NewItemMessagesMessageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
