package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i02a04b6e43081491c2abcb830b305cfc3ab4dc9d4259a002176c3c2eb3740686 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmanagedapppolicies"
    i0905e0ac4dfce182193508f037473078789db3a50a06a30286a2ee5948722c2b "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/sendmail"
    i09356101381ad72f71353e79635a568c15aa4784845da2174da11ecc16e9f22b "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i1222d7b752526ff89b69d30a38a102e3057bebc68233ecc507d1b60a089dbd6a "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/assignlicense"
    i46776dd0f100b41fc690e244132d19b63e39e3187f349fe7dc9e5175a0847766 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/checkmemberobjects"
    i64851b465f265f3d2ba47c84d63e5649caf3419fccfc349dbca432923ca5c0be "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/removealldevicesfrommanagement"
    i69c7e1fb162d17de7eacd3f2a04d77fa3afa733ba6c0c4ffbd0da6a67d5ba782 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/checkmembergroups"
    i6c491168c6efd75266ea68626bded35484615999f37c612b5a4afe88923c080b "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/translateexchangeids"
    i8868dcff164c1492b1cfa63c550117eadc50f0b66ca7bb8b112e21310da38b47 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/exportpersonaldata"
    i8fb8b409549290207400b5a2138df19314f12b29a7770ae1dbe294b75c145d76 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmembergroups"
    i9107af457a49fa9c687760d70096284aefcbdb0fdc7f6551045cb2eee425faa4 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/wipemanagedappregistrationsbydevicetag"
    i91c34a330a3d91cd397e97d5a823a9a8986c73008d35049981674742827acf6d "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmemberobjects"
    i9a811f9770e0662b1131dbb8ee538867cda62d634c0997b985183d6a0ffc51ba "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/reprocesslicenseassignment"
    idcd4143c7fa6d67cfc1960110ca8161cb83afc44a5a341bea6649b64b26b9187 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmanagedappdiagnosticstatuses"
    ie2e6034515a33f1a0a2aac0319c7f107f2d7bb31fe97ad710c01c0c871f01878 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/revokesigninsessions"
    ie699760d92d96f2a1f6f7924f9e0f982bc8fedf6b3b9aa28cc2ca67dc6c8c0a3 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmailtips"
    ieab0a84ba22d2a97b2a2853a9c6130a8bb03c06119f68d792d16d68741378952 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/restore"
    if01aa1f6c7d43938207fc7d11f106de560219b5891bd5c58e05f874e49b5da45 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/findmeetingtimes"
    iff6bbdc35918a4d33b3eb6d3d9e45db7b95ec047eaed56a5e59e781e21610c63 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/changepassword"
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
func (m *UserRequestBuilder) AssignLicense()(*i1222d7b752526ff89b69d30a38a102e3057bebc68233ecc507d1b60a089dbd6a.AssignLicenseRequestBuilder) {
    return i1222d7b752526ff89b69d30a38a102e3057bebc68233ecc507d1b60a089dbd6a.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*iff6bbdc35918a4d33b3eb6d3d9e45db7b95ec047eaed56a5e59e781e21610c63.ChangePasswordRequestBuilder) {
    return iff6bbdc35918a4d33b3eb6d3d9e45db7b95ec047eaed56a5e59e781e21610c63.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i69c7e1fb162d17de7eacd3f2a04d77fa3afa733ba6c0c4ffbd0da6a67d5ba782.CheckMemberGroupsRequestBuilder) {
    return i69c7e1fb162d17de7eacd3f2a04d77fa3afa733ba6c0c4ffbd0da6a67d5ba782.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i46776dd0f100b41fc690e244132d19b63e39e3187f349fe7dc9e5175a0847766.CheckMemberObjectsRequestBuilder) {
    return i46776dd0f100b41fc690e244132d19b63e39e3187f349fe7dc9e5175a0847766.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/memberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*i8868dcff164c1492b1cfa63c550117eadc50f0b66ca7bb8b112e21310da38b47.ExportPersonalDataRequestBuilder) {
    return i8868dcff164c1492b1cfa63c550117eadc50f0b66ca7bb8b112e21310da38b47.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*if01aa1f6c7d43938207fc7d11f106de560219b5891bd5c58e05f874e49b5da45.FindMeetingTimesRequestBuilder) {
    return if01aa1f6c7d43938207fc7d11f106de560219b5891bd5c58e05f874e49b5da45.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*ie699760d92d96f2a1f6f7924f9e0f982bc8fedf6b3b9aa28cc2ca67dc6c8c0a3.GetMailTipsRequestBuilder) {
    return ie699760d92d96f2a1f6f7924f9e0f982bc8fedf6b3b9aa28cc2ca67dc6c8c0a3.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*idcd4143c7fa6d67cfc1960110ca8161cb83afc44a5a341bea6649b64b26b9187.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return idcd4143c7fa6d67cfc1960110ca8161cb83afc44a5a341bea6649b64b26b9187.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i02a04b6e43081491c2abcb830b305cfc3ab4dc9d4259a002176c3c2eb3740686.GetManagedAppPoliciesRequestBuilder) {
    return i02a04b6e43081491c2abcb830b305cfc3ab4dc9d4259a002176c3c2eb3740686.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i8fb8b409549290207400b5a2138df19314f12b29a7770ae1dbe294b75c145d76.GetMemberGroupsRequestBuilder) {
    return i8fb8b409549290207400b5a2138df19314f12b29a7770ae1dbe294b75c145d76.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i91c34a330a3d91cd397e97d5a823a9a8986c73008d35049981674742827acf6d.GetMemberObjectsRequestBuilder) {
    return i91c34a330a3d91cd397e97d5a823a9a8986c73008d35049981674742827acf6d.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i09356101381ad72f71353e79635a568c15aa4784845da2174da11ecc16e9f22b.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i09356101381ad72f71353e79635a568c15aa4784845da2174da11ecc16e9f22b.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i64851b465f265f3d2ba47c84d63e5649caf3419fccfc349dbca432923ca5c0be.RemoveAllDevicesFromManagementRequestBuilder) {
    return i64851b465f265f3d2ba47c84d63e5649caf3419fccfc349dbca432923ca5c0be.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i9a811f9770e0662b1131dbb8ee538867cda62d634c0997b985183d6a0ffc51ba.ReprocessLicenseAssignmentRequestBuilder) {
    return i9a811f9770e0662b1131dbb8ee538867cda62d634c0997b985183d6a0ffc51ba.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*ieab0a84ba22d2a97b2a2853a9c6130a8bb03c06119f68d792d16d68741378952.RestoreRequestBuilder) {
    return ieab0a84ba22d2a97b2a2853a9c6130a8bb03c06119f68d792d16d68741378952.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*ie2e6034515a33f1a0a2aac0319c7f107f2d7bb31fe97ad710c01c0c871f01878.RevokeSignInSessionsRequestBuilder) {
    return ie2e6034515a33f1a0a2aac0319c7f107f2d7bb31fe97ad710c01c0c871f01878.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i0905e0ac4dfce182193508f037473078789db3a50a06a30286a2ee5948722c2b.SendMailRequestBuilder) {
    return i0905e0ac4dfce182193508f037473078789db3a50a06a30286a2ee5948722c2b.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i6c491168c6efd75266ea68626bded35484615999f37c612b5a4afe88923c080b.TranslateExchangeIdsRequestBuilder) {
    return i6c491168c6efd75266ea68626bded35484615999f37c612b5a4afe88923c080b.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i9107af457a49fa9c687760d70096284aefcbdb0fdc7f6551045cb2eee425faa4.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i9107af457a49fa9c687760d70096284aefcbdb0fdc7f6551045cb2eee425faa4.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
