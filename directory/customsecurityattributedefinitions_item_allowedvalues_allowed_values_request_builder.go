package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder provides operations to manage the allowedValues property of the microsoft.graph.customSecurityAttributeDefinition entity.
type CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilderGetQueryParameters get a list of the allowedValue objects and their properties.
type CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilderGetQueryParameters struct {
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
// CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilderGetQueryParameters
}
// CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAllowedValueId provides operations to manage the allowedValues property of the microsoft.graph.customSecurityAttributeDefinition entity.
// returns a *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder when successful
func (m *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder) ByAllowedValueId(allowedValueId string)(*CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if allowedValueId != "" {
        urlTplParams["allowedValue%2Did"] = allowedValueId
    }
    return NewCustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilderInternal instantiates a new CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder and sets the default values.
func NewCustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder) {
    m := &CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/customSecurityAttributeDefinitions/{customSecurityAttributeDefinition%2Did}/allowedValues{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder instantiates a new CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder and sets the default values.
func NewCustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CustomsecurityattributedefinitionsItemAllowedvaluesCountRequestBuilder when successful
func (m *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder) Count()(*CustomsecurityattributedefinitionsItemAllowedvaluesCountRequestBuilder) {
    return NewCustomsecurityattributedefinitionsItemAllowedvaluesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the allowedValue objects and their properties.
// returns a AllowedValueCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/customsecurityattributedefinition-list-allowedvalues?view=graph-rest-1.0
func (m *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder) Get(ctx context.Context, requestConfiguration *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AllowedValueCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAllowedValueCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AllowedValueCollectionResponseable), nil
}
// Post create a new allowedValue object.
// returns a AllowedValueable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/customsecurityattributedefinition-post-allowedvalues?view=graph-rest-1.0
func (m *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AllowedValueable, requestConfiguration *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AllowedValueable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAllowedValueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AllowedValueable), nil
}
// ToGetRequestInformation get a list of the allowedValue objects and their properties.
// returns a *RequestInformation when successful
func (m *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new allowedValue object.
// returns a *RequestInformation when successful
func (m *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AllowedValueable, requestConfiguration *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder when successful
func (m *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder) WithUrl(rawUrl string)(*CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder) {
    return NewCustomsecurityattributedefinitionsItemAllowedvaluesAllowedValuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
