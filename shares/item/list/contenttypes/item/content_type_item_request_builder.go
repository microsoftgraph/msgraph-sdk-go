package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0df9b813acca9f52cb58191b981da2631e2ce7ec1e0f39baea115430a4d91ec7 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/columnlinks"
    i1ec18c0a4fbb8f904f34e72c387a8703eb3af7790657ffef68bd11ec2a9905ef "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/columns"
    i49bc8900bfdd8b3b1d3f1b52a9d8d36be89f0d252009782f35e12ba5c6fae675 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/associatewithhubsites"
    i4a1ce58e21fb551499684168220c8c6eee92bb5bc6a4b99639b513c5cfe71a1b "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/basetypes"
    i553035798d555384c1d78519561934573cbdb4b0f2bd9da038ce9a71c4991d05 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/columnpositions"
    i7f1730a5f677932cd121090203a3bc4480b7f148f3eae6dd7de05441a796ed7c "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/unpublish"
    i88052e6ec6afde3eb4c8a4e647edd6f1b4d0d4c30b1be1a17f71fdd90f644c90 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/ispublished"
    i934cd417c891824167d128199e27f353bdac8d75faa96747cb1f2f67634a829e "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/publish"
    ibda66b96fbe7db9cdba0b0f6fb218ff35d517fc7b3ac53ebf9bdc4c45957d5c8 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/copytodefaultcontentlocation"
    if20ce4a9c8987ec45c17d11c97d9255ead0b63c6bb3cb2a948e79d243584e4de "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/base"
    i5a14dba7acc9a67d9b71d060c0e0de130dbeaf1185def884644520254fd0a2bf "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/columns/item"
    i684c0ce66c376c68b7eda568c4a0e38ff6f9905b99e380e5c7d6735c0ce5813b "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/basetypes/item"
    i798aa8d3eb2beecb739a9f0d6422dbcd303b5e2935628722cd84d0e6fbea4bb1 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/columnpositions/item"
    i8459e0e9fdbef8a8ad4c7a315789b0cca6cc0c5d53c4d9f685c083f3c7d8381d "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/columnlinks/item"
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
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i49bc8900bfdd8b3b1d3f1b52a9d8d36be89f0d252009782f35e12ba5c6fae675.AssociateWithHubSitesRequestBuilder) {
    return i49bc8900bfdd8b3b1d3f1b52a9d8d36be89f0d252009782f35e12ba5c6fae675.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Base()(*if20ce4a9c8987ec45c17d11c97d9255ead0b63c6bb3cb2a948e79d243584e4de.BaseRequestBuilder) {
    return if20ce4a9c8987ec45c17d11c97d9255ead0b63c6bb3cb2a948e79d243584e4de.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i4a1ce58e21fb551499684168220c8c6eee92bb5bc6a4b99639b513c5cfe71a1b.BaseTypesRequestBuilder) {
    return i4a1ce58e21fb551499684168220c8c6eee92bb5bc6a4b99639b513c5cfe71a1b.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.shares.item.list.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*i684c0ce66c376c68b7eda568c4a0e38ff6f9905b99e380e5c7d6735c0ce5813b.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id1"] = id
    }
    return i684c0ce66c376c68b7eda568c4a0e38ff6f9905b99e380e5c7d6735c0ce5813b.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i0df9b813acca9f52cb58191b981da2631e2ce7ec1e0f39baea115430a4d91ec7.ColumnLinksRequestBuilder) {
    return i0df9b813acca9f52cb58191b981da2631e2ce7ec1e0f39baea115430a4d91ec7.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.shares.item.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i8459e0e9fdbef8a8ad4c7a315789b0cca6cc0c5d53c4d9f685c083f3c7d8381d.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return i8459e0e9fdbef8a8ad4c7a315789b0cca6cc0c5d53c4d9f685c083f3c7d8381d.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*i553035798d555384c1d78519561934573cbdb4b0f2bd9da038ce9a71c4991d05.ColumnPositionsRequestBuilder) {
    return i553035798d555384c1d78519561934573cbdb4b0f2bd9da038ce9a71c4991d05.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.shares.item.list.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*i798aa8d3eb2beecb739a9f0d6422dbcd303b5e2935628722cd84d0e6fbea4bb1.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i798aa8d3eb2beecb739a9f0d6422dbcd303b5e2935628722cd84d0e6fbea4bb1.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Columns()(*i1ec18c0a4fbb8f904f34e72c387a8703eb3af7790657ffef68bd11ec2a9905ef.ColumnsRequestBuilder) {
    return i1ec18c0a4fbb8f904f34e72c387a8703eb3af7790657ffef68bd11ec2a9905ef.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.shares.item.list.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i5a14dba7acc9a67d9b71d060c0e0de130dbeaf1185def884644520254fd0a2bf.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i5a14dba7acc9a67d9b71d060c0e0de130dbeaf1185def884644520254fd0a2bf.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem_id}/list/contentTypes/{contentType_id}{?select,expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*ibda66b96fbe7db9cdba0b0f6fb218ff35d517fc7b3ac53ebf9bdc4c45957d5c8.CopyToDefaultContentLocationRequestBuilder) {
    return ibda66b96fbe7db9cdba0b0f6fb218ff35d517fc7b3ac53ebf9bdc4c45957d5c8.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for shares
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
// CreatePatchRequestInformation update the navigation property contentTypes in shares
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
// Delete delete navigation property contentTypes for shares
func (m *ContentTypeItemRequestBuilder) Delete(options *ContentTypeItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
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
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateContentTypeFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
func (m *ContentTypeItemRequestBuilder) IsPublished()(*i88052e6ec6afde3eb4c8a4e647edd6f1b4d0d4c30b1be1a17f71fdd90f644c90.IsPublishedRequestBuilder) {
    return i88052e6ec6afde3eb4c8a4e647edd6f1b4d0d4c30b1be1a17f71fdd90f644c90.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in shares
func (m *ContentTypeItemRequestBuilder) Patch(options *ContentTypeItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContentTypeItemRequestBuilder) Publish()(*i934cd417c891824167d128199e27f353bdac8d75faa96747cb1f2f67634a829e.PublishRequestBuilder) {
    return i934cd417c891824167d128199e27f353bdac8d75faa96747cb1f2f67634a829e.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i7f1730a5f677932cd121090203a3bc4480b7f148f3eae6dd7de05441a796ed7c.UnpublishRequestBuilder) {
    return i7f1730a5f677932cd121090203a3bc4480b7f148f3eae6dd7de05441a796ed7c.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
