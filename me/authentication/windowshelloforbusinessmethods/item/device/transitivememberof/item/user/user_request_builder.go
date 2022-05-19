package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i168a6408fb68bdc563b6cc43f48b64bf35c74b775b16ffcb5fa504e78418a695 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/removealldevicesfrommanagement"
    i1b0f90d2526cf73c39178260adaace0bf2d5a8475a428ec9053051b515a8bcf0 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/exportpersonaldata"
    i3957f7a02a2a2335548d7d7e42ff8c2ca3e67d6d4f3931a9782e005345fb7a78 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/wipemanagedappregistrationsbydevicetag"
    i4b87d2663e0e099eb89412e370b7c6bd9f77cf50c7eaf1cf7000e6cd9f8b6fef "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/checkmemberobjects"
    i4f93b3a3fe42468e24365e13440a1fd62f2e3869b43955322a21ac6b275f314d "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/assignlicense"
    i55e9454f9145bf4acbb8f84cc8afd68eb8a88bc12a4ec1da1279d8fab8b78cae "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmailtips"
    i5c2ae46db87fba76405ae00c7ad40d21f7bee7f4ab52dd506cd97c6a84e739c3 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmemberobjects"
    i65f28eb1d8bfb195fd3e8a399fe56e1495e03049e7d06ec957cebecab8da9eaf "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmanagedappdiagnosticstatuses"
    i6680f40e185ef002adf51e6ad9e3737227cc42ad1b1333ae9a1bb0e82ba1f1d1 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmembergroups"
    i9987d80f7cb242491c0b319e38bdd3718ebcbf1bec0089ea61859b1b8047073d "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/reprocesslicenseassignment"
    i9a8eedc78e8d355af0082e15f59b3beda7812da45e024d9d445c8ffc541de4f6 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/changepassword"
    ia67c2d8f3634952b867678538e5b75fc4677ff31f12c980a2064ece8ffb16ad1 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmanagedapppolicies"
    ia7133f3d869b42eb7496aa3a3a17c3af3560e6919f5c692b5873ef1d4e021d28 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/restore"
    ib409f27019eab1067a2c3210bb8b29bad478ee8a561f9e935510c1e099276b5b "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/checkmembergroups"
    icd814027060116b65c27a376e73487e1cae32b533350d555933e77f11d5e673d "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    ie60a472aabc7b8b68e9313e8699bc42fee081c146cb8149c2f559bf7ea58690f "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/findmeetingtimes"
    ief4623d813bbd7cc1b19765cd475a1b2e9188d71fb9f40bf88cd53ece15ad49e "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/translateexchangeids"
    if4ededd4467fd49d8e347ee5706006fd526090102cf360b04256fc34fac0349b "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/revokesigninsessions"
    ifc46f53cfafe1f93f33fa6d2cb7831110ec144869089aa2cfb5b62f11d1a00b9 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/sendmail"
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
func (m *UserRequestBuilder) AssignLicense()(*i4f93b3a3fe42468e24365e13440a1fd62f2e3869b43955322a21ac6b275f314d.AssignLicenseRequestBuilder) {
    return i4f93b3a3fe42468e24365e13440a1fd62f2e3869b43955322a21ac6b275f314d.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i9a8eedc78e8d355af0082e15f59b3beda7812da45e024d9d445c8ffc541de4f6.ChangePasswordRequestBuilder) {
    return i9a8eedc78e8d355af0082e15f59b3beda7812da45e024d9d445c8ffc541de4f6.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*ib409f27019eab1067a2c3210bb8b29bad478ee8a561f9e935510c1e099276b5b.CheckMemberGroupsRequestBuilder) {
    return ib409f27019eab1067a2c3210bb8b29bad478ee8a561f9e935510c1e099276b5b.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i4b87d2663e0e099eb89412e370b7c6bd9f77cf50c7eaf1cf7000e6cd9f8b6fef.CheckMemberObjectsRequestBuilder) {
    return i4b87d2663e0e099eb89412e370b7c6bd9f77cf50c7eaf1cf7000e6cd9f8b6fef.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*i1b0f90d2526cf73c39178260adaace0bf2d5a8475a428ec9053051b515a8bcf0.ExportPersonalDataRequestBuilder) {
    return i1b0f90d2526cf73c39178260adaace0bf2d5a8475a428ec9053051b515a8bcf0.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*ie60a472aabc7b8b68e9313e8699bc42fee081c146cb8149c2f559bf7ea58690f.FindMeetingTimesRequestBuilder) {
    return ie60a472aabc7b8b68e9313e8699bc42fee081c146cb8149c2f559bf7ea58690f.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i55e9454f9145bf4acbb8f84cc8afd68eb8a88bc12a4ec1da1279d8fab8b78cae.GetMailTipsRequestBuilder) {
    return i55e9454f9145bf4acbb8f84cc8afd68eb8a88bc12a4ec1da1279d8fab8b78cae.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i65f28eb1d8bfb195fd3e8a399fe56e1495e03049e7d06ec957cebecab8da9eaf.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i65f28eb1d8bfb195fd3e8a399fe56e1495e03049e7d06ec957cebecab8da9eaf.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*ia67c2d8f3634952b867678538e5b75fc4677ff31f12c980a2064ece8ffb16ad1.GetManagedAppPoliciesRequestBuilder) {
    return ia67c2d8f3634952b867678538e5b75fc4677ff31f12c980a2064ece8ffb16ad1.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i6680f40e185ef002adf51e6ad9e3737227cc42ad1b1333ae9a1bb0e82ba1f1d1.GetMemberGroupsRequestBuilder) {
    return i6680f40e185ef002adf51e6ad9e3737227cc42ad1b1333ae9a1bb0e82ba1f1d1.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i5c2ae46db87fba76405ae00c7ad40d21f7bee7f4ab52dd506cd97c6a84e739c3.GetMemberObjectsRequestBuilder) {
    return i5c2ae46db87fba76405ae00c7ad40d21f7bee7f4ab52dd506cd97c6a84e739c3.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*icd814027060116b65c27a376e73487e1cae32b533350d555933e77f11d5e673d.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return icd814027060116b65c27a376e73487e1cae32b533350d555933e77f11d5e673d.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i168a6408fb68bdc563b6cc43f48b64bf35c74b775b16ffcb5fa504e78418a695.RemoveAllDevicesFromManagementRequestBuilder) {
    return i168a6408fb68bdc563b6cc43f48b64bf35c74b775b16ffcb5fa504e78418a695.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i9987d80f7cb242491c0b319e38bdd3718ebcbf1bec0089ea61859b1b8047073d.ReprocessLicenseAssignmentRequestBuilder) {
    return i9987d80f7cb242491c0b319e38bdd3718ebcbf1bec0089ea61859b1b8047073d.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*ia7133f3d869b42eb7496aa3a3a17c3af3560e6919f5c692b5873ef1d4e021d28.RestoreRequestBuilder) {
    return ia7133f3d869b42eb7496aa3a3a17c3af3560e6919f5c692b5873ef1d4e021d28.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*if4ededd4467fd49d8e347ee5706006fd526090102cf360b04256fc34fac0349b.RevokeSignInSessionsRequestBuilder) {
    return if4ededd4467fd49d8e347ee5706006fd526090102cf360b04256fc34fac0349b.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*ifc46f53cfafe1f93f33fa6d2cb7831110ec144869089aa2cfb5b62f11d1a00b9.SendMailRequestBuilder) {
    return ifc46f53cfafe1f93f33fa6d2cb7831110ec144869089aa2cfb5b62f11d1a00b9.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*ief4623d813bbd7cc1b19765cd475a1b2e9188d71fb9f40bf88cd53ece15ad49e.TranslateExchangeIdsRequestBuilder) {
    return ief4623d813bbd7cc1b19765cd475a1b2e9188d71fb9f40bf88cd53ece15ad49e.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i3957f7a02a2a2335548d7d7e42ff8c2ca3e67d6d4f3931a9782e005345fb7a78.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i3957f7a02a2a2335548d7d7e42ff8c2ca3e67d6d4f3931a9782e005345fb7a78.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
