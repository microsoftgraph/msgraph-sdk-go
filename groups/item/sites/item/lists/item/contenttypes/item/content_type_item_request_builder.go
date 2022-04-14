package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i1bf7e274532170f1bce30fa85891cd1060aaf4432c850c3e6f600c74225358dd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/publish"
    i4fd99e30bd7e1ec5146299a811b6699b7a48158aff453a0ca332de49c3afb49e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/columnlinks"
    i5980d4c6089817dc360b0282a3de8f025a947b489914b96333ae7bb03e59b88a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/base"
    i8d5d0e8815604e277a1a9310a20dd65c29f73cdd754a3ce39aac678b60860ca1 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/basetypes"
    i965a40c98e9504af392fa55c0e043f32c195f17afac01638dd363f4492bea4e9 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/ispublished"
    ibbce353b9cf3173510a3ab5b4232bab32639d0afdeb30cb6fa7be20ae2f9db7e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/copytodefaultcontentlocation"
    ic897225cd3a6c5e1636ce8c17f93a0c8b2150f74edeb8ac1991c2c3c584636f6 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/columnpositions"
    id483c2708d28edda81621fb56e0347a6ae606808dcbb1ae7f49bb596e0150406 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/columns"
    id4f65ae8ee4c57e2f7ab4326616b2d20b8bfcc68216f9b5f5cdbeaf6a6099b33 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/associatewithhubsites"
    iee2737a01517790f43d05681c34a6491d2832aa82ce6073946c0001bc2a818f9 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/unpublish"
    i1149c500ddccfb7c4db5f677250807d83df1d279a796576eb9613fedd3b2578f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/columnlinks/item"
    i93827d02ff2765cecf70f15c1b4ff586989d30c2b0051a346ee15cd23c2ba844 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/basetypes/item"
    icf1aa423f6e90fa0f5e91966c342d736b4bf0d2e71c543272f6a32ecd56802ce "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/columnpositions/item"
    ie281909f6939213db2ea9fee97bafef5874b4d351854a3f510c03c428183b137 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/columns/item"
)

// ContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type ContentTypeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ContentTypeItemRequestBuilderDeleteOptions options for Delete
type ContentTypeItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ContentTypeItemRequestBuilderGetOptions options for Get
type ContentTypeItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ContentTypeItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ContentTypeItemRequestBuilderGetQueryParameters the collection of content types present in this list.
type ContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ContentTypeItemRequestBuilderPatchOptions options for Patch
type ContentTypeItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AssociateWithHubSites the associateWithHubSites property
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*id4f65ae8ee4c57e2f7ab4326616b2d20b8bfcc68216f9b5f5cdbeaf6a6099b33.AssociateWithHubSitesRequestBuilder) {
    return id4f65ae8ee4c57e2f7ab4326616b2d20b8bfcc68216f9b5f5cdbeaf6a6099b33.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Base the base property
func (m *ContentTypeItemRequestBuilder) Base()(*i5980d4c6089817dc360b0282a3de8f025a947b489914b96333ae7bb03e59b88a.BaseRequestBuilder) {
    return i5980d4c6089817dc360b0282a3de8f025a947b489914b96333ae7bb03e59b88a.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypes the baseTypes property
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i8d5d0e8815604e277a1a9310a20dd65c29f73cdd754a3ce39aac678b60860ca1.BaseTypesRequestBuilder) {
    return i8d5d0e8815604e277a1a9310a20dd65c29f73cdd754a3ce39aac678b60860ca1.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.lists.item.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*i93827d02ff2765cecf70f15c1b4ff586989d30c2b0051a346ee15cd23c2ba844.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did1"] = id
    }
    return i93827d02ff2765cecf70f15c1b4ff586989d30c2b0051a346ee15cd23c2ba844.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnLinks the columnLinks property
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i4fd99e30bd7e1ec5146299a811b6699b7a48158aff453a0ca332de49c3afb49e.ColumnLinksRequestBuilder) {
    return i4fd99e30bd7e1ec5146299a811b6699b7a48158aff453a0ca332de49c3afb49e.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.lists.item.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i1149c500ddccfb7c4db5f677250807d83df1d279a796576eb9613fedd3b2578f.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink%2Did"] = id
    }
    return i1149c500ddccfb7c4db5f677250807d83df1d279a796576eb9613fedd3b2578f.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnPositions the columnPositions property
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*ic897225cd3a6c5e1636ce8c17f93a0c8b2150f74edeb8ac1991c2c3c584636f6.ColumnPositionsRequestBuilder) {
    return ic897225cd3a6c5e1636ce8c17f93a0c8b2150f74edeb8ac1991c2c3c584636f6.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.lists.item.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*icf1aa423f6e90fa0f5e91966c342d736b4bf0d2e71c543272f6a32ecd56802ce.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return icf1aa423f6e90fa0f5e91966c342d736b4bf0d2e71c543272f6a32ecd56802ce.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Columns the columns property
func (m *ContentTypeItemRequestBuilder) Columns()(*id483c2708d28edda81621fb56e0347a6ae606808dcbb1ae7f49bb596e0150406.ColumnsRequestBuilder) {
    return id483c2708d28edda81621fb56e0347a6ae606808dcbb1ae7f49bb596e0150406.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.lists.item.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*ie281909f6939213db2ea9fee97bafef5874b4d351854a3f510c03c428183b137.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return ie281909f6939213db2ea9fee97bafef5874b4d351854a3f510c03c428183b137.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/{contentType%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContentTypeItemRequestBuilder instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CopyToDefaultContentLocation the copyToDefaultContentLocation property
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*ibbce353b9cf3173510a3ab5b4232bab32639d0afdeb30cb6fa7be20ae2f9db7e.CopyToDefaultContentLocationRequestBuilder) {
    return ibbce353b9cf3173510a3ab5b4232bab32639d0afdeb30cb6fa7be20ae2f9db7e.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for groups
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformation(options *ContentTypeItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformation(options *ContentTypeItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property contentTypes in groups
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformation(options *ContentTypeItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property contentTypes for groups
func (m *ContentTypeItemRequestBuilder) Delete(options *ContentTypeItemRequestBuilderDeleteOptions)(error) {
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
// Get the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) Get(options *ContentTypeItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContentTypeFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
func (m *ContentTypeItemRequestBuilder) IsPublished()(*i965a40c98e9504af392fa55c0e043f32c195f17afac01638dd363f4492bea4e9.IsPublishedRequestBuilder) {
    return i965a40c98e9504af392fa55c0e043f32c195f17afac01638dd363f4492bea4e9.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in groups
func (m *ContentTypeItemRequestBuilder) Patch(options *ContentTypeItemRequestBuilderPatchOptions)(error) {
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
// Publish the publish property
func (m *ContentTypeItemRequestBuilder) Publish()(*i1bf7e274532170f1bce30fa85891cd1060aaf4432c850c3e6f600c74225358dd.PublishRequestBuilder) {
    return i1bf7e274532170f1bce30fa85891cd1060aaf4432c850c3e6f600c74225358dd.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unpublish the unpublish property
func (m *ContentTypeItemRequestBuilder) Unpublish()(*iee2737a01517790f43d05681c34a6491d2832aa82ce6073946c0001bc2a818f9.UnpublishRequestBuilder) {
    return iee2737a01517790f43d05681c34a6491d2832aa82ce6073946c0001bc2a818f9.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
