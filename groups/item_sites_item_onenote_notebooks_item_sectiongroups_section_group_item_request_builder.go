package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder provides operations to manage the sectionGroups property of the microsoft.graph.notebook entity.
type ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderGetQueryParameters the section groups in the notebook. Read-only. Nullable.
type ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderGetQueryParameters
}
// ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderInternal instantiates a new ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder) {
    m := &ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote/notebooks/{notebook%2Did}/sectionGroups/{sectionGroup%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder instantiates a new ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property sectionGroups for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the section groups in the notebook. Read-only. Nullable.
// returns a SectionGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable, error) {
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
// returns a *ItemSitesItemOnenoteNotebooksItemSectiongroupsItemParentnotebookParentNotebookRequestBuilder when successful
func (m *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder) ParentNotebook()(*ItemSitesItemOnenoteNotebooksItemSectiongroupsItemParentnotebookParentNotebookRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectiongroupsItemParentnotebookParentNotebookRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ParentSectionGroup provides operations to manage the parentSectionGroup property of the microsoft.graph.sectionGroup entity.
// returns a *ItemSitesItemOnenoteNotebooksItemSectiongroupsItemParentsectiongroupParentSectionGroupRequestBuilder when successful
func (m *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder) ParentSectionGroup()(*ItemSitesItemOnenoteNotebooksItemSectiongroupsItemParentsectiongroupParentSectionGroupRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectiongroupsItemParentsectiongroupParentSectionGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property sectionGroups in groups
// returns a SectionGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable, error) {
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
// returns a *ItemSitesItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder when successful
func (m *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder) SectionGroups()(*ItemSitesItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sections provides operations to manage the sections property of the microsoft.graph.sectionGroup entity.
// returns a *ItemSitesItemOnenoteNotebooksItemSectiongroupsItemSectionsRequestBuilder when successful
func (m *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder) Sections()(*ItemSitesItemOnenoteNotebooksItemSectiongroupsItemSectionsRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectiongroupsItemSectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property sectionGroups for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the section groups in the notebook. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder when successful
func (m *ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectiongroupsSectionGroupItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
