package base

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i33e21fbb7cc711988c7e5185df0630a4391e5502f821c5bc48a2f03e02f58418 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/base/ref"
    i5b1252a152d28480d7cd282f8e48e03c320e810cd1c1b0ddff56bbc863e78957 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/base/unpublish"
    iab46021f29f0137c239df157977a40072476e6e29535aea4f6d7b4a1d3d9c00e "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/base/associatewithhubsites"
    ibb709e977091722307e148234d2fd040df7821f88cb53176880432c206587f79 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/base/copytodefaultcontentlocation"
    ic411a64481a6bb2e7667659a622f5cce7499310ad9d9d93ee8bcae4782528b45 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/base/publish"
    ieae5c63a5818e285af36c2aea41680b3995f15175255dae27d689674f00a14a8 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/base/ispublished"
)

// Builds and executes requests for operations under \shares\{sharedDriveItem-id}\list\contentTypes\{contentType-id}\base
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
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *BaseRequestBuilder) AssociateWithHubSites()(*iab46021f29f0137c239df157977a40072476e6e29535aea4f6d7b4a1d3d9c00e.AssociateWithHubSitesRequestBuilder) {
    return iab46021f29f0137c239df157977a40072476e6e29535aea4f6d7b4a1d3d9c00e.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new BaseRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewBaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseRequestBuilder) {
    m := &BaseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem_id}/list/contentTypes/{contentType_id}/base{?select,expand}";
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
func (m *BaseRequestBuilder) CopyToDefaultContentLocation()(*ibb709e977091722307e148234d2fd040df7821f88cb53176880432c206587f79.CopyToDefaultContentLocationRequestBuilder) {
    return ibb709e977091722307e148234d2fd040df7821f88cb53176880432c206587f79.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// Builds and executes requests for operations under \shares\{sharedDriveItem-id}\list\contentTypes\{contentType-id}\base\microsoft.graph.isPublished()
func (m *BaseRequestBuilder) IsPublished()(*ieae5c63a5818e285af36c2aea41680b3995f15175255dae27d689674f00a14a8.IsPublishedRequestBuilder) {
    return ieae5c63a5818e285af36c2aea41680b3995f15175255dae27d689674f00a14a8.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Publish()(*ic411a64481a6bb2e7667659a622f5cce7499310ad9d9d93ee8bcae4782528b45.PublishRequestBuilder) {
    return ic411a64481a6bb2e7667659a622f5cce7499310ad9d9d93ee8bcae4782528b45.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Ref()(*i33e21fbb7cc711988c7e5185df0630a4391e5502f821c5bc48a2f03e02f58418.RefRequestBuilder) {
    return i33e21fbb7cc711988c7e5185df0630a4391e5502f821c5bc48a2f03e02f58418.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Unpublish()(*i5b1252a152d28480d7cd282f8e48e03c320e810cd1c1b0ddff56bbc863e78957.UnpublishRequestBuilder) {
    return i5b1252a152d28480d7cd282f8e48e03c320e810cd1c1b0ddff56bbc863e78957.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
