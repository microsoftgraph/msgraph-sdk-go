package solutions

import (
    i6d5b6940d5e24f8aee3eb7eef5499c5a97c9c71485b644e5b94efadbc56609b0 "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingcurrencies"
    ic2077c84f4cea40593f46a097344f757e71b10d42626a06348d8a6856a4a6408 "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingbusinesses"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i30929f1a7f5c848dd5ee1e9d3e9b1dc2926ba5d92f9f5ca889f41edee949d39a "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingcurrencies/item"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i98532a0d43958ac46a84969d106d7877583a8e83cd026d67fbc826e462b1e8c7 "github.com/microsoftgraph/msgraph-sdk-go/solutions/bookingbusinesses/item"
)

// SolutionsRequestBuilder builds and executes requests for operations under \solutions
type SolutionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SolutionsRequestBuilderGetOptions options for Get
type SolutionsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SolutionsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SolutionsRequestBuilderGetQueryParameters get solutions
type SolutionsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SolutionsRequestBuilderPatchOptions options for Patch
type SolutionsRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Solutions;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *SolutionsRequestBuilder) BookingBusinesses()(*ic2077c84f4cea40593f46a097344f757e71b10d42626a06348d8a6856a4a6408.BookingBusinessesRequestBuilder) {
    return ic2077c84f4cea40593f46a097344f757e71b10d42626a06348d8a6856a4a6408.NewBookingBusinessesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BookingBusinessesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.solutions.bookingBusinesses.item collection
func (m *SolutionsRequestBuilder) BookingBusinessesById(id string)(*i98532a0d43958ac46a84969d106d7877583a8e83cd026d67fbc826e462b1e8c7.BookingBusinessRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingBusiness_id"] = id
    }
    return i98532a0d43958ac46a84969d106d7877583a8e83cd026d67fbc826e462b1e8c7.NewBookingBusinessRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SolutionsRequestBuilder) BookingCurrencies()(*i6d5b6940d5e24f8aee3eb7eef5499c5a97c9c71485b644e5b94efadbc56609b0.BookingCurrenciesRequestBuilder) {
    return i6d5b6940d5e24f8aee3eb7eef5499c5a97c9c71485b644e5b94efadbc56609b0.NewBookingCurrenciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BookingCurrenciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.solutions.bookingCurrencies.item collection
func (m *SolutionsRequestBuilder) BookingCurrenciesById(id string)(*i30929f1a7f5c848dd5ee1e9d3e9b1dc2926ba5d92f9f5ca889f41edee949d39a.BookingCurrencyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingCurrency_id"] = id
    }
    return i30929f1a7f5c848dd5ee1e9d3e9b1dc2926ba5d92f9f5ca889f41edee949d39a.NewBookingCurrencyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSolutionsRequestBuilderInternal instantiates a new SolutionsRequestBuilder and sets the default values.
func NewSolutionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SolutionsRequestBuilder) {
    m := &SolutionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/solutions{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSolutionsRequestBuilder instantiates a new SolutionsRequestBuilder and sets the default values.
func NewSolutionsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SolutionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSolutionsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get solutions
func (m *SolutionsRequestBuilder) CreateGetRequestInformation(options *SolutionsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update solutions
func (m *SolutionsRequestBuilder) CreatePatchRequestInformation(options *SolutionsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
// Get get solutions
func (m *SolutionsRequestBuilder) Get(options *SolutionsRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SolutionsRoot, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewSolutionsRoot() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SolutionsRoot), nil
}
// Patch update solutions
func (m *SolutionsRequestBuilder) Patch(options *SolutionsRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
