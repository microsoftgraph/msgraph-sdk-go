package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    ib50fe719b41297b226c038c3403f9fc36fc99fbb9cd41b4ffb1c75c3c3e244e0 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/parentgroup"
    ib67061c9a5ec7f54c1fceca2aa7982053498cb3be7809cad62a033dc35b973c3 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/terms"
    ieec65ad1ca1d01c340556d91630b2f2e19f4283f01b845b7eb8f2b149f1239d3 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/children"
    if8ad04d6afd501e6c191842628bbb732b74c6e6135337da4b3a96a9ab4ced94a "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/relations"
    i285174a663ce0c2b2b0057b47643838a60db3b04b36225f624e5a9063261d49d "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/terms/item"
    i6add9f2ac9e1c03c41f3593c206bb30367db44710541f6111e7233fbe3710d7a "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/children/item"
    i6bbacacf32786b9a7dc430d497d260fe18ca363ecc82921008d5fdc715d0d70d "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/relations/item"
)

// Builds and executes requests for operations under \sites\{site-id}\termStore\sets\{set-id}
type SetRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type SetRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type SetRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SetRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Collection of all sets available in the term store.
type SetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type SetRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Set;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *SetRequestBuilder) Children()(*ieec65ad1ca1d01c340556d91630b2f2e19f4283f01b845b7eb8f2b149f1239d3.ChildrenRequestBuilder) {
    return ieec65ad1ca1d01c340556d91630b2f2e19f4283f01b845b7eb8f2b149f1239d3.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStore.sets.item.children.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SetRequestBuilder) ChildrenById(id string)(*i6add9f2ac9e1c03c41f3593c206bb30367db44710541f6111e7233fbe3710d7a.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id"] = id
    }
    return i6add9f2ac9e1c03c41f3593c206bb30367db44710541f6111e7233fbe3710d7a.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new SetRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetRequestBuilder) {
    m := &SetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/termStore/sets/{set_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new SetRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetRequestBuilderInternal(urlParams, requestAdapter)
}
// Collection of all sets available in the term store.
// Parameters:
//  - options : Options for the request
func (m *SetRequestBuilder) CreateDeleteRequestInformation(options *SetRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Collection of all sets available in the term store.
// Parameters:
//  - options : Options for the request
func (m *SetRequestBuilder) CreateGetRequestInformation(options *SetRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Collection of all sets available in the term store.
// Parameters:
//  - options : Options for the request
func (m *SetRequestBuilder) CreatePatchRequestInformation(options *SetRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Collection of all sets available in the term store.
// Parameters:
//  - options : Options for the request
func (m *SetRequestBuilder) Delete(options *SetRequestBuilderDeleteOptions)(error) {
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
// Collection of all sets available in the term store.
// Parameters:
//  - options : Options for the request
func (m *SetRequestBuilder) Get(options *SetRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Set, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewSet() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Set), nil
}
func (m *SetRequestBuilder) ParentGroup()(*ib50fe719b41297b226c038c3403f9fc36fc99fbb9cd41b4ffb1c75c3c3e244e0.ParentGroupRequestBuilder) {
    return ib50fe719b41297b226c038c3403f9fc36fc99fbb9cd41b4ffb1c75c3c3e244e0.NewParentGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Collection of all sets available in the term store.
// Parameters:
//  - options : Options for the request
func (m *SetRequestBuilder) Patch(options *SetRequestBuilderPatchOptions)(error) {
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
func (m *SetRequestBuilder) Relations()(*if8ad04d6afd501e6c191842628bbb732b74c6e6135337da4b3a96a9ab4ced94a.RelationsRequestBuilder) {
    return if8ad04d6afd501e6c191842628bbb732b74c6e6135337da4b3a96a9ab4ced94a.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStore.sets.item.relations.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SetRequestBuilder) RelationsById(id string)(*i6bbacacf32786b9a7dc430d497d260fe18ca363ecc82921008d5fdc715d0d70d.RelationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation_id"] = id
    }
    return i6bbacacf32786b9a7dc430d497d260fe18ca363ecc82921008d5fdc715d0d70d.NewRelationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SetRequestBuilder) Terms()(*ib67061c9a5ec7f54c1fceca2aa7982053498cb3be7809cad62a033dc35b973c3.TermsRequestBuilder) {
    return ib67061c9a5ec7f54c1fceca2aa7982053498cb3be7809cad62a033dc35b973c3.NewTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStore.sets.item.terms.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SetRequestBuilder) TermsById(id string)(*i285174a663ce0c2b2b0057b47643838a60db3b04b36225f624e5a9063261d49d.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id"] = id
    }
    return i285174a663ce0c2b2b0057b47643838a60db3b04b36225f624e5a9063261d49d.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
