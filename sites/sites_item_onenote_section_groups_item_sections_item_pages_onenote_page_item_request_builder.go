package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder provides operations to manage the pages property of the microsoft.graph.onenoteSection entity.
type SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters the collection of pages in the section.  Read-only. Nullable.
type SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters
}
// SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderInternal instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewSitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) {
    m := &SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/onenote/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewSitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the site entity.
func (m *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Content()(*SitesItemOnenoteSectionGroupsItemSectionsItemPagesItemContentRequestBuilder) {
    return NewSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CopyToSection provides operations to call the copyToSection method.
func (m *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) CopyToSection()(*SitesItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionRequestBuilder) {
    return NewSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property pages for sites
func (m *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the collection of pages in the section.  Read-only. Nullable.
func (m *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property pages in sites
func (m *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, requestConfiguration *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property pages for sites
func (m *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of pages in the section.  Read-only. Nullable.
func (m *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenotePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable), nil
}
// OnenotePatchContent provides operations to call the onenotePatchContent method.
func (m *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) OnenotePatchContent()(*SitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentRequestBuilder) {
    return NewSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentNotebook provides operations to manage the parentNotebook property of the microsoft.graph.onenotePage entity.
func (m *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) ParentNotebook()(*SitesItemOnenoteSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilder) {
    return NewSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentSection provides operations to manage the parentSection property of the microsoft.graph.onenotePage entity.
func (m *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) ParentSection()(*SitesItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder) {
    return NewSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property pages in sites
func (m *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, requestConfiguration *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenotePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable), nil
}
// Preview provides operations to call the preview method.
func (m *SitesItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Preview()(*SitesItemOnenoteSectionGroupsItemSectionsItemPagesItemPreviewRequestBuilder) {
    return NewSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
