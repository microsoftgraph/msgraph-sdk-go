package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GroupsItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder provides operations to count the resources in the collection.
type GroupsItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupsItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewGroupsItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder) {
    m := &GroupsItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/children/{term%2Did}/children/{term%2Did1}/relations/$count";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewGroupsItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the number of the resource
func (m *GroupsItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "text/plain"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get the number of the resource
func (m *GroupsItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
