package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContentTypeItemRequestBuilderDeleteOptions options for Delete
type ContentTypeItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContentTypeItemRequestBuilderGetOptions options for Get
type ContentTypeItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContentTypeItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContentTypeItemRequestBuilderGetQueryParameters the collection of content types present in this list.
type ContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ContentTypeItemRequestBuilderPatchOptions options for Patch
type ContentTypeItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentTypeable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*id4f65ae8ee4c57e2f7ab4326616b2d20b8bfcc68216f9b5f5cdbeaf6a6099b33.AssociateWithHubSitesRequestBuilder) {
    return id4f65ae8ee4c57e2f7ab4326616b2d20b8bfcc68216f9b5f5cdbeaf6a6099b33.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Base()(*i5980d4c6089817dc360b0282a3de8f025a947b489914b96333ae7bb03e59b88a.BaseRequestBuilder) {
    return i5980d4c6089817dc360b0282a3de8f025a947b489914b96333ae7bb03e59b88a.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["contentType_id1"] = id
    }
    return i93827d02ff2765cecf70f15c1b4ff586989d30c2b0051a346ee15cd23c2ba844.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["columnLink_id"] = id
    }
    return i1149c500ddccfb7c4db5f677250807d83df1d279a796576eb9613fedd3b2578f.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["columnDefinition_id"] = id
    }
    return icf1aa423f6e90fa0f5e91966c342d736b4bf0d2e71c543272f6a32ecd56802ce.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["columnDefinition_id"] = id
    }
    return ie281909f6939213db2ea9fee97bafef5874b4d351854a3f510c03c428183b137.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/sites/{site_id}/lists/{list_id}/contentTypes/{contentType_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContentTypeItemRequestBuilder instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*ibbce353b9cf3173510a3ab5b4232bab32639d0afdeb30cb6fa7be20ae2f9db7e.CopyToDefaultContentLocationRequestBuilder) {
    return ibbce353b9cf3173510a3ab5b4232bab32639d0afdeb30cb6fa7be20ae2f9db7e.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for groups
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformation(options *ContentTypeItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformation(options *ContentTypeItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property contentTypes in groups
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformation(options *ContentTypeItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
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
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) Get(options *ContentTypeItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateContentTypeFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentTypeable), nil
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
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContentTypeItemRequestBuilder) Publish()(*i1bf7e274532170f1bce30fa85891cd1060aaf4432c850c3e6f600c74225358dd.PublishRequestBuilder) {
    return i1bf7e274532170f1bce30fa85891cd1060aaf4432c850c3e6f600c74225358dd.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Unpublish()(*iee2737a01517790f43d05681c34a6491d2832aa82ce6073946c0001bc2a818f9.UnpublishRequestBuilder) {
    return iee2737a01517790f43d05681c34a6491d2832aa82ce6073946c0001bc2a818f9.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
