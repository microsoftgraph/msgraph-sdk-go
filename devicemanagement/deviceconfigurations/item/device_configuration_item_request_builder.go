package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
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

// DeviceConfigurationItemRequestBuilder provides operations to manage the deviceConfigurations property of the microsoft.graph.deviceManagement entity.
type DeviceConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceConfigurationItemRequestBuilderDeleteOptions options for Delete
type DeviceConfigurationItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceConfigurationItemRequestBuilderGetOptions options for Get
type DeviceConfigurationItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceConfigurationItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceConfigurationItemRequestBuilderGetQueryParameters the device configurations.
type DeviceConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceConfigurationItemRequestBuilderPatchOptions options for Patch
type DeviceConfigurationItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceConfigurationable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceConfigurationItemRequestBuilder) Assign()(*i4cd13eb5be3949210a98b3e74a48e803b388b1057dc5b76fa296250220177115.AssignRequestBuilder) {
    return i4cd13eb5be3949210a98b3e74a48e803b388b1057dc5b76fa296250220177115.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceConfigurationItemRequestBuilder) Assignments()(*i1acaef5ee6ae1997bcbb07d275900c5d95eb14266a9c79c9ee811bf4c0edefb4.AssignmentsRequestBuilder) {
    return i1acaef5ee6ae1997bcbb07d275900c5d95eb14266a9c79c9ee811bf4c0edefb4.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceConfigurations.item.assignments.item collection
func (m *DeviceConfigurationItemRequestBuilder) AssignmentsById(id string)(*i847fd9044721470530d680766ffdf92095b6c54818a689f742e81f854e71cbd3.DeviceConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationAssignment_id"] = id
    }
    return i847fd9044721470530d680766ffdf92095b6c54818a689f742e81f854e71cbd3.NewDeviceConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceConfigurationItemRequestBuilderInternal instantiates a new DeviceConfigurationItemRequestBuilder and sets the default values.
func NewDeviceConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceConfigurationItemRequestBuilder) {
    m := &DeviceConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceConfigurationItemRequestBuilder instantiates a new DeviceConfigurationItemRequestBuilder and sets the default values.
func NewDeviceConfigurationItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceConfigurations for deviceManagement
func (m *DeviceConfigurationItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceConfigurationItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the device configurations.
func (m *DeviceConfigurationItemRequestBuilder) CreateGetRequestInformation(options *DeviceConfigurationItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceConfigurations in deviceManagement
func (m *DeviceConfigurationItemRequestBuilder) CreatePatchRequestInformation(options *DeviceConfigurationItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property deviceConfigurations for deviceManagement
func (m *DeviceConfigurationItemRequestBuilder) Delete(options *DeviceConfigurationItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceConfigurationItemRequestBuilder) DeviceSettingStateSummaries()(*ib0c9f5570d692d204085c3256ffa0914246699c56424649fc82be68680943fb2.DeviceSettingStateSummariesRequestBuilder) {
    return ib0c9f5570d692d204085c3256ffa0914246699c56424649fc82be68680943fb2.NewDeviceSettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceSettingStateSummariesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceConfigurations.item.deviceSettingStateSummaries.item collection
func (m *DeviceConfigurationItemRequestBuilder) DeviceSettingStateSummariesById(id string)(*i37d5c64fb898f2ba64f90d5a72b1da2a1813227f63002283be7b64bced3b78de.SettingStateDeviceSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["settingStateDeviceSummary_id"] = id
    }
    return i37d5c64fb898f2ba64f90d5a72b1da2a1813227f63002283be7b64bced3b78de.NewSettingStateDeviceSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceConfigurationItemRequestBuilder) DeviceStatuses()(*i2052556f4fb1ecc2b8520d93accde420e5a85783ea309802622a14a2965cac8c.DeviceStatusesRequestBuilder) {
    return i2052556f4fb1ecc2b8520d93accde420e5a85783ea309802622a14a2965cac8c.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceConfigurations.item.deviceStatuses.item collection
func (m *DeviceConfigurationItemRequestBuilder) DeviceStatusesById(id string)(*ie1759fa3aebfd08bd0a0973770111e293804186cb2bccc14cfdf3f11a34c8f75.DeviceConfigurationDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationDeviceStatus_id"] = id
    }
    return ie1759fa3aebfd08bd0a0973770111e293804186cb2bccc14cfdf3f11a34c8f75.NewDeviceConfigurationDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceConfigurationItemRequestBuilder) DeviceStatusOverview()(*i419f93c2382cc7ee59ef6aee59654c413d9feb258f5f656a8f44dc8fb4ad2b61.DeviceStatusOverviewRequestBuilder) {
    return i419f93c2382cc7ee59ef6aee59654c413d9feb258f5f656a8f44dc8fb4ad2b61.NewDeviceStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the device configurations.
func (m *DeviceConfigurationItemRequestBuilder) Get(options *DeviceConfigurationItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDeviceConfigurationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceConfigurationable), nil
}
// GetOmaSettingPlainTextValueWithSecretReferenceValueId provides operations to call the getOmaSettingPlainTextValue method.
func (m *DeviceConfigurationItemRequestBuilder) GetOmaSettingPlainTextValueWithSecretReferenceValueId(secretReferenceValueId *string)(*i5d3bb0910f46428b82e3ef51c35a9f795351e7ee61bfb08c89e5de44026ecaa2.GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    return i5d3bb0910f46428b82e3ef51c35a9f795351e7ee61bfb08c89e5de44026ecaa2.NewGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, secretReferenceValueId);
}
// Patch update the navigation property deviceConfigurations in deviceManagement
func (m *DeviceConfigurationItemRequestBuilder) Patch(options *DeviceConfigurationItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceConfigurationItemRequestBuilder) UserStatuses()(*ie37508b74757cd3bd2b6dc98d9db2ba9651032937fe4385be4b4c611ade0f8fc.UserStatusesRequestBuilder) {
    return ie37508b74757cd3bd2b6dc98d9db2ba9651032937fe4385be4b4c611ade0f8fc.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatusesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceConfigurations.item.userStatuses.item collection
func (m *DeviceConfigurationItemRequestBuilder) UserStatusesById(id string)(*i13d5635acea3d1befd38097b223c984a262de351ddcc99002d56e0362ecf8c7c.DeviceConfigurationUserStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationUserStatus_id"] = id
    }
    return i13d5635acea3d1befd38097b223c984a262de351ddcc99002d56e0362ecf8c7c.NewDeviceConfigurationUserStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceConfigurationItemRequestBuilder) UserStatusOverview()(*i1cbc049d4f1b0a0c7493469fdf2a72ede1929262173246b80d83717fb3509836.UserStatusOverviewRequestBuilder) {
    return i1cbc049d4f1b0a0c7493469fdf2a72ede1929262173246b80d83717fb3509836.NewUserStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
