package contacts

import (
    i56415aaf9483a2e973df3b1a2fe488e7b2d8844c5d37e84bfa9789655c1de234 "github.com/microsoftgraph/msgraph-sdk-go/contacts/validateproperties"
    ibd6b6137a854fdc0082d98413fd3f4475189c1b69ecd839ed41d01e72ef32af8 "github.com/microsoftgraph/msgraph-sdk-go/contacts/getbyids"
    icd7c2e08bbb98a59371e3c25cbc42e41894aefcd3c07a9dda51cde0fbcbda0a1 "github.com/microsoftgraph/msgraph-sdk-go/contacts/getavailableextensionproperties"
    id6a5f8a255008661c4f3891226b4705b4315a990d81bf1bdc10e110af7ce62f9 "github.com/microsoftgraph/msgraph-sdk-go/contacts/delta"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// Builds and executes requests for operations under \contacts
type ContactsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type ContactsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContactsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get entities from contacts
type ContactsRequestBuilderGetQueryParameters struct {
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
// Options for Post
type ContactsRequestBuilderPostOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OrgContact;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new ContactsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewContactsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactsRequestBuilder) {
    m := &ContactsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contacts{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ContactsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewContactsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get entities from contacts
// Parameters:
//  - options : Options for the request
func (m *ContactsRequestBuilder) CreateGetRequestInformation(options *ContactsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Add new entity to contacts
// Parameters:
//  - options : Options for the request
func (m *ContactsRequestBuilder) CreatePostRequestInformation(options *ContactsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Builds and executes requests for operations under \contacts\microsoft.graph.delta()
func (m *ContactsRequestBuilder) Delta()(*id6a5f8a255008661c4f3891226b4705b4315a990d81bf1bdc10e110af7ce62f9.DeltaRequestBuilder) {
    return id6a5f8a255008661c4f3891226b4705b4315a990d81bf1bdc10e110af7ce62f9.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get entities from contacts
// Parameters:
//  - options : Options for the request
func (m *ContactsRequestBuilder) Get(options *ContactsRequestBuilderGetOptions)(*ContactsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContactsResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ContactsResponse), nil
}
func (m *ContactsRequestBuilder) GetAvailableExtensionProperties()(*icd7c2e08bbb98a59371e3c25cbc42e41894aefcd3c07a9dda51cde0fbcbda0a1.GetAvailableExtensionPropertiesRequestBuilder) {
    return icd7c2e08bbb98a59371e3c25cbc42e41894aefcd3c07a9dda51cde0fbcbda0a1.NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactsRequestBuilder) GetByIds()(*ibd6b6137a854fdc0082d98413fd3f4475189c1b69ecd839ed41d01e72ef32af8.GetByIdsRequestBuilder) {
    return ibd6b6137a854fdc0082d98413fd3f4475189c1b69ecd839ed41d01e72ef32af8.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Add new entity to contacts
// Parameters:
//  - options : Options for the request
func (m *ContactsRequestBuilder) Post(options *ContactsRequestBuilderPostOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OrgContact, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOrgContact() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OrgContact), nil
}
func (m *ContactsRequestBuilder) ValidateProperties()(*i56415aaf9483a2e973df3b1a2fe488e7b2d8844c5d37e84bfa9789655c1de234.ValidatePropertiesRequestBuilder) {
    return i56415aaf9483a2e973df3b1a2fe488e7b2d8844c5d37e84bfa9789655c1de234.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
