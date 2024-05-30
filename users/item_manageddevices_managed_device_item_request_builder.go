package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemManageddevicesManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.user entity.
type ItemManageddevicesManagedDeviceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemManageddevicesManagedDeviceItemRequestBuilderGetQueryParameters the managed devices associated with the user.
type ItemManageddevicesManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemManageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManageddevicesManagedDeviceItemRequestBuilderGetQueryParameters
}
// ItemManageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BypassActivationLock provides operations to call the bypassActivationLock method.
// returns a *ItemManageddevicesItemBypassactivationlockBypassActivationLockRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) BypassActivationLock()(*ItemManageddevicesItemBypassactivationlockBypassActivationLockRequestBuilder) {
    return NewItemManageddevicesItemBypassactivationlockBypassActivationLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CleanWindowsDevice provides operations to call the cleanWindowsDevice method.
// returns a *ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder) {
    return NewItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemManageddevicesManagedDeviceItemRequestBuilderInternal instantiates a new ItemManageddevicesManagedDeviceItemRequestBuilder and sets the default values.
func NewItemManageddevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesManagedDeviceItemRequestBuilder) {
    m := &ItemManageddevicesManagedDeviceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemManageddevicesManagedDeviceItemRequestBuilder instantiates a new ItemManageddevicesManagedDeviceItemRequestBuilder and sets the default values.
func NewItemManageddevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managedDevices for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemManageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *ItemManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*ItemManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewItemManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemDevicecategoryDeviceCategoryRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*ItemManageddevicesItemDevicecategoryDeviceCategoryRequestBuilder) {
    return NewItemManageddevicesItemDevicecategoryDeviceCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*ItemManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder) {
    return NewItemManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStatesRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStatesRequestBuilder) {
    return NewItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DisableLostMode provides operations to call the disableLostMode method.
// returns a *ItemManageddevicesItemDisablelostmodeDisableLostModeRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) DisableLostMode()(*ItemManageddevicesItemDisablelostmodeDisableLostModeRequestBuilder) {
    return NewItemManageddevicesItemDisablelostmodeDisableLostModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the managed devices associated with the user.
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, error) {
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
// returns a *ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) LocateDevice()(*ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder) {
    return NewItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogCollectionRequests provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) LogCollectionRequests()(*ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder) {
    return NewItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
// returns a *ItemManageddevicesItemLogoutsharedappledeviceactiveuserLogoutSharedAppleDeviceActiveUserRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*ItemManageddevicesItemLogoutsharedappledeviceactiveuserLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewItemManageddevicesItemLogoutsharedappledeviceactiveuserLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property managedDevices in users
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, requestConfiguration *ItemManageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, error) {
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
// returns a *ItemManageddevicesItemRebootnowRebootNowRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) RebootNow()(*ItemManageddevicesItemRebootnowRebootNowRequestBuilder) {
    return NewItemManageddevicesItemRebootnowRebootNowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RecoverPasscode provides operations to call the recoverPasscode method.
// returns a *ItemManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) RecoverPasscode()(*ItemManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder) {
    return NewItemManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoteLock provides operations to call the remoteLock method.
// returns a *ItemManageddevicesItemRemotelockRemoteLockRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) RemoteLock()(*ItemManageddevicesItemRemotelockRemoteLockRequestBuilder) {
    return NewItemManageddevicesItemRemotelockRemoteLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
// returns a *ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder) {
    return NewItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResetPasscode provides operations to call the resetPasscode method.
// returns a *ItemManageddevicesItemResetpasscodeResetPasscodeRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) ResetPasscode()(*ItemManageddevicesItemResetpasscodeResetPasscodeRequestBuilder) {
    return NewItemManageddevicesItemResetpasscodeResetPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Retire provides operations to call the retire method.
// returns a *ItemManageddevicesItemRetireRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) Retire()(*ItemManageddevicesItemRetireRequestBuilder) {
    return NewItemManageddevicesItemRetireRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShutDown provides operations to call the shutDown method.
// returns a *ItemManageddevicesItemShutdownShutDownRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) ShutDown()(*ItemManageddevicesItemShutdownShutDownRequestBuilder) {
    return NewItemManageddevicesItemShutdownShutDownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SyncDevice provides operations to call the syncDevice method.
// returns a *ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) SyncDevice()(*ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder) {
    return NewItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property managedDevices for users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the managed devices associated with the user.
// returns a *RequestInformation when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managedDevices in users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, requestConfiguration *ItemManageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*ItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder) {
    return NewItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemUsersRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) Users()(*ItemManageddevicesItemUsersRequestBuilder) {
    return NewItemManageddevicesItemUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderScan provides operations to call the windowsDefenderScan method.
// returns a *ItemManageddevicesItemWindowsdefenderscanWindowsDefenderScanRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*ItemManageddevicesItemWindowsdefenderscanWindowsDefenderScanRequestBuilder) {
    return NewItemManageddevicesItemWindowsdefenderscanWindowsDefenderScanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
// returns a *ItemManageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*ItemManageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewItemManageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) {
    return NewItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Wipe provides operations to call the wipe method.
// returns a *ItemManageddevicesItemWipeRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) Wipe()(*ItemManageddevicesItemWipeRequestBuilder) {
    return NewItemManageddevicesItemWipeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemManageddevicesManagedDeviceItemRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesManagedDeviceItemRequestBuilder) {
    return NewItemManageddevicesManagedDeviceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
