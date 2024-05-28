package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManageddevicesManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
type ManageddevicesManagedDeviceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManageddevicesManagedDeviceItemRequestBuilderGetQueryParameters the list of managed devices.
type ManageddevicesManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManageddevicesManagedDeviceItemRequestBuilderGetQueryParameters
}
// ManageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BypassActivationLock provides operations to call the bypassActivationLock method.
// returns a *ManageddevicesItemBypassactivationlockBypassActivationLockRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) BypassActivationLock()(*ManageddevicesItemBypassactivationlockBypassActivationLockRequestBuilder) {
    return NewManageddevicesItemBypassactivationlockBypassActivationLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CleanWindowsDevice provides operations to call the cleanWindowsDevice method.
// returns a *ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder) {
    return NewManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewManageddevicesManagedDeviceItemRequestBuilderInternal instantiates a new ManageddevicesManagedDeviceItemRequestBuilder and sets the default values.
func NewManageddevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesManagedDeviceItemRequestBuilder) {
    m := &ManageddevicesManagedDeviceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManageddevicesManagedDeviceItemRequestBuilder instantiates a new ManageddevicesManagedDeviceItemRequestBuilder and sets the default values.
func NewManageddevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a managedDevice.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-manageddevice-delete?view=graph-rest-1.0
func (m *ManageddevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeleteUserFromSharedAppleDevice provides operations to call the deleteUserFromSharedAppleDevice method.
// returns a *ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemDevicecategoryDeviceCategoryRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*ManageddevicesItemDevicecategoryDeviceCategoryRequestBuilder) {
    return NewManageddevicesItemDevicecategoryDeviceCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder) {
    return NewManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStatesRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStatesRequestBuilder) {
    return NewManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DisableLostMode provides operations to call the disableLostMode method.
// returns a *ManageddevicesItemDisablelostmodeDisableLostModeRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) DisableLostMode()(*ManageddevicesItemDisablelostmodeDisableLostModeRequestBuilder) {
    return NewManageddevicesItemDisablelostmodeDisableLostModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of managed devices.
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable), nil
}
// LocateDevice provides operations to call the locateDevice method.
// returns a *ManageddevicesItemLocatedeviceLocateDeviceRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) LocateDevice()(*ManageddevicesItemLocatedeviceLocateDeviceRequestBuilder) {
    return NewManageddevicesItemLocatedeviceLocateDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogCollectionRequests provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) LogCollectionRequests()(*ManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder) {
    return NewManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
// returns a *ManageddevicesItemLogoutsharedappledeviceactiveuserLogoutSharedAppleDeviceActiveUserRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*ManageddevicesItemLogoutsharedappledeviceactiveuserLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewManageddevicesItemLogoutsharedappledeviceactiveuserLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property managedDevices in deviceManagement
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, requestConfiguration *ManageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable), nil
}
// RebootNow provides operations to call the rebootNow method.
// returns a *ManageddevicesItemRebootnowRebootNowRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) RebootNow()(*ManageddevicesItemRebootnowRebootNowRequestBuilder) {
    return NewManageddevicesItemRebootnowRebootNowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RecoverPasscode provides operations to call the recoverPasscode method.
// returns a *ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) RecoverPasscode()(*ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder) {
    return NewManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoteLock provides operations to call the remoteLock method.
// returns a *ManageddevicesItemRemotelockRemoteLockRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) RemoteLock()(*ManageddevicesItemRemotelockRemoteLockRequestBuilder) {
    return NewManageddevicesItemRemotelockRemoteLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
// returns a *ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder) {
    return NewManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResetPasscode provides operations to call the resetPasscode method.
// returns a *ManageddevicesItemResetpasscodeResetPasscodeRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) ResetPasscode()(*ManageddevicesItemResetpasscodeResetPasscodeRequestBuilder) {
    return NewManageddevicesItemResetpasscodeResetPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Retire provides operations to call the retire method.
// returns a *ManageddevicesItemRetireRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) Retire()(*ManageddevicesItemRetireRequestBuilder) {
    return NewManageddevicesItemRetireRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShutDown provides operations to call the shutDown method.
// returns a *ManageddevicesItemShutdownShutDownRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) ShutDown()(*ManageddevicesItemShutdownShutDownRequestBuilder) {
    return NewManageddevicesItemShutdownShutDownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SyncDevice provides operations to call the syncDevice method.
// returns a *ManageddevicesItemSyncdeviceSyncDeviceRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) SyncDevice()(*ManageddevicesItemSyncdeviceSyncDeviceRequestBuilder) {
    return NewManageddevicesItemSyncdeviceSyncDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes a managedDevice.
// returns a *RequestInformation when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of managed devices.
// returns a *RequestInformation when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property managedDevices in deviceManagement
// returns a *RequestInformation when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, requestConfiguration *ManageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// UpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
// returns a *ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder) {
    return NewManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemUsersRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) Users()(*ManageddevicesItemUsersRequestBuilder) {
    return NewManageddevicesItemUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderScan provides operations to call the windowsDefenderScan method.
// returns a *ManageddevicesItemWindowsdefenderscanWindowsDefenderScanRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*ManageddevicesItemWindowsdefenderscanWindowsDefenderScanRequestBuilder) {
    return NewManageddevicesItemWindowsdefenderscanWindowsDefenderScanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
// returns a *ManageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*ManageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewManageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) {
    return NewManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Wipe provides operations to call the wipe method.
// returns a *ManageddevicesItemWipeRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) Wipe()(*ManageddevicesItemWipeRequestBuilder) {
    return NewManageddevicesItemWipeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManageddevicesManagedDeviceItemRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesManagedDeviceItemRequestBuilder) {
    return NewManageddevicesManagedDeviceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
