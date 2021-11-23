package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i4b46996c82e46abbb82f3e42feb2374dc2329d86c1ffc76177139a4283da6ff7 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/children/item/set"
    i7a080172d6e32d5c7712419d63f89750cab2263f64dd4fbebcd951d38418b7cd "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/children/item/children"
    ibe800dbf2e735c822f05a7aa4f19e58fb17559233abbcf570b7add4f52b54b40 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/children/item/relations"
    i8a752431742e1ac1a636d4e59cd329f70ab3dab6b1779b4e0d42b805fa101e93 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/children/item/relations/item"
    ia766aa6a5716710aa5da06f16059a67780af116d7ac1a24e6a2939b9772bb4df "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/children/item/children/item"
)

// TermRequestBuilder builds and executes requests for operations under \sites\{site-id}\termStore\sets\{set-id}\children\{term-id}
type TermRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TermRequestBuilderDeleteOptions options for Delete
type TermRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TermRequestBuilderGetOptions options for Get
type TermRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TermRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TermRequestBuilderGetQueryParameters children terms of set in term [store].
type TermRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// TermRequestBuilderPatchOptions options for Patch
type TermRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Term;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TermRequestBuilder) Children()(*i7a080172d6e32d5c7712419d63f89750cab2263f64dd4fbebcd951d38418b7cd.ChildrenRequestBuilder) {
    return i7a080172d6e32d5c7712419d63f89750cab2263f64dd4fbebcd951d38418b7cd.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStore.sets.item.children.item.children.item collection
func (m *TermRequestBuilder) ChildrenById(id string)(*ia766aa6a5716710aa5da06f16059a67780af116d7ac1a24e6a2939b9772bb4df.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id1"] = id
    }
    return ia766aa6a5716710aa5da06f16059a67780af116d7ac1a24e6a2939b9772bb4df.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTermRequestBuilderInternal instantiates a new TermRequestBuilder and sets the default values.
func NewTermRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermRequestBuilder) {
    m := &TermRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/termStore/sets/{set_id}/children/{term_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTermRequestBuilder instantiates a new TermRequestBuilder and sets the default values.
func NewTermRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation children terms of set in term [store].
func (m *TermRequestBuilder) CreateDeleteRequestInformation(options *TermRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation children terms of set in term [store].
func (m *TermRequestBuilder) CreateGetRequestInformation(options *TermRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation children terms of set in term [store].
func (m *TermRequestBuilder) CreatePatchRequestInformation(options *TermRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete children terms of set in term [store].
func (m *TermRequestBuilder) Delete(options *TermRequestBuilderDeleteOptions)(error) {
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
// Get children terms of set in term [store].
func (m *TermRequestBuilder) Get(options *TermRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Term, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewTerm() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Term), nil
}
// Patch children terms of set in term [store].
func (m *TermRequestBuilder) Patch(options *TermRequestBuilderPatchOptions)(error) {
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
func (m *TermRequestBuilder) Relations()(*ibe800dbf2e735c822f05a7aa4f19e58fb17559233abbcf570b7add4f52b54b40.RelationsRequestBuilder) {
    return ibe800dbf2e735c822f05a7aa4f19e58fb17559233abbcf570b7add4f52b54b40.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStore.sets.item.children.item.relations.item collection
func (m *TermRequestBuilder) RelationsById(id string)(*i8a752431742e1ac1a636d4e59cd329f70ab3dab6b1779b4e0d42b805fa101e93.RelationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation_id"] = id
    }
    return i8a752431742e1ac1a636d4e59cd329f70ab3dab6b1779b4e0d42b805fa101e93.NewRelationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TermRequestBuilder) Set()(*i4b46996c82e46abbb82f3e42feb2374dc2329d86c1ffc76177139a4283da6ff7.SetRequestBuilder) {
    return i4b46996c82e46abbb82f3e42feb2374dc2329d86c1ffc76177139a4283da6ff7.NewSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
