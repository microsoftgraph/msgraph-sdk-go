package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilder provides operations to count the resources in the collection.
type ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilderGetQueryParameters get the number of the resource
type ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilderGetQueryParameters
}
// NewItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilderInternal instantiates a new ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilder and sets the default values.
func NewItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilder) {
    m := &ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/channels/{channel%2Did}/messages/{chatMessage%2Did}/hostedContents/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilder instantiates a new ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilder and sets the default values.
func NewItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilder when successful
func (m *ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilder) {
    return NewItemJoinedteamsItemChannelsItemMessagesItemHostedcontentsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
