package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder provides operations to manage the externalColumns property of the microsoft.graph.site entity.
type ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilderGetQueryParameters get externalColumns from groups
type ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilderGetQueryParameters
}
// NewItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilderInternal instantiates a new ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder and sets the default values.
func NewItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder) {
    m := &ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/getByPath(path='{path}')/getByPath(path='{path1}')/externalColumns{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder instantiates a new ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder and sets the default values.
func NewItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get externalColumns from groups
// returns a ColumnDefinitionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ColumnDefinitionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateColumnDefinitionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ColumnDefinitionCollectionResponseable), nil
}
// ToGetRequestInformation get externalColumns from groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
