package transitivemembers

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i38c46eb33f2d36dc59f0ae913c99556a8126694bb4665e210ef3d49ba179a204 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/transitivemembers/ref"
)

// TransitiveMembersRequestBuilder builds and executes requests for operations under \groups\{group-id}\transitiveMembers
type TransitiveMembersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TransitiveMembersRequestBuilderGetOptions options for Get
type TransitiveMembersRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TransitiveMembersRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TransitiveMembersRequestBuilderGetQueryParameters get transitiveMembers from groups
type TransitiveMembersRequestBuilderGetQueryParameters struct {
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
// NewTransitiveMembersRequestBuilderInternal instantiates a new TransitiveMembersRequestBuilder and sets the default values.
func NewTransitiveMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TransitiveMembersRequestBuilder) {
    m := &TransitiveMembersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/transitiveMembers{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTransitiveMembersRequestBuilder instantiates a new TransitiveMembersRequestBuilder and sets the default values.
func NewTransitiveMembersRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TransitiveMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTransitiveMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get transitiveMembers from groups
func (m *TransitiveMembersRequestBuilder) CreateGetRequestInformation(options *TransitiveMembersRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get transitiveMembers from groups
func (m *TransitiveMembersRequestBuilder) Get(options *TransitiveMembersRequestBuilderGetOptions)(*TransitiveMembersResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTransitiveMembersResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*TransitiveMembersResponse), nil
}
func (m *TransitiveMembersRequestBuilder) Ref()(*i38c46eb33f2d36dc59f0ae913c99556a8126694bb4665e210ef3d49ba179a204.RefRequestBuilder) {
    return i38c46eb33f2d36dc59f0ae913c99556a8126694bb4665e210ef3d49ba179a204.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
