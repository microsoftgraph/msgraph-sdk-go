package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i2070c445678bb477a0e6ece4fba1feede3e05505713b95a26e0f64fe0e523d1e "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmanagedappdiagnosticstatuses"
    i226ecc447edd24870796d566cc7ccfe04b5f8a423bd59c81b8d3b01905cd5851 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/restore"
    i27530cf646a7a9a035e401f73f898080c9813017863ea6bd12c4f503bfd11616 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/exportpersonaldata"
    i3980cf70ba7e8face0024e6f88ab6d601e20ef50a6dc80273b1f137db76e3b58 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/removealldevicesfrommanagement"
    i3f45a7b80ff83b592e07186902f1ad793e98c2e4fda644b434cd8ce5ad1b5753 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmemberobjects"
    i438745a1dc6035d947cc9f37018ac163f7b2440754dbf4021981faec8e65ed2d "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i44321e047f2bffa6dedbcd81fbb202039651b96cf54221e970864695312bf8c8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/sendmail"
    i5da3c9b5d13937b807d7a73a61395e121e4310b4eb5a67ef3cadcee33b2be8b0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmembergroups"
    i60c6937f870f584727c6e5c0f941408a726c03782b7994ba3e8b9f4da1af4363 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/checkmemberobjects"
    i8d5eadb561a6d6575825101b2000c489a8903463013fed6757ab18961756672d "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/assignlicense"
    i935ab712cf0ac065f1a2230d67ec4d88f419853ee69995a2269207f034b38628 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/findmeetingtimes"
    i970432633ec594d6a3e4a0bd81eba9c95572deb0fc5ff84583aa5501b6848ddf "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/checkmembergroups"
    i975a586754ce75ade353e1dd4cbae6d55a340a3ff84f8c99021c0c21d8a5763a "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmailtips"
    i97c8d47ed59692a49a68bb5bce2312e7c6ef8adb2361e08234843a766b24b7ab "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/reprocesslicenseassignment"
    i9bd9aa1f1c14bf232f9c4be6a3e0b57c95d9e769f89eea4e8c6952de359b068b "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/translateexchangeids"
    i9e0cd25a5c9025760969e6373728a1c941e3b961f69a56a5eb2896f92f6406af "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/revokesigninsessions"
    ibba99e4ff027a7768152295e623d73b4c3f6c528d300afecd81bda19c6fa3c35 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmanagedapppolicies"
    ic484710ebdbf9ece00c455e7f05e48d51734007068f3f75620ce93c2f64df57c "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/changepassword"
    idb107a94d38b7af1781cb08bfebaa444649d8e1ee65934f2f83fca5808e58c8b "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/wipemanagedappregistrationsbydevicetag"
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
func (m *UserRequestBuilder) AssignLicense()(*i8d5eadb561a6d6575825101b2000c489a8903463013fed6757ab18961756672d.AssignLicenseRequestBuilder) {
    return i8d5eadb561a6d6575825101b2000c489a8903463013fed6757ab18961756672d.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*ic484710ebdbf9ece00c455e7f05e48d51734007068f3f75620ce93c2f64df57c.ChangePasswordRequestBuilder) {
    return ic484710ebdbf9ece00c455e7f05e48d51734007068f3f75620ce93c2f64df57c.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i970432633ec594d6a3e4a0bd81eba9c95572deb0fc5ff84583aa5501b6848ddf.CheckMemberGroupsRequestBuilder) {
    return i970432633ec594d6a3e4a0bd81eba9c95572deb0fc5ff84583aa5501b6848ddf.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i60c6937f870f584727c6e5c0f941408a726c03782b7994ba3e8b9f4da1af4363.CheckMemberObjectsRequestBuilder) {
    return i60c6937f870f584727c6e5c0f941408a726c03782b7994ba3e8b9f4da1af4363.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/memberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*i27530cf646a7a9a035e401f73f898080c9813017863ea6bd12c4f503bfd11616.ExportPersonalDataRequestBuilder) {
    return i27530cf646a7a9a035e401f73f898080c9813017863ea6bd12c4f503bfd11616.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i935ab712cf0ac065f1a2230d67ec4d88f419853ee69995a2269207f034b38628.FindMeetingTimesRequestBuilder) {
    return i935ab712cf0ac065f1a2230d67ec4d88f419853ee69995a2269207f034b38628.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i975a586754ce75ade353e1dd4cbae6d55a340a3ff84f8c99021c0c21d8a5763a.GetMailTipsRequestBuilder) {
    return i975a586754ce75ade353e1dd4cbae6d55a340a3ff84f8c99021c0c21d8a5763a.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i2070c445678bb477a0e6ece4fba1feede3e05505713b95a26e0f64fe0e523d1e.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i2070c445678bb477a0e6ece4fba1feede3e05505713b95a26e0f64fe0e523d1e.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*ibba99e4ff027a7768152295e623d73b4c3f6c528d300afecd81bda19c6fa3c35.GetManagedAppPoliciesRequestBuilder) {
    return ibba99e4ff027a7768152295e623d73b4c3f6c528d300afecd81bda19c6fa3c35.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i5da3c9b5d13937b807d7a73a61395e121e4310b4eb5a67ef3cadcee33b2be8b0.GetMemberGroupsRequestBuilder) {
    return i5da3c9b5d13937b807d7a73a61395e121e4310b4eb5a67ef3cadcee33b2be8b0.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i3f45a7b80ff83b592e07186902f1ad793e98c2e4fda644b434cd8ce5ad1b5753.GetMemberObjectsRequestBuilder) {
    return i3f45a7b80ff83b592e07186902f1ad793e98c2e4fda644b434cd8ce5ad1b5753.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i438745a1dc6035d947cc9f37018ac163f7b2440754dbf4021981faec8e65ed2d.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i438745a1dc6035d947cc9f37018ac163f7b2440754dbf4021981faec8e65ed2d.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i3980cf70ba7e8face0024e6f88ab6d601e20ef50a6dc80273b1f137db76e3b58.RemoveAllDevicesFromManagementRequestBuilder) {
    return i3980cf70ba7e8face0024e6f88ab6d601e20ef50a6dc80273b1f137db76e3b58.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i97c8d47ed59692a49a68bb5bce2312e7c6ef8adb2361e08234843a766b24b7ab.ReprocessLicenseAssignmentRequestBuilder) {
    return i97c8d47ed59692a49a68bb5bce2312e7c6ef8adb2361e08234843a766b24b7ab.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i226ecc447edd24870796d566cc7ccfe04b5f8a423bd59c81b8d3b01905cd5851.RestoreRequestBuilder) {
    return i226ecc447edd24870796d566cc7ccfe04b5f8a423bd59c81b8d3b01905cd5851.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i9e0cd25a5c9025760969e6373728a1c941e3b961f69a56a5eb2896f92f6406af.RevokeSignInSessionsRequestBuilder) {
    return i9e0cd25a5c9025760969e6373728a1c941e3b961f69a56a5eb2896f92f6406af.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i44321e047f2bffa6dedbcd81fbb202039651b96cf54221e970864695312bf8c8.SendMailRequestBuilder) {
    return i44321e047f2bffa6dedbcd81fbb202039651b96cf54221e970864695312bf8c8.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i9bd9aa1f1c14bf232f9c4be6a3e0b57c95d9e769f89eea4e8c6952de359b068b.TranslateExchangeIdsRequestBuilder) {
    return i9bd9aa1f1c14bf232f9c4be6a3e0b57c95d9e769f89eea4e8c6952de359b068b.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*idb107a94d38b7af1781cb08bfebaa444649d8e1ee65934f2f83fca5808e58c8b.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return idb107a94d38b7af1781cb08bfebaa444649d8e1ee65934f2f83fca5808e58c8b.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
