package base

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0f41d613e43ebb75a5256f4e9fa0a26801eb0f970049cd54b38553a93e4d7583 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/base/copytodefaultcontentlocation"
    i298a83c5dfdbdcc2fee8f2e4175f23ed520fffd5614aee42f5ff7a82103ab660 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/base/ref"
    i54f27f04a6c3f63e912c66b46a1861003db14caee53bd2407d8af0a587a9a0d9 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/base/publish"
    i662ea0f4d893f61f005e23e62b72156cd39ad88e8c990bc431b36ac5c1063c4a "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/base/associatewithhubsites"
    id9c5c5d747414e50788c8641fe3cbc3445ec4e625ea3e6a375edc67e81597e08 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/base/ispublished"
    idba48be55ae6c72be4696fd197b78f035d4e1f7f6b196a4abebb04b5054522b8 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/base/unpublish"
)

// Builds and executes requests for operations under \sites\{site-id}\lists\{list-id}\contentTypes\{contentType-id}\base
type BaseRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type BaseRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *BaseRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Parent contentType from which this content type is derived.
type BaseRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *BaseRequestBuilder) AssociateWithHubSites()(*i662ea0f4d893f61f005e23e62b72156cd39ad88e8c990bc431b36ac5c1063c4a.AssociateWithHubSitesRequestBuilder) {
    return i662ea0f4d893f61f005e23e62b72156cd39ad88e8c990bc431b36ac5c1063c4a.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new BaseRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewBaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseRequestBuilder) {
    m := &BaseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/lists/{list_id}/contentTypes/{contentType_id}/base{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new BaseRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewBaseRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBaseRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *BaseRequestBuilder) CopyToDefaultContentLocation()(*i0f41d613e43ebb75a5256f4e9fa0a26801eb0f970049cd54b38553a93e4d7583.CopyToDefaultContentLocationRequestBuilder) {
    return i0f41d613e43ebb75a5256f4e9fa0a26801eb0f970049cd54b38553a93e4d7583.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Parent contentType from which this content type is derived.
// Parameters:
//  - options : Options for the request
func (m *BaseRequestBuilder) CreateGetRequestInformation(options *BaseRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Parent contentType from which this content type is derived.
// Parameters:
//  - options : Options for the request
func (m *BaseRequestBuilder) Get(options *BaseRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentType, error) {
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
// Builds and executes requests for operations under \sites\{site-id}\lists\{list-id}\contentTypes\{contentType-id}\base\microsoft.graph.isPublished()
func (m *BaseRequestBuilder) IsPublished()(*id9c5c5d747414e50788c8641fe3cbc3445ec4e625ea3e6a375edc67e81597e08.IsPublishedRequestBuilder) {
    return id9c5c5d747414e50788c8641fe3cbc3445ec4e625ea3e6a375edc67e81597e08.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Publish()(*i54f27f04a6c3f63e912c66b46a1861003db14caee53bd2407d8af0a587a9a0d9.PublishRequestBuilder) {
    return i54f27f04a6c3f63e912c66b46a1861003db14caee53bd2407d8af0a587a9a0d9.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Ref()(*i298a83c5dfdbdcc2fee8f2e4175f23ed520fffd5614aee42f5ff7a82103ab660.RefRequestBuilder) {
    return i298a83c5dfdbdcc2fee8f2e4175f23ed520fffd5614aee42f5ff7a82103ab660.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Unpublish()(*idba48be55ae6c72be4696fd197b78f035d4e1f7f6b196a4abebb04b5054522b8.UnpublishRequestBuilder) {
    return idba48be55ae6c72be4696fd197b78f035d4e1f7f6b196a4abebb04b5054522b8.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
