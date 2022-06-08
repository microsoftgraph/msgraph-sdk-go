package group

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i043c9450388a10671f7e3e61e161be55bdf8bbb29eabc2be372abfdfcd7ccf05 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/members/item/group/resetunseencount"
    i08ba4eb40c63695b4eb8e07e5b49201fd8e485871b3ace93b9810652b8465c12 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/members/item/group/checkgrantedpermissionsforapp"
    i0d5137e949d97e392e7b443f7449a0732f6afbe7da44707852fb7f47e19a88a4 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/members/item/group/validateproperties"
    i0f975e5874fa1706d07aa5a08d71664aa5df289872cef0505f2766e1bb1ffb3f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/members/item/group/addfavorite"
    i19557340a69d1ba626938aede2fe3711a5ba04b8e7ea3b6597e7403973ce636c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/members/item/group/getmemberobjects"
    i21315d4319b67a9501a29f1c594e6b07153693b74a34b29b0455446ae20da529 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/members/item/group/getmembergroups"
    i2177bc162e5cad209fce97a4e5c1cc016301ada22360316fea6d4556a5f07d11 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/members/item/group/assignlicense"
    i4dd53a0dae0e79b39f38d76376d9fdc283f357bb9dc7bc1b22bcff517c09e50c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/members/item/group/restore"
    i850813c4b5d8df96f6a8db6f448603c055b2142ceb010ba2d96df996181e8ac6 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/members/item/group/checkmemberobjects"
    i92c375fe0d0e9fa961f0e46fbdff0b5a9d445035b79bd4ae5509279b573d0a68 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/members/item/group/subscribebymail"
    i94c32ee903d8842c113067be39e6c64b2d10c011a9f47fc930e812256ef69538 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/members/item/group/removefavorite"
    i9f35685b5b691577dfe36e8d00707997537b734fc9455c30e4a6b94a09070fed "github.com/microsoftgraph/msgraph-sdk-go/groups/item/members/item/group/renew"
    ib9fa47507f892503ee392d1efc9f539d9c03559ef368b1cb54ee0c2bb5129a18 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/members/item/group/checkmembergroups"
    ifb5a57f4327933d987eadf3a9bbf4cb6d74ce02c81081933d43912c65e5ad956 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/members/item/group/unsubscribebymail"
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
func (m *GroupRequestBuilder) AddFavorite()(*i0f975e5874fa1706d07aa5a08d71664aa5df289872cef0505f2766e1bb1ffb3f.AddFavoriteRequestBuilder) {
    return i0f975e5874fa1706d07aa5a08d71664aa5df289872cef0505f2766e1bb1ffb3f.NewAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *GroupRequestBuilder) AssignLicense()(*i2177bc162e5cad209fce97a4e5c1cc016301ada22360316fea6d4556a5f07d11.AssignLicenseRequestBuilder) {
    return i2177bc162e5cad209fce97a4e5c1cc016301ada22360316fea6d4556a5f07d11.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckGrantedPermissionsForApp the checkGrantedPermissionsForApp property
func (m *GroupRequestBuilder) CheckGrantedPermissionsForApp()(*i08ba4eb40c63695b4eb8e07e5b49201fd8e485871b3ace93b9810652b8465c12.CheckGrantedPermissionsForAppRequestBuilder) {
    return i08ba4eb40c63695b4eb8e07e5b49201fd8e485871b3ace93b9810652b8465c12.NewCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *GroupRequestBuilder) CheckMemberGroups()(*ib9fa47507f892503ee392d1efc9f539d9c03559ef368b1cb54ee0c2bb5129a18.CheckMemberGroupsRequestBuilder) {
    return ib9fa47507f892503ee392d1efc9f539d9c03559ef368b1cb54ee0c2bb5129a18.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *GroupRequestBuilder) CheckMemberObjects()(*i850813c4b5d8df96f6a8db6f448603c055b2142ceb010ba2d96df996181e8ac6.CheckMemberObjectsRequestBuilder) {
    return i850813c4b5d8df96f6a8db6f448603c055b2142ceb010ba2d96df996181e8ac6.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    m := &GroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/members/{directoryObject%2Did}/microsoft.graph.group{?%24select,%24expand}";
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
func (m *GroupRequestBuilder) GetMemberGroups()(*i21315d4319b67a9501a29f1c594e6b07153693b74a34b29b0455446ae20da529.GetMemberGroupsRequestBuilder) {
    return i21315d4319b67a9501a29f1c594e6b07153693b74a34b29b0455446ae20da529.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *GroupRequestBuilder) GetMemberObjects()(*i19557340a69d1ba626938aede2fe3711a5ba04b8e7ea3b6597e7403973ce636c.GetMemberObjectsRequestBuilder) {
    return i19557340a69d1ba626938aede2fe3711a5ba04b8e7ea3b6597e7403973ce636c.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *GroupRequestBuilder) RemoveFavorite()(*i94c32ee903d8842c113067be39e6c64b2d10c011a9f47fc930e812256ef69538.RemoveFavoriteRequestBuilder) {
    return i94c32ee903d8842c113067be39e6c64b2d10c011a9f47fc930e812256ef69538.NewRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Renew the renew property
func (m *GroupRequestBuilder) Renew()(*i9f35685b5b691577dfe36e8d00707997537b734fc9455c30e4a6b94a09070fed.RenewRequestBuilder) {
    return i9f35685b5b691577dfe36e8d00707997537b734fc9455c30e4a6b94a09070fed.NewRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetUnseenCount the resetUnseenCount property
func (m *GroupRequestBuilder) ResetUnseenCount()(*i043c9450388a10671f7e3e61e161be55bdf8bbb29eabc2be372abfdfcd7ccf05.ResetUnseenCountRequestBuilder) {
    return i043c9450388a10671f7e3e61e161be55bdf8bbb29eabc2be372abfdfcd7ccf05.NewResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *GroupRequestBuilder) Restore()(*i4dd53a0dae0e79b39f38d76376d9fdc283f357bb9dc7bc1b22bcff517c09e50c.RestoreRequestBuilder) {
    return i4dd53a0dae0e79b39f38d76376d9fdc283f357bb9dc7bc1b22bcff517c09e50c.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribeByMail the subscribeByMail property
func (m *GroupRequestBuilder) SubscribeByMail()(*i92c375fe0d0e9fa961f0e46fbdff0b5a9d445035b79bd4ae5509279b573d0a68.SubscribeByMailRequestBuilder) {
    return i92c375fe0d0e9fa961f0e46fbdff0b5a9d445035b79bd4ae5509279b573d0a68.NewSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsubscribeByMail the unsubscribeByMail property
func (m *GroupRequestBuilder) UnsubscribeByMail()(*ifb5a57f4327933d987eadf3a9bbf4cb6d74ce02c81081933d43912c65e5ad956.UnsubscribeByMailRequestBuilder) {
    return ifb5a57f4327933d987eadf3a9bbf4cb6d74ce02c81081933d43912c65e5ad956.NewUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *GroupRequestBuilder) ValidateProperties()(*i0d5137e949d97e392e7b443f7449a0732f6afbe7da44707852fb7f47e19a88a4.ValidatePropertiesRequestBuilder) {
    return i0d5137e949d97e392e7b443f7449a0732f6afbe7da44707852fb7f47e19a88a4.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
