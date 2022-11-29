package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SitesSiteItemRequestBuilder provides operations to manage the collection of site entities.
type SitesSiteItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SitesSiteItemRequestBuilderGetQueryParameters retrieve properties and relationships for a [site][] resource.A **site** resource represents a team site in SharePoint.
type SitesSiteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SitesSiteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SitesSiteItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SitesSiteItemRequestBuilderGetQueryParameters
}
// SitesSiteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SitesSiteItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) Analytics()(*SitesItemAnalyticsRequestBuilder) {
    return NewSitesItemAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Columns provides operations to manage the columns property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) Columns()(*SitesItemColumnsRequestBuilder) {
    return NewSitesItemColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById provides operations to manage the columns property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) ColumnsById(id string)(*SitesItemColumnsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewSitesItemColumnsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSitesSiteItemRequestBuilderInternal instantiates a new SiteItemRequestBuilder and sets the default values.
func NewSitesSiteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SitesSiteItemRequestBuilder) {
    m := &SitesSiteItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSitesSiteItemRequestBuilder instantiates a new SiteItemRequestBuilder and sets the default values.
func NewSitesSiteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SitesSiteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSitesSiteItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentTypes provides operations to manage the contentTypes property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) ContentTypes()(*SitesItemContentTypesRequestBuilder) {
    return NewSitesItemContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById provides operations to manage the contentTypes property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) ContentTypesById(id string)(*SitesItemContentTypesContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did"] = id
    }
    return NewSitesItemContentTypesContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateGetRequestInformation retrieve properties and relationships for a [site][] resource.A **site** resource represents a team site in SharePoint.
func (m *SitesSiteItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SitesSiteItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in sites
func (m *SitesSiteItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, requestConfiguration *SitesSiteItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *SitesSiteItemRequestBuilder) Drive()(*SitesItemDriveRequestBuilder) {
    return NewSitesItemDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives provides operations to manage the drives property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) Drives()(*SitesItemDrivesRequestBuilder) {
    return NewSitesItemDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById provides operations to manage the drives property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) DrivesById(id string)(*SitesItemDrivesDriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return NewSitesItemDrivesDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExternalColumns provides operations to manage the externalColumns property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) ExternalColumns()(*SitesItemExternalColumnsRequestBuilder) {
    return NewSitesItemExternalColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalColumnsById provides operations to manage the externalColumns property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) ExternalColumnsById(id string)(*SitesItemExternalColumnsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewSitesItemExternalColumnsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get retrieve properties and relationships for a [site][] resource.A **site** resource represents a team site in SharePoint.
func (m *SitesSiteItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SitesSiteItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, error) {
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
func (m *SitesSiteItemRequestBuilder) GetActivitiesByInterval()(*SitesItemGetActivitiesByIntervalRequestBuilder) {
    return NewSitesItemGetActivitiesByIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *SitesSiteItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*SitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// GetApplicableContentTypesForListWithListId provides operations to call the getApplicableContentTypesForList method.
func (m *SitesSiteItemRequestBuilder) GetApplicableContentTypesForListWithListId(listId *string)(*SitesItemGetApplicableContentTypesForListWithListIdRequestBuilder) {
    return NewSitesItemGetApplicableContentTypesForListWithListIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, listId);
}
// GetByPathWithPath provides operations to call the getByPath method.
func (m *SitesSiteItemRequestBuilder) GetByPathWithPath(path *string)(*SitesItemGetByPathWithPathRequestBuilder) {
    return NewSitesItemGetByPathWithPathRequestBuilderInternal(m.pathParameters, m.requestAdapter, path);
}
// Items provides operations to manage the items property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) Items()(*SitesItemItemsRequestBuilder) {
    return NewSitesItemItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById provides operations to manage the items property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) ItemsById(id string)(*SitesItemItemsBaseItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["baseItem%2Did"] = id
    }
    return NewSitesItemItemsBaseItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Lists provides operations to manage the lists property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) Lists()(*SitesItemListsRequestBuilder) {
    return NewSitesItemListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListsById provides operations to manage the lists property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) ListsById(id string)(*SitesItemListsListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["list%2Did"] = id
    }
    return NewSitesItemListsListItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) Onenote()(*SitesItemOnenoteRequestBuilder) {
    return NewSitesItemOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) Operations()(*SitesItemOperationsRequestBuilder) {
    return NewSitesItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) OperationsById(id string)(*SitesItemOperationsRichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation%2Did"] = id
    }
    return NewSitesItemOperationsRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in sites
func (m *SitesSiteItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, requestConfiguration *SitesSiteItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, error) {
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
func (m *SitesSiteItemRequestBuilder) Permissions()(*SitesItemPermissionsRequestBuilder) {
    return NewSitesItemPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) PermissionsById(id string)(*SitesItemPermissionsPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return NewSitesItemPermissionsPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites provides operations to manage the sites property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) Sites()(*SitesItemSitesRequestBuilder) {
    return NewSitesItemSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById provides operations to manage the sites property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) SitesById(id string)(*SitesItemSitesSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did1"] = id
    }
    return NewSitesItemSitesSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TermStore provides operations to manage the termStore property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) TermStore()(*SitesItemTermStoreRequestBuilder) {
    return NewSitesItemTermStoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermStores provides operations to manage the termStores property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) TermStores()(*SitesItemTermStoresRequestBuilder) {
    return NewSitesItemTermStoresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermStoresById provides operations to manage the termStores property of the microsoft.graph.site entity.
func (m *SitesSiteItemRequestBuilder) TermStoresById(id string)(*SitesItemTermStoresStoreItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["store%2Did"] = id
    }
    return NewSitesItemTermStoresStoreItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
