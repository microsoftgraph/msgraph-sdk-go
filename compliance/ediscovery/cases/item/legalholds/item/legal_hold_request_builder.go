package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i69aec157e52758a0945b48bcb47b6d4479cfd9fd7daf48bc04b98f6a9cb235bb "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/legalholds/item/usersources"
    icf19f57741dd687a544069f5a98ed14a297becf697826e2b74c31147376e3c55 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/legalholds/item/sitesources"
    i9980d1ce0e24630cdd6399fe1c804147f29ed1ed65e117bd6eefd4fe89e6d162 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/legalholds/item/usersources/item"
    ia0148ae1539b87fb39be32b9c20fe8150aec9f673cfb5b9da8dbc5bfddd7414d "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/legalholds/item/sitesources/item"
)

// Builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\legalHolds\{legalHold-id}
type LegalHoldRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type LegalHoldRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type LegalHoldRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *LegalHoldRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Returns a list of case legalHold objects for this case.  Nullable.
type LegalHoldRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type LegalHoldRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.LegalHold;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new LegalHoldRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewLegalHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LegalHoldRequestBuilder) {
    m := &LegalHoldRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/legalHolds/{legalHold_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new LegalHoldRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewLegalHoldRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LegalHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLegalHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// Returns a list of case legalHold objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *LegalHoldRequestBuilder) CreateDeleteRequestInformation(options *LegalHoldRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Returns a list of case legalHold objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *LegalHoldRequestBuilder) CreateGetRequestInformation(options *LegalHoldRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Returns a list of case legalHold objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *LegalHoldRequestBuilder) CreatePatchRequestInformation(options *LegalHoldRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Returns a list of case legalHold objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *LegalHoldRequestBuilder) Delete(options *LegalHoldRequestBuilderDeleteOptions)(error) {
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
// Returns a list of case legalHold objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *LegalHoldRequestBuilder) Get(options *LegalHoldRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.LegalHold, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewLegalHold() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.LegalHold), nil
}
// Returns a list of case legalHold objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *LegalHoldRequestBuilder) Patch(options *LegalHoldRequestBuilderPatchOptions)(error) {
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
func (m *LegalHoldRequestBuilder) SiteSources()(*icf19f57741dd687a544069f5a98ed14a297becf697826e2b74c31147376e3c55.SiteSourcesRequestBuilder) {
    return icf19f57741dd687a544069f5a98ed14a297becf697826e2b74c31147376e3c55.NewSiteSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.compliance.ediscovery.cases.item.legalHolds.item.siteSources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *LegalHoldRequestBuilder) SiteSourcesById(id string)(*ia0148ae1539b87fb39be32b9c20fe8150aec9f673cfb5b9da8dbc5bfddd7414d.SiteSourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["siteSource_id"] = id
    }
    return ia0148ae1539b87fb39be32b9c20fe8150aec9f673cfb5b9da8dbc5bfddd7414d.NewSiteSourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *LegalHoldRequestBuilder) UserSources()(*i69aec157e52758a0945b48bcb47b6d4479cfd9fd7daf48bc04b98f6a9cb235bb.UserSourcesRequestBuilder) {
    return i69aec157e52758a0945b48bcb47b6d4479cfd9fd7daf48bc04b98f6a9cb235bb.NewUserSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.compliance.ediscovery.cases.item.legalHolds.item.userSources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *LegalHoldRequestBuilder) UserSourcesById(id string)(*i9980d1ce0e24630cdd6399fe1c804147f29ed1ed65e117bd6eefd4fe89e6d162.UserSourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userSource_id"] = id
    }
    return i9980d1ce0e24630cdd6399fe1c804147f29ed1ed65e117bd6eefd4fe89e6d162.NewUserSourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
