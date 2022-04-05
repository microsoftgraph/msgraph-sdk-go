package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i340d5d9b1c93dce1b5a110cf5bbac8114c9a4b58e229c9166305cebe2c92e1cb "github.com/microsoftgraph/msgraph-sdk-go/agreements/item/acceptances"
    i534405002975c09989779fd4d2b641edcec6af4615b71a008c4aad75a339f35b "github.com/microsoftgraph/msgraph-sdk-go/agreements/item/files"
    ibfafeebed2dd285e7305e7475c15f9e74a95731aa8cf08121e9f6a0d20ce2d92 "github.com/microsoftgraph/msgraph-sdk-go/agreements/item/file"
    ic5f55850233142ac7d28f19d006edf4c80dee8f569419b3a03d839aeeaba1991 "github.com/microsoftgraph/msgraph-sdk-go/agreements/item/acceptances/item"
    id334b355b0e9c1a808648dc507c9c2722906e5d1665f782a1e2496b2a183da94 "github.com/microsoftgraph/msgraph-sdk-go/agreements/item/files/item"
)

// AgreementItemRequestBuilder provides operations to manage the collection of agreement entities.
type AgreementItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AgreementItemRequestBuilderDeleteOptions options for Delete
type AgreementItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AgreementItemRequestBuilderGetOptions options for Get
type AgreementItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *AgreementItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AgreementItemRequestBuilderGetQueryParameters get entity from agreements by key
type AgreementItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// AgreementItemRequestBuilderPatchOptions options for Patch
type AgreementItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Agreementable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Acceptances the acceptances property
func (m *AgreementItemRequestBuilder) Acceptances()(*i340d5d9b1c93dce1b5a110cf5bbac8114c9a4b58e229c9166305cebe2c92e1cb.AcceptancesRequestBuilder) {
    return i340d5d9b1c93dce1b5a110cf5bbac8114c9a4b58e229c9166305cebe2c92e1cb.NewAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcceptancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.agreements.item.acceptances.item collection
func (m *AgreementItemRequestBuilder) AcceptancesById(id string)(*ic5f55850233142ac7d28f19d006edf4c80dee8f569419b3a03d839aeeaba1991.AgreementAcceptanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance_id"] = id
    }
    return ic5f55850233142ac7d28f19d006edf4c80dee8f569419b3a03d839aeeaba1991.NewAgreementAcceptanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAgreementItemRequestBuilderInternal instantiates a new AgreementItemRequestBuilder and sets the default values.
func NewAgreementItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AgreementItemRequestBuilder) {
    m := &AgreementItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/agreements/{agreement_id}{?select}";
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
// CreateDeleteRequestInformation delete entity from agreements
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
// CreateGetRequestInformation get entity from agreements by key
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
// CreatePatchRequestInformation update entity in agreements
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
// Delete delete entity from agreements
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
func (m *AgreementItemRequestBuilder) File()(*ibfafeebed2dd285e7305e7475c15f9e74a95731aa8cf08121e9f6a0d20ce2d92.FileRequestBuilder) {
    return ibfafeebed2dd285e7305e7475c15f9e74a95731aa8cf08121e9f6a0d20ce2d92.NewFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Files the files property
func (m *AgreementItemRequestBuilder) Files()(*i534405002975c09989779fd4d2b641edcec6af4615b71a008c4aad75a339f35b.FilesRequestBuilder) {
    return i534405002975c09989779fd4d2b641edcec6af4615b71a008c4aad75a339f35b.NewFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.agreements.item.files.item collection
func (m *AgreementItemRequestBuilder) FilesById(id string)(*id334b355b0e9c1a808648dc507c9c2722906e5d1665f782a1e2496b2a183da94.AgreementFileLocalizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementFileLocalization_id"] = id
    }
    return id334b355b0e9c1a808648dc507c9c2722906e5d1665f782a1e2496b2a183da94.NewAgreementFileLocalizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from agreements by key
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
// Patch update entity in agreements
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
