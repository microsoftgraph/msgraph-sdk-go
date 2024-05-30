package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilder provides operations to call the getAttackSimulationTrainingUserCoverage method.
type SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilderGetQueryParameters list training coverage for tenant users in attack simulation and training campaigns. This function supports @odata.nextLink for pagination.
type SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilderGetQueryParameters struct {
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
// SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilderGetQueryParameters
}
// NewSecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilderInternal instantiates a new SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilder and sets the default values.
func NewSecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilder) {
    m := &SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/security/getAttackSimulationTrainingUserCoverage(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewSecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilder instantiates a new SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilder and sets the default values.
func NewSecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list training coverage for tenant users in attack simulation and training campaigns. This function supports @odata.nextLink for pagination.
// Deprecated: This method is obsolete. Use GetAsGetAttackSimulationTrainingUserCoverageGetResponse instead.
// returns a SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/securityreportsroot-getattacksimulationtrainingusercoverage?view=graph-rest-1.0
func (m *SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilder) Get(ctx context.Context, requestConfiguration *SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilderGetRequestConfiguration)(SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageResponseable), nil
}
// GetAsGetAttackSimulationTrainingUserCoverageGetResponse list training coverage for tenant users in attack simulation and training campaigns. This function supports @odata.nextLink for pagination.
// returns a SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/securityreportsroot-getattacksimulationtrainingusercoverage?view=graph-rest-1.0
func (m *SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilder) GetAsGetAttackSimulationTrainingUserCoverageGetResponse(ctx context.Context, requestConfiguration *SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilderGetRequestConfiguration)(SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageGetResponseable), nil
}
// ToGetRequestInformation list training coverage for tenant users in attack simulation and training campaigns. This function supports @odata.nextLink for pagination.
// returns a *RequestInformation when successful
func (m *SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilder when successful
func (m *SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilder) WithUrl(rawUrl string)(*SecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilder) {
    return NewSecurityGetattacksimulationtrainingusercoverageGetAttackSimulationTrainingUserCoverageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
