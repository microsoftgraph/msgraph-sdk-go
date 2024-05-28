package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder provides operations to call the getActivitiesByInterval method.
type ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetQueryParameters invoke function getActivitiesByInterval
type ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetQueryParameters
}
// NewItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderInternal instantiates a new ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder and sets the default values.
func NewItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) {
    m := &ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/getByPath(path='{path}')/getActivitiesByInterval(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder instantiates a new ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder and sets the default values.
func NewItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getActivitiesByInterval
// Deprecated: This method is obsolete. Use GetAsGetActivitiesByIntervalGetResponse instead.
// returns a ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetRequestConfiguration)(ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalResponseable), nil
}
// GetAsGetActivitiesByIntervalGetResponse invoke function getActivitiesByInterval
// returns a ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) GetAsGetActivitiesByIntervalGetResponse(ctx context.Context, requestConfiguration *ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetRequestConfiguration)(ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalGetResponseable), nil
}
// ToGetRequestInformation invoke function getActivitiesByInterval
// returns a *RequestInformation when successful
func (m *ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
