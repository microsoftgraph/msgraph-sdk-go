package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOnenotePagesOnenotePageItemRequestBuilder provides operations to manage the pages property of the microsoft.graph.onenote entity.
type ItemOnenotePagesOnenotePageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnenotePagesOnenotePageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenotePagesOnenotePageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOnenotePagesOnenotePageItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a page object. **Getting page information** Access a page's metadata by page identifier: **Getting page content** You can use the page's `content` endpoint to get the HTML content of a page: The `includeIDs=true` query option is used to update pages.
type ItemOnenotePagesOnenotePageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOnenotePagesOnenotePageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenotePagesOnenotePageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnenotePagesOnenotePageItemRequestBuilderGetQueryParameters
}
// ItemOnenotePagesOnenotePageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenotePagesOnenotePageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnenotePagesOnenotePageItemRequestBuilderInternal instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewItemOnenotePagesOnenotePageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenotePagesOnenotePageItemRequestBuilder) {
    m := &ItemOnenotePagesOnenotePageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/onenote/pages/{onenotePage%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemOnenotePagesOnenotePageItemRequestBuilder instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewItemOnenotePagesOnenotePageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenotePagesOnenotePageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenotePagesOnenotePageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the site entity.
func (m *ItemOnenotePagesOnenotePageItemRequestBuilder) Content()(*ItemOnenotePagesItemContentRequestBuilder) {
    return NewItemOnenotePagesItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CopyToSection provides operations to call the copyToSection method.
func (m *ItemOnenotePagesOnenotePageItemRequestBuilder) CopyToSection()(*ItemOnenotePagesItemCopyToSectionRequestBuilder) {
    return NewItemOnenotePagesItemCopyToSectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete a OneNote page.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/page-delete?view=graph-rest-1.0
func (m *ItemOnenotePagesOnenotePageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemOnenotePagesOnenotePageItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get retrieve the properties and relationships of a page object. **Getting page information** Access a page's metadata by page identifier: **Getting page content** You can use the page's `content` endpoint to get the HTML content of a page: The `includeIDs=true` query option is used to update pages.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/page-get?view=graph-rest-1.0
func (m *ItemOnenotePagesOnenotePageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnenotePagesOnenotePageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *ItemOnenotePagesOnenotePageItemRequestBuilder) OnenotePatchContent()(*ItemOnenotePagesItemOnenotePatchContentRequestBuilder) {
    return NewItemOnenotePagesItemOnenotePatchContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ParentNotebook provides operations to manage the parentNotebook property of the microsoft.graph.onenotePage entity.
func (m *ItemOnenotePagesOnenotePageItemRequestBuilder) ParentNotebook()(*ItemOnenotePagesItemParentNotebookRequestBuilder) {
    return NewItemOnenotePagesItemParentNotebookRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ParentSection provides operations to manage the parentSection property of the microsoft.graph.onenotePage entity.
func (m *ItemOnenotePagesOnenotePageItemRequestBuilder) ParentSection()(*ItemOnenotePagesItemParentSectionRequestBuilder) {
    return NewItemOnenotePagesItemParentSectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property pages in sites
func (m *ItemOnenotePagesOnenotePageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, requestConfiguration *ItemOnenotePagesOnenotePageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *ItemOnenotePagesOnenotePageItemRequestBuilder) Preview()(*ItemOnenotePagesItemPreviewRequestBuilder) {
    return NewItemOnenotePagesItemPreviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a OneNote page.
func (m *ItemOnenotePagesOnenotePageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemOnenotePagesOnenotePageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of a page object. **Getting page information** Access a page's metadata by page identifier: **Getting page content** You can use the page's `content` endpoint to get the HTML content of a page: The `includeIDs=true` query option is used to update pages.
func (m *ItemOnenotePagesOnenotePageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnenotePagesOnenotePageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property pages in sites
func (m *ItemOnenotePagesOnenotePageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, requestConfiguration *ItemOnenotePagesOnenotePageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
