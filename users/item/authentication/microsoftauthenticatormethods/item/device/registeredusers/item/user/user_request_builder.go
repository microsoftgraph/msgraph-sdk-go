package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i094c94a330b1acc96a82dc1ec2ccde86db88e6077f9cc5a231b2ae80ffd67e7d "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/reprocesslicenseassignment"
    i1fdd1fd3e73d5c58198e2d1bb1a652ffb2d4857c666c90aedc9b2c6455407070 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/translateexchangeids"
    i2f6a472d09940df120988615ea9526c04dbcb2add48ce2bfd98179da866e84b9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/findmeetingtimes"
    i33aea86d0b18ffc58f21a7b1f1a077c95eb520f4421e53f637d3401b1ea23e19 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmemberobjects"
    i358a53d0b81a626348851273d83def11ca89bba08ed53ae0a4b65ce039d2b918 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/revokesigninsessions"
    i3d5885b70374de0863101d7d787aeb8b424950f07b0b37d1f5f7dab2c8395252 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/removealldevicesfrommanagement"
    i446e072c70268491dac834d3679063a3a50829cdb6014def5c494d61b59d4435 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/checkmembergroups"
    i4a757c4701e155e681ee8c89aee54dcba734cff0f9c70f88f0d534718fa0ee88 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmembergroups"
    i54db0918f94c75cc28df6f3081743d27a837d2a1f8b2fa90559f05023ec7084b "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/exportpersonaldata"
    i60e5b0145dbce1b193009e18caca5321e7752e3183570c6c3c7b2b602c2af78b "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/assignlicense"
    i75c542414f8fae790418b3fa64acfe0258633426d0ccee37b99dba398e0e61d3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/reminderviewwithstartdatetimewithenddatetime"
    i861c6a924bb5d11191be318c0137904ba1497b19ccd1910d61706c34359a3ad5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/restore"
    i879963576cebb1f6c256eb9e3a67fd6a4d7dc95b34a0410d5b761bc96e5014ec "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/sendmail"
    i90fcbc1acdff37470f3c985048a0ded17ece2d2d65b9b7cd006fc3c043cef287 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/changepassword"
    i9d1e33a3032dccb53e48b922e49269d1a3eded57731872bb43365efe8ec784de "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/checkmemberobjects"
    ia674383743f82f3a933de1cdc54955623e7be2fc5657abe8dab79d03696d5c5f "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmanagedapppolicies"
    iccbb0ad360ea42c7371d4ef2f368d3c31807225494b77d02d828c7e41a604228 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmanagedappdiagnosticstatuses"
    id4ba95d18790aca59a956f736a10ee4b31692fa49cbc646825e53fa27d823daf "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmailtips"
    if37c0d765b4296445b40044d157cd43b55add47ffcd243a09ab1a7c5b0f1e0ee "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/wipemanagedappregistrationsbydevicetag"
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
func (m *UserRequestBuilder) AssignLicense()(*i60e5b0145dbce1b193009e18caca5321e7752e3183570c6c3c7b2b602c2af78b.AssignLicenseRequestBuilder) {
    return i60e5b0145dbce1b193009e18caca5321e7752e3183570c6c3c7b2b602c2af78b.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i90fcbc1acdff37470f3c985048a0ded17ece2d2d65b9b7cd006fc3c043cef287.ChangePasswordRequestBuilder) {
    return i90fcbc1acdff37470f3c985048a0ded17ece2d2d65b9b7cd006fc3c043cef287.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i446e072c70268491dac834d3679063a3a50829cdb6014def5c494d61b59d4435.CheckMemberGroupsRequestBuilder) {
    return i446e072c70268491dac834d3679063a3a50829cdb6014def5c494d61b59d4435.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i9d1e33a3032dccb53e48b922e49269d1a3eded57731872bb43365efe8ec784de.CheckMemberObjectsRequestBuilder) {
    return i9d1e33a3032dccb53e48b922e49269d1a3eded57731872bb43365efe8ec784de.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredUsers/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*i54db0918f94c75cc28df6f3081743d27a837d2a1f8b2fa90559f05023ec7084b.ExportPersonalDataRequestBuilder) {
    return i54db0918f94c75cc28df6f3081743d27a837d2a1f8b2fa90559f05023ec7084b.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i2f6a472d09940df120988615ea9526c04dbcb2add48ce2bfd98179da866e84b9.FindMeetingTimesRequestBuilder) {
    return i2f6a472d09940df120988615ea9526c04dbcb2add48ce2bfd98179da866e84b9.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*id4ba95d18790aca59a956f736a10ee4b31692fa49cbc646825e53fa27d823daf.GetMailTipsRequestBuilder) {
    return id4ba95d18790aca59a956f736a10ee4b31692fa49cbc646825e53fa27d823daf.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*iccbb0ad360ea42c7371d4ef2f368d3c31807225494b77d02d828c7e41a604228.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return iccbb0ad360ea42c7371d4ef2f368d3c31807225494b77d02d828c7e41a604228.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*ia674383743f82f3a933de1cdc54955623e7be2fc5657abe8dab79d03696d5c5f.GetManagedAppPoliciesRequestBuilder) {
    return ia674383743f82f3a933de1cdc54955623e7be2fc5657abe8dab79d03696d5c5f.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i4a757c4701e155e681ee8c89aee54dcba734cff0f9c70f88f0d534718fa0ee88.GetMemberGroupsRequestBuilder) {
    return i4a757c4701e155e681ee8c89aee54dcba734cff0f9c70f88f0d534718fa0ee88.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i33aea86d0b18ffc58f21a7b1f1a077c95eb520f4421e53f637d3401b1ea23e19.GetMemberObjectsRequestBuilder) {
    return i33aea86d0b18ffc58f21a7b1f1a077c95eb520f4421e53f637d3401b1ea23e19.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i75c542414f8fae790418b3fa64acfe0258633426d0ccee37b99dba398e0e61d3.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i75c542414f8fae790418b3fa64acfe0258633426d0ccee37b99dba398e0e61d3.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i3d5885b70374de0863101d7d787aeb8b424950f07b0b37d1f5f7dab2c8395252.RemoveAllDevicesFromManagementRequestBuilder) {
    return i3d5885b70374de0863101d7d787aeb8b424950f07b0b37d1f5f7dab2c8395252.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i094c94a330b1acc96a82dc1ec2ccde86db88e6077f9cc5a231b2ae80ffd67e7d.ReprocessLicenseAssignmentRequestBuilder) {
    return i094c94a330b1acc96a82dc1ec2ccde86db88e6077f9cc5a231b2ae80ffd67e7d.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i861c6a924bb5d11191be318c0137904ba1497b19ccd1910d61706c34359a3ad5.RestoreRequestBuilder) {
    return i861c6a924bb5d11191be318c0137904ba1497b19ccd1910d61706c34359a3ad5.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i358a53d0b81a626348851273d83def11ca89bba08ed53ae0a4b65ce039d2b918.RevokeSignInSessionsRequestBuilder) {
    return i358a53d0b81a626348851273d83def11ca89bba08ed53ae0a4b65ce039d2b918.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i879963576cebb1f6c256eb9e3a67fd6a4d7dc95b34a0410d5b761bc96e5014ec.SendMailRequestBuilder) {
    return i879963576cebb1f6c256eb9e3a67fd6a4d7dc95b34a0410d5b761bc96e5014ec.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i1fdd1fd3e73d5c58198e2d1bb1a652ffb2d4857c666c90aedc9b2c6455407070.TranslateExchangeIdsRequestBuilder) {
    return i1fdd1fd3e73d5c58198e2d1bb1a652ffb2d4857c666c90aedc9b2c6455407070.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*if37c0d765b4296445b40044d157cd43b55add47ffcd243a09ab1a7c5b0f1e0ee.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return if37c0d765b4296445b40044d157cd43b55add47ffcd243a09ab1a7c5b0f1e0ee.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
