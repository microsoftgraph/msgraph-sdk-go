package group

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0289889d7f7cafbe71d06ce1f8309f0d9f69876bb1f1d8a2e3d3b3d6d8363a43 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivememberof/item/group/getmemberobjects"
    i134d2e492131bf79cb66261f143677ce9a0c477eff55feeb83f815812fbd6350 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivememberof/item/group/assignlicense"
    i329dcd9a020e5f099745cf75effd0b0c234c5576435e49794c08aeaab2d23de2 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivememberof/item/group/checkmemberobjects"
    i37b867b22a7cb19bd8602352d592586464b7b44e58b8a4d30a4a427cb41d3e0e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivememberof/item/group/renew"
    i4a99b4fd2fd2c9a65cc24b42f16180f66c2e288602bc045506b5507ec2c3126e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivememberof/item/group/getmembergroups"
    i4acf6a5bc982008105b83bb668ad775931b24ee13374c6750bd6001e79cc3bfe "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivememberof/item/group/subscribebymail"
    i8532092b84d776df8fd75430eeecf433110a42887c6db0c9aae45fe5c2522a34 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivememberof/item/group/checkmembergroups"
    i955193f439ffa50835f6594c3fd8b62f6ea1bd37b7f1dad19fec21fe46aefb80 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivememberof/item/group/checkgrantedpermissionsforapp"
    ic607ff04ba35edb01feae3bbafe1e97044e4b8a12da9bb57a49bafa3123316b1 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivememberof/item/group/removefavorite"
    id9d2b0577ea08383f56b5dbb3b6553659e4d10b60b9336057bac487bc2f52c1c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivememberof/item/group/addfavorite"
    ie180e4894c2301720dd0d5715cdcc743a2ce2561333e69b0381b9a548f599226 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivememberof/item/group/unsubscribebymail"
    if37dcec580f3bf51bd363294b922d266ad98d63cd5fa5fc6bedc5618aa194775 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivememberof/item/group/resetunseencount"
    if82dfd618a39fd9510cb2e311b92293cde7be0df73e79887c2d1b452beec6c7b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivememberof/item/group/validateproperties"
    iff9f70c8e0188266a0058aafe9028b73e761bb28f509644dd1d6f8fb21a1b7cd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivememberof/item/group/restore"
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
func (m *GroupRequestBuilder) AddFavorite()(*id9d2b0577ea08383f56b5dbb3b6553659e4d10b60b9336057bac487bc2f52c1c.AddFavoriteRequestBuilder) {
    return id9d2b0577ea08383f56b5dbb3b6553659e4d10b60b9336057bac487bc2f52c1c.NewAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *GroupRequestBuilder) AssignLicense()(*i134d2e492131bf79cb66261f143677ce9a0c477eff55feeb83f815812fbd6350.AssignLicenseRequestBuilder) {
    return i134d2e492131bf79cb66261f143677ce9a0c477eff55feeb83f815812fbd6350.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckGrantedPermissionsForApp the checkGrantedPermissionsForApp property
func (m *GroupRequestBuilder) CheckGrantedPermissionsForApp()(*i955193f439ffa50835f6594c3fd8b62f6ea1bd37b7f1dad19fec21fe46aefb80.CheckGrantedPermissionsForAppRequestBuilder) {
    return i955193f439ffa50835f6594c3fd8b62f6ea1bd37b7f1dad19fec21fe46aefb80.NewCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *GroupRequestBuilder) CheckMemberGroups()(*i8532092b84d776df8fd75430eeecf433110a42887c6db0c9aae45fe5c2522a34.CheckMemberGroupsRequestBuilder) {
    return i8532092b84d776df8fd75430eeecf433110a42887c6db0c9aae45fe5c2522a34.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *GroupRequestBuilder) CheckMemberObjects()(*i329dcd9a020e5f099745cf75effd0b0c234c5576435e49794c08aeaab2d23de2.CheckMemberObjectsRequestBuilder) {
    return i329dcd9a020e5f099745cf75effd0b0c234c5576435e49794c08aeaab2d23de2.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    m := &GroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.group{?%24select,%24expand}";
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
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.group
func (m *GroupRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *GroupRequestBuilder) GetMemberGroups()(*i4a99b4fd2fd2c9a65cc24b42f16180f66c2e288602bc045506b5507ec2c3126e.GetMemberGroupsRequestBuilder) {
    return i4a99b4fd2fd2c9a65cc24b42f16180f66c2e288602bc045506b5507ec2c3126e.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *GroupRequestBuilder) GetMemberObjects()(*i0289889d7f7cafbe71d06ce1f8309f0d9f69876bb1f1d8a2e3d3b3d6d8363a43.GetMemberObjectsRequestBuilder) {
    return i0289889d7f7cafbe71d06ce1f8309f0d9f69876bb1f1d8a2e3d3b3d6d8363a43.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *GroupRequestBuilder) RemoveFavorite()(*ic607ff04ba35edb01feae3bbafe1e97044e4b8a12da9bb57a49bafa3123316b1.RemoveFavoriteRequestBuilder) {
    return ic607ff04ba35edb01feae3bbafe1e97044e4b8a12da9bb57a49bafa3123316b1.NewRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Renew the renew property
func (m *GroupRequestBuilder) Renew()(*i37b867b22a7cb19bd8602352d592586464b7b44e58b8a4d30a4a427cb41d3e0e.RenewRequestBuilder) {
    return i37b867b22a7cb19bd8602352d592586464b7b44e58b8a4d30a4a427cb41d3e0e.NewRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetUnseenCount the resetUnseenCount property
func (m *GroupRequestBuilder) ResetUnseenCount()(*if37dcec580f3bf51bd363294b922d266ad98d63cd5fa5fc6bedc5618aa194775.ResetUnseenCountRequestBuilder) {
    return if37dcec580f3bf51bd363294b922d266ad98d63cd5fa5fc6bedc5618aa194775.NewResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *GroupRequestBuilder) Restore()(*iff9f70c8e0188266a0058aafe9028b73e761bb28f509644dd1d6f8fb21a1b7cd.RestoreRequestBuilder) {
    return iff9f70c8e0188266a0058aafe9028b73e761bb28f509644dd1d6f8fb21a1b7cd.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribeByMail the subscribeByMail property
func (m *GroupRequestBuilder) SubscribeByMail()(*i4acf6a5bc982008105b83bb668ad775931b24ee13374c6750bd6001e79cc3bfe.SubscribeByMailRequestBuilder) {
    return i4acf6a5bc982008105b83bb668ad775931b24ee13374c6750bd6001e79cc3bfe.NewSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsubscribeByMail the unsubscribeByMail property
func (m *GroupRequestBuilder) UnsubscribeByMail()(*ie180e4894c2301720dd0d5715cdcc743a2ce2561333e69b0381b9a548f599226.UnsubscribeByMailRequestBuilder) {
    return ie180e4894c2301720dd0d5715cdcc743a2ce2561333e69b0381b9a548f599226.NewUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *GroupRequestBuilder) ValidateProperties()(*if82dfd618a39fd9510cb2e311b92293cde7be0df73e79887c2d1b452beec6c7b.ValidatePropertiesRequestBuilder) {
    return if82dfd618a39fd9510cb2e311b92293cde7be0df73e79887c2d1b452beec6c7b.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
