package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i1bd9f6fcec8cad5cf057c73a59840be7a45f6df4cd6e9cf1680b44c2bc8d79e3 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/termsofuse/agreements/item/files"
    i497cbf408523a7308e7d74d59fe980755bb06512cff2a1505fefee52ffc5f482 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/termsofuse/agreements/item/acceptances"
    i70f02eb4dfaff6e6823bf1a53edf2c9ba3ecdf930d6d4a42d9df9b5a5a136631 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/termsofuse/agreements/item/file"
    i22d600b0c4aaca0bdb6422ea093131f6b276800efda0ef10e38f62d8c5ee6950 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/termsofuse/agreements/item/acceptances/item"
    iaac740f8096136ae314c81411ffdcd09c6315014dc38c6f4f71a1f25f94237bd "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/termsofuse/agreements/item/files/item"
)

// AgreementItemRequestBuilder provides operations to manage the agreements property of the microsoft.graph.termsOfUseContainer entity.
type AgreementItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AgreementItemRequestBuilderDeleteOptions options for Delete
type AgreementItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AgreementItemRequestBuilderGetOptions options for Get
type AgreementItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AgreementItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AgreementItemRequestBuilderGetQueryParameters represents a tenant's customizable terms of use agreement that's created and managed with Azure Active Directory (Azure AD).
type AgreementItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AgreementItemRequestBuilderPatchOptions options for Patch
type AgreementItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Agreementable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Acceptances the acceptances property
func (m *AgreementItemRequestBuilder) Acceptances()(*i497cbf408523a7308e7d74d59fe980755bb06512cff2a1505fefee52ffc5f482.AcceptancesRequestBuilder) {
    return i497cbf408523a7308e7d74d59fe980755bb06512cff2a1505fefee52ffc5f482.NewAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcceptancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.termsOfUse.agreements.item.acceptances.item collection
func (m *AgreementItemRequestBuilder) AcceptancesById(id string)(*i22d600b0c4aaca0bdb6422ea093131f6b276800efda0ef10e38f62d8c5ee6950.AgreementAcceptanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance%2Did"] = id
    }
    return i22d600b0c4aaca0bdb6422ea093131f6b276800efda0ef10e38f62d8c5ee6950.NewAgreementAcceptanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAgreementItemRequestBuilderInternal instantiates a new AgreementItemRequestBuilder and sets the default values.
func NewAgreementItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AgreementItemRequestBuilder) {
    m := &AgreementItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/termsOfUse/agreements/{agreement%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAgreementItemRequestBuilder instantiates a new AgreementItemRequestBuilder and sets the default values.
func NewAgreementItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AgreementItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAgreementItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property agreements for identityGovernance
func (m *AgreementItemRequestBuilder) CreateDeleteRequestInformation(options *AgreementItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
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
// CreateGetRequestInformation represents a tenant's customizable terms of use agreement that's created and managed with Azure Active Directory (Azure AD).
func (m *AgreementItemRequestBuilder) CreateGetRequestInformation(options *AgreementItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property agreements in identityGovernance
func (m *AgreementItemRequestBuilder) CreatePatchRequestInformation(options *AgreementItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property agreements for identityGovernance
func (m *AgreementItemRequestBuilder) Delete(options *AgreementItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
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
// File the file property
func (m *AgreementItemRequestBuilder) File()(*i70f02eb4dfaff6e6823bf1a53edf2c9ba3ecdf930d6d4a42d9df9b5a5a136631.FileRequestBuilder) {
    return i70f02eb4dfaff6e6823bf1a53edf2c9ba3ecdf930d6d4a42d9df9b5a5a136631.NewFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Files the files property
func (m *AgreementItemRequestBuilder) Files()(*i1bd9f6fcec8cad5cf057c73a59840be7a45f6df4cd6e9cf1680b44c2bc8d79e3.FilesRequestBuilder) {
    return i1bd9f6fcec8cad5cf057c73a59840be7a45f6df4cd6e9cf1680b44c2bc8d79e3.NewFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.termsOfUse.agreements.item.files.item collection
func (m *AgreementItemRequestBuilder) FilesById(id string)(*iaac740f8096136ae314c81411ffdcd09c6315014dc38c6f4f71a1f25f94237bd.AgreementFileLocalizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementFileLocalization%2Did"] = id
    }
    return iaac740f8096136ae314c81411ffdcd09c6315014dc38c6f4f71a1f25f94237bd.NewAgreementFileLocalizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get represents a tenant's customizable terms of use agreement that's created and managed with Azure Active Directory (Azure AD).
func (m *AgreementItemRequestBuilder) Get(options *AgreementItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Agreementable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAgreementFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Agreementable), nil
}
// Patch update the navigation property agreements in identityGovernance
func (m *AgreementItemRequestBuilder) Patch(options *AgreementItemRequestBuilderPatchOptions)(error) {
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
