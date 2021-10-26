package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1acaef5ee6ae1997bcbb07d275900c5d95eb14266a9c79c9ee811bf4c0edefb4 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceconfigurations/item/assignments"
    i1cbc049d4f1b0a0c7493469fdf2a72ede1929262173246b80d83717fb3509836 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceconfigurations/item/userstatusoverview"
    i2052556f4fb1ecc2b8520d93accde420e5a85783ea309802622a14a2965cac8c "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceconfigurations/item/devicestatuses"
    i419f93c2382cc7ee59ef6aee59654c413d9feb258f5f656a8f44dc8fb4ad2b61 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceconfigurations/item/devicestatusoverview"
    i4cd13eb5be3949210a98b3e74a48e803b388b1057dc5b76fa296250220177115 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceconfigurations/item/assign"
    i5d3bb0910f46428b82e3ef51c35a9f795351e7ee61bfb08c89e5de44026ecaa2 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceconfigurations/item/getomasettingplaintextvaluewithsecretreferencevalueid"
    ib0c9f5570d692d204085c3256ffa0914246699c56424649fc82be68680943fb2 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceconfigurations/item/devicesettingstatesummaries"
    ie37508b74757cd3bd2b6dc98d9db2ba9651032937fe4385be4b4c611ade0f8fc "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceconfigurations/item/userstatuses"
    i13d5635acea3d1befd38097b223c984a262de351ddcc99002d56e0362ecf8c7c "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceconfigurations/item/userstatuses/item"
    i37d5c64fb898f2ba64f90d5a72b1da2a1813227f63002283be7b64bced3b78de "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceconfigurations/item/devicesettingstatesummaries/item"
    i847fd9044721470530d680766ffdf92095b6c54818a689f742e81f854e71cbd3 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceconfigurations/item/assignments/item"
    ie1759fa3aebfd08bd0a0973770111e293804186cb2bccc14cfdf3f11a34c8f75 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceconfigurations/item/devicestatuses/item"
)

