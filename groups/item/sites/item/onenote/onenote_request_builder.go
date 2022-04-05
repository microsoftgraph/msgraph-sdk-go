package onenote

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i1dc5a32fdef093887ac28cc0c8bca5a0d25d23501cd27dcca8cbc290540705dd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/sections"
    i441838dd29b242650bd7039dd08b5e34d6b39a9a3d74aa998b4ac41c5f3b31d8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/pages"
    i4833aa4798bac6e2d1b27413ff8bdbe8ce66a1ae991a6d4f674557800dd9314c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/resources"
    i51da7de877f44ebb046999f6c850aa6f2979b5de0ebf5c92cc602d63bb51d409 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/operations"
    i5ad7f337b9255100fb3a698996b5b25aa72d3af6a1db88a2befb9b86d7f0c125 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/notebooks"
    ib06c8196d06bc7fe6a64f371a7b1bddb2a7d8f44eaba7531b314fcef6f8bb977 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/sectiongroups"
    i31c8ff0ce5826f214b3c99ffed6b10e2491455c11d64a9546d47b02e96c08ce9 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/notebooks/item"
    i491bcdaec023edf3c376e0613168616caa148137e322b28771e11840fa698083 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/pages/item"
    i8c0b8b51b87672413a18d94f087c250716f466a20fcb74c70d8e99428e7884a4 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/sections/item"
    i9b5b1c5e7066c902b878df7cd7030266dc898dbfbba4fa5afb5d8d421c3182de "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/operations/item"
    icb75c59060050d546263152fcb819987bb25d9bc5948bba2c924282950950679 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/resources/item"
    ifbd90d50469a7f26b8a8adedfe4d30f41af421568ce824def28a45f309b51ec9 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote/sectiongroups/item"
)

// OnenoteRequestBuilder provides operations to manage the onenote property of the microsoft.graph.site entity.
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
// OnenoteRequestBuilderGetQueryParameters calls the OneNote service for notebook related operations.
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
    m.urlTemplate = "{+baseurl}/groups/{group_id}/sites/{site_id}/onenote{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property onenote for groups
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
// CreateGetRequestInformation calls the OneNote service for notebook related operations.
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
// CreatePatchRequestInformation update the navigation property onenote in groups
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
// Delete delete navigation property onenote for groups
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
// Get calls the OneNote service for notebook related operations.
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
func (m *OnenoteRequestBuilder) Notebooks()(*i5ad7f337b9255100fb3a698996b5b25aa72d3af6a1db88a2befb9b86d7f0c125.NotebooksRequestBuilder) {
    return i5ad7f337b9255100fb3a698996b5b25aa72d3af6a1db88a2befb9b86d7f0c125.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotebooksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.onenote.notebooks.item collection
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*i31c8ff0ce5826f214b3c99ffed6b10e2491455c11d64a9546d47b02e96c08ce9.NotebookItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook_id"] = id
    }
    return i31c8ff0ce5826f214b3c99ffed6b10e2491455c11d64a9546d47b02e96c08ce9.NewNotebookItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *OnenoteRequestBuilder) Operations()(*i51da7de877f44ebb046999f6c850aa6f2979b5de0ebf5c92cc602d63bb51d409.OperationsRequestBuilder) {
    return i51da7de877f44ebb046999f6c850aa6f2979b5de0ebf5c92cc602d63bb51d409.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.onenote.operations.item collection
func (m *OnenoteRequestBuilder) OperationsById(id string)(*i9b5b1c5e7066c902b878df7cd7030266dc898dbfbba4fa5afb5d8d421c3182de.OnenoteOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation_id"] = id
    }
    return i9b5b1c5e7066c902b878df7cd7030266dc898dbfbba4fa5afb5d8d421c3182de.NewOnenoteOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Pages the pages property
func (m *OnenoteRequestBuilder) Pages()(*i441838dd29b242650bd7039dd08b5e34d6b39a9a3d74aa998b4ac41c5f3b31d8.PagesRequestBuilder) {
    return i441838dd29b242650bd7039dd08b5e34d6b39a9a3d74aa998b4ac41c5f3b31d8.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.onenote.pages.item collection
func (m *OnenoteRequestBuilder) PagesById(id string)(*i491bcdaec023edf3c376e0613168616caa148137e322b28771e11840fa698083.OnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return i491bcdaec023edf3c376e0613168616caa148137e322b28771e11840fa698083.NewOnenotePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property onenote in groups
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
func (m *OnenoteRequestBuilder) Resources()(*i4833aa4798bac6e2d1b27413ff8bdbe8ce66a1ae991a6d4f674557800dd9314c.ResourcesRequestBuilder) {
    return i4833aa4798bac6e2d1b27413ff8bdbe8ce66a1ae991a6d4f674557800dd9314c.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.onenote.resources.item collection
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*icb75c59060050d546263152fcb819987bb25d9bc5948bba2c924282950950679.OnenoteResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource_id"] = id
    }
    return icb75c59060050d546263152fcb819987bb25d9bc5948bba2c924282950950679.NewOnenoteResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SectionGroups the sectionGroups property
func (m *OnenoteRequestBuilder) SectionGroups()(*ib06c8196d06bc7fe6a64f371a7b1bddb2a7d8f44eaba7531b314fcef6f8bb977.SectionGroupsRequestBuilder) {
    return ib06c8196d06bc7fe6a64f371a7b1bddb2a7d8f44eaba7531b314fcef6f8bb977.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.onenote.sectionGroups.item collection
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*ifbd90d50469a7f26b8a8adedfe4d30f41af421568ce824def28a45f309b51ec9.SectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return ifbd90d50469a7f26b8a8adedfe4d30f41af421568ce824def28a45f309b51ec9.NewSectionGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sections the sections property
func (m *OnenoteRequestBuilder) Sections()(*i1dc5a32fdef093887ac28cc0c8bca5a0d25d23501cd27dcca8cbc290540705dd.SectionsRequestBuilder) {
    return i1dc5a32fdef093887ac28cc0c8bca5a0d25d23501cd27dcca8cbc290540705dd.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.onenote.sections.item collection
func (m *OnenoteRequestBuilder) SectionsById(id string)(*i8c0b8b51b87672413a18d94f087c250716f466a20fcb74c70d8e99428e7884a4.OnenoteSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i8c0b8b51b87672413a18d94f087c250716f466a20fcb74c70d8e99428e7884a4.NewOnenoteSectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
