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

// AuditLogsRequestBuilder builds and executes requests for operations under \auditLogs
type AuditLogsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AuditLogsRequestBuilderGetOptions options for Get
type AuditLogsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AuditLogsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AuditLogsRequestBuilderGetQueryParameters get auditLogs
type AuditLogsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AuditLogsRequestBuilderPatchOptions options for Patch
type AuditLogsRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AuditLogRoot;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAuditLogsRequestBuilderInternal instantiates a new AuditLogsRequestBuilder and sets the default values.
func NewAuditLogsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuditLogsRequestBuilder) {
    m := &AuditLogsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/auditLogs{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuditLogsRequestBuilder instantiates a new AuditLogsRequestBuilder and sets the default values.
func NewAuditLogsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuditLogsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuditLogsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get auditLogs
func (m *AuditLogsRequestBuilder) CreateGetRequestInformation(options *AuditLogsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update auditLogs
func (m *AuditLogsRequestBuilder) CreatePatchRequestInformation(options *AuditLogsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AuditLogsRequestBuilder) DirectoryAudits()(*i0ffcbd4df9c49ebdf92c325c07b99e822be86d1d7c8ef35696fe1f4c2ebe8ddd.DirectoryAuditsRequestBuilder) {
    return i0ffcbd4df9c49ebdf92c325c07b99e822be86d1d7c8ef35696fe1f4c2ebe8ddd.NewDirectoryAuditsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryAuditsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.auditLogs.directoryAudits.item collection
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
// Get get auditLogs
func (m *AuditLogsRequestBuilder) Get(options *AuditLogsRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AuditLogRoot, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewAuditLogRoot() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AuditLogRoot), nil
}
// Patch update auditLogs
func (m *AuditLogsRequestBuilder) Patch(options *AuditLogsRequestBuilderPatchOptions)(error) {
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
func (m *AuditLogsRequestBuilder) Provisioning()(*i097180ed0235aaf99aba24bd4318eaecbd7dc30663c7e16f228dcaf593408050.ProvisioningRequestBuilder) {
    return i097180ed0235aaf99aba24bd4318eaecbd7dc30663c7e16f228dcaf593408050.NewProvisioningRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProvisioningById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.auditLogs.provisioning.item collection
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
// RestrictedSignInsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.auditLogs.restrictedSignIns.item collection
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
// SignInsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.auditLogs.signIns.item collection
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
