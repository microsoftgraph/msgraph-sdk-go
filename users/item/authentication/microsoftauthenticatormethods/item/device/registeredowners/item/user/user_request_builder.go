package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i42b6cbbb96e43f6b4a702f1213f84934b54ff0d59d1e9f3f6eb090f3c7df6d49 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/checkmembergroups"
    i460459fa38e0e46706757f32cd8fafc8b39d6a65e829684057f1d4339ec943bb "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/revokesigninsessions"
    i4ac7c771d326a82a1e887e9fc95b44df1ec0eb62652c3e955cd86c94f98a1fd3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmanagedappdiagnosticstatuses"
    i53824c04caa2e75314b4185b4b894cd6d3cda4b374c0c7b6c6d9d5daff25b090 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/assignlicense"
    i56fbdea521800c1be2a951e3dc5e2ffdaa33fb8fa154f96d6a0ec348c229b930 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/reminderviewwithstartdatetimewithenddatetime"
    i6b02f0f707574f9ecaa0b1f816b60fec8be648483c00a91749b62947df6a811f "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmailtips"
    i8ffb6c53a48bf5d5dfa30768495887232b83e5521e0ba68aaef8a8757b9de1fb "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/wipemanagedappregistrationsbydevicetag"
    i91bd16431d50f3bb6a63b89e87ed7e28eeda81ae4008046373026a8ea39660c3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/reprocesslicenseassignment"
    i9232397521e7cb5e41b0450ef3a85eb40f2191b570be1a74e1cfeae4c021cafd "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/removealldevicesfrommanagement"
    i9325417a7dec248ff14b87a203ace4bd023a244dcb1b1f1c5dd7e4e130e7c401 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/checkmemberobjects"
    i9647905cc94ec1c1fa6274e97a49cb2112d88b9332db815a919bf3f0b739bcb3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/translateexchangeids"
    ib4a63432ff8e3369484a62278d17791662add897af036ce35ab28b65e6922763 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmembergroups"
    ib6180c6cbcc56404eb309a6de00da62ef541ca7f69452a4c5fa5c1f65bf29f2d "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/restore"
    icbb2e48d697c332bcd60339175f7024fb9f77567e6bf0b3aa099ba0dff079239 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/sendmail"
    icd0d81e85258426ca59714ca85c5cc1e6ed33e2b3f24d94ef88f5406ca5cd369 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmanagedapppolicies"
    id3e092ac6fc8a5f7a6271a246556b0114db0da973987038a8c7a2ba307e63cf4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/changepassword"
    ie297f016b1fd3e4b1ad485773062f0c2a39b81d8fbe294d9dabdf0abfc485f8f "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/exportpersonaldata"
    ie5ff498d60e7ad840f2ed33f9b76dd5b9ffc714015c47eaeb8d94947e8538fa2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmemberobjects"
    ie6db44fb040db75eed47f09062dba2507f20737057a6295d94546442bec5fb23 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/findmeetingtimes"
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
func (m *UserRequestBuilder) AssignLicense()(*i53824c04caa2e75314b4185b4b894cd6d3cda4b374c0c7b6c6d9d5daff25b090.AssignLicenseRequestBuilder) {
    return i53824c04caa2e75314b4185b4b894cd6d3cda4b374c0c7b6c6d9d5daff25b090.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*id3e092ac6fc8a5f7a6271a246556b0114db0da973987038a8c7a2ba307e63cf4.ChangePasswordRequestBuilder) {
    return id3e092ac6fc8a5f7a6271a246556b0114db0da973987038a8c7a2ba307e63cf4.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i42b6cbbb96e43f6b4a702f1213f84934b54ff0d59d1e9f3f6eb090f3c7df6d49.CheckMemberGroupsRequestBuilder) {
    return i42b6cbbb96e43f6b4a702f1213f84934b54ff0d59d1e9f3f6eb090f3c7df6d49.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i9325417a7dec248ff14b87a203ace4bd023a244dcb1b1f1c5dd7e4e130e7c401.CheckMemberObjectsRequestBuilder) {
    return i9325417a7dec248ff14b87a203ace4bd023a244dcb1b1f1c5dd7e4e130e7c401.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*ie297f016b1fd3e4b1ad485773062f0c2a39b81d8fbe294d9dabdf0abfc485f8f.ExportPersonalDataRequestBuilder) {
    return ie297f016b1fd3e4b1ad485773062f0c2a39b81d8fbe294d9dabdf0abfc485f8f.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*ie6db44fb040db75eed47f09062dba2507f20737057a6295d94546442bec5fb23.FindMeetingTimesRequestBuilder) {
    return ie6db44fb040db75eed47f09062dba2507f20737057a6295d94546442bec5fb23.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i6b02f0f707574f9ecaa0b1f816b60fec8be648483c00a91749b62947df6a811f.GetMailTipsRequestBuilder) {
    return i6b02f0f707574f9ecaa0b1f816b60fec8be648483c00a91749b62947df6a811f.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i4ac7c771d326a82a1e887e9fc95b44df1ec0eb62652c3e955cd86c94f98a1fd3.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i4ac7c771d326a82a1e887e9fc95b44df1ec0eb62652c3e955cd86c94f98a1fd3.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*icd0d81e85258426ca59714ca85c5cc1e6ed33e2b3f24d94ef88f5406ca5cd369.GetManagedAppPoliciesRequestBuilder) {
    return icd0d81e85258426ca59714ca85c5cc1e6ed33e2b3f24d94ef88f5406ca5cd369.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*ib4a63432ff8e3369484a62278d17791662add897af036ce35ab28b65e6922763.GetMemberGroupsRequestBuilder) {
    return ib4a63432ff8e3369484a62278d17791662add897af036ce35ab28b65e6922763.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*ie5ff498d60e7ad840f2ed33f9b76dd5b9ffc714015c47eaeb8d94947e8538fa2.GetMemberObjectsRequestBuilder) {
    return ie5ff498d60e7ad840f2ed33f9b76dd5b9ffc714015c47eaeb8d94947e8538fa2.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i56fbdea521800c1be2a951e3dc5e2ffdaa33fb8fa154f96d6a0ec348c229b930.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i56fbdea521800c1be2a951e3dc5e2ffdaa33fb8fa154f96d6a0ec348c229b930.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i9232397521e7cb5e41b0450ef3a85eb40f2191b570be1a74e1cfeae4c021cafd.RemoveAllDevicesFromManagementRequestBuilder) {
    return i9232397521e7cb5e41b0450ef3a85eb40f2191b570be1a74e1cfeae4c021cafd.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i91bd16431d50f3bb6a63b89e87ed7e28eeda81ae4008046373026a8ea39660c3.ReprocessLicenseAssignmentRequestBuilder) {
    return i91bd16431d50f3bb6a63b89e87ed7e28eeda81ae4008046373026a8ea39660c3.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*ib6180c6cbcc56404eb309a6de00da62ef541ca7f69452a4c5fa5c1f65bf29f2d.RestoreRequestBuilder) {
    return ib6180c6cbcc56404eb309a6de00da62ef541ca7f69452a4c5fa5c1f65bf29f2d.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i460459fa38e0e46706757f32cd8fafc8b39d6a65e829684057f1d4339ec943bb.RevokeSignInSessionsRequestBuilder) {
    return i460459fa38e0e46706757f32cd8fafc8b39d6a65e829684057f1d4339ec943bb.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*icbb2e48d697c332bcd60339175f7024fb9f77567e6bf0b3aa099ba0dff079239.SendMailRequestBuilder) {
    return icbb2e48d697c332bcd60339175f7024fb9f77567e6bf0b3aa099ba0dff079239.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i9647905cc94ec1c1fa6274e97a49cb2112d88b9332db815a919bf3f0b739bcb3.TranslateExchangeIdsRequestBuilder) {
    return i9647905cc94ec1c1fa6274e97a49cb2112d88b9332db815a919bf3f0b739bcb3.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i8ffb6c53a48bf5d5dfa30768495887232b83e5521e0ba68aaef8a8757b9de1fb.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i8ffb6c53a48bf5d5dfa30768495887232b83e5521e0ba68aaef8a8757b9de1fb.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
