package onenote

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i54a3bfcdd648d6ee47e87c311695b29c4737b8921bb8f2a506c0864844df5dba "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sections"
    i97f70d655aa735279e41729aa6800f01397d2d7e8d41164003eff9b47387d3a5 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/notebooks"
    iac72d35f9c20cb360839134d86d535f87e1cdbe10b719f8c792d8017c06edd1d "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/operations"
    iae2d6340b64994bda52b72186cf936e932db0b6d72196a85d99b8d63141db19b "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/pages"
    ic51747a34f2e47436a56cc441d712037e581ac2bd968645ce4fe67a401a2f4bf "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/resources"
    id71911f826ddd0ad6edad424ad81ad5b822e25039b7e26730ec6b2f6de37ad3b "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups"
    i1927275756e9aeb82a189ce572418c67c8a998c4353590bc94f2a62a876dae2d "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/notebooks/item"
    i1941888ff368eb1fd130194825ce82fb010616c26e7af09e43370fa6bb2cfd71 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item"
    i2b8ad7bd2ae87c92c8852bbeef996852d7685df93cb12b50d58ecc094ecfd124 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/operations/item"
    i9ffb03d414d6abd9f024bf2d1f19d7b7fa97959d1245465fdaef4d163ddd6de4 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/pages/item"
    ibe60e1baafd5c0c7f830c971df51a300df66822516c0cfcd8a43fd2adc385c80 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sections/item"
    id67049e3fa5525a6ba50b6b8bc9f60107ac06f61256f8cf55bb655b4552ecd21 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/resources/item"
)

// OnenoteRequestBuilder provides operations to manage the onenote property of the microsoft.graph.user entity.
type OnenoteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnenoteRequestBuilderDeleteOptions options for Delete
type OnenoteRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// OnenoteRequestBuilderGetOptions options for Get
type OnenoteRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *OnenoteRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// OnenoteRequestBuilderGetQueryParameters read-only.
type OnenoteRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenoteRequestBuilderPatchOptions options for Patch
type OnenoteRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Onenoteable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewOnenoteRequestBuilderInternal instantiates a new OnenoteRequestBuilder and sets the default values.
func NewOnenoteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenoteRequestBuilder) {
    m := &OnenoteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onenote{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteRequestBuilder instantiates a new OnenoteRequestBuilder and sets the default values.
func NewOnenoteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenoteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property onenote for me
func (m *OnenoteRequestBuilder) CreateDeleteRequestInformation(options *OnenoteRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation read-only.
func (m *OnenoteRequestBuilder) CreateGetRequestInformation(options *OnenoteRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property onenote in me
func (m *OnenoteRequestBuilder) CreatePatchRequestInformation(options *OnenoteRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property onenote for me
func (m *OnenoteRequestBuilder) Delete(options *OnenoteRequestBuilderDeleteOptions)(error) {
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
// Get read-only.
func (m *OnenoteRequestBuilder) Get(options *OnenoteRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Onenoteable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenoteFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Onenoteable), nil
}
// Notebooks the notebooks property
func (m *OnenoteRequestBuilder) Notebooks()(*i97f70d655aa735279e41729aa6800f01397d2d7e8d41164003eff9b47387d3a5.NotebooksRequestBuilder) {
    return i97f70d655aa735279e41729aa6800f01397d2d7e8d41164003eff9b47387d3a5.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotebooksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.onenote.notebooks.item collection
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*i1927275756e9aeb82a189ce572418c67c8a998c4353590bc94f2a62a876dae2d.NotebookItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook_id"] = id
    }
    return i1927275756e9aeb82a189ce572418c67c8a998c4353590bc94f2a62a876dae2d.NewNotebookItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *OnenoteRequestBuilder) Operations()(*iac72d35f9c20cb360839134d86d535f87e1cdbe10b719f8c792d8017c06edd1d.OperationsRequestBuilder) {
    return iac72d35f9c20cb360839134d86d535f87e1cdbe10b719f8c792d8017c06edd1d.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.onenote.operations.item collection
func (m *OnenoteRequestBuilder) OperationsById(id string)(*i2b8ad7bd2ae87c92c8852bbeef996852d7685df93cb12b50d58ecc094ecfd124.OnenoteOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation_id"] = id
    }
    return i2b8ad7bd2ae87c92c8852bbeef996852d7685df93cb12b50d58ecc094ecfd124.NewOnenoteOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Pages the pages property
func (m *OnenoteRequestBuilder) Pages()(*iae2d6340b64994bda52b72186cf936e932db0b6d72196a85d99b8d63141db19b.PagesRequestBuilder) {
    return iae2d6340b64994bda52b72186cf936e932db0b6d72196a85d99b8d63141db19b.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.onenote.pages.item collection
func (m *OnenoteRequestBuilder) PagesById(id string)(*i9ffb03d414d6abd9f024bf2d1f19d7b7fa97959d1245465fdaef4d163ddd6de4.OnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return i9ffb03d414d6abd9f024bf2d1f19d7b7fa97959d1245465fdaef4d163ddd6de4.NewOnenotePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property onenote in me
func (m *OnenoteRequestBuilder) Patch(options *OnenoteRequestBuilderPatchOptions)(error) {
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
// Resources the resources property
func (m *OnenoteRequestBuilder) Resources()(*ic51747a34f2e47436a56cc441d712037e581ac2bd968645ce4fe67a401a2f4bf.ResourcesRequestBuilder) {
    return ic51747a34f2e47436a56cc441d712037e581ac2bd968645ce4fe67a401a2f4bf.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.onenote.resources.item collection
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*id67049e3fa5525a6ba50b6b8bc9f60107ac06f61256f8cf55bb655b4552ecd21.OnenoteResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource_id"] = id
    }
    return id67049e3fa5525a6ba50b6b8bc9f60107ac06f61256f8cf55bb655b4552ecd21.NewOnenoteResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SectionGroups the sectionGroups property
func (m *OnenoteRequestBuilder) SectionGroups()(*id71911f826ddd0ad6edad424ad81ad5b822e25039b7e26730ec6b2f6de37ad3b.SectionGroupsRequestBuilder) {
    return id71911f826ddd0ad6edad424ad81ad5b822e25039b7e26730ec6b2f6de37ad3b.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.onenote.sectionGroups.item collection
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*i1941888ff368eb1fd130194825ce82fb010616c26e7af09e43370fa6bb2cfd71.SectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return i1941888ff368eb1fd130194825ce82fb010616c26e7af09e43370fa6bb2cfd71.NewSectionGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sections the sections property
func (m *OnenoteRequestBuilder) Sections()(*i54a3bfcdd648d6ee47e87c311695b29c4737b8921bb8f2a506c0864844df5dba.SectionsRequestBuilder) {
    return i54a3bfcdd648d6ee47e87c311695b29c4737b8921bb8f2a506c0864844df5dba.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.onenote.sections.item collection
func (m *OnenoteRequestBuilder) SectionsById(id string)(*ibe60e1baafd5c0c7f830c971df51a300df66822516c0cfcd8a43fd2adc385c80.OnenoteSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return ibe60e1baafd5c0c7f830c971df51a300df66822516c0cfcd8a43fd2adc385c80.NewOnenoteSectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
