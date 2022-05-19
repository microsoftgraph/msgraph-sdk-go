package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0083e74708852878fb22762d7144ed8d06aee434471e1c9b28a81954ce1e7218 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i07cd499cd1414f135f86cb8562b609752b2822ac29b48a791420863400c2dcf1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/exportpersonaldata"
    i0e1bba855a3cc9fedd6811f2bb7c8d4e174bd5bb23f6fc98943a715702f3583c "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/reprocesslicenseassignment"
    i236d23924ea917fd15d3365320e6675051f4f949a3ad920f53469bb46730b779 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmanagedappdiagnosticstatuses"
    i3105393eade2f8c4596eceb1a84d040f107a2f2da7c92658d32ca47a7483a8e2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/findmeetingtimes"
    i3ab204b03316ca0c8e4ce2dc4f6a46bc28a31caba07f8400ef7e50ddd0949882 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/sendmail"
    i5acfe3f6d7c847c73302cc8c8591ac5e0193bc449ac95984b3abbec585be5318 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/checkmembergroups"
    i6377dbf56cf4f8015ddf4bf8a9529b4f7d0d509366c02a45bafaa1f3d2f01206 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmembergroups"
    i6c4cc4f4cc46edd344589cffc1591818aa60c2a0c2e12ed5d3c4b7fc5e7c3638 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/changepassword"
    i6d3f0fb153f1d47631b4928423e7589f450f4d9182367489241f1d9ae9cfccfb "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/translateexchangeids"
    i8ba4d6d27355e621c88f5dbbca08868aee43ad19fb2ad4d1a530c4c1397373ce "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/wipemanagedappregistrationsbydevicetag"
    i96cb296c1271c195174fd0868aacf2cd9abdef041f9bca5cb0a9e3d2bed263a8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmailtips"
    i99c7a9c5edec8f9cc4221d56384edfcf50173d89e73c3eeaedb573b884a10bb9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmemberobjects"
    i9ed9907e56159364501a3e537ce974c41cfdfb7d19cc48fdbf8c2c9d79522008 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/restore"
    ia4146d2eaa1d7a29e34fdfda2140081446782a907824ab0a414178cd524062c1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/revokesigninsessions"
    ib447773cdc57b2fa8412b6cdc0d2b9564ded1c1c06bd71004fdf426348acdbab "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmanagedapppolicies"
    ibab30e34a07409d2222a4fce5eb4db00bc9a581e3e23626c1063417748a64092 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/assignlicense"
    if0b11feb31654c35877b1715d3bb034b0cd1dbaf8956c23fec78acae33c3737a "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/checkmemberobjects"
    if35aa188c819453ea3bb0c0525b0ef7bd71a10ae6ec65daf42557cc1ea529788 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/removealldevicesfrommanagement"
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
func (m *UserRequestBuilder) AssignLicense()(*ibab30e34a07409d2222a4fce5eb4db00bc9a581e3e23626c1063417748a64092.AssignLicenseRequestBuilder) {
    return ibab30e34a07409d2222a4fce5eb4db00bc9a581e3e23626c1063417748a64092.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i6c4cc4f4cc46edd344589cffc1591818aa60c2a0c2e12ed5d3c4b7fc5e7c3638.ChangePasswordRequestBuilder) {
    return i6c4cc4f4cc46edd344589cffc1591818aa60c2a0c2e12ed5d3c4b7fc5e7c3638.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i5acfe3f6d7c847c73302cc8c8591ac5e0193bc449ac95984b3abbec585be5318.CheckMemberGroupsRequestBuilder) {
    return i5acfe3f6d7c847c73302cc8c8591ac5e0193bc449ac95984b3abbec585be5318.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*if0b11feb31654c35877b1715d3bb034b0cd1dbaf8956c23fec78acae33c3737a.CheckMemberObjectsRequestBuilder) {
    return if0b11feb31654c35877b1715d3bb034b0cd1dbaf8956c23fec78acae33c3737a.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*i07cd499cd1414f135f86cb8562b609752b2822ac29b48a791420863400c2dcf1.ExportPersonalDataRequestBuilder) {
    return i07cd499cd1414f135f86cb8562b609752b2822ac29b48a791420863400c2dcf1.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i3105393eade2f8c4596eceb1a84d040f107a2f2da7c92658d32ca47a7483a8e2.FindMeetingTimesRequestBuilder) {
    return i3105393eade2f8c4596eceb1a84d040f107a2f2da7c92658d32ca47a7483a8e2.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i96cb296c1271c195174fd0868aacf2cd9abdef041f9bca5cb0a9e3d2bed263a8.GetMailTipsRequestBuilder) {
    return i96cb296c1271c195174fd0868aacf2cd9abdef041f9bca5cb0a9e3d2bed263a8.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i236d23924ea917fd15d3365320e6675051f4f949a3ad920f53469bb46730b779.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i236d23924ea917fd15d3365320e6675051f4f949a3ad920f53469bb46730b779.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*ib447773cdc57b2fa8412b6cdc0d2b9564ded1c1c06bd71004fdf426348acdbab.GetManagedAppPoliciesRequestBuilder) {
    return ib447773cdc57b2fa8412b6cdc0d2b9564ded1c1c06bd71004fdf426348acdbab.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i6377dbf56cf4f8015ddf4bf8a9529b4f7d0d509366c02a45bafaa1f3d2f01206.GetMemberGroupsRequestBuilder) {
    return i6377dbf56cf4f8015ddf4bf8a9529b4f7d0d509366c02a45bafaa1f3d2f01206.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i99c7a9c5edec8f9cc4221d56384edfcf50173d89e73c3eeaedb573b884a10bb9.GetMemberObjectsRequestBuilder) {
    return i99c7a9c5edec8f9cc4221d56384edfcf50173d89e73c3eeaedb573b884a10bb9.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i0083e74708852878fb22762d7144ed8d06aee434471e1c9b28a81954ce1e7218.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i0083e74708852878fb22762d7144ed8d06aee434471e1c9b28a81954ce1e7218.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*if35aa188c819453ea3bb0c0525b0ef7bd71a10ae6ec65daf42557cc1ea529788.RemoveAllDevicesFromManagementRequestBuilder) {
    return if35aa188c819453ea3bb0c0525b0ef7bd71a10ae6ec65daf42557cc1ea529788.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i0e1bba855a3cc9fedd6811f2bb7c8d4e174bd5bb23f6fc98943a715702f3583c.ReprocessLicenseAssignmentRequestBuilder) {
    return i0e1bba855a3cc9fedd6811f2bb7c8d4e174bd5bb23f6fc98943a715702f3583c.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i9ed9907e56159364501a3e537ce974c41cfdfb7d19cc48fdbf8c2c9d79522008.RestoreRequestBuilder) {
    return i9ed9907e56159364501a3e537ce974c41cfdfb7d19cc48fdbf8c2c9d79522008.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*ia4146d2eaa1d7a29e34fdfda2140081446782a907824ab0a414178cd524062c1.RevokeSignInSessionsRequestBuilder) {
    return ia4146d2eaa1d7a29e34fdfda2140081446782a907824ab0a414178cd524062c1.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i3ab204b03316ca0c8e4ce2dc4f6a46bc28a31caba07f8400ef7e50ddd0949882.SendMailRequestBuilder) {
    return i3ab204b03316ca0c8e4ce2dc4f6a46bc28a31caba07f8400ef7e50ddd0949882.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i6d3f0fb153f1d47631b4928423e7589f450f4d9182367489241f1d9ae9cfccfb.TranslateExchangeIdsRequestBuilder) {
    return i6d3f0fb153f1d47631b4928423e7589f450f4d9182367489241f1d9ae9cfccfb.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i8ba4d6d27355e621c88f5dbbca08868aee43ad19fb2ad4d1a530c4c1397373ce.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i8ba4d6d27355e621c88f5dbbca08868aee43ad19fb2ad4d1a530c4c1397373ce.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
