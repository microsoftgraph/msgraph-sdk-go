package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeletedteamsGetallmessagesGetAllMessagesRequestBuilder provides operations to call the getAllMessages method.
type DeletedteamsGetallmessagesGetAllMessagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedteamsGetallmessagesGetAllMessagesRequestBuilderGetQueryParameters invoke function getAllMessages
type DeletedteamsGetallmessagesGetAllMessagesRequestBuilderGetQueryParameters struct {
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
// DeletedteamsGetallmessagesGetAllMessagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedteamsGetallmessagesGetAllMessagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeletedteamsGetallmessagesGetAllMessagesRequestBuilderGetQueryParameters
}
// NewDeletedteamsGetallmessagesGetAllMessagesRequestBuilderInternal instantiates a new DeletedteamsGetallmessagesGetAllMessagesRequestBuilder and sets the default values.
func NewDeletedteamsGetallmessagesGetAllMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsGetallmessagesGetAllMessagesRequestBuilder) {
    m := &DeletedteamsGetallmessagesGetAllMessagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/deletedTeams/getAllMessages(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top,model*}", pathParameters),
    }
    return m
}
// NewDeletedteamsGetallmessagesGetAllMessagesRequestBuilder instantiates a new DeletedteamsGetallmessagesGetAllMessagesRequestBuilder and sets the default values.
func NewDeletedteamsGetallmessagesGetAllMessagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsGetallmessagesGetAllMessagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedteamsGetallmessagesGetAllMessagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getAllMessages
// Deprecated: This method is obsolete. Use GetAsGetAllMessagesGetResponse instead.
// returns a DeletedteamsGetallmessagesGetAllMessagesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedteamsGetallmessagesGetAllMessagesRequestBuilder) Get(ctx context.Context, requestConfiguration *DeletedteamsGetallmessagesGetAllMessagesRequestBuilderGetRequestConfiguration)(DeletedteamsGetallmessagesGetAllMessagesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeletedteamsGetallmessagesGetAllMessagesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeletedteamsGetallmessagesGetAllMessagesResponseable), nil
}
// GetAsGetAllMessagesGetResponse invoke function getAllMessages
// returns a DeletedteamsGetallmessagesGetAllMessagesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedteamsGetallmessagesGetAllMessagesRequestBuilder) GetAsGetAllMessagesGetResponse(ctx context.Context, requestConfiguration *DeletedteamsGetallmessagesGetAllMessagesRequestBuilderGetRequestConfiguration)(DeletedteamsGetallmessagesGetAllMessagesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeletedteamsGetallmessagesGetAllMessagesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeletedteamsGetallmessagesGetAllMessagesGetResponseable), nil
}
// ToGetRequestInformation invoke function getAllMessages
// returns a *RequestInformation when successful
func (m *DeletedteamsGetallmessagesGetAllMessagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeletedteamsGetallmessagesGetAllMessagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeletedteamsGetallmessagesGetAllMessagesRequestBuilder when successful
func (m *DeletedteamsGetallmessagesGetAllMessagesRequestBuilder) WithUrl(rawUrl string)(*DeletedteamsGetallmessagesGetAllMessagesRequestBuilder) {
    return NewDeletedteamsGetallmessagesGetAllMessagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
