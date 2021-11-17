package schools

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i9bff590a460b52e5f40b634c15d18595cddb033ef6d43ef01edef967ac8013a3 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/schools/delta"
    ie96248ccc61767ca0f9ce692dcb287e1e3fe4d846626806247e1208a266f28b0 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/schools/ref"
)

// Builds and executes requests for operations under \education\classes\{educationClass-id}\schools
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
// All schools that this class is associated with. Nullable.
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
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass_id}/schools{?top,skip,search,filter,count,orderby,select,expand}";
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
// All schools that this class is associated with. Nullable.
// Parameters:
//  - options : Options for the request
func (m *SchoolsRequestBuilder) CreateGetRequestInformation(options *SchoolsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Builds and executes requests for operations under \education\classes\{educationClass-id}\schools\microsoft.graph.delta()
func (m *SchoolsRequestBuilder) Delta()(*i9bff590a460b52e5f40b634c15d18595cddb033ef6d43ef01edef967ac8013a3.DeltaRequestBuilder) {
    return i9bff590a460b52e5f40b634c15d18595cddb033ef6d43ef01edef967ac8013a3.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// All schools that this class is associated with. Nullable.
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
func (m *SchoolsRequestBuilder) Ref()(*ie96248ccc61767ca0f9ce692dcb287e1e3fe4d846626806247e1208a266f28b0.RefRequestBuilder) {
    return ie96248ccc61767ca0f9ce692dcb287e1e3fe4d846626806247e1208a266f28b0.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
