package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i034fa792d8d732212dd65b01933bc86521305a49c30f7a027777d676991b9745 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmanagedapppolicies"
    i04f4ac8679766f59b030c208d7155434ceea796e0f1425c8f396b305c6f65d14 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/exportpersonaldata"
    i080c93f0fc3e2252ccd0d360b957f43630307da6d4ccd5e218976e35a0a93343 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/checkmemberobjects"
    i0e2465baf30ad8b5daeb6f95db5bb79827e78a6e86a75c442d8f54723c6e762e "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmanagedappdiagnosticstatuses"
    i17331ea7b59f3efbbb4ded48dbaaca3b2810885e8796e38510e7793b080b079c "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/removealldevicesfrommanagement"
    i2d74c84278ccd3d971a381edafe8449b845bde6d12ddb154a001c55ce49be7cd "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/reprocesslicenseassignment"
    i2d950fe7bf1b680578f13be1d5b2bf638265e1bc4bda960aca5f2291653ee965 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/translateexchangeids"
    i426f11402e25f086996da0d809d506d0b0978d9d83878bc5f169e6cab5b889d2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/wipemanagedappregistrationsbydevicetag"
    i53334a2d7b6a6feb7fd45cb667c9fa0b1794391241d9c0fa729c9b5ccfc18b2c "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/findmeetingtimes"
    i648af854c4485da70b361f9659df836960cc57d69b78915689a059504a57498a "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmemberobjects"
    i78a4b2eb7c96134e7961b11034f9101279a9286248a01a1c94cd560b85c5c16f "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/revokesigninsessions"
    i8b6b3e36ede74ad351e2f9dfd8debe78e83ff91e7502a134f273d8c73c0a4d09 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/restore"
    i901df12953e4d2f9818ac828e3c1185db28f86de4680aef4dd14ebf51bb79b07 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/assignlicense"
    ia4df0b2fbeab316a150165e7e3759595df042429e76fe5222aec70c13902009c "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/checkmembergroups"
    icbabe7b460b30d441e9cfa5a45121995adcb634ab8a628751df6838a8bb82c30 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmailtips"
    icfc51e276dcb2b0e7912a60fde2b7ccfac0b350c6619752510e618b869856a90 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmembergroups"
    if1db4527f1413d3c2e8b1cf05da9ff6d087bce7356972fdc57aee3a87ca03c47 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/sendmail"
    if5540577fe6052644675e8c07cc1352727f132d6a9b82dd638433322b84de6be "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/changepassword"
    if8e09a3b559d78053d47729b71d441a78954af2b6b5d089773019954401c9e69 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/reminderviewwithstartdatetimewithenddatetime"
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
func (m *UserRequestBuilder) AssignLicense()(*i901df12953e4d2f9818ac828e3c1185db28f86de4680aef4dd14ebf51bb79b07.AssignLicenseRequestBuilder) {
    return i901df12953e4d2f9818ac828e3c1185db28f86de4680aef4dd14ebf51bb79b07.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*if5540577fe6052644675e8c07cc1352727f132d6a9b82dd638433322b84de6be.ChangePasswordRequestBuilder) {
    return if5540577fe6052644675e8c07cc1352727f132d6a9b82dd638433322b84de6be.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*ia4df0b2fbeab316a150165e7e3759595df042429e76fe5222aec70c13902009c.CheckMemberGroupsRequestBuilder) {
    return ia4df0b2fbeab316a150165e7e3759595df042429e76fe5222aec70c13902009c.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i080c93f0fc3e2252ccd0d360b957f43630307da6d4ccd5e218976e35a0a93343.CheckMemberObjectsRequestBuilder) {
    return i080c93f0fc3e2252ccd0d360b957f43630307da6d4ccd5e218976e35a0a93343.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/registeredUsers/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*i04f4ac8679766f59b030c208d7155434ceea796e0f1425c8f396b305c6f65d14.ExportPersonalDataRequestBuilder) {
    return i04f4ac8679766f59b030c208d7155434ceea796e0f1425c8f396b305c6f65d14.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i53334a2d7b6a6feb7fd45cb667c9fa0b1794391241d9c0fa729c9b5ccfc18b2c.FindMeetingTimesRequestBuilder) {
    return i53334a2d7b6a6feb7fd45cb667c9fa0b1794391241d9c0fa729c9b5ccfc18b2c.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*icbabe7b460b30d441e9cfa5a45121995adcb634ab8a628751df6838a8bb82c30.GetMailTipsRequestBuilder) {
    return icbabe7b460b30d441e9cfa5a45121995adcb634ab8a628751df6838a8bb82c30.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i0e2465baf30ad8b5daeb6f95db5bb79827e78a6e86a75c442d8f54723c6e762e.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i0e2465baf30ad8b5daeb6f95db5bb79827e78a6e86a75c442d8f54723c6e762e.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i034fa792d8d732212dd65b01933bc86521305a49c30f7a027777d676991b9745.GetManagedAppPoliciesRequestBuilder) {
    return i034fa792d8d732212dd65b01933bc86521305a49c30f7a027777d676991b9745.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*icfc51e276dcb2b0e7912a60fde2b7ccfac0b350c6619752510e618b869856a90.GetMemberGroupsRequestBuilder) {
    return icfc51e276dcb2b0e7912a60fde2b7ccfac0b350c6619752510e618b869856a90.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i648af854c4485da70b361f9659df836960cc57d69b78915689a059504a57498a.GetMemberObjectsRequestBuilder) {
    return i648af854c4485da70b361f9659df836960cc57d69b78915689a059504a57498a.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*if8e09a3b559d78053d47729b71d441a78954af2b6b5d089773019954401c9e69.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return if8e09a3b559d78053d47729b71d441a78954af2b6b5d089773019954401c9e69.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i17331ea7b59f3efbbb4ded48dbaaca3b2810885e8796e38510e7793b080b079c.RemoveAllDevicesFromManagementRequestBuilder) {
    return i17331ea7b59f3efbbb4ded48dbaaca3b2810885e8796e38510e7793b080b079c.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i2d74c84278ccd3d971a381edafe8449b845bde6d12ddb154a001c55ce49be7cd.ReprocessLicenseAssignmentRequestBuilder) {
    return i2d74c84278ccd3d971a381edafe8449b845bde6d12ddb154a001c55ce49be7cd.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i8b6b3e36ede74ad351e2f9dfd8debe78e83ff91e7502a134f273d8c73c0a4d09.RestoreRequestBuilder) {
    return i8b6b3e36ede74ad351e2f9dfd8debe78e83ff91e7502a134f273d8c73c0a4d09.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i78a4b2eb7c96134e7961b11034f9101279a9286248a01a1c94cd560b85c5c16f.RevokeSignInSessionsRequestBuilder) {
    return i78a4b2eb7c96134e7961b11034f9101279a9286248a01a1c94cd560b85c5c16f.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*if1db4527f1413d3c2e8b1cf05da9ff6d087bce7356972fdc57aee3a87ca03c47.SendMailRequestBuilder) {
    return if1db4527f1413d3c2e8b1cf05da9ff6d087bce7356972fdc57aee3a87ca03c47.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i2d950fe7bf1b680578f13be1d5b2bf638265e1bc4bda960aca5f2291653ee965.TranslateExchangeIdsRequestBuilder) {
    return i2d950fe7bf1b680578f13be1d5b2bf638265e1bc4bda960aca5f2291653ee965.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i426f11402e25f086996da0d809d506d0b0978d9d83878bc5f169e6cab5b889d2.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i426f11402e25f086996da0d809d506d0b0978d9d83878bc5f169e6cab5b889d2.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
