package group

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i036314544fa6afaf43a05ad63f0244fae2c8a5d48c26e1b308cbca3ac104e74f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/item/group/checkmemberobjects"
    i199d58dae82bb0e1b033cd701e9b835eb8d63a139bac54a07ebc2642fa18b840 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/item/group/checkgrantedpermissionsforapp"
    i31e35f8d974071de8a684d36c6f96e8b682ee5bb32534578e22fc7f759511fb6 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/item/group/assignlicense"
    i4b46a90ea736a649ef0f4b81aaa0659a1b77da42bd63b79d7bd93483b436893b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/item/group/checkmembergroups"
    i7358bc6c3e20576d674984b0b660dffe4e6f9806d9d11daf9c5a19125aa7c2ca "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/item/group/addfavorite"
    i7c04c49c0b6cd51c322cde7d2d2e816a948926347a9fe1fe12f060cb41cba348 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/item/group/unsubscribebymail"
    i845440e59a19e6857afd6797d3898f2f4ec35a3bbf17a89d055d978bbed9edea "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/item/group/restore"
    i88f0fd8d7e6f9dd48e796c2a3de670f50635f4e4958e08c1a6ddedaa706b9c7a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/item/group/subscribebymail"
    i951f97794b5580e43be237bbce8dd19405286ef6c163349492e718df4278bcec "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/item/group/renew"
    i9e7ab46a25d34ded89090283e7ba418873ab3320b01f0b9e5c5aa6bb21ee159e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/item/group/resetunseencount"
    iaa8126c3b14076e9ea388e6ecd3aa4f47604ca508742ea0504e345e166bbefd2 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/item/group/getmembergroups"
    ib813aef2ae48ce63d6b75361bcc1d9f3cdcb6ec1fa6cca6c674a25523d93e740 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/item/group/removefavorite"
    ic8607685c6ce251b1879e3858441a1e8e7cc37a0108220a9be825bde29190fc8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/item/group/getmemberobjects"
    iefbf78819eea61805fd04744a6a7a8c065560af11ddfb4e6152d21b4097488b8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/item/group/validateproperties"
)

