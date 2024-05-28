package shares

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
type ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderGetQueryParameters column order information in a content type.
type ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderGetQueryParameters
}
// NewItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderInternal instantiates a new ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder and sets the default values.
func NewItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder) {
    m := &ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shares/{sharedDriveItem%2Did}/list/contentTypes/{contentType%2Did}/columnPositions/{columnDefinition%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder instantiates a new ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder and sets the default values.
func NewItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get column order information in a content type.
// returns a ColumnDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ColumnDefinitionable, error) {
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
// ToGetRequestInformation column order information in a content type.
// returns a *RequestInformation when successful
func (m *ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder when successful
func (m *ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*ItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder) {
    return NewItemListContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
