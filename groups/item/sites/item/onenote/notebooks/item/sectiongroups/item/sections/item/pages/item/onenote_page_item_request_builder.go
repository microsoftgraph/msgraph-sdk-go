package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0751f0ef955f097a6e1d28cf633304900fd1d31f3661a913768b93eb5d4b29ff "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/notebooks/item/sectiongroups/item/sections/item/pages/item/copytosection"
    i14378ec5676c963676924aeb88bc35e83d10f976067ab7b3e5a5687056fa953e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/notebooks/item/sectiongroups/item/sections/item/pages/item/content"
    i15e366600667a0ef904369fd5d5dff55b31b941646ead717bc0ced010056f279 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/notebooks/item/sectiongroups/item/sections/item/pages/item/parentnotebook"
    i65c576d3d246b37516aed72753fd6bb7f255f109f2cea0e98569eb2ae722284f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/notebooks/item/sectiongroups/item/sections/item/pages/item/parentsection"
    ida3b90ef3f99b5be7aa97f247e60ff4bb2d8f8818fe10021409353d52f16c5b9 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/notebooks/item/sectiongroups/item/sections/item/pages/item/preview"
    ie7ae16bbe76be7a72d711e474d435f6f706de6eed9643fa4acfde1976c4a07b8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/notebooks/item/sectiongroups/item/sections/item/pages/item/onenotepatchcontent"
)

// OnenotePageItemRequestBuilder provides operations to manage the pages property of the microsoft.graph.onenoteSection entity.
type OnenotePageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnenotePageItemRequestBuilderDeleteOptions options for Delete
type OnenotePageItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// OnenotePageItemRequestBuilderGetOptions options for Get
type OnenotePageItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *OnenotePageItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// OnenotePageItemRequestBuilderGetQueryParameters the collection of pages in the section.  Read-only. Nullable.
type OnenotePageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenotePageItemRequestBuilderPatchOptions options for Patch
type OnenotePageItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewOnenotePageItemRequestBuilderInternal instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewOnenotePageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenotePageItemRequestBuilder) {
    m := &OnenotePageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/sites/{site_id}/onenote/notebooks/{notebook_id}/sectionGroups/{sectionGroup_id}/sections/{onenoteSection_id}/pages/{onenotePage_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenotePageItemRequestBuilder instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewOnenotePageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenotePageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenotePageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the content property
func (m *OnenotePageItemRequestBuilder) Content()(*i14378ec5676c963676924aeb88bc35e83d10f976067ab7b3e5a5687056fa953e.ContentRequestBuilder) {
    return i14378ec5676c963676924aeb88bc35e83d10f976067ab7b3e5a5687056fa953e.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CopyToSection the copyToSection property
func (m *OnenotePageItemRequestBuilder) CopyToSection()(*i0751f0ef955f097a6e1d28cf633304900fd1d31f3661a913768b93eb5d4b29ff.CopyToSectionRequestBuilder) {
    return i0751f0ef955f097a6e1d28cf633304900fd1d31f3661a913768b93eb5d4b29ff.NewCopyToSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property pages for groups
func (m *OnenotePageItemRequestBuilder) CreateDeleteRequestInformation(options *OnenotePageItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the collection of pages in the section.  Read-only. Nullable.
func (m *OnenotePageItemRequestBuilder) CreateGetRequestInformation(options *OnenotePageItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property pages in groups
func (m *OnenotePageItemRequestBuilder) CreatePatchRequestInformation(options *OnenotePageItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property pages for groups
func (m *OnenotePageItemRequestBuilder) Delete(options *OnenotePageItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of pages in the section.  Read-only. Nullable.
func (m *OnenotePageItemRequestBuilder) Get(options *OnenotePageItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenotePageFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable), nil
}
// OnenotePatchContent the onenotePatchContent property
func (m *OnenotePageItemRequestBuilder) OnenotePatchContent()(*ie7ae16bbe76be7a72d711e474d435f6f706de6eed9643fa4acfde1976c4a07b8.OnenotePatchContentRequestBuilder) {
    return ie7ae16bbe76be7a72d711e474d435f6f706de6eed9643fa4acfde1976c4a07b8.NewOnenotePatchContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentNotebook the parentNotebook property
func (m *OnenotePageItemRequestBuilder) ParentNotebook()(*i15e366600667a0ef904369fd5d5dff55b31b941646ead717bc0ced010056f279.ParentNotebookRequestBuilder) {
    return i15e366600667a0ef904369fd5d5dff55b31b941646ead717bc0ced010056f279.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentSection the parentSection property
func (m *OnenotePageItemRequestBuilder) ParentSection()(*i65c576d3d246b37516aed72753fd6bb7f255f109f2cea0e98569eb2ae722284f.ParentSectionRequestBuilder) {
    return i65c576d3d246b37516aed72753fd6bb7f255f109f2cea0e98569eb2ae722284f.NewParentSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property pages in groups
func (m *OnenotePageItemRequestBuilder) Patch(options *OnenotePageItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Preview provides operations to call the preview method.
func (m *OnenotePageItemRequestBuilder) Preview()(*ida3b90ef3f99b5be7aa97f247e60ff4bb2d8f8818fe10021409353d52f16c5b9.PreviewRequestBuilder) {
    return ida3b90ef3f99b5be7aa97f247e60ff4bb2d8f8818fe10021409353d52f16c5b9.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
