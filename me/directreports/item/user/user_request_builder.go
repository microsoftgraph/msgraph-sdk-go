package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i038032eb3182aebea05f30b9d2a48c12ed7c03f0758e0e9442d1ef41be459fbf "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/translateexchangeids"
    i11f0ed663691bb720ff37fc15de1e1e3a4b3dd644be18f9fee7f41485767db0c "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/sendmail"
    i144aaa65767797066ee034bf05bc73080dd1af2e6cd3898f276dfbce21bc87b4 "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/getmemberobjects"
    i16020c49a366e9b08ff7ff8913b1913e6880bc79fd12a7ac91370d0a6d2d3bef "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/getmanagedapppolicies"
    i2ade7aa107a29856e4f66d5e7060b7caf7819cb9421f497b9e1ad4fc61601a67 "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/revokesigninsessions"
    i318ec20383922500425e1aafeea339b446c40d5c15ae4b7c3528aeb7f20c2115 "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/getmailtips"
    i44a5ebc499e7aea4fc9317e0042ff68744071626d241c5f04523d70631282c42 "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/wipemanagedappregistrationsbydevicetag"
    i5dbf677435d2447b88af18b8832c4f86f65b22a5a4bebab2046a44457e0fce2f "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/getmembergroups"
    i6d58c7ff25c94928688ef0561e29629b56f54bd58c81afc1258a64fc82eab16d "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/assignlicense"
    i74344819b6fb722450592fd690add55c70132ed1743ee87c033486a897d43ec2 "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/changepassword"
    i85e15b4056a779878362d29bfd349eaf406f4068d095cc50ff412fe4e61fb3d2 "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/exportpersonaldata"
    i8aefc2ddaff09e0f332e6cd9ff38f57ce0d3568f803473478832f79b65113a45 "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/reprocesslicenseassignment"
    ia52be143d9f6fe6cbb5042ba17bb1c637098c492ca0912f7b5d3d00382081525 "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/getmanagedappdiagnosticstatuses"
    ic277fead3ef9a76e31cbd08290f93669a59574e1e8d64888d4de2fc4483d5487 "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/findmeetingtimes"
    ice6009e6caf0a2b3d9a268576300f61ec65028df3221eec90034990a9f37d508 "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/checkmemberobjects"
    id0de9cdf816bb7b176e03e4217d8b8f537802ef18df93974418b9ffe7cb44ea7 "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/checkmembergroups"
    id9c08c6f18db6726cba3e637a7ebb590d89606b64a4ef6378d66ed16aa8b3f9c "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/reminderviewwithstartdatetimewithenddatetime"
    ie51721fac761a5894cd41b20ca4aeef8a8c3430aa665139511464176384b7e64 "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/removealldevicesfrommanagement"
    ieab5e357a059bc8642609444ae11c8bb1548b87600221889a98ed659012763c3 "github.com/microsoftgraph/msgraph-sdk-go/me/directreports/item/user/restore"
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
func (m *UserRequestBuilder) AssignLicense()(*i6d58c7ff25c94928688ef0561e29629b56f54bd58c81afc1258a64fc82eab16d.AssignLicenseRequestBuilder) {
    return i6d58c7ff25c94928688ef0561e29629b56f54bd58c81afc1258a64fc82eab16d.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i74344819b6fb722450592fd690add55c70132ed1743ee87c033486a897d43ec2.ChangePasswordRequestBuilder) {
    return i74344819b6fb722450592fd690add55c70132ed1743ee87c033486a897d43ec2.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*id0de9cdf816bb7b176e03e4217d8b8f537802ef18df93974418b9ffe7cb44ea7.CheckMemberGroupsRequestBuilder) {
    return id0de9cdf816bb7b176e03e4217d8b8f537802ef18df93974418b9ffe7cb44ea7.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*ice6009e6caf0a2b3d9a268576300f61ec65028df3221eec90034990a9f37d508.CheckMemberObjectsRequestBuilder) {
    return ice6009e6caf0a2b3d9a268576300f61ec65028df3221eec90034990a9f37d508.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/directReports/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*i85e15b4056a779878362d29bfd349eaf406f4068d095cc50ff412fe4e61fb3d2.ExportPersonalDataRequestBuilder) {
    return i85e15b4056a779878362d29bfd349eaf406f4068d095cc50ff412fe4e61fb3d2.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*ic277fead3ef9a76e31cbd08290f93669a59574e1e8d64888d4de2fc4483d5487.FindMeetingTimesRequestBuilder) {
    return ic277fead3ef9a76e31cbd08290f93669a59574e1e8d64888d4de2fc4483d5487.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i318ec20383922500425e1aafeea339b446c40d5c15ae4b7c3528aeb7f20c2115.GetMailTipsRequestBuilder) {
    return i318ec20383922500425e1aafeea339b446c40d5c15ae4b7c3528aeb7f20c2115.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*ia52be143d9f6fe6cbb5042ba17bb1c637098c492ca0912f7b5d3d00382081525.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return ia52be143d9f6fe6cbb5042ba17bb1c637098c492ca0912f7b5d3d00382081525.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i16020c49a366e9b08ff7ff8913b1913e6880bc79fd12a7ac91370d0a6d2d3bef.GetManagedAppPoliciesRequestBuilder) {
    return i16020c49a366e9b08ff7ff8913b1913e6880bc79fd12a7ac91370d0a6d2d3bef.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i5dbf677435d2447b88af18b8832c4f86f65b22a5a4bebab2046a44457e0fce2f.GetMemberGroupsRequestBuilder) {
    return i5dbf677435d2447b88af18b8832c4f86f65b22a5a4bebab2046a44457e0fce2f.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i144aaa65767797066ee034bf05bc73080dd1af2e6cd3898f276dfbce21bc87b4.GetMemberObjectsRequestBuilder) {
    return i144aaa65767797066ee034bf05bc73080dd1af2e6cd3898f276dfbce21bc87b4.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*id9c08c6f18db6726cba3e637a7ebb590d89606b64a4ef6378d66ed16aa8b3f9c.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return id9c08c6f18db6726cba3e637a7ebb590d89606b64a4ef6378d66ed16aa8b3f9c.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*ie51721fac761a5894cd41b20ca4aeef8a8c3430aa665139511464176384b7e64.RemoveAllDevicesFromManagementRequestBuilder) {
    return ie51721fac761a5894cd41b20ca4aeef8a8c3430aa665139511464176384b7e64.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i8aefc2ddaff09e0f332e6cd9ff38f57ce0d3568f803473478832f79b65113a45.ReprocessLicenseAssignmentRequestBuilder) {
    return i8aefc2ddaff09e0f332e6cd9ff38f57ce0d3568f803473478832f79b65113a45.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*ieab5e357a059bc8642609444ae11c8bb1548b87600221889a98ed659012763c3.RestoreRequestBuilder) {
    return ieab5e357a059bc8642609444ae11c8bb1548b87600221889a98ed659012763c3.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i2ade7aa107a29856e4f66d5e7060b7caf7819cb9421f497b9e1ad4fc61601a67.RevokeSignInSessionsRequestBuilder) {
    return i2ade7aa107a29856e4f66d5e7060b7caf7819cb9421f497b9e1ad4fc61601a67.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i11f0ed663691bb720ff37fc15de1e1e3a4b3dd644be18f9fee7f41485767db0c.SendMailRequestBuilder) {
    return i11f0ed663691bb720ff37fc15de1e1e3a4b3dd644be18f9fee7f41485767db0c.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i038032eb3182aebea05f30b9d2a48c12ed7c03f0758e0e9442d1ef41be459fbf.TranslateExchangeIdsRequestBuilder) {
    return i038032eb3182aebea05f30b9d2a48c12ed7c03f0758e0e9442d1ef41be459fbf.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i44a5ebc499e7aea4fc9317e0042ff68744071626d241c5f04523d70631282c42.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i44a5ebc499e7aea4fc9317e0042ff68744071626d241c5f04523d70631282c42.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
