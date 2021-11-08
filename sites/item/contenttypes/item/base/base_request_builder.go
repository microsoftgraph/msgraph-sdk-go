package base

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i016659afa037990cd2b4c9c2d75c899b57b8832eb50ec4a55fca3ec39777fe79 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/base/unpublish"
    i3a3960bea2237464248f789220d6aa417a8f98ca84aaf2206ba377283e81d87a "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/base/ref"
    i5f43292b379c598605c02c3341a9cd03235d89c9edb778813c3909fb5ab4515b "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/base/associatewithhubsites"
    ia7002c94a0af2a5b7a665f396a598e30903acb571020f2b0a80d0a3c2021e33f "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/base/publish"
    ib05cb7d2d9b6f0e63c5218c33f9ba6750c8337f35f8f73f8f26e339c357da1ce "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/base/ispublished"
    icf04b80c1e65629c90c62e890266ef6edd310d5450faf2b98319ee20c44acadc "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/base/copytodefaultcontentlocation"
)

// Builds and executes requests for operations under \sites\{site-id}\contentTypes\{contentType-id}\base
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
func (m *BaseRequestBuilder) AssociateWithHubSites()(*i5f43292b379c598605c02c3341a9cd03235d89c9edb778813c3909fb5ab4515b.AssociateWithHubSitesRequestBuilder) {
    return i5f43292b379c598605c02c3341a9cd03235d89c9edb778813c3909fb5ab4515b.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new BaseRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewBaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseRequestBuilder) {
    m := &BaseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/contentTypes/{contentType_id}/base{?select,expand}";
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
func (m *BaseRequestBuilder) CopyToDefaultContentLocation()(*icf04b80c1e65629c90c62e890266ef6edd310d5450faf2b98319ee20c44acadc.CopyToDefaultContentLocationRequestBuilder) {
    return icf04b80c1e65629c90c62e890266ef6edd310d5450faf2b98319ee20c44acadc.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// Builds and executes requests for operations under \sites\{site-id}\contentTypes\{contentType-id}\base\microsoft.graph.isPublished()
func (m *BaseRequestBuilder) IsPublished()(*ib05cb7d2d9b6f0e63c5218c33f9ba6750c8337f35f8f73f8f26e339c357da1ce.IsPublishedRequestBuilder) {
    return ib05cb7d2d9b6f0e63c5218c33f9ba6750c8337f35f8f73f8f26e339c357da1ce.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Publish()(*ia7002c94a0af2a5b7a665f396a598e30903acb571020f2b0a80d0a3c2021e33f.PublishRequestBuilder) {
    return ia7002c94a0af2a5b7a665f396a598e30903acb571020f2b0a80d0a3c2021e33f.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Ref()(*i3a3960bea2237464248f789220d6aa417a8f98ca84aaf2206ba377283e81d87a.RefRequestBuilder) {
    return i3a3960bea2237464248f789220d6aa417a8f98ca84aaf2206ba377283e81d87a.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Unpublish()(*i016659afa037990cd2b4c9c2d75c899b57b8832eb50ec4a55fca3ec39777fe79.UnpublishRequestBuilder) {
    return i016659afa037990cd2b4c9c2d75c899b57b8832eb50ec4a55fca3ec39777fe79.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
