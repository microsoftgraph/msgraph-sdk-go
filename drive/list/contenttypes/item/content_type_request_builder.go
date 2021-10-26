package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
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
    i15ae992f8c06ad8a7ba1ec98d7e1e47ad2aadb64d711a178a018a0753d67fbcf "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/columns/item"
    ib8f5fc0530361dc8f60f00c441136b066b0ef8c5599697f14d1092db89ae8857 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/columnlinks/item"
)

// Builds and executes requests for operations under \drive\list\contentTypes\{contentType-id}
type ContentTypeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The collection of content types present in this list.
type ContentTypeRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *ContentTypeRequestBuilder) AssociateWithHubSites()(*idb984764a96aa1929d9736abf97514ee454b1e8782c03102f624a72dec321d85.AssociateWithHubSitesRequestBuilder) {
    return idb984764a96aa1929d9736abf97514ee454b1e8782c03102f624a72dec321d85.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Base()(*i0b98bee81fa5a4fb2b5fa646275edac4f17fd1135e34963ffdb081964f70fa29.BaseRequestBuilder) {
    return i0b98bee81fa5a4fb2b5fa646275edac4f17fd1135e34963ffdb081964f70fa29.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) BaseTypes()(*ibbfe9e07ff81e3a1ac4b3ed908b90bd28a86fd235c5c1fcf568ba59eabeca6ec.BaseTypesRequestBuilder) {
    return ibbfe9e07ff81e3a1ac4b3ed908b90bd28a86fd235c5c1fcf568ba59eabeca6ec.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnLinks()(*i350e95307b635e7cbd9b80327f06108e3c2d26f1b80f2ffe02ab9e26040919c0.ColumnLinksRequestBuilder) {
    return i350e95307b635e7cbd9b80327f06108e3c2d26f1b80f2ffe02ab9e26040919c0.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.list.contentTypes.item.columnLinks.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContentTypeRequestBuilder) ColumnLinksById(id string)(*ib8f5fc0530361dc8f60f00c441136b066b0ef8c5599697f14d1092db89ae8857.ColumnLinkRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return ib8f5fc0530361dc8f60f00c441136b066b0ef8c5599697f14d1092db89ae8857.NewColumnLinkRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnPositions()(*i45c3970936281df10318dee6941b2dd766057e61fc3af1a2061b427dcc3f44b0.ColumnPositionsRequestBuilder) {
    return i45c3970936281df10318dee6941b2dd766057e61fc3af1a2061b427dcc3f44b0.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Columns()(*i37ed59fef414ed542b467a743d67ed5f7b35e828361fcee35f794973680f0273.ColumnsRequestBuilder) {
    return i37ed59fef414ed542b467a743d67ed5f7b35e828361fcee35f794973680f0273.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.list.contentTypes.item.columns.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContentTypeRequestBuilder) ColumnsById(id string)(*i15ae992f8c06ad8a7ba1ec98d7e1e47ad2aadb64d711a178a018a0753d67fbcf.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i15ae992f8c06ad8a7ba1ec98d7e1e47ad2aadb64d711a178a018a0753d67fbcf.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new ContentTypeRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewContentTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeRequestBuilder) {
    m := &ContentTypeRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/drive/list/contentTypes/{contentType_id}{?select,expand}";
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
func (m *ContentTypeRequestBuilder) CopyToDefaultContentLocation()(*i8a8731025cbe5699c843f367738d2c9ee2193b6d64796d2224611f93056402c4.CopyToDefaultContentLocationRequestBuilder) {
    return i8a8731025cbe5699c843f367738d2c9ee2193b6d64796d2224611f93056402c4.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The collection of content types present in this list.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *ContentTypeRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The collection of content types present in this list.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *ContentTypeRequestBuilder) CreateGetRequestInformation(q func (value *ContentTypeRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ContentTypeRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The collection of content types present in this list.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *ContentTypeRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentType, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The collection of content types present in this list.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ContentTypeRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
// The collection of content types present in this list.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ContentTypeRequestBuilder) Get(q func (value *ContentTypeRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentType, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewContentType() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentType), nil
}
// Builds and executes requests for operations under \drive\list\contentTypes\{contentType-id}\microsoft.graph.isPublished()
func (m *ContentTypeRequestBuilder) IsPublished()(*i0e4a4f160e936ed0b701995d280a973c677a1c132f35feaca1aae10c863ea1b4.IsPublishedRequestBuilder) {
    return i0e4a4f160e936ed0b701995d280a973c677a1c132f35feaca1aae10c863ea1b4.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The collection of content types present in this list.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ContentTypeRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentType, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContentTypeRequestBuilder) Publish()(*i2bab590cdf67b22b3d8348c3477a9ef52e970dff41681d07f85c1a7e2f18726d.PublishRequestBuilder) {
    return i2bab590cdf67b22b3d8348c3477a9ef52e970dff41681d07f85c1a7e2f18726d.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Unpublish()(*iaf680a4c4cb484ea82e42a31267fda11f665a0e1c57cc2bc8d76dacd06dbc5f6.UnpublishRequestBuilder) {
    return iaf680a4c4cb484ea82e42a31267fda11f665a0e1c57cc2bc8d76dacd06dbc5f6.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
