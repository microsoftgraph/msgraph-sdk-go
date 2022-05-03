package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i02fb6a209ffec3d5fa70a95db4705fc7c75cfc91437c2fa31a8e146503996769 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies/item/devicestatuses"
    i160c96d8c11dddcb0d327977c26d60c04f4b6c89fe2e2313696b0a10646d4cce "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies/item/scheduledactionsforrule"
    i19eac586d5d9cdaadb80098949fbb380feb5000e24ac5c8fe81ccde9477b5e2a "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies/item/devicesettingstatesummaries"
    i2484457c7c0545ba42fb55ebfdec5c3410449670887d9a9294c738eb9760f4a8 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies/item/devicestatusoverview"
    i2e2c72e21cb0c8cdaec10a654e10b331974312d189a1909be8f3858df23bdb23 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies/item/assignments"
    i3030ab3b219bf633819d468663df791c240786ab0585d8a9d53bea3fca9b560e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies/item/assign"
    i724a222824d9cdbc36a1de8ffc1291a4febccc391ed1175a2c5bea1b79164a83 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies/item/userstatusoverview"
    i7279341ed743f35022dbfaf4fecf2797848a84b5e767ce5329eac709d2cd27d0 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies/item/scheduleactionsforrules"
    i9bc4838e6f15dbc4846618bfd7af51ca2f4b8feade4b5c921daf6d262ac1f5b6 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies/item/userstatuses"
    i322065cfad3bb35a1e313e9c3ea9020cbd58f06bf7ec5f4c92fbf517ea5d4439 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies/item/devicestatuses/item"
    i524f1aabf095e7a89f8751206668b1131dc6a571119b861cc6a82957868b66bb "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies/item/devicesettingstatesummaries/item"
    i52eb2204ffb8d3ff80719b2c0c36fc7747e810561f4265399b1b3fe48ce24e2e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies/item/userstatuses/item"
    i6900b8b299999e708b620425dc8c6ce05ee63496c68d74d9f14fdd267e9722fd "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies/item/assignments/item"
    i7fb30de09cd8a009b11319ff0df1e3a97b0ee6a475979d0c276ba707f9024a6b "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies/item/scheduledactionsforrule/item"
)

