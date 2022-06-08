package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i002301f9155631a22e209658d43d84acabcddb22107bc9a2dd55d49821a614c3 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/wipemanagedappregistrationsbydevicetag"
    i005aeff1450145cd5a835ca47ea766eb7c4d048a7e6ec58d98ba4df00c8336cc "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/reprocesslicenseassignment"
    i01c6fe25fb18951a18b498c6dbb1eea99b821e1ba3a634694c23ada41c61cf4b "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmailtips"
    i03b97e53844bede4ad32b7aadb06cb0f8b8ac968519b3a5bf577d8c7d14bbc39 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/sendmail"
    i05ba256dcf1de9c2dcfe297257a77df96ba37682ae28e092062b7b61a408df1b "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmembergroups"
    i2a5b8ab1c7f79eee3484cf3df992da2815c342bf613efb756e323072bdc9e37b "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/removealldevicesfrommanagement"
    i38cbe0781b77a1237e92718b37522db39d89083acdb7a7d96631e2077961e55b "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/reminderviewwithstartdatetimewithenddatetime"
    i418856f2e73d5f5e8c4b67160d79e6dc3c3a7d860789d061eb7ca83522cf4032 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmanagedapppolicies"
    i4c4701d5296da1c741ac578b95f01e424961d5e85e4733772cd7f61bd4a17bb2 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/findmeetingtimes"
    i4ef6a5852699cd9ca1df1defea208a73abd68212f2c020b5f1a531950b4017ed "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/translateexchangeids"
    i510e0539d5089aefe5d124b70e57c5cc2773ef9b49d683a4d8df114bbf064e62 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/changepassword"
    i71425f22e324268f189c3e87edd3aa502ca2182c7105246ef0aa43ad8483dbe3 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmemberobjects"
    i8aa541794c105b8539e4b06d8bbe2ff25954563787965a65f3a2a7b78f5f3eb6 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/checkmembergroups"
    i964b7a2380bb88720427dbd19c41c75f35605f068e0cfa53d1cfeeec12a64afb "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/exportpersonaldata"
    ia6a778f89374d7e18aedba1bbe69a1719c3acf9d9ebdbcdad034fcc38fd53c18 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmanagedappdiagnosticstatuses"
    ic872f626c2d89679fe11215c017168c156de82cb7d2368395669388d8ecf6081 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/revokesigninsessions"
    id691d180df5be1d55e647ca80d420478a269a2ebebb628ee4586bdf711e2e4b9 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/checkmemberobjects"
    id863bd94f75c8cd6e3c47ded78881e3bfbd93ee5fd678c213ed412eb02341e7e "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/restore"
    if4f1a9469d7e6b4d1cea44ee19b036e848baef5707d6eba8de9871df1aae75a3 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/assignlicense"
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
func (m *UserRequestBuilder) AssignLicense()(*if4f1a9469d7e6b4d1cea44ee19b036e848baef5707d6eba8de9871df1aae75a3.AssignLicenseRequestBuilder) {
    return if4f1a9469d7e6b4d1cea44ee19b036e848baef5707d6eba8de9871df1aae75a3.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i510e0539d5089aefe5d124b70e57c5cc2773ef9b49d683a4d8df114bbf064e62.ChangePasswordRequestBuilder) {
    return i510e0539d5089aefe5d124b70e57c5cc2773ef9b49d683a4d8df114bbf064e62.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i8aa541794c105b8539e4b06d8bbe2ff25954563787965a65f3a2a7b78f5f3eb6.CheckMemberGroupsRequestBuilder) {
    return i8aa541794c105b8539e4b06d8bbe2ff25954563787965a65f3a2a7b78f5f3eb6.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*id691d180df5be1d55e647ca80d420478a269a2ebebb628ee4586bdf711e2e4b9.CheckMemberObjectsRequestBuilder) {
    return id691d180df5be1d55e647ca80d420478a269a2ebebb628ee4586bdf711e2e4b9.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*i964b7a2380bb88720427dbd19c41c75f35605f068e0cfa53d1cfeeec12a64afb.ExportPersonalDataRequestBuilder) {
    return i964b7a2380bb88720427dbd19c41c75f35605f068e0cfa53d1cfeeec12a64afb.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i4c4701d5296da1c741ac578b95f01e424961d5e85e4733772cd7f61bd4a17bb2.FindMeetingTimesRequestBuilder) {
    return i4c4701d5296da1c741ac578b95f01e424961d5e85e4733772cd7f61bd4a17bb2.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i01c6fe25fb18951a18b498c6dbb1eea99b821e1ba3a634694c23ada41c61cf4b.GetMailTipsRequestBuilder) {
    return i01c6fe25fb18951a18b498c6dbb1eea99b821e1ba3a634694c23ada41c61cf4b.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*ia6a778f89374d7e18aedba1bbe69a1719c3acf9d9ebdbcdad034fcc38fd53c18.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return ia6a778f89374d7e18aedba1bbe69a1719c3acf9d9ebdbcdad034fcc38fd53c18.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i418856f2e73d5f5e8c4b67160d79e6dc3c3a7d860789d061eb7ca83522cf4032.GetManagedAppPoliciesRequestBuilder) {
    return i418856f2e73d5f5e8c4b67160d79e6dc3c3a7d860789d061eb7ca83522cf4032.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i05ba256dcf1de9c2dcfe297257a77df96ba37682ae28e092062b7b61a408df1b.GetMemberGroupsRequestBuilder) {
    return i05ba256dcf1de9c2dcfe297257a77df96ba37682ae28e092062b7b61a408df1b.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i71425f22e324268f189c3e87edd3aa502ca2182c7105246ef0aa43ad8483dbe3.GetMemberObjectsRequestBuilder) {
    return i71425f22e324268f189c3e87edd3aa502ca2182c7105246ef0aa43ad8483dbe3.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i38cbe0781b77a1237e92718b37522db39d89083acdb7a7d96631e2077961e55b.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i38cbe0781b77a1237e92718b37522db39d89083acdb7a7d96631e2077961e55b.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i2a5b8ab1c7f79eee3484cf3df992da2815c342bf613efb756e323072bdc9e37b.RemoveAllDevicesFromManagementRequestBuilder) {
    return i2a5b8ab1c7f79eee3484cf3df992da2815c342bf613efb756e323072bdc9e37b.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i005aeff1450145cd5a835ca47ea766eb7c4d048a7e6ec58d98ba4df00c8336cc.ReprocessLicenseAssignmentRequestBuilder) {
    return i005aeff1450145cd5a835ca47ea766eb7c4d048a7e6ec58d98ba4df00c8336cc.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*id863bd94f75c8cd6e3c47ded78881e3bfbd93ee5fd678c213ed412eb02341e7e.RestoreRequestBuilder) {
    return id863bd94f75c8cd6e3c47ded78881e3bfbd93ee5fd678c213ed412eb02341e7e.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*ic872f626c2d89679fe11215c017168c156de82cb7d2368395669388d8ecf6081.RevokeSignInSessionsRequestBuilder) {
    return ic872f626c2d89679fe11215c017168c156de82cb7d2368395669388d8ecf6081.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i03b97e53844bede4ad32b7aadb06cb0f8b8ac968519b3a5bf577d8c7d14bbc39.SendMailRequestBuilder) {
    return i03b97e53844bede4ad32b7aadb06cb0f8b8ac968519b3a5bf577d8c7d14bbc39.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i4ef6a5852699cd9ca1df1defea208a73abd68212f2c020b5f1a531950b4017ed.TranslateExchangeIdsRequestBuilder) {
    return i4ef6a5852699cd9ca1df1defea208a73abd68212f2c020b5f1a531950b4017ed.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i002301f9155631a22e209658d43d84acabcddb22107bc9a2dd55d49821a614c3.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i002301f9155631a22e209658d43d84acabcddb22107bc9a2dd55d49821a614c3.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
