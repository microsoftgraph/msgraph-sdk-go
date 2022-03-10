package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/termstore"
    i2696ae20356e7a7bd523a7c9ecbe7f1d75673920fbc84261856981408b9d68ba "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item/parentgroup/sets/item/children/item/relations/item/fromterm"
    i586fceed288258be07ac8ea286f420c601837b5e64db88c03aba7ccd423c75f3 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item/parentgroup/sets/item/children/item/relations/item/set"
    i85176a849393755d393c7ffa4e83c98b21915ded2172ba0d83ece1c732db4181 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item/parentgroup/sets/item/children/item/relations/item/toterm"
)

// RelationItemRequestBuilder provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
type RelationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RelationItemRequestBuilderDeleteOptions options for Delete
type RelationItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RelationItemRequestBuilderGetOptions options for Get
type RelationItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RelationItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RelationItemRequestBuilderGetQueryParameters to indicate which terms are related to the current term as either pinned or reused.
type RelationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// RelationItemRequestBuilderPatchOptions options for Patch
type RelationItemRequestBuilderPatchOptions struct {
    // 
    Body id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.Relationable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewRelationItemRequestBuilderInternal instantiates a new RelationItemRequestBuilder and sets the default values.
func NewRelationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RelationItemRequestBuilder) {
    m := &RelationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/termStores/{store_id}/sets/{set_id}/parentGroup/sets/{set_id1}/children/{term_id}/relations/{relation_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRelationItemRequestBuilder instantiates a new RelationItemRequestBuilder and sets the default values.
func NewRelationItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RelationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRelationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property relations for sites
func (m *RelationItemRequestBuilder) CreateDeleteRequestInformation(options *RelationItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation to indicate which terms are related to the current term as either pinned or reused.
func (m *RelationItemRequestBuilder) CreateGetRequestInformation(options *RelationItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property relations in sites
func (m *RelationItemRequestBuilder) CreatePatchRequestInformation(options *RelationItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property relations for sites
func (m *RelationItemRequestBuilder) Delete(options *RelationItemRequestBuilderDeleteOptions)(error) {
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
func (m *RelationItemRequestBuilder) FromTerm()(*i2696ae20356e7a7bd523a7c9ecbe7f1d75673920fbc84261856981408b9d68ba.FromTermRequestBuilder) {
    return i2696ae20356e7a7bd523a7c9ecbe7f1d75673920fbc84261856981408b9d68ba.NewFromTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get to indicate which terms are related to the current term as either pinned or reused.
func (m *RelationItemRequestBuilder) Get(options *RelationItemRequestBuilderGetOptions)(id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.Relationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.CreateRelationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.Relationable), nil
}
// Patch update the navigation property relations in sites
func (m *RelationItemRequestBuilder) Patch(options *RelationItemRequestBuilderPatchOptions)(error) {
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
func (m *RelationItemRequestBuilder) Set()(*i586fceed288258be07ac8ea286f420c601837b5e64db88c03aba7ccd423c75f3.SetRequestBuilder) {
    return i586fceed288258be07ac8ea286f420c601837b5e64db88c03aba7ccd423c75f3.NewSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RelationItemRequestBuilder) ToTerm()(*i85176a849393755d393c7ffa4e83c98b21915ded2172ba0d83ece1c732db4181.ToTermRequestBuilder) {
    return i85176a849393755d393c7ffa4e83c98b21915ded2172ba0d83ece1c732db4181.NewToTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
