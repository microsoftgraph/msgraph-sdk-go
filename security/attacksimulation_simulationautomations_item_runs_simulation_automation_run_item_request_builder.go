package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder provides operations to manage the runs property of the microsoft.graph.simulationAutomation entity.
type AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderGetQueryParameters a collection of simulation automation runs.
type AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderGetQueryParameters
}
// AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderInternal instantiates a new AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder and sets the default values.
func NewAttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder) {
    m := &AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/attackSimulation/simulationAutomations/{simulationAutomation%2Did}/runs/{simulationAutomationRun%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder instantiates a new AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder and sets the default values.
func NewAttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property runs for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of simulation automation runs.
// returns a SimulationAutomationRunable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationRunable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSimulationAutomationRunFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationRunable), nil
}
// Patch update the navigation property runs in security
// returns a SimulationAutomationRunable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationRunable, requestConfiguration *AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationRunable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSimulationAutomationRunFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationRunable), nil
}
// ToDeleteRequestInformation delete navigation property runs for security
// returns a *RequestInformation when successful
func (m *AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of simulation automation runs.
// returns a *RequestInformation when successful
func (m *AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property runs in security
// returns a *RequestInformation when successful
func (m *AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationRunable, requestConfiguration *AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder when successful
func (m *AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder) WithUrl(rawUrl string)(*AttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder) {
    return NewAttacksimulationSimulationautomationsItemRunsSimulationAutomationRunItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
