package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
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

// ManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.user entity.
type ManagedDeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedDeviceItemRequestBuilderDeleteOptions options for Delete
type ManagedDeviceItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedDeviceItemRequestBuilderGetOptions options for Get
type ManagedDeviceItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedDeviceItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedDeviceItemRequestBuilderGetQueryParameters the managed devices associated with the user.
type ManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ManagedDeviceItemRequestBuilderPatchOptions options for Patch
type ManagedDeviceItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDeviceable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ManagedDeviceItemRequestBuilder) BypassActivationLock()(*i97119b774733e620c79e77fe13b9bb32fe545bb2ed890cab4c2a70c7c3e2a442.BypassActivationLockRequestBuilder) {
    return i97119b774733e620c79e77fe13b9bb32fe545bb2ed890cab4c2a70c7c3e2a442.NewBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*i1cea6eb12ff2031dedeb2ca8e7862c20336acfdade144ef517a8e5f962fde1f6.CleanWindowsDeviceRequestBuilder) {
    return i1cea6eb12ff2031dedeb2ca8e7862c20336acfdade144ef517a8e5f962fde1f6.NewCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceItemRequestBuilder) {
    m := &ManagedDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/managedDevices/{managedDevice_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedDeviceItemRequestBuilder instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedDevices for me
func (m *ManagedDeviceItemRequestBuilder) CreateDeleteRequestInformation(options *ManagedDeviceItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the managed devices associated with the user.
func (m *ManagedDeviceItemRequestBuilder) CreateGetRequestInformation(options *ManagedDeviceItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property managedDevices in me
func (m *ManagedDeviceItemRequestBuilder) CreatePatchRequestInformation(options *ManagedDeviceItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property managedDevices for me
func (m *ManagedDeviceItemRequestBuilder) Delete(options *ManagedDeviceItemRequestBuilderDeleteOptions)(error) {
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
func (m *ManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*id98973b9a80b6d1c37c9c0f1386e5d814f433103bc9e0e7fba3ee12787a54dd8.DeleteUserFromSharedAppleDeviceRequestBuilder) {
    return id98973b9a80b6d1c37c9c0f1386e5d814f433103bc9e0e7fba3ee12787a54dd8.NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) DeviceCategory()(*ic7e7f352418887585e7f8a86e825eb5a78d5d8b112d8eb01d1373aec67367a26.DeviceCategoryRequestBuilder) {
    return ic7e7f352418887585e7f8a86e825eb5a78d5d8b112d8eb01d1373aec67367a26.NewDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*i83acc3dd442c42abdbabb9f4fb5b7f362e8c2133c09e2ab7c8978a58dcaba49a.DeviceCompliancePolicyStatesRequestBuilder) {
    return i83acc3dd442c42abdbabb9f4fb5b7f362e8c2133c09e2ab7c8978a58dcaba49a.NewDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStatesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.managedDevices.item.deviceCompliancePolicyStates.item collection
func (m *ManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*i04ce841e69bdde74a0a5ab2a5f5b64b7e52f890d6dce32e3f002ee145b68f803.DeviceCompliancePolicyStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState_id"] = id
    }
    return i04ce841e69bdde74a0a5ab2a5f5b64b7e52f890d6dce32e3f002ee145b68f803.NewDeviceCompliancePolicyStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ia9eca82f666e7935a7636c1a6840da82b973be896a102904ee6f6cb28bc3f48a.DeviceConfigurationStatesRequestBuilder) {
    return ia9eca82f666e7935a7636c1a6840da82b973be896a102904ee6f6cb28bc3f48a.NewDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationStatesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.managedDevices.item.deviceConfigurationStates.item collection
func (m *ManagedDeviceItemRequestBuilder) DeviceConfigurationStatesById(id string)(*i635fc9cf6abf1cd6f28dfa5885addac5c0d71ed597869597de4dc6569fcace0a.DeviceConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState_id"] = id
    }
    return i635fc9cf6abf1cd6f28dfa5885addac5c0d71ed597869597de4dc6569fcace0a.NewDeviceConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) DisableLostMode()(*i227d0481e3969744ed37e693056ba632a2619393c65c499dec2bb5818ee14670.DisableLostModeRequestBuilder) {
    return i227d0481e3969744ed37e693056ba632a2619393c65c499dec2bb5818ee14670.NewDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the managed devices associated with the user.
func (m *ManagedDeviceItemRequestBuilder) Get(options *ManagedDeviceItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDeviceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateManagedDeviceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDeviceable), nil
}
func (m *ManagedDeviceItemRequestBuilder) LocateDevice()(*i2e55400e3a6b1518964ef2369bbc57f4e9a9c4a253afb8263d0c7dd315d4c634.LocateDeviceRequestBuilder) {
    return i2e55400e3a6b1518964ef2369bbc57f4e9a9c4a253afb8263d0c7dd315d4c634.NewLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*i4910726fdcc07627f867a8093698926b8b1ebb256d1ec4e04b21c86a7ee852eb.LogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return i4910726fdcc07627f867a8093698926b8b1ebb256d1ec4e04b21c86a7ee852eb.NewLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property managedDevices in me
func (m *ManagedDeviceItemRequestBuilder) Patch(options *ManagedDeviceItemRequestBuilderPatchOptions)(error) {
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
func (m *ManagedDeviceItemRequestBuilder) RebootNow()(*i6c92c5c6ee7fb6a63e6a0166f2ec31aa4b47828c1a5ff577350d48023fb89b7a.RebootNowRequestBuilder) {
    return i6c92c5c6ee7fb6a63e6a0166f2ec31aa4b47828c1a5ff577350d48023fb89b7a.NewRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) RecoverPasscode()(*iceb6710bce94d2cc31507d20a496d678d25042f6be1b3ab09b691e833abcccc9.RecoverPasscodeRequestBuilder) {
    return iceb6710bce94d2cc31507d20a496d678d25042f6be1b3ab09b691e833abcccc9.NewRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) RemoteLock()(*i674214359fe29326a03866451c37e7e463c6460897eff0b1713788bcd244994d.RemoteLockRequestBuilder) {
    return i674214359fe29326a03866451c37e7e463c6460897eff0b1713788bcd244994d.NewRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*i4d4ffdac6d6f354742f9f83e0d75cd1f102d1dd45a13d732b9a0c59657c5b53e.RequestRemoteAssistanceRequestBuilder) {
    return i4d4ffdac6d6f354742f9f83e0d75cd1f102d1dd45a13d732b9a0c59657c5b53e.NewRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) ResetPasscode()(*i13f9a52f0faf606b01bfa3955e81ee30242ef21fa0101aa087fbede7948f1b27.ResetPasscodeRequestBuilder) {
    return i13f9a52f0faf606b01bfa3955e81ee30242ef21fa0101aa087fbede7948f1b27.NewResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) Retire()(*i48553f07197f253783753adbaf7d41660e2a12cf4e334a15ce0a642355263b5e.RetireRequestBuilder) {
    return i48553f07197f253783753adbaf7d41660e2a12cf4e334a15ce0a642355263b5e.NewRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) ShutDown()(*id8aabeb9ac25ed8c755eee299962060c1b32b52e8491a650fa00ecfd2bcb5f28.ShutDownRequestBuilder) {
    return id8aabeb9ac25ed8c755eee299962060c1b32b52e8491a650fa00ecfd2bcb5f28.NewShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) SyncDevice()(*ibece143defa494d6c2d968eb6a807859f4c7077b7f099cd2c37993aec1843123.SyncDeviceRequestBuilder) {
    return ibece143defa494d6c2d968eb6a807859f4c7077b7f099cd2c37993aec1843123.NewSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*i93be5ba1e2ed1dcc8ae36f63d85f3edb6388941e09d0e839b2f1a5f66d4919f4.UpdateWindowsDeviceAccountRequestBuilder) {
    return i93be5ba1e2ed1dcc8ae36f63d85f3edb6388941e09d0e839b2f1a5f66d4919f4.NewUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*i1e119e421440451e7a53df7e5ca93960ef8e6779e31ca34c2612ddf6e4f5328e.WindowsDefenderScanRequestBuilder) {
    return i1e119e421440451e7a53df7e5ca93960ef8e6779e31ca34c2612ddf6e4f5328e.NewWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*i2f0f1986c49cb55444d8f3078b1690316ddcf576fc6f835186d573e09b7ecc57.WindowsDefenderUpdateSignaturesRequestBuilder) {
    return i2f0f1986c49cb55444d8f3078b1690316ddcf576fc6f835186d573e09b7ecc57.NewWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) Wipe()(*i62df4e156ab0e11abc901553dd3d14a35d775f80cadfec4b46046b962d16f910.WipeRequestBuilder) {
    return i62df4e156ab0e11abc901553dd3d14a35d775f80cadfec4b46046b962d16f910.NewWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
