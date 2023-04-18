package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
type DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetQueryParameters the device compliance policies.
type DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetQueryParameters
}
// DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) Assign()(*DeviceCompliancePoliciesItemAssignRequestBuilder) {
    return NewDeviceCompliancePoliciesItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) Assignments()(*DeviceCompliancePoliciesItemAssignmentsRequestBuilder) {
    return NewDeviceCompliancePoliciesItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderInternal instantiates a new DeviceCompliancePolicyItemRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) {
    m := &DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder instantiates a new DeviceCompliancePolicyItemRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceCompliancePolicies for deviceManagement
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceSettingStateSummaries provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) DeviceSettingStateSummaries()(*DeviceCompliancePoliciesItemDeviceSettingStateSummariesRequestBuilder) {
    return NewDeviceCompliancePoliciesItemDeviceSettingStateSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) DeviceStatuses()(*DeviceCompliancePoliciesItemDeviceStatusesRequestBuilder) {
    return NewDeviceCompliancePoliciesItemDeviceStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceStatusOverview provides operations to manage the deviceStatusOverview property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) DeviceStatusOverview()(*DeviceCompliancePoliciesItemDeviceStatusOverviewRequestBuilder) {
    return NewDeviceCompliancePoliciesItemDeviceStatusOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the device compliance policies.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// Patch update the navigation property deviceCompliancePolicies in deviceManagement
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyable, requestConfiguration *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) ScheduleActionsForRules()(*DeviceCompliancePoliciesItemScheduleActionsForRulesRequestBuilder) {
    return NewDeviceCompliancePoliciesItemScheduleActionsForRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ScheduledActionsForRule provides operations to manage the scheduledActionsForRule property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) ScheduledActionsForRule()(*DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder) {
    return NewDeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property deviceCompliancePolicies for deviceManagement
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the device compliance policies.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceCompliancePolicies in deviceManagement
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyable, requestConfiguration *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// UserStatuses provides operations to manage the userStatuses property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) UserStatuses()(*DeviceCompliancePoliciesItemUserStatusesRequestBuilder) {
    return NewDeviceCompliancePoliciesItemUserStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserStatusOverview provides operations to manage the userStatusOverview property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) UserStatusOverview()(*DeviceCompliancePoliciesItemUserStatusOverviewRequestBuilder) {
    return NewDeviceCompliancePoliciesItemUserStatusOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
