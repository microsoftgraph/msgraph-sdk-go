package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i01fefd20a716f509a0059f1f839faadc9e4de1045dc1dba63a10ededffd308ab "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/checkmembergroups"
    i16d80f33722f81a9d00614f3c8bd4242732c7a1278da72dc646b0a31fa41a5b0 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/assignlicense"
    i2bc42c0148def38705b9e9d5124f8119a7f2df026c9652f9cad139803f0e6e73 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/checkmemberobjects"
    i40e2067052230848aeff5ba0626772c2c2bc1a0f2df1b41951f1af8b7e6327a6 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/sendmail"
    i4518a1659833f65ae456846ba720f161b32f49e1f9a93024804a24a61fcec8d6 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmanagedappdiagnosticstatuses"
    i45c64dcbdcb1ba3709d233ba288832bd17a9f13fc0f47a315d05fe503e568e5d "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/wipemanagedappregistrationsbydevicetag"
    i4ad61960adead1812cffa62750d87d64602ee06a9ebc29f325bddb93cf79e519 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/reminderviewwithstartdatetimewithenddatetime"
    i7f28e175de1049e97efeeb91e0d1335a1ba7649a1f8f7411842f3a3158e65bba "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmailtips"
    i8e6ae30717b99801203f5e6933a063a0d60df18d43c4d75255ae07cd7936c941 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/revokesigninsessions"
    i913032f1265c753b0761330e3803cd0e47dad5ecf2a56fb041aa714eb253ab7b "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmembergroups"
    i96741dc37adb750fd11b0ff44ed749aaf54eb8a02c1263c880249966f28495c5 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/restore"
    i9a1272dd9d51f172ed5df48d7741fd52d69297eb2b62dff5f3f8101bebcbdfbc "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/translateexchangeids"
    ia381ac27535bac6c9328dc626b7eabe5148a34d77316dcb3b0ebd116839b85c0 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/changepassword"
    ibaad7da5c0850536d714c0664e7a70d63f71ee20d07d49df7be65dabaf031675 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmemberobjects"
    ic4bfd067568ab7da377c9e5fc3a429e0f641db6f15e239c660ed4b1d49d14354 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmanagedapppolicies"
    icd243b40bc934494d33c1449e4fd7baccf9ca58b22d0afb406b8cd4a779491db "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/exportpersonaldata"
    ice89d24641221f302f2f2daf89b97b39c3a36204b2c689635ae1fbc3fc503936 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/findmeetingtimes"
    id036a2aec3f132e46e8d18eaf088dbdc272322b065ec1a27c77f87645970f39b "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/removealldevicesfrommanagement"
    idecf32ce6137514de6b56f843a3cc8fffce1cda9cfec50509af6168d5b536006 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/reprocesslicenseassignment"
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
func (m *UserRequestBuilder) AssignLicense()(*i16d80f33722f81a9d00614f3c8bd4242732c7a1278da72dc646b0a31fa41a5b0.AssignLicenseRequestBuilder) {
    return i16d80f33722f81a9d00614f3c8bd4242732c7a1278da72dc646b0a31fa41a5b0.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*ia381ac27535bac6c9328dc626b7eabe5148a34d77316dcb3b0ebd116839b85c0.ChangePasswordRequestBuilder) {
    return ia381ac27535bac6c9328dc626b7eabe5148a34d77316dcb3b0ebd116839b85c0.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i01fefd20a716f509a0059f1f839faadc9e4de1045dc1dba63a10ededffd308ab.CheckMemberGroupsRequestBuilder) {
    return i01fefd20a716f509a0059f1f839faadc9e4de1045dc1dba63a10ededffd308ab.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i2bc42c0148def38705b9e9d5124f8119a7f2df026c9652f9cad139803f0e6e73.CheckMemberObjectsRequestBuilder) {
    return i2bc42c0148def38705b9e9d5124f8119a7f2df026c9652f9cad139803f0e6e73.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/registeredUsers/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
    requestInfo.Headers["Accept"] = "application/json"
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
func (m *UserRequestBuilder) ExportPersonalData()(*icd243b40bc934494d33c1449e4fd7baccf9ca58b22d0afb406b8cd4a779491db.ExportPersonalDataRequestBuilder) {
    return icd243b40bc934494d33c1449e4fd7baccf9ca58b22d0afb406b8cd4a779491db.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*ice89d24641221f302f2f2daf89b97b39c3a36204b2c689635ae1fbc3fc503936.FindMeetingTimesRequestBuilder) {
    return ice89d24641221f302f2f2daf89b97b39c3a36204b2c689635ae1fbc3fc503936.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i7f28e175de1049e97efeeb91e0d1335a1ba7649a1f8f7411842f3a3158e65bba.GetMailTipsRequestBuilder) {
    return i7f28e175de1049e97efeeb91e0d1335a1ba7649a1f8f7411842f3a3158e65bba.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i4518a1659833f65ae456846ba720f161b32f49e1f9a93024804a24a61fcec8d6.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i4518a1659833f65ae456846ba720f161b32f49e1f9a93024804a24a61fcec8d6.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*ic4bfd067568ab7da377c9e5fc3a429e0f641db6f15e239c660ed4b1d49d14354.GetManagedAppPoliciesRequestBuilder) {
    return ic4bfd067568ab7da377c9e5fc3a429e0f641db6f15e239c660ed4b1d49d14354.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i913032f1265c753b0761330e3803cd0e47dad5ecf2a56fb041aa714eb253ab7b.GetMemberGroupsRequestBuilder) {
    return i913032f1265c753b0761330e3803cd0e47dad5ecf2a56fb041aa714eb253ab7b.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*ibaad7da5c0850536d714c0664e7a70d63f71ee20d07d49df7be65dabaf031675.GetMemberObjectsRequestBuilder) {
    return ibaad7da5c0850536d714c0664e7a70d63f71ee20d07d49df7be65dabaf031675.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i4ad61960adead1812cffa62750d87d64602ee06a9ebc29f325bddb93cf79e519.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i4ad61960adead1812cffa62750d87d64602ee06a9ebc29f325bddb93cf79e519.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*id036a2aec3f132e46e8d18eaf088dbdc272322b065ec1a27c77f87645970f39b.RemoveAllDevicesFromManagementRequestBuilder) {
    return id036a2aec3f132e46e8d18eaf088dbdc272322b065ec1a27c77f87645970f39b.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*idecf32ce6137514de6b56f843a3cc8fffce1cda9cfec50509af6168d5b536006.ReprocessLicenseAssignmentRequestBuilder) {
    return idecf32ce6137514de6b56f843a3cc8fffce1cda9cfec50509af6168d5b536006.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i96741dc37adb750fd11b0ff44ed749aaf54eb8a02c1263c880249966f28495c5.RestoreRequestBuilder) {
    return i96741dc37adb750fd11b0ff44ed749aaf54eb8a02c1263c880249966f28495c5.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i8e6ae30717b99801203f5e6933a063a0d60df18d43c4d75255ae07cd7936c941.RevokeSignInSessionsRequestBuilder) {
    return i8e6ae30717b99801203f5e6933a063a0d60df18d43c4d75255ae07cd7936c941.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i40e2067052230848aeff5ba0626772c2c2bc1a0f2df1b41951f1af8b7e6327a6.SendMailRequestBuilder) {
    return i40e2067052230848aeff5ba0626772c2c2bc1a0f2df1b41951f1af8b7e6327a6.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i9a1272dd9d51f172ed5df48d7741fd52d69297eb2b62dff5f3f8101bebcbdfbc.TranslateExchangeIdsRequestBuilder) {
    return i9a1272dd9d51f172ed5df48d7741fd52d69297eb2b62dff5f3f8101bebcbdfbc.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i45c64dcbdcb1ba3709d233ba288832bd17a9f13fc0f47a315d05fe503e568e5d.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i45c64dcbdcb1ba3709d233ba288832bd17a9f13fc0f47a315d05fe503e568e5d.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
