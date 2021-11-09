package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i216680e4546e630eb3cbf2fcf969b793e1a5b3d1fdc99e7019519e14005c6ff2 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/reviewsets/item/export"
    i48a563666fce9357ae7a387c08104ebea6c2da139b47ccdb29c7569a0935242d "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/reviewsets/item/addtoreviewset"
    i7209cca27bc5c0646c6dee17a9f11942dc01a9c9631005ca9c05b7ed53b0dd1b "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/reviewsets/item/queries"
    i347f500aa366008d63004070310f906e17971d948a5d6891fa17cfb9a4978776 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/reviewsets/item/queries/item"
)

// Builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\reviewSets\{reviewSet-id}
type ReviewSetRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ReviewSetRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ReviewSetRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ReviewSetRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Returns a list of reviewSet objects in the case. Read-only. Nullable.
type ReviewSetRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ReviewSetRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ReviewSet;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ReviewSetRequestBuilder) AddToReviewSet()(*i48a563666fce9357ae7a387c08104ebea6c2da139b47ccdb29c7569a0935242d.AddToReviewSetRequestBuilder) {
    return i48a563666fce9357ae7a387c08104ebea6c2da139b47ccdb29c7569a0935242d.NewAddToReviewSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ReviewSetRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewReviewSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ReviewSetRequestBuilder) {
    m := &ReviewSetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/reviewSets/{reviewSet_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ReviewSetRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewReviewSetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ReviewSetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReviewSetRequestBuilderInternal(urlParams, requestAdapter)
}
// Returns a list of reviewSet objects in the case. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ReviewSetRequestBuilder) CreateDeleteRequestInformation(options *ReviewSetRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// Returns a list of reviewSet objects in the case. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ReviewSetRequestBuilder) CreateGetRequestInformation(options *ReviewSetRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Returns a list of reviewSet objects in the case. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ReviewSetRequestBuilder) CreatePatchRequestInformation(options *ReviewSetRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Returns a list of reviewSet objects in the case. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ReviewSetRequestBuilder) Delete(options *ReviewSetRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ReviewSetRequestBuilder) Export()(*i216680e4546e630eb3cbf2fcf969b793e1a5b3d1fdc99e7019519e14005c6ff2.ExportRequestBuilder) {
    return i216680e4546e630eb3cbf2fcf969b793e1a5b3d1fdc99e7019519e14005c6ff2.NewExportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Returns a list of reviewSet objects in the case. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ReviewSetRequestBuilder) Get(options *ReviewSetRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ReviewSet, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewReviewSet() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ReviewSet), nil
}
// Returns a list of reviewSet objects in the case. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ReviewSetRequestBuilder) Patch(options *ReviewSetRequestBuilderPatchOptions)(error) {
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
func (m *ReviewSetRequestBuilder) Queries()(*i7209cca27bc5c0646c6dee17a9f11942dc01a9c9631005ca9c05b7ed53b0dd1b.QueriesRequestBuilder) {
    return i7209cca27bc5c0646c6dee17a9f11942dc01a9c9631005ca9c05b7ed53b0dd1b.NewQueriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.compliance.ediscovery.cases.item.reviewSets.item.queries.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ReviewSetRequestBuilder) QueriesById(id string)(*i347f500aa366008d63004070310f906e17971d948a5d6891fa17cfb9a4978776.ReviewSetQueryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["reviewSetQuery_id"] = id
    }
    return i347f500aa366008d63004070310f906e17971d948a5d6891fa17cfb9a4978776.NewReviewSetQueryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
