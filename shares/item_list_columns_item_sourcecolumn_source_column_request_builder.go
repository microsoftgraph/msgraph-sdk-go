package shares

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemListColumnsItemSourcecolumnSourceColumnRequestBuilder provides operations to manage the sourceColumn property of the microsoft.graph.columnDefinition entity.
type ItemListColumnsItemSourcecolumnSourceColumnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListColumnsItemSourcecolumnSourceColumnRequestBuilderGetQueryParameters the source column for the content type column.
type ItemListColumnsItemSourcecolumnSourceColumnRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemListColumnsItemSourcecolumnSourceColumnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListColumnsItemSourcecolumnSourceColumnRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemListColumnsItemSourcecolumnSourceColumnRequestBuilderGetQueryParameters
}
// NewItemListColumnsItemSourcecolumnSourceColumnRequestBuilderInternal instantiates a new ItemListColumnsItemSourcecolumnSourceColumnRequestBuilder and sets the default values.
func NewItemListColumnsItemSourcecolumnSourceColumnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListColumnsItemSourcecolumnSourceColumnRequestBuilder) {
    m := &ItemListColumnsItemSourcecolumnSourceColumnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shares/{sharedDriveItem%2Did}/list/columns/{columnDefinition%2Did}/sourceColumn{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemListColumnsItemSourcecolumnSourceColumnRequestBuilder instantiates a new ItemListColumnsItemSourcecolumnSourceColumnRequestBuilder and sets the default values.
func NewItemListColumnsItemSourcecolumnSourceColumnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListColumnsItemSourcecolumnSourceColumnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListColumnsItemSourcecolumnSourceColumnRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the source column for the content type column.
// returns a ColumnDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListColumnsItemSourcecolumnSourceColumnRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListColumnsItemSourcecolumnSourceColumnRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ColumnDefinitionable, error) {
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
func (m *ItemListColumnsItemSourcecolumnSourceColumnRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListColumnsItemSourcecolumnSourceColumnRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemListColumnsItemSourcecolumnSourceColumnRequestBuilder when successful
func (m *ItemListColumnsItemSourcecolumnSourceColumnRequestBuilder) WithUrl(rawUrl string)(*ItemListColumnsItemSourcecolumnSourceColumnRequestBuilder) {
    return NewItemListColumnsItemSourcecolumnSourceColumnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
