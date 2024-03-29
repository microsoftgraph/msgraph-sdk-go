package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilder provides operations to manage the media for the team entity.
type ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilderGetQueryParameters get media content for the navigation property hostedContents from teams
type ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilderGetQueryParameters struct {
    // Format of the content
    Format *string `uriparametername:"%24format"`
}
// ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilderGetQueryParameters
}
// ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilderInternal instantiates a new ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilder and sets the default values.
func NewItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilder) {
    m := &ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/primaryChannel/messages/{chatMessage%2Did}/hostedContents/{chatMessageHostedContent%2Did}/$value{?%24format*}", pathParameters),
    }
    return m
}
// NewItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilder instantiates a new ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilder and sets the default values.
func NewItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get media content for the navigation property hostedContents from teams
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chatmessage-list-hostedcontents?view=graph-rest-1.0
func (m *ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put update media content for the navigation property hostedContents in teams
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilderPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation get media content for the navigation property hostedContents from teams
// returns a *RequestInformation when successful
func (m *ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation update media content for the navigation property hostedContents in teams
// returns a *RequestInformation when successful
func (m *ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, "{+baseurl}/teams/{team%2Did}/primaryChannel/messages/{chatMessage%2Did}/hostedContents/{chatMessageHostedContent%2Did}/$value", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetStreamContentAndContentType(body, "application/octet-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilder when successful
func (m *ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilder) WithUrl(rawUrl string)(*ItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilder) {
    return NewItemPrimaryChannelMessagesItemHostedContentsItemValueContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
