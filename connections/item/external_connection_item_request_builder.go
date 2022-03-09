package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i272adc03cf71cf334688c5afeb2ca23015a19e248696b4eab7d158e92d85a73c "github.com/microsoftgraph/msgraph-sdk-go/connections/item/schema"
    i9f701836f6f064da245ba53c4a8cab511ce5c7805cdaa836b4aef8c9af6852a6 "github.com/microsoftgraph/msgraph-sdk-go/connections/item/operations"
    ia84cec0e2ab16fa98bdba4c35c9559797c1bb782a6a6075b268c17aa2dcc36ca "github.com/microsoftgraph/msgraph-sdk-go/connections/item/groups"
    icf71281ca8434d6f3f31e6093b19ce2721c2209cc2f0fed59881aad425ce8fd9 "github.com/microsoftgraph/msgraph-sdk-go/connections/item/items"
    i0587295e9b0a8f2babd875f43631037e119426cc9b33dc66829b00e0d66f13af "github.com/microsoftgraph/msgraph-sdk-go/connections/item/groups/item"
    i2a254f378bc8eba0badf934e4cc3f16461077f8f3e402ec2f5ec50847304c54a "github.com/microsoftgraph/msgraph-sdk-go/connections/item/items/item"
    i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/externalconnectors"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i9d1670091fc52082d729a9f897ef1d49013bba55d168ff02986d174ec33643a0 "github.com/microsoftgraph/msgraph-sdk-go/connections/item/operations/item"
)

// ExternalConnectionItemRequestBuilder provides operations to manage the collection of externalConnection entities.
type ExternalConnectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ExternalConnectionItemRequestBuilderDeleteOptions options for Delete
type ExternalConnectionItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ExternalConnectionItemRequestBuilderGetOptions options for Get
type ExternalConnectionItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ExternalConnectionItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ExternalConnectionItemRequestBuilderGetQueryParameters get entity from connections by key
type ExternalConnectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ExternalConnectionItemRequestBuilderPatchOptions options for Patch
type ExternalConnectionItemRequestBuilderPatchOptions struct {
    // 
    Body i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ExternalConnectionable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewExternalConnectionItemRequestBuilderInternal instantiates a new ExternalConnectionItemRequestBuilder and sets the default values.
func NewExternalConnectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExternalConnectionItemRequestBuilder) {
    m := &ExternalConnectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/connections/{externalConnection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExternalConnectionItemRequestBuilder instantiates a new ExternalConnectionItemRequestBuilder and sets the default values.
func NewExternalConnectionItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExternalConnectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExternalConnectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from connections
func (m *ExternalConnectionItemRequestBuilder) CreateDeleteRequestInformation(options *ExternalConnectionItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from connections by key
func (m *ExternalConnectionItemRequestBuilder) CreateGetRequestInformation(options *ExternalConnectionItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in connections
func (m *ExternalConnectionItemRequestBuilder) CreatePatchRequestInformation(options *ExternalConnectionItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from connections
func (m *ExternalConnectionItemRequestBuilder) Delete(options *ExternalConnectionItemRequestBuilderDeleteOptions)(error) {
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
// Get get entity from connections by key
func (m *ExternalConnectionItemRequestBuilder) Get(options *ExternalConnectionItemRequestBuilderGetOptions)(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ExternalConnectionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.CreateExternalConnectionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ExternalConnectionable), nil
}
func (m *ExternalConnectionItemRequestBuilder) Groups()(*ia84cec0e2ab16fa98bdba4c35c9559797c1bb782a6a6075b268c17aa2dcc36ca.GroupsRequestBuilder) {
    return ia84cec0e2ab16fa98bdba4c35c9559797c1bb782a6a6075b268c17aa2dcc36ca.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.connections.item.groups.item collection
func (m *ExternalConnectionItemRequestBuilder) GroupsById(id string)(*i0587295e9b0a8f2babd875f43631037e119426cc9b33dc66829b00e0d66f13af.ExternalGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalGroup_id"] = id
    }
    return i0587295e9b0a8f2babd875f43631037e119426cc9b33dc66829b00e0d66f13af.NewExternalGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ExternalConnectionItemRequestBuilder) Items()(*icf71281ca8434d6f3f31e6093b19ce2721c2209cc2f0fed59881aad425ce8fd9.ItemsRequestBuilder) {
    return icf71281ca8434d6f3f31e6093b19ce2721c2209cc2f0fed59881aad425ce8fd9.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.connections.item.items.item collection
func (m *ExternalConnectionItemRequestBuilder) ItemsById(id string)(*i2a254f378bc8eba0badf934e4cc3f16461077f8f3e402ec2f5ec50847304c54a.ExternalItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalItem_id"] = id
    }
    return i2a254f378bc8eba0badf934e4cc3f16461077f8f3e402ec2f5ec50847304c54a.NewExternalItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ExternalConnectionItemRequestBuilder) Operations()(*i9f701836f6f064da245ba53c4a8cab511ce5c7805cdaa836b4aef8c9af6852a6.OperationsRequestBuilder) {
    return i9f701836f6f064da245ba53c4a8cab511ce5c7805cdaa836b4aef8c9af6852a6.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.connections.item.operations.item collection
func (m *ExternalConnectionItemRequestBuilder) OperationsById(id string)(*i9d1670091fc52082d729a9f897ef1d49013bba55d168ff02986d174ec33643a0.ConnectionOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connectionOperation_id"] = id
    }
    return i9d1670091fc52082d729a9f897ef1d49013bba55d168ff02986d174ec33643a0.NewConnectionOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in connections
func (m *ExternalConnectionItemRequestBuilder) Patch(options *ExternalConnectionItemRequestBuilderPatchOptions)(error) {
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
func (m *ExternalConnectionItemRequestBuilder) Schema()(*i272adc03cf71cf334688c5afeb2ca23015a19e248696b4eab7d158e92d85a73c.SchemaRequestBuilder) {
    return i272adc03cf71cf334688c5afeb2ca23015a19e248696b4eab7d158e92d85a73c.NewSchemaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
