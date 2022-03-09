package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i0b98bee81fa5a4fb2b5fa646275edac4f17fd1135e34963ffdb081964f70fa29 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/base"
    i0e4a4f160e936ed0b701995d280a973c677a1c132f35feaca1aae10c863ea1b4 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/ispublished"
    i2bab590cdf67b22b3d8348c3477a9ef52e970dff41681d07f85c1a7e2f18726d "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/publish"
    i350e95307b635e7cbd9b80327f06108e3c2d26f1b80f2ffe02ab9e26040919c0 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/columnlinks"
    i37ed59fef414ed542b467a743d67ed5f7b35e828361fcee35f794973680f0273 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/columns"
    i45c3970936281df10318dee6941b2dd766057e61fc3af1a2061b427dcc3f44b0 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/columnpositions"
    i8a8731025cbe5699c843f367738d2c9ee2193b6d64796d2224611f93056402c4 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/copytodefaultcontentlocation"
    iaf680a4c4cb484ea82e42a31267fda11f665a0e1c57cc2bc8d76dacd06dbc5f6 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/unpublish"
    ibbfe9e07ff81e3a1ac4b3ed908b90bd28a86fd235c5c1fcf568ba59eabeca6ec "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/basetypes"
    idb984764a96aa1929d9736abf97514ee454b1e8782c03102f624a72dec321d85 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/associatewithhubsites"
    i12c62140c5b78ceed5111ee524200560ab5e66ec6923f649782e9d66e3cd955e "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/columnpositions/item"
    i15ae992f8c06ad8a7ba1ec98d7e1e47ad2aadb64d711a178a018a0753d67fbcf "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/columns/item"
    i27f7157489366635a4d889a21bcf438a31aaa8b9e8f1fc16c1fe83d3b58b85aa "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/basetypes/item"
    ib8f5fc0530361dc8f60f00c441136b066b0ef8c5599697f14d1092db89ae8857 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/columnlinks/item"
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
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*idb984764a96aa1929d9736abf97514ee454b1e8782c03102f624a72dec321d85.AssociateWithHubSitesRequestBuilder) {
    return idb984764a96aa1929d9736abf97514ee454b1e8782c03102f624a72dec321d85.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Base()(*i0b98bee81fa5a4fb2b5fa646275edac4f17fd1135e34963ffdb081964f70fa29.BaseRequestBuilder) {
    return i0b98bee81fa5a4fb2b5fa646275edac4f17fd1135e34963ffdb081964f70fa29.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*ibbfe9e07ff81e3a1ac4b3ed908b90bd28a86fd235c5c1fcf568ba59eabeca6ec.BaseTypesRequestBuilder) {
    return ibbfe9e07ff81e3a1ac4b3ed908b90bd28a86fd235c5c1fcf568ba59eabeca6ec.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.list.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*i27f7157489366635a4d889a21bcf438a31aaa8b9e8f1fc16c1fe83d3b58b85aa.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id1"] = id
    }
    return i27f7157489366635a4d889a21bcf438a31aaa8b9e8f1fc16c1fe83d3b58b85aa.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i350e95307b635e7cbd9b80327f06108e3c2d26f1b80f2ffe02ab9e26040919c0.ColumnLinksRequestBuilder) {
    return i350e95307b635e7cbd9b80327f06108e3c2d26f1b80f2ffe02ab9e26040919c0.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*ib8f5fc0530361dc8f60f00c441136b066b0ef8c5599697f14d1092db89ae8857.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return ib8f5fc0530361dc8f60f00c441136b066b0ef8c5599697f14d1092db89ae8857.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*i45c3970936281df10318dee6941b2dd766057e61fc3af1a2061b427dcc3f44b0.ColumnPositionsRequestBuilder) {
    return i45c3970936281df10318dee6941b2dd766057e61fc3af1a2061b427dcc3f44b0.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.list.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*i12c62140c5b78ceed5111ee524200560ab5e66ec6923f649782e9d66e3cd955e.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i12c62140c5b78ceed5111ee524200560ab5e66ec6923f649782e9d66e3cd955e.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Columns()(*i37ed59fef414ed542b467a743d67ed5f7b35e828361fcee35f794973680f0273.ColumnsRequestBuilder) {
    return i37ed59fef414ed542b467a743d67ed5f7b35e828361fcee35f794973680f0273.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.list.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i15ae992f8c06ad8a7ba1ec98d7e1e47ad2aadb64d711a178a018a0753d67fbcf.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i15ae992f8c06ad8a7ba1ec98d7e1e47ad2aadb64d711a178a018a0753d67fbcf.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/list/contentTypes/{contentType_id}{?select,expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*i8a8731025cbe5699c843f367738d2c9ee2193b6d64796d2224611f93056402c4.CopyToDefaultContentLocationRequestBuilder) {
    return i8a8731025cbe5699c843f367738d2c9ee2193b6d64796d2224611f93056402c4.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for drive
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
// CreatePatchRequestInformation update the navigation property contentTypes in drive
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
// Delete delete navigation property contentTypes for drive
func (m *ContentTypeItemRequestBuilder) Delete(options *ContentTypeItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
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
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateContentTypeFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
func (m *ContentTypeItemRequestBuilder) IsPublished()(*i0e4a4f160e936ed0b701995d280a973c677a1c132f35feaca1aae10c863ea1b4.IsPublishedRequestBuilder) {
    return i0e4a4f160e936ed0b701995d280a973c677a1c132f35feaca1aae10c863ea1b4.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in drive
func (m *ContentTypeItemRequestBuilder) Patch(options *ContentTypeItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContentTypeItemRequestBuilder) Publish()(*i2bab590cdf67b22b3d8348c3477a9ef52e970dff41681d07f85c1a7e2f18726d.PublishRequestBuilder) {
    return i2bab590cdf67b22b3d8348c3477a9ef52e970dff41681d07f85c1a7e2f18726d.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Unpublish()(*iaf680a4c4cb484ea82e42a31267fda11f665a0e1c57cc2bc8d76dacd06dbc5f6.UnpublishRequestBuilder) {
    return iaf680a4c4cb484ea82e42a31267fda11f665a0e1c57cc2bc8d76dacd06dbc5f6.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
