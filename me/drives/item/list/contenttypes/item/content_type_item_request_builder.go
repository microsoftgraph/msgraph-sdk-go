package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i40b40c14afe20846df7697563943ff4147fcf3513ce68277ca2de88e464c473d "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/columnlinks"
    i53f314f0c3675cac45d21567fca93ab1edb06078259cfaea249d2ec6d2c1fdc3 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/copytodefaultcontentlocation"
    i9ce8000fb091cd3a5804314f05a98b6a927489500c5f2f1144c7f99e190c4067 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/associatewithhubsites"
    iaa8801033902add70125689951f2f36de2346e26374e4686d5f9c9de194cc05c "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/columns"
    iab9d147d4fb3d030a6b4abf8d305a14336264a5196d363e98b75cbf009a5334b "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/ispublished"
    iad6947c553ec5a4ce509828e095618d07a6d717e1de82221d777154d9192e208 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/unpublish"
    ic7b6bd6de956f12a8a38cb22bd58e71544ae32b7980a25b0536b301ca19a7bed "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/basetypes"
    idc99e113446b35e57939736e74d537663450beb1feaebdb6606c46591d5b5fb9 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/base"
    ide567c8b90b031fe74c6884dcbedb4274d796cf26b54c0e2d34d12cfef78aa17 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/columnpositions"
    ifebd77c70e2b2b28ca8e8b783ee20b065141c868e8a8be9109b2e89c805d5cf1 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/publish"
    i46bd691a897439fa677fdb51d5324696838d658e10fc675de7973db14cd34306 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/columnpositions/item"
    i637aca2c30344f22f5f4c7820f1574d5fe884d7b3bdb14a043cdcf1fd79e1dbb "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/columns/item"
    i6e6310aa8e8c7f0e53fc51025248ae3a6a4a30cb8bc899c5eae4a34bcf3b7dfb "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/columnlinks/item"
    if0d46cb9718fba761345686cf114220325cd78ab599073b77cfccc0543a3036f "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/basetypes/item"
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
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i9ce8000fb091cd3a5804314f05a98b6a927489500c5f2f1144c7f99e190c4067.AssociateWithHubSitesRequestBuilder) {
    return i9ce8000fb091cd3a5804314f05a98b6a927489500c5f2f1144c7f99e190c4067.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Base()(*idc99e113446b35e57939736e74d537663450beb1feaebdb6606c46591d5b5fb9.BaseRequestBuilder) {
    return idc99e113446b35e57939736e74d537663450beb1feaebdb6606c46591d5b5fb9.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*ic7b6bd6de956f12a8a38cb22bd58e71544ae32b7980a25b0536b301ca19a7bed.BaseTypesRequestBuilder) {
    return ic7b6bd6de956f12a8a38cb22bd58e71544ae32b7980a25b0536b301ca19a7bed.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.list.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*if0d46cb9718fba761345686cf114220325cd78ab599073b77cfccc0543a3036f.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id1"] = id
    }
    return if0d46cb9718fba761345686cf114220325cd78ab599073b77cfccc0543a3036f.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i40b40c14afe20846df7697563943ff4147fcf3513ce68277ca2de88e464c473d.ColumnLinksRequestBuilder) {
    return i40b40c14afe20846df7697563943ff4147fcf3513ce68277ca2de88e464c473d.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i6e6310aa8e8c7f0e53fc51025248ae3a6a4a30cb8bc899c5eae4a34bcf3b7dfb.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return i6e6310aa8e8c7f0e53fc51025248ae3a6a4a30cb8bc899c5eae4a34bcf3b7dfb.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*ide567c8b90b031fe74c6884dcbedb4274d796cf26b54c0e2d34d12cfef78aa17.ColumnPositionsRequestBuilder) {
    return ide567c8b90b031fe74c6884dcbedb4274d796cf26b54c0e2d34d12cfef78aa17.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.list.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*i46bd691a897439fa677fdb51d5324696838d658e10fc675de7973db14cd34306.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i46bd691a897439fa677fdb51d5324696838d658e10fc675de7973db14cd34306.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Columns()(*iaa8801033902add70125689951f2f36de2346e26374e4686d5f9c9de194cc05c.ColumnsRequestBuilder) {
    return iaa8801033902add70125689951f2f36de2346e26374e4686d5f9c9de194cc05c.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.list.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i637aca2c30344f22f5f4c7820f1574d5fe884d7b3bdb14a043cdcf1fd79e1dbb.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i637aca2c30344f22f5f4c7820f1574d5fe884d7b3bdb14a043cdcf1fd79e1dbb.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive_id}/list/contentTypes/{contentType_id}{?select,expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*i53f314f0c3675cac45d21567fca93ab1edb06078259cfaea249d2ec6d2c1fdc3.CopyToDefaultContentLocationRequestBuilder) {
    return i53f314f0c3675cac45d21567fca93ab1edb06078259cfaea249d2ec6d2c1fdc3.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for me
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
// CreatePatchRequestInformation update the navigation property contentTypes in me
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
// Delete delete navigation property contentTypes for me
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
func (m *ContentTypeItemRequestBuilder) IsPublished()(*iab9d147d4fb3d030a6b4abf8d305a14336264a5196d363e98b75cbf009a5334b.IsPublishedRequestBuilder) {
    return iab9d147d4fb3d030a6b4abf8d305a14336264a5196d363e98b75cbf009a5334b.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in me
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
func (m *ContentTypeItemRequestBuilder) Publish()(*ifebd77c70e2b2b28ca8e8b783ee20b065141c868e8a8be9109b2e89c805d5cf1.PublishRequestBuilder) {
    return ifebd77c70e2b2b28ca8e8b783ee20b065141c868e8a8be9109b2e89c805d5cf1.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Unpublish()(*iad6947c553ec5a4ce509828e095618d07a6d717e1de82221d777154d9192e208.UnpublishRequestBuilder) {
    return iad6947c553ec5a4ce509828e095618d07a6d717e1de82221d777154d9192e208.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
