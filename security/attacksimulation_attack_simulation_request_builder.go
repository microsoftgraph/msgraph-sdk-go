package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AttacksimulationAttackSimulationRequestBuilder provides operations to manage the attackSimulation property of the microsoft.graph.security entity.
type AttacksimulationAttackSimulationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AttacksimulationAttackSimulationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationAttackSimulationRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AttacksimulationAttackSimulationRequestBuilderGetQueryParameters get attackSimulation from security
type AttacksimulationAttackSimulationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AttacksimulationAttackSimulationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationAttackSimulationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttacksimulationAttackSimulationRequestBuilderGetQueryParameters
}
// AttacksimulationAttackSimulationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationAttackSimulationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAttacksimulationAttackSimulationRequestBuilderInternal instantiates a new AttacksimulationAttackSimulationRequestBuilder and sets the default values.
func NewAttacksimulationAttackSimulationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationAttackSimulationRequestBuilder) {
    m := &AttacksimulationAttackSimulationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/attackSimulation{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAttacksimulationAttackSimulationRequestBuilder instantiates a new AttacksimulationAttackSimulationRequestBuilder and sets the default values.
func NewAttacksimulationAttackSimulationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationAttackSimulationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttacksimulationAttackSimulationRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property attackSimulation for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationAttackSimulationRequestBuilder) Delete(ctx context.Context, requestConfiguration *AttacksimulationAttackSimulationRequestBuilderDeleteRequestConfiguration)(error) {
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
// EndUserNotifications provides operations to manage the endUserNotifications property of the microsoft.graph.attackSimulationRoot entity.
// returns a *AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder when successful
func (m *AttacksimulationAttackSimulationRequestBuilder) EndUserNotifications()(*AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder) {
    return NewAttacksimulationEndusernotificationsEndUserNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get attackSimulation from security
// returns a AttackSimulationRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationAttackSimulationRequestBuilder) Get(ctx context.Context, requestConfiguration *AttacksimulationAttackSimulationRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttackSimulationRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAttackSimulationRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttackSimulationRootable), nil
}
// LandingPages provides operations to manage the landingPages property of the microsoft.graph.attackSimulationRoot entity.
// returns a *AttacksimulationLandingpagesLandingPagesRequestBuilder when successful
func (m *AttacksimulationAttackSimulationRequestBuilder) LandingPages()(*AttacksimulationLandingpagesLandingPagesRequestBuilder) {
    return NewAttacksimulationLandingpagesLandingPagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LoginPages provides operations to manage the loginPages property of the microsoft.graph.attackSimulationRoot entity.
// returns a *AttacksimulationLoginpagesLoginPagesRequestBuilder when successful
func (m *AttacksimulationAttackSimulationRequestBuilder) LoginPages()(*AttacksimulationLoginpagesLoginPagesRequestBuilder) {
    return NewAttacksimulationLoginpagesLoginPagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.attackSimulationRoot entity.
// returns a *AttacksimulationOperationsRequestBuilder when successful
func (m *AttacksimulationAttackSimulationRequestBuilder) Operations()(*AttacksimulationOperationsRequestBuilder) {
    return NewAttacksimulationOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property attackSimulation in security
// returns a AttackSimulationRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationAttackSimulationRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttackSimulationRootable, requestConfiguration *AttacksimulationAttackSimulationRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttackSimulationRootable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAttackSimulationRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttackSimulationRootable), nil
}
// Payloads provides operations to manage the payloads property of the microsoft.graph.attackSimulationRoot entity.
// returns a *AttacksimulationPayloadsRequestBuilder when successful
func (m *AttacksimulationAttackSimulationRequestBuilder) Payloads()(*AttacksimulationPayloadsRequestBuilder) {
    return NewAttacksimulationPayloadsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SimulationAutomations provides operations to manage the simulationAutomations property of the microsoft.graph.attackSimulationRoot entity.
// returns a *AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder when successful
func (m *AttacksimulationAttackSimulationRequestBuilder) SimulationAutomations()(*AttacksimulationSimulationautomationsSimulationAutomationsRequestBuilder) {
    return NewAttacksimulationSimulationautomationsSimulationAutomationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Simulations provides operations to manage the simulations property of the microsoft.graph.attackSimulationRoot entity.
// returns a *AttacksimulationSimulationsRequestBuilder when successful
func (m *AttacksimulationAttackSimulationRequestBuilder) Simulations()(*AttacksimulationSimulationsRequestBuilder) {
    return NewAttacksimulationSimulationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property attackSimulation for security
// returns a *RequestInformation when successful
func (m *AttacksimulationAttackSimulationRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AttacksimulationAttackSimulationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get attackSimulation from security
// returns a *RequestInformation when successful
func (m *AttacksimulationAttackSimulationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AttacksimulationAttackSimulationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property attackSimulation in security
// returns a *RequestInformation when successful
func (m *AttacksimulationAttackSimulationRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttackSimulationRootable, requestConfiguration *AttacksimulationAttackSimulationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Trainings provides operations to manage the trainings property of the microsoft.graph.attackSimulationRoot entity.
// returns a *AttacksimulationTrainingsRequestBuilder when successful
func (m *AttacksimulationAttackSimulationRequestBuilder) Trainings()(*AttacksimulationTrainingsRequestBuilder) {
    return NewAttacksimulationTrainingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AttacksimulationAttackSimulationRequestBuilder when successful
func (m *AttacksimulationAttackSimulationRequestBuilder) WithUrl(rawUrl string)(*AttacksimulationAttackSimulationRequestBuilder) {
    return NewAttacksimulationAttackSimulationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
