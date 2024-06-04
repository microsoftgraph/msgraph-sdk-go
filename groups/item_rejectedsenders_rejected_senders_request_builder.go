package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemRejectedsendersRejectedSendersRequestBuilder provides operations to manage the rejectedSenders property of the microsoft.graph.group entity.
type ItemRejectedsendersRejectedSendersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRejectedsendersRejectedSendersRequestBuilderGetQueryParameters users in the rejected senders list can't post to conversations of the group (identified in the GET request URL). Make sure you don't specify the same user or group in the rejected senders and accepted senders lists, otherwise you get an error.
type ItemRejectedsendersRejectedSendersRequestBuilderGetQueryParameters struct {
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
// ItemRejectedsendersRejectedSendersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRejectedsendersRejectedSendersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRejectedsendersRejectedSendersRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.rejectedSenders.item collection
// returns a *ItemRejectedsendersDirectoryObjectItemRequestBuilder when successful
func (m *ItemRejectedsendersRejectedSendersRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*ItemRejectedsendersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewItemRejectedsendersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemRejectedsendersRejectedSendersRequestBuilderInternal instantiates a new ItemRejectedsendersRejectedSendersRequestBuilder and sets the default values.
func NewItemRejectedsendersRejectedSendersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRejectedsendersRejectedSendersRequestBuilder) {
    m := &ItemRejectedsendersRejectedSendersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/rejectedSenders{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemRejectedsendersRejectedSendersRequestBuilder instantiates a new ItemRejectedsendersRejectedSendersRequestBuilder and sets the default values.
func NewItemRejectedsendersRejectedSendersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRejectedsendersRejectedSendersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRejectedsendersRejectedSendersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemRejectedsendersCountRequestBuilder when successful
func (m *ItemRejectedsendersRejectedSendersRequestBuilder) Count()(*ItemRejectedsendersCountRequestBuilder) {
    return NewItemRejectedsendersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get users in the rejected senders list can't post to conversations of the group (identified in the GET request URL). Make sure you don't specify the same user or group in the rejected senders and accepted senders lists, otherwise you get an error.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/group-list-rejectedsenders?view=graph-rest-1.0
func (m *ItemRejectedsendersRejectedSendersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRejectedsendersRejectedSendersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable), nil
}
// Ref provides operations to manage the collection of group entities.
// returns a *ItemRejectedsendersRefRequestBuilder when successful
func (m *ItemRejectedsendersRejectedSendersRequestBuilder) Ref()(*ItemRejectedsendersRefRequestBuilder) {
    return NewItemRejectedsendersRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation users in the rejected senders list can't post to conversations of the group (identified in the GET request URL). Make sure you don't specify the same user or group in the rejected senders and accepted senders lists, otherwise you get an error.
// returns a *RequestInformation when successful
func (m *ItemRejectedsendersRejectedSendersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRejectedsendersRejectedSendersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemRejectedsendersRejectedSendersRequestBuilder when successful
func (m *ItemRejectedsendersRejectedSendersRequestBuilder) WithUrl(rawUrl string)(*ItemRejectedsendersRejectedSendersRequestBuilder) {
    return NewItemRejectedsendersRejectedSendersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
