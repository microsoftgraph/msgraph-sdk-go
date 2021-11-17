package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i20a49af14049cce13bbcbf77639ba345f23a979ba735144f96caf7cc46ac52d4 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatuses"
    i6bbd09e352a0e87235e9c9775845f897603974356216a5236acfbea5f16155b0 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/assignments"
    i8d9da7a408610e6b2da98617dc04ec0862fb37929855ca63ba17948a3e7c6869 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatussummary"
    ia2a6efbe6de075696db4a436ca6bc9745df0eb596dd5de9281c86bf2bebcd940 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/assign"
    id36e539f78070cadf5b4f1fcefe8eeff25266761b7e3541b45788010e5cdfb9b "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatussummary"
    iebd1b47b377fdbd0e32c261a4ebc852aebb3266e58bde42f93309b4e2edc6d16 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatuses"
    i5644f6cefb6e37ba63ec8582d3a158ec48759832b4146d5ffd2d350288b5e5ca "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatuses/item"
    i72a4318a7e8f28adce62f7520318129090631e70fa23a9371561c34fac1f0582 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatuses/item"
    i88ff03a861b0c98567d5cfa6d93d8d9d0de34bb73fdd948e55765d9858a275c3 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/assignments/item"
)

// Builds and executes requests for operations under \deviceAppManagement\mobileAppConfigurations\{managedDeviceMobileAppConfiguration-id}
type ManagedDeviceMobileAppConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ManagedDeviceMobileAppConfigurationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ManagedDeviceMobileAppConfigurationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedDeviceMobileAppConfigurationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The Managed Device Mobile Application Configurations.
type ManagedDeviceMobileAppConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ManagedDeviceMobileAppConfigurationRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDeviceMobileAppConfiguration;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Assign()(*ia2a6efbe6de075696db4a436ca6bc9745df0eb596dd5de9281c86bf2bebcd940.AssignRequestBuilder) {
    return ia2a6efbe6de075696db4a436ca6bc9745df0eb596dd5de9281c86bf2bebcd940.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Assignments()(*i6bbd09e352a0e87235e9c9775845f897603974356216a5236acfbea5f16155b0.AssignmentsRequestBuilder) {
    return i6bbd09e352a0e87235e9c9775845f897603974356216a5236acfbea5f16155b0.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.mobileAppConfigurations.item.assignments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) AssignmentsById(id string)(*i88ff03a861b0c98567d5cfa6d93d8d9d0de34bb73fdd948e55765d9858a275c3.ManagedDeviceMobileAppConfigurationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationAssignment_id"] = id
    }
    return i88ff03a861b0c98567d5cfa6d93d8d9d0de34bb73fdd948e55765d9858a275c3.NewManagedDeviceMobileAppConfigurationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new ManagedDeviceMobileAppConfigurationRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewManagedDeviceMobileAppConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceMobileAppConfigurationRequestBuilder) {
    m := &ManagedDeviceMobileAppConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileAppConfigurations/{managedDeviceMobileAppConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ManagedDeviceMobileAppConfigurationRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewManagedDeviceMobileAppConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceMobileAppConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceMobileAppConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// The Managed Device Mobile Application Configurations.
// Parameters:
//  - options : Options for the request
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) CreateDeleteRequestInformation(options *ManagedDeviceMobileAppConfigurationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The Managed Device Mobile Application Configurations.
// Parameters:
//  - options : Options for the request
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) CreateGetRequestInformation(options *ManagedDeviceMobileAppConfigurationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The Managed Device Mobile Application Configurations.
// Parameters:
//  - options : Options for the request
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) CreatePatchRequestInformation(options *ManagedDeviceMobileAppConfigurationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The Managed Device Mobile Application Configurations.
// Parameters:
//  - options : Options for the request
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Delete(options *ManagedDeviceMobileAppConfigurationRequestBuilderDeleteOptions)(error) {
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
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) DeviceStatuses()(*iebd1b47b377fdbd0e32c261a4ebc852aebb3266e58bde42f93309b4e2edc6d16.DeviceStatusesRequestBuilder) {
    return iebd1b47b377fdbd0e32c261a4ebc852aebb3266e58bde42f93309b4e2edc6d16.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.mobileAppConfigurations.item.deviceStatuses.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) DeviceStatusesById(id string)(*i5644f6cefb6e37ba63ec8582d3a158ec48759832b4146d5ffd2d350288b5e5ca.ManagedDeviceMobileAppConfigurationDeviceStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationDeviceStatus_id"] = id
    }
    return i5644f6cefb6e37ba63ec8582d3a158ec48759832b4146d5ffd2d350288b5e5ca.NewManagedDeviceMobileAppConfigurationDeviceStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) DeviceStatusSummary()(*i8d9da7a408610e6b2da98617dc04ec0862fb37929855ca63ba17948a3e7c6869.DeviceStatusSummaryRequestBuilder) {
    return i8d9da7a408610e6b2da98617dc04ec0862fb37929855ca63ba17948a3e7c6869.NewDeviceStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The Managed Device Mobile Application Configurations.
// Parameters:
//  - options : Options for the request
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Get(options *ManagedDeviceMobileAppConfigurationRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDeviceMobileAppConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewManagedDeviceMobileAppConfiguration() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDeviceMobileAppConfiguration), nil
}
// The Managed Device Mobile Application Configurations.
// Parameters:
//  - options : Options for the request
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Patch(options *ManagedDeviceMobileAppConfigurationRequestBuilderPatchOptions)(error) {
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
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) UserStatuses()(*i20a49af14049cce13bbcbf77639ba345f23a979ba735144f96caf7cc46ac52d4.UserStatusesRequestBuilder) {
    return i20a49af14049cce13bbcbf77639ba345f23a979ba735144f96caf7cc46ac52d4.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.mobileAppConfigurations.item.userStatuses.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) UserStatusesById(id string)(*i72a4318a7e8f28adce62f7520318129090631e70fa23a9371561c34fac1f0582.ManagedDeviceMobileAppConfigurationUserStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationUserStatus_id"] = id
    }
    return i72a4318a7e8f28adce62f7520318129090631e70fa23a9371561c34fac1f0582.NewManagedDeviceMobileAppConfigurationUserStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) UserStatusSummary()(*id36e539f78070cadf5b4f1fcefe8eeff25266761b7e3541b45788010e5cdfb9b.UserStatusSummaryRequestBuilder) {
    return id36e539f78070cadf5b4f1fcefe8eeff25266761b7e3541b45788010e5cdfb9b.NewUserStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
