package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
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

// deviceCompliancePolicyRequestBuilder builds and executes requests for operations under \deviceManagement\deviceCompliancePolicies\{deviceCompliancePolicy-id}
type DeviceCompliancePolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceCompliancePolicyRequestBuilderDeleteOptions options for Delete
type DeviceCompliancePolicyRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceCompliancePolicyRequestBuilderGetOptions options for Get
type DeviceCompliancePolicyRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceCompliancePolicyRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// deviceCompliancePolicyRequestBuilderGetQueryParameters the device compliance policies.
type DeviceCompliancePolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// DeviceCompliancePolicyRequestBuilderPatchOptions options for Patch
type DeviceCompliancePolicyRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceCompliancePolicy;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceCompliancePolicyRequestBuilder) Assign()(*i3030ab3b219bf633819d468663df791c240786ab0585d8a9d53bea3fca9b560e.AssignRequestBuilder) {
    return i3030ab3b219bf633819d468663df791c240786ab0585d8a9d53bea3fca9b560e.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) Assignments()(*i2e2c72e21cb0c8cdaec10a654e10b331974312d189a1909be8f3858df23bdb23.AssignmentsRequestBuilder) {
    return i2e2c72e21cb0c8cdaec10a654e10b331974312d189a1909be8f3858df23bdb23.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceCompliancePolicies.item.assignments.item collection
func (m *DeviceCompliancePolicyRequestBuilder) AssignmentsById(id string)(*i6900b8b299999e708b620425dc8c6ce05ee63496c68d74d9f14fdd267e9722fd.DeviceCompliancePolicyAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyAssignment_id"] = id
    }
    return i6900b8b299999e708b620425dc8c6ce05ee63496c68d74d9f14fdd267e9722fd.NewDeviceCompliancePolicyAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceCompliancePolicyRequestBuilderInternal instantiates a new DeviceCompliancePolicyRequestBuilder and sets the default values.
func NewDeviceCompliancePolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceCompliancePolicyRequestBuilder) {
    m := &DeviceCompliancePolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceCompliancePolicyRequestBuilder instantiates a new DeviceCompliancePolicyRequestBuilder and sets the default values.
func NewDeviceCompliancePolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceCompliancePolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the device compliance policies.
func (m *DeviceCompliancePolicyRequestBuilder) CreateDeleteRequestInformation(options *DeviceCompliancePolicyRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceCompliancePolicyRequestBuilder) CreateGetRequestInformation(options *DeviceCompliancePolicyRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the device compliance policies.
func (m *DeviceCompliancePolicyRequestBuilder) CreatePatchRequestInformation(options *DeviceCompliancePolicyRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the device compliance policies.
func (m *DeviceCompliancePolicyRequestBuilder) Delete(options *DeviceCompliancePolicyRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceCompliancePolicyRequestBuilder) DeviceSettingStateSummaries()(*i19eac586d5d9cdaadb80098949fbb380feb5000e24ac5c8fe81ccde9477b5e2a.DeviceSettingStateSummariesRequestBuilder) {
    return i19eac586d5d9cdaadb80098949fbb380feb5000e24ac5c8fe81ccde9477b5e2a.NewDeviceSettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceSettingStateSummariesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceCompliancePolicies.item.deviceSettingStateSummaries.item collection
func (m *DeviceCompliancePolicyRequestBuilder) DeviceSettingStateSummariesById(id string)(*i524f1aabf095e7a89f8751206668b1131dc6a571119b861cc6a82957868b66bb.SettingStateDeviceSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["settingStateDeviceSummary_id"] = id
    }
    return i524f1aabf095e7a89f8751206668b1131dc6a571119b861cc6a82957868b66bb.NewSettingStateDeviceSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) DeviceStatuses()(*i02fb6a209ffec3d5fa70a95db4705fc7c75cfc91437c2fa31a8e146503996769.DeviceStatusesRequestBuilder) {
    return i02fb6a209ffec3d5fa70a95db4705fc7c75cfc91437c2fa31a8e146503996769.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceCompliancePolicies.item.deviceStatuses.item collection
func (m *DeviceCompliancePolicyRequestBuilder) DeviceStatusesById(id string)(*i322065cfad3bb35a1e313e9c3ea9020cbd58f06bf7ec5f4c92fbf517ea5d4439.DeviceComplianceDeviceStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceDeviceStatus_id"] = id
    }
    return i322065cfad3bb35a1e313e9c3ea9020cbd58f06bf7ec5f4c92fbf517ea5d4439.NewDeviceComplianceDeviceStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) DeviceStatusOverview()(*i2484457c7c0545ba42fb55ebfdec5c3410449670887d9a9294c738eb9760f4a8.DeviceStatusOverviewRequestBuilder) {
    return i2484457c7c0545ba42fb55ebfdec5c3410449670887d9a9294c738eb9760f4a8.NewDeviceStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the device compliance policies.
func (m *DeviceCompliancePolicyRequestBuilder) Get(options *DeviceCompliancePolicyRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceCompliancePolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDeviceCompliancePolicy() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceCompliancePolicy), nil
}
// Patch the device compliance policies.
func (m *DeviceCompliancePolicyRequestBuilder) Patch(options *DeviceCompliancePolicyRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceCompliancePolicyRequestBuilder) ScheduleActionsForRules()(*i7279341ed743f35022dbfaf4fecf2797848a84b5e767ce5329eac709d2cd27d0.ScheduleActionsForRulesRequestBuilder) {
    return i7279341ed743f35022dbfaf4fecf2797848a84b5e767ce5329eac709d2cd27d0.NewScheduleActionsForRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) ScheduledActionsForRule()(*i160c96d8c11dddcb0d327977c26d60c04f4b6c89fe2e2313696b0a10646d4cce.ScheduledActionsForRuleRequestBuilder) {
    return i160c96d8c11dddcb0d327977c26d60c04f4b6c89fe2e2313696b0a10646d4cce.NewScheduledActionsForRuleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScheduledActionsForRuleById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceCompliancePolicies.item.scheduledActionsForRule.item collection
func (m *DeviceCompliancePolicyRequestBuilder) ScheduledActionsForRuleById(id string)(*i7fb30de09cd8a009b11319ff0df1e3a97b0ee6a475979d0c276ba707f9024a6b.DeviceComplianceScheduledActionForRuleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceScheduledActionForRule_id"] = id
    }
    return i7fb30de09cd8a009b11319ff0df1e3a97b0ee6a475979d0c276ba707f9024a6b.NewDeviceComplianceScheduledActionForRuleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) UserStatuses()(*i9bc4838e6f15dbc4846618bfd7af51ca2f4b8feade4b5c921daf6d262ac1f5b6.UserStatusesRequestBuilder) {
    return i9bc4838e6f15dbc4846618bfd7af51ca2f4b8feade4b5c921daf6d262ac1f5b6.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatusesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceCompliancePolicies.item.userStatuses.item collection
func (m *DeviceCompliancePolicyRequestBuilder) UserStatusesById(id string)(*i52eb2204ffb8d3ff80719b2c0c36fc7747e810561f4265399b1b3fe48ce24e2e.DeviceComplianceUserStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceUserStatus_id"] = id
    }
    return i52eb2204ffb8d3ff80719b2c0c36fc7747e810561f4265399b1b3fe48ce24e2e.NewDeviceComplianceUserStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) UserStatusOverview()(*i724a222824d9cdbc36a1de8ffc1291a4febccc391ed1175a2c5bea1b79164a83.UserStatusOverviewRequestBuilder) {
    return i724a222824d9cdbc36a1de8ffc1291a4febccc391ed1175a2c5bea1b79164a83.NewUserStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
