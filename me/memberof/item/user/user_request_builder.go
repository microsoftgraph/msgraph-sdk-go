package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0a880454c3cf095af96e4732025c8391c1240909e00cc84011051df214ca20bd "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/exportpersonaldata"
    i1c113401bbbd9b9a33f8f1f9e923237a2380c9a1678df84c51a9d0e000d60661 "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/getmembergroups"
    i218afeb040b3a0f29a652b52544ea2d385100151732655aa314d1e586982c754 "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/removealldevicesfrommanagement"
    i22744e2a74b7337e6007d2bca8ab359ad2aca503420c68af527fef009a9e6fda "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/getmemberobjects"
    i2c22e248137c16e49d971edff978ebc33d46947837f77201d007c3ee72064473 "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/checkmembergroups"
    i2db25dce6fc84282dd56d7ec887a95dacc09797c7e39e766b5e064a5822ae2fb "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/wipemanagedappregistrationsbydevicetag"
    i486f4976e7419e1470cae1193d1fc675c25959d0a888a3654f5d661468046a47 "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/sendmail"
    i6f98e4daa40b8095e1d30800bea0e796101119c045d7ac1cbf2ea8dd5f4ba2a1 "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/findmeetingtimes"
    i7b6666dde87e27206b1caa6e5e09a14104b1566fb5e50cb4cbef98cd232435c8 "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/changepassword"
    i7da34d9616f7e656a96142e917768d2396376c354a7816b0238f91854ecbb4d2 "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/getmailtips"
    i92cd43a6d941dde4c3c7d6ceca07409178163b71cf47390290c13ce0610c3684 "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/assignlicense"
    i952c95a72ca3bacfb3d070ea325bf625f0057b1316e53d9b22a8e2b5c3f4a3ba "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    ia3432185a93e8df09b4aff43526c29fca5bd9af692c4677285da34b5037ebbea "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/restore"
    ia48f9500bc7180ef91be35f25b00bbbf3b94f11c8d0102806740c785116f5bf8 "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/checkmemberobjects"
    ie1d756528f18a619a1dc4f148c439f2284d33381940730620b3a4dfe21681713 "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/translateexchangeids"
    ief32e86f33133c7840e0d5b7b69e252448d7d0fbe817ddd147c0ffe06964eb1b "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/getmanagedappdiagnosticstatuses"
    if30db53bf657ef68d267ed99e175f0a3b352042c7053ccc0fdb4b9b2eed8e9c0 "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/reprocesslicenseassignment"
    ifa195037031359998aebbeeac0f62a75d2811c70bf1c0e4373a6ab63ef8c9417 "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/revokesigninsessions"
    ife0df36d89a1d7c4392611bd234dde03549240a60372d54a92488824535a1de5 "github.com/microsoftgraph/msgraph-sdk-go/me/memberof/item/user/getmanagedapppolicies"
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
func (m *UserRequestBuilder) AssignLicense()(*i92cd43a6d941dde4c3c7d6ceca07409178163b71cf47390290c13ce0610c3684.AssignLicenseRequestBuilder) {
    return i92cd43a6d941dde4c3c7d6ceca07409178163b71cf47390290c13ce0610c3684.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i7b6666dde87e27206b1caa6e5e09a14104b1566fb5e50cb4cbef98cd232435c8.ChangePasswordRequestBuilder) {
    return i7b6666dde87e27206b1caa6e5e09a14104b1566fb5e50cb4cbef98cd232435c8.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i2c22e248137c16e49d971edff978ebc33d46947837f77201d007c3ee72064473.CheckMemberGroupsRequestBuilder) {
    return i2c22e248137c16e49d971edff978ebc33d46947837f77201d007c3ee72064473.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*ia48f9500bc7180ef91be35f25b00bbbf3b94f11c8d0102806740c785116f5bf8.CheckMemberObjectsRequestBuilder) {
    return ia48f9500bc7180ef91be35f25b00bbbf3b94f11c8d0102806740c785116f5bf8.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/memberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*i0a880454c3cf095af96e4732025c8391c1240909e00cc84011051df214ca20bd.ExportPersonalDataRequestBuilder) {
    return i0a880454c3cf095af96e4732025c8391c1240909e00cc84011051df214ca20bd.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i6f98e4daa40b8095e1d30800bea0e796101119c045d7ac1cbf2ea8dd5f4ba2a1.FindMeetingTimesRequestBuilder) {
    return i6f98e4daa40b8095e1d30800bea0e796101119c045d7ac1cbf2ea8dd5f4ba2a1.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i7da34d9616f7e656a96142e917768d2396376c354a7816b0238f91854ecbb4d2.GetMailTipsRequestBuilder) {
    return i7da34d9616f7e656a96142e917768d2396376c354a7816b0238f91854ecbb4d2.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*ief32e86f33133c7840e0d5b7b69e252448d7d0fbe817ddd147c0ffe06964eb1b.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return ief32e86f33133c7840e0d5b7b69e252448d7d0fbe817ddd147c0ffe06964eb1b.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*ife0df36d89a1d7c4392611bd234dde03549240a60372d54a92488824535a1de5.GetManagedAppPoliciesRequestBuilder) {
    return ife0df36d89a1d7c4392611bd234dde03549240a60372d54a92488824535a1de5.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i1c113401bbbd9b9a33f8f1f9e923237a2380c9a1678df84c51a9d0e000d60661.GetMemberGroupsRequestBuilder) {
    return i1c113401bbbd9b9a33f8f1f9e923237a2380c9a1678df84c51a9d0e000d60661.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i22744e2a74b7337e6007d2bca8ab359ad2aca503420c68af527fef009a9e6fda.GetMemberObjectsRequestBuilder) {
    return i22744e2a74b7337e6007d2bca8ab359ad2aca503420c68af527fef009a9e6fda.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i952c95a72ca3bacfb3d070ea325bf625f0057b1316e53d9b22a8e2b5c3f4a3ba.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i952c95a72ca3bacfb3d070ea325bf625f0057b1316e53d9b22a8e2b5c3f4a3ba.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i218afeb040b3a0f29a652b52544ea2d385100151732655aa314d1e586982c754.RemoveAllDevicesFromManagementRequestBuilder) {
    return i218afeb040b3a0f29a652b52544ea2d385100151732655aa314d1e586982c754.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*if30db53bf657ef68d267ed99e175f0a3b352042c7053ccc0fdb4b9b2eed8e9c0.ReprocessLicenseAssignmentRequestBuilder) {
    return if30db53bf657ef68d267ed99e175f0a3b352042c7053ccc0fdb4b9b2eed8e9c0.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*ia3432185a93e8df09b4aff43526c29fca5bd9af692c4677285da34b5037ebbea.RestoreRequestBuilder) {
    return ia3432185a93e8df09b4aff43526c29fca5bd9af692c4677285da34b5037ebbea.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*ifa195037031359998aebbeeac0f62a75d2811c70bf1c0e4373a6ab63ef8c9417.RevokeSignInSessionsRequestBuilder) {
    return ifa195037031359998aebbeeac0f62a75d2811c70bf1c0e4373a6ab63ef8c9417.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i486f4976e7419e1470cae1193d1fc675c25959d0a888a3654f5d661468046a47.SendMailRequestBuilder) {
    return i486f4976e7419e1470cae1193d1fc675c25959d0a888a3654f5d661468046a47.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*ie1d756528f18a619a1dc4f148c439f2284d33381940730620b3a4dfe21681713.TranslateExchangeIdsRequestBuilder) {
    return ie1d756528f18a619a1dc4f148c439f2284d33381940730620b3a4dfe21681713.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i2db25dce6fc84282dd56d7ec887a95dacc09797c7e39e766b5e064a5822ae2fb.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i2db25dce6fc84282dd56d7ec887a95dacc09797c7e39e766b5e064a5822ae2fb.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
