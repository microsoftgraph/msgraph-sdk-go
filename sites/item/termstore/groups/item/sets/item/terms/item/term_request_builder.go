package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    ib387154f1adfee1012fa83037220b7890e544dbcaebdc83227c131fda461d9cc "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/groups/item/sets/item/terms/item/children"
    ibd4072fbdcfe2f687bf9eb50c1c9cab17ba7feb6b9e477bc6903701a72813cdc "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/groups/item/sets/item/terms/item/set"
    if2f81ca9fd32dce408153070ba100dec3bbba3f4a78999c825bbc872e5204c37 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/groups/item/sets/item/terms/item/relations"
    i830382e02f940987ffd4007dfcdd55cfd855f15c1563327664bc6b5c21ac048d "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/groups/item/sets/item/terms/item/children/item"
    ia02cb2beba7e2dc73eabfbe7ae6a47fdf42caeb7ee73fb65f4fad04ef9424930 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/groups/item/sets/item/terms/item/relations/item"
)

// TermRequestBuilder builds and executes requests for operations under \sites\{site-id}\termStore\groups\{group-id}\sets\{set-id}\terms\{term-id}
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
// TermRequestBuilderGetQueryParameters all the terms under the set.
type TermRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
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
func (m *TermRequestBuilder) Children()(*ib387154f1adfee1012fa83037220b7890e544dbcaebdc83227c131fda461d9cc.ChildrenRequestBuilder) {
    return ib387154f1adfee1012fa83037220b7890e544dbcaebdc83227c131fda461d9cc.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStore.groups.item.sets.item.terms.item.children.item collection
func (m *TermRequestBuilder) ChildrenById(id string)(*i830382e02f940987ffd4007dfcdd55cfd855f15c1563327664bc6b5c21ac048d.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id1"] = id
    }
    return i830382e02f940987ffd4007dfcdd55cfd855f15c1563327664bc6b5c21ac048d.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTermRequestBuilderInternal instantiates a new TermRequestBuilder and sets the default values.
func NewTermRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermRequestBuilder) {
    m := &TermRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/termStore/groups/{group_id}/sets/{set_id}/terms/{term_id}{?select,expand}";
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
// CreateDeleteRequestInformation all the terms under the set.
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
// CreateGetRequestInformation all the terms under the set.
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
// CreatePatchRequestInformation all the terms under the set.
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
// Delete all the terms under the set.
func (m *TermRequestBuilder) Delete(options *TermRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get all the terms under the set.
func (m *TermRequestBuilder) Get(options *TermRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Term, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewTerm() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Term), nil
}
// Patch all the terms under the set.
func (m *TermRequestBuilder) Patch(options *TermRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *TermRequestBuilder) Relations()(*if2f81ca9fd32dce408153070ba100dec3bbba3f4a78999c825bbc872e5204c37.RelationsRequestBuilder) {
    return if2f81ca9fd32dce408153070ba100dec3bbba3f4a78999c825bbc872e5204c37.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStore.groups.item.sets.item.terms.item.relations.item collection
func (m *TermRequestBuilder) RelationsById(id string)(*ia02cb2beba7e2dc73eabfbe7ae6a47fdf42caeb7ee73fb65f4fad04ef9424930.RelationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation_id"] = id
    }
    return ia02cb2beba7e2dc73eabfbe7ae6a47fdf42caeb7ee73fb65f4fad04ef9424930.NewRelationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TermRequestBuilder) Set()(*ibd4072fbdcfe2f687bf9eb50c1c9cab17ba7feb6b9e477bc6903701a72813cdc.SetRequestBuilder) {
    return ibd4072fbdcfe2f687bf9eb50c1c9cab17ba7feb6b9e477bc6903701a72813cdc.NewSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
