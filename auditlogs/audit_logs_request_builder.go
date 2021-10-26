package auditlogs

import (
    i097180ed0235aaf99aba24bd4318eaecbd7dc30663c7e16f228dcaf593408050 "github.com/microsoftgraph/msgraph-sdk-go/auditlogs/provisioning"
    i0ffcbd4df9c49ebdf92c325c07b99e822be86d1d7c8ef35696fe1f4c2ebe8ddd "github.com/microsoftgraph/msgraph-sdk-go/auditlogs/directoryaudits"
    i6856047a3d7bd4c97c5dd10b2289247f2be8a26b5d67f5c0130850c9e19185eb "github.com/microsoftgraph/msgraph-sdk-go/auditlogs/signins"
    i7625071e63047df95c3e60a95e06a7af646aea10ec9f35f446c2da5700131f2b "github.com/microsoftgraph/msgraph-sdk-go/auditlogs/restrictedsignins"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i1a750e38628ba45883889a3ea54a13c3e3772956da70ded7ae26f13e0bbbedd2 "github.com/microsoftgraph/msgraph-sdk-go/auditlogs/provisioning/item"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i72f75127bd05e9eb686f047ff40a72bc034bcb41296fc9a2b3b1cc35881c4f3e "github.com/microsoftgraph/msgraph-sdk-go/auditlogs/signins/item"
    i7f3e11d84b38e07c2ad13447d9693f755cb3ec92773d58ec7b0f68b62669249f "github.com/microsoftgraph/msgraph-sdk-go/auditlogs/restrictedsignins/item"
    i8ab9c44a4a415e76c346c3f6eda2e19ce5ffba6ea1f21df106cb58013b2ec53f "github.com/microsoftgraph/msgraph-sdk-go/auditlogs/directoryaudits/item"
)

// Builds and executes requests for operations under \auditLogs
type AuditLogsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Get auditLogs
type AuditLogsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Instantiates a new AuditLogsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAuditLogsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuditLogsRequestBuilder) {
    m := &AuditLogsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/auditLogs{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new AuditLogsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAuditLogsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuditLogsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuditLogsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get auditLogs
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *AuditLogsRequestBuilder) CreateGetRequestInformation(q func (value *AuditLogsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(AuditLogsRequestBuilderGetQueryParameters)
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
// Update auditLogs
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *AuditLogsRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AuditLogRoot, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AuditLogsRequestBuilder) DirectoryAudits()(*i0ffcbd4df9c49ebdf92c325c07b99e822be86d1d7c8ef35696fe1f4c2ebe8ddd.DirectoryAuditsRequestBuilder) {
    return i0ffcbd4df9c49ebdf92c325c07b99e822be86d1d7c8ef35696fe1f4c2ebe8ddd.NewDirectoryAuditsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.auditLogs.directoryAudits.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AuditLogsRequestBuilder) DirectoryAuditsById(id string)(*i8ab9c44a4a415e76c346c3f6eda2e19ce5ffba6ea1f21df106cb58013b2ec53f.DirectoryAuditRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryAudit_id"] = id
    }
    return i8ab9c44a4a415e76c346c3f6eda2e19ce5ffba6ea1f21df106cb58013b2ec53f.NewDirectoryAuditRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get auditLogs
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *AuditLogsRequestBuilder) Get(q func (value *AuditLogsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AuditLogRoot, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewAuditLogRoot() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AuditLogRoot), nil
}
// Update auditLogs
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *AuditLogsRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AuditLogRoot, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *AuditLogsRequestBuilder) Provisioning()(*i097180ed0235aaf99aba24bd4318eaecbd7dc30663c7e16f228dcaf593408050.ProvisioningRequestBuilder) {
    return i097180ed0235aaf99aba24bd4318eaecbd7dc30663c7e16f228dcaf593408050.NewProvisioningRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.auditLogs.provisioning.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AuditLogsRequestBuilder) ProvisioningById(id string)(*i1a750e38628ba45883889a3ea54a13c3e3772956da70ded7ae26f13e0bbbedd2.ProvisioningObjectSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["provisioningObjectSummary_id"] = id
    }
    return i1a750e38628ba45883889a3ea54a13c3e3772956da70ded7ae26f13e0bbbedd2.NewProvisioningObjectSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuditLogsRequestBuilder) RestrictedSignIns()(*i7625071e63047df95c3e60a95e06a7af646aea10ec9f35f446c2da5700131f2b.RestrictedSignInsRequestBuilder) {
    return i7625071e63047df95c3e60a95e06a7af646aea10ec9f35f446c2da5700131f2b.NewRestrictedSignInsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.auditLogs.restrictedSignIns.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AuditLogsRequestBuilder) RestrictedSignInsById(id string)(*i7f3e11d84b38e07c2ad13447d9693f755cb3ec92773d58ec7b0f68b62669249f.RestrictedSignInRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["restrictedSignIn_id"] = id
    }
    return i7f3e11d84b38e07c2ad13447d9693f755cb3ec92773d58ec7b0f68b62669249f.NewRestrictedSignInRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuditLogsRequestBuilder) SignIns()(*i6856047a3d7bd4c97c5dd10b2289247f2be8a26b5d67f5c0130850c9e19185eb.SignInsRequestBuilder) {
    return i6856047a3d7bd4c97c5dd10b2289247f2be8a26b5d67f5c0130850c9e19185eb.NewSignInsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.auditLogs.signIns.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AuditLogsRequestBuilder) SignInsById(id string)(*i72f75127bd05e9eb686f047ff40a72bc034bcb41296fc9a2b3b1cc35881c4f3e.SignInRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["signIn_id"] = id
    }
    return i72f75127bd05e9eb686f047ff40a72bc034bcb41296fc9a2b3b1cc35881c4f3e.NewSignInRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