// DeviceCompliancePolicyItemRequestBuilder provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
type DeviceCompliancePolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceCompliancePolicyItemRequestBuilderGetQueryParameters the device compliance policies.
type DeviceCompliancePolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceCompliancePolicyItemRequestBuilderGetQueryParameters
}
// DeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign the assign property
func (m *DeviceCompliancePolicyItemRequestBuilder) Assign()(*i3030ab3b219bf633819d468663df791c240786ab0585d8a9d53bea3fca9b560e.AssignRequestBuilder) {
    return i3030ab3b219bf633819d468663df791c240786ab0585d8a9d53bea3fca9b560e.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *DeviceCompliancePolicyItemRequestBuilder) Assignments()(*i2e2c72e21cb0c8cdaec10a654e10b331974312d189a1909be8f3858df23bdb23.AssignmentsRequestBuilder) {
    return i2e2c72e21cb0c8cdaec10a654e10b331974312d189a1909be8f3858df23bdb23.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceCompliancePolicies.item.assignments.item collection
func (m *DeviceCompliancePolicyItemRequestBuilder) AssignmentsById(id string)(*i6900b8b299999e708b620425dc8c6ce05ee63496c68d74d9f14fdd267e9722fd.DeviceCompliancePolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyAssignment%2Did"] = id
    }
    return i6900b8b299999e708b620425dc8c6ce05ee63496c68d74d9f14fdd267e9722fd.NewDeviceCompliancePolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceCompliancePolicyItemRequestBuilderInternal instantiates a new DeviceCompliancePolicyItemRequestBuilder and sets the default values.
func NewDeviceCompliancePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePolicyItemRequestBuilder) {
    m := &DeviceCompliancePolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceCompliancePolicyItemRequestBuilder instantiates a new DeviceCompliancePolicyItemRequestBuilder and sets the default values.
func NewDeviceCompliancePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceCompliancePolicies for deviceManagement
func (m *DeviceCompliancePolicyItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property deviceCompliancePolicies for deviceManagement
func (m *DeviceCompliancePolicyItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the device compliance policies.
func (m *DeviceCompliancePolicyItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the device compliance policies.
func (m *DeviceCompliancePolicyItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property deviceCompliancePolicies in deviceManagement
func (m *DeviceCompliancePolicyItemRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property deviceCompliancePolicies in deviceManagement
func (m *DeviceCompliancePolicyItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyable, requestConfiguration *DeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property deviceCompliancePolicies for deviceManagement
func (m *DeviceCompliancePolicyItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property deviceCompliancePolicies for deviceManagement
func (m *DeviceCompliancePolicyItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *DeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceSettingStateSummaries the deviceSettingStateSummaries property
func (m *DeviceCompliancePolicyItemRequestBuilder) DeviceSettingStateSummaries()(*i19eac586d5d9cdaadb80098949fbb380feb5000e24ac5c8fe81ccde9477b5e2a.DeviceSettingStateSummariesRequestBuilder) {
    return i19eac586d5d9cdaadb80098949fbb380feb5000e24ac5c8fe81ccde9477b5e2a.NewDeviceSettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceSettingStateSummariesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceCompliancePolicies.item.deviceSettingStateSummaries.item collection
func (m *DeviceCompliancePolicyItemRequestBuilder) DeviceSettingStateSummariesById(id string)(*i524f1aabf095e7a89f8751206668b1131dc6a571119b861cc6a82957868b66bb.SettingStateDeviceSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["settingStateDeviceSummary%2Did"] = id
    }
    return i524f1aabf095e7a89f8751206668b1131dc6a571119b861cc6a82957868b66bb.NewSettingStateDeviceSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceStatuses the deviceStatuses property
func (m *DeviceCompliancePolicyItemRequestBuilder) DeviceStatuses()(*i02fb6a209ffec3d5fa70a95db4705fc7c75cfc91437c2fa31a8e146503996769.DeviceStatusesRequestBuilder) {
    return i02fb6a209ffec3d5fa70a95db4705fc7c75cfc91437c2fa31a8e146503996769.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceCompliancePolicies.item.deviceStatuses.item collection
func (m *DeviceCompliancePolicyItemRequestBuilder) DeviceStatusesById(id string)(*i322065cfad3bb35a1e313e9c3ea9020cbd58f06bf7ec5f4c92fbf517ea5d4439.DeviceComplianceDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceDeviceStatus%2Did"] = id
    }
    return i322065cfad3bb35a1e313e9c3ea9020cbd58f06bf7ec5f4c92fbf517ea5d4439.NewDeviceComplianceDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceStatusOverview the deviceStatusOverview property
func (m *DeviceCompliancePolicyItemRequestBuilder) DeviceStatusOverview()(*i2484457c7c0545ba42fb55ebfdec5c3410449670887d9a9294c738eb9760f4a8.DeviceStatusOverviewRequestBuilder) {
    return i2484457c7c0545ba42fb55ebfdec5c3410449670887d9a9294c738eb9760f4a8.NewDeviceStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the device compliance policies.
func (m *DeviceCompliancePolicyItemRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the device compliance policies.
func (m *DeviceCompliancePolicyItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceCompliancePolicyFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyable), nil
}
// Patch update the navigation property deviceCompliancePolicies in deviceManagement
func (m *DeviceCompliancePolicyItemRequestBuilder) Patch(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property deviceCompliancePolicies in deviceManagement
func (m *DeviceCompliancePolicyItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyable, requestConfiguration *DeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ScheduleActionsForRules the scheduleActionsForRules property
func (m *DeviceCompliancePolicyItemRequestBuilder) ScheduleActionsForRules()(*i7279341ed743f35022dbfaf4fecf2797848a84b5e767ce5329eac709d2cd27d0.ScheduleActionsForRulesRequestBuilder) {
    return i7279341ed743f35022dbfaf4fecf2797848a84b5e767ce5329eac709d2cd27d0.NewScheduleActionsForRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScheduledActionsForRule the scheduledActionsForRule property
func (m *DeviceCompliancePolicyItemRequestBuilder) ScheduledActionsForRule()(*i160c96d8c11dddcb0d327977c26d60c04f4b6c89fe2e2313696b0a10646d4cce.ScheduledActionsForRuleRequestBuilder) {
    return i160c96d8c11dddcb0d327977c26d60c04f4b6c89fe2e2313696b0a10646d4cce.NewScheduledActionsForRuleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScheduledActionsForRuleById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceCompliancePolicies.item.scheduledActionsForRule.item collection
func (m *DeviceCompliancePolicyItemRequestBuilder) ScheduledActionsForRuleById(id string)(*i7fb30de09cd8a009b11319ff0df1e3a97b0ee6a475979d0c276ba707f9024a6b.DeviceComplianceScheduledActionForRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceScheduledActionForRule%2Did"] = id
    }
    return i7fb30de09cd8a009b11319ff0df1e3a97b0ee6a475979d0c276ba707f9024a6b.NewDeviceComplianceScheduledActionForRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserStatuses the userStatuses property
func (m *DeviceCompliancePolicyItemRequestBuilder) UserStatuses()(*i9bc4838e6f15dbc4846618bfd7af51ca2f4b8feade4b5c921daf6d262ac1f5b6.UserStatusesRequestBuilder) {
    return i9bc4838e6f15dbc4846618bfd7af51ca2f4b8feade4b5c921daf6d262ac1f5b6.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatusesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceCompliancePolicies.item.userStatuses.item collection
func (m *DeviceCompliancePolicyItemRequestBuilder) UserStatusesById(id string)(*i52eb2204ffb8d3ff80719b2c0c36fc7747e810561f4265399b1b3fe48ce24e2e.DeviceComplianceUserStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceUserStatus%2Did"] = id
    }
    return i52eb2204ffb8d3ff80719b2c0c36fc7747e810561f4265399b1b3fe48ce24e2e.NewDeviceComplianceUserStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserStatusOverview the userStatusOverview property
func (m *DeviceCompliancePolicyItemRequestBuilder) UserStatusOverview()(*i724a222824d9cdbc36a1de8ffc1291a4febccc391ed1175a2c5bea1b79164a83.UserStatusOverviewRequestBuilder) {
    return i724a222824d9cdbc36a1de8ffc1291a4febccc391ed1175a2c5bea1b79164a83.NewUserStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
