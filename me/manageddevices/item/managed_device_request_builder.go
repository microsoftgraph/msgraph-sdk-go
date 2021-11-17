package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i13f9a52f0faf606b01bfa3955e81ee30242ef21fa0101aa087fbede7948f1b27 "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/resetpasscode"
    i1cea6eb12ff2031dedeb2ca8e7862c20336acfdade144ef517a8e5f962fde1f6 "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/cleanwindowsdevice"
    i1e119e421440451e7a53df7e5ca93960ef8e6779e31ca34c2612ddf6e4f5328e "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/windowsdefenderscan"
    i227d0481e3969744ed37e693056ba632a2619393c65c499dec2bb5818ee14670 "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/disablelostmode"
    i2e55400e3a6b1518964ef2369bbc57f4e9a9c4a253afb8263d0c7dd315d4c634 "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/locatedevice"
    i2f0f1986c49cb55444d8f3078b1690316ddcf576fc6f835186d573e09b7ecc57 "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/windowsdefenderupdatesignatures"
    i48553f07197f253783753adbaf7d41660e2a12cf4e334a15ce0a642355263b5e "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/retire"
    i4910726fdcc07627f867a8093698926b8b1ebb256d1ec4e04b21c86a7ee852eb "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/logoutsharedappledeviceactiveuser"
    i4d4ffdac6d6f354742f9f83e0d75cd1f102d1dd45a13d732b9a0c59657c5b53e "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/requestremoteassistance"
    i62df4e156ab0e11abc901553dd3d14a35d775f80cadfec4b46046b962d16f910 "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/wipe"
    i674214359fe29326a03866451c37e7e463c6460897eff0b1713788bcd244994d "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/remotelock"
    i6c92c5c6ee7fb6a63e6a0166f2ec31aa4b47828c1a5ff577350d48023fb89b7a "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/rebootnow"
    i83acc3dd442c42abdbabb9f4fb5b7f362e8c2133c09e2ab7c8978a58dcaba49a "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/devicecompliancepolicystates"
    i93be5ba1e2ed1dcc8ae36f63d85f3edb6388941e09d0e839b2f1a5f66d4919f4 "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/updatewindowsdeviceaccount"
    i97119b774733e620c79e77fe13b9bb32fe545bb2ed890cab4c2a70c7c3e2a442 "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/bypassactivationlock"
    ia9eca82f666e7935a7636c1a6840da82b973be896a102904ee6f6cb28bc3f48a "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/deviceconfigurationstates"
    ibece143defa494d6c2d968eb6a807859f4c7077b7f099cd2c37993aec1843123 "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/syncdevice"
    ic7e7f352418887585e7f8a86e825eb5a78d5d8b112d8eb01d1373aec67367a26 "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/devicecategory"
    iceb6710bce94d2cc31507d20a496d678d25042f6be1b3ab09b691e833abcccc9 "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/recoverpasscode"
    id8aabeb9ac25ed8c755eee299962060c1b32b52e8491a650fa00ecfd2bcb5f28 "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/shutdown"
    id98973b9a80b6d1c37c9c0f1386e5d814f433103bc9e0e7fba3ee12787a54dd8 "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/deleteuserfromsharedappledevice"
    i04ce841e69bdde74a0a5ab2a5f5b64b7e52f890d6dce32e3f002ee145b68f803 "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/devicecompliancepolicystates/item"
    i635fc9cf6abf1cd6f28dfa5885addac5c0d71ed597869597de4dc6569fcace0a "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item/deviceconfigurationstates/item"
)

