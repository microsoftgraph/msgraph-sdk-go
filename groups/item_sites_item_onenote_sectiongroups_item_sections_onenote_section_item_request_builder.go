package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder provides operations to manage the sections property of the microsoft.graph.sectionGroup entity.
type ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderGetQueryParameters the sections in the section group. Read-only. Nullable.
type ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderGetQueryParameters
}
// ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderInternal instantiates a new ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder) {
    m := &ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder instantiates a new ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CopyToNotebook provides operations to call the copyToNotebook method.
// returns a *ItemSitesItemOnenoteSectiongroupsItemSectionsItemCopytonotebookCopyToNotebookRequestBuilder when successful
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder) CopyToNotebook()(*ItemSitesItemOnenoteSectiongroupsItemSectionsItemCopytonotebookCopyToNotebookRequestBuilder) {
    return NewItemSitesItemOnenoteSectiongroupsItemSectionsItemCopytonotebookCopyToNotebookRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CopyToSectionGroup provides operations to call the copyToSectionGroup method.
// returns a *ItemSitesItemOnenoteSectiongroupsItemSectionsItemCopytosectiongroupCopyToSectionGroupRequestBuilder when successful
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder) CopyToSectionGroup()(*ItemSitesItemOnenoteSectiongroupsItemSectionsItemCopytosectiongroupCopyToSectionGroupRequestBuilder) {
    return NewItemSitesItemOnenoteSectiongroupsItemSectionsItemCopytosectiongroupCopyToSectionGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property sections for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the sections in the section group. Read-only. Nullable.
// returns a OnenoteSectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenoteSectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenoteSectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenoteSectionable), nil
}
// Pages provides operations to manage the pages property of the microsoft.graph.onenoteSection entity.
// returns a *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder when successful
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder) Pages()(*ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder) {
    return NewItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ParentNotebook provides operations to manage the parentNotebook property of the microsoft.graph.onenoteSection entity.
// returns a *ItemSitesItemOnenoteSectiongroupsItemSectionsItemParentnotebookParentNotebookRequestBuilder when successful
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder) ParentNotebook()(*ItemSitesItemOnenoteSectiongroupsItemSectionsItemParentnotebookParentNotebookRequestBuilder) {
    return NewItemSitesItemOnenoteSectiongroupsItemSectionsItemParentnotebookParentNotebookRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ParentSectionGroup provides operations to manage the parentSectionGroup property of the microsoft.graph.onenoteSection entity.
// returns a *ItemSitesItemOnenoteSectiongroupsItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilder when successful
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder) ParentSectionGroup()(*ItemSitesItemOnenoteSectiongroupsItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilder) {
    return NewItemSitesItemOnenoteSectiongroupsItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property sections in groups
// returns a OnenoteSectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenoteSectionable, requestConfiguration *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenoteSectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenoteSectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenoteSectionable), nil
}
// ToDeleteRequestInformation delete navigation property sections for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the sections in the section group. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property sections in groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenoteSectionable, requestConfiguration *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder when successful
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder) {
    return NewItemSitesItemOnenoteSectiongroupsItemSectionsOnenoteSectionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
