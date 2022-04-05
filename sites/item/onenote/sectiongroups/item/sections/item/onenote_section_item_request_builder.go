package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0f2676d82370f6a55380863b65b8bb740d78b68c0bf2eed5487f88a691fb3a0f "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/sectiongroups/item/sections/item/parentsectiongroup"
    i2d94e1a60f6e1ddcfb93dbf70214d26a7ba6f27916fe9d0c1f4d9f58267891c0 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/sectiongroups/item/sections/item/parentnotebook"
    i41bc329878db5931bc722e89716417fd5e6574257e738b090a16bfc7a6ae576a "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/sectiongroups/item/sections/item/pages"
    i66552648433b2d958aa923a792a721316006d0154cc2fe41b17742d3f2f2801b "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/sectiongroups/item/sections/item/copytosectiongroup"
    ia9bc885d9447c2a26504b4a81f6a942f066471265e9caabb83df985e05eff939 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/sectiongroups/item/sections/item/copytonotebook"
    i49326ed096b83738236176e336c6019f491d91d1f7abfca60eaa5a85032c14ce "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/sectiongroups/item/sections/item/pages/item"
)

// OnenoteSectionItemRequestBuilder provides operations to manage the sections property of the microsoft.graph.sectionGroup entity.
type OnenoteSectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnenoteSectionItemRequestBuilderDeleteOptions options for Delete
type OnenoteSectionItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// OnenoteSectionItemRequestBuilderGetOptions options for Get
type OnenoteSectionItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *OnenoteSectionItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// OnenoteSectionItemRequestBuilderGetQueryParameters the sections in the section group. Read-only. Nullable.
type OnenoteSectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenoteSectionItemRequestBuilderPatchOptions options for Patch
type OnenoteSectionItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenoteSectionable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewOnenoteSectionItemRequestBuilderInternal instantiates a new OnenoteSectionItemRequestBuilder and sets the default values.
func NewOnenoteSectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenoteSectionItemRequestBuilder) {
    m := &OnenoteSectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/onenote/sectionGroups/{sectionGroup_id}/sections/{onenoteSection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteSectionItemRequestBuilder instantiates a new OnenoteSectionItemRequestBuilder and sets the default values.
func NewOnenoteSectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenoteSectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteSectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CopyToNotebook the copyToNotebook property
func (m *OnenoteSectionItemRequestBuilder) CopyToNotebook()(*ia9bc885d9447c2a26504b4a81f6a942f066471265e9caabb83df985e05eff939.CopyToNotebookRequestBuilder) {
    return ia9bc885d9447c2a26504b4a81f6a942f066471265e9caabb83df985e05eff939.NewCopyToNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CopyToSectionGroup the copyToSectionGroup property
func (m *OnenoteSectionItemRequestBuilder) CopyToSectionGroup()(*i66552648433b2d958aa923a792a721316006d0154cc2fe41b17742d3f2f2801b.CopyToSectionGroupRequestBuilder) {
    return i66552648433b2d958aa923a792a721316006d0154cc2fe41b17742d3f2f2801b.NewCopyToSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property sections for sites
func (m *OnenoteSectionItemRequestBuilder) CreateDeleteRequestInformation(options *OnenoteSectionItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) CreateGetRequestInformation(options *OnenoteSectionItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property sections in sites
func (m *OnenoteSectionItemRequestBuilder) CreatePatchRequestInformation(options *OnenoteSectionItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property sections for sites
func (m *OnenoteSectionItemRequestBuilder) Delete(options *OnenoteSectionItemRequestBuilderDeleteOptions)(error) {
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
// Get the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) Get(options *OnenoteSectionItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenoteSectionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenoteSectionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenoteSectionable), nil
}
// Pages the pages property
func (m *OnenoteSectionItemRequestBuilder) Pages()(*i41bc329878db5931bc722e89716417fd5e6574257e738b090a16bfc7a6ae576a.PagesRequestBuilder) {
    return i41bc329878db5931bc722e89716417fd5e6574257e738b090a16bfc7a6ae576a.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.sectionGroups.item.sections.item.pages.item collection
func (m *OnenoteSectionItemRequestBuilder) PagesById(id string)(*i49326ed096b83738236176e336c6019f491d91d1f7abfca60eaa5a85032c14ce.OnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return i49326ed096b83738236176e336c6019f491d91d1f7abfca60eaa5a85032c14ce.NewOnenotePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ParentNotebook the parentNotebook property
func (m *OnenoteSectionItemRequestBuilder) ParentNotebook()(*i2d94e1a60f6e1ddcfb93dbf70214d26a7ba6f27916fe9d0c1f4d9f58267891c0.ParentNotebookRequestBuilder) {
    return i2d94e1a60f6e1ddcfb93dbf70214d26a7ba6f27916fe9d0c1f4d9f58267891c0.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentSectionGroup the parentSectionGroup property
func (m *OnenoteSectionItemRequestBuilder) ParentSectionGroup()(*i0f2676d82370f6a55380863b65b8bb740d78b68c0bf2eed5487f88a691fb3a0f.ParentSectionGroupRequestBuilder) {
    return i0f2676d82370f6a55380863b65b8bb740d78b68c0bf2eed5487f88a691fb3a0f.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property sections in sites
func (m *OnenoteSectionItemRequestBuilder) Patch(options *OnenoteSectionItemRequestBuilderPatchOptions)(error) {
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
