package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i26fe70e594b43427646ebf5c1a30088de63363845fdd8d965f6cb9ba0749f80e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicysettingstatesummaries/item/devicecompliancesettingstates"
    i526f8febb10205573c5dd2e17b0c9f845b7b7aa8f86cebe17e6e7b83e19bbf56 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicysettingstatesummaries/item/devicecompliancesettingstates/item"
)

// DeviceCompliancePolicySettingStateSummaryItemRequestBuilder builds and executes requests for operations under \deviceManagement\deviceCompliancePolicySettingStateSummaries\{deviceCompliancePolicySettingStateSummary-id}
type DeviceCompliancePolicySettingStateSummaryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceCompliancePolicySettingStateSummaryItemRequestBuilderDeleteOptions options for Delete
type DeviceCompliancePolicySettingStateSummaryItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceCompliancePolicySettingStateSummaryItemRequestBuilderGetOptions options for Get
type DeviceCompliancePolicySettingStateSummaryItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceCompliancePolicySettingStateSummaryItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceCompliancePolicySettingStateSummaryItemRequestBuilderGetQueryParameters the summary states of compliance policy settings for this account.
type DeviceCompliancePolicySettingStateSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceCompliancePolicySettingStateSummaryItemRequestBuilderPatchOptions options for Patch
type DeviceCompliancePolicySettingStateSummaryItemRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceCompliancePolicySettingStateSummary;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDeviceCompliancePolicySettingStateSummaryItemRequestBuilderInternal instantiates a new DeviceCompliancePolicySettingStateSummaryItemRequestBuilder and sets the default values.
func NewDeviceCompliancePolicySettingStateSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceCompliancePolicySettingStateSummaryItemRequestBuilder) {
    m := &DeviceCompliancePolicySettingStateSummaryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceCompliancePolicySettingStateSummaryItemRequestBuilder instantiates a new DeviceCompliancePolicySettingStateSummaryItemRequestBuilder and sets the default values.
func NewDeviceCompliancePolicySettingStateSummaryItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceCompliancePolicySettingStateSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePolicySettingStateSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the summary states of compliance policy settings for this account.
func (m *DeviceCompliancePolicySettingStateSummaryItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceCompliancePolicySettingStateSummaryItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the summary states of compliance policy settings for this account.
func (m *DeviceCompliancePolicySettingStateSummaryItemRequestBuilder) CreateGetRequestInformation(options *DeviceCompliancePolicySettingStateSummaryItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the summary states of compliance policy settings for this account.
func (m *DeviceCompliancePolicySettingStateSummaryItemRequestBuilder) CreatePatchRequestInformation(options *DeviceCompliancePolicySettingStateSummaryItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the summary states of compliance policy settings for this account.
func (m *DeviceCompliancePolicySettingStateSummaryItemRequestBuilder) Delete(options *DeviceCompliancePolicySettingStateSummaryItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceCompliancePolicySettingStateSummaryItemRequestBuilder) DeviceComplianceSettingStates()(*i26fe70e594b43427646ebf5c1a30088de63363845fdd8d965f6cb9ba0749f80e.DeviceComplianceSettingStatesRequestBuilder) {
    return i26fe70e594b43427646ebf5c1a30088de63363845fdd8d965f6cb9ba0749f80e.NewDeviceComplianceSettingStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceComplianceSettingStatesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceCompliancePolicySettingStateSummaries.item.deviceComplianceSettingStates.item collection
func (m *DeviceCompliancePolicySettingStateSummaryItemRequestBuilder) DeviceComplianceSettingStatesById(id string)(*i526f8febb10205573c5dd2e17b0c9f845b7b7aa8f86cebe17e6e7b83e19bbf56.DeviceComplianceSettingStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceSettingState_id"] = id
    }
    return i526f8febb10205573c5dd2e17b0c9f845b7b7aa8f86cebe17e6e7b83e19bbf56.NewDeviceComplianceSettingStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the summary states of compliance policy settings for this account.
func (m *DeviceCompliancePolicySettingStateSummaryItemRequestBuilder) Get(options *DeviceCompliancePolicySettingStateSummaryItemRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceCompliancePolicySettingStateSummary, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDeviceCompliancePolicySettingStateSummary() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceCompliancePolicySettingStateSummary), nil
}
// Patch the summary states of compliance policy settings for this account.
func (m *DeviceCompliancePolicySettingStateSummaryItemRequestBuilder) Patch(options *DeviceCompliancePolicySettingStateSummaryItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
