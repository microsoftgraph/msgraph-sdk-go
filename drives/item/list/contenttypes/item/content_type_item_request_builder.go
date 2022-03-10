package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i05d9428439d35e52cb2ce2ba5c3c7be1cd010f18c1f354cb859cec016c46ad56 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/columns"
    i3060c755378bef956392441c325290f6ea178acc828301e16119a470f9192ffe "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/copytodefaultcontentlocation"
    i3b8ee80d094176753a5065e732b8fe371a09654d0bd329c3db9b51d7203fd853 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/publish"
    i43f44f7d18ca3faffb3453324604e7cc5c111263d5e22f2c445d29e2eb8218e9 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/base"
    i4f8ea4604387554b83af4225698949b2df687d40f43c743c6c46462ded20d3f9 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/ispublished"
    i582b9b1101bb3849bfd2fa23dfed26494d51af218a6e426e47df1c8cfbfa5554 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/basetypes"
    i7539bdc5f9b19a28873835c6b9f2f1c5a304eaf2c846cb1c025784c17b81241f "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/associatewithhubsites"
    ia9f43a3a505aa44f104fa23f977e28b34f2065aad158175d2503a0cc251a6f38 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/unpublish"
    ib1e7c6dac6ea9f3518585a39144c5b871a51fc839296cea92c1bd56467d4b734 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/columnpositions"
    ibb8519e802c21fcd2fd2594af481ca8371cfb5ded0d4b614dbcf185b47bf3d59 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/columnlinks"
    i2f2b0214ea0fbc249fc7576a95de950d0d9844dc794add5aad2d189d732227ed "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/columnpositions/item"
    i47a65a6f94278006557555cc162ea76107be8a3f3deedeec2b3f8fd743a4e3ef "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/columns/item"
    i849a233b89dacb52395d4caffa0734d8ee95f4a3f072a6bb085645174e84959c "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/basetypes/item"
    i9d04d98cacfd97673fbb11e0191a5450a9dc4bb49e2376d572d9a0639d40abde "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/columnlinks/item"
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
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i7539bdc5f9b19a28873835c6b9f2f1c5a304eaf2c846cb1c025784c17b81241f.AssociateWithHubSitesRequestBuilder) {
    return i7539bdc5f9b19a28873835c6b9f2f1c5a304eaf2c846cb1c025784c17b81241f.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Base()(*i43f44f7d18ca3faffb3453324604e7cc5c111263d5e22f2c445d29e2eb8218e9.BaseRequestBuilder) {
    return i43f44f7d18ca3faffb3453324604e7cc5c111263d5e22f2c445d29e2eb8218e9.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i582b9b1101bb3849bfd2fa23dfed26494d51af218a6e426e47df1c8cfbfa5554.BaseTypesRequestBuilder) {
    return i582b9b1101bb3849bfd2fa23dfed26494d51af218a6e426e47df1c8cfbfa5554.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.list.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*i849a233b89dacb52395d4caffa0734d8ee95f4a3f072a6bb085645174e84959c.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id1"] = id
    }
    return i849a233b89dacb52395d4caffa0734d8ee95f4a3f072a6bb085645174e84959c.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*ibb8519e802c21fcd2fd2594af481ca8371cfb5ded0d4b614dbcf185b47bf3d59.ColumnLinksRequestBuilder) {
    return ibb8519e802c21fcd2fd2594af481ca8371cfb5ded0d4b614dbcf185b47bf3d59.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i9d04d98cacfd97673fbb11e0191a5450a9dc4bb49e2376d572d9a0639d40abde.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return i9d04d98cacfd97673fbb11e0191a5450a9dc4bb49e2376d572d9a0639d40abde.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*ib1e7c6dac6ea9f3518585a39144c5b871a51fc839296cea92c1bd56467d4b734.ColumnPositionsRequestBuilder) {
    return ib1e7c6dac6ea9f3518585a39144c5b871a51fc839296cea92c1bd56467d4b734.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.list.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*i2f2b0214ea0fbc249fc7576a95de950d0d9844dc794add5aad2d189d732227ed.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i2f2b0214ea0fbc249fc7576a95de950d0d9844dc794add5aad2d189d732227ed.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Columns()(*i05d9428439d35e52cb2ce2ba5c3c7be1cd010f18c1f354cb859cec016c46ad56.ColumnsRequestBuilder) {
    return i05d9428439d35e52cb2ce2ba5c3c7be1cd010f18c1f354cb859cec016c46ad56.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.list.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i47a65a6f94278006557555cc162ea76107be8a3f3deedeec2b3f8fd743a4e3ef.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i47a65a6f94278006557555cc162ea76107be8a3f3deedeec2b3f8fd743a4e3ef.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/list/contentTypes/{contentType_id}{?select,expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*i3060c755378bef956392441c325290f6ea178acc828301e16119a470f9192ffe.CopyToDefaultContentLocationRequestBuilder) {
    return i3060c755378bef956392441c325290f6ea178acc828301e16119a470f9192ffe.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for drives
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
// CreatePatchRequestInformation update the navigation property contentTypes in drives
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
// Delete delete navigation property contentTypes for drives
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
func (m *ContentTypeItemRequestBuilder) IsPublished()(*i4f8ea4604387554b83af4225698949b2df687d40f43c743c6c46462ded20d3f9.IsPublishedRequestBuilder) {
    return i4f8ea4604387554b83af4225698949b2df687d40f43c743c6c46462ded20d3f9.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in drives
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
func (m *ContentTypeItemRequestBuilder) Publish()(*i3b8ee80d094176753a5065e732b8fe371a09654d0bd329c3db9b51d7203fd853.PublishRequestBuilder) {
    return i3b8ee80d094176753a5065e732b8fe371a09654d0bd329c3db9b51d7203fd853.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Unpublish()(*ia9f43a3a505aa44f104fa23f977e28b34f2065aad158175d2503a0cc251a6f38.UnpublishRequestBuilder) {
    return ia9f43a3a505aa44f104fa23f977e28b34f2065aad158175d2503a0cc251a6f38.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
