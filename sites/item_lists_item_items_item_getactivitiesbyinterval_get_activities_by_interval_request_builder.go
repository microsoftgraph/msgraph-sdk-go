package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder provides operations to call the getActivitiesByInterval method.
type ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetQueryParameters invoke function getActivitiesByInterval
type ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetQueryParameters struct {
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
// ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetQueryParameters
}
// NewItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderInternal instantiates a new ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder and sets the default values.
func NewItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) {
    m := &ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}/getActivitiesByInterval(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder instantiates a new ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder and sets the default values.
func NewItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getActivitiesByInterval
// Deprecated: This method is obsolete. Use GetAsGetActivitiesByIntervalGetResponse instead.
// returns a ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetRequestConfiguration)(ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponseable), nil
}
// GetAsGetActivitiesByIntervalGetResponse invoke function getActivitiesByInterval
// returns a ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) GetAsGetActivitiesByIntervalGetResponse(ctx context.Context, requestConfiguration *ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetRequestConfiguration)(ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponseable), nil
}
// ToGetRequestInformation invoke function getActivitiesByInterval
// returns a *RequestInformation when successful
func (m *ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder when successful
func (m *ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) WithUrl(rawUrl string)(*ItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) {
    return NewItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
