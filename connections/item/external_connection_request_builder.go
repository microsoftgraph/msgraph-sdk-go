package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i272adc03cf71cf334688c5afeb2ca23015a19e248696b4eab7d158e92d85a73c "github.com/microsoftgraph/msgraph-sdk-go/connections/item/schema"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i9f701836f6f064da245ba53c4a8cab511ce5c7805cdaa836b4aef8c9af6852a6 "github.com/microsoftgraph/msgraph-sdk-go/connections/item/operations"
    ia84cec0e2ab16fa98bdba4c35c9559797c1bb782a6a6075b268c17aa2dcc36ca "github.com/microsoftgraph/msgraph-sdk-go/connections/item/groups"
    icf71281ca8434d6f3f31e6093b19ce2721c2209cc2f0fed59881aad425ce8fd9 "github.com/microsoftgraph/msgraph-sdk-go/connections/item/items"
    i0587295e9b0a8f2babd875f43631037e119426cc9b33dc66829b00e0d66f13af "github.com/microsoftgraph/msgraph-sdk-go/connections/item/groups/item"
    i2a254f378bc8eba0badf934e4cc3f16461077f8f3e402ec2f5ec50847304c54a "github.com/microsoftgraph/msgraph-sdk-go/connections/item/items/item"
    i9d1670091fc52082d729a9f897ef1d49013bba55d168ff02986d174ec33643a0 "github.com/microsoftgraph/msgraph-sdk-go/connections/item/operations/item"
)

// Builds and executes requests for operations under \connections\{externalConnection-id}
type ExternalConnectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ExternalConnectionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ExternalConnectionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ExternalConnectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get entity from connections by key
type ExternalConnectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ExternalConnectionRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ExternalConnection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new ExternalConnectionRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewExternalConnectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExternalConnectionRequestBuilder) {
    m := &ExternalConnectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/connections/{externalConnection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ExternalConnectionRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewExternalConnectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExternalConnectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExternalConnectionRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete entity from connections
// Parameters:
//  - options : Options for the request
func (m *ExternalConnectionRequestBuilder) CreateDeleteRequestInformation(options *ExternalConnectionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get entity from connections by key
// Parameters:
//  - options : Options for the request
func (m *ExternalConnectionRequestBuilder) CreateGetRequestInformation(options *ExternalConnectionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Update entity in connections
// Parameters:
//  - options : Options for the request
func (m *ExternalConnectionRequestBuilder) CreatePatchRequestInformation(options *ExternalConnectionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete entity from connections
// Parameters:
//  - options : Options for the request
func (m *ExternalConnectionRequestBuilder) Delete(options *ExternalConnectionRequestBuilderDeleteOptions)(error) {
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
// Get entity from connections by key
// Parameters:
//  - options : Options for the request
func (m *ExternalConnectionRequestBuilder) Get(options *ExternalConnectionRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ExternalConnection, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewExternalConnection() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ExternalConnection), nil
}
func (m *ExternalConnectionRequestBuilder) Groups()(*ia84cec0e2ab16fa98bdba4c35c9559797c1bb782a6a6075b268c17aa2dcc36ca.GroupsRequestBuilder) {
    return ia84cec0e2ab16fa98bdba4c35c9559797c1bb782a6a6075b268c17aa2dcc36ca.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.connections.item.groups.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ExternalConnectionRequestBuilder) GroupsById(id string)(*i0587295e9b0a8f2babd875f43631037e119426cc9b33dc66829b00e0d66f13af.ExternalGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalGroup_id"] = id
    }
    return i0587295e9b0a8f2babd875f43631037e119426cc9b33dc66829b00e0d66f13af.NewExternalGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ExternalConnectionRequestBuilder) Items()(*icf71281ca8434d6f3f31e6093b19ce2721c2209cc2f0fed59881aad425ce8fd9.ItemsRequestBuilder) {
    return icf71281ca8434d6f3f31e6093b19ce2721c2209cc2f0fed59881aad425ce8fd9.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.connections.item.items.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ExternalConnectionRequestBuilder) ItemsById(id string)(*i2a254f378bc8eba0badf934e4cc3f16461077f8f3e402ec2f5ec50847304c54a.ExternalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalItem_id"] = id
    }
    return i2a254f378bc8eba0badf934e4cc3f16461077f8f3e402ec2f5ec50847304c54a.NewExternalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ExternalConnectionRequestBuilder) Operations()(*i9f701836f6f064da245ba53c4a8cab511ce5c7805cdaa836b4aef8c9af6852a6.OperationsRequestBuilder) {
    return i9f701836f6f064da245ba53c4a8cab511ce5c7805cdaa836b4aef8c9af6852a6.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.connections.item.operations.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ExternalConnectionRequestBuilder) OperationsById(id string)(*i9d1670091fc52082d729a9f897ef1d49013bba55d168ff02986d174ec33643a0.ConnectionOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connectionOperation_id"] = id
    }
    return i9d1670091fc52082d729a9f897ef1d49013bba55d168ff02986d174ec33643a0.NewConnectionOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Update entity in connections
// Parameters:
//  - options : Options for the request
func (m *ExternalConnectionRequestBuilder) Patch(options *ExternalConnectionRequestBuilderPatchOptions)(error) {
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
func (m *ExternalConnectionRequestBuilder) Schema()(*i272adc03cf71cf334688c5afeb2ca23015a19e248696b4eab7d158e92d85a73c.SchemaRequestBuilder) {
    return i272adc03cf71cf334688c5afeb2ca23015a19e248696b4eab7d158e92d85a73c.NewSchemaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
