package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder provides operations to call the getApplicableContentTypesForList method.
type ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderGetQueryParameters get site contentTypes that can be added to a list.
type ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderGetQueryParameters
}
// NewItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderInternal instantiates a new ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder and sets the default values.
func NewItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, listId *string)(*ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder) {
    m := &ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/getByPath(path='{path}')/getApplicableContentTypesForList(listId='{listId}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if listId != nil {
        m.BaseRequestBuilder.PathParameters["listId"] = *listId
    }
    return m
}
// NewItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder instantiates a new ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder and sets the default values.
func NewItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get site contentTypes that can be added to a list.
// Deprecated: This method is obsolete. Use GetAsGetApplicableContentTypesForListWithListIdGetResponse instead.
// returns a ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/site-getapplicablecontenttypesforlist?view=graph-rest-1.0
func (m *ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderGetRequestConfiguration)(ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponseable), nil
}
// GetAsGetApplicableContentTypesForListWithListIdGetResponse get site contentTypes that can be added to a list.
// returns a ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/site-getapplicablecontenttypesforlist?view=graph-rest-1.0
func (m *ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder) GetAsGetApplicableContentTypesForListWithListIdGetResponse(ctx context.Context, requestConfiguration *ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderGetRequestConfiguration)(ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdGetResponseable), nil
}
// ToGetRequestInformation get site contentTypes that can be added to a list.
// returns a *RequestInformation when successful
func (m *ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
