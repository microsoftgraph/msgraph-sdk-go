package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder provides operations to manage the hostedContents property of the microsoft.graph.chatMessage entity.
type ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilderGetQueryParameters content in a message hosted by Microsoft Teams - for example, images or code snippets.
type ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilderGetQueryParameters
}
// ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByChatMessageHostedContentId provides operations to manage the hostedContents property of the microsoft.graph.chatMessage entity.
// returns a *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder) ByChatMessageHostedContentId(chatMessageHostedContentId string)(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if chatMessageHostedContentId != "" {
        urlTplParams["chatMessageHostedContent%2Did"] = chatMessageHostedContentId
    }
    return NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsChatMessageHostedContentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilderInternal instantiates a new ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder and sets the default values.
func NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder) {
    m := &ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/primaryChannel/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/hostedContents{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder instantiates a new ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder and sets the default values.
func NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsCountRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder) Count()(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsCountRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get content in a message hosted by Microsoft Teams - for example, images or code snippets.
// returns a ChatMessageHostedContentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatMessageHostedContentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentCollectionResponseable), nil
}
// Post create new navigation property to hostedContents for users
// returns a ChatMessageHostedContentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, requestConfiguration *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation content in a message hosted by Microsoft Teams - for example, images or code snippets.
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to hostedContents for users
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, requestConfiguration *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemHostedcontentsHostedContentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
