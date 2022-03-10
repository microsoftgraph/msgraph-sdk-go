package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/termstore"
    i5c7181a57af37d74c03c0bac93b5fb2e1e78a27e94f9b2cde1023e86e6475105 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/parentgroup/sets/item/relations"
    i7e8ed2eaba08fa0920c3e293bf0bb904dce59da2e1618b06a6e24cf81a5bbaeb "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/parentgroup/sets/item/terms"
    ib4cfb0a713da9bffcf1e91a7f74fb55f54174223e8372ffdfaf3c250a3f7b888 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/parentgroup/sets/item/children"
    i214088727280531f73c8bacb42fd6c2346364377cfabc71d66b4af778e51e7ca "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/parentgroup/sets/item/terms/item"
    id9eaefb5aaca69eea2b2f459ff7eddfd1bb7462c3a6af56b10481aa2d556ec2e "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/parentgroup/sets/item/children/item"
    ied24f74b1e494ab5f96ea80e8f278d294acd87cb7edb884cf45fef2b13368afe "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/sets/item/parentgroup/sets/item/relations/item"
)

// SetItemRequestBuilder provides operations to manage the sets property of the microsoft.graph.termStore.group entity.
type SetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SetItemRequestBuilderDeleteOptions options for Delete
type SetItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SetItemRequestBuilderGetOptions options for Get
type SetItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SetItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SetItemRequestBuilderGetQueryParameters all sets under the group in a term [store].
type SetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SetItemRequestBuilderPatchOptions options for Patch
type SetItemRequestBuilderPatchOptions struct {
    // 
    Body id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.Setable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *SetItemRequestBuilder) Children()(*ib4cfb0a713da9bffcf1e91a7f74fb55f54174223e8372ffdfaf3c250a3f7b888.ChildrenRequestBuilder) {
    return ib4cfb0a713da9bffcf1e91a7f74fb55f54174223e8372ffdfaf3c250a3f7b888.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStore.sets.item.parentGroup.sets.item.children.item collection
func (m *SetItemRequestBuilder) ChildrenById(id string)(*id9eaefb5aaca69eea2b2f459ff7eddfd1bb7462c3a6af56b10481aa2d556ec2e.TermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id"] = id
    }
    return id9eaefb5aaca69eea2b2f459ff7eddfd1bb7462c3a6af56b10481aa2d556ec2e.NewTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSetItemRequestBuilderInternal instantiates a new SetItemRequestBuilder and sets the default values.
func NewSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetItemRequestBuilder) {
    m := &SetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/termStore/sets/{set_id}/parentGroup/sets/{set_id1}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSetItemRequestBuilder instantiates a new SetItemRequestBuilder and sets the default values.
func NewSetItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property sets for sites
func (m *SetItemRequestBuilder) CreateDeleteRequestInformation(options *SetItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SetItemRequestBuilder) CreateGetRequestInformation(options *SetItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property sets in sites
func (m *SetItemRequestBuilder) CreatePatchRequestInformation(options *SetItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property sets for sites
func (m *SetItemRequestBuilder) Delete(options *SetItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get all sets under the group in a term [store].
func (m *SetItemRequestBuilder) Get(options *SetItemRequestBuilderGetOptions)(id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.Setable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.CreateSetFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.Setable), nil
}
// Patch update the navigation property sets in sites
func (m *SetItemRequestBuilder) Patch(options *SetItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *SetItemRequestBuilder) Relations()(*i5c7181a57af37d74c03c0bac93b5fb2e1e78a27e94f9b2cde1023e86e6475105.RelationsRequestBuilder) {
    return i5c7181a57af37d74c03c0bac93b5fb2e1e78a27e94f9b2cde1023e86e6475105.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStore.sets.item.parentGroup.sets.item.relations.item collection
func (m *SetItemRequestBuilder) RelationsById(id string)(*ied24f74b1e494ab5f96ea80e8f278d294acd87cb7edb884cf45fef2b13368afe.RelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation_id"] = id
    }
    return ied24f74b1e494ab5f96ea80e8f278d294acd87cb7edb884cf45fef2b13368afe.NewRelationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SetItemRequestBuilder) Terms()(*i7e8ed2eaba08fa0920c3e293bf0bb904dce59da2e1618b06a6e24cf81a5bbaeb.TermsRequestBuilder) {
    return i7e8ed2eaba08fa0920c3e293bf0bb904dce59da2e1618b06a6e24cf81a5bbaeb.NewTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStore.sets.item.parentGroup.sets.item.terms.item collection
func (m *SetItemRequestBuilder) TermsById(id string)(*i214088727280531f73c8bacb42fd6c2346364377cfabc71d66b4af778e51e7ca.TermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id"] = id
    }
    return i214088727280531f73c8bacb42fd6c2346364377cfabc71d66b4af778e51e7ca.NewTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
