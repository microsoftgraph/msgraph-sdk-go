package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i4f8dd6fbbbd8e5eac4c018e252d494b346d6344a7e2f2cf4d26e0787cfc50d17 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sectiongroups/item/sections"
    ib4e3139e711c23865c0e217671361df83d26fdce68565cc23cedb903bb5b8eff "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sectiongroups/item/parentsectiongroup"
    ibb6e449347a883aed7ef7d46e6fde8269b6148c357bcb3eee0fea66f175d3b29 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sectiongroups/item/sectiongroups"
    iecb427cf10e33fed277f8759f6db9ebd96d9f26a571e9a50f5f0c15563ed1f94 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sectiongroups/item/parentnotebook"
    i572c7a753a0704b855af1cd5e3531b9eaa2442243b3635858b7a519c4d054813 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sectiongroups/item/sections/item"
    i8533efee27114835d72ec3a21661ae353e4b65d8107ece04573cc5a2bac2c299 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sectiongroups/item/sectiongroups/item"
)

// SectionGroupItemRequestBuilder provides operations to manage the sectionGroups property of the microsoft.graph.notebook entity.
type SectionGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SectionGroupItemRequestBuilderDeleteOptions options for Delete
type SectionGroupItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// SectionGroupItemRequestBuilderGetOptions options for Get
type SectionGroupItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *SectionGroupItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// SectionGroupItemRequestBuilderGetQueryParameters the section groups in the notebook. Read-only. Nullable.
type SectionGroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SectionGroupItemRequestBuilderPatchOptions options for Patch
type SectionGroupItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewSectionGroupItemRequestBuilderInternal instantiates a new SectionGroupItemRequestBuilder and sets the default values.
func NewSectionGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SectionGroupItemRequestBuilder) {
    m := &SectionGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/onenote/notebooks/{notebook_id}/sectionGroups/{sectionGroup_id}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property sectionGroups for users
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
// CreateGetRequestInformation the section groups in the notebook. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property sectionGroups in users
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
// Delete delete navigation property sectionGroups for users
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
// Get the section groups in the notebook. Read-only. Nullable.
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
func (m *SectionGroupItemRequestBuilder) ParentNotebook()(*iecb427cf10e33fed277f8759f6db9ebd96d9f26a571e9a50f5f0c15563ed1f94.ParentNotebookRequestBuilder) {
    return iecb427cf10e33fed277f8759f6db9ebd96d9f26a571e9a50f5f0c15563ed1f94.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentSectionGroup the parentSectionGroup property
func (m *SectionGroupItemRequestBuilder) ParentSectionGroup()(*ib4e3139e711c23865c0e217671361df83d26fdce68565cc23cedb903bb5b8eff.ParentSectionGroupRequestBuilder) {
    return ib4e3139e711c23865c0e217671361df83d26fdce68565cc23cedb903bb5b8eff.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property sectionGroups in users
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
func (m *SectionGroupItemRequestBuilder) SectionGroups()(*ibb6e449347a883aed7ef7d46e6fde8269b6148c357bcb3eee0fea66f175d3b29.SectionGroupsRequestBuilder) {
    return ibb6e449347a883aed7ef7d46e6fde8269b6148c357bcb3eee0fea66f175d3b29.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.onenote.notebooks.item.sectionGroups.item.sectionGroups.item collection
func (m *SectionGroupItemRequestBuilder) SectionGroupsById(id string)(*i8533efee27114835d72ec3a21661ae353e4b65d8107ece04573cc5a2bac2c299.SectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id1"] = id
    }
    return i8533efee27114835d72ec3a21661ae353e4b65d8107ece04573cc5a2bac2c299.NewSectionGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sections the sections property
func (m *SectionGroupItemRequestBuilder) Sections()(*i4f8dd6fbbbd8e5eac4c018e252d494b346d6344a7e2f2cf4d26e0787cfc50d17.SectionsRequestBuilder) {
    return i4f8dd6fbbbd8e5eac4c018e252d494b346d6344a7e2f2cf4d26e0787cfc50d17.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.onenote.notebooks.item.sectionGroups.item.sections.item collection
func (m *SectionGroupItemRequestBuilder) SectionsById(id string)(*i572c7a753a0704b855af1cd5e3531b9eaa2442243b3635858b7a519c4d054813.OnenoteSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i572c7a753a0704b855af1cd5e3531b9eaa2442243b3635858b7a519c4d054813.NewOnenoteSectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
