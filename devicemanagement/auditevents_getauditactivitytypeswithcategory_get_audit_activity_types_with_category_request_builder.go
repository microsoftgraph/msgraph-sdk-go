package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder provides operations to call the getAuditActivityTypes method.
type AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilderGetQueryParameters invoke function getAuditActivityTypes
type AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilderGetQueryParameters
}
// NewAuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilderInternal instantiates a new AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder and sets the default values.
func NewAuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, category *string)(*AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder) {
    m := &AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/auditEvents/getAuditActivityTypes(category='{category}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if category != nil {
        m.BaseRequestBuilder.PathParameters["category"] = *category
    }
    return m
}
// NewAuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder instantiates a new AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder and sets the default values.
func NewAuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getAuditActivityTypes
// Deprecated: This method is obsolete. Use GetAsGetAuditActivityTypesWithCategoryGetResponse instead.
// returns a AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder) Get(ctx context.Context, requestConfiguration *AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilderGetRequestConfiguration)(AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryResponseable), nil
}
// GetAsGetAuditActivityTypesWithCategoryGetResponse invoke function getAuditActivityTypes
// returns a AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder) GetAsGetAuditActivityTypesWithCategoryGetResponse(ctx context.Context, requestConfiguration *AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilderGetRequestConfiguration)(AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryGetResponseable), nil
}
// ToGetRequestInformation invoke function getAuditActivityTypes
// returns a *RequestInformation when successful
func (m *AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder when successful
func (m *AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder) WithUrl(rawUrl string)(*AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder) {
    return NewAuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
