package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
type ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderGetQueryParameters column order information in a content type.
type ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderGetQueryParameters
}
// NewItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderInternal instantiates a new ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder and sets the default values.
func NewItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder) {
    m := &ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/contentTypes/{contentType%2Did}/columnPositions/{columnDefinition%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder instantiates a new ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder and sets the default values.
func NewItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get column order information in a content type.
// returns a ColumnDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ColumnDefinitionable, error) {
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
func (m *ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder when successful
func (m *ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*ItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder) {
    return NewItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
