package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilder provides operations to call the getAttackSimulationSimulationUserCoverage method.
type SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilderGetQueryParameters list training coverage for each tenant user in attack simulation and training campaigns. This function supports @odata.nextLink for pagination.
type SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilderGetQueryParameters struct {
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
// SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilderGetQueryParameters
}
// NewSecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilderInternal instantiates a new SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilder and sets the default values.
func NewSecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilder) {
    m := &SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/security/getAttackSimulationSimulationUserCoverage(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewSecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilder instantiates a new SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilder and sets the default values.
func NewSecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list training coverage for each tenant user in attack simulation and training campaigns. This function supports @odata.nextLink for pagination.
// Deprecated: This method is obsolete. Use GetAsGetAttackSimulationSimulationUserCoverageGetResponse instead.
// returns a SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/securityreportsroot-getattacksimulationsimulationusercoverage?view=graph-rest-1.0
func (m *SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilder) Get(ctx context.Context, requestConfiguration *SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilderGetRequestConfiguration)(SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageResponseable), nil
}
// GetAsGetAttackSimulationSimulationUserCoverageGetResponse list training coverage for each tenant user in attack simulation and training campaigns. This function supports @odata.nextLink for pagination.
// returns a SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/securityreportsroot-getattacksimulationsimulationusercoverage?view=graph-rest-1.0
func (m *SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilder) GetAsGetAttackSimulationSimulationUserCoverageGetResponse(ctx context.Context, requestConfiguration *SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilderGetRequestConfiguration)(SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageGetResponseable), nil
}
// ToGetRequestInformation list training coverage for each tenant user in attack simulation and training campaigns. This function supports @odata.nextLink for pagination.
// returns a *RequestInformation when successful
func (m *SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilder when successful
func (m *SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilder) WithUrl(rawUrl string)(*SecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilder) {
    return NewSecurityGetattacksimulationsimulationusercoverageGetAttackSimulationSimulationUserCoverageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
