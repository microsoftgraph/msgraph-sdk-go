package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i077b9469338cd6fe76c9dfeda6243e6b5d520b8243a664903766f7f5af28b90b "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/noncustodialdatasources"
    i096bda9ef0ded6406764a7b5528f793d5934749454ff835b4a6273d0070c920e "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/tags"
    i1758a6fc539f48df5f433ed46227df12bcd9a6541a52ef55a75f4bc807d811dd "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/close"
    i3217f9d981f136fb37dfc5f7ab7227bd1458edbd9470d914f230f842675a5eda "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/operations"
    i57c09eece2e1de02831fff718f46ace1616f4ea94ad1d234c59bcfadfb1afab6 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/reopen"
    i91be10d76ad65752f5a439b6614e10588f62477a29742473de51fe9594cb8ad3 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/legalholds"
    ib48b5bf0deeb8c5ce0bef85f4a7be5abf1af6eb678b0b896e7ed38c0d005135e "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/reviewsets"
    id798b04ba84cb0cb17a30b9cd4e27f23d7cfd2716289b8db290b971dcf198c87 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/sourcecollections"
    ie3fac1241b58c1888b0acfd23e41de8a9c6b6e4b6dbaf26591077227f0c8bf69 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/custodians"
    if4e4bf311140c1cc077500acee99dedd4f5552fe94bfc81fe710cab65bfbaa33 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/settings"
    i25a74a234f89f3c7d72202262fe2486f396192162bfd2ddaaa38419906519095 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/reviewsets/item"
    i446437d64b3005964e44a434329d694612081d1ce27e8660be0424dbabd26dee "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/tags/item"
    i5678950134df902871479649dcc98de02d8e2edbecbc25c3b5a5907ca1b9115c "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/legalholds/item"
    i708067d6776642a4aacabd9ea917f6b8275a1f15cc6c35141c931aff1bb42156 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/custodians/item"
    id76a53de1db93f10adcd7de19ca0e597fc304ca23f5ac8d9324d16ffb3e4b578 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/noncustodialdatasources/item"
    iea59213ebdd82950a7fb82ce688e93cf81050dcc2c557572cf5940d272df5ed0 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/operations/item"
    if85e073d14e518a57564e86e007651fd9029f94947d243843900ea08d7231ec7 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item"
)

// Builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}
type CaseRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type CaseRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type CaseRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CaseRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get cases from compliance
type CaseRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type CaseRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Case_escaped;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CaseRequestBuilder) Close()(*i1758a6fc539f48df5f433ed46227df12bcd9a6541a52ef55a75f4bc807d811dd.CloseRequestBuilder) {
    return i1758a6fc539f48df5f433ed46227df12bcd9a6541a52ef55a75f4bc807d811dd.NewCloseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new CaseRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CaseRequestBuilder) {
    m := &CaseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new CaseRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCaseRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CaseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCaseRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete navigation property cases for compliance
// Parameters:
//  - options : Options for the request
func (m *CaseRequestBuilder) CreateDeleteRequestInformation(options *CaseRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get cases from compliance
// Parameters:
//  - options : Options for the request
func (m *CaseRequestBuilder) CreateGetRequestInformation(options *CaseRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Update the navigation property cases in compliance
// Parameters:
//  - options : Options for the request
func (m *CaseRequestBuilder) CreatePatchRequestInformation(options *CaseRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CaseRequestBuilder) Custodians()(*ie3fac1241b58c1888b0acfd23e41de8a9c6b6e4b6dbaf26591077227f0c8bf69.CustodiansRequestBuilder) {
    return ie3fac1241b58c1888b0acfd23e41de8a9c6b6e4b6dbaf26591077227f0c8bf69.NewCustodiansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.compliance.ediscovery.cases.item.custodians.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CaseRequestBuilder) CustodiansById(id string)(*i708067d6776642a4aacabd9ea917f6b8275a1f15cc6c35141c931aff1bb42156.CustodianRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["custodian_id"] = id
    }
    return i708067d6776642a4aacabd9ea917f6b8275a1f15cc6c35141c931aff1bb42156.NewCustodianRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete navigation property cases for compliance
// Parameters:
//  - options : Options for the request
func (m *CaseRequestBuilder) Delete(options *CaseRequestBuilderDeleteOptions)(error) {
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
// Get cases from compliance
// Parameters:
//  - options : Options for the request
func (m *CaseRequestBuilder) Get(options *CaseRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Case_escaped, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewCase_escaped() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Case_escaped), nil
}
func (m *CaseRequestBuilder) LegalHolds()(*i91be10d76ad65752f5a439b6614e10588f62477a29742473de51fe9594cb8ad3.LegalHoldsRequestBuilder) {
    return i91be10d76ad65752f5a439b6614e10588f62477a29742473de51fe9594cb8ad3.NewLegalHoldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.compliance.ediscovery.cases.item.legalHolds.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CaseRequestBuilder) LegalHoldsById(id string)(*i5678950134df902871479649dcc98de02d8e2edbecbc25c3b5a5907ca1b9115c.LegalHoldRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["legalHold_id"] = id
    }
    return i5678950134df902871479649dcc98de02d8e2edbecbc25c3b5a5907ca1b9115c.NewLegalHoldRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CaseRequestBuilder) NoncustodialDataSources()(*i077b9469338cd6fe76c9dfeda6243e6b5d520b8243a664903766f7f5af28b90b.NoncustodialDataSourcesRequestBuilder) {
    return i077b9469338cd6fe76c9dfeda6243e6b5d520b8243a664903766f7f5af28b90b.NewNoncustodialDataSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.compliance.ediscovery.cases.item.noncustodialDataSources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CaseRequestBuilder) NoncustodialDataSourcesById(id string)(*id76a53de1db93f10adcd7de19ca0e597fc304ca23f5ac8d9324d16ffb3e4b578.NoncustodialDataSourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["noncustodialDataSource_id"] = id
    }
    return id76a53de1db93f10adcd7de19ca0e597fc304ca23f5ac8d9324d16ffb3e4b578.NewNoncustodialDataSourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CaseRequestBuilder) Operations()(*i3217f9d981f136fb37dfc5f7ab7227bd1458edbd9470d914f230f842675a5eda.OperationsRequestBuilder) {
    return i3217f9d981f136fb37dfc5f7ab7227bd1458edbd9470d914f230f842675a5eda.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.compliance.ediscovery.cases.item.operations.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CaseRequestBuilder) OperationsById(id string)(*iea59213ebdd82950a7fb82ce688e93cf81050dcc2c557572cf5940d272df5ed0.CaseOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["caseOperation_id"] = id
    }
    return iea59213ebdd82950a7fb82ce688e93cf81050dcc2c557572cf5940d272df5ed0.NewCaseOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Update the navigation property cases in compliance
// Parameters:
//  - options : Options for the request
func (m *CaseRequestBuilder) Patch(options *CaseRequestBuilderPatchOptions)(error) {
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
func (m *CaseRequestBuilder) Reopen()(*i57c09eece2e1de02831fff718f46ace1616f4ea94ad1d234c59bcfadfb1afab6.ReopenRequestBuilder) {
    return i57c09eece2e1de02831fff718f46ace1616f4ea94ad1d234c59bcfadfb1afab6.NewReopenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CaseRequestBuilder) ReviewSets()(*ib48b5bf0deeb8c5ce0bef85f4a7be5abf1af6eb678b0b896e7ed38c0d005135e.ReviewSetsRequestBuilder) {
    return ib48b5bf0deeb8c5ce0bef85f4a7be5abf1af6eb678b0b896e7ed38c0d005135e.NewReviewSetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.compliance.ediscovery.cases.item.reviewSets.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CaseRequestBuilder) ReviewSetsById(id string)(*i25a74a234f89f3c7d72202262fe2486f396192162bfd2ddaaa38419906519095.ReviewSetRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["reviewSet_id"] = id
    }
    return i25a74a234f89f3c7d72202262fe2486f396192162bfd2ddaaa38419906519095.NewReviewSetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CaseRequestBuilder) Settings()(*if4e4bf311140c1cc077500acee99dedd4f5552fe94bfc81fe710cab65bfbaa33.SettingsRequestBuilder) {
    return if4e4bf311140c1cc077500acee99dedd4f5552fe94bfc81fe710cab65bfbaa33.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CaseRequestBuilder) SourceCollections()(*id798b04ba84cb0cb17a30b9cd4e27f23d7cfd2716289b8db290b971dcf198c87.SourceCollectionsRequestBuilder) {
    return id798b04ba84cb0cb17a30b9cd4e27f23d7cfd2716289b8db290b971dcf198c87.NewSourceCollectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.compliance.ediscovery.cases.item.sourceCollections.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CaseRequestBuilder) SourceCollectionsById(id string)(*if85e073d14e518a57564e86e007651fd9029f94947d243843900ea08d7231ec7.SourceCollectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sourceCollection_id"] = id
    }
    return if85e073d14e518a57564e86e007651fd9029f94947d243843900ea08d7231ec7.NewSourceCollectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CaseRequestBuilder) Tags()(*i096bda9ef0ded6406764a7b5528f793d5934749454ff835b4a6273d0070c920e.TagsRequestBuilder) {
    return i096bda9ef0ded6406764a7b5528f793d5934749454ff835b4a6273d0070c920e.NewTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.compliance.ediscovery.cases.item.tags.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CaseRequestBuilder) TagsById(id string)(*i446437d64b3005964e44a434329d694612081d1ce27e8660be0424dbabd26dee.TagRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tag_id"] = id
    }
    return i446437d64b3005964e44a434329d694612081d1ce27e8660be0424dbabd26dee.NewTagRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
