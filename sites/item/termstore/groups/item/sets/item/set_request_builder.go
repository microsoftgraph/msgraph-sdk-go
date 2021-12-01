package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i072b257cbcd70f6e2c89e10ab190d29e5f1cbee922d2b3424c46647f3ccab851 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/groups/item/sets/item/parentgroup"
    i7c235ec7022979e55577adf275aa571630468ed734b3853bb0efa207f0823131 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/groups/item/sets/item/relations"
    i7d3a55057ab8f971a4d92f1e74401911270f92c6acf10c9a99e65ceb566987fd "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/groups/item/sets/item/children"
    i7ffce1ba10aea0f4863b232e8685b9a32b705c164cac74c0df9207f77c54966f "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/groups/item/sets/item/terms"
    i0c833bef238bdedc35009cd060d7757b7364a0ecd956e30cfa332ca4baeb3563 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/groups/item/sets/item/children/item"
    i255bf375c218134404e98cd1c793110101c5f53a820b0e75433e1b1cb6cea84a "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/groups/item/sets/item/terms/item"
    ie8096a99df2895285481f59980d79fd875d310e46b53d21b91efbaa334bf2f72 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore/groups/item/sets/item/relations/item"
)

// SetRequestBuilder builds and executes requests for operations under \sites\{site-id}\termStore\groups\{group-id}\sets\{set-id}
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
// SetRequestBuilderGetQueryParameters all sets under the group in a term [store].
type SetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
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
func (m *SetRequestBuilder) Children()(*i7d3a55057ab8f971a4d92f1e74401911270f92c6acf10c9a99e65ceb566987fd.ChildrenRequestBuilder) {
    return i7d3a55057ab8f971a4d92f1e74401911270f92c6acf10c9a99e65ceb566987fd.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStore.groups.item.sets.item.children.item collection
func (m *SetRequestBuilder) ChildrenById(id string)(*i0c833bef238bdedc35009cd060d7757b7364a0ecd956e30cfa332ca4baeb3563.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id"] = id
    }
    return i0c833bef238bdedc35009cd060d7757b7364a0ecd956e30cfa332ca4baeb3563.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSetRequestBuilderInternal instantiates a new SetRequestBuilder and sets the default values.
func NewSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetRequestBuilder) {
    m := &SetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/termStore/groups/{group_id}/sets/{set_id}{?select,expand}";
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
func (m *SetRequestBuilder) ParentGroup()(*i072b257cbcd70f6e2c89e10ab190d29e5f1cbee922d2b3424c46647f3ccab851.ParentGroupRequestBuilder) {
    return i072b257cbcd70f6e2c89e10ab190d29e5f1cbee922d2b3424c46647f3ccab851.NewParentGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *SetRequestBuilder) Relations()(*i7c235ec7022979e55577adf275aa571630468ed734b3853bb0efa207f0823131.RelationsRequestBuilder) {
    return i7c235ec7022979e55577adf275aa571630468ed734b3853bb0efa207f0823131.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStore.groups.item.sets.item.relations.item collection
func (m *SetRequestBuilder) RelationsById(id string)(*ie8096a99df2895285481f59980d79fd875d310e46b53d21b91efbaa334bf2f72.RelationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation_id"] = id
    }
    return ie8096a99df2895285481f59980d79fd875d310e46b53d21b91efbaa334bf2f72.NewRelationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SetRequestBuilder) Terms()(*i7ffce1ba10aea0f4863b232e8685b9a32b705c164cac74c0df9207f77c54966f.TermsRequestBuilder) {
    return i7ffce1ba10aea0f4863b232e8685b9a32b705c164cac74c0df9207f77c54966f.NewTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStore.groups.item.sets.item.terms.item collection
func (m *SetRequestBuilder) TermsById(id string)(*i255bf375c218134404e98cd1c793110101c5f53a820b0e75433e1b1cb6cea84a.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id"] = id
    }
    return i255bf375c218134404e98cd1c793110101c5f53a820b0e75433e1b1cb6cea84a.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
