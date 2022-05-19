package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0b7050f71890250c44b155e10dead034e6c1cfa152dc03fe1bcb1fd0bdbee98e "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/checkmemberobjects"
    i0d18ca633ec32d8c575dfe4410fee8619d93553f6806f102323889e3fb9247f5 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/removealldevicesfrommanagement"
    i0f459488cb0501199f5086ba55df0b7167159ed2986ed0e032c1f8ea37f246cf "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/reprocesslicenseassignment"
    i1b61640fd0c85285e78a563343e709ff7a6adac9a6a309fc87104c19d3049c09 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/exportpersonaldata"
    i1efba902b65b18adea13a9469c11ca1c35342ec97f62b0fd5052f4af85ab2b51 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/reminderviewwithstartdatetimewithenddatetime"
    i40147940fb5f41edb94bd1ec3d6d0663f46c35f752e54afe06b7ac06169ba91c "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/translateexchangeids"
    i4c3cfac570bbf857b547a9ecec4e5a43064db617f319792a4dbac12e28b0171f "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmanagedapppolicies"
    i539f07573bf24ab0d6cec1f69296e43ec4e40ae4f2a306511b0ce6cac9ad2289 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/assignlicense"
    i667d3a721457d4ca03e734493d5c5fc73bbc9f4e3e9877d0c6684640ad0d4b17 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/findmeetingtimes"
    i80f22358710504f705e7834d9e44ce8740e6c78dff4ea001dcf8ef2ea8ea3f7d "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/revokesigninsessions"
    i91ce69fcb3cc343ef7138ff41ded6baa7ebf9d9adb31759f8814b5a849d52172 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmembergroups"
    i97a36b9a1e027cde01dc2beee6b6d996542a5776c9de5768536499eda118dc66 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/restore"
    ia7cb9ee65a25490287923f39b9a4f8600c714023b1648b1723e971df1c263099 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/changepassword"
    ib8513e2a59c6f8c3b2ef2bceaf607c449d6cb0d4de1209a2cc4e2f249b7e3990 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/checkmembergroups"
    id5a5234bc8e65b9fc1a6e0c5832051369d823676a665940e0cb46a94bdb56c4e "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmailtips"
    idb72fe7ce7f4b5a52d1bb257860dcb78f5809a7a736ef8159871ead4c74067b4 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmemberobjects"
    ie51c6fa073ee4b27dd97395a72a963acf7609e590c25ac02b7c862159b843585 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/wipemanagedappregistrationsbydevicetag"
    ie6147f956ce8f19018d6143b287a254f0e18db9b5a5ce359ae45646a57db5be6 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmanagedappdiagnosticstatuses"
    ifc4bd76be7da393d692397de11446a6e3e58e4574d4bf029b22f4cd7bb03e3f0 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/sendmail"
)

