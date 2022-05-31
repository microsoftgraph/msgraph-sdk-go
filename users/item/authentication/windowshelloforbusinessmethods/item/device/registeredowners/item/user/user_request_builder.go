package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i2b94965342c590caca486657cb9ae971945216ef38c7391a32f4068d4047d27b "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/wipemanagedappregistrationsbydevicetag"
    i4e16fb53b327d9d6bc6f51d69dd6ecd43c81744974f33e0e881da079b44a2603 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/exportpersonaldata"
    i502f01d12445945c67e18298a0851f9e333fcc37bfe57b5f95c3968052b78169 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/removealldevicesfrommanagement"
    i679143ff0efa5b93f0d6014ab76dfa0679de3afd65fc264f70d9f1f27e057d14 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmanagedappdiagnosticstatuses"
    i77af6a10d58a289e60faf752e61d7681f35e4c15e1f9e7062e8532c705336f45 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/restore"
    i7e8258131a5816993e47983cbbdebb422496e823ff36f1875f22d753c2d906c3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmembergroups"
    i9629fe814e9ea1d2c7a65c8841464cbf1742cc2e1a44c2e94850a1e1dc50020e "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/translateexchangeids"
    ia0184f6568ba9a50998c5182b0733c657e92dd30a3f55b542109b2ece3a8adfc "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/checkmembergroups"
    iaa962c7bcf9e30a48f6d16b9abef32f63a1a24515901deead6aee9fa6854a38d "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmemberobjects"
    iae311b9db275d0dd122c27a0d110f37e5a89e1456a1d9724d79d3516d339b8eb "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/assignlicense"
    iafda0a277d54e6ce1d590a71f76efc6d728e0c48b9e25e1b33e2b6b64f27487e "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/checkmemberobjects"
    ibb2a26056340ed5e95865e4128eb09118237dbcea1edf7797369c12b3a48459c "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/reprocesslicenseassignment"
    ic80f0bf9ee630f0cdab09e3edb1bfa43211fb6670472a1c60fed48094f237728 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmanagedapppolicies"
    icd092446d33825167d472222381d048995338f3b0c4c1c33cc3ac66dbe9a33e0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/revokesigninsessions"
    idabd528318e9212e0618f90e19937ce79437fc419db22ea16751c6a2c18869f3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmailtips"
    ie6bb51cc96bf1307b0d1ee8b32001597ad7baa8b2719c4fb4299cbec89749019 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/reminderviewwithstartdatetimewithenddatetime"
    ie814dd556b3a52ed2c93451527233f9ede1715166cdb8e34e0ca6406a62c058a "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/sendmail"
    ie948546b329dbb1f3a4bd90da5e1f59944cabf7a094656c09d7437959cdc4320 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/findmeetingtimes"
    ifa48dc2676e69043ea49cb9a02abb64cfed09279d415ab573b50a860595d1373 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/changepassword"
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
func (m *UserRequestBuilder) AssignLicense()(*iae311b9db275d0dd122c27a0d110f37e5a89e1456a1d9724d79d3516d339b8eb.AssignLicenseRequestBuilder) {
    return iae311b9db275d0dd122c27a0d110f37e5a89e1456a1d9724d79d3516d339b8eb.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*ifa48dc2676e69043ea49cb9a02abb64cfed09279d415ab573b50a860595d1373.ChangePasswordRequestBuilder) {
    return ifa48dc2676e69043ea49cb9a02abb64cfed09279d415ab573b50a860595d1373.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*ia0184f6568ba9a50998c5182b0733c657e92dd30a3f55b542109b2ece3a8adfc.CheckMemberGroupsRequestBuilder) {
    return ia0184f6568ba9a50998c5182b0733c657e92dd30a3f55b542109b2ece3a8adfc.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*iafda0a277d54e6ce1d590a71f76efc6d728e0c48b9e25e1b33e2b6b64f27487e.CheckMemberObjectsRequestBuilder) {
    return iafda0a277d54e6ce1d590a71f76efc6d728e0c48b9e25e1b33e2b6b64f27487e.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*i4e16fb53b327d9d6bc6f51d69dd6ecd43c81744974f33e0e881da079b44a2603.ExportPersonalDataRequestBuilder) {
    return i4e16fb53b327d9d6bc6f51d69dd6ecd43c81744974f33e0e881da079b44a2603.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*ie948546b329dbb1f3a4bd90da5e1f59944cabf7a094656c09d7437959cdc4320.FindMeetingTimesRequestBuilder) {
    return ie948546b329dbb1f3a4bd90da5e1f59944cabf7a094656c09d7437959cdc4320.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*idabd528318e9212e0618f90e19937ce79437fc419db22ea16751c6a2c18869f3.GetMailTipsRequestBuilder) {
    return idabd528318e9212e0618f90e19937ce79437fc419db22ea16751c6a2c18869f3.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i679143ff0efa5b93f0d6014ab76dfa0679de3afd65fc264f70d9f1f27e057d14.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i679143ff0efa5b93f0d6014ab76dfa0679de3afd65fc264f70d9f1f27e057d14.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*ic80f0bf9ee630f0cdab09e3edb1bfa43211fb6670472a1c60fed48094f237728.GetManagedAppPoliciesRequestBuilder) {
    return ic80f0bf9ee630f0cdab09e3edb1bfa43211fb6670472a1c60fed48094f237728.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i7e8258131a5816993e47983cbbdebb422496e823ff36f1875f22d753c2d906c3.GetMemberGroupsRequestBuilder) {
    return i7e8258131a5816993e47983cbbdebb422496e823ff36f1875f22d753c2d906c3.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*iaa962c7bcf9e30a48f6d16b9abef32f63a1a24515901deead6aee9fa6854a38d.GetMemberObjectsRequestBuilder) {
    return iaa962c7bcf9e30a48f6d16b9abef32f63a1a24515901deead6aee9fa6854a38d.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*ie6bb51cc96bf1307b0d1ee8b32001597ad7baa8b2719c4fb4299cbec89749019.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return ie6bb51cc96bf1307b0d1ee8b32001597ad7baa8b2719c4fb4299cbec89749019.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i502f01d12445945c67e18298a0851f9e333fcc37bfe57b5f95c3968052b78169.RemoveAllDevicesFromManagementRequestBuilder) {
    return i502f01d12445945c67e18298a0851f9e333fcc37bfe57b5f95c3968052b78169.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*ibb2a26056340ed5e95865e4128eb09118237dbcea1edf7797369c12b3a48459c.ReprocessLicenseAssignmentRequestBuilder) {
    return ibb2a26056340ed5e95865e4128eb09118237dbcea1edf7797369c12b3a48459c.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i77af6a10d58a289e60faf752e61d7681f35e4c15e1f9e7062e8532c705336f45.RestoreRequestBuilder) {
    return i77af6a10d58a289e60faf752e61d7681f35e4c15e1f9e7062e8532c705336f45.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*icd092446d33825167d472222381d048995338f3b0c4c1c33cc3ac66dbe9a33e0.RevokeSignInSessionsRequestBuilder) {
    return icd092446d33825167d472222381d048995338f3b0c4c1c33cc3ac66dbe9a33e0.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*ie814dd556b3a52ed2c93451527233f9ede1715166cdb8e34e0ca6406a62c058a.SendMailRequestBuilder) {
    return ie814dd556b3a52ed2c93451527233f9ede1715166cdb8e34e0ca6406a62c058a.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i9629fe814e9ea1d2c7a65c8841464cbf1742cc2e1a44c2e94850a1e1dc50020e.TranslateExchangeIdsRequestBuilder) {
    return i9629fe814e9ea1d2c7a65c8841464cbf1742cc2e1a44c2e94850a1e1dc50020e.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i2b94965342c590caca486657cb9ae971945216ef38c7391a32f4068d4047d27b.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i2b94965342c590caca486657cb9ae971945216ef38c7391a32f4068d4047d27b.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
