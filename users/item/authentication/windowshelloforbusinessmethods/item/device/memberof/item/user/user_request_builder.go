package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i27656711765d54a3589234008dc4ad3e6157523684aa07fcf1fa647d31c1737b "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmailtips"
    i29a64d0c231d25aece4888051b4d712bed7a91ac84a2f79b6f58ec35ab0abdc2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/checkmemberobjects"
    i3a6e7fc67dcbe5c727a57046c31387bfba9691b311e31a6310f2b8455d67c9c2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmanagedapppolicies"
    i3f72e87e09bf220e736d42be3e11b3886981c6247a8e5026998b2015787f9db9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmemberobjects"
    i4228441fd1ea845d1591ab5dd3d0d9978a7e11dae17fc875ec5f21c17482a9c5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/findmeetingtimes"
    i5da7208ad43143cff1704d8f41229e4a9963fe880d1ad11aa1a08c2a515d59c8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/changepassword"
    i7091ea07152f1189ef846791c5d42acd550808a104d38d492ccd050e85b17f8e "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/revokesigninsessions"
    i76792c3a0a4dfb31454482fe2cadc07fdbef17979f0eb3b3a8dc38842c9e1928 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmembergroups"
    i78bcd3b9a441ba3ccb12aa6b3e528565c2dfe604db99f56ff8dc0e3ac4710d52 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/reprocesslicenseassignment"
    i8ed52174ad48ac48aa0d02de6a070a549781d608a9b4d478fa3b160435861f56 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/wipemanagedappregistrationsbydevicetag"
    i989327e1b623a455855e8d0a09a7fe5dfaa44a35367030fb9c29c2cbc69a9f87 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/restore"
    i9afbbb359058b7f6d87e0449ec86e13b5845502011a2ba96b66026e9f7e8bd47 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    ic134608a89c2b4a727d36832e03666eb6d70732a591bd15bd65bf7cbb3405510 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/sendmail"
    ic4d00f9bb525517fadbed9794791fcc4a529adaf364f07ecfa1dafcf10e60069 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/removealldevicesfrommanagement"
    icb04b165057d4c082d5a601d42f506cdea104a54b5aa9257623de34913f258c1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/checkmembergroups"
    id389f98ff368923a3b029712a02abd2f698fe296bdc25c2e502a4c531917105a "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/exportpersonaldata"
    idb0e03b6302ecc78d9cfce9b814567c219d1dd89a9cfc6f3b3c79cb90a278c9d "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/translateexchangeids"
    ie4ae5086541b134d78296a077da7a28500650e04bbe2de5d1d8e131fea02632a "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmanagedappdiagnosticstatuses"
    if25d55f0c36152664dd4fa21c13baf07355eab348f6d4e82949f5c226454c82b "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/assignlicense"
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
func (m *UserRequestBuilder) AssignLicense()(*if25d55f0c36152664dd4fa21c13baf07355eab348f6d4e82949f5c226454c82b.AssignLicenseRequestBuilder) {
    return if25d55f0c36152664dd4fa21c13baf07355eab348f6d4e82949f5c226454c82b.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i5da7208ad43143cff1704d8f41229e4a9963fe880d1ad11aa1a08c2a515d59c8.ChangePasswordRequestBuilder) {
    return i5da7208ad43143cff1704d8f41229e4a9963fe880d1ad11aa1a08c2a515d59c8.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*icb04b165057d4c082d5a601d42f506cdea104a54b5aa9257623de34913f258c1.CheckMemberGroupsRequestBuilder) {
    return icb04b165057d4c082d5a601d42f506cdea104a54b5aa9257623de34913f258c1.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i29a64d0c231d25aece4888051b4d712bed7a91ac84a2f79b6f58ec35ab0abdc2.CheckMemberObjectsRequestBuilder) {
    return i29a64d0c231d25aece4888051b4d712bed7a91ac84a2f79b6f58ec35ab0abdc2.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/memberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*id389f98ff368923a3b029712a02abd2f698fe296bdc25c2e502a4c531917105a.ExportPersonalDataRequestBuilder) {
    return id389f98ff368923a3b029712a02abd2f698fe296bdc25c2e502a4c531917105a.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i4228441fd1ea845d1591ab5dd3d0d9978a7e11dae17fc875ec5f21c17482a9c5.FindMeetingTimesRequestBuilder) {
    return i4228441fd1ea845d1591ab5dd3d0d9978a7e11dae17fc875ec5f21c17482a9c5.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i27656711765d54a3589234008dc4ad3e6157523684aa07fcf1fa647d31c1737b.GetMailTipsRequestBuilder) {
    return i27656711765d54a3589234008dc4ad3e6157523684aa07fcf1fa647d31c1737b.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*ie4ae5086541b134d78296a077da7a28500650e04bbe2de5d1d8e131fea02632a.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return ie4ae5086541b134d78296a077da7a28500650e04bbe2de5d1d8e131fea02632a.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i3a6e7fc67dcbe5c727a57046c31387bfba9691b311e31a6310f2b8455d67c9c2.GetManagedAppPoliciesRequestBuilder) {
    return i3a6e7fc67dcbe5c727a57046c31387bfba9691b311e31a6310f2b8455d67c9c2.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i76792c3a0a4dfb31454482fe2cadc07fdbef17979f0eb3b3a8dc38842c9e1928.GetMemberGroupsRequestBuilder) {
    return i76792c3a0a4dfb31454482fe2cadc07fdbef17979f0eb3b3a8dc38842c9e1928.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i3f72e87e09bf220e736d42be3e11b3886981c6247a8e5026998b2015787f9db9.GetMemberObjectsRequestBuilder) {
    return i3f72e87e09bf220e736d42be3e11b3886981c6247a8e5026998b2015787f9db9.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i9afbbb359058b7f6d87e0449ec86e13b5845502011a2ba96b66026e9f7e8bd47.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i9afbbb359058b7f6d87e0449ec86e13b5845502011a2ba96b66026e9f7e8bd47.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*ic4d00f9bb525517fadbed9794791fcc4a529adaf364f07ecfa1dafcf10e60069.RemoveAllDevicesFromManagementRequestBuilder) {
    return ic4d00f9bb525517fadbed9794791fcc4a529adaf364f07ecfa1dafcf10e60069.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i78bcd3b9a441ba3ccb12aa6b3e528565c2dfe604db99f56ff8dc0e3ac4710d52.ReprocessLicenseAssignmentRequestBuilder) {
    return i78bcd3b9a441ba3ccb12aa6b3e528565c2dfe604db99f56ff8dc0e3ac4710d52.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i989327e1b623a455855e8d0a09a7fe5dfaa44a35367030fb9c29c2cbc69a9f87.RestoreRequestBuilder) {
    return i989327e1b623a455855e8d0a09a7fe5dfaa44a35367030fb9c29c2cbc69a9f87.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i7091ea07152f1189ef846791c5d42acd550808a104d38d492ccd050e85b17f8e.RevokeSignInSessionsRequestBuilder) {
    return i7091ea07152f1189ef846791c5d42acd550808a104d38d492ccd050e85b17f8e.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*ic134608a89c2b4a727d36832e03666eb6d70732a591bd15bd65bf7cbb3405510.SendMailRequestBuilder) {
    return ic134608a89c2b4a727d36832e03666eb6d70732a591bd15bd65bf7cbb3405510.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*idb0e03b6302ecc78d9cfce9b814567c219d1dd89a9cfc6f3b3c79cb90a278c9d.TranslateExchangeIdsRequestBuilder) {
    return idb0e03b6302ecc78d9cfce9b814567c219d1dd89a9cfc6f3b3c79cb90a278c9d.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i8ed52174ad48ac48aa0d02de6a070a549781d608a9b4d478fa3b160435861f56.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i8ed52174ad48ac48aa0d02de6a070a549781d608a9b4d478fa3b160435861f56.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
