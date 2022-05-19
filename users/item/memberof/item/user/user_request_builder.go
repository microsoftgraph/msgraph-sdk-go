package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i326394449db0a190f0a2fc474b2a490675d74783e92d65b49067d0d30f78fd22 "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/assignlicense"
    i45de06b1f5a9abeed74b8c95e82c981451cd0ab816d7167324b763a72c339357 "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i4b72766e0b8b1aa132b8a4f13dfc7b0058530145c4261d70e644554128b8e399 "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/removealldevicesfrommanagement"
    i4e27058d8f79e692a1947cfb44e71bec11640181e9986ffcdeabaceb8c7af4c5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/getmemberobjects"
    i54ccdd10e1babdf7a720ac0f0ee6fc672fcb354e879b96f221442e162f52d852 "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/reprocesslicenseassignment"
    i59d9e41e5e6a3b4a93da15b3d40653e21265ab3d9ae0e17b04976db5512964ef "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/getmanagedapppolicies"
    i59f17085687504d2eb6620b2058ff580e0aa57be0923a4a8ce3d844e01af02ef "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/exportpersonaldata"
    i61cc66f6561fe8559fcabff44e7cf82443e56b35785bcaf90ff6484a4493bdc8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/getmembergroups"
    i662db12ac01d7038a5b93f0f1ab3c8e498be98d04a1dc41c7f29e8be9b3e2249 "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/findmeetingtimes"
    i7ad434fbe1ad63417dd1eb08e62642968ea0c3a31ea31f74c6c6c7b3593733ba "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/getmailtips"
    i7f352500093baef1b19a05301d65737715f88454230ea216b5487620e23f9bcb "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/restore"
    i85c043082ce6971a9408255031ce028bcfcf317cd395b030c97bc4e333e1fd7e "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/revokesigninsessions"
    i94bbd38de045f07dd4203aaeaba941d25dc5107a8bc877d7be84e4f37851c888 "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/getmanagedappdiagnosticstatuses"
    i9d8c1789bc721e6f87af390e2c121fef26320b3f4f4b02cdd3d3f92564822b7d "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/checkmembergroups"
    ia676b0c8452eaa27680a5f824494afcd62ca91248d5d1435c27962dafbd089ef "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/translateexchangeids"
    ic2fd0cc125167b6efbade089bb78f3b71befeecf38130be5ce3fddc6c93e991c "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/wipemanagedappregistrationsbydevicetag"
    icc66bcf11474ad72318c4c646ba46e93f45c40b8ca05c9f0ce5bd068c5528b72 "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/checkmemberobjects"
    ie0a189d9450f7c12c239fba7ef90e148ed3a7c55a97b1a1b16dd24ddaf3af05d "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/sendmail"
    ieaadcc6e31da29f93473f27ed1f874889abcf9cdc2d42aad81085a18255791b8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/memberof/item/user/changepassword"
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
func (m *UserRequestBuilder) AssignLicense()(*i326394449db0a190f0a2fc474b2a490675d74783e92d65b49067d0d30f78fd22.AssignLicenseRequestBuilder) {
    return i326394449db0a190f0a2fc474b2a490675d74783e92d65b49067d0d30f78fd22.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*ieaadcc6e31da29f93473f27ed1f874889abcf9cdc2d42aad81085a18255791b8.ChangePasswordRequestBuilder) {
    return ieaadcc6e31da29f93473f27ed1f874889abcf9cdc2d42aad81085a18255791b8.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i9d8c1789bc721e6f87af390e2c121fef26320b3f4f4b02cdd3d3f92564822b7d.CheckMemberGroupsRequestBuilder) {
    return i9d8c1789bc721e6f87af390e2c121fef26320b3f4f4b02cdd3d3f92564822b7d.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*icc66bcf11474ad72318c4c646ba46e93f45c40b8ca05c9f0ce5bd068c5528b72.CheckMemberObjectsRequestBuilder) {
    return icc66bcf11474ad72318c4c646ba46e93f45c40b8ca05c9f0ce5bd068c5528b72.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/memberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportPersonalData()(*i59f17085687504d2eb6620b2058ff580e0aa57be0923a4a8ce3d844e01af02ef.ExportPersonalDataRequestBuilder) {
    return i59f17085687504d2eb6620b2058ff580e0aa57be0923a4a8ce3d844e01af02ef.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i662db12ac01d7038a5b93f0f1ab3c8e498be98d04a1dc41c7f29e8be9b3e2249.FindMeetingTimesRequestBuilder) {
    return i662db12ac01d7038a5b93f0f1ab3c8e498be98d04a1dc41c7f29e8be9b3e2249.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i7ad434fbe1ad63417dd1eb08e62642968ea0c3a31ea31f74c6c6c7b3593733ba.GetMailTipsRequestBuilder) {
    return i7ad434fbe1ad63417dd1eb08e62642968ea0c3a31ea31f74c6c6c7b3593733ba.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i94bbd38de045f07dd4203aaeaba941d25dc5107a8bc877d7be84e4f37851c888.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i94bbd38de045f07dd4203aaeaba941d25dc5107a8bc877d7be84e4f37851c888.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i59d9e41e5e6a3b4a93da15b3d40653e21265ab3d9ae0e17b04976db5512964ef.GetManagedAppPoliciesRequestBuilder) {
    return i59d9e41e5e6a3b4a93da15b3d40653e21265ab3d9ae0e17b04976db5512964ef.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i61cc66f6561fe8559fcabff44e7cf82443e56b35785bcaf90ff6484a4493bdc8.GetMemberGroupsRequestBuilder) {
    return i61cc66f6561fe8559fcabff44e7cf82443e56b35785bcaf90ff6484a4493bdc8.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i4e27058d8f79e692a1947cfb44e71bec11640181e9986ffcdeabaceb8c7af4c5.GetMemberObjectsRequestBuilder) {
    return i4e27058d8f79e692a1947cfb44e71bec11640181e9986ffcdeabaceb8c7af4c5.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i45de06b1f5a9abeed74b8c95e82c981451cd0ab816d7167324b763a72c339357.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i45de06b1f5a9abeed74b8c95e82c981451cd0ab816d7167324b763a72c339357.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i4b72766e0b8b1aa132b8a4f13dfc7b0058530145c4261d70e644554128b8e399.RemoveAllDevicesFromManagementRequestBuilder) {
    return i4b72766e0b8b1aa132b8a4f13dfc7b0058530145c4261d70e644554128b8e399.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i54ccdd10e1babdf7a720ac0f0ee6fc672fcb354e879b96f221442e162f52d852.ReprocessLicenseAssignmentRequestBuilder) {
    return i54ccdd10e1babdf7a720ac0f0ee6fc672fcb354e879b96f221442e162f52d852.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i7f352500093baef1b19a05301d65737715f88454230ea216b5487620e23f9bcb.RestoreRequestBuilder) {
    return i7f352500093baef1b19a05301d65737715f88454230ea216b5487620e23f9bcb.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i85c043082ce6971a9408255031ce028bcfcf317cd395b030c97bc4e333e1fd7e.RevokeSignInSessionsRequestBuilder) {
    return i85c043082ce6971a9408255031ce028bcfcf317cd395b030c97bc4e333e1fd7e.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*ie0a189d9450f7c12c239fba7ef90e148ed3a7c55a97b1a1b16dd24ddaf3af05d.SendMailRequestBuilder) {
    return ie0a189d9450f7c12c239fba7ef90e148ed3a7c55a97b1a1b16dd24ddaf3af05d.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*ia676b0c8452eaa27680a5f824494afcd62ca91248d5d1435c27962dafbd089ef.TranslateExchangeIdsRequestBuilder) {
    return ia676b0c8452eaa27680a5f824494afcd62ca91248d5d1435c27962dafbd089ef.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*ic2fd0cc125167b6efbade089bb78f3b71befeecf38130be5ce3fddc6c93e991c.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return ic2fd0cc125167b6efbade089bb78f3b71befeecf38130be5ce3fddc6c93e991c.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