// GroupRequestBuilder casts the previous resource to group.
type GroupRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.group
type GroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupRequestBuilderGetQueryParameters
}
// AddFavorite the addFavorite property
func (m *GroupRequestBuilder) AddFavorite()(*i7358bc6c3e20576d674984b0b660dffe4e6f9806d9d11daf9c5a19125aa7c2ca.AddFavoriteRequestBuilder) {
    return i7358bc6c3e20576d674984b0b660dffe4e6f9806d9d11daf9c5a19125aa7c2ca.NewAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *GroupRequestBuilder) AssignLicense()(*i31e35f8d974071de8a684d36c6f96e8b682ee5bb32534578e22fc7f759511fb6.AssignLicenseRequestBuilder) {
    return i31e35f8d974071de8a684d36c6f96e8b682ee5bb32534578e22fc7f759511fb6.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckGrantedPermissionsForApp the checkGrantedPermissionsForApp property
func (m *GroupRequestBuilder) CheckGrantedPermissionsForApp()(*i199d58dae82bb0e1b033cd701e9b835eb8d63a139bac54a07ebc2642fa18b840.CheckGrantedPermissionsForAppRequestBuilder) {
    return i199d58dae82bb0e1b033cd701e9b835eb8d63a139bac54a07ebc2642fa18b840.NewCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *GroupRequestBuilder) CheckMemberGroups()(*i4b46a90ea736a649ef0f4b81aaa0659a1b77da42bd63b79d7bd93483b436893b.CheckMemberGroupsRequestBuilder) {
    return i4b46a90ea736a649ef0f4b81aaa0659a1b77da42bd63b79d7bd93483b436893b.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *GroupRequestBuilder) CheckMemberObjects()(*i036314544fa6afaf43a05ad63f0244fae2c8a5d48c26e1b308cbca3ac104e74f.CheckMemberObjectsRequestBuilder) {
    return i036314544fa6afaf43a05ad63f0244fae2c8a5d48c26e1b308cbca3ac104e74f.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    m := &GroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/owners/{directoryObject%2Did}/microsoft.graph.group{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupRequestBuilder instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.group
func (m *GroupRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get the item of type microsoft.graph.directoryObject as microsoft.graph.group
func (m *GroupRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.group
func (m *GroupRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *GroupRequestBuilder) GetMemberGroups()(*iaa8126c3b14076e9ea388e6ecd3aa4f47604ca508742ea0504e345e166bbefd2.GetMemberGroupsRequestBuilder) {
    return iaa8126c3b14076e9ea388e6ecd3aa4f47604ca508742ea0504e345e166bbefd2.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *GroupRequestBuilder) GetMemberObjects()(*ic8607685c6ce251b1879e3858441a1e8e7cc37a0108220a9be825bde29190fc8.GetMemberObjectsRequestBuilder) {
    return ic8607685c6ce251b1879e3858441a1e8e7cc37a0108220a9be825bde29190fc8.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler get the item of type microsoft.graph.directoryObject as microsoft.graph.group
func (m *GroupRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GroupRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateGroupFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable), nil
}
// RemoveFavorite the removeFavorite property
func (m *GroupRequestBuilder) RemoveFavorite()(*ib813aef2ae48ce63d6b75361bcc1d9f3cdcb6ec1fa6cca6c674a25523d93e740.RemoveFavoriteRequestBuilder) {
    return ib813aef2ae48ce63d6b75361bcc1d9f3cdcb6ec1fa6cca6c674a25523d93e740.NewRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Renew the renew property
func (m *GroupRequestBuilder) Renew()(*i951f97794b5580e43be237bbce8dd19405286ef6c163349492e718df4278bcec.RenewRequestBuilder) {
    return i951f97794b5580e43be237bbce8dd19405286ef6c163349492e718df4278bcec.NewRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetUnseenCount the resetUnseenCount property
func (m *GroupRequestBuilder) ResetUnseenCount()(*i9e7ab46a25d34ded89090283e7ba418873ab3320b01f0b9e5c5aa6bb21ee159e.ResetUnseenCountRequestBuilder) {
    return i9e7ab46a25d34ded89090283e7ba418873ab3320b01f0b9e5c5aa6bb21ee159e.NewResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *GroupRequestBuilder) Restore()(*i845440e59a19e6857afd6797d3898f2f4ec35a3bbf17a89d055d978bbed9edea.RestoreRequestBuilder) {
    return i845440e59a19e6857afd6797d3898f2f4ec35a3bbf17a89d055d978bbed9edea.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribeByMail the subscribeByMail property
func (m *GroupRequestBuilder) SubscribeByMail()(*i88f0fd8d7e6f9dd48e796c2a3de670f50635f4e4958e08c1a6ddedaa706b9c7a.SubscribeByMailRequestBuilder) {
    return i88f0fd8d7e6f9dd48e796c2a3de670f50635f4e4958e08c1a6ddedaa706b9c7a.NewSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsubscribeByMail the unsubscribeByMail property
func (m *GroupRequestBuilder) UnsubscribeByMail()(*i7c04c49c0b6cd51c322cde7d2d2e816a948926347a9fe1fe12f060cb41cba348.UnsubscribeByMailRequestBuilder) {
    return i7c04c49c0b6cd51c322cde7d2d2e816a948926347a9fe1fe12f060cb41cba348.NewUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *GroupRequestBuilder) ValidateProperties()(*iefbf78819eea61805fd04744a6a7a8c065560af11ddfb4e6152d21b4097488b8.ValidatePropertiesRequestBuilder) {
    return iefbf78819eea61805fd04744a6a7a8c065560af11ddfb4e6152d21b4097488b8.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
