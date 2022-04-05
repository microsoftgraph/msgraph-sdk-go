package identityprotection

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    i734828c35d0209cc42d94bd3da5e79564f27866a2e7dd505cec26e26172ddd00 "github.com/microsoftgraph/msgraph-sdk-go/identityprotection/riskdetections"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    if659cd8a8fa7e51b07cf2aa293599e63216814779670d31cfb9fa4a2facb3f81 "github.com/microsoftgraph/msgraph-sdk-go/identityprotection/riskyusers"
    i4fd6061886571b3d64bdb991095eb35ae40fd05f187e9565e0e84610c3889be0 "github.com/microsoftgraph/msgraph-sdk-go/identityprotection/riskyusers/item"
    i919d73984bc501fa850d0ae7fc63eed2d0af650bd7a764d682b904e3619f0c40 "github.com/microsoftgraph/msgraph-sdk-go/identityprotection/riskdetections/item"
)

// IdentityProtectionRequestBuilder provides operations to manage the identityProtectionRoot singleton.
type IdentityProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// IdentityProtectionRequestBuilderGetOptions options for Get
type IdentityProtectionRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *IdentityProtectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// IdentityProtectionRequestBuilderGetQueryParameters get identityProtection
type IdentityProtectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// IdentityProtectionRequestBuilderPatchOptions options for Patch
type IdentityProtectionRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityProtectionRootable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewIdentityProtectionRequestBuilderInternal instantiates a new IdentityProtectionRequestBuilder and sets the default values.
func NewIdentityProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityProtectionRequestBuilder) {
    m := &IdentityProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityProtection{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityProtectionRequestBuilder instantiates a new IdentityProtectionRequestBuilder and sets the default values.
func NewIdentityProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get identityProtection
func (m *IdentityProtectionRequestBuilder) CreateGetRequestInformation(options *IdentityProtectionRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update identityProtection
func (m *IdentityProtectionRequestBuilder) CreatePatchRequestInformation(options *IdentityProtectionRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get get identityProtection
func (m *IdentityProtectionRequestBuilder) Get(options *IdentityProtectionRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityProtectionRootable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIdentityProtectionRootFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityProtectionRootable), nil
}
// Patch update identityProtection
func (m *IdentityProtectionRequestBuilder) Patch(options *IdentityProtectionRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// RiskDetections the riskDetections property
func (m *IdentityProtectionRequestBuilder) RiskDetections()(*i734828c35d0209cc42d94bd3da5e79564f27866a2e7dd505cec26e26172ddd00.RiskDetectionsRequestBuilder) {
    return i734828c35d0209cc42d94bd3da5e79564f27866a2e7dd505cec26e26172ddd00.NewRiskDetectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RiskDetectionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityProtection.riskDetections.item collection
func (m *IdentityProtectionRequestBuilder) RiskDetectionsById(id string)(*i919d73984bc501fa850d0ae7fc63eed2d0af650bd7a764d682b904e3619f0c40.RiskDetectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["riskDetection_id"] = id
    }
    return i919d73984bc501fa850d0ae7fc63eed2d0af650bd7a764d682b904e3619f0c40.NewRiskDetectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RiskyUsers the riskyUsers property
func (m *IdentityProtectionRequestBuilder) RiskyUsers()(*if659cd8a8fa7e51b07cf2aa293599e63216814779670d31cfb9fa4a2facb3f81.RiskyUsersRequestBuilder) {
    return if659cd8a8fa7e51b07cf2aa293599e63216814779670d31cfb9fa4a2facb3f81.NewRiskyUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RiskyUsersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityProtection.riskyUsers.item collection
func (m *IdentityProtectionRequestBuilder) RiskyUsersById(id string)(*i4fd6061886571b3d64bdb991095eb35ae40fd05f187e9565e0e84610c3889be0.RiskyUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["riskyUser_id"] = id
    }
    return i4fd6061886571b3d64bdb991095eb35ae40fd05f187e9565e0e84610c3889be0.NewRiskyUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
