package teachers

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i60e69df95455a4cf0e9fa34ce11b749c74034755913c76c1052dc484024c615c "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/teachers/ref"
    ia879f83be5eb0abe21bec0d5b30131bf56ffb6a18677cf143109ff098ad5ea3a "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/teachers/delta"
)

// Builds and executes requests for operations under \education\classes\{educationClass-id}\teachers
type TeachersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// All teachers in the class. Nullable.
type TeachersRequestBuilderGetQueryParameters struct {
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
// Instantiates a new TeachersRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTeachersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeachersRequestBuilder) {
    m := &TeachersRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/education/classes/{educationClass_id}/teachers{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new TeachersRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTeachersRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeachersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeachersRequestBuilderInternal(urlParams, requestAdapter)
}
// All teachers in the class. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *TeachersRequestBuilder) CreateGetRequestInformation(q func (value *TeachersRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(TeachersRequestBuilderGetQueryParameters)
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
// Builds and executes requests for operations under \education\classes\{educationClass-id}\teachers\microsoft.graph.delta()
func (m *TeachersRequestBuilder) Delta()(*ia879f83be5eb0abe21bec0d5b30131bf56ffb6a18677cf143109ff098ad5ea3a.DeltaRequestBuilder) {
    return ia879f83be5eb0abe21bec0d5b30131bf56ffb6a18677cf143109ff098ad5ea3a.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// All teachers in the class. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *TeachersRequestBuilder) Get(q func (value *TeachersRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*TeachersResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeachersResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*TeachersResponse), nil
}
func (m *TeachersRequestBuilder) Ref()(*i60e69df95455a4cf0e9fa34ce11b749c74034755913c76c1052dc484024c615c.RefRequestBuilder) {
    return i60e69df95455a4cf0e9fa34ce11b749c74034755913c76c1052dc484024c615c.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
