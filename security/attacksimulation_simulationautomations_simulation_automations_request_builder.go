package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder provides operations to manage the simulationAutomations property of the microsoft.graph.attackSimulationRoot entity.
type AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilderGetQueryParameters get a list of attack simulation automations for a tenant.
type AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilderGetQueryParameters struct {
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
// AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilderGetQueryParameters
}
// AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySimulationAutomationId provides operations to manage the simulationAutomations property of the microsoft.graph.attackSimulationRoot entity.
// returns a *AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder when successful
func (m *AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder) BySimulationAutomationId(simulationAutomationId string)(*AttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if simulationAutomationId != "" {
        urlTplParams["simulationAutomation%2Did"] = simulationAutomationId
    }
    return NewAttacksimulationSimulationautomationsSimulationAutomationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAttacksimulationSimulationautomationsSimulationAutomationsRequestBuilderInternal instantiates a new AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder and sets the default values.
func NewAttacksimulationSimulationautomationsSimulationAutomationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder) {
    m := &AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/attackSimulation/simulationAutomations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder instantiates a new AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder and sets the default values.
func NewAttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttacksimulationSimulationautomationsSimulationAutomationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AttacksimulationSimulationautomationsCountRequestBuilder when successful
func (m *AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder) Count()(*AttacksimulationSimulationautomationsCountRequestBuilder) {
    return NewAttacksimulationSimulationautomationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of attack simulation automations for a tenant.
// returns a SimulationAutomationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/attacksimulationroot-list-simulationautomations?view=graph-rest-1.0
func (m *AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder) Get(ctx context.Context, requestConfiguration *AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSimulationAutomationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationCollectionResponseable), nil
}
// Post create new navigation property to simulationAutomations for security
// returns a SimulationAutomationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationable, requestConfiguration *AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation get a list of attack simulation automations for a tenant.
// returns a *RequestInformation when successful
func (m *AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to simulationAutomations for security
// returns a *RequestInformation when successful
func (m *AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationable, requestConfiguration *AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder when successful
func (m *AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder) WithUrl(rawUrl string)(*AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder) {
    return NewAttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
