package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i02e61d7c24506eb3092637591a2fe281450dbe31d56751354fa7a05f72a89a6d "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/disablelostmode"
    i15a1a10a0c024f930afef3b930ac6a0a48505b8f50dc765c9208ba2816bfb8fb "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/resetpasscode"
    i35e5de1b6d1217fb83cdfe34b93910dd53c07bcbaed491f2d9dd30bf097a2e27 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/cleanwindowsdevice"
    i4be1f6e23df553f3d25539b06b6e72e653c54673be6c256636259b2dfc46b028 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/deleteuserfromsharedappledevice"
    i5626258f553fa5ab744196016cddf3bdb99eca3ad27f9716b75a1816ddcac4e4 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/windowsdefenderupdatesignatures"
    i579420ac2331f8db64f4efb8f6b2f9b7e59da9e53c2f60c7189463b2f6729a66 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/retire"
    i61111df9932978d316b7c2d5bf743126bd3dfc765577371722fdfd9540e546c6 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/deviceconfigurationstates"
    i73fe8c13857767fd45ca26708823e85bca6083a2a2ebdb0dda0d4fe4e91cdf17 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/updatewindowsdeviceaccount"
    i957d1faa5f562ab89577eb46e79203c7e50075bcfa0818779d21777aac9c4774 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/locatedevice"
    ib2246236511e9e6bd5f149ec83d2768718af4bee5cf6b48497a15ba9859004da "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/devicecompliancepolicystates"
    ib4e430bec489c09a72c51ad3159e98695c04be6d6cbd03769b4074193f098ce8 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/rebootnow"
    ib6fe4058a826f6d827fec2a61b53730458811701046b26ceb4044feeb7ddf7d1 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/recoverpasscode"
    ib975ed5a3045ece0bc94c5dd2f4c5a4ea484dc80b1830219bb365eea8fd973c0 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/logoutsharedappledeviceactiveuser"
    ic1c10c22035a19935689244731094da0d49c68bb4707c06ce0ecf896730268c2 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/windowsdefenderscan"
    ic2df49236f6d10a558da4728bb3e2ce53874fd50cfc571bf9e691b66d446221c "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/syncdevice"
    ic61092c249f07b9e411d18b3b8310e71210bdf5c5de775c13112cab083c5d082 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/wipe"
    icc2068c1350cb53b049a71e87eb1cf48456fa9e9b48d284b8abaa38d4207ffc2 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/shutdown"
    id96b7ae916d8c8e465b8a821c2fe41d4d020ec0a7610953b779987f93a48d88e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/devicecategory"
    ie6232acc82dadd162e8e29367892609c026901eba827c7fdbb14386e5d36b99d "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/remotelock"
    if1e16de23b7ea9301d8d5208c98e249a57e7e1e4df4785a249e053121429dd1e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/requestremoteassistance"
    if245b21847517f2bdf21fefeaac5356812b1d5b2cad61e07f1e97c61a72adfeb "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/bypassactivationlock"
    i257ff0bc433efcedf2c2cd195e3b06a2fbc6131ac66e39a2a9444f64b497cbf5 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/deviceconfigurationstates/item"
    i30c735a31ceee62863c65dfaa2404ac107c81c4dfd9ccc96d5eee45ca24f04f2 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item/devicecompliancepolicystates/item"
)

