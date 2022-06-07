package group

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i09d663178f88b97ff1ca3137bf89d8b330a495aca6bd041f645e83c47abc3c50 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberswithlicenseerrors/item/group/validateproperties"
    i0f7627f57ee9a74eacb4c8f78c3e8c5c416bbda0301a8dd934e184230c634d6d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberswithlicenseerrors/item/group/checkmemberobjects"
    i1de374b03b616d6c22a3ad798c70968d12ca4c947bd8167f4101a63e33f803de "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberswithlicenseerrors/item/group/checkgrantedpermissionsforapp"
    i2b27a6bdac9739e3fe8a32e70ed0341e088cdaee1c67a28c905de4f2a1445b34 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberswithlicenseerrors/item/group/unsubscribebymail"
    i4e2e4e3c9efe9781bc90193dc20090ffa20875d3e3f9b73f72b6ab259626816f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberswithlicenseerrors/item/group/resetunseencount"
    i7620daee6a9af165a10449aef368de20be61310fc3e6d692904562515372ee08 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberswithlicenseerrors/item/group/subscribebymail"
    i88ca3e4527e11449127bb11821828aae7b8ff3d97085a27f0adf3807055f062e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberswithlicenseerrors/item/group/removefavorite"
    ia8e04ec09ff83b8f0eec18fc3a66128d613407a0794875872ce0eb386fb25911 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberswithlicenseerrors/item/group/renew"
    iba8956f722f0acc825002a84e2f0ab40ffd90d8ef3843cde583f2e88b1171db9 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberswithlicenseerrors/item/group/restore"
    ic8e64db8b92ea5f57fe1e57a17f7d5a384a9987c5642e8c0d07431084d612109 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberswithlicenseerrors/item/group/assignlicense"
    id9605d363735a7a78dcd2423cc9216739b41f2ef276af0d65c2dcffa49550d68 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberswithlicenseerrors/item/group/checkmembergroups"
    idca07e4305abff078f3de4d53111f5f903e2240c9130ec1f019445d97522ff90 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberswithlicenseerrors/item/group/addfavorite"
    ieadceb79a3e070a2137a5d77c269c3e7af1753763621696da9ddb6c766da36a2 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberswithlicenseerrors/item/group/getmemberobjects"
    iefd7b4492059d1b5e796f8e08f52ff725b8a9e473e99c4afb444e3cea719dd6b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/memberswithlicenseerrors/item/group/getmembergroups"
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
func (m *GroupRequestBuilder) AddFavorite()(*idca07e4305abff078f3de4d53111f5f903e2240c9130ec1f019445d97522ff90.AddFavoriteRequestBuilder) {
    return idca07e4305abff078f3de4d53111f5f903e2240c9130ec1f019445d97522ff90.NewAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *GroupRequestBuilder) AssignLicense()(*ic8e64db8b92ea5f57fe1e57a17f7d5a384a9987c5642e8c0d07431084d612109.AssignLicenseRequestBuilder) {
    return ic8e64db8b92ea5f57fe1e57a17f7d5a384a9987c5642e8c0d07431084d612109.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckGrantedPermissionsForApp the checkGrantedPermissionsForApp property
func (m *GroupRequestBuilder) CheckGrantedPermissionsForApp()(*i1de374b03b616d6c22a3ad798c70968d12ca4c947bd8167f4101a63e33f803de.CheckGrantedPermissionsForAppRequestBuilder) {
    return i1de374b03b616d6c22a3ad798c70968d12ca4c947bd8167f4101a63e33f803de.NewCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *GroupRequestBuilder) CheckMemberGroups()(*id9605d363735a7a78dcd2423cc9216739b41f2ef276af0d65c2dcffa49550d68.CheckMemberGroupsRequestBuilder) {
    return id9605d363735a7a78dcd2423cc9216739b41f2ef276af0d65c2dcffa49550d68.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *GroupRequestBuilder) CheckMemberObjects()(*i0f7627f57ee9a74eacb4c8f78c3e8c5c416bbda0301a8dd934e184230c634d6d.CheckMemberObjectsRequestBuilder) {
    return i0f7627f57ee9a74eacb4c8f78c3e8c5c416bbda0301a8dd934e184230c634d6d.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    m := &GroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/membersWithLicenseErrors/{directoryObject%2Did}/microsoft.graph.group{?%24select,%24expand}";
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
func (m *GroupRequestBuilder) GetMemberGroups()(*iefd7b4492059d1b5e796f8e08f52ff725b8a9e473e99c4afb444e3cea719dd6b.GetMemberGroupsRequestBuilder) {
    return iefd7b4492059d1b5e796f8e08f52ff725b8a9e473e99c4afb444e3cea719dd6b.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *GroupRequestBuilder) GetMemberObjects()(*ieadceb79a3e070a2137a5d77c269c3e7af1753763621696da9ddb6c766da36a2.GetMemberObjectsRequestBuilder) {
    return ieadceb79a3e070a2137a5d77c269c3e7af1753763621696da9ddb6c766da36a2.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *GroupRequestBuilder) RemoveFavorite()(*i88ca3e4527e11449127bb11821828aae7b8ff3d97085a27f0adf3807055f062e.RemoveFavoriteRequestBuilder) {
    return i88ca3e4527e11449127bb11821828aae7b8ff3d97085a27f0adf3807055f062e.NewRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Renew the renew property
func (m *GroupRequestBuilder) Renew()(*ia8e04ec09ff83b8f0eec18fc3a66128d613407a0794875872ce0eb386fb25911.RenewRequestBuilder) {
    return ia8e04ec09ff83b8f0eec18fc3a66128d613407a0794875872ce0eb386fb25911.NewRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetUnseenCount the resetUnseenCount property
func (m *GroupRequestBuilder) ResetUnseenCount()(*i4e2e4e3c9efe9781bc90193dc20090ffa20875d3e3f9b73f72b6ab259626816f.ResetUnseenCountRequestBuilder) {
    return i4e2e4e3c9efe9781bc90193dc20090ffa20875d3e3f9b73f72b6ab259626816f.NewResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *GroupRequestBuilder) Restore()(*iba8956f722f0acc825002a84e2f0ab40ffd90d8ef3843cde583f2e88b1171db9.RestoreRequestBuilder) {
    return iba8956f722f0acc825002a84e2f0ab40ffd90d8ef3843cde583f2e88b1171db9.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribeByMail the subscribeByMail property
func (m *GroupRequestBuilder) SubscribeByMail()(*i7620daee6a9af165a10449aef368de20be61310fc3e6d692904562515372ee08.SubscribeByMailRequestBuilder) {
    return i7620daee6a9af165a10449aef368de20be61310fc3e6d692904562515372ee08.NewSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsubscribeByMail the unsubscribeByMail property
func (m *GroupRequestBuilder) UnsubscribeByMail()(*i2b27a6bdac9739e3fe8a32e70ed0341e088cdaee1c67a28c905de4f2a1445b34.UnsubscribeByMailRequestBuilder) {
    return i2b27a6bdac9739e3fe8a32e70ed0341e088cdaee1c67a28c905de4f2a1445b34.NewUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *GroupRequestBuilder) ValidateProperties()(*i09d663178f88b97ff1ca3137bf89d8b330a495aca6bd041f645e83c47abc3c50.ValidatePropertiesRequestBuilder) {
    return i09d663178f88b97ff1ca3137bf89d8b330a495aca6bd041f645e83c47abc3c50.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
