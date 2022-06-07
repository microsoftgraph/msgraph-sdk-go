package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i00d8448b6b1607e1447f6c5706990b1bc956db81fb6844665dd53599a404cf6d "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/exportpersonaldata"
    i10e2e0def10a553738c2bd48c141a3bf7c50763a0d2c7643b877f9d73042142c "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/getmemberobjects"
    i16406dce7ef93a529c1afe49498737386718d97ebb78af4f79a038fc550b5bd9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/findmeetingtimes"
    i23014d2087b70772b0cccb9e9d6b3b0679f8e992b007818d9c2f2f85cd1818db "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/getmailtips"
    i2955873b1aec44e68d1e7d5746b62bb191c1b8e249c2d5fe8275ee2236bd388f "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/changepassword"
    i305186a396818b9b5c143c7103156f0f04d161300654a19dcdc3c7addf6a1096 "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/restore"
    i67c5d581f202527ad30e56a2c1cb97b1f60dc123ae810acf2159cf64e4c3be08 "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/checkmemberobjects"
    i849945e889839d680979e0c957d337ea2e4846d5b05086bec39f18e1ea478935 "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/getmanagedappdiagnosticstatuses"
    i867f2389b042d7aeb1318c9ee1fe640e5a38a5538c53de7da7f581296f833658 "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/reprocesslicenseassignment"
    i8c00256649d220a27a75bb77d0dc41512451130ad6e35775617350a0771bbce5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/getmembergroups"
    i8ca30bb0899ff38d6bd88d2b7106ac4c5c296c196632715417a38f5a08104192 "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/translateexchangeids"
    i9c40fb04136abdb492acf0bfe828594e9a8e42153a595ccb4f6f9a7ce7cb059e "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/revokesigninsessions"
    iad6e7cf1c646d873b1d50ca39233ee024e44e57a666e6071cb9f209c2e2b26f7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/wipemanagedappregistrationsbydevicetag"
    ic29e73702e7f404cb7321eee37b68fdd662427bee0f0086adcdc8733fc46d87a "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/reminderviewwithstartdatetimewithenddatetime"
    icb3924bea97cd4c9c2194bd608efa6b240cbd2809d6c609edb3648e3c3d847fd "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/assignlicense"
    id86431f1fc4da624b5b9f4d98e103edb42221a22e19dd878702fe4bc6e07e803 "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/removealldevicesfrommanagement"
    idf42cdd272eb67df0af6d3b3eefa0c675b13e3ac94270752582b15f0ed484473 "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/checkmembergroups"
    idf567b32cd236785f2bc078829a07357ad80d4ae6bb4d371a9d944391faf9172 "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/sendmail"
    if0cf9929ae20abe7ea032922fcc2a245a97db023d06dc05ffdeff5753c6ac445 "github.com/microsoftgraph/msgraph-sdk-go/users/item/directreports/item/user/getmanagedapppolicies"
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
func (m *UserRequestBuilder) AssignLicense()(*icb3924bea97cd4c9c2194bd608efa6b240cbd2809d6c609edb3648e3c3d847fd.AssignLicenseRequestBuilder) {
    return icb3924bea97cd4c9c2194bd608efa6b240cbd2809d6c609edb3648e3c3d847fd.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i2955873b1aec44e68d1e7d5746b62bb191c1b8e249c2d5fe8275ee2236bd388f.ChangePasswordRequestBuilder) {
    return i2955873b1aec44e68d1e7d5746b62bb191c1b8e249c2d5fe8275ee2236bd388f.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*idf42cdd272eb67df0af6d3b3eefa0c675b13e3ac94270752582b15f0ed484473.CheckMemberGroupsRequestBuilder) {
    return idf42cdd272eb67df0af6d3b3eefa0c675b13e3ac94270752582b15f0ed484473.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i67c5d581f202527ad30e56a2c1cb97b1f60dc123ae810acf2159cf64e4c3be08.CheckMemberObjectsRequestBuilder) {
    return i67c5d581f202527ad30e56a2c1cb97b1f60dc123ae810acf2159cf64e4c3be08.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/directReports/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*i00d8448b6b1607e1447f6c5706990b1bc956db81fb6844665dd53599a404cf6d.ExportPersonalDataRequestBuilder) {
    return i00d8448b6b1607e1447f6c5706990b1bc956db81fb6844665dd53599a404cf6d.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i16406dce7ef93a529c1afe49498737386718d97ebb78af4f79a038fc550b5bd9.FindMeetingTimesRequestBuilder) {
    return i16406dce7ef93a529c1afe49498737386718d97ebb78af4f79a038fc550b5bd9.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i23014d2087b70772b0cccb9e9d6b3b0679f8e992b007818d9c2f2f85cd1818db.GetMailTipsRequestBuilder) {
    return i23014d2087b70772b0cccb9e9d6b3b0679f8e992b007818d9c2f2f85cd1818db.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i849945e889839d680979e0c957d337ea2e4846d5b05086bec39f18e1ea478935.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i849945e889839d680979e0c957d337ea2e4846d5b05086bec39f18e1ea478935.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*if0cf9929ae20abe7ea032922fcc2a245a97db023d06dc05ffdeff5753c6ac445.GetManagedAppPoliciesRequestBuilder) {
    return if0cf9929ae20abe7ea032922fcc2a245a97db023d06dc05ffdeff5753c6ac445.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i8c00256649d220a27a75bb77d0dc41512451130ad6e35775617350a0771bbce5.GetMemberGroupsRequestBuilder) {
    return i8c00256649d220a27a75bb77d0dc41512451130ad6e35775617350a0771bbce5.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i10e2e0def10a553738c2bd48c141a3bf7c50763a0d2c7643b877f9d73042142c.GetMemberObjectsRequestBuilder) {
    return i10e2e0def10a553738c2bd48c141a3bf7c50763a0d2c7643b877f9d73042142c.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*ic29e73702e7f404cb7321eee37b68fdd662427bee0f0086adcdc8733fc46d87a.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return ic29e73702e7f404cb7321eee37b68fdd662427bee0f0086adcdc8733fc46d87a.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*id86431f1fc4da624b5b9f4d98e103edb42221a22e19dd878702fe4bc6e07e803.RemoveAllDevicesFromManagementRequestBuilder) {
    return id86431f1fc4da624b5b9f4d98e103edb42221a22e19dd878702fe4bc6e07e803.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i867f2389b042d7aeb1318c9ee1fe640e5a38a5538c53de7da7f581296f833658.ReprocessLicenseAssignmentRequestBuilder) {
    return i867f2389b042d7aeb1318c9ee1fe640e5a38a5538c53de7da7f581296f833658.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i305186a396818b9b5c143c7103156f0f04d161300654a19dcdc3c7addf6a1096.RestoreRequestBuilder) {
    return i305186a396818b9b5c143c7103156f0f04d161300654a19dcdc3c7addf6a1096.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i9c40fb04136abdb492acf0bfe828594e9a8e42153a595ccb4f6f9a7ce7cb059e.RevokeSignInSessionsRequestBuilder) {
    return i9c40fb04136abdb492acf0bfe828594e9a8e42153a595ccb4f6f9a7ce7cb059e.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*idf567b32cd236785f2bc078829a07357ad80d4ae6bb4d371a9d944391faf9172.SendMailRequestBuilder) {
    return idf567b32cd236785f2bc078829a07357ad80d4ae6bb4d371a9d944391faf9172.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i8ca30bb0899ff38d6bd88d2b7106ac4c5c296c196632715417a38f5a08104192.TranslateExchangeIdsRequestBuilder) {
    return i8ca30bb0899ff38d6bd88d2b7106ac4c5c296c196632715417a38f5a08104192.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*iad6e7cf1c646d873b1d50ca39233ee024e44e57a666e6071cb9f209c2e2b26f7.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return iad6e7cf1c646d873b1d50ca39233ee024e44e57a666e6071cb9f209c2e2b26f7.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