// UserRequestBuilder casts the previous resource to user.
type UserRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UserRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.user
type UserRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserRequestBuilderGetQueryParameters
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i539f07573bf24ab0d6cec1f69296e43ec4e40ae4f2a306511b0ce6cac9ad2289.AssignLicenseRequestBuilder) {
    return i539f07573bf24ab0d6cec1f69296e43ec4e40ae4f2a306511b0ce6cac9ad2289.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*ia7cb9ee65a25490287923f39b9a4f8600c714023b1648b1723e971df1c263099.ChangePasswordRequestBuilder) {
    return ia7cb9ee65a25490287923f39b9a4f8600c714023b1648b1723e971df1c263099.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*ib8513e2a59c6f8c3b2ef2bceaf607c449d6cb0d4de1209a2cc4e2f249b7e3990.CheckMemberGroupsRequestBuilder) {
    return ib8513e2a59c6f8c3b2ef2bceaf607c449d6cb0d4de1209a2cc4e2f249b7e3990.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i0b7050f71890250c44b155e10dead034e6c1cfa152dc03fe1bcb1fd0bdbee98e.CheckMemberObjectsRequestBuilder) {
    return i0b7050f71890250c44b155e10dead034e6c1cfa152dc03fe1bcb1fd0bdbee98e.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredUsers/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserRequestBuilder instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *UserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i1b61640fd0c85285e78a563343e709ff7a6adac9a6a309fc87104c19d3049c09.ExportPersonalDataRequestBuilder) {
    return i1b61640fd0c85285e78a563343e709ff7a6adac9a6a309fc87104c19d3049c09.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i667d3a721457d4ca03e734493d5c5fc73bbc9f4e3e9877d0c6684640ad0d4b17.FindMeetingTimesRequestBuilder) {
    return i667d3a721457d4ca03e734493d5c5fc73bbc9f4e3e9877d0c6684640ad0d4b17.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*id5a5234bc8e65b9fc1a6e0c5832051369d823676a665940e0cb46a94bdb56c4e.GetMailTipsRequestBuilder) {
    return id5a5234bc8e65b9fc1a6e0c5832051369d823676a665940e0cb46a94bdb56c4e.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*ie6147f956ce8f19018d6143b287a254f0e18db9b5a5ce359ae45646a57db5be6.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return ie6147f956ce8f19018d6143b287a254f0e18db9b5a5ce359ae45646a57db5be6.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i4c3cfac570bbf857b547a9ecec4e5a43064db617f319792a4dbac12e28b0171f.GetManagedAppPoliciesRequestBuilder) {
    return i4c3cfac570bbf857b547a9ecec4e5a43064db617f319792a4dbac12e28b0171f.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i91ce69fcb3cc343ef7138ff41ded6baa7ebf9d9adb31759f8814b5a849d52172.GetMemberGroupsRequestBuilder) {
    return i91ce69fcb3cc343ef7138ff41ded6baa7ebf9d9adb31759f8814b5a849d52172.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*idb72fe7ce7f4b5a52d1bb257860dcb78f5809a7a736ef8159871ead4c74067b4.GetMemberObjectsRequestBuilder) {
    return idb72fe7ce7f4b5a52d1bb257860dcb78f5809a7a736ef8159871ead4c74067b4.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *UserRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable), nil
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i1efba902b65b18adea13a9469c11ca1c35342ec97f62b0fd5052f4af85ab2b51.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i1efba902b65b18adea13a9469c11ca1c35342ec97f62b0fd5052f4af85ab2b51.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i0d18ca633ec32d8c575dfe4410fee8619d93553f6806f102323889e3fb9247f5.RemoveAllDevicesFromManagementRequestBuilder) {
    return i0d18ca633ec32d8c575dfe4410fee8619d93553f6806f102323889e3fb9247f5.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i0f459488cb0501199f5086ba55df0b7167159ed2986ed0e032c1f8ea37f246cf.ReprocessLicenseAssignmentRequestBuilder) {
    return i0f459488cb0501199f5086ba55df0b7167159ed2986ed0e032c1f8ea37f246cf.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i97a36b9a1e027cde01dc2beee6b6d996542a5776c9de5768536499eda118dc66.RestoreRequestBuilder) {
    return i97a36b9a1e027cde01dc2beee6b6d996542a5776c9de5768536499eda118dc66.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i80f22358710504f705e7834d9e44ce8740e6c78dff4ea001dcf8ef2ea8ea3f7d.RevokeSignInSessionsRequestBuilder) {
    return i80f22358710504f705e7834d9e44ce8740e6c78dff4ea001dcf8ef2ea8ea3f7d.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*ifc4bd76be7da393d692397de11446a6e3e58e4574d4bf029b22f4cd7bb03e3f0.SendMailRequestBuilder) {
    return ifc4bd76be7da393d692397de11446a6e3e58e4574d4bf029b22f4cd7bb03e3f0.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i40147940fb5f41edb94bd1ec3d6d0663f46c35f752e54afe06b7ac06169ba91c.TranslateExchangeIdsRequestBuilder) {
    return i40147940fb5f41edb94bd1ec3d6d0663f46c35f752e54afe06b7ac06169ba91c.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*ie51c6fa073ee4b27dd97395a72a963acf7609e590c25ac02b7c862159b843585.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return ie51c6fa073ee4b27dd97395a72a963acf7609e590c25ac02b7c862159b843585.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
