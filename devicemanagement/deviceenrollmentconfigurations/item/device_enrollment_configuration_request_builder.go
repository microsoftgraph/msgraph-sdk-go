package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    id9a7398da3afd65325bb4b76c51dfd745cf1949c7305f7305c55987303e7c110 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceenrollmentconfigurations/item/assign"
    ie89e35b2b9b66f7bb09b458829876684f00494c74c7722c84d86c9ce3b2ebf71 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceenrollmentconfigurations/item/setpriority"
    if663dc8479f3b9f5e6a2081e3f608c0970726f4a06ad253da9b3de49b9527f3f "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceenrollmentconfigurations/item/assignments"
    i5bae63f9ea126428555bb4338d2c597964628babb784c2a5f30e3ffc29c7e6bb "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceenrollmentconfigurations/item/assignments/item"
)

// Builds and executes requests for operations under \deviceManagement\deviceEnrollmentConfigurations\{deviceEnrollmentConfiguration-id}
type DeviceEnrollmentConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type DeviceEnrollmentConfigurationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type DeviceEnrollmentConfigurationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceEnrollmentConfigurationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The list of device enrollment configurations
type DeviceEnrollmentConfigurationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type DeviceEnrollmentConfigurationRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceEnrollmentConfiguration;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceEnrollmentConfigurationRequestBuilder) Assign()(*id9a7398da3afd65325bb4b76c51dfd745cf1949c7305f7305c55987303e7c110.AssignRequestBuilder) {
    return id9a7398da3afd65325bb4b76c51dfd745cf1949c7305f7305c55987303e7c110.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceEnrollmentConfigurationRequestBuilder) Assignments()(*if663dc8479f3b9f5e6a2081e3f608c0970726f4a06ad253da9b3de49b9527f3f.AssignmentsRequestBuilder) {
    return if663dc8479f3b9f5e6a2081e3f608c0970726f4a06ad253da9b3de49b9527f3f.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceEnrollmentConfigurations.item.assignments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceEnrollmentConfigurationRequestBuilder) AssignmentsById(id string)(*i5bae63f9ea126428555bb4338d2c597964628babb784c2a5f30e3ffc29c7e6bb.EnrollmentConfigurationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["enrollmentConfigurationAssignment_id"] = id
    }
    return i5bae63f9ea126428555bb4338d2c597964628babb784c2a5f30e3ffc29c7e6bb.NewEnrollmentConfigurationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new DeviceEnrollmentConfigurationRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDeviceEnrollmentConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceEnrollmentConfigurationRequestBuilder) {
    m := &DeviceEnrollmentConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new DeviceEnrollmentConfigurationRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDeviceEnrollmentConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceEnrollmentConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceEnrollmentConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// The list of device enrollment configurations
// Parameters:
//  - options : Options for the request
func (m *DeviceEnrollmentConfigurationRequestBuilder) CreateDeleteRequestInformation(options *DeviceEnrollmentConfigurationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The list of device enrollment configurations
// Parameters:
//  - options : Options for the request
func (m *DeviceEnrollmentConfigurationRequestBuilder) CreateGetRequestInformation(options *DeviceEnrollmentConfigurationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// The list of device enrollment configurations
// Parameters:
//  - options : Options for the request
func (m *DeviceEnrollmentConfigurationRequestBuilder) CreatePatchRequestInformation(options *DeviceEnrollmentConfigurationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The list of device enrollment configurations
// Parameters:
//  - options : Options for the request
func (m *DeviceEnrollmentConfigurationRequestBuilder) Delete(options *DeviceEnrollmentConfigurationRequestBuilderDeleteOptions)(error) {
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
// The list of device enrollment configurations
// Parameters:
//  - options : Options for the request
func (m *DeviceEnrollmentConfigurationRequestBuilder) Get(options *DeviceEnrollmentConfigurationRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceEnrollmentConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDeviceEnrollmentConfiguration() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceEnrollmentConfiguration), nil
}
// The list of device enrollment configurations
// Parameters:
//  - options : Options for the request
func (m *DeviceEnrollmentConfigurationRequestBuilder) Patch(options *DeviceEnrollmentConfigurationRequestBuilderPatchOptions)(error) {
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
func (m *DeviceEnrollmentConfigurationRequestBuilder) SetPriority()(*ie89e35b2b9b66f7bb09b458829876684f00494c74c7722c84d86c9ce3b2ebf71.SetPriorityRequestBuilder) {
    return ie89e35b2b9b66f7bb09b458829876684f00494c74c7722c84d86c9ce3b2ebf71.NewSetPriorityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
