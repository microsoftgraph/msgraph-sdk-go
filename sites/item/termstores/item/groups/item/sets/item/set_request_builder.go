package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i28411d489ad3cd729f1892303dcb2505651bcd5f6acc3e5cf227bc4118eea651 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/relations"
    i571a44d28bde16f052863de7c5cf0d88446d9f63065509564f1a776efa65f403 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/terms"
    i827e50e0e704e89f39c6ad2c9dd28f427f51f016d3cd9a70c29413bf073622a7 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/children"
    ida7529d11c61cb15d10ee4276f7eef55f69d00bec0595927684536c75e82d402 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/parentgroup"
    i193200b3734d30a01f56084d8677802018b1f9f0bbb205294e6e271745ac81a8 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/terms/item"
    ia984aaf284a0f001442f674627bbd07e2de13686abd06c482db59f9e652b98f2 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/children/item"
    ibb602acd9415bff5d888f98c2aba16005dd06d930f81c6a0360e1614b8bc2aaf "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/relations/item"
)

// setRequestBuilder builds and executes requests for operations under \sites\{site-id}\termStores\{store-id}\groups\{group-id}\sets\{set-id}
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
// setRequestBuilderGetQueryParameters all sets under the group in a term [store].
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
func (m *SetRequestBuilder) Children()(*i827e50e0e704e89f39c6ad2c9dd28f427f51f016d3cd9a70c29413bf073622a7.ChildrenRequestBuilder) {
    return i827e50e0e704e89f39c6ad2c9dd28f427f51f016d3cd9a70c29413bf073622a7.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStores.item.groups.item.sets.item.children.item collection
func (m *SetRequestBuilder) ChildrenById(id string)(*ia984aaf284a0f001442f674627bbd07e2de13686abd06c482db59f9e652b98f2.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id"] = id
    }
    return ia984aaf284a0f001442f674627bbd07e2de13686abd06c482db59f9e652b98f2.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSetRequestBuilderInternal instantiates a new SetRequestBuilder and sets the default values.
func NewSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetRequestBuilder) {
    m := &SetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/termStores/{store_id}/groups/{group_id}/sets/{set_id}{?select,expand}";
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
// CreateDeleteRequestInformation all sets under the group in a term [store].
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
// CreateGetRequestInformation all sets under the group in a term [store].
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
// CreatePatchRequestInformation all sets under the group in a term [store].
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
// Delete all sets under the group in a term [store].
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
// Get all sets under the group in a term [store].
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
func (m *SetRequestBuilder) ParentGroup()(*ida7529d11c61cb15d10ee4276f7eef55f69d00bec0595927684536c75e82d402.ParentGroupRequestBuilder) {
    return ida7529d11c61cb15d10ee4276f7eef55f69d00bec0595927684536c75e82d402.NewParentGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch all sets under the group in a term [store].
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
func (m *SetRequestBuilder) Relations()(*i28411d489ad3cd729f1892303dcb2505651bcd5f6acc3e5cf227bc4118eea651.RelationsRequestBuilder) {
    return i28411d489ad3cd729f1892303dcb2505651bcd5f6acc3e5cf227bc4118eea651.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStores.item.groups.item.sets.item.relations.item collection
func (m *SetRequestBuilder) RelationsById(id string)(*ibb602acd9415bff5d888f98c2aba16005dd06d930f81c6a0360e1614b8bc2aaf.RelationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation_id"] = id
    }
    return ibb602acd9415bff5d888f98c2aba16005dd06d930f81c6a0360e1614b8bc2aaf.NewRelationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SetRequestBuilder) Terms()(*i571a44d28bde16f052863de7c5cf0d88446d9f63065509564f1a776efa65f403.TermsRequestBuilder) {
    return i571a44d28bde16f052863de7c5cf0d88446d9f63065509564f1a776efa65f403.NewTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStores.item.groups.item.sets.item.terms.item collection
func (m *SetRequestBuilder) TermsById(id string)(*i193200b3734d30a01f56084d8677802018b1f9f0bbb205294e6e271745ac81a8.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id"] = id
    }
    return i193200b3734d30a01f56084d8677802018b1f9f0bbb205294e6e271745ac81a8.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
