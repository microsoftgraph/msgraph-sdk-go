package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i84942b2bfcb9178388a82b249df98f4e6789b1a2d3acfd6f00794f6a5f0d2672 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/terms/item/relations"
    ic1ef48fed9f7e7747ae011a011d6b502d0acf9bb8670138bc54a9f2c7a23f67d "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/terms/item/children"
    ifc76778cb35efc6ff7ef94cc2bfe606c429a742d6f55ef9c88c613c9adbcb49c "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/terms/item/set"
    i4895978c3d10923dc950fa7a114a6291c0a1894d47982c8b05d7d6af604e2357 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/terms/item/relations/item"
    i8cc4ef0bae4a8f952ee86f40d69f634609d4c6887d2f5a4f08a05fe6b47cd99e "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/terms/item/children/item"
)

// Builds and executes requests for operations under \sites\{site-id}\termStores\{store-id}\groups\{group-id}\sets\{set-id}\terms\{term-id}
type TermRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// All the terms under the set.
type TermRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *TermRequestBuilder) Children()(*ic1ef48fed9f7e7747ae011a011d6b502d0acf9bb8670138bc54a9f2c7a23f67d.ChildrenRequestBuilder) {
    return ic1ef48fed9f7e7747ae011a011d6b502d0acf9bb8670138bc54a9f2c7a23f67d.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStores.item.groups.item.sets.item.terms.item.children.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *TermRequestBuilder) ChildrenById(id string)(*i8cc4ef0bae4a8f952ee86f40d69f634609d4c6887d2f5a4f08a05fe6b47cd99e.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id1"] = id
    }
    return i8cc4ef0bae4a8f952ee86f40d69f634609d4c6887d2f5a4f08a05fe6b47cd99e.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new TermRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTermRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermRequestBuilder) {
    m := &TermRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/sites/{site_id}/termStores/{store_id}/groups/{group_id}/sets/{set_id}/terms/{term_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new TermRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTermRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermRequestBuilderInternal(urlParams, requestAdapter)
}
// All the terms under the set.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *TermRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// All the terms under the set.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *TermRequestBuilder) CreateGetRequestInformation(q func (value *TermRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(TermRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// All the terms under the set.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *TermRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Term, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// All the terms under the set.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *TermRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
// All the terms under the set.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *TermRequestBuilder) Get(q func (value *TermRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Term, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewTerm() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Term), nil
}
// All the terms under the set.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *TermRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Term, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *TermRequestBuilder) Relations()(*i84942b2bfcb9178388a82b249df98f4e6789b1a2d3acfd6f00794f6a5f0d2672.RelationsRequestBuilder) {
    return i84942b2bfcb9178388a82b249df98f4e6789b1a2d3acfd6f00794f6a5f0d2672.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStores.item.groups.item.sets.item.terms.item.relations.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *TermRequestBuilder) RelationsById(id string)(*i4895978c3d10923dc950fa7a114a6291c0a1894d47982c8b05d7d6af604e2357.RelationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation_id"] = id
    }
    return i4895978c3d10923dc950fa7a114a6291c0a1894d47982c8b05d7d6af604e2357.NewRelationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TermRequestBuilder) Set()(*ifc76778cb35efc6ff7ef94cc2bfe606c429a742d6f55ef9c88c613c9adbcb49c.SetRequestBuilder) {
    return ifc76778cb35efc6ff7ef94cc2bfe606c429a742d6f55ef9c88c613c9adbcb49c.NewSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
