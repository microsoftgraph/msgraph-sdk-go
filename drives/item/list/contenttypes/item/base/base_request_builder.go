package base

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i07bdc42ac2dfc520f6f21d728420f445c545c09c072e332673bf36d1e557dabd "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/base/copytodefaultcontentlocation"
    i24022dc8f1d4fbcda6fccaefaa52cd5f07d66a7a26a683fb56943c8c02fff958 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/base/publish"
    i52fc1031503665ffb5435070345c2ab63a9250f772516bb30435d0c67a3f4889 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/base/ref"
    i6155334405b1fd4bff004bf74a3483dda04ffe0f347075a40980117443dd9c5b "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/base/unpublish"
    i9640ad4c5f221fa7cf3d22da3713c65bc604c71b2a426c9f22a9b7afcb36eb73 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/base/associatewithhubsites"
    ic781016c5523c999c4b1cd671cfb53ed6b10a54462724f42e2aab645055f1897 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/base/ispublished"
)

// BaseRequestBuilder builds and executes requests for operations under \drives\{drive-id}\list\contentTypes\{contentType-id}\base
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
    Select_escaped []string;
}
func (m *BaseRequestBuilder) AssociateWithHubSites()(*i9640ad4c5f221fa7cf3d22da3713c65bc604c71b2a426c9f22a9b7afcb36eb73.AssociateWithHubSitesRequestBuilder) {
    return i9640ad4c5f221fa7cf3d22da3713c65bc604c71b2a426c9f22a9b7afcb36eb73.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBaseRequestBuilderInternal instantiates a new BaseRequestBuilder and sets the default values.
func NewBaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseRequestBuilder) {
    m := &BaseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/list/contentTypes/{contentType_id}/base{?select,expand}";
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
func (m *BaseRequestBuilder) CopyToDefaultContentLocation()(*i07bdc42ac2dfc520f6f21d728420f445c545c09c072e332673bf36d1e557dabd.CopyToDefaultContentLocationRequestBuilder) {
    return i07bdc42ac2dfc520f6f21d728420f445c545c09c072e332673bf36d1e557dabd.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewContentType() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentType), nil
}
// IsPublished builds and executes requests for operations under \drives\{drive-id}\list\contentTypes\{contentType-id}\base\microsoft.graph.isPublished()
func (m *BaseRequestBuilder) IsPublished()(*ic781016c5523c999c4b1cd671cfb53ed6b10a54462724f42e2aab645055f1897.IsPublishedRequestBuilder) {
    return ic781016c5523c999c4b1cd671cfb53ed6b10a54462724f42e2aab645055f1897.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Publish()(*i24022dc8f1d4fbcda6fccaefaa52cd5f07d66a7a26a683fb56943c8c02fff958.PublishRequestBuilder) {
    return i24022dc8f1d4fbcda6fccaefaa52cd5f07d66a7a26a683fb56943c8c02fff958.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Ref()(*i52fc1031503665ffb5435070345c2ab63a9250f772516bb30435d0c67a3f4889.RefRequestBuilder) {
    return i52fc1031503665ffb5435070345c2ab63a9250f772516bb30435d0c67a3f4889.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Unpublish()(*i6155334405b1fd4bff004bf74a3483dda04ffe0f347075a40980117443dd9c5b.UnpublishRequestBuilder) {
    return i6155334405b1fd4bff004bf74a3483dda04ffe0f347075a40980117443dd9c5b.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
