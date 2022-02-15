package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
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
    i8459e0e9fdbef8a8ad4c7a315789b0cca6cc0c5d53c4d9f685c083f3c7d8381d "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/columnlinks/item"
)

// ContentTypeRequestBuilder builds and executes requests for operations under \shares\{sharedDriveItem-id}\list\contentTypes\{contentType-id}
type ContentTypeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContentTypeRequestBuilderDeleteOptions options for Delete
type ContentTypeRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContentTypeRequestBuilderGetOptions options for Get
type ContentTypeRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContentTypeRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContentTypeRequestBuilderGetQueryParameters the collection of content types present in this list.
type ContentTypeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ContentTypeRequestBuilderPatchOptions options for Patch
type ContentTypeRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentType;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ContentTypeRequestBuilder) AssociateWithHubSites()(*i49bc8900bfdd8b3b1d3f1b52a9d8d36be89f0d252009782f35e12ba5c6fae675.AssociateWithHubSitesRequestBuilder) {
    return i49bc8900bfdd8b3b1d3f1b52a9d8d36be89f0d252009782f35e12ba5c6fae675.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Base()(*if20ce4a9c8987ec45c17d11c97d9255ead0b63c6bb3cb2a948e79d243584e4de.BaseRequestBuilder) {
    return if20ce4a9c8987ec45c17d11c97d9255ead0b63c6bb3cb2a948e79d243584e4de.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) BaseTypes()(*i4a1ce58e21fb551499684168220c8c6eee92bb5bc6a4b99639b513c5cfe71a1b.BaseTypesRequestBuilder) {
    return i4a1ce58e21fb551499684168220c8c6eee92bb5bc6a4b99639b513c5cfe71a1b.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnLinks()(*i0df9b813acca9f52cb58191b981da2631e2ce7ec1e0f39baea115430a4d91ec7.ColumnLinksRequestBuilder) {
    return i0df9b813acca9f52cb58191b981da2631e2ce7ec1e0f39baea115430a4d91ec7.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.shares.item.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeRequestBuilder) ColumnLinksById(id string)(*i8459e0e9fdbef8a8ad4c7a315789b0cca6cc0c5d53c4d9f685c083f3c7d8381d.ColumnLinkRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return i8459e0e9fdbef8a8ad4c7a315789b0cca6cc0c5d53c4d9f685c083f3c7d8381d.NewColumnLinkRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnPositions()(*i553035798d555384c1d78519561934573cbdb4b0f2bd9da038ce9a71c4991d05.ColumnPositionsRequestBuilder) {
    return i553035798d555384c1d78519561934573cbdb4b0f2bd9da038ce9a71c4991d05.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Columns()(*i1ec18c0a4fbb8f904f34e72c387a8703eb3af7790657ffef68bd11ec2a9905ef.ColumnsRequestBuilder) {
    return i1ec18c0a4fbb8f904f34e72c387a8703eb3af7790657ffef68bd11ec2a9905ef.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.shares.item.list.contentTypes.item.columns.item collection
func (m *ContentTypeRequestBuilder) ColumnsById(id string)(*i5a14dba7acc9a67d9b71d060c0e0de130dbeaf1185def884644520254fd0a2bf.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i5a14dba7acc9a67d9b71d060c0e0de130dbeaf1185def884644520254fd0a2bf.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeRequestBuilderInternal instantiates a new ContentTypeRequestBuilder and sets the default values.
func NewContentTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeRequestBuilder) {
    m := &ContentTypeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem_id}/list/contentTypes/{contentType_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContentTypeRequestBuilder instantiates a new ContentTypeRequestBuilder and sets the default values.
func NewContentTypeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContentTypeRequestBuilder) CopyToDefaultContentLocation()(*ibda66b96fbe7db9cdba0b0f6fb218ff35d517fc7b3ac53ebf9bdc4c45957d5c8.CopyToDefaultContentLocationRequestBuilder) {
    return ibda66b96fbe7db9cdba0b0f6fb218ff35d517fc7b3ac53ebf9bdc4c45957d5c8.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the collection of content types present in this list.
func (m *ContentTypeRequestBuilder) CreateDeleteRequestInformation(options *ContentTypeRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ContentTypeRequestBuilder) CreateGetRequestInformation(options *ContentTypeRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the collection of content types present in this list.
func (m *ContentTypeRequestBuilder) CreatePatchRequestInformation(options *ContentTypeRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the collection of content types present in this list.
func (m *ContentTypeRequestBuilder) Delete(options *ContentTypeRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of content types present in this list.
func (m *ContentTypeRequestBuilder) Get(options *ContentTypeRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentType, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewContentType() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentType), nil
}
// IsPublished builds and executes requests for operations under \shares\{sharedDriveItem-id}\list\contentTypes\{contentType-id}\microsoft.graph.isPublished()
func (m *ContentTypeRequestBuilder) IsPublished()(*i88052e6ec6afde3eb4c8a4e647edd6f1b4d0d4c30b1be1a17f71fdd90f644c90.IsPublishedRequestBuilder) {
    return i88052e6ec6afde3eb4c8a4e647edd6f1b4d0d4c30b1be1a17f71fdd90f644c90.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the collection of content types present in this list.
func (m *ContentTypeRequestBuilder) Patch(options *ContentTypeRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContentTypeRequestBuilder) Publish()(*i934cd417c891824167d128199e27f353bdac8d75faa96747cb1f2f67634a829e.PublishRequestBuilder) {
    return i934cd417c891824167d128199e27f353bdac8d75faa96747cb1f2f67634a829e.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Unpublish()(*i7f1730a5f677932cd121090203a3bc4480b7f148f3eae6dd7de05441a796ed7c.UnpublishRequestBuilder) {
    return i7f1730a5f677932cd121090203a3bc4480b7f148f3eae6dd7de05441a796ed7c.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