// Builds and executes requests for operations under \deviceManagement\deviceConfigurations\{deviceConfiguration-id}
type DeviceConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The device configurations.
type DeviceConfigurationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *DeviceConfigurationRequestBuilder) Assign()(*i4cd13eb5be3949210a98b3e74a48e803b388b1057dc5b76fa296250220177115.AssignRequestBuilder) {
    return i4cd13eb5be3949210a98b3e74a48e803b388b1057dc5b76fa296250220177115.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceConfigurationRequestBuilder) Assignments()(*i1acaef5ee6ae1997bcbb07d275900c5d95eb14266a9c79c9ee811bf4c0edefb4.AssignmentsRequestBuilder) {
    return i1acaef5ee6ae1997bcbb07d275900c5d95eb14266a9c79c9ee811bf4c0edefb4.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceConfigurations.item.assignments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceConfigurationRequestBuilder) AssignmentsById(id string)(*i847fd9044721470530d680766ffdf92095b6c54818a689f742e81f854e71cbd3.DeviceConfigurationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationAssignment_id"] = id
    }
    return i847fd9044721470530d680766ffdf92095b6c54818a689f742e81f854e71cbd3.NewDeviceConfigurationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new DeviceConfigurationRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDeviceConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceConfigurationRequestBuilder) {
    m := &DeviceConfigurationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/deviceManagement/deviceConfigurations/{deviceConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new DeviceConfigurationRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDeviceConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// The device configurations.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *DeviceConfigurationRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The device configurations.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *DeviceConfigurationRequestBuilder) CreateGetRequestInformation(q func (value *DeviceConfigurationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(DeviceConfigurationRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The device configurations.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *DeviceConfigurationRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceConfiguration, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The device configurations.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *DeviceConfigurationRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceConfigurationRequestBuilder) DeviceSettingStateSummaries()(*ib0c9f5570d692d204085c3256ffa0914246699c56424649fc82be68680943fb2.DeviceSettingStateSummariesRequestBuilder) {
    return ib0c9f5570d692d204085c3256ffa0914246699c56424649fc82be68680943fb2.NewDeviceSettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceConfigurations.item.deviceSettingStateSummaries.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceConfigurationRequestBuilder) DeviceSettingStateSummariesById(id string)(*i37d5c64fb898f2ba64f90d5a72b1da2a1813227f63002283be7b64bced3b78de.SettingStateDeviceSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["settingStateDeviceSummary_id"] = id
    }
    return i37d5c64fb898f2ba64f90d5a72b1da2a1813227f63002283be7b64bced3b78de.NewSettingStateDeviceSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceConfigurationRequestBuilder) DeviceStatuses()(*i2052556f4fb1ecc2b8520d93accde420e5a85783ea309802622a14a2965cac8c.DeviceStatusesRequestBuilder) {
    return i2052556f4fb1ecc2b8520d93accde420e5a85783ea309802622a14a2965cac8c.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceConfigurations.item.deviceStatuses.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceConfigurationRequestBuilder) DeviceStatusesById(id string)(*ie1759fa3aebfd08bd0a0973770111e293804186cb2bccc14cfdf3f11a34c8f75.DeviceConfigurationDeviceStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationDeviceStatus_id"] = id
    }
    return ie1759fa3aebfd08bd0a0973770111e293804186cb2bccc14cfdf3f11a34c8f75.NewDeviceConfigurationDeviceStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceConfigurationRequestBuilder) DeviceStatusOverview()(*i419f93c2382cc7ee59ef6aee59654c413d9feb258f5f656a8f44dc8fb4ad2b61.DeviceStatusOverviewRequestBuilder) {
    return i419f93c2382cc7ee59ef6aee59654c413d9feb258f5f656a8f44dc8fb4ad2b61.NewDeviceStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The device configurations.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *DeviceConfigurationRequestBuilder) Get(q func (value *DeviceConfigurationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDeviceConfiguration() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceConfiguration), nil
}
// Builds and executes requests for operations under \deviceManagement\deviceConfigurations\{deviceConfiguration-id}\microsoft.graph.getOmaSettingPlainTextValue(secretReferenceValueId='{secretReferenceValueId}')
// Parameters:
//  - secretReferenceValueId : Usage: secretReferenceValueId={secretReferenceValueId}
func (m *DeviceConfigurationRequestBuilder) GetOmaSettingPlainTextValueWithSecretReferenceValueId(secretReferenceValueId *string)(*i5d3bb0910f46428b82e3ef51c35a9f795351e7ee61bfb08c89e5de44026ecaa2.GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    return i5d3bb0910f46428b82e3ef51c35a9f795351e7ee61bfb08c89e5de44026ecaa2.NewGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, secretReferenceValueId);
}
// The device configurations.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *DeviceConfigurationRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceConfiguration, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceConfigurationRequestBuilder) UserStatuses()(*ie37508b74757cd3bd2b6dc98d9db2ba9651032937fe4385be4b4c611ade0f8fc.UserStatusesRequestBuilder) {
    return ie37508b74757cd3bd2b6dc98d9db2ba9651032937fe4385be4b4c611ade0f8fc.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceConfigurations.item.userStatuses.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceConfigurationRequestBuilder) UserStatusesById(id string)(*i13d5635acea3d1befd38097b223c984a262de351ddcc99002d56e0362ecf8c7c.DeviceConfigurationUserStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationUserStatus_id"] = id
    }
    return i13d5635acea3d1befd38097b223c984a262de351ddcc99002d56e0362ecf8c7c.NewDeviceConfigurationUserStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceConfigurationRequestBuilder) UserStatusOverview()(*i1cbc049d4f1b0a0c7493469fdf2a72ede1929262173246b80d83717fb3509836.UserStatusOverviewRequestBuilder) {
    return i1cbc049d4f1b0a0c7493469fdf2a72ede1929262173246b80d83717fb3509836.NewUserStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
