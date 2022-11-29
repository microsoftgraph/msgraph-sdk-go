package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GroupsItemSitesSiteItemRequestBuilder provides operations to manage the sites property of the microsoft.graph.group entity.
type GroupsItemSitesSiteItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemSitesSiteItemRequestBuilderGetQueryParameters the list of SharePoint sites in this group. Access the default site with /sites/root.
type GroupsItemSitesSiteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsItemSitesSiteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemSitesSiteItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsItemSitesSiteItemRequestBuilderGetQueryParameters
}
// GroupsItemSitesSiteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemSitesSiteItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) Analytics()(*GroupsItemSitesItemAnalyticsRequestBuilder) {
    return NewGroupsItemSitesItemAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Columns provides operations to manage the columns property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) Columns()(*GroupsItemSitesItemColumnsRequestBuilder) {
    return NewGroupsItemSitesItemColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById provides operations to manage the columns property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) ColumnsById(id string)(*GroupsItemSitesItemColumnsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewGroupsItemSitesItemColumnsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewGroupsItemSitesSiteItemRequestBuilderInternal instantiates a new SiteItemRequestBuilder and sets the default values.
func NewGroupsItemSitesSiteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemSitesSiteItemRequestBuilder) {
    m := &GroupsItemSitesSiteItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemSitesSiteItemRequestBuilder instantiates a new SiteItemRequestBuilder and sets the default values.
func NewGroupsItemSitesSiteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemSitesSiteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemSitesSiteItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentTypes provides operations to manage the contentTypes property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) ContentTypes()(*GroupsItemSitesItemContentTypesRequestBuilder) {
    return NewGroupsItemSitesItemContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById provides operations to manage the contentTypes property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) ContentTypesById(id string)(*GroupsItemSitesItemContentTypesContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did"] = id
    }
    return NewGroupsItemSitesItemContentTypesContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateGetRequestInformation the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *GroupsItemSitesSiteItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemSitesSiteItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property sites in groups
func (m *GroupsItemSitesSiteItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, requestConfiguration *GroupsItemSitesSiteItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Drive provides operations to manage the drive property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) Drive()(*GroupsItemSitesItemDriveRequestBuilder) {
    return NewGroupsItemSitesItemDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives provides operations to manage the drives property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) Drives()(*GroupsItemSitesItemDrivesRequestBuilder) {
    return NewGroupsItemSitesItemDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById provides operations to manage the drives property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) DrivesById(id string)(*GroupsItemSitesItemDrivesDriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return NewGroupsItemSitesItemDrivesDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExternalColumns provides operations to manage the externalColumns property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) ExternalColumns()(*GroupsItemSitesItemExternalColumnsRequestBuilder) {
    return NewGroupsItemSitesItemExternalColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalColumnsById provides operations to manage the externalColumns property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) ExternalColumnsById(id string)(*GroupsItemSitesItemExternalColumnsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewGroupsItemSitesItemExternalColumnsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *GroupsItemSitesSiteItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemSitesSiteItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable), nil
}
// GetActivitiesByInterval provides operations to call the getActivitiesByInterval method.
func (m *GroupsItemSitesSiteItemRequestBuilder) GetActivitiesByInterval()(*GroupsItemSitesItemGetActivitiesByIntervalRequestBuilder) {
    return NewGroupsItemSitesItemGetActivitiesByIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *GroupsItemSitesSiteItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*GroupsItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewGroupsItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// GetApplicableContentTypesForListWithListId provides operations to call the getApplicableContentTypesForList method.
func (m *GroupsItemSitesSiteItemRequestBuilder) GetApplicableContentTypesForListWithListId(listId *string)(*GroupsItemSitesItemGetApplicableContentTypesForListWithListIdRequestBuilder) {
    return NewGroupsItemSitesItemGetApplicableContentTypesForListWithListIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, listId);
}
// GetByPathWithPath provides operations to call the getByPath method.
func (m *GroupsItemSitesSiteItemRequestBuilder) GetByPathWithPath(path *string)(*GroupsItemSitesItemGetByPathWithPathRequestBuilder) {
    return NewGroupsItemSitesItemGetByPathWithPathRequestBuilderInternal(m.pathParameters, m.requestAdapter, path);
}
// Items provides operations to manage the items property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) Items()(*GroupsItemSitesItemItemsRequestBuilder) {
    return NewGroupsItemSitesItemItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById provides operations to manage the items property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) ItemsById(id string)(*GroupsItemSitesItemItemsBaseItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["baseItem%2Did"] = id
    }
    return NewGroupsItemSitesItemItemsBaseItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Lists provides operations to manage the lists property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) Lists()(*GroupsItemSitesItemListsRequestBuilder) {
    return NewGroupsItemSitesItemListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListsById provides operations to manage the lists property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) ListsById(id string)(*GroupsItemSitesItemListsListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["list%2Did"] = id
    }
    return NewGroupsItemSitesItemListsListItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) Onenote()(*GroupsItemSitesItemOnenoteRequestBuilder) {
    return NewGroupsItemSitesItemOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) Operations()(*GroupsItemSitesItemOperationsRequestBuilder) {
    return NewGroupsItemSitesItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) OperationsById(id string)(*GroupsItemSitesItemOperationsRichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation%2Did"] = id
    }
    return NewGroupsItemSitesItemOperationsRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property sites in groups
func (m *GroupsItemSitesSiteItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, requestConfiguration *GroupsItemSitesSiteItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable), nil
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) Permissions()(*GroupsItemSitesItemPermissionsRequestBuilder) {
    return NewGroupsItemSitesItemPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) PermissionsById(id string)(*GroupsItemSitesItemPermissionsPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return NewGroupsItemSitesItemPermissionsPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites provides operations to manage the sites property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) Sites()(*GroupsItemSitesItemSitesRequestBuilder) {
    return NewGroupsItemSitesItemSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById provides operations to manage the sites property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) SitesById(id string)(*GroupsItemSitesItemSitesSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did1"] = id
    }
    return NewGroupsItemSitesItemSitesSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TermStore provides operations to manage the termStore property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) TermStore()(*GroupsItemSitesItemTermStoreRequestBuilder) {
    return NewGroupsItemSitesItemTermStoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermStores provides operations to manage the termStores property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) TermStores()(*GroupsItemSitesItemTermStoresRequestBuilder) {
    return NewGroupsItemSitesItemTermStoresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermStoresById provides operations to manage the termStores property of the microsoft.graph.site entity.
func (m *GroupsItemSitesSiteItemRequestBuilder) TermStoresById(id string)(*GroupsItemSitesItemTermStoresStoreItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["store%2Did"] = id
    }
    return NewGroupsItemSitesItemTermStoresStoreItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
