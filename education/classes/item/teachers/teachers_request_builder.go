package teachers

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i60e69df95455a4cf0e9fa34ce11b749c74034755913c76c1052dc484024c615c "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/teachers/ref"
    ia879f83be5eb0abe21bec0d5b30131bf56ffb6a18677cf143109ff098ad5ea3a "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/teachers/delta"
)

// TeachersRequestBuilder builds and executes requests for operations under \education\classes\{educationClass-id}\teachers
type TeachersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TeachersRequestBuilderGetOptions options for Get
type TeachersRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TeachersRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TeachersRequestBuilderGetQueryParameters all teachers in the class. Nullable.
type TeachersRequestBuilderGetQueryParameters struct {
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
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// NewTeachersRequestBuilderInternal instantiates a new TeachersRequestBuilder and sets the default values.
func NewTeachersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeachersRequestBuilder) {
    m := &TeachersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass_id}/teachers{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeachersRequestBuilder instantiates a new TeachersRequestBuilder and sets the default values.
func NewTeachersRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeachersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeachersRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation all teachers in the class. Nullable.
func (m *TeachersRequestBuilder) CreateGetRequestInformation(options *TeachersRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delta builds and executes requests for operations under \education\classes\{educationClass-id}\teachers\microsoft.graph.delta()
func (m *TeachersRequestBuilder) Delta()(*ia879f83be5eb0abe21bec0d5b30131bf56ffb6a18677cf143109ff098ad5ea3a.DeltaRequestBuilder) {
    return ia879f83be5eb0abe21bec0d5b30131bf56ffb6a18677cf143109ff098ad5ea3a.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get all teachers in the class. Nullable.
func (m *TeachersRequestBuilder) Get(options *TeachersRequestBuilderGetOptions)(*TeachersResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeachersResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*TeachersResponse), nil
}
func (m *TeachersRequestBuilder) Ref()(*i60e69df95455a4cf0e9fa34ce11b749c74034755913c76c1052dc484024c615c.RefRequestBuilder) {
    return i60e69df95455a4cf0e9fa34ce11b749c74034755913c76c1052dc484024c615c.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
