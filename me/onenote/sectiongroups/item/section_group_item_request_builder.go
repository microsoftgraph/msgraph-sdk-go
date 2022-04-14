package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i1e7a6fa63a3aa1a30a3f603f9d8e5e8209eb27cefdf999460ff279bb7f728dfc "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item/parentnotebook"
    ib98f9d64f1e81452e6781379404033a621251161412ccb2bbf577bb535be194d "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item/parentsectiongroup"
    ibe31feeb46a70d408fa28b6d36b2e15d5e5bcb7c5551fab265996d88d9782169 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item/sections"
    iee9f95785d7c785b1f1294693ca429c811af6695d07b63514b421240e65421d4 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item/sectiongroups"
    i10dbc5d719db7ec75326404ecda829e986c01e79190f6c0030372eb05feaebee "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item/sections/item"
    ie3dc997ceb1c48bcfbc46b5096a7af15491b262834d0661de1716546dedad04b "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item/sectiongroups/item"
)

// SectionGroupItemRequestBuilder provides operations to manage the sectionGroups property of the microsoft.graph.onenote entity.
type SectionGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SectionGroupItemRequestBuilderDeleteOptions options for Delete
type SectionGroupItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SectionGroupItemRequestBuilderGetOptions options for Get
type SectionGroupItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SectionGroupItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SectionGroupItemRequestBuilderGetQueryParameters the section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
type SectionGroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SectionGroupItemRequestBuilderPatchOptions options for Patch
type SectionGroupItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewSectionGroupItemRequestBuilderInternal instantiates a new SectionGroupItemRequestBuilder and sets the default values.
func NewSectionGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SectionGroupItemRequestBuilder) {
    m := &SectionGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onenote/sectionGroups/{sectionGroup%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSectionGroupItemRequestBuilder instantiates a new SectionGroupItemRequestBuilder and sets the default values.
func NewSectionGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SectionGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSectionGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property sectionGroups for me
func (m *SectionGroupItemRequestBuilder) CreateDeleteRequestInformation(options *SectionGroupItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *SectionGroupItemRequestBuilder) CreateGetRequestInformation(options *SectionGroupItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property sectionGroups in me
func (m *SectionGroupItemRequestBuilder) CreatePatchRequestInformation(options *SectionGroupItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property sectionGroups for me
func (m *SectionGroupItemRequestBuilder) Delete(options *SectionGroupItemRequestBuilderDeleteOptions)(error) {
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
// Get the section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *SectionGroupItemRequestBuilder) Get(options *SectionGroupItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSectionGroupFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable), nil
}
// ParentNotebook the parentNotebook property
func (m *SectionGroupItemRequestBuilder) ParentNotebook()(*i1e7a6fa63a3aa1a30a3f603f9d8e5e8209eb27cefdf999460ff279bb7f728dfc.ParentNotebookRequestBuilder) {
    return i1e7a6fa63a3aa1a30a3f603f9d8e5e8209eb27cefdf999460ff279bb7f728dfc.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentSectionGroup the parentSectionGroup property
func (m *SectionGroupItemRequestBuilder) ParentSectionGroup()(*ib98f9d64f1e81452e6781379404033a621251161412ccb2bbf577bb535be194d.ParentSectionGroupRequestBuilder) {
    return ib98f9d64f1e81452e6781379404033a621251161412ccb2bbf577bb535be194d.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property sectionGroups in me
func (m *SectionGroupItemRequestBuilder) Patch(options *SectionGroupItemRequestBuilderPatchOptions)(error) {
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
// SectionGroups the sectionGroups property
func (m *SectionGroupItemRequestBuilder) SectionGroups()(*iee9f95785d7c785b1f1294693ca429c811af6695d07b63514b421240e65421d4.SectionGroupsRequestBuilder) {
    return iee9f95785d7c785b1f1294693ca429c811af6695d07b63514b421240e65421d4.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.onenote.sectionGroups.item.sectionGroups.item collection
func (m *SectionGroupItemRequestBuilder) SectionGroupsById(id string)(*ie3dc997ceb1c48bcfbc46b5096a7af15491b262834d0661de1716546dedad04b.SectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup%2Did1"] = id
    }
    return ie3dc997ceb1c48bcfbc46b5096a7af15491b262834d0661de1716546dedad04b.NewSectionGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sections the sections property
func (m *SectionGroupItemRequestBuilder) Sections()(*ibe31feeb46a70d408fa28b6d36b2e15d5e5bcb7c5551fab265996d88d9782169.SectionsRequestBuilder) {
    return ibe31feeb46a70d408fa28b6d36b2e15d5e5bcb7c5551fab265996d88d9782169.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.onenote.sectionGroups.item.sections.item collection
func (m *SectionGroupItemRequestBuilder) SectionsById(id string)(*i10dbc5d719db7ec75326404ecda829e986c01e79190f6c0030372eb05feaebee.OnenoteSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection%2Did"] = id
    }
    return i10dbc5d719db7ec75326404ecda829e986c01e79190f6c0030372eb05feaebee.NewOnenoteSectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
