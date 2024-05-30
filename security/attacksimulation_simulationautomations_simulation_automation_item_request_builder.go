package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder provides operations to manage the simulationAutomations property of the microsoft.graph.attackSimulationRoot entity.
type AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderGetQueryParameters get an attack simulation automation for a tenant.
type AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderGetQueryParameters
}
// AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderInternal instantiates a new AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder and sets the default values.
func NewAttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder) {
    m := &AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/attackSimulation/simulationAutomations/{simulationAutomation%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder instantiates a new AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder and sets the default values.
func NewAttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property simulationAutomations for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get an attack simulation automation for a tenant.
// returns a SimulationAutomationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/simulationautomation-get?view=graph-rest-1.0
func (m *AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSimulationAutomationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationable), nil
}
// Patch update the navigation property simulationAutomations in security
// returns a SimulationAutomationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationable, requestConfiguration *AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSimulationAutomationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationable), nil
}
// Runs provides operations to manage the runs property of the microsoft.graph.simulationAutomation entity.
// returns a *AttacksimulationSimulationautomationsItemRunsRequestBuilder when successful
func (m *AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder) Runs()(*AttacksimulationSimulationautomationsItemRunsRequestBuilder) {
    return NewAttacksimulationSimulationautomationsItemRunsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property simulationAutomations for security
// returns a *RequestInformation when successful
func (m *AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get an attack simulation automation for a tenant.
// returns a *RequestInformation when successful
func (m *AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property simulationAutomations in security
// returns a *RequestInformation when successful
func (m *AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationable, requestConfiguration *AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder when successful
func (m *AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder) WithUrl(rawUrl string)(*AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder) {
    return NewAttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
