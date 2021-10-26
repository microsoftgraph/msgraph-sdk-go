package messages

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0a1ac23d2ad4f7421bfe2ad68e1ae9861244745d3a0eb9b2339dfa4bf574093b "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/messages/favorite"
    i184768f41f49a81af236b2a9bc61c9295feb546cfe69781d725d312259ea39b5 "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/messages/markunread"
    i34e6b5bfcdf3235a943d804a8018471597527eea943054880b2e367767db738f "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/messages/markread"
    i6fe59a721c674e36b7fd2811e2df3f93a4dd9bf30a3958c160ecfa9b1bde6ad7 "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/messages/unfavorite"
    i7c8002d50d3d01307f13e1746331c6c44c1d8adb28af73e57a6058f41851d2c0 "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/messages/archive"
    if4711e0cd2dcd5bcd3b2fe5ff1a931184507edf584a1b39b0489f9d289e9acff "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/messages/unarchive"
)

// Builds and executes requests for operations under \admin\serviceAnnouncement\messages
type MessagesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
type MessagesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select_escaped []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
func (m *MessagesRequestBuilder) Archive()(*i7c8002d50d3d01307f13e1746331c6c44c1d8adb28af73e57a6058f41851d2c0.ArchiveRequestBuilder) {
    return i7c8002d50d3d01307f13e1746331c6c44c1d8adb28af73e57a6058f41851d2c0.NewArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new MessagesRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessagesRequestBuilder) {
    m := &MessagesRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/admin/serviceAnnouncement/messages{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new MessagesRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewMessagesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessagesRequestBuilderInternal(urlParams, requestAdapter)
}
// A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *MessagesRequestBuilder) CreateGetRequestInformation(q func (value *MessagesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(MessagesRequestBuilderGetQueryParameters)
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
// A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *MessagesRequestBuilder) CreatePostRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServiceUpdateMessage, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
func (m *MessagesRequestBuilder) Favorite()(*i0a1ac23d2ad4f7421bfe2ad68e1ae9861244745d3a0eb9b2339dfa4bf574093b.FavoriteRequestBuilder) {
    return i0a1ac23d2ad4f7421bfe2ad68e1ae9861244745d3a0eb9b2339dfa4bf574093b.NewFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *MessagesRequestBuilder) Get(q func (value *MessagesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*MessagesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMessagesResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*MessagesResponse), nil
}
func (m *MessagesRequestBuilder) MarkRead()(*i34e6b5bfcdf3235a943d804a8018471597527eea943054880b2e367767db738f.MarkReadRequestBuilder) {
    return i34e6b5bfcdf3235a943d804a8018471597527eea943054880b2e367767db738f.NewMarkReadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessagesRequestBuilder) MarkUnread()(*i184768f41f49a81af236b2a9bc61c9295feb546cfe69781d725d312259ea39b5.MarkUnreadRequestBuilder) {
    return i184768f41f49a81af236b2a9bc61c9295feb546cfe69781d725d312259ea39b5.NewMarkUnreadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *MessagesRequestBuilder) Post(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServiceUpdateMessage, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServiceUpdateMessage, error) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewServiceUpdateMessage() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServiceUpdateMessage), nil
}
func (m *MessagesRequestBuilder) Unarchive()(*if4711e0cd2dcd5bcd3b2fe5ff1a931184507edf584a1b39b0489f9d289e9acff.UnarchiveRequestBuilder) {
    return if4711e0cd2dcd5bcd3b2fe5ff1a931184507edf584a1b39b0489f9d289e9acff.NewUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessagesRequestBuilder) Unfavorite()(*i6fe59a721c674e36b7fd2811e2df3f93a4dd9bf30a3958c160ecfa9b1bde6ad7.UnfavoriteRequestBuilder) {
    return i6fe59a721c674e36b7fd2811e2df3f93a4dd9bf30a3958c160ecfa9b1bde6ad7.NewUnfavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
