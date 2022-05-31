package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i03e60c19632aacc9a3981602e7cbf40a1424177636c9d88d75478396f317eb47 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/restore"
    i0da19ed127d67a21cd33d76cd76c5c742b66859bf69d71342cb560c64bee07b3 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/assignlicense"
    i2200593b17c89b2f46c6c8f46b96b33717dd84fa3a2e886712d3e125bc8b89fd "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/sendmail"
    i2907ea9cef0667991cecbb7595e850fe5db47ac7b8ed1b1cea2abc20ad07d489 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/findmeetingtimes"
    i374278169a58cf18715f6a67e3f2c0627470d4de2412d24110ea1100c3463e26 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/reminderviewwithstartdatetimewithenddatetime"
    i38e53a0b5d3661f3cc01179035affeb45f930d3253e93942261bfc0a9d1d4f82 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/checkmemberobjects"
    i47852983a6885f8a266d9336029a431691910c9c45a676a6273edda41db26251 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmailtips"
    i57604645655956a7f24a4274f53b54d80508db76055b70ed5f9350f19ede3622 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/reprocesslicenseassignment"
    i58bc98daf4c48a0e34b160e87e4f7d94667e650db42a2969bc2e6cdeeb6e9211 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/exportpersonaldata"
    i5d82254a04d082f0654bc29328866c728e6808a035fa7fd89a5359a3827bb055 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/removealldevicesfrommanagement"
    i70e3864f7abd02bfce3bf6fcbb46dd2c2a5cce904c6505f7753c4614c5472a80 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmemberobjects"
    i71a287eda40518d95080fecc728bc69b2302afe704d95e4ee48f348058346e21 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmanagedappdiagnosticstatuses"
    i79047f95d8286d8ff6e1658dea22583da4c4f45bd28c0968ad3ab08baf8e8f00 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/checkmembergroups"
    i7a734caed5f350205090e1d19cb1240964a958a0c181dc7a2a4c5d30048dad88 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/changepassword"
    i907d89776132cbe2b1c78f35fc71a94663b396523eb05997a6fa040b0ee662d8 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmanagedapppolicies"
    ic29d7151accdc0b9a3f2f868ee7ebee862d0e229ff77c8cf62303d757b422943 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/wipemanagedappregistrationsbydevicetag"
    ic4b6724575708f7ff45dc00cf5e0046b9216f1212a1dbce55c4fd4a2951dc82c "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/translateexchangeids"
    id0d8d1c3705845a6a6c27e27d85a50f85402c501e1943dd938135070f7b8323b "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/revokesigninsessions"
    if7c3fd1e2a6e24b6f99287ba2d23bd2cd6022b0d7aa3ff678ec9aac99d37ca01 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmembergroups"
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
func (m *UserRequestBuilder) AssignLicense()(*i0da19ed127d67a21cd33d76cd76c5c742b66859bf69d71342cb560c64bee07b3.AssignLicenseRequestBuilder) {
    return i0da19ed127d67a21cd33d76cd76c5c742b66859bf69d71342cb560c64bee07b3.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i7a734caed5f350205090e1d19cb1240964a958a0c181dc7a2a4c5d30048dad88.ChangePasswordRequestBuilder) {
    return i7a734caed5f350205090e1d19cb1240964a958a0c181dc7a2a4c5d30048dad88.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i79047f95d8286d8ff6e1658dea22583da4c4f45bd28c0968ad3ab08baf8e8f00.CheckMemberGroupsRequestBuilder) {
    return i79047f95d8286d8ff6e1658dea22583da4c4f45bd28c0968ad3ab08baf8e8f00.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i38e53a0b5d3661f3cc01179035affeb45f930d3253e93942261bfc0a9d1d4f82.CheckMemberObjectsRequestBuilder) {
    return i38e53a0b5d3661f3cc01179035affeb45f930d3253e93942261bfc0a9d1d4f82.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*i58bc98daf4c48a0e34b160e87e4f7d94667e650db42a2969bc2e6cdeeb6e9211.ExportPersonalDataRequestBuilder) {
    return i58bc98daf4c48a0e34b160e87e4f7d94667e650db42a2969bc2e6cdeeb6e9211.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i2907ea9cef0667991cecbb7595e850fe5db47ac7b8ed1b1cea2abc20ad07d489.FindMeetingTimesRequestBuilder) {
    return i2907ea9cef0667991cecbb7595e850fe5db47ac7b8ed1b1cea2abc20ad07d489.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i47852983a6885f8a266d9336029a431691910c9c45a676a6273edda41db26251.GetMailTipsRequestBuilder) {
    return i47852983a6885f8a266d9336029a431691910c9c45a676a6273edda41db26251.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i71a287eda40518d95080fecc728bc69b2302afe704d95e4ee48f348058346e21.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i71a287eda40518d95080fecc728bc69b2302afe704d95e4ee48f348058346e21.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i907d89776132cbe2b1c78f35fc71a94663b396523eb05997a6fa040b0ee662d8.GetManagedAppPoliciesRequestBuilder) {
    return i907d89776132cbe2b1c78f35fc71a94663b396523eb05997a6fa040b0ee662d8.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*if7c3fd1e2a6e24b6f99287ba2d23bd2cd6022b0d7aa3ff678ec9aac99d37ca01.GetMemberGroupsRequestBuilder) {
    return if7c3fd1e2a6e24b6f99287ba2d23bd2cd6022b0d7aa3ff678ec9aac99d37ca01.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i70e3864f7abd02bfce3bf6fcbb46dd2c2a5cce904c6505f7753c4614c5472a80.GetMemberObjectsRequestBuilder) {
    return i70e3864f7abd02bfce3bf6fcbb46dd2c2a5cce904c6505f7753c4614c5472a80.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i374278169a58cf18715f6a67e3f2c0627470d4de2412d24110ea1100c3463e26.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i374278169a58cf18715f6a67e3f2c0627470d4de2412d24110ea1100c3463e26.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i5d82254a04d082f0654bc29328866c728e6808a035fa7fd89a5359a3827bb055.RemoveAllDevicesFromManagementRequestBuilder) {
    return i5d82254a04d082f0654bc29328866c728e6808a035fa7fd89a5359a3827bb055.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i57604645655956a7f24a4274f53b54d80508db76055b70ed5f9350f19ede3622.ReprocessLicenseAssignmentRequestBuilder) {
    return i57604645655956a7f24a4274f53b54d80508db76055b70ed5f9350f19ede3622.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i03e60c19632aacc9a3981602e7cbf40a1424177636c9d88d75478396f317eb47.RestoreRequestBuilder) {
    return i03e60c19632aacc9a3981602e7cbf40a1424177636c9d88d75478396f317eb47.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*id0d8d1c3705845a6a6c27e27d85a50f85402c501e1943dd938135070f7b8323b.RevokeSignInSessionsRequestBuilder) {
    return id0d8d1c3705845a6a6c27e27d85a50f85402c501e1943dd938135070f7b8323b.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i2200593b17c89b2f46c6c8f46b96b33717dd84fa3a2e886712d3e125bc8b89fd.SendMailRequestBuilder) {
    return i2200593b17c89b2f46c6c8f46b96b33717dd84fa3a2e886712d3e125bc8b89fd.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*ic4b6724575708f7ff45dc00cf5e0046b9216f1212a1dbce55c4fd4a2951dc82c.TranslateExchangeIdsRequestBuilder) {
    return ic4b6724575708f7ff45dc00cf5e0046b9216f1212a1dbce55c4fd4a2951dc82c.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*ic29d7151accdc0b9a3f2f868ee7ebee862d0e229ff77c8cf62303d757b422943.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return ic29d7151accdc0b9a3f2f868ee7ebee862d0e229ff77c8cf62303d757b422943.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
