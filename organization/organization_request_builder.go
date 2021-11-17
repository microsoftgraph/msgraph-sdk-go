package organization

import (
    i0abaee8a058d22e1d581799c8740be61e9b6d10d3cc9df8a86b2c94d7c5d3836 "github.com/microsoftgraph/msgraph-sdk-go/organization/getavailableextensionproperties"
    i18ef64004d4e7ee2e8b15bbed5d6f937fc44f7035c463548e93e34e31361ffa4 "github.com/microsoftgraph/msgraph-sdk-go/organization/getbyids"
    i9a82bd7e2eb74542be8be13d437c57a66f6f2c71a200bd1483ca2b221ae825f2 "github.com/microsoftgraph/msgraph-sdk-go/organization/validateproperties"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// Builds and executes requests for operations under \organization
type OrganizationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type OrganizationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OrganizationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get entities from organization
type OrganizationRequestBuilderGetQueryParameters struct {
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
type OrganizationRequestBuilderPostOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Organization;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new OrganizationRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOrganizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OrganizationRequestBuilder) {
    m := &OrganizationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new OrganizationRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOrganizationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OrganizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get entities from organization
// Parameters:
//  - options : Options for the request
func (m *OrganizationRequestBuilder) CreateGetRequestInformation(options *OrganizationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Add new entity to organization
// Parameters:
//  - options : Options for the request
func (m *OrganizationRequestBuilder) CreatePostRequestInformation(options *OrganizationRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get entities from organization
// Parameters:
//  - options : Options for the request
func (m *OrganizationRequestBuilder) Get(options *OrganizationRequestBuilderGetOptions)(*OrganizationResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOrganizationResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*OrganizationResponse), nil
}
func (m *OrganizationRequestBuilder) GetAvailableExtensionProperties()(*i0abaee8a058d22e1d581799c8740be61e9b6d10d3cc9df8a86b2c94d7c5d3836.GetAvailableExtensionPropertiesRequestBuilder) {
    return i0abaee8a058d22e1d581799c8740be61e9b6d10d3cc9df8a86b2c94d7c5d3836.NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) GetByIds()(*i18ef64004d4e7ee2e8b15bbed5d6f937fc44f7035c463548e93e34e31361ffa4.GetByIdsRequestBuilder) {
    return i18ef64004d4e7ee2e8b15bbed5d6f937fc44f7035c463548e93e34e31361ffa4.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Add new entity to organization
// Parameters:
//  - options : Options for the request
func (m *OrganizationRequestBuilder) Post(options *OrganizationRequestBuilderPostOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Organization, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOrganization() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Organization), nil
}
func (m *OrganizationRequestBuilder) ValidateProperties()(*i9a82bd7e2eb74542be8be13d437c57a66f6f2c71a200bd1483ca2b221ae825f2.ValidatePropertiesRequestBuilder) {
    return i9a82bd7e2eb74542be8be13d437c57a66f6f2c71a200bd1483ca2b221ae825f2.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
