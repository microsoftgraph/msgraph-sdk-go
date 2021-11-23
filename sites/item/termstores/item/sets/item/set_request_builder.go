package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i492f2c0deb9e84f7d9a7743fda4e1d1f3a1791c71cda353ebb7023d385689e39 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item/children"
    i62a9e863e325b0aabc2b399fe176787f5e007e5d5ebde360ce5496bb73b60b1f "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item/terms"
    i6bc6f7534ec2cd0ea31ebe97205d014f30992fabe1575742019d5cb73a56227b "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item/parentgroup"
    ifb3c619491821827a3e3ee9e96a8b7fee32dcae5099fa00e741a8272d0353454 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item/relations"
    i2b9fdf09ba8d0e3f8e4ec726b485908e535dd730cd31df3088c246390559cf0c "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item/children/item"
    id29b486b9de74a5e228634cbc0b78a3e93bf36959bcca334777fa52e4dd5ba76 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item/relations/item"
    id6ef0c2679ba27db5107619f8479aebb6116d87d37a91efddc176e4a725bef98 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item/terms/item"
)

// setRequestBuilder builds and executes requests for operations under \sites\{site-id}\termStores\{store-id}\sets\{set-id}
type SetRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SetRequestBuilderDeleteOptions options for Delete
type SetRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SetRequestBuilderGetOptions options for Get
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
// setRequestBuilderGetQueryParameters collection of all sets available in the term store.
type SetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// SetRequestBuilderPatchOptions options for Patch
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
func (m *SetRequestBuilder) Children()(*i492f2c0deb9e84f7d9a7743fda4e1d1f3a1791c71cda353ebb7023d385689e39.ChildrenRequestBuilder) {
    return i492f2c0deb9e84f7d9a7743fda4e1d1f3a1791c71cda353ebb7023d385689e39.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStores.item.sets.item.children.item collection
func (m *SetRequestBuilder) ChildrenById(id string)(*i2b9fdf09ba8d0e3f8e4ec726b485908e535dd730cd31df3088c246390559cf0c.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id"] = id
    }
    return i2b9fdf09ba8d0e3f8e4ec726b485908e535dd730cd31df3088c246390559cf0c.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSetRequestBuilderInternal instantiates a new SetRequestBuilder and sets the default values.
func NewSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetRequestBuilder) {
    m := &SetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/termStores/{store_id}/sets/{set_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSetRequestBuilder instantiates a new SetRequestBuilder and sets the default values.
func NewSetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation collection of all sets available in the term store.
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
// CreateGetRequestInformation collection of all sets available in the term store.
func (m *SetRequestBuilder) CreateGetRequestInformation(options *SetRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation collection of all sets available in the term store.
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
// Delete collection of all sets available in the term store.
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
// Get collection of all sets available in the term store.
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
func (m *SetRequestBuilder) ParentGroup()(*i6bc6f7534ec2cd0ea31ebe97205d014f30992fabe1575742019d5cb73a56227b.ParentGroupRequestBuilder) {
    return i6bc6f7534ec2cd0ea31ebe97205d014f30992fabe1575742019d5cb73a56227b.NewParentGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch collection of all sets available in the term store.
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
func (m *SetRequestBuilder) Relations()(*ifb3c619491821827a3e3ee9e96a8b7fee32dcae5099fa00e741a8272d0353454.RelationsRequestBuilder) {
    return ifb3c619491821827a3e3ee9e96a8b7fee32dcae5099fa00e741a8272d0353454.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStores.item.sets.item.relations.item collection
func (m *SetRequestBuilder) RelationsById(id string)(*id29b486b9de74a5e228634cbc0b78a3e93bf36959bcca334777fa52e4dd5ba76.RelationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation_id"] = id
    }
    return id29b486b9de74a5e228634cbc0b78a3e93bf36959bcca334777fa52e4dd5ba76.NewRelationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SetRequestBuilder) Terms()(*i62a9e863e325b0aabc2b399fe176787f5e007e5d5ebde360ce5496bb73b60b1f.TermsRequestBuilder) {
    return i62a9e863e325b0aabc2b399fe176787f5e007e5d5ebde360ce5496bb73b60b1f.NewTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStores.item.sets.item.terms.item collection
func (m *SetRequestBuilder) TermsById(id string)(*id6ef0c2679ba27db5107619f8479aebb6116d87d37a91efddc176e4a725bef98.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id"] = id
    }
    return id6ef0c2679ba27db5107619f8479aebb6116d87d37a91efddc176e4a725bef98.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
