package taughtclasses

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i9a82ba081a87a0b82fe7518f2699a91f6cdabe7560d9041a4bbd68ca17595c04 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/taughtclasses/delta"
    ibfebe8e1821c93dda8997c543dd7161494427c674556776df0eee3b6720e6f3c "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/taughtclasses/ref"
)

// TaughtClassesRequestBuilder builds and executes requests for operations under \education\users\{educationUser-id}\taughtClasses
type TaughtClassesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TaughtClassesRequestBuilderGetOptions options for Get
type TaughtClassesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TaughtClassesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TaughtClassesRequestBuilderGetQueryParameters classes for which the user is a teacher.
type TaughtClassesRequestBuilderGetQueryParameters struct {
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
// NewTaughtClassesRequestBuilderInternal instantiates a new TaughtClassesRequestBuilder and sets the default values.
func NewTaughtClassesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TaughtClassesRequestBuilder) {
    m := &TaughtClassesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser_id}/taughtClasses{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTaughtClassesRequestBuilder instantiates a new TaughtClassesRequestBuilder and sets the default values.
func NewTaughtClassesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TaughtClassesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTaughtClassesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation classes for which the user is a teacher.
func (m *TaughtClassesRequestBuilder) CreateGetRequestInformation(options *TaughtClassesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delta builds and executes requests for operations under \education\users\{educationUser-id}\taughtClasses\microsoft.graph.delta()
func (m *TaughtClassesRequestBuilder) Delta()(*i9a82ba081a87a0b82fe7518f2699a91f6cdabe7560d9041a4bbd68ca17595c04.DeltaRequestBuilder) {
    return i9a82ba081a87a0b82fe7518f2699a91f6cdabe7560d9041a4bbd68ca17595c04.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get classes for which the user is a teacher.
func (m *TaughtClassesRequestBuilder) Get(options *TaughtClassesRequestBuilderGetOptions)(*TaughtClassesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTaughtClassesResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*TaughtClassesResponse), nil
}
func (m *TaughtClassesRequestBuilder) Ref()(*ibfebe8e1821c93dda8997c543dd7161494427c674556776df0eee3b6720e6f3c.RefRequestBuilder) {
    return ibfebe8e1821c93dda8997c543dd7161494427c674556776df0eee3b6720e6f3c.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
