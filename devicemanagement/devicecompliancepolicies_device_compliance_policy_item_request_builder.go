package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
type DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderGetQueryParameters read properties and relationships of the windowsPhone81CompliancePolicy object.
type DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderGetQueryParameters
}
// DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *DevicecompliancepoliciesItemAssignRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) Assign()(*DevicecompliancepoliciesItemAssignRequestBuilder) {
    return NewDevicecompliancepoliciesItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceCompliancePolicy entity.
// returns a *DevicecompliancepoliciesItemAssignmentsRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) Assignments()(*DevicecompliancepoliciesItemAssignmentsRequestBuilder) {
    return NewDevicecompliancepoliciesItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderInternal instantiates a new DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) {
    m := &DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder instantiates a new DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a iosCompliancePolicy.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-ioscompliancepolicy-delete?view=graph-rest-1.0
func (m *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceSettingStateSummaries provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceCompliancePolicy entity.
// returns a *DevicecompliancepoliciesItemDevicesettingstatesummariesDeviceSettingStateSummariesRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) DeviceSettingStateSummaries()(*DevicecompliancepoliciesItemDevicesettingstatesummariesDeviceSettingStateSummariesRequestBuilder) {
    return NewDevicecompliancepoliciesItemDevicesettingstatesummariesDeviceSettingStateSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.deviceCompliancePolicy entity.
// returns a *DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) DeviceStatuses()(*DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder) {
    return NewDevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceStatusOverview provides operations to manage the deviceStatusOverview property of the microsoft.graph.deviceCompliancePolicy entity.
// returns a *DevicecompliancepoliciesItemDevicestatusoverviewDeviceStatusOverviewRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) DeviceStatusOverview()(*DevicecompliancepoliciesItemDevicestatusoverviewDeviceStatusOverviewRequestBuilder) {
    return NewDevicecompliancepoliciesItemDevicestatusoverviewDeviceStatusOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read properties and relationships of the windowsPhone81CompliancePolicy object.
// returns a DeviceCompliancePolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-windowsphone81compliancepolicy-get?view=graph-rest-1.0
func (m *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceCompliancePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyable), nil
}
// Patch update the properties of a macOSCompliancePolicy object.
// returns a DeviceCompliancePolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-macoscompliancepolicy-update?view=graph-rest-1.0
func (m *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyable, requestConfiguration *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceCompliancePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyable), nil
}
// ScheduleActionsForRules provides operations to call the scheduleActionsForRules method.
// returns a *DevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) ScheduleActionsForRules()(*DevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesRequestBuilder) {
    return NewDevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ScheduledActionsForRule provides operations to manage the scheduledActionsForRule property of the microsoft.graph.deviceCompliancePolicy entity.
// returns a *DevicecompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) ScheduledActionsForRule()(*DevicecompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder) {
    return NewDevicecompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes a iosCompliancePolicy.
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the windowsPhone81CompliancePolicy object.
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a macOSCompliancePolicy object.
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyable, requestConfiguration *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserStatuses provides operations to manage the userStatuses property of the microsoft.graph.deviceCompliancePolicy entity.
// returns a *DevicecompliancepoliciesItemUserstatusesUserStatusesRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) UserStatuses()(*DevicecompliancepoliciesItemUserstatusesUserStatusesRequestBuilder) {
    return NewDevicecompliancepoliciesItemUserstatusesUserStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserStatusOverview provides operations to manage the userStatusOverview property of the microsoft.graph.deviceCompliancePolicy entity.
// returns a *DevicecompliancepoliciesItemUserstatusoverviewUserStatusOverviewRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) UserStatusOverview()(*DevicecompliancepoliciesItemUserstatusoverviewUserStatusOverviewRequestBuilder) {
    return NewDevicecompliancepoliciesItemUserstatusoverviewUserStatusOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) WithUrl(rawUrl string)(*DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) {
    return NewDevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
