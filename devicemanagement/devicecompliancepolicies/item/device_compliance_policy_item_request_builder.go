package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceCompliancePolicyItemRequestBuilderDeleteOptions options for Delete
type DeviceCompliancePolicyItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceCompliancePolicyItemRequestBuilderGetOptions options for Get
type DeviceCompliancePolicyItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceCompliancePolicyItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceCompliancePolicyItemRequestBuilderGetQueryParameters the device compliance policies.
type DeviceCompliancePolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceCompliancePolicyItemRequestBuilderPatchOptions options for Patch
type DeviceCompliancePolicyItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceCompliancePolicyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceCompliancePolicyItemRequestBuilder) Assign()(*i3030ab3b219bf633819d468663df791c240786ab0585d8a9d53bea3fca9b560e.AssignRequestBuilder) {
    return i3030ab3b219bf633819d468663df791c240786ab0585d8a9d53bea3fca9b560e.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["deviceCompliancePolicyAssignment_id"] = id
    }
    return i6900b8b299999e708b620425dc8c6ce05ee63496c68d74d9f14fdd267e9722fd.NewDeviceCompliancePolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceCompliancePolicyItemRequestBuilderInternal instantiates a new DeviceCompliancePolicyItemRequestBuilder and sets the default values.
func NewDeviceCompliancePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceCompliancePolicyItemRequestBuilder) {
    m := &DeviceCompliancePolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceCompliancePolicyItemRequestBuilder instantiates a new DeviceCompliancePolicyItemRequestBuilder and sets the default values.
func NewDeviceCompliancePolicyItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceCompliancePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceCompliancePolicies for deviceManagement
func (m *DeviceCompliancePolicyItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceCompliancePolicyItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the device compliance policies.
func (m *DeviceCompliancePolicyItemRequestBuilder) CreateGetRequestInformation(options *DeviceCompliancePolicyItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property deviceCompliancePolicies in deviceManagement
func (m *DeviceCompliancePolicyItemRequestBuilder) CreatePatchRequestInformation(options *DeviceCompliancePolicyItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property deviceCompliancePolicies for deviceManagement
func (m *DeviceCompliancePolicyItemRequestBuilder) Delete(options *DeviceCompliancePolicyItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["settingStateDeviceSummary_id"] = id
    }
    return i524f1aabf095e7a89f8751206668b1131dc6a571119b861cc6a82957868b66bb.NewSettingStateDeviceSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["deviceComplianceDeviceStatus_id"] = id
    }
    return i322065cfad3bb35a1e313e9c3ea9020cbd58f06bf7ec5f4c92fbf517ea5d4439.NewDeviceComplianceDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceCompliancePolicyItemRequestBuilder) DeviceStatusOverview()(*i2484457c7c0545ba42fb55ebfdec5c3410449670887d9a9294c738eb9760f4a8.DeviceStatusOverviewRequestBuilder) {
    return i2484457c7c0545ba42fb55ebfdec5c3410449670887d9a9294c738eb9760f4a8.NewDeviceStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the device compliance policies.
func (m *DeviceCompliancePolicyItemRequestBuilder) Get(options *DeviceCompliancePolicyItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceCompliancePolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDeviceCompliancePolicyFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceCompliancePolicyable), nil
}
// Patch update the navigation property deviceCompliancePolicies in deviceManagement
func (m *DeviceCompliancePolicyItemRequestBuilder) Patch(options *DeviceCompliancePolicyItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceCompliancePolicyItemRequestBuilder) ScheduleActionsForRules()(*i7279341ed743f35022dbfaf4fecf2797848a84b5e767ce5329eac709d2cd27d0.ScheduleActionsForRulesRequestBuilder) {
    return i7279341ed743f35022dbfaf4fecf2797848a84b5e767ce5329eac709d2cd27d0.NewScheduleActionsForRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["deviceComplianceScheduledActionForRule_id"] = id
    }
    return i7fb30de09cd8a009b11319ff0df1e3a97b0ee6a475979d0c276ba707f9024a6b.NewDeviceComplianceScheduledActionForRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["deviceComplianceUserStatus_id"] = id
    }
    return i52eb2204ffb8d3ff80719b2c0c36fc7747e810561f4265399b1b3fe48ce24e2e.NewDeviceComplianceUserStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceCompliancePolicyItemRequestBuilder) UserStatusOverview()(*i724a222824d9cdbc36a1de8ffc1291a4febccc391ed1175a2c5bea1b79164a83.UserStatusOverviewRequestBuilder) {
    return i724a222824d9cdbc36a1de8ffc1291a4febccc391ed1175a2c5bea1b79164a83.NewUserStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
