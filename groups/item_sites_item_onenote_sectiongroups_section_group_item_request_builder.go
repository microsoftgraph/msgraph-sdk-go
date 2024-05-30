package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder provides operations to manage the sectionGroups property of the microsoft.graph.onenote entity.
type ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderGetQueryParameters the section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
type ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderGetQueryParameters
}
// ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderInternal instantiates a new ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder) {
    m := &ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote/sectionGroups/{sectionGroup%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder instantiates a new ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property sectionGroups for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
// returns a SectionGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSectionGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable), nil
}
// ParentNotebook provides operations to manage the parentNotebook property of the microsoft.graph.sectionGroup entity.
// returns a *ItemSitesItemOnenoteSectiongroupsItemParentnotebookParentNotebookRequestBuilder when successful
func (m *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder) ParentNotebook()(*ItemSitesItemOnenoteSectiongroupsItemParentnotebookParentNotebookRequestBuilder) {
    return NewItemSitesItemOnenoteSectiongroupsItemParentnotebookParentNotebookRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ParentSectionGroup provides operations to manage the parentSectionGroup property of the microsoft.graph.sectionGroup entity.
// returns a *ItemSitesItemOnenoteSectiongroupsItemParentsectiongroupParentSectionGroupRequestBuilder when successful
func (m *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder) ParentSectionGroup()(*ItemSitesItemOnenoteSectiongroupsItemParentsectiongroupParentSectionGroupRequestBuilder) {
    return NewItemSitesItemOnenoteSectiongroupsItemParentsectiongroupParentSectionGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property sectionGroups in groups
// returns a SectionGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable, requestConfiguration *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSectionGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable), nil
}
// SectionGroups provides operations to manage the sectionGroups property of the microsoft.graph.sectionGroup entity.
// returns a *ItemSitesItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder when successful
func (m *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder) SectionGroups()(*ItemSitesItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) {
    return NewItemSitesItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sections provides operations to manage the sections property of the microsoft.graph.sectionGroup entity.
// returns a *ItemSitesItemOnenoteSectiongroupsItemSectionsRequestBuilder when successful
func (m *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder) Sections()(*ItemSitesItemOnenoteSectiongroupsItemSectionsRequestBuilder) {
    return NewItemSitesItemOnenoteSectiongroupsItemSectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property sectionGroups for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property sectionGroups in groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable, requestConfiguration *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder when successful
func (m *ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder) {
    return NewItemSitesItemOnenoteSectiongroupsSectionGroupItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
