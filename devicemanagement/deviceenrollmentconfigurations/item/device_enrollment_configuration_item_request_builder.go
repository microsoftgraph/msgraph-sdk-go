package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    id9a7398da3afd65325bb4b76c51dfd745cf1949c7305f7305c55987303e7c110 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceenrollmentconfigurations/item/assign"
    ie89e35b2b9b66f7bb09b458829876684f00494c74c7722c84d86c9ce3b2ebf71 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceenrollmentconfigurations/item/setpriority"
    if663dc8479f3b9f5e6a2081e3f608c0970726f4a06ad253da9b3de49b9527f3f "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceenrollmentconfigurations/item/assignments"
    i5bae63f9ea126428555bb4338d2c597964628babb784c2a5f30e3ffc29c7e6bb "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceenrollmentconfigurations/item/assignments/item"
)

// DeviceEnrollmentConfigurationItemRequestBuilder provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.deviceManagement entity.
type DeviceEnrollmentConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceEnrollmentConfigurationItemRequestBuilderDeleteOptions options for Delete
type DeviceEnrollmentConfigurationItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceEnrollmentConfigurationItemRequestBuilderGetOptions options for Get
type DeviceEnrollmentConfigurationItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceEnrollmentConfigurationItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceEnrollmentConfigurationItemRequestBuilderGetQueryParameters the list of device enrollment configurations
type DeviceEnrollmentConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceEnrollmentConfigurationItemRequestBuilderPatchOptions options for Patch
type DeviceEnrollmentConfigurationItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceEnrollmentConfigurationable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) Assign()(*id9a7398da3afd65325bb4b76c51dfd745cf1949c7305f7305c55987303e7c110.AssignRequestBuilder) {
    return id9a7398da3afd65325bb4b76c51dfd745cf1949c7305f7305c55987303e7c110.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) Assignments()(*if663dc8479f3b9f5e6a2081e3f608c0970726f4a06ad253da9b3de49b9527f3f.AssignmentsRequestBuilder) {
    return if663dc8479f3b9f5e6a2081e3f608c0970726f4a06ad253da9b3de49b9527f3f.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceEnrollmentConfigurations.item.assignments.item collection
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) AssignmentsById(id string)(*i5bae63f9ea126428555bb4338d2c597964628babb784c2a5f30e3ffc29c7e6bb.EnrollmentConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["enrollmentConfigurationAssignment_id"] = id
    }
    return i5bae63f9ea126428555bb4338d2c597964628babb784c2a5f30e3ffc29c7e6bb.NewEnrollmentConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceEnrollmentConfigurationItemRequestBuilderInternal instantiates a new DeviceEnrollmentConfigurationItemRequestBuilder and sets the default values.
func NewDeviceEnrollmentConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceEnrollmentConfigurationItemRequestBuilder) {
    m := &DeviceEnrollmentConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceEnrollmentConfigurationItemRequestBuilder instantiates a new DeviceEnrollmentConfigurationItemRequestBuilder and sets the default values.
func NewDeviceEnrollmentConfigurationItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceEnrollmentConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceEnrollmentConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceEnrollmentConfigurations for deviceManagement
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceEnrollmentConfigurationItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of device enrollment configurations
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) CreateGetRequestInformation(options *DeviceEnrollmentConfigurationItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceEnrollmentConfigurations in deviceManagement
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) CreatePatchRequestInformation(options *DeviceEnrollmentConfigurationItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property deviceEnrollmentConfigurations for deviceManagement
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) Delete(options *DeviceEnrollmentConfigurationItemRequestBuilderDeleteOptions)(error) {
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
// Get the list of device enrollment configurations
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) Get(options *DeviceEnrollmentConfigurationItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceEnrollmentConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDeviceEnrollmentConfigurationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceEnrollmentConfigurationable), nil
}
// Patch update the navigation property deviceEnrollmentConfigurations in deviceManagement
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) Patch(options *DeviceEnrollmentConfigurationItemRequestBuilderPatchOptions)(error) {
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
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) SetPriority()(*ie89e35b2b9b66f7bb09b458829876684f00494c74c7722c84d86c9ce3b2ebf71.SetPriorityRequestBuilder) {
    return ie89e35b2b9b66f7bb09b458829876684f00494c74c7722c84d86c9ce3b2ebf71.NewSetPriorityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
