package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i0ede09ae8367dd5c31dd4a3d4897afa20a790181f02805a095f055e76fe2c387 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/contenttypes/item/basetypes"
    i41d83938157760e988cd228cf1244dd7bf0a549d5435f863193bdd932fd68ba3 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/contenttypes/item/ispublished"
    i44e7e3d2deb30875ef4628b9609c18dc4d1908db455833604327226140d028ac "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/contenttypes/item/copytodefaultcontentlocation"
    i45172cf09e83bad281fcb6464a1bbe805ea0f0d44e695167c8ad2581ac05202c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/contenttypes/item/associatewithhubsites"
    i531499f2acb1bde0ca2cc736a2d39c930eeffe76c5a01e85f0bfba643a194a20 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/contenttypes/item/columns"
    i6281d1793bf2139c3f8f2bc02b8b8f50ba4a3db6fff4c1154b5d7f8db7cc7c29 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/contenttypes/item/columnpositions"
    i9ba6603d8f89296fc650b4ef3660c9e1f76ea2ee3638792514c371a5df7c7b42 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/contenttypes/item/unpublish"
    ib86359ce3c7056dc7b940eab81d8b7cf2ba8d5db34be3a488f755fc360c2dfcc "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/contenttypes/item/columnlinks"
    id847dfab08bdf9d61ee185a46e5358145c28a6c00c1acbc9fb2ff61f2c03c09d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/contenttypes/item/base"
    ie3d120eb9ebeb3f728edfc28599120c13836d5b0e69365485148d7ea38fd6b82 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/contenttypes/item/publish"
    i0545e3b14912db622ebe37c26a41b32e2d67d6c9b1e93252f04bc6a5f5ac4975 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/contenttypes/item/columnlinks/item"
    i25b5535e0a33ffe2264128e8503e833333e122474edd346c75abcd795cff67e4 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/contenttypes/item/columns/item"
    i814e6afe9257c8f1c82fb630146a8daf0f670848f46c4bf867cb3102d5e84a28 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/contenttypes/item/basetypes/item"
    if5f72596c440ff55e8dab37638fea02c1a89816e858394131bb84f94ddf431db "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/contenttypes/item/columnpositions/item"
)

// ContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.site entity.
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
// ContentTypeItemRequestBuilderGetQueryParameters the collection of content types defined for this site.
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
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i45172cf09e83bad281fcb6464a1bbe805ea0f0d44e695167c8ad2581ac05202c.AssociateWithHubSitesRequestBuilder) {
    return i45172cf09e83bad281fcb6464a1bbe805ea0f0d44e695167c8ad2581ac05202c.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Base()(*id847dfab08bdf9d61ee185a46e5358145c28a6c00c1acbc9fb2ff61f2c03c09d.BaseRequestBuilder) {
    return id847dfab08bdf9d61ee185a46e5358145c28a6c00c1acbc9fb2ff61f2c03c09d.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i0ede09ae8367dd5c31dd4a3d4897afa20a790181f02805a095f055e76fe2c387.BaseTypesRequestBuilder) {
    return i0ede09ae8367dd5c31dd4a3d4897afa20a790181f02805a095f055e76fe2c387.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*i814e6afe9257c8f1c82fb630146a8daf0f670848f46c4bf867cb3102d5e84a28.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id1"] = id
    }
    return i814e6afe9257c8f1c82fb630146a8daf0f670848f46c4bf867cb3102d5e84a28.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*ib86359ce3c7056dc7b940eab81d8b7cf2ba8d5db34be3a488f755fc360c2dfcc.ColumnLinksRequestBuilder) {
    return ib86359ce3c7056dc7b940eab81d8b7cf2ba8d5db34be3a488f755fc360c2dfcc.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i0545e3b14912db622ebe37c26a41b32e2d67d6c9b1e93252f04bc6a5f5ac4975.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return i0545e3b14912db622ebe37c26a41b32e2d67d6c9b1e93252f04bc6a5f5ac4975.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*i6281d1793bf2139c3f8f2bc02b8b8f50ba4a3db6fff4c1154b5d7f8db7cc7c29.ColumnPositionsRequestBuilder) {
    return i6281d1793bf2139c3f8f2bc02b8b8f50ba4a3db6fff4c1154b5d7f8db7cc7c29.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*if5f72596c440ff55e8dab37638fea02c1a89816e858394131bb84f94ddf431db.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return if5f72596c440ff55e8dab37638fea02c1a89816e858394131bb84f94ddf431db.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Columns()(*i531499f2acb1bde0ca2cc736a2d39c930eeffe76c5a01e85f0bfba643a194a20.ColumnsRequestBuilder) {
    return i531499f2acb1bde0ca2cc736a2d39c930eeffe76c5a01e85f0bfba643a194a20.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i25b5535e0a33ffe2264128e8503e833333e122474edd346c75abcd795cff67e4.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i25b5535e0a33ffe2264128e8503e833333e122474edd346c75abcd795cff67e4.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/sites/{site_id}/contentTypes/{contentType_id}{?select,expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*i44e7e3d2deb30875ef4628b9609c18dc4d1908db455833604327226140d028ac.CopyToDefaultContentLocationRequestBuilder) {
    return i44e7e3d2deb30875ef4628b9609c18dc4d1908db455833604327226140d028ac.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// CreateGetRequestInformation the collection of content types defined for this site.
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
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of content types defined for this site.
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
func (m *ContentTypeItemRequestBuilder) IsPublished()(*i41d83938157760e988cd228cf1244dd7bf0a549d5435f863193bdd932fd68ba3.IsPublishedRequestBuilder) {
    return i41d83938157760e988cd228cf1244dd7bf0a549d5435f863193bdd932fd68ba3.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in groups
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
func (m *ContentTypeItemRequestBuilder) Publish()(*ie3d120eb9ebeb3f728edfc28599120c13836d5b0e69365485148d7ea38fd6b82.PublishRequestBuilder) {
    return ie3d120eb9ebeb3f728edfc28599120c13836d5b0e69365485148d7ea38fd6b82.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i9ba6603d8f89296fc650b4ef3660c9e1f76ea2ee3638792514c371a5df7c7b42.UnpublishRequestBuilder) {
    return i9ba6603d8f89296fc650b4ef3660c9e1f76ea2ee3638792514c371a5df7c7b42.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
