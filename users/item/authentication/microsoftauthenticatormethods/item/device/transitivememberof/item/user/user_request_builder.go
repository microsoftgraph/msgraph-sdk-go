package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i028d856c2596c76bb8481cd7176554314372830778ef1f7f93491bada305047f "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmanagedappdiagnosticstatuses"
    i035c3fedb3beb427daaa7dedafe349a3469c5461081a2e0a05e66aadf225ad9d "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/changepassword"
    i08d85f847a9c61cf4a02bd6bd208f7aa28761e0c0288121ddce5a97ca04fd099 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmanagedapppolicies"
    i09c1c513fb0d981e2aa2e7bf5a75a66cc55bee3bdb3ce0817f08963870c43cba "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/findmeetingtimes"
    i42faebd32c89233709f77a4d566642f109ef1b2cdebb1324d710abf09e52479a "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/revokesigninsessions"
    i59f9b776e0651f494da0ad201b797697e2569a6dd36c9e33b6d3e99c5324a4ce "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmembergroups"
    i7d9fa46f03024d59a8f5577936dd271d01e9b7972aeb1f0acc53bea1244903f5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/wipemanagedappregistrationsbydevicetag"
    i8e643544b6784ff9f85089399da87182bd329385e36057981c620d275e35f5ef "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/checkmemberobjects"
    i91042dde85929c44b5e0adccd8172593d771d7693a397880f7e8ea8dd4f04944 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/exportpersonaldata"
    i9736cbc307c53a31b4af940061fc238b16f33ea212931aa7dc1544b85fa845db "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/checkmembergroups"
    iae61de378726b803ed455c3f3442e73902981ee6aa71ae324ff1bf5ed0dd676f "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/assignlicense"
    ib34bfb3089084a49c666d8942c60f0eeae20fc2ae677e8c26bff648f33ba2990 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmailtips"
    ib79e0da5aa3c56a5a292cbc289d1607689a82d8ff3e62cbccdd58ce77341f178 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/removealldevicesfrommanagement"
    ic0cb7a8bb6f3efe1b072f9d07f7b01e06a3bc4be9d8e402e8a6b1ff185794943 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/reprocesslicenseassignment"
    id94a08d5c8bf7c0b7f85c1b432087b864e8606b38ac774d7ddd5811c20ba5253 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/restore"
    ie33beadfa03a986c0b81ce4af89e7516b3a6984dcca4b141015144ecaeb9c52c "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmemberobjects"
    ie35688e89ee7796da3b82ddaf17a7bc13d440b3cd8586eddc732345304869b7b "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/translateexchangeids"
    ie873b85c188afcb2b2717d000325594b09c30dc82842dce0bb38fa1ea5e4333b "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/sendmail"
    ieeacdd389aa89d3d95831ca3af26ff52a26532771487a827100b0a0a3c1665c2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/reminderviewwithstartdatetimewithenddatetime"
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
func (m *UserRequestBuilder) AssignLicense()(*iae61de378726b803ed455c3f3442e73902981ee6aa71ae324ff1bf5ed0dd676f.AssignLicenseRequestBuilder) {
    return iae61de378726b803ed455c3f3442e73902981ee6aa71ae324ff1bf5ed0dd676f.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i035c3fedb3beb427daaa7dedafe349a3469c5461081a2e0a05e66aadf225ad9d.ChangePasswordRequestBuilder) {
    return i035c3fedb3beb427daaa7dedafe349a3469c5461081a2e0a05e66aadf225ad9d.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i9736cbc307c53a31b4af940061fc238b16f33ea212931aa7dc1544b85fa845db.CheckMemberGroupsRequestBuilder) {
    return i9736cbc307c53a31b4af940061fc238b16f33ea212931aa7dc1544b85fa845db.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i8e643544b6784ff9f85089399da87182bd329385e36057981c620d275e35f5ef.CheckMemberObjectsRequestBuilder) {
    return i8e643544b6784ff9f85089399da87182bd329385e36057981c620d275e35f5ef.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*i91042dde85929c44b5e0adccd8172593d771d7693a397880f7e8ea8dd4f04944.ExportPersonalDataRequestBuilder) {
    return i91042dde85929c44b5e0adccd8172593d771d7693a397880f7e8ea8dd4f04944.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i09c1c513fb0d981e2aa2e7bf5a75a66cc55bee3bdb3ce0817f08963870c43cba.FindMeetingTimesRequestBuilder) {
    return i09c1c513fb0d981e2aa2e7bf5a75a66cc55bee3bdb3ce0817f08963870c43cba.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*ib34bfb3089084a49c666d8942c60f0eeae20fc2ae677e8c26bff648f33ba2990.GetMailTipsRequestBuilder) {
    return ib34bfb3089084a49c666d8942c60f0eeae20fc2ae677e8c26bff648f33ba2990.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i028d856c2596c76bb8481cd7176554314372830778ef1f7f93491bada305047f.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i028d856c2596c76bb8481cd7176554314372830778ef1f7f93491bada305047f.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i08d85f847a9c61cf4a02bd6bd208f7aa28761e0c0288121ddce5a97ca04fd099.GetManagedAppPoliciesRequestBuilder) {
    return i08d85f847a9c61cf4a02bd6bd208f7aa28761e0c0288121ddce5a97ca04fd099.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i59f9b776e0651f494da0ad201b797697e2569a6dd36c9e33b6d3e99c5324a4ce.GetMemberGroupsRequestBuilder) {
    return i59f9b776e0651f494da0ad201b797697e2569a6dd36c9e33b6d3e99c5324a4ce.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*ie33beadfa03a986c0b81ce4af89e7516b3a6984dcca4b141015144ecaeb9c52c.GetMemberObjectsRequestBuilder) {
    return ie33beadfa03a986c0b81ce4af89e7516b3a6984dcca4b141015144ecaeb9c52c.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*ieeacdd389aa89d3d95831ca3af26ff52a26532771487a827100b0a0a3c1665c2.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return ieeacdd389aa89d3d95831ca3af26ff52a26532771487a827100b0a0a3c1665c2.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*ib79e0da5aa3c56a5a292cbc289d1607689a82d8ff3e62cbccdd58ce77341f178.RemoveAllDevicesFromManagementRequestBuilder) {
    return ib79e0da5aa3c56a5a292cbc289d1607689a82d8ff3e62cbccdd58ce77341f178.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*ic0cb7a8bb6f3efe1b072f9d07f7b01e06a3bc4be9d8e402e8a6b1ff185794943.ReprocessLicenseAssignmentRequestBuilder) {
    return ic0cb7a8bb6f3efe1b072f9d07f7b01e06a3bc4be9d8e402e8a6b1ff185794943.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*id94a08d5c8bf7c0b7f85c1b432087b864e8606b38ac774d7ddd5811c20ba5253.RestoreRequestBuilder) {
    return id94a08d5c8bf7c0b7f85c1b432087b864e8606b38ac774d7ddd5811c20ba5253.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i42faebd32c89233709f77a4d566642f109ef1b2cdebb1324d710abf09e52479a.RevokeSignInSessionsRequestBuilder) {
    return i42faebd32c89233709f77a4d566642f109ef1b2cdebb1324d710abf09e52479a.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*ie873b85c188afcb2b2717d000325594b09c30dc82842dce0bb38fa1ea5e4333b.SendMailRequestBuilder) {
    return ie873b85c188afcb2b2717d000325594b09c30dc82842dce0bb38fa1ea5e4333b.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*ie35688e89ee7796da3b82ddaf17a7bc13d440b3cd8586eddc732345304869b7b.TranslateExchangeIdsRequestBuilder) {
    return ie35688e89ee7796da3b82ddaf17a7bc13d440b3cd8586eddc732345304869b7b.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i7d9fa46f03024d59a8f5577936dd271d01e9b7972aeb1f0acc53bea1244903f5.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i7d9fa46f03024d59a8f5577936dd271d01e9b7972aeb1f0acc53bea1244903f5.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
