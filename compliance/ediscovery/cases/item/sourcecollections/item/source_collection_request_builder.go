package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0426a45634f5b14e742c49498f0b0ec5b33febf5abe60bbdf4439a02f523b4c1 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/additionalsources"
    i397384c3371da59e35781624fb28c664cf8f9285cffd37e0b3aedd31beeeb0d5 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/noncustodialsources"
    i5373a604ce27cdb2e4458d435a42a71b181cef3bc4c9a4d5acfb05fd0bbbcb89 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/custodiansources"
    i8ee97cc2b03e3a21e13b440800b4b970ba62e8b85d3b7c815350bc3cfb3cfb21 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/addtoreviewsetoperation"
    ic9dd3e3c2d60f618ce4ad0ae74f4648d5f556a37e3e370cd2e9c9cab844b73ad "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/estimatestatistics"
    ife8c15f320d3cbd567a00fa1d6da5a77600d35f2836385e56173b84c2f7377bc "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/lastestimatestatisticsoperation"
    i463d3f6bee9db9c690fce70fa3d76bdd2833e3a55a00725883f6b5f3a77fbd23 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/additionalsources/item"
)

// Builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\sourceCollections\{sourceCollection-id}
type SourceCollectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type SourceCollectionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type SourceCollectionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SourceCollectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Returns a list of sourceCollection objects associated with this case.
type SourceCollectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type SourceCollectionRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SourceCollection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *SourceCollectionRequestBuilder) AdditionalSources()(*i0426a45634f5b14e742c49498f0b0ec5b33febf5abe60bbdf4439a02f523b4c1.AdditionalSourcesRequestBuilder) {
    return i0426a45634f5b14e742c49498f0b0ec5b33febf5abe60bbdf4439a02f523b4c1.NewAdditionalSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.compliance.ediscovery.cases.item.sourceCollections.item.additionalSources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SourceCollectionRequestBuilder) AdditionalSourcesById(id string)(*i463d3f6bee9db9c690fce70fa3d76bdd2833e3a55a00725883f6b5f3a77fbd23.DataSourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataSource_id"] = id
    }
    return i463d3f6bee9db9c690fce70fa3d76bdd2833e3a55a00725883f6b5f3a77fbd23.NewDataSourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SourceCollectionRequestBuilder) AddToReviewSetOperation()(*i8ee97cc2b03e3a21e13b440800b4b970ba62e8b85d3b7c815350bc3cfb3cfb21.AddToReviewSetOperationRequestBuilder) {
    return i8ee97cc2b03e3a21e13b440800b4b970ba62e8b85d3b7c815350bc3cfb3cfb21.NewAddToReviewSetOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new SourceCollectionRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSourceCollectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SourceCollectionRequestBuilder) {
    m := &SourceCollectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/sourceCollections/{sourceCollection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new SourceCollectionRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSourceCollectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SourceCollectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSourceCollectionRequestBuilderInternal(urlParams, requestAdapter)
}
// Returns a list of sourceCollection objects associated with this case.
// Parameters:
//  - options : Options for the request
func (m *SourceCollectionRequestBuilder) CreateDeleteRequestInformation(options *SourceCollectionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Returns a list of sourceCollection objects associated with this case.
// Parameters:
//  - options : Options for the request
func (m *SourceCollectionRequestBuilder) CreateGetRequestInformation(options *SourceCollectionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Returns a list of sourceCollection objects associated with this case.
// Parameters:
//  - options : Options for the request
func (m *SourceCollectionRequestBuilder) CreatePatchRequestInformation(options *SourceCollectionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SourceCollectionRequestBuilder) CustodianSources()(*i5373a604ce27cdb2e4458d435a42a71b181cef3bc4c9a4d5acfb05fd0bbbcb89.CustodianSourcesRequestBuilder) {
    return i5373a604ce27cdb2e4458d435a42a71b181cef3bc4c9a4d5acfb05fd0bbbcb89.NewCustodianSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Returns a list of sourceCollection objects associated with this case.
// Parameters:
//  - options : Options for the request
func (m *SourceCollectionRequestBuilder) Delete(options *SourceCollectionRequestBuilderDeleteOptions)(error) {
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
func (m *SourceCollectionRequestBuilder) EstimateStatistics()(*ic9dd3e3c2d60f618ce4ad0ae74f4648d5f556a37e3e370cd2e9c9cab844b73ad.EstimateStatisticsRequestBuilder) {
    return ic9dd3e3c2d60f618ce4ad0ae74f4648d5f556a37e3e370cd2e9c9cab844b73ad.NewEstimateStatisticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Returns a list of sourceCollection objects associated with this case.
// Parameters:
//  - options : Options for the request
func (m *SourceCollectionRequestBuilder) Get(options *SourceCollectionRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SourceCollection, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewSourceCollection() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SourceCollection), nil
}
func (m *SourceCollectionRequestBuilder) LastEstimateStatisticsOperation()(*ife8c15f320d3cbd567a00fa1d6da5a77600d35f2836385e56173b84c2f7377bc.LastEstimateStatisticsOperationRequestBuilder) {
    return ife8c15f320d3cbd567a00fa1d6da5a77600d35f2836385e56173b84c2f7377bc.NewLastEstimateStatisticsOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SourceCollectionRequestBuilder) NoncustodialSources()(*i397384c3371da59e35781624fb28c664cf8f9285cffd37e0b3aedd31beeeb0d5.NoncustodialSourcesRequestBuilder) {
    return i397384c3371da59e35781624fb28c664cf8f9285cffd37e0b3aedd31beeeb0d5.NewNoncustodialSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Returns a list of sourceCollection objects associated with this case.
// Parameters:
//  - options : Options for the request
func (m *SourceCollectionRequestBuilder) Patch(options *SourceCollectionRequestBuilderPatchOptions)(error) {
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
