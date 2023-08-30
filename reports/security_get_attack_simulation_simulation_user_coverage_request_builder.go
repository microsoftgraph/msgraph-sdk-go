package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SecurityGetAttackSimulationSimulationUserCoverageRequestBuilder provides operations to call the getAttackSimulationSimulationUserCoverage method.
type SecurityGetAttackSimulationSimulationUserCoverageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SecurityGetAttackSimulationSimulationUserCoverageRequestBuilderGetQueryParameters invoke function getAttackSimulationSimulationUserCoverage
type SecurityGetAttackSimulationSimulationUserCoverageRequestBuilderGetQueryParameters struct {
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
// SecurityGetAttackSimulationSimulationUserCoverageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityGetAttackSimulationSimulationUserCoverageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurityGetAttackSimulationSimulationUserCoverageRequestBuilderGetQueryParameters
}
// NewSecurityGetAttackSimulationSimulationUserCoverageRequestBuilderInternal instantiates a new GetAttackSimulationSimulationUserCoverageRequestBuilder and sets the default values.
func NewSecurityGetAttackSimulationSimulationUserCoverageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityGetAttackSimulationSimulationUserCoverageRequestBuilder) {
    m := &SecurityGetAttackSimulationSimulationUserCoverageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/security/getAttackSimulationSimulationUserCoverage(){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    return m
}
// NewSecurityGetAttackSimulationSimulationUserCoverageRequestBuilder instantiates a new GetAttackSimulationSimulationUserCoverageRequestBuilder and sets the default values.
func NewSecurityGetAttackSimulationSimulationUserCoverageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityGetAttackSimulationSimulationUserCoverageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityGetAttackSimulationSimulationUserCoverageRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getAttackSimulationSimulationUserCoverage
func (m *SecurityGetAttackSimulationSimulationUserCoverageRequestBuilder) Get(ctx context.Context, requestConfiguration *SecurityGetAttackSimulationSimulationUserCoverageRequestBuilderGetRequestConfiguration)(SecurityGetAttackSimulationSimulationUserCoverageResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSecurityGetAttackSimulationSimulationUserCoverageResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SecurityGetAttackSimulationSimulationUserCoverageResponseable), nil
}
// ToGetRequestInformation invoke function getAttackSimulationSimulationUserCoverage
func (m *SecurityGetAttackSimulationSimulationUserCoverageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SecurityGetAttackSimulationSimulationUserCoverageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *SecurityGetAttackSimulationSimulationUserCoverageRequestBuilder) WithUrl(rawUrl string)(*SecurityGetAttackSimulationSimulationUserCoverageRequestBuilder) {
    return NewSecurityGetAttackSimulationSimulationUserCoverageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
