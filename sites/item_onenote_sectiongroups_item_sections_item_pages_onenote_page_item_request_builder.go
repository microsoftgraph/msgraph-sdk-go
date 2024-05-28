package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder provides operations to manage the pages property of the microsoft.graph.onenoteSection entity.
type ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters the collection of pages in the section.  Read-only. Nullable.
type ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters
}
// ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderInternal instantiates a new ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder and sets the default values.
func NewItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) {
    m := &ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/onenote/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder instantiates a new ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder and sets the default values.
func NewItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the site entity.
// returns a *ItemOnenoteSectiongroupsItemSectionsItemPagesItemContentRequestBuilder when successful
func (m *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Content()(*ItemOnenoteSectiongroupsItemSectionsItemPagesItemContentRequestBuilder) {
    return NewItemOnenoteSectiongroupsItemSectionsItemPagesItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CopyToSection provides operations to call the copyToSection method.
// returns a *ItemOnenoteSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilder when successful
func (m *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) CopyToSection()(*ItemOnenoteSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilder) {
    return NewItemOnenoteSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property pages for sites
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of pages in the section.  Read-only. Nullable.
// returns a OnenotePageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenotePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable), nil
}
// OnenotePatchContent provides operations to call the onenotePatchContent method.
// returns a *ItemOnenoteSectiongroupsItemSectionsItemPagesItemOnenotepatchcontentOnenotePatchContentRequestBuilder when successful
func (m *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) OnenotePatchContent()(*ItemOnenoteSectiongroupsItemSectionsItemPagesItemOnenotepatchcontentOnenotePatchContentRequestBuilder) {
    return NewItemOnenoteSectiongroupsItemSectionsItemPagesItemOnenotepatchcontentOnenotePatchContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ParentNotebook provides operations to manage the parentNotebook property of the microsoft.graph.onenotePage entity.
// returns a *ItemOnenoteSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilder when successful
func (m *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) ParentNotebook()(*ItemOnenoteSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilder) {
    return NewItemOnenoteSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ParentSection provides operations to manage the parentSection property of the microsoft.graph.onenotePage entity.
// returns a *ItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilder when successful
func (m *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) ParentSection()(*ItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilder) {
    return NewItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property pages in sites
// returns a OnenotePageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, requestConfiguration *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenotePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable), nil
}
// Preview provides operations to call the preview method.
// returns a *ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilder when successful
func (m *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Preview()(*ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilder) {
    return NewItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property pages for sites
// returns a *RequestInformation when successful
func (m *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the collection of pages in the section.  Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property pages in sites
// returns a *RequestInformation when successful
func (m *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, requestConfiguration *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder when successful
func (m *ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) WithUrl(rawUrl string)(*ItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) {
    return NewItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
