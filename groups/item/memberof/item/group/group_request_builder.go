package group

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i03181d1cf618c0807409aa724e7ab3227d5bb8f878b66a892899a58dfcfd2093 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberof/item/group/restore"
    i083e40cc10079622b79666573c2ba71604696ac7ec5baafe387bc8d9c6e1646e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberof/item/group/checkmemberobjects"
    i0ac56495301fbf585de6ec2f446f7b8cefb41464d5a1bf3915a2d898028201a8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberof/item/group/removefavorite"
    i0f59702538604271bc93e2b4ac6dad94f2c272f171e5ab8c49be5c1b7b4be6aa "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberof/item/group/resetunseencount"
    i1a8cc365e06256437dbe571b0141e3220b1aacb7292a2b0e546e3980a615e1a4 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberof/item/group/getmembergroups"
    i253fe37675760fd3d7457563f50128cab7b89191eea40e12ae3122f74d603dc9 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberof/item/group/unsubscribebymail"
    i30520783d26ba493ab8ce24df3c229b48ad166a7786eed04d9fc138b250d2f6e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberof/item/group/getmemberobjects"
    i32720636efcaf3cac9a04acddcb415b8c375732166938de7f03456fdf2933db1 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberof/item/group/subscribebymail"
    i4cc29a9bdd44b6f930dcca82400b0b3647dd0a8a886ff8d2db104dbb30e09d1a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberof/item/group/validateproperties"
    i696de52f2b63869661b2daf8c5cbdf2d09cdf295f79ed338ea519d946cf6ee60 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberof/item/group/renew"
    i805e8bd9283ccd3ef1c240427c9a5ef43af4ea21666de7e3509cf20c0b9559da "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberof/item/group/addfavorite"
    i8152472190a7e0a690ce35584c8ce15420a916c210817c9cac35a02d06968fed "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberof/item/group/checkmembergroups"
    ide445454aeb8413f8e8cd2a7c7084358fef03395e1dee97b0a8b8a9a97cb9b6a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberof/item/group/assignlicense"
    ie598f6696a6bb2845ee35b288438d46830fbe96111e144c4a2a372845220b1a0 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberof/item/group/checkgrantedpermissionsforapp"
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
func (m *GroupRequestBuilder) AddFavorite()(*i805e8bd9283ccd3ef1c240427c9a5ef43af4ea21666de7e3509cf20c0b9559da.AddFavoriteRequestBuilder) {
    return i805e8bd9283ccd3ef1c240427c9a5ef43af4ea21666de7e3509cf20c0b9559da.NewAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *GroupRequestBuilder) AssignLicense()(*ide445454aeb8413f8e8cd2a7c7084358fef03395e1dee97b0a8b8a9a97cb9b6a.AssignLicenseRequestBuilder) {
    return ide445454aeb8413f8e8cd2a7c7084358fef03395e1dee97b0a8b8a9a97cb9b6a.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckGrantedPermissionsForApp the checkGrantedPermissionsForApp property
func (m *GroupRequestBuilder) CheckGrantedPermissionsForApp()(*ie598f6696a6bb2845ee35b288438d46830fbe96111e144c4a2a372845220b1a0.CheckGrantedPermissionsForAppRequestBuilder) {
    return ie598f6696a6bb2845ee35b288438d46830fbe96111e144c4a2a372845220b1a0.NewCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *GroupRequestBuilder) CheckMemberGroups()(*i8152472190a7e0a690ce35584c8ce15420a916c210817c9cac35a02d06968fed.CheckMemberGroupsRequestBuilder) {
    return i8152472190a7e0a690ce35584c8ce15420a916c210817c9cac35a02d06968fed.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *GroupRequestBuilder) CheckMemberObjects()(*i083e40cc10079622b79666573c2ba71604696ac7ec5baafe387bc8d9c6e1646e.CheckMemberObjectsRequestBuilder) {
    return i083e40cc10079622b79666573c2ba71604696ac7ec5baafe387bc8d9c6e1646e.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    m := &GroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/memberOf/{directoryObject%2Did}/microsoft.graph.group{?%24select,%24expand}";
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
func (m *GroupRequestBuilder) GetMemberGroups()(*i1a8cc365e06256437dbe571b0141e3220b1aacb7292a2b0e546e3980a615e1a4.GetMemberGroupsRequestBuilder) {
    return i1a8cc365e06256437dbe571b0141e3220b1aacb7292a2b0e546e3980a615e1a4.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *GroupRequestBuilder) GetMemberObjects()(*i30520783d26ba493ab8ce24df3c229b48ad166a7786eed04d9fc138b250d2f6e.GetMemberObjectsRequestBuilder) {
    return i30520783d26ba493ab8ce24df3c229b48ad166a7786eed04d9fc138b250d2f6e.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *GroupRequestBuilder) RemoveFavorite()(*i0ac56495301fbf585de6ec2f446f7b8cefb41464d5a1bf3915a2d898028201a8.RemoveFavoriteRequestBuilder) {
    return i0ac56495301fbf585de6ec2f446f7b8cefb41464d5a1bf3915a2d898028201a8.NewRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Renew the renew property
func (m *GroupRequestBuilder) Renew()(*i696de52f2b63869661b2daf8c5cbdf2d09cdf295f79ed338ea519d946cf6ee60.RenewRequestBuilder) {
    return i696de52f2b63869661b2daf8c5cbdf2d09cdf295f79ed338ea519d946cf6ee60.NewRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetUnseenCount the resetUnseenCount property
func (m *GroupRequestBuilder) ResetUnseenCount()(*i0f59702538604271bc93e2b4ac6dad94f2c272f171e5ab8c49be5c1b7b4be6aa.ResetUnseenCountRequestBuilder) {
    return i0f59702538604271bc93e2b4ac6dad94f2c272f171e5ab8c49be5c1b7b4be6aa.NewResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *GroupRequestBuilder) Restore()(*i03181d1cf618c0807409aa724e7ab3227d5bb8f878b66a892899a58dfcfd2093.RestoreRequestBuilder) {
    return i03181d1cf618c0807409aa724e7ab3227d5bb8f878b66a892899a58dfcfd2093.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribeByMail the subscribeByMail property
func (m *GroupRequestBuilder) SubscribeByMail()(*i32720636efcaf3cac9a04acddcb415b8c375732166938de7f03456fdf2933db1.SubscribeByMailRequestBuilder) {
    return i32720636efcaf3cac9a04acddcb415b8c375732166938de7f03456fdf2933db1.NewSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsubscribeByMail the unsubscribeByMail property
func (m *GroupRequestBuilder) UnsubscribeByMail()(*i253fe37675760fd3d7457563f50128cab7b89191eea40e12ae3122f74d603dc9.UnsubscribeByMailRequestBuilder) {
    return i253fe37675760fd3d7457563f50128cab7b89191eea40e12ae3122f74d603dc9.NewUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *GroupRequestBuilder) ValidateProperties()(*i4cc29a9bdd44b6f930dcca82400b0b3647dd0a8a886ff8d2db104dbb30e09d1a.ValidatePropertiesRequestBuilder) {
    return i4cc29a9bdd44b6f930dcca82400b0b3647dd0a8a886ff8d2db104dbb30e09d1a.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
