package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0ab5bb55f01625cef47df421c7468109fcf5edf925821053202c13a35c5b8549 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmanagedapppolicies"
    i12a423d29c94c48a28e3c9d5ae8a35fac1df61f94ef79e81ae13016d49f48fcb "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmemberobjects"
    i154c953d34f9f1d845bc01e280e539fcb8b0372f7f3fb72eba9fdbcdf5312d0c "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i19c0d2bc3fcc24df70bfc7cf0ba58f170319f695f08dc7920b834dc3021af0f3 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/checkmemberobjects"
    i280e105d81984ed7323f6f0b7dbeb6e82fdb0581df616d7116099ceb32986917 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/revokesigninsessions"
    i2da1ac6dc91c0fa29beefab0c429724a9280d308a9b3e8e5df88d91d82ab5483 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmanagedappdiagnosticstatuses"
    i42f9209bae2554da9ee23a6ab22c20bba04f3f43172e75a0a1053cea240adedf "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/restore"
    i4e485fa06b82da57b28e8e5ecaf7ea219128d31e430758ea43f7d7cf1c9e5c6c "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/findmeetingtimes"
    i53d57a2c2e4e9e4f96549e1905c6d95520fb187a94948050f26748ad843d62d3 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/changepassword"
    i636bb86f8822abc5fd73661f5e835d833f91f915c918bd888d172e196586cda6 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmembergroups"
    i7a7a45d82e85a505520c1ccbe755db8633aeb2078ee38f9ae1d82f4e38a56cff "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/assignlicense"
    i99b96cc7e13c22785316b830164312654bd7da23181924865f35a169dcceefd8 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/sendmail"
    i9fadb215481fc5e32128553fee48f758454a29498917672f5ff8f78de85d4f7c "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/reprocesslicenseassignment"
    ibebc726109ef74cbb743e459450724337fea9d79ccb781628e71ed7432992459 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/removealldevicesfrommanagement"
    ice4508cbc1180211665b3a5598adfdfc200bcd08687a0cc7622225ad576cd230 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/translateexchangeids"
    id383f7ff3bf452725ec74b6fb5f01fd84c96e744f70c9194944ceea9650a9aa8 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/wipemanagedappregistrationsbydevicetag"
    id4b718942195c18534c137972ff0d57b14568182a5418577d10d98703b3b37dd "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/exportpersonaldata"
    ie27abddee84104a9b59e398e949b5df2bc8091518bdb3368e03542d502b46dd6 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/checkmembergroups"
    if5b9efd2f546539df4513a0a51debc499d730e73b074aeb62972b5670202ceae "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmailtips"
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
func (m *UserRequestBuilder) AssignLicense()(*i7a7a45d82e85a505520c1ccbe755db8633aeb2078ee38f9ae1d82f4e38a56cff.AssignLicenseRequestBuilder) {
    return i7a7a45d82e85a505520c1ccbe755db8633aeb2078ee38f9ae1d82f4e38a56cff.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i53d57a2c2e4e9e4f96549e1905c6d95520fb187a94948050f26748ad843d62d3.ChangePasswordRequestBuilder) {
    return i53d57a2c2e4e9e4f96549e1905c6d95520fb187a94948050f26748ad843d62d3.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*ie27abddee84104a9b59e398e949b5df2bc8091518bdb3368e03542d502b46dd6.CheckMemberGroupsRequestBuilder) {
    return ie27abddee84104a9b59e398e949b5df2bc8091518bdb3368e03542d502b46dd6.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i19c0d2bc3fcc24df70bfc7cf0ba58f170319f695f08dc7920b834dc3021af0f3.CheckMemberObjectsRequestBuilder) {
    return i19c0d2bc3fcc24df70bfc7cf0ba58f170319f695f08dc7920b834dc3021af0f3.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/memberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*id4b718942195c18534c137972ff0d57b14568182a5418577d10d98703b3b37dd.ExportPersonalDataRequestBuilder) {
    return id4b718942195c18534c137972ff0d57b14568182a5418577d10d98703b3b37dd.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i4e485fa06b82da57b28e8e5ecaf7ea219128d31e430758ea43f7d7cf1c9e5c6c.FindMeetingTimesRequestBuilder) {
    return i4e485fa06b82da57b28e8e5ecaf7ea219128d31e430758ea43f7d7cf1c9e5c6c.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*if5b9efd2f546539df4513a0a51debc499d730e73b074aeb62972b5670202ceae.GetMailTipsRequestBuilder) {
    return if5b9efd2f546539df4513a0a51debc499d730e73b074aeb62972b5670202ceae.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i2da1ac6dc91c0fa29beefab0c429724a9280d308a9b3e8e5df88d91d82ab5483.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i2da1ac6dc91c0fa29beefab0c429724a9280d308a9b3e8e5df88d91d82ab5483.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i0ab5bb55f01625cef47df421c7468109fcf5edf925821053202c13a35c5b8549.GetManagedAppPoliciesRequestBuilder) {
    return i0ab5bb55f01625cef47df421c7468109fcf5edf925821053202c13a35c5b8549.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i636bb86f8822abc5fd73661f5e835d833f91f915c918bd888d172e196586cda6.GetMemberGroupsRequestBuilder) {
    return i636bb86f8822abc5fd73661f5e835d833f91f915c918bd888d172e196586cda6.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i12a423d29c94c48a28e3c9d5ae8a35fac1df61f94ef79e81ae13016d49f48fcb.GetMemberObjectsRequestBuilder) {
    return i12a423d29c94c48a28e3c9d5ae8a35fac1df61f94ef79e81ae13016d49f48fcb.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i154c953d34f9f1d845bc01e280e539fcb8b0372f7f3fb72eba9fdbcdf5312d0c.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i154c953d34f9f1d845bc01e280e539fcb8b0372f7f3fb72eba9fdbcdf5312d0c.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*ibebc726109ef74cbb743e459450724337fea9d79ccb781628e71ed7432992459.RemoveAllDevicesFromManagementRequestBuilder) {
    return ibebc726109ef74cbb743e459450724337fea9d79ccb781628e71ed7432992459.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i9fadb215481fc5e32128553fee48f758454a29498917672f5ff8f78de85d4f7c.ReprocessLicenseAssignmentRequestBuilder) {
    return i9fadb215481fc5e32128553fee48f758454a29498917672f5ff8f78de85d4f7c.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i42f9209bae2554da9ee23a6ab22c20bba04f3f43172e75a0a1053cea240adedf.RestoreRequestBuilder) {
    return i42f9209bae2554da9ee23a6ab22c20bba04f3f43172e75a0a1053cea240adedf.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i280e105d81984ed7323f6f0b7dbeb6e82fdb0581df616d7116099ceb32986917.RevokeSignInSessionsRequestBuilder) {
    return i280e105d81984ed7323f6f0b7dbeb6e82fdb0581df616d7116099ceb32986917.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i99b96cc7e13c22785316b830164312654bd7da23181924865f35a169dcceefd8.SendMailRequestBuilder) {
    return i99b96cc7e13c22785316b830164312654bd7da23181924865f35a169dcceefd8.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*ice4508cbc1180211665b3a5598adfdfc200bcd08687a0cc7622225ad576cd230.TranslateExchangeIdsRequestBuilder) {
    return ice4508cbc1180211665b3a5598adfdfc200bcd08687a0cc7622225ad576cd230.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*id383f7ff3bf452725ec74b6fb5f01fd84c96e744f70c9194944ceea9650a9aa8.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return id383f7ff3bf452725ec74b6fb5f01fd84c96e744f70c9194944ceea9650a9aa8.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
