package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemManagedDevicesManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.user entity.
type ItemManagedDevicesManagedDeviceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters the managed devices associated with the user.
type ItemManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters
}
// ItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BypassActivationLock provides operations to call the bypassActivationLock method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) BypassActivationLock()(*ItemManagedDevicesItemBypassActivationLockRequestBuilder) {
    return NewItemManagedDevicesItemBypassActivationLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CleanWindowsDevice provides operations to call the cleanWindowsDevice method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*ItemManagedDevicesItemCleanWindowsDeviceRequestBuilder) {
    return NewItemManagedDevicesItemCleanWindowsDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemManagedDevicesManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewItemManagedDevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesManagedDeviceItemRequestBuilder) {
    m := &ItemManagedDevicesManagedDeviceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemManagedDevicesManagedDeviceItemRequestBuilder instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewItemManagedDevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managedDevices for users
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeleteUserFromSharedAppleDevice provides operations to call the deleteUserFromSharedAppleDevice method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*ItemManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewItemManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*ItemManagedDevicesItemDeviceCategoryRequestBuilder) {
    return NewItemManagedDevicesItemDeviceCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*ItemManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder) {
    return NewItemManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ItemManagedDevicesItemDeviceConfigurationStatesRequestBuilder) {
    return NewItemManagedDevicesItemDeviceConfigurationStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DisableLostMode provides operations to call the disableLostMode method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DisableLostMode()(*ItemManagedDevicesItemDisableLostModeRequestBuilder) {
    return NewItemManagedDevicesItemDisableLostModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the managed devices associated with the user.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) LocateDevice()(*ItemManagedDevicesItemLocateDeviceRequestBuilder) {
    return NewItemManagedDevicesItemLocateDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*ItemManagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewItemManagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property managedDevices in users
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RebootNow()(*ItemManagedDevicesItemRebootNowRequestBuilder) {
    return NewItemManagedDevicesItemRebootNowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RecoverPasscode provides operations to call the recoverPasscode method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RecoverPasscode()(*ItemManagedDevicesItemRecoverPasscodeRequestBuilder) {
    return NewItemManagedDevicesItemRecoverPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoteLock provides operations to call the remoteLock method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RemoteLock()(*ItemManagedDevicesItemRemoteLockRequestBuilder) {
    return NewItemManagedDevicesItemRemoteLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*ItemManagedDevicesItemRequestRemoteAssistanceRequestBuilder) {
    return NewItemManagedDevicesItemRequestRemoteAssistanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResetPasscode provides operations to call the resetPasscode method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ResetPasscode()(*ItemManagedDevicesItemResetPasscodeRequestBuilder) {
    return NewItemManagedDevicesItemResetPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Retire provides operations to call the retire method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Retire()(*ItemManagedDevicesItemRetireRequestBuilder) {
    return NewItemManagedDevicesItemRetireRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShutDown provides operations to call the shutDown method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ShutDown()(*ItemManagedDevicesItemShutDownRequestBuilder) {
    return NewItemManagedDevicesItemShutDownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SyncDevice provides operations to call the syncDevice method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) SyncDevice()(*ItemManagedDevicesItemSyncDeviceRequestBuilder) {
    return NewItemManagedDevicesItemSyncDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property managedDevices for users
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the managed devices associated with the user.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property managedDevices in users
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// UpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*ItemManagedDevicesItemUpdateWindowsDeviceAccountRequestBuilder) {
    return NewItemManagedDevicesItemUpdateWindowsDeviceAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Users()(*ItemManagedDevicesItemUsersRequestBuilder) {
    return NewItemManagedDevicesItemUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderScan provides operations to call the windowsDefenderScan method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*ItemManagedDevicesItemWindowsDefenderScanRequestBuilder) {
    return NewItemManagedDevicesItemWindowsDefenderScanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*ItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Wipe provides operations to call the wipe method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Wipe()(*ItemManagedDevicesItemWipeRequestBuilder) {
    return NewItemManagedDevicesItemWipeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
