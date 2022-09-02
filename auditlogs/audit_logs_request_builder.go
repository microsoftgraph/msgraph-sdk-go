package auditlogs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    i097180ed0235aaf99aba24bd4318eaecbd7dc30663c7e16f228dcaf593408050 "github.com/microsoftgraph/msgraph-sdk-go/auditlogs/provisioning"
    i0ffcbd4df9c49ebdf92c325c07b99e822be86d1d7c8ef35696fe1f4c2ebe8ddd "github.com/microsoftgraph/msgraph-sdk-go/auditlogs/directoryaudits"
    i6856047a3d7bd4c97c5dd10b2289247f2be8a26b5d67f5c0130850c9e19185eb "github.com/microsoftgraph/msgraph-sdk-go/auditlogs/signins"
    i7625071e63047df95c3e60a95e06a7af646aea10ec9f35f446c2da5700131f2b "github.com/microsoftgraph/msgraph-sdk-go/auditlogs/restrictedsignins"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i1a750e38628ba45883889a3ea54a13c3e3772956da70ded7ae26f13e0bbbedd2 "github.com/microsoftgraph/msgraph-sdk-go/auditlogs/provisioning/item"
    i72f75127bd05e9eb686f047ff40a72bc034bcb41296fc9a2b3b1cc35881c4f3e "github.com/microsoftgraph/msgraph-sdk-go/auditlogs/signins/item"
    i7f3e11d84b38e07c2ad13447d9693f755cb3ec92773d58ec7b0f68b62669249f "github.com/microsoftgraph/msgraph-sdk-go/auditlogs/restrictedsignins/item"
    i8ab9c44a4a415e76c346c3f6eda2e19ce5ffba6ea1f21df106cb58013b2ec53f "github.com/microsoftgraph/msgraph-sdk-go/auditlogs/directoryaudits/item"
)

// AuditLogsRequestBuilder provides operations to manage the auditLogRoot singleton.
type AuditLogsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuditLogsRequestBuilderGetQueryParameters get auditLogs
type AuditLogsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuditLogsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuditLogsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuditLogsRequestBuilderGetQueryParameters
}
// AuditLogsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuditLogsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuditLogsRequestBuilderInternal instantiates a new AuditLogsRequestBuilder and sets the default values.
func NewAuditLogsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditLogsRequestBuilder) {
    m := &AuditLogsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/auditLogs{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuditLogsRequestBuilder instantiates a new AuditLogsRequestBuilder and sets the default values.
func NewAuditLogsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditLogsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuditLogsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get auditLogs
func (m *AuditLogsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get auditLogs
func (m *AuditLogsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AuditLogsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update auditLogs
func (m *AuditLogsRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuditLogRootable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update auditLogs
func (m *AuditLogsRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuditLogRootable, requestConfiguration *AuditLogsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DirectoryAudits the directoryAudits property
func (m *AuditLogsRequestBuilder) DirectoryAudits()(*i0ffcbd4df9c49ebdf92c325c07b99e822be86d1d7c8ef35696fe1f4c2ebe8ddd.DirectoryAuditsRequestBuilder) {
    return i0ffcbd4df9c49ebdf92c325c07b99e822be86d1d7c8ef35696fe1f4c2ebe8ddd.NewDirectoryAuditsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryAuditsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.auditLogs.directoryAudits.item collection
func (m *AuditLogsRequestBuilder) DirectoryAuditsById(id string)(*i8ab9c44a4a415e76c346c3f6eda2e19ce5ffba6ea1f21df106cb58013b2ec53f.DirectoryAuditItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryAudit%2Did"] = id
    }
    return i8ab9c44a4a415e76c346c3f6eda2e19ce5ffba6ea1f21df106cb58013b2ec53f.NewDirectoryAuditItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get auditLogs
func (m *AuditLogsRequestBuilder) Get(ctx context.Context, requestConfiguration *AuditLogsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuditLogRootable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuditLogRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuditLogRootable), nil
}
// Patch update auditLogs
func (m *AuditLogsRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuditLogRootable, requestConfiguration *AuditLogsRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Provisioning the provisioning property
func (m *AuditLogsRequestBuilder) Provisioning()(*i097180ed0235aaf99aba24bd4318eaecbd7dc30663c7e16f228dcaf593408050.ProvisioningRequestBuilder) {
    return i097180ed0235aaf99aba24bd4318eaecbd7dc30663c7e16f228dcaf593408050.NewProvisioningRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProvisioningById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.auditLogs.provisioning.item collection
func (m *AuditLogsRequestBuilder) ProvisioningById(id string)(*i1a750e38628ba45883889a3ea54a13c3e3772956da70ded7ae26f13e0bbbedd2.ProvisioningObjectSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["provisioningObjectSummary%2Did"] = id
    }
    return i1a750e38628ba45883889a3ea54a13c3e3772956da70ded7ae26f13e0bbbedd2.NewProvisioningObjectSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RestrictedSignIns the restrictedSignIns property
func (m *AuditLogsRequestBuilder) RestrictedSignIns()(*i7625071e63047df95c3e60a95e06a7af646aea10ec9f35f446c2da5700131f2b.RestrictedSignInsRequestBuilder) {
    return i7625071e63047df95c3e60a95e06a7af646aea10ec9f35f446c2da5700131f2b.NewRestrictedSignInsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RestrictedSignInsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.auditLogs.restrictedSignIns.item collection
func (m *AuditLogsRequestBuilder) RestrictedSignInsById(id string)(*i7f3e11d84b38e07c2ad13447d9693f755cb3ec92773d58ec7b0f68b62669249f.RestrictedSignInItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["restrictedSignIn%2Did"] = id
    }
    return i7f3e11d84b38e07c2ad13447d9693f755cb3ec92773d58ec7b0f68b62669249f.NewRestrictedSignInItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SignIns the signIns property
func (m *AuditLogsRequestBuilder) SignIns()(*i6856047a3d7bd4c97c5dd10b2289247f2be8a26b5d67f5c0130850c9e19185eb.SignInsRequestBuilder) {
    return i6856047a3d7bd4c97c5dd10b2289247f2be8a26b5d67f5c0130850c9e19185eb.NewSignInsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SignInsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.auditLogs.signIns.item collection
func (m *AuditLogsRequestBuilder) SignInsById(id string)(*i72f75127bd05e9eb686f047ff40a72bc034bcb41296fc9a2b3b1cc35881c4f3e.SignInItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["signIn%2Did"] = id
    }
    return i72f75127bd05e9eb686f047ff40a72bc034bcb41296fc9a2b3b1cc35881c4f3e.NewSignInItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
