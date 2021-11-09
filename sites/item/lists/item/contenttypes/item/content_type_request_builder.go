package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0adbaf2bb9aa63a5e604a5fa371b072519032b803f737301b0b52921cc1d3a35 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/unpublish"
    i1a7695ad2fabf538093408204c95d6189330219c052c9f612b44a85734601b1c "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/ispublished"
    i278379c29b4a8817ca06b917afe4d8c15eeafd2fd0f5640b5269d826d7333584 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/columnlinks"
    i3a448b37181c7678264208ecae87fe0e94cc7ea93cf9ce950d8c69413aaec6db "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/associatewithhubsites"
    i6e3f707a321441169d0528a8f9d4e00d29f3fdc3ce146c739c6537af4b41e473 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/base"
    i91d7fb036b4d816c4d7231c78f8e9680172ba5a52ca457e5ae197cfc261d3361 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/publish"
    ib477091a8a0eb17f6038d6809d39677533b56c9feeae3f59b9fc6af836d4766e "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/columnpositions"
    iba33625fbf7112f94866f86d5bf89ec9e16e3e3d0e4f825b748afa7b67c6fb15 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/copytodefaultcontentlocation"
    ida9964b630004248401c792f6f5212cb1ddb19922f1fb202be62edb471410340 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/basetypes"
    iff966f1c31329b7dc1c1986153d73717d35370de251cc9894dff12910c064f2a "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/columns"
    i2150d109d18cf045a75c8d5444fead62bb38570ae3bd07afea3c979c0b6f7858 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/columns/item"
    ia4e76c4ea17d224738d3bd333eae6742adbcd53e331ba5959a9aed631af5d7c2 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/columnlinks/item"
)

// Builds and executes requests for operations under \sites\{site-id}\lists\{list-id}\contentTypes\{contentType-id}
type ContentTypeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ContentTypeRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
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
// The collection of content types present in this list.
type ContentTypeRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
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
func (m *ContentTypeRequestBuilder) AssociateWithHubSites()(*i3a448b37181c7678264208ecae87fe0e94cc7ea93cf9ce950d8c69413aaec6db.AssociateWithHubSitesRequestBuilder) {
    return i3a448b37181c7678264208ecae87fe0e94cc7ea93cf9ce950d8c69413aaec6db.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Base()(*i6e3f707a321441169d0528a8f9d4e00d29f3fdc3ce146c739c6537af4b41e473.BaseRequestBuilder) {
    return i6e3f707a321441169d0528a8f9d4e00d29f3fdc3ce146c739c6537af4b41e473.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) BaseTypes()(*ida9964b630004248401c792f6f5212cb1ddb19922f1fb202be62edb471410340.BaseTypesRequestBuilder) {
    return ida9964b630004248401c792f6f5212cb1ddb19922f1fb202be62edb471410340.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnLinks()(*i278379c29b4a8817ca06b917afe4d8c15eeafd2fd0f5640b5269d826d7333584.ColumnLinksRequestBuilder) {
    return i278379c29b4a8817ca06b917afe4d8c15eeafd2fd0f5640b5269d826d7333584.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.lists.item.contentTypes.item.columnLinks.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContentTypeRequestBuilder) ColumnLinksById(id string)(*ia4e76c4ea17d224738d3bd333eae6742adbcd53e331ba5959a9aed631af5d7c2.ColumnLinkRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return ia4e76c4ea17d224738d3bd333eae6742adbcd53e331ba5959a9aed631af5d7c2.NewColumnLinkRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnPositions()(*ib477091a8a0eb17f6038d6809d39677533b56c9feeae3f59b9fc6af836d4766e.ColumnPositionsRequestBuilder) {
    return ib477091a8a0eb17f6038d6809d39677533b56c9feeae3f59b9fc6af836d4766e.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Columns()(*iff966f1c31329b7dc1c1986153d73717d35370de251cc9894dff12910c064f2a.ColumnsRequestBuilder) {
    return iff966f1c31329b7dc1c1986153d73717d35370de251cc9894dff12910c064f2a.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.lists.item.contentTypes.item.columns.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContentTypeRequestBuilder) ColumnsById(id string)(*i2150d109d18cf045a75c8d5444fead62bb38570ae3bd07afea3c979c0b6f7858.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i2150d109d18cf045a75c8d5444fead62bb38570ae3bd07afea3c979c0b6f7858.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new ContentTypeRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewContentTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeRequestBuilder) {
    m := &ContentTypeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/lists/{list_id}/contentTypes/{contentType_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ContentTypeRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewContentTypeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContentTypeRequestBuilder) CopyToDefaultContentLocation()(*iba33625fbf7112f94866f86d5bf89ec9e16e3e3d0e4f825b748afa7b67c6fb15.CopyToDefaultContentLocationRequestBuilder) {
    return iba33625fbf7112f94866f86d5bf89ec9e16e3e3d0e4f825b748afa7b67c6fb15.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The collection of content types present in this list.
// Parameters:
//  - options : Options for the request
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
// The collection of content types present in this list.
// Parameters:
//  - options : Options for the request
func (m *ContentTypeRequestBuilder) CreateGetRequestInformation(options *ContentTypeRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// The collection of content types present in this list.
// Parameters:
//  - options : Options for the request
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
// The collection of content types present in this list.
// Parameters:
//  - options : Options for the request
func (m *ContentTypeRequestBuilder) Delete(options *ContentTypeRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// The collection of content types present in this list.
// Parameters:
//  - options : Options for the request
func (m *ContentTypeRequestBuilder) Get(options *ContentTypeRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentType, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewContentType() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentType), nil
}
// Builds and executes requests for operations under \sites\{site-id}\lists\{list-id}\contentTypes\{contentType-id}\microsoft.graph.isPublished()
func (m *ContentTypeRequestBuilder) IsPublished()(*i1a7695ad2fabf538093408204c95d6189330219c052c9f612b44a85734601b1c.IsPublishedRequestBuilder) {
    return i1a7695ad2fabf538093408204c95d6189330219c052c9f612b44a85734601b1c.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The collection of content types present in this list.
// Parameters:
//  - options : Options for the request
func (m *ContentTypeRequestBuilder) Patch(options *ContentTypeRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContentTypeRequestBuilder) Publish()(*i91d7fb036b4d816c4d7231c78f8e9680172ba5a52ca457e5ae197cfc261d3361.PublishRequestBuilder) {
    return i91d7fb036b4d816c4d7231c78f8e9680172ba5a52ca457e5ae197cfc261d3361.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Unpublish()(*i0adbaf2bb9aa63a5e604a5fa371b072519032b803f737301b0b52921cc1d3a35.UnpublishRequestBuilder) {
    return i0adbaf2bb9aa63a5e604a5fa371b072519032b803f737301b0b52921cc1d3a35.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
