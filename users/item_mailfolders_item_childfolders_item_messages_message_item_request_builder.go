package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
type ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderGetQueryParameters the collection of messages in the mailFolder.
type ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderGetQueryParameters
}
// ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.message entity.
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) Attachments()(*ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderInternal instantiates a new ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder and sets the default values.
func NewItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) {
    m := &ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder instantiates a new ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder and sets the default values.
func NewItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemValueContentRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) Content()(*ItemMailfoldersItemChildfoldersItemMessagesItemValueContentRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesItemValueContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Copy provides operations to call the copy method.
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) Copy()(*ItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesItemCopyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateForward provides operations to call the createForward method.
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemCreateforwardCreateForwardRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) CreateForward()(*ItemMailfoldersItemChildfoldersItemMessagesItemCreateforwardCreateForwardRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesItemCreateforwardCreateForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateReply provides operations to call the createReply method.
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) CreateReply()(*ItemMailfoldersItemChildfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesItemCreatereplyCreateReplyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateReplyAll provides operations to call the createReplyAll method.
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemCreatereplyallCreateReplyAllRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) CreateReplyAll()(*ItemMailfoldersItemChildfoldersItemMessagesItemCreatereplyallCreateReplyAllRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesItemCreatereplyallCreateReplyAllRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property messages for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Extensions provides operations to manage the extensions property of the microsoft.graph.message entity.
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemExtensionsRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) Extensions()(*ItemMailfoldersItemChildfoldersItemMessagesItemExtensionsRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemForwardRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) Forward()(*ItemMailfoldersItemChildfoldersItemMessagesItemForwardRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of messages in the mailFolder.
// returns a Messageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemMoveRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) Move()(*ItemMailfoldersItemChildfoldersItemMessagesItemMoveRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesItemMoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property messages in users
// returns a Messageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, requestConfiguration *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemReplyRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) Reply()(*ItemMailfoldersItemChildfoldersItemMessagesItemReplyRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesItemReplyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReplyAll provides operations to call the replyAll method.
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemReplyallReplyAllRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) ReplyAll()(*ItemMailfoldersItemChildfoldersItemMessagesItemReplyallReplyAllRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesItemReplyallReplyAllRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Send provides operations to call the send method.
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemSendRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) Send()(*ItemMailfoldersItemChildfoldersItemMessagesItemSendRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesItemSendRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property messages for users
// returns a *RequestInformation when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the collection of messages in the mailFolder.
// returns a *RequestInformation when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property messages in users
// returns a *RequestInformation when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, requestConfiguration *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) WithUrl(rawUrl string)(*ItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesMessageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
