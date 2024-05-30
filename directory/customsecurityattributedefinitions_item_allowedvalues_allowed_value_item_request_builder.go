package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder provides operations to manage the allowedValues property of the microsoft.graph.customSecurityAttributeDefinition entity.
type CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderGetQueryParameters read the properties and relationships of an allowedValue object.
type CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderGetQueryParameters
}
// CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderInternal instantiates a new CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder and sets the default values.
func NewCustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder) {
    m := &CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/customSecurityAttributeDefinitions/{customSecurityAttributeDefinition%2Did}/allowedValues/{allowedValue%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder instantiates a new CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder and sets the default values.
func NewCustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property allowedValues for directory
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read the properties and relationships of an allowedValue object.
// returns a AllowedValueable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/allowedvalue-get?view=graph-rest-1.0
func (m *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AllowedValueable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the properties of an allowedValue object.
// returns a AllowedValueable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/allowedvalue-update?view=graph-rest-1.0
func (m *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AllowedValueable, requestConfiguration *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AllowedValueable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property allowedValues for directory
// returns a *RequestInformation when successful
func (m *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an allowedValue object.
// returns a *RequestInformation when successful
func (m *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an allowedValue object.
// returns a *RequestInformation when successful
func (m *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AllowedValueable, requestConfiguration *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder when successful
func (m *CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder) WithUrl(rawUrl string)(*CustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder) {
    return NewCustomsecurityattributedefinitionsItemAllowedvaluesAllowedValueItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
