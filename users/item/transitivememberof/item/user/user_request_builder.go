package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i123124cb32d0f116ced0191259dd89964a6600e0eec6a9672c8eaca3f67b93de "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/changepassword"
    i16ed5d21fdccfad7787d56f75021a43ec816bafe85e6b4752316c53c520a378c "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/checkmembergroups"
    i2f12309df84b205ea95798556379785ed5ac4f78931f40889755d5930059b359 "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/sendmail"
    i303fdd752f971a25163315b4b82b1aaba9f11f70e27afd5be0b78b6b6372251e "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i36ae18c516a0856590ea83935aeaf28cb0a870da1342f39d52fb54224ef6fe44 "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/reprocesslicenseassignment"
    i3819afc8f934ca14d05c453c7a87fd8e6548f8f7253b0553b0ff6c8f77c69c67 "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/getmanagedapppolicies"
    i62de15d3e420465e313d0febdb721170caf935ca578616e8ea4a9d3d4ed89268 "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/restore"
    i7117965810d51221a43236e04af0df860d0e0f5b8f445e9b7823ae30b1bcb8d0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/checkmemberobjects"
    i73b0bcfed071eef60ac693f663adfc184ba5e5d1ddf0a60aecb6939670e5572f "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/translateexchangeids"
    i74d9ef94b7b4a930794d999b43aad249bb44da90d4e6d6669d74d82710ce1cb3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/removealldevicesfrommanagement"
    i797060766fe3f73f0a543c051d6d8472cba09ebeba820c0b2f869bc833ca3f39 "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/revokesigninsessions"
    i8cc877f6c0c01c7cb8c58c1ba7cb9bc40b9294fe934a3d5a4a8cff311aa18407 "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/getmembergroups"
    i92317bf18db5101823e5970b7641c253c213b6fc2b78b1f4a5aac551d5312bcd "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/assignlicense"
    i9b6c45101ab81e9ae8a093985115b0312db744456b9e3aec4b1c2ef3cf1a3edf "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/getmemberobjects"
    ibcf0acd7b393fbd65417aa4f37b45047e3dc248a82e01adedebfd38cca554764 "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/findmeetingtimes"
    ibcfe409e6d5f36c2881957688252a766d1ac802ca66238b7242844d81833c8e3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/exportpersonaldata"
    id68feaecafe791d2fe9a98ebe249d82212d56c3f36a7556b303283a6469471cb "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/getmanagedappdiagnosticstatuses"
    id7e56d3284fe41b16e8c0cc24d120373394d2aeb21eac05f45fdfa49fc0b8428 "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/getmailtips"
    ie0a524fdd3b0ca0e26822de4dd5e8a79507d73966700054a6d021e950d1e3bb4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/transitivememberof/item/user/wipemanagedappregistrationsbydevicetag"
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
func (m *UserRequestBuilder) AssignLicense()(*i92317bf18db5101823e5970b7641c253c213b6fc2b78b1f4a5aac551d5312bcd.AssignLicenseRequestBuilder) {
    return i92317bf18db5101823e5970b7641c253c213b6fc2b78b1f4a5aac551d5312bcd.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i123124cb32d0f116ced0191259dd89964a6600e0eec6a9672c8eaca3f67b93de.ChangePasswordRequestBuilder) {
    return i123124cb32d0f116ced0191259dd89964a6600e0eec6a9672c8eaca3f67b93de.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i16ed5d21fdccfad7787d56f75021a43ec816bafe85e6b4752316c53c520a378c.CheckMemberGroupsRequestBuilder) {
    return i16ed5d21fdccfad7787d56f75021a43ec816bafe85e6b4752316c53c520a378c.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i7117965810d51221a43236e04af0df860d0e0f5b8f445e9b7823ae30b1bcb8d0.CheckMemberObjectsRequestBuilder) {
    return i7117965810d51221a43236e04af0df860d0e0f5b8f445e9b7823ae30b1bcb8d0.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*ibcfe409e6d5f36c2881957688252a766d1ac802ca66238b7242844d81833c8e3.ExportPersonalDataRequestBuilder) {
    return ibcfe409e6d5f36c2881957688252a766d1ac802ca66238b7242844d81833c8e3.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*ibcf0acd7b393fbd65417aa4f37b45047e3dc248a82e01adedebfd38cca554764.FindMeetingTimesRequestBuilder) {
    return ibcf0acd7b393fbd65417aa4f37b45047e3dc248a82e01adedebfd38cca554764.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*id7e56d3284fe41b16e8c0cc24d120373394d2aeb21eac05f45fdfa49fc0b8428.GetMailTipsRequestBuilder) {
    return id7e56d3284fe41b16e8c0cc24d120373394d2aeb21eac05f45fdfa49fc0b8428.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*id68feaecafe791d2fe9a98ebe249d82212d56c3f36a7556b303283a6469471cb.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return id68feaecafe791d2fe9a98ebe249d82212d56c3f36a7556b303283a6469471cb.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i3819afc8f934ca14d05c453c7a87fd8e6548f8f7253b0553b0ff6c8f77c69c67.GetManagedAppPoliciesRequestBuilder) {
    return i3819afc8f934ca14d05c453c7a87fd8e6548f8f7253b0553b0ff6c8f77c69c67.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i8cc877f6c0c01c7cb8c58c1ba7cb9bc40b9294fe934a3d5a4a8cff311aa18407.GetMemberGroupsRequestBuilder) {
    return i8cc877f6c0c01c7cb8c58c1ba7cb9bc40b9294fe934a3d5a4a8cff311aa18407.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i9b6c45101ab81e9ae8a093985115b0312db744456b9e3aec4b1c2ef3cf1a3edf.GetMemberObjectsRequestBuilder) {
    return i9b6c45101ab81e9ae8a093985115b0312db744456b9e3aec4b1c2ef3cf1a3edf.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i303fdd752f971a25163315b4b82b1aaba9f11f70e27afd5be0b78b6b6372251e.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i303fdd752f971a25163315b4b82b1aaba9f11f70e27afd5be0b78b6b6372251e.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i74d9ef94b7b4a930794d999b43aad249bb44da90d4e6d6669d74d82710ce1cb3.RemoveAllDevicesFromManagementRequestBuilder) {
    return i74d9ef94b7b4a930794d999b43aad249bb44da90d4e6d6669d74d82710ce1cb3.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i36ae18c516a0856590ea83935aeaf28cb0a870da1342f39d52fb54224ef6fe44.ReprocessLicenseAssignmentRequestBuilder) {
    return i36ae18c516a0856590ea83935aeaf28cb0a870da1342f39d52fb54224ef6fe44.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i62de15d3e420465e313d0febdb721170caf935ca578616e8ea4a9d3d4ed89268.RestoreRequestBuilder) {
    return i62de15d3e420465e313d0febdb721170caf935ca578616e8ea4a9d3d4ed89268.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i797060766fe3f73f0a543c051d6d8472cba09ebeba820c0b2f869bc833ca3f39.RevokeSignInSessionsRequestBuilder) {
    return i797060766fe3f73f0a543c051d6d8472cba09ebeba820c0b2f869bc833ca3f39.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i2f12309df84b205ea95798556379785ed5ac4f78931f40889755d5930059b359.SendMailRequestBuilder) {
    return i2f12309df84b205ea95798556379785ed5ac4f78931f40889755d5930059b359.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i73b0bcfed071eef60ac693f663adfc184ba5e5d1ddf0a60aecb6939670e5572f.TranslateExchangeIdsRequestBuilder) {
    return i73b0bcfed071eef60ac693f663adfc184ba5e5d1ddf0a60aecb6939670e5572f.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*ie0a524fdd3b0ca0e26822de4dd5e8a79507d73966700054a6d021e950d1e3bb4.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return ie0a524fdd3b0ca0e26822de4dd5e8a79507d73966700054a6d021e950d1e3bb4.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
