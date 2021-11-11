package schools

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i31c55a4d38e851a2de14c96cf6d2fede7b768f082f53a4448032a19bba7c7c1e "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/schools/delta"
    i59640742c9e94ac289bae1f030f8b678df7249c049506acce82ced5002851d43 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/schools/ref"
)

// Builds and executes requests for operations under \education\users\{educationUser-id}\schools
type SchoolsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type SchoolsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SchoolsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Schools to which the user belongs. Nullable.
type SchoolsRequestBuilderGetQueryParameters struct {
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
// Instantiates a new SchoolsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSchoolsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SchoolsRequestBuilder) {
    m := &SchoolsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser_id}/schools{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new SchoolsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSchoolsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SchoolsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchoolsRequestBuilderInternal(urlParams, requestAdapter)
}
// Schools to which the user belongs. Nullable.
// Parameters:
//  - options : Options for the request
func (m *SchoolsRequestBuilder) CreateGetRequestInformation(options *SchoolsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(options.Q)
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
// Builds and executes requests for operations under \education\users\{educationUser-id}\schools\microsoft.graph.delta()
func (m *SchoolsRequestBuilder) Delta()(*i31c55a4d38e851a2de14c96cf6d2fede7b768f082f53a4448032a19bba7c7c1e.DeltaRequestBuilder) {
    return i31c55a4d38e851a2de14c96cf6d2fede7b768f082f53a4448032a19bba7c7c1e.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schools to which the user belongs. Nullable.
// Parameters:
//  - options : Options for the request
func (m *SchoolsRequestBuilder) Get(options *SchoolsRequestBuilderGetOptions)(*SchoolsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSchoolsResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*SchoolsResponse), nil
}
func (m *SchoolsRequestBuilder) Ref()(*i59640742c9e94ac289bae1f030f8b678df7249c049506acce82ced5002851d43.RefRequestBuilder) {
    return i59640742c9e94ac289bae1f030f8b678df7249c049506acce82ced5002851d43.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
