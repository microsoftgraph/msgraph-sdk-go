package group

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i03b9e39cf33ffa2ba457a8107e1237d34c795591373ec5ad5080324faab9a90c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivemembers/item/group/unsubscribebymail"
    i27c56f2bece4e3c31cdd923cf87dad529009ea861b8aa703e45002eacc688f3a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivemembers/item/group/resetunseencount"
    i2edba13c1535bab4b07a153bfd9e54c32fffc7787bde08543ee23e82bdd3fff8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivemembers/item/group/subscribebymail"
    i3e26308c324717037a5d1ed07325316e39b1099055ed05b9cbb274eda85a5512 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivemembers/item/group/validateproperties"
    i4d00d7e745c6dc2c4ef9394130e0442bf715d8a11171b080a3dd627bd963d433 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivemembers/item/group/checkmemberobjects"
    i50d699751762a9f20f96165fff40aebc52601ae379a4f3bd9e75346ec04f747a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivemembers/item/group/getmembergroups"
    i51569a2b789ad9ad11b3fd047e619cd19b835d10d490c309127717df98a22a43 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivemembers/item/group/checkmembergroups"
    i7c6396a963d61d64ae9408cb4cc1b00bc3eb992263a080c8554f54ee65e6810d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivemembers/item/group/renew"
    i81f8bedb7517eb2d205e1a609e215dafc2d235817b5f46a4c98e4158bb5127dd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivemembers/item/group/addfavorite"
    ib2915da8c20423a31428493de4177790b9c1e566719b3cc384ee4c0b8483c18e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivemembers/item/group/getmemberobjects"
    ib2b6152940f6b1669203e891d29a07c1105c922268210f9f6f4dabc57a9895a5 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivemembers/item/group/assignlicense"
    id1177d98fcc5af528d5d04788a6023cdf708385fe01c1381ce2a7bd390d1a6e1 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivemembers/item/group/checkgrantedpermissionsforapp"
    ie2203924d441e539b2810b318c6c2531ae66a5f7b42e2759a510e33b3080205f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivemembers/item/group/restore"
    ie5b5dd2c984ac06427580dfd08a840b015aaa6fb5773791e80c5f17b22f48c4a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivemembers/item/group/removefavorite"
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
func (m *GroupRequestBuilder) AddFavorite()(*i81f8bedb7517eb2d205e1a609e215dafc2d235817b5f46a4c98e4158bb5127dd.AddFavoriteRequestBuilder) {
    return i81f8bedb7517eb2d205e1a609e215dafc2d235817b5f46a4c98e4158bb5127dd.NewAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *GroupRequestBuilder) AssignLicense()(*ib2b6152940f6b1669203e891d29a07c1105c922268210f9f6f4dabc57a9895a5.AssignLicenseRequestBuilder) {
    return ib2b6152940f6b1669203e891d29a07c1105c922268210f9f6f4dabc57a9895a5.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckGrantedPermissionsForApp the checkGrantedPermissionsForApp property
func (m *GroupRequestBuilder) CheckGrantedPermissionsForApp()(*id1177d98fcc5af528d5d04788a6023cdf708385fe01c1381ce2a7bd390d1a6e1.CheckGrantedPermissionsForAppRequestBuilder) {
    return id1177d98fcc5af528d5d04788a6023cdf708385fe01c1381ce2a7bd390d1a6e1.NewCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *GroupRequestBuilder) CheckMemberGroups()(*i51569a2b789ad9ad11b3fd047e619cd19b835d10d490c309127717df98a22a43.CheckMemberGroupsRequestBuilder) {
    return i51569a2b789ad9ad11b3fd047e619cd19b835d10d490c309127717df98a22a43.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *GroupRequestBuilder) CheckMemberObjects()(*i4d00d7e745c6dc2c4ef9394130e0442bf715d8a11171b080a3dd627bd963d433.CheckMemberObjectsRequestBuilder) {
    return i4d00d7e745c6dc2c4ef9394130e0442bf715d8a11171b080a3dd627bd963d433.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    m := &GroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/transitiveMembers/{directoryObject%2Did}/microsoft.graph.group{?%24select,%24expand}";
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
func (m *GroupRequestBuilder) GetMemberGroups()(*i50d699751762a9f20f96165fff40aebc52601ae379a4f3bd9e75346ec04f747a.GetMemberGroupsRequestBuilder) {
    return i50d699751762a9f20f96165fff40aebc52601ae379a4f3bd9e75346ec04f747a.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *GroupRequestBuilder) GetMemberObjects()(*ib2915da8c20423a31428493de4177790b9c1e566719b3cc384ee4c0b8483c18e.GetMemberObjectsRequestBuilder) {
    return ib2915da8c20423a31428493de4177790b9c1e566719b3cc384ee4c0b8483c18e.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *GroupRequestBuilder) RemoveFavorite()(*ie5b5dd2c984ac06427580dfd08a840b015aaa6fb5773791e80c5f17b22f48c4a.RemoveFavoriteRequestBuilder) {
    return ie5b5dd2c984ac06427580dfd08a840b015aaa6fb5773791e80c5f17b22f48c4a.NewRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Renew the renew property
func (m *GroupRequestBuilder) Renew()(*i7c6396a963d61d64ae9408cb4cc1b00bc3eb992263a080c8554f54ee65e6810d.RenewRequestBuilder) {
    return i7c6396a963d61d64ae9408cb4cc1b00bc3eb992263a080c8554f54ee65e6810d.NewRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetUnseenCount the resetUnseenCount property
func (m *GroupRequestBuilder) ResetUnseenCount()(*i27c56f2bece4e3c31cdd923cf87dad529009ea861b8aa703e45002eacc688f3a.ResetUnseenCountRequestBuilder) {
    return i27c56f2bece4e3c31cdd923cf87dad529009ea861b8aa703e45002eacc688f3a.NewResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *GroupRequestBuilder) Restore()(*ie2203924d441e539b2810b318c6c2531ae66a5f7b42e2759a510e33b3080205f.RestoreRequestBuilder) {
    return ie2203924d441e539b2810b318c6c2531ae66a5f7b42e2759a510e33b3080205f.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribeByMail the subscribeByMail property
func (m *GroupRequestBuilder) SubscribeByMail()(*i2edba13c1535bab4b07a153bfd9e54c32fffc7787bde08543ee23e82bdd3fff8.SubscribeByMailRequestBuilder) {
    return i2edba13c1535bab4b07a153bfd9e54c32fffc7787bde08543ee23e82bdd3fff8.NewSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsubscribeByMail the unsubscribeByMail property
func (m *GroupRequestBuilder) UnsubscribeByMail()(*i03b9e39cf33ffa2ba457a8107e1237d34c795591373ec5ad5080324faab9a90c.UnsubscribeByMailRequestBuilder) {
    return i03b9e39cf33ffa2ba457a8107e1237d34c795591373ec5ad5080324faab9a90c.NewUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *GroupRequestBuilder) ValidateProperties()(*i3e26308c324717037a5d1ed07325316e39b1099055ed05b9cbb274eda85a5512.ValidatePropertiesRequestBuilder) {
    return i3e26308c324717037a5d1ed07325316e39b1099055ed05b9cbb274eda85a5512.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
