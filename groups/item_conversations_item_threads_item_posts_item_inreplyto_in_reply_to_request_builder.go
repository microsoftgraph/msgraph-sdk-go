package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder provides operations to manage the inReplyTo property of the microsoft.graph.post entity.
type ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderGetQueryParameters read-only. Supports $expand.
type ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.post entity.
// returns a *ItemConversationsItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder when successful
func (m *ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) Attachments()(*ItemConversationsItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderInternal instantiates a new ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder and sets the default values.
func NewItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) {
    m := &ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/conversations/{conversation%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}/inReplyTo{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder instantiates a new ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder and sets the default values.
func NewItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderInternal(urlParams, requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.post entity.
// returns a *ItemConversationsItemThreadsItemPostsItemInreplytoExtensionsRequestBuilder when successful
func (m *ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) Extensions()(*ItemConversationsItemThreadsItemPostsItemInreplytoExtensionsRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInreplytoExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilder when successful
func (m *ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) Forward()(*ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only. Supports $expand.
// returns a Postable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Postable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePostFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Postable), nil
}
// Reply provides operations to call the reply method.
// returns a *ItemConversationsItemThreadsItemPostsItemInreplytoReplyRequestBuilder when successful
func (m *ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) Reply()(*ItemConversationsItemThreadsItemPostsItemInreplytoReplyRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInreplytoReplyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation read-only. Supports $expand.
// returns a *RequestInformation when successful
func (m *ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder when successful
func (m *ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) WithUrl(rawUrl string)(*ItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