// Builds and executes requests for operations under \me\managedDevices\{managedDevice-id}
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
// The managed devices associated with the user.
type ManagedDeviceRequestBuilderGetQueryParameters struct {
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
func (m *ManagedDeviceRequestBuilder) BypassActivationLock()(*i97119b774733e620c79e77fe13b9bb32fe545bb2ed890cab4c2a70c7c3e2a442.BypassActivationLockRequestBuilder) {
    return i97119b774733e620c79e77fe13b9bb32fe545bb2ed890cab4c2a70c7c3e2a442.NewBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) CleanWindowsDevice()(*i1cea6eb12ff2031dedeb2ca8e7862c20336acfdade144ef517a8e5f962fde1f6.CleanWindowsDeviceRequestBuilder) {
    return i1cea6eb12ff2031dedeb2ca8e7862c20336acfdade144ef517a8e5f962fde1f6.NewCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ManagedDeviceRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewManagedDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceRequestBuilder) {
    m := &ManagedDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/managedDevices/{managedDevice_id}{?select,expand}";
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
// The managed devices associated with the user.
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
// The managed devices associated with the user.
// Parameters:
//  - options : Options for the request
func (m *ManagedDeviceRequestBuilder) CreateGetRequestInformation(options *ManagedDeviceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The managed devices associated with the user.
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
// The managed devices associated with the user.
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
func (m *ManagedDeviceRequestBuilder) DeleteUserFromSharedAppleDevice()(*id98973b9a80b6d1c37c9c0f1386e5d814f433103bc9e0e7fba3ee12787a54dd8.DeleteUserFromSharedAppleDeviceRequestBuilder) {
    return id98973b9a80b6d1c37c9c0f1386e5d814f433103bc9e0e7fba3ee12787a54dd8.NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DeviceCategory()(*ic7e7f352418887585e7f8a86e825eb5a78d5d8b112d8eb01d1373aec67367a26.DeviceCategoryRequestBuilder) {
    return ic7e7f352418887585e7f8a86e825eb5a78d5d8b112d8eb01d1373aec67367a26.NewDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DeviceCompliancePolicyStates()(*i83acc3dd442c42abdbabb9f4fb5b7f362e8c2133c09e2ab7c8978a58dcaba49a.DeviceCompliancePolicyStatesRequestBuilder) {
    return i83acc3dd442c42abdbabb9f4fb5b7f362e8c2133c09e2ab7c8978a58dcaba49a.NewDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.managedDevices.item.deviceCompliancePolicyStates.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ManagedDeviceRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*i04ce841e69bdde74a0a5ab2a5f5b64b7e52f890d6dce32e3f002ee145b68f803.DeviceCompliancePolicyStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState_id"] = id
    }
    return i04ce841e69bdde74a0a5ab2a5f5b64b7e52f890d6dce32e3f002ee145b68f803.NewDeviceCompliancePolicyStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DeviceConfigurationStates()(*ia9eca82f666e7935a7636c1a6840da82b973be896a102904ee6f6cb28bc3f48a.DeviceConfigurationStatesRequestBuilder) {
    return ia9eca82f666e7935a7636c1a6840da82b973be896a102904ee6f6cb28bc3f48a.NewDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.managedDevices.item.deviceConfigurationStates.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ManagedDeviceRequestBuilder) DeviceConfigurationStatesById(id string)(*i635fc9cf6abf1cd6f28dfa5885addac5c0d71ed597869597de4dc6569fcace0a.DeviceConfigurationStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState_id"] = id
    }
    return i635fc9cf6abf1cd6f28dfa5885addac5c0d71ed597869597de4dc6569fcace0a.NewDeviceConfigurationStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DisableLostMode()(*i227d0481e3969744ed37e693056ba632a2619393c65c499dec2bb5818ee14670.DisableLostModeRequestBuilder) {
    return i227d0481e3969744ed37e693056ba632a2619393c65c499dec2bb5818ee14670.NewDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The managed devices associated with the user.
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
func (m *ManagedDeviceRequestBuilder) LocateDevice()(*i2e55400e3a6b1518964ef2369bbc57f4e9a9c4a253afb8263d0c7dd315d4c634.LocateDeviceRequestBuilder) {
    return i2e55400e3a6b1518964ef2369bbc57f4e9a9c4a253afb8263d0c7dd315d4c634.NewLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*i4910726fdcc07627f867a8093698926b8b1ebb256d1ec4e04b21c86a7ee852eb.LogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return i4910726fdcc07627f867a8093698926b8b1ebb256d1ec4e04b21c86a7ee852eb.NewLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The managed devices associated with the user.
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
func (m *ManagedDeviceRequestBuilder) RebootNow()(*i6c92c5c6ee7fb6a63e6a0166f2ec31aa4b47828c1a5ff577350d48023fb89b7a.RebootNowRequestBuilder) {
    return i6c92c5c6ee7fb6a63e6a0166f2ec31aa4b47828c1a5ff577350d48023fb89b7a.NewRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RecoverPasscode()(*iceb6710bce94d2cc31507d20a496d678d25042f6be1b3ab09b691e833abcccc9.RecoverPasscodeRequestBuilder) {
    return iceb6710bce94d2cc31507d20a496d678d25042f6be1b3ab09b691e833abcccc9.NewRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RemoteLock()(*i674214359fe29326a03866451c37e7e463c6460897eff0b1713788bcd244994d.RemoteLockRequestBuilder) {
    return i674214359fe29326a03866451c37e7e463c6460897eff0b1713788bcd244994d.NewRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RequestRemoteAssistance()(*i4d4ffdac6d6f354742f9f83e0d75cd1f102d1dd45a13d732b9a0c59657c5b53e.RequestRemoteAssistanceRequestBuilder) {
    return i4d4ffdac6d6f354742f9f83e0d75cd1f102d1dd45a13d732b9a0c59657c5b53e.NewRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResetPasscode()(*i13f9a52f0faf606b01bfa3955e81ee30242ef21fa0101aa087fbede7948f1b27.ResetPasscodeRequestBuilder) {
    return i13f9a52f0faf606b01bfa3955e81ee30242ef21fa0101aa087fbede7948f1b27.NewResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Retire()(*i48553f07197f253783753adbaf7d41660e2a12cf4e334a15ce0a642355263b5e.RetireRequestBuilder) {
    return i48553f07197f253783753adbaf7d41660e2a12cf4e334a15ce0a642355263b5e.NewRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ShutDown()(*id8aabeb9ac25ed8c755eee299962060c1b32b52e8491a650fa00ecfd2bcb5f28.ShutDownRequestBuilder) {
    return id8aabeb9ac25ed8c755eee299962060c1b32b52e8491a650fa00ecfd2bcb5f28.NewShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SyncDevice()(*ibece143defa494d6c2d968eb6a807859f4c7077b7f099cd2c37993aec1843123.SyncDeviceRequestBuilder) {
    return ibece143defa494d6c2d968eb6a807859f4c7077b7f099cd2c37993aec1843123.NewSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) UpdateWindowsDeviceAccount()(*i93be5ba1e2ed1dcc8ae36f63d85f3edb6388941e09d0e839b2f1a5f66d4919f4.UpdateWindowsDeviceAccountRequestBuilder) {
    return i93be5ba1e2ed1dcc8ae36f63d85f3edb6388941e09d0e839b2f1a5f66d4919f4.NewUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderScan()(*i1e119e421440451e7a53df7e5ca93960ef8e6779e31ca34c2612ddf6e4f5328e.WindowsDefenderScanRequestBuilder) {
    return i1e119e421440451e7a53df7e5ca93960ef8e6779e31ca34c2612ddf6e4f5328e.NewWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderUpdateSignatures()(*i2f0f1986c49cb55444d8f3078b1690316ddcf576fc6f835186d573e09b7ecc57.WindowsDefenderUpdateSignaturesRequestBuilder) {
    return i2f0f1986c49cb55444d8f3078b1690316ddcf576fc6f835186d573e09b7ecc57.NewWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Wipe()(*i62df4e156ab0e11abc901553dd3d14a35d775f80cadfec4b46046b962d16f910.WipeRequestBuilder) {
    return i62df4e156ab0e11abc901553dd3d14a35d775f80cadfec4b46046b962d16f910.NewWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
