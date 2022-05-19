package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i077dc1b76e9df6790a1219ae159404d3871e463af279de0b0bef93146dd17814 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmemberobjects"
    i0a08b8b35c5b071608057455ddc98143d4be611eccb5c5932dafcc69e91b27fb "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmailtips"
    i19b848c644a27c598565502d3060f6cfa1014ab6a26b368b47e44fc13cc7e0d4 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/checkmemberobjects"
    i21669926752a6db800b90be69f3ad9da7fb119267818c68f6540b9e987f60b54 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i246b40eb4f600536a208b7814171bf3e90ebbc0a2884990263d367bb3d30ab9c "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/checkmembergroups"
    i4a7b0354e079e6159b791ab5c813ad77dba307663b763f54395173ad8d6a2d1d "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/exportpersonaldata"
    i53ca150ad8b5d81054de858323f50860d79559e6f23946e5446ee74c93c463b6 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmembergroups"
    i621cce69ae2f00662ba2693fa3c6fc6d67ecae63bbb1c1f6aac6c7372d3eb3e1 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmanagedapppolicies"
    i8d67209cc01e8e62826dc127a78b8feef61c3626c21382d8de8a0c87c8db3b77 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/findmeetingtimes"
    i9a72bf0a709b9a1264a44d179fbd7697f8951f317a8c6857ac35ccfc7e030669 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/sendmail"
    i9b8705ac1cb1540ab47f7830654c3357c8b4ad41f88bc3acebdef171cc60e6cc "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/wipemanagedappregistrationsbydevicetag"
    iacae38eba8ac2458a718e857e2caafdafd55646e8488497bd60135f81a0cf351 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/reprocesslicenseassignment"
    iad3f1638310677f43f8e286dfd1f21a3e2fa9463683bbb9e8688e54e0a3c068f "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmanagedappdiagnosticstatuses"
    ib7308f26d0f9f7a555f3e90e1e5cb1e6a994f2baf8ebd34ac8d18b4a5b37a2dc "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/changepassword"
    ic0f9ce9ed1a638c5cc6e51ae35377581735a83e122197204463ba0a8865aea10 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/translateexchangeids"
    ie03646b662a708ade3d9cf7fca1785d6515ef8b6533db435fe8fa11ebe0d5c79 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/removealldevicesfrommanagement"
    ie0a31115bc7c22a24bdd5031c6ca9aeb2d57ec4ad7cde59a0824b252f320b7d6 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/assignlicense"
    ie96899382e1a4ded44eac2d16102755eec34c309f967fa57612b83c0403b2800 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/revokesigninsessions"
    ifffa551972f00a57f3c53d01126817ccfee4a47b42e30e7ca44d5462532addc0 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/restore"
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
func (m *UserRequestBuilder) AssignLicense()(*ie0a31115bc7c22a24bdd5031c6ca9aeb2d57ec4ad7cde59a0824b252f320b7d6.AssignLicenseRequestBuilder) {
    return ie0a31115bc7c22a24bdd5031c6ca9aeb2d57ec4ad7cde59a0824b252f320b7d6.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*ib7308f26d0f9f7a555f3e90e1e5cb1e6a994f2baf8ebd34ac8d18b4a5b37a2dc.ChangePasswordRequestBuilder) {
    return ib7308f26d0f9f7a555f3e90e1e5cb1e6a994f2baf8ebd34ac8d18b4a5b37a2dc.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i246b40eb4f600536a208b7814171bf3e90ebbc0a2884990263d367bb3d30ab9c.CheckMemberGroupsRequestBuilder) {
    return i246b40eb4f600536a208b7814171bf3e90ebbc0a2884990263d367bb3d30ab9c.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i19b848c644a27c598565502d3060f6cfa1014ab6a26b368b47e44fc13cc7e0d4.CheckMemberObjectsRequestBuilder) {
    return i19b848c644a27c598565502d3060f6cfa1014ab6a26b368b47e44fc13cc7e0d4.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*i4a7b0354e079e6159b791ab5c813ad77dba307663b763f54395173ad8d6a2d1d.ExportPersonalDataRequestBuilder) {
    return i4a7b0354e079e6159b791ab5c813ad77dba307663b763f54395173ad8d6a2d1d.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i8d67209cc01e8e62826dc127a78b8feef61c3626c21382d8de8a0c87c8db3b77.FindMeetingTimesRequestBuilder) {
    return i8d67209cc01e8e62826dc127a78b8feef61c3626c21382d8de8a0c87c8db3b77.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i0a08b8b35c5b071608057455ddc98143d4be611eccb5c5932dafcc69e91b27fb.GetMailTipsRequestBuilder) {
    return i0a08b8b35c5b071608057455ddc98143d4be611eccb5c5932dafcc69e91b27fb.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*iad3f1638310677f43f8e286dfd1f21a3e2fa9463683bbb9e8688e54e0a3c068f.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return iad3f1638310677f43f8e286dfd1f21a3e2fa9463683bbb9e8688e54e0a3c068f.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i621cce69ae2f00662ba2693fa3c6fc6d67ecae63bbb1c1f6aac6c7372d3eb3e1.GetManagedAppPoliciesRequestBuilder) {
    return i621cce69ae2f00662ba2693fa3c6fc6d67ecae63bbb1c1f6aac6c7372d3eb3e1.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i53ca150ad8b5d81054de858323f50860d79559e6f23946e5446ee74c93c463b6.GetMemberGroupsRequestBuilder) {
    return i53ca150ad8b5d81054de858323f50860d79559e6f23946e5446ee74c93c463b6.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i077dc1b76e9df6790a1219ae159404d3871e463af279de0b0bef93146dd17814.GetMemberObjectsRequestBuilder) {
    return i077dc1b76e9df6790a1219ae159404d3871e463af279de0b0bef93146dd17814.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i21669926752a6db800b90be69f3ad9da7fb119267818c68f6540b9e987f60b54.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i21669926752a6db800b90be69f3ad9da7fb119267818c68f6540b9e987f60b54.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*ie03646b662a708ade3d9cf7fca1785d6515ef8b6533db435fe8fa11ebe0d5c79.RemoveAllDevicesFromManagementRequestBuilder) {
    return ie03646b662a708ade3d9cf7fca1785d6515ef8b6533db435fe8fa11ebe0d5c79.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*iacae38eba8ac2458a718e857e2caafdafd55646e8488497bd60135f81a0cf351.ReprocessLicenseAssignmentRequestBuilder) {
    return iacae38eba8ac2458a718e857e2caafdafd55646e8488497bd60135f81a0cf351.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*ifffa551972f00a57f3c53d01126817ccfee4a47b42e30e7ca44d5462532addc0.RestoreRequestBuilder) {
    return ifffa551972f00a57f3c53d01126817ccfee4a47b42e30e7ca44d5462532addc0.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*ie96899382e1a4ded44eac2d16102755eec34c309f967fa57612b83c0403b2800.RevokeSignInSessionsRequestBuilder) {
    return ie96899382e1a4ded44eac2d16102755eec34c309f967fa57612b83c0403b2800.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i9a72bf0a709b9a1264a44d179fbd7697f8951f317a8c6857ac35ccfc7e030669.SendMailRequestBuilder) {
    return i9a72bf0a709b9a1264a44d179fbd7697f8951f317a8c6857ac35ccfc7e030669.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*ic0f9ce9ed1a638c5cc6e51ae35377581735a83e122197204463ba0a8865aea10.TranslateExchangeIdsRequestBuilder) {
    return ic0f9ce9ed1a638c5cc6e51ae35377581735a83e122197204463ba0a8865aea10.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i9b8705ac1cb1540ab47f7830654c3357c8b4ad41f88bc3acebdef171cc60e6cc.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i9b8705ac1cb1540ab47f7830654c3357c8b4ad41f88bc3acebdef171cc60e6cc.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
