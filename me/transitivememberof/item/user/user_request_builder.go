package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i20a96ec08ff94c1f036dfe9a815703f5ac4db45e25e22744b13273eab981d8b6 "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/changepassword"
    i313e97e04cfcbc2585ed51edea2bacb9b3588415f9a6e4413fbf41047d7463da "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/removealldevicesfrommanagement"
    i36709a19bd6d5c9e9c4e8450d5e0d0a6762318c02b4395de30735d54780f40b3 "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/getmemberobjects"
    i5c8797c86b54d04a585d978daafe4175d678a6f5119131643d8940a8a8bcf8b3 "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/assignlicense"
    i65e2d4f204badbe0e89490ed022c388727c9878fe03d38dbf7a4698d77978b12 "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/wipemanagedappregistrationsbydevicetag"
    i698bfe6ea2111af4fa2e14650b46c942841930f18768fb7a1156a5043665a18c "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/revokesigninsessions"
    i6bea8bfc24b35e5738217a401c1e7644579a5920fa4b3051dcbf4a236593c1fb "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/checkmemberobjects"
    i956abe3613d5a84125674df5bbbe893363b506aecc5daab367db1d81147513ef "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/getmanagedappdiagnosticstatuses"
    i97edca610819d472c0449bfb8264945ff70743e5abb7cc7bcb6c23d872dbe28b "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/checkmembergroups"
    iaaec1b74cb5ec627286454fc1d786c140964149f61735038de339eaf525e1f66 "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/restore"
    iae5fe37a8e5e695606156a196fc19fc55ab7df0a120bcb642632696c8ca3fdc7 "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/getmembergroups"
    ib2335cec0589b10a3d8d65cb69e1c52aef0c4de6ba0240b4cd9cc8d035af57ec "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/translateexchangeids"
    ibc461a7540cd2c4d5c29f383b6885a6f0b2740fe747869d44df91bbae2725f0b "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/getmanagedapppolicies"
    icc0a1e3089684e362187bc2024d81bf6bfc8683811452118b811e529c673446e "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/getmailtips"
    icd6e9cecb213f3fcf24b6a466010248241e19f406b6506a68e97106203cd96cf "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/sendmail"
    id4b9c76ace5ec84254fa58989097db9162bf3d5a59c3e9a665fc5de264dd879c "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/reprocesslicenseassignment"
    ideeba3da8804cb3ad1f2c21deead722f71b9b7f4c18d066f99c05dbb9da5e0da "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/findmeetingtimes"
    if0e79908d51c2c052b70c6ef213a24e5dc8960b2e1da138f5724d95fc5533ea2 "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/exportpersonaldata"
    if7188ad9bf40cd284ce846ddb5130e7b3caadb388ba8295ba27d2f40972e8f99 "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof/item/user/reminderviewwithstartdatetimewithenddatetime"
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
func (m *UserRequestBuilder) AssignLicense()(*i5c8797c86b54d04a585d978daafe4175d678a6f5119131643d8940a8a8bcf8b3.AssignLicenseRequestBuilder) {
    return i5c8797c86b54d04a585d978daafe4175d678a6f5119131643d8940a8a8bcf8b3.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i20a96ec08ff94c1f036dfe9a815703f5ac4db45e25e22744b13273eab981d8b6.ChangePasswordRequestBuilder) {
    return i20a96ec08ff94c1f036dfe9a815703f5ac4db45e25e22744b13273eab981d8b6.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i97edca610819d472c0449bfb8264945ff70743e5abb7cc7bcb6c23d872dbe28b.CheckMemberGroupsRequestBuilder) {
    return i97edca610819d472c0449bfb8264945ff70743e5abb7cc7bcb6c23d872dbe28b.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i6bea8bfc24b35e5738217a401c1e7644579a5920fa4b3051dcbf4a236593c1fb.CheckMemberObjectsRequestBuilder) {
    return i6bea8bfc24b35e5738217a401c1e7644579a5920fa4b3051dcbf4a236593c1fb.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*if0e79908d51c2c052b70c6ef213a24e5dc8960b2e1da138f5724d95fc5533ea2.ExportPersonalDataRequestBuilder) {
    return if0e79908d51c2c052b70c6ef213a24e5dc8960b2e1da138f5724d95fc5533ea2.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*ideeba3da8804cb3ad1f2c21deead722f71b9b7f4c18d066f99c05dbb9da5e0da.FindMeetingTimesRequestBuilder) {
    return ideeba3da8804cb3ad1f2c21deead722f71b9b7f4c18d066f99c05dbb9da5e0da.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*icc0a1e3089684e362187bc2024d81bf6bfc8683811452118b811e529c673446e.GetMailTipsRequestBuilder) {
    return icc0a1e3089684e362187bc2024d81bf6bfc8683811452118b811e529c673446e.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i956abe3613d5a84125674df5bbbe893363b506aecc5daab367db1d81147513ef.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i956abe3613d5a84125674df5bbbe893363b506aecc5daab367db1d81147513ef.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*ibc461a7540cd2c4d5c29f383b6885a6f0b2740fe747869d44df91bbae2725f0b.GetManagedAppPoliciesRequestBuilder) {
    return ibc461a7540cd2c4d5c29f383b6885a6f0b2740fe747869d44df91bbae2725f0b.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*iae5fe37a8e5e695606156a196fc19fc55ab7df0a120bcb642632696c8ca3fdc7.GetMemberGroupsRequestBuilder) {
    return iae5fe37a8e5e695606156a196fc19fc55ab7df0a120bcb642632696c8ca3fdc7.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i36709a19bd6d5c9e9c4e8450d5e0d0a6762318c02b4395de30735d54780f40b3.GetMemberObjectsRequestBuilder) {
    return i36709a19bd6d5c9e9c4e8450d5e0d0a6762318c02b4395de30735d54780f40b3.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*if7188ad9bf40cd284ce846ddb5130e7b3caadb388ba8295ba27d2f40972e8f99.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return if7188ad9bf40cd284ce846ddb5130e7b3caadb388ba8295ba27d2f40972e8f99.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i313e97e04cfcbc2585ed51edea2bacb9b3588415f9a6e4413fbf41047d7463da.RemoveAllDevicesFromManagementRequestBuilder) {
    return i313e97e04cfcbc2585ed51edea2bacb9b3588415f9a6e4413fbf41047d7463da.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*id4b9c76ace5ec84254fa58989097db9162bf3d5a59c3e9a665fc5de264dd879c.ReprocessLicenseAssignmentRequestBuilder) {
    return id4b9c76ace5ec84254fa58989097db9162bf3d5a59c3e9a665fc5de264dd879c.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*iaaec1b74cb5ec627286454fc1d786c140964149f61735038de339eaf525e1f66.RestoreRequestBuilder) {
    return iaaec1b74cb5ec627286454fc1d786c140964149f61735038de339eaf525e1f66.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i698bfe6ea2111af4fa2e14650b46c942841930f18768fb7a1156a5043665a18c.RevokeSignInSessionsRequestBuilder) {
    return i698bfe6ea2111af4fa2e14650b46c942841930f18768fb7a1156a5043665a18c.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*icd6e9cecb213f3fcf24b6a466010248241e19f406b6506a68e97106203cd96cf.SendMailRequestBuilder) {
    return icd6e9cecb213f3fcf24b6a466010248241e19f406b6506a68e97106203cd96cf.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*ib2335cec0589b10a3d8d65cb69e1c52aef0c4de6ba0240b4cd9cc8d035af57ec.TranslateExchangeIdsRequestBuilder) {
    return ib2335cec0589b10a3d8d65cb69e1c52aef0c4de6ba0240b4cd9cc8d035af57ec.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i65e2d4f204badbe0e89490ed022c388727c9878fe03d38dbf7a4698d77978b12.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i65e2d4f204badbe0e89490ed022c388727c9878fe03d38dbf7a4698d77978b12.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
