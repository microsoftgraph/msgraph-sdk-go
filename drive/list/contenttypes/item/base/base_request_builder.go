package base

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i079860cbe13058cd118170891f87551a3c9b5088d55d36a85ec4a073431a7d63 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/base/associatewithhubsites"
    i886cf40d7d56f9258ff0fe7066e5e5ca8eeb0d9cf60782ad9d0f640674c4f202 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/base/ref"
    i8ac43ceefdbeb993d1246bee07a9e84697bcc562527580e8fc05d8d7bd3b2057 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/base/ispublished"
    id25f80f3c54b25d2032c03fe7c91be340ea06671ae2a0b7b344b88f798ec40f2 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/base/publish"
    iddbe26954f00db92fceb2c7dd0a8f1a7fb906190e8154f86afa932b5aa772609 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/base/copytodefaultcontentlocation"
    ie54539325a708f46f600b1a1cf834b3caafbcd1695e53878df17ae63943818bc "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item/base/unpublish"
)

// BaseRequestBuilder builds and executes requests for operations under \drive\list\contentTypes\{contentType-id}\base
type BaseRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BaseRequestBuilderGetOptions options for Get
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
// BaseRequestBuilderGetQueryParameters parent contentType from which this content type is derived.
type BaseRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
func (m *BaseRequestBuilder) AssociateWithHubSites()(*i079860cbe13058cd118170891f87551a3c9b5088d55d36a85ec4a073431a7d63.AssociateWithHubSitesRequestBuilder) {
    return i079860cbe13058cd118170891f87551a3c9b5088d55d36a85ec4a073431a7d63.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBaseRequestBuilderInternal instantiates a new BaseRequestBuilder and sets the default values.
func NewBaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseRequestBuilder) {
    m := &BaseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/list/contentTypes/{contentType_id}/base{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBaseRequestBuilder instantiates a new BaseRequestBuilder and sets the default values.
func NewBaseRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBaseRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *BaseRequestBuilder) CopyToDefaultContentLocation()(*iddbe26954f00db92fceb2c7dd0a8f1a7fb906190e8154f86afa932b5aa772609.CopyToDefaultContentLocationRequestBuilder) {
    return iddbe26954f00db92fceb2c7dd0a8f1a7fb906190e8154f86afa932b5aa772609.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation parent contentType from which this content type is derived.
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
// Get parent contentType from which this content type is derived.
func (m *BaseRequestBuilder) Get(options *BaseRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentType, error) {
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
// IsPublished builds and executes requests for operations under \drive\list\contentTypes\{contentType-id}\base\microsoft.graph.isPublished()
func (m *BaseRequestBuilder) IsPublished()(*i8ac43ceefdbeb993d1246bee07a9e84697bcc562527580e8fc05d8d7bd3b2057.IsPublishedRequestBuilder) {
    return i8ac43ceefdbeb993d1246bee07a9e84697bcc562527580e8fc05d8d7bd3b2057.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Publish()(*id25f80f3c54b25d2032c03fe7c91be340ea06671ae2a0b7b344b88f798ec40f2.PublishRequestBuilder) {
    return id25f80f3c54b25d2032c03fe7c91be340ea06671ae2a0b7b344b88f798ec40f2.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Ref()(*i886cf40d7d56f9258ff0fe7066e5e5ca8eeb0d9cf60782ad9d0f640674c4f202.RefRequestBuilder) {
    return i886cf40d7d56f9258ff0fe7066e5e5ca8eeb0d9cf60782ad9d0f640674c4f202.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Unpublish()(*ie54539325a708f46f600b1a1cf834b3caafbcd1695e53878df17ae63943818bc.UnpublishRequestBuilder) {
    return ie54539325a708f46f600b1a1cf834b3caafbcd1695e53878df17ae63943818bc.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
