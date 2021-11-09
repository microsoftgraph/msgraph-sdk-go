package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i2820c9edb98271b07d8e6c82803b18c5fd396dfdb829a54c4e29f38c27929eaf "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/custodians/item/activate"
    i64bc074d4248009f3f943b7082d0570c5212d84eb48c3bb9747afa7539372b89 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/custodians/item/updateindex"
    i8d2e14d221f199197439ef56b2d86f70e5a82a6b723a34825e5bd776da1d3068 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/custodians/item/sitesources"
    id950a9b18a86bcc619de6aa3f7d8be72ec9923464f61fd8de209347599e114c6 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/custodians/item/unifiedgroupsources"
    idaa9015c41d3bbd1ca2da0007dbdfac1631455b88b293af816aa086696ff2101 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/custodians/item/release"
    iff888f6d890df58fa1b2acb0c7c2f374c96b4aa996afb6086badbf69a2c18e22 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/custodians/item/usersources"
    i3d68438633e939c8a8182b4bd49ace94b8af18198c26539a46c3fb993fce98ce "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/custodians/item/unifiedgroupsources/item"
    i590b8799079072a8392ba8e789481494a803c908c9d0d0aac42bb5038f174e67 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/custodians/item/sitesources/item"
    i7abe6a82dd72a789b196d0e65d1e575134514972a9d9483c1fc378eb8003aed3 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/custodians/item/usersources/item"
)

// Builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\custodians\{custodian-id}
type CustodianRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type CustodianRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type CustodianRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CustodianRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Returns a list of case custodian objects for this case.  Nullable.
type CustodianRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type CustodianRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Custodian;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CustodianRequestBuilder) Activate()(*i2820c9edb98271b07d8e6c82803b18c5fd396dfdb829a54c4e29f38c27929eaf.ActivateRequestBuilder) {
    return i2820c9edb98271b07d8e6c82803b18c5fd396dfdb829a54c4e29f38c27929eaf.NewActivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new CustodianRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCustodianRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustodianRequestBuilder) {
    m := &CustodianRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/custodians/{custodian_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new CustodianRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCustodianRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustodianRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustodianRequestBuilderInternal(urlParams, requestAdapter)
}
// Returns a list of case custodian objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *CustodianRequestBuilder) CreateDeleteRequestInformation(options *CustodianRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Returns a list of case custodian objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *CustodianRequestBuilder) CreateGetRequestInformation(options *CustodianRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Returns a list of case custodian objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *CustodianRequestBuilder) CreatePatchRequestInformation(options *CustodianRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Returns a list of case custodian objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *CustodianRequestBuilder) Delete(options *CustodianRequestBuilderDeleteOptions)(error) {
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
// Returns a list of case custodian objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *CustodianRequestBuilder) Get(options *CustodianRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Custodian, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewCustodian() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Custodian), nil
}
// Returns a list of case custodian objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *CustodianRequestBuilder) Patch(options *CustodianRequestBuilderPatchOptions)(error) {
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
func (m *CustodianRequestBuilder) Release()(*idaa9015c41d3bbd1ca2da0007dbdfac1631455b88b293af816aa086696ff2101.ReleaseRequestBuilder) {
    return idaa9015c41d3bbd1ca2da0007dbdfac1631455b88b293af816aa086696ff2101.NewReleaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CustodianRequestBuilder) SiteSources()(*i8d2e14d221f199197439ef56b2d86f70e5a82a6b723a34825e5bd776da1d3068.SiteSourcesRequestBuilder) {
    return i8d2e14d221f199197439ef56b2d86f70e5a82a6b723a34825e5bd776da1d3068.NewSiteSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.compliance.ediscovery.cases.item.custodians.item.siteSources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CustodianRequestBuilder) SiteSourcesById(id string)(*i590b8799079072a8392ba8e789481494a803c908c9d0d0aac42bb5038f174e67.SiteSourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["siteSource_id"] = id
    }
    return i590b8799079072a8392ba8e789481494a803c908c9d0d0aac42bb5038f174e67.NewSiteSourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CustodianRequestBuilder) UnifiedGroupSources()(*id950a9b18a86bcc619de6aa3f7d8be72ec9923464f61fd8de209347599e114c6.UnifiedGroupSourcesRequestBuilder) {
    return id950a9b18a86bcc619de6aa3f7d8be72ec9923464f61fd8de209347599e114c6.NewUnifiedGroupSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.compliance.ediscovery.cases.item.custodians.item.unifiedGroupSources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CustodianRequestBuilder) UnifiedGroupSourcesById(id string)(*i3d68438633e939c8a8182b4bd49ace94b8af18198c26539a46c3fb993fce98ce.UnifiedGroupSourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedGroupSource_id"] = id
    }
    return i3d68438633e939c8a8182b4bd49ace94b8af18198c26539a46c3fb993fce98ce.NewUnifiedGroupSourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CustodianRequestBuilder) UpdateIndex()(*i64bc074d4248009f3f943b7082d0570c5212d84eb48c3bb9747afa7539372b89.UpdateIndexRequestBuilder) {
    return i64bc074d4248009f3f943b7082d0570c5212d84eb48c3bb9747afa7539372b89.NewUpdateIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CustodianRequestBuilder) UserSources()(*iff888f6d890df58fa1b2acb0c7c2f374c96b4aa996afb6086badbf69a2c18e22.UserSourcesRequestBuilder) {
    return iff888f6d890df58fa1b2acb0c7c2f374c96b4aa996afb6086badbf69a2c18e22.NewUserSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.compliance.ediscovery.cases.item.custodians.item.userSources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CustodianRequestBuilder) UserSourcesById(id string)(*i7abe6a82dd72a789b196d0e65d1e575134514972a9d9483c1fc378eb8003aed3.UserSourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userSource_id"] = id
    }
    return i7abe6a82dd72a789b196d0e65d1e575134514972a9d9483c1fc378eb8003aed3.NewUserSourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
