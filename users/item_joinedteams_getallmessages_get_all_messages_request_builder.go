package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilder provides operations to call the getAllMessages method.
type ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilderGetQueryParameters invoke function getAllMessages
type ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // The payment model for the API
    Model *string `uriparametername:"model"`
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
// ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilderGetQueryParameters
}
// NewItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilderInternal instantiates a new ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilder and sets the default values.
func NewItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilder) {
    m := &ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/getAllMessages(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top,model*}", pathParameters),
    }
    return m
}
// NewItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilder instantiates a new ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilder and sets the default values.
func NewItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getAllMessages
// Deprecated: This method is obsolete. Use GetAsGetAllMessagesGetResponse instead.
// returns a ItemJoinedteamsGetallmessagesGetAllMessagesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilderGetRequestConfiguration)(ItemJoinedteamsGetallmessagesGetAllMessagesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemJoinedteamsGetallmessagesGetAllMessagesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemJoinedteamsGetallmessagesGetAllMessagesResponseable), nil
}
// GetAsGetAllMessagesGetResponse invoke function getAllMessages
// returns a ItemJoinedteamsGetallmessagesGetAllMessagesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilder) GetAsGetAllMessagesGetResponse(ctx context.Context, requestConfiguration *ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilderGetRequestConfiguration)(ItemJoinedteamsGetallmessagesGetAllMessagesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemJoinedteamsGetallmessagesGetAllMessagesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemJoinedteamsGetallmessagesGetAllMessagesGetResponseable), nil
}
// ToGetRequestInformation invoke function getAllMessages
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilder when successful
func (m *ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilder) {
    return NewItemJoinedteamsGetallmessagesGetAllMessagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