// Builds and executes requests for operations under \deviceManagement\managedDevices\{managedDevice-id}
type ManagedDeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ManagedDeviceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ManagedDeviceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedDeviceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The list of managed devices.
type ManagedDeviceRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ManagedDeviceRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDevice;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ManagedDeviceRequestBuilder) BypassActivationLock()(*if245b21847517f2bdf21fefeaac5356812b1d5b2cad61e07f1e97c61a72adfeb.BypassActivationLockRequestBuilder) {
    return if245b21847517f2bdf21fefeaac5356812b1d5b2cad61e07f1e97c61a72adfeb.NewBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) CleanWindowsDevice()(*i35e5de1b6d1217fb83cdfe34b93910dd53c07bcbaed491f2d9dd30bf097a2e27.CleanWindowsDeviceRequestBuilder) {
    return i35e5de1b6d1217fb83cdfe34b93910dd53c07bcbaed491f2d9dd30bf097a2e27.NewCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ManagedDeviceRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewManagedDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceRequestBuilder) {
    m := &ManagedDeviceRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/deviceManagement/managedDevices/{managedDevice_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ManagedDeviceRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewManagedDeviceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// The list of managed devices.
// Parameters:
//  - options : Options for the request
func (m *ManagedDeviceRequestBuilder) CreateDeleteRequestInformation(options *ManagedDeviceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The list of managed devices.
// Parameters:
//  - options : Options for the request
func (m *ManagedDeviceRequestBuilder) CreateGetRequestInformation(options *ManagedDeviceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The list of managed devices.
// Parameters:
//  - options : Options for the request
func (m *ManagedDeviceRequestBuilder) CreatePatchRequestInformation(options *ManagedDeviceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The list of managed devices.
// Parameters:
//  - options : Options for the request
func (m *ManagedDeviceRequestBuilder) Delete(options *ManagedDeviceRequestBuilderDeleteOptions)(error) {
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
func (m *ManagedDeviceRequestBuilder) DeleteUserFromSharedAppleDevice()(*i4be1f6e23df553f3d25539b06b6e72e653c54673be6c256636259b2dfc46b028.DeleteUserFromSharedAppleDeviceRequestBuilder) {
    return i4be1f6e23df553f3d25539b06b6e72e653c54673be6c256636259b2dfc46b028.NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DeviceCategory()(*id96b7ae916d8c8e465b8a821c2fe41d4d020ec0a7610953b779987f93a48d88e.DeviceCategoryRequestBuilder) {
    return id96b7ae916d8c8e465b8a821c2fe41d4d020ec0a7610953b779987f93a48d88e.NewDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DeviceCompliancePolicyStates()(*ib2246236511e9e6bd5f149ec83d2768718af4bee5cf6b48497a15ba9859004da.DeviceCompliancePolicyStatesRequestBuilder) {
    return ib2246236511e9e6bd5f149ec83d2768718af4bee5cf6b48497a15ba9859004da.NewDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.deviceManagement.managedDevices.item.deviceCompliancePolicyStates.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ManagedDeviceRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*i30c735a31ceee62863c65dfaa2404ac107c81c4dfd9ccc96d5eee45ca24f04f2.DeviceCompliancePolicyStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState_id"] = id
    }
    return i30c735a31ceee62863c65dfaa2404ac107c81c4dfd9ccc96d5eee45ca24f04f2.NewDeviceCompliancePolicyStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DeviceConfigurationStates()(*i61111df9932978d316b7c2d5bf743126bd3dfc765577371722fdfd9540e546c6.DeviceConfigurationStatesRequestBuilder) {
    return i61111df9932978d316b7c2d5bf743126bd3dfc765577371722fdfd9540e546c6.NewDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.deviceManagement.managedDevices.item.deviceConfigurationStates.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ManagedDeviceRequestBuilder) DeviceConfigurationStatesById(id string)(*i257ff0bc433efcedf2c2cd195e3b06a2fbc6131ac66e39a2a9444f64b497cbf5.DeviceConfigurationStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState_id"] = id
    }
    return i257ff0bc433efcedf2c2cd195e3b06a2fbc6131ac66e39a2a9444f64b497cbf5.NewDeviceConfigurationStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DisableLostMode()(*i02e61d7c24506eb3092637591a2fe281450dbe31d56751354fa7a05f72a89a6d.DisableLostModeRequestBuilder) {
    return i02e61d7c24506eb3092637591a2fe281450dbe31d56751354fa7a05f72a89a6d.NewDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The list of managed devices.
// Parameters:
//  - options : Options for the request
func (m *ManagedDeviceRequestBuilder) Get(options *ManagedDeviceRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDevice, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewManagedDevice() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDevice), nil
}
func (m *ManagedDeviceRequestBuilder) LocateDevice()(*i957d1faa5f562ab89577eb46e79203c7e50075bcfa0818779d21777aac9c4774.LocateDeviceRequestBuilder) {
    return i957d1faa5f562ab89577eb46e79203c7e50075bcfa0818779d21777aac9c4774.NewLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*ib975ed5a3045ece0bc94c5dd2f4c5a4ea484dc80b1830219bb365eea8fd973c0.LogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return ib975ed5a3045ece0bc94c5dd2f4c5a4ea484dc80b1830219bb365eea8fd973c0.NewLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The list of managed devices.
// Parameters:
//  - options : Options for the request
func (m *ManagedDeviceRequestBuilder) Patch(options *ManagedDeviceRequestBuilderPatchOptions)(error) {
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
func (m *ManagedDeviceRequestBuilder) RebootNow()(*ib4e430bec489c09a72c51ad3159e98695c04be6d6cbd03769b4074193f098ce8.RebootNowRequestBuilder) {
    return ib4e430bec489c09a72c51ad3159e98695c04be6d6cbd03769b4074193f098ce8.NewRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RecoverPasscode()(*ib6fe4058a826f6d827fec2a61b53730458811701046b26ceb4044feeb7ddf7d1.RecoverPasscodeRequestBuilder) {
    return ib6fe4058a826f6d827fec2a61b53730458811701046b26ceb4044feeb7ddf7d1.NewRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RemoteLock()(*ie6232acc82dadd162e8e29367892609c026901eba827c7fdbb14386e5d36b99d.RemoteLockRequestBuilder) {
    return ie6232acc82dadd162e8e29367892609c026901eba827c7fdbb14386e5d36b99d.NewRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RequestRemoteAssistance()(*if1e16de23b7ea9301d8d5208c98e249a57e7e1e4df4785a249e053121429dd1e.RequestRemoteAssistanceRequestBuilder) {
    return if1e16de23b7ea9301d8d5208c98e249a57e7e1e4df4785a249e053121429dd1e.NewRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResetPasscode()(*i15a1a10a0c024f930afef3b930ac6a0a48505b8f50dc765c9208ba2816bfb8fb.ResetPasscodeRequestBuilder) {
    return i15a1a10a0c024f930afef3b930ac6a0a48505b8f50dc765c9208ba2816bfb8fb.NewResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Retire()(*i579420ac2331f8db64f4efb8f6b2f9b7e59da9e53c2f60c7189463b2f6729a66.RetireRequestBuilder) {
    return i579420ac2331f8db64f4efb8f6b2f9b7e59da9e53c2f60c7189463b2f6729a66.NewRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ShutDown()(*icc2068c1350cb53b049a71e87eb1cf48456fa9e9b48d284b8abaa38d4207ffc2.ShutDownRequestBuilder) {
    return icc2068c1350cb53b049a71e87eb1cf48456fa9e9b48d284b8abaa38d4207ffc2.NewShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SyncDevice()(*ic2df49236f6d10a558da4728bb3e2ce53874fd50cfc571bf9e691b66d446221c.SyncDeviceRequestBuilder) {
    return ic2df49236f6d10a558da4728bb3e2ce53874fd50cfc571bf9e691b66d446221c.NewSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) UpdateWindowsDeviceAccount()(*i73fe8c13857767fd45ca26708823e85bca6083a2a2ebdb0dda0d4fe4e91cdf17.UpdateWindowsDeviceAccountRequestBuilder) {
    return i73fe8c13857767fd45ca26708823e85bca6083a2a2ebdb0dda0d4fe4e91cdf17.NewUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderScan()(*ic1c10c22035a19935689244731094da0d49c68bb4707c06ce0ecf896730268c2.WindowsDefenderScanRequestBuilder) {
    return ic1c10c22035a19935689244731094da0d49c68bb4707c06ce0ecf896730268c2.NewWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderUpdateSignatures()(*i5626258f553fa5ab744196016cddf3bdb99eca3ad27f9716b75a1816ddcac4e4.WindowsDefenderUpdateSignaturesRequestBuilder) {
    return i5626258f553fa5ab744196016cddf3bdb99eca3ad27f9716b75a1816ddcac4e4.NewWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Wipe()(*ic61092c249f07b9e411d18b3b8310e71210bdf5c5de775c13112cab083c5d082.WipeRequestBuilder) {
    return ic61092c249f07b9e411d18b3b8310e71210bdf5c5de775c13112cab083c5d082.NewWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
