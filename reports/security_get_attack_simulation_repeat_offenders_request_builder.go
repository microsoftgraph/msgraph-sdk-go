package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SecurityGetAttackSimulationRepeatOffendersRequestBuilder provides operations to call the getAttackSimulationRepeatOffenders method.
type SecurityGetAttackSimulationRepeatOffendersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SecurityGetAttackSimulationRepeatOffendersRequestBuilderGetQueryParameters invoke function getAttackSimulationRepeatOffenders
type SecurityGetAttackSimulationRepeatOffendersRequestBuilderGetQueryParameters struct {
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
// SecurityGetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityGetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurityGetAttackSimulationRepeatOffendersRequestBuilderGetQueryParameters
}
// NewSecurityGetAttackSimulationRepeatOffendersRequestBuilderInternal instantiates a new GetAttackSimulationRepeatOffendersRequestBuilder and sets the default values.
func NewSecurityGetAttackSimulationRepeatOffendersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityGetAttackSimulationRepeatOffendersRequestBuilder) {
    m := &SecurityGetAttackSimulationRepeatOffendersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/security/getAttackSimulationRepeatOffenders(){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    return m
}
// NewSecurityGetAttackSimulationRepeatOffendersRequestBuilder instantiates a new GetAttackSimulationRepeatOffendersRequestBuilder and sets the default values.
func NewSecurityGetAttackSimulationRepeatOffendersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityGetAttackSimulationRepeatOffendersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityGetAttackSimulationRepeatOffendersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getAttackSimulationRepeatOffenders
func (m *SecurityGetAttackSimulationRepeatOffendersRequestBuilder) Get(ctx context.Context, requestConfiguration *SecurityGetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration)(SecurityGetAttackSimulationRepeatOffendersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSecurityGetAttackSimulationRepeatOffendersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SecurityGetAttackSimulationRepeatOffendersResponseable), nil
}
// ToGetRequestInformation invoke function getAttackSimulationRepeatOffenders
func (m *SecurityGetAttackSimulationRepeatOffendersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SecurityGetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *SecurityGetAttackSimulationRepeatOffendersRequestBuilder) WithUrl(rawUrl string)(*SecurityGetAttackSimulationRepeatOffendersRequestBuilder) {
    return NewSecurityGetAttackSimulationRepeatOffendersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
