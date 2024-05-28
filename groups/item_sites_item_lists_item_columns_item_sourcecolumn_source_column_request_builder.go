package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilder provides operations to manage the sourceColumn property of the microsoft.graph.columnDefinition entity.
type ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilderGetQueryParameters the source column for the content type column.
type ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilderGetQueryParameters
}
// NewItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilderInternal instantiates a new ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilder and sets the default values.
func NewItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilder) {
    m := &ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/columns/{columnDefinition%2Did}/sourceColumn{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilder instantiates a new ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilder and sets the default values.
func NewItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the source column for the content type column.
// returns a ColumnDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ColumnDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateColumnDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ColumnDefinitionable), nil
}
// ToGetRequestInformation the source column for the content type column.
// returns a *RequestInformation when successful
func (m *ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilder when successful
func (m *ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilder) {
    return NewItemSitesItemListsItemColumnsItemSourcecolumnSourceColumnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
