package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i219dc1f6f7a6c6bc20554b924796c145ac68ed6855192bd1ce4325416fcbfc75 "github.com/microsoftgraph/msgraph-sdk-go/contracts/item/restore"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i634333242efa9a29a6ddf99350f54e41bbbf0133e0686bf1f7528b7eefc8aa1f "github.com/microsoftgraph/msgraph-sdk-go/contracts/item/checkmemberobjects"
    i68d315812c96beaf0bb656713cc7f6ec30a202d0bd23957081408380e1e7c72b "github.com/microsoftgraph/msgraph-sdk-go/contracts/item/getmemberobjects"
    i9816d6d5fdd8c2b627440ac47aefe7ebcea7ecbad8b64b00685c0dd443cabc73 "github.com/microsoftgraph/msgraph-sdk-go/contracts/item/checkmembergroups"
    iabffea946aaf9953004feea739fa785cf07d700c84085f7a935b4867be8cfa06 "github.com/microsoftgraph/msgraph-sdk-go/contracts/item/getmembergroups"
)

// Builds and executes requests for operations under \contracts\{contract-id}
type ContractRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ContractRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ContractRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContractRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get entity from contracts by key
type ContractRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ContractRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Contract;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ContractRequestBuilder) CheckMemberGroups()(*i9816d6d5fdd8c2b627440ac47aefe7ebcea7ecbad8b64b00685c0dd443cabc73.CheckMemberGroupsRequestBuilder) {
    return i9816d6d5fdd8c2b627440ac47aefe7ebcea7ecbad8b64b00685c0dd443cabc73.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContractRequestBuilder) CheckMemberObjects()(*i634333242efa9a29a6ddf99350f54e41bbbf0133e0686bf1f7528b7eefc8aa1f.CheckMemberObjectsRequestBuilder) {
    return i634333242efa9a29a6ddf99350f54e41bbbf0133e0686bf1f7528b7eefc8aa1f.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ContractRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewContractRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContractRequestBuilder) {
    m := &ContractRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contracts/{contract_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ContractRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewContractRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContractRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContractRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete entity from contracts
// Parameters:
//  - options : Options for the request
func (m *ContractRequestBuilder) CreateDeleteRequestInformation(options *ContractRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get entity from contracts by key
// Parameters:
//  - options : Options for the request
func (m *ContractRequestBuilder) CreateGetRequestInformation(options *ContractRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Update entity in contracts
// Parameters:
//  - options : Options for the request
func (m *ContractRequestBuilder) CreatePatchRequestInformation(options *ContractRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete entity from contracts
// Parameters:
//  - options : Options for the request
func (m *ContractRequestBuilder) Delete(options *ContractRequestBuilderDeleteOptions)(error) {
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
// Get entity from contracts by key
// Parameters:
//  - options : Options for the request
func (m *ContractRequestBuilder) Get(options *ContractRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Contract, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewContract() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Contract), nil
}
func (m *ContractRequestBuilder) GetMemberGroups()(*iabffea946aaf9953004feea739fa785cf07d700c84085f7a935b4867be8cfa06.GetMemberGroupsRequestBuilder) {
    return iabffea946aaf9953004feea739fa785cf07d700c84085f7a935b4867be8cfa06.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContractRequestBuilder) GetMemberObjects()(*i68d315812c96beaf0bb656713cc7f6ec30a202d0bd23957081408380e1e7c72b.GetMemberObjectsRequestBuilder) {
    return i68d315812c96beaf0bb656713cc7f6ec30a202d0bd23957081408380e1e7c72b.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Update entity in contracts
// Parameters:
//  - options : Options for the request
func (m *ContractRequestBuilder) Patch(options *ContractRequestBuilderPatchOptions)(error) {
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
func (m *ContractRequestBuilder) Restore()(*i219dc1f6f7a6c6bc20554b924796c145ac68ed6855192bd1ce4325416fcbfc75.RestoreRequestBuilder) {
    return i219dc1f6f7a6c6bc20554b924796c145ac68ed6855192bd1ce4325416fcbfc75.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
