package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i1881cafcb69c2a03b070eab982c3f40a50f879626e4f5a433adf3d53764e612f "github.com/microsoftgraph/msgraph-sdk-go/domains/item/forcedelete"
    iacad0a6834b5884a8d686d9e187dfb93e979024e3471c842fc86db3dd765352a "github.com/microsoftgraph/msgraph-sdk-go/domains/item/verify"
    ic9cf916b096ab628ac94150eded9fad2da688dcaced99eb8551954b4daabff85 "github.com/microsoftgraph/msgraph-sdk-go/domains/item/federationconfiguration"
    icf5b8e1a8b1c94340467eae996f377e8f09f9616367b7d5309238180fac55b36 "github.com/microsoftgraph/msgraph-sdk-go/domains/item/verificationdnsrecords"
    id6628f59c761c8d4beb30bc461b4a9033a596bf14d71905302fee3144bffe677 "github.com/microsoftgraph/msgraph-sdk-go/domains/item/domainnamereferences"
    if6c517fa3faff0b607727add22bace8751787df881a7f48b830e8ee8854493e1 "github.com/microsoftgraph/msgraph-sdk-go/domains/item/serviceconfigurationrecords"
    i4eeff4aea5c653c30cf74b43987a1c46ab8d4723858e572edb5f60ff78363be0 "github.com/microsoftgraph/msgraph-sdk-go/domains/item/domainnamereferences/item"
    i920c7da1a3bc0cc91d66e6345a870d4afe1c3ff7eadb815aed8f74ac233de19e "github.com/microsoftgraph/msgraph-sdk-go/domains/item/verificationdnsrecords/item"
    ibe4d7e9a1d86de76f2fb6d99d1444ab26d987f71d4d6f5054eebc0aebb2d6687 "github.com/microsoftgraph/msgraph-sdk-go/domains/item/federationconfiguration/item"
    if238474bb023ee1ae7b1e221337d6649adb30bb8a532c530686588c4d8f149e4 "github.com/microsoftgraph/msgraph-sdk-go/domains/item/serviceconfigurationrecords/item"
)

// DomainItemRequestBuilder provides operations to manage the collection of domain entities.
type DomainItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DomainItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DomainItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DomainItemRequestBuilderGetQueryParameters retrieve the properties and relationships of domain object.
type DomainItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DomainItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DomainItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DomainItemRequestBuilderGetQueryParameters
}
// DomainItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DomainItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDomainItemRequestBuilderInternal instantiates a new DomainItemRequestBuilder and sets the default values.
func NewDomainItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DomainItemRequestBuilder) {
    m := &DomainItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/domains/{domain%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDomainItemRequestBuilder instantiates a new DomainItemRequestBuilder and sets the default values.
func NewDomainItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DomainItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDomainItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation deletes a domain from a tenant.
func (m *DomainItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration deletes a domain from a tenant.
func (m *DomainItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DomainItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation retrieve the properties and relationships of domain object.
func (m *DomainItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration retrieve the properties and relationships of domain object.
func (m *DomainItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DomainItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the properties of domain object.
func (m *DomainItemRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Domainable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the properties of domain object.
func (m *DomainItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Domainable, requestConfiguration *DomainItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete deletes a domain from a tenant.
func (m *DomainItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler deletes a domain from a tenant.
func (m *DomainItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *DomainItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DomainNameReferences the domainNameReferences property
func (m *DomainItemRequestBuilder) DomainNameReferences()(*id6628f59c761c8d4beb30bc461b4a9033a596bf14d71905302fee3144bffe677.DomainNameReferencesRequestBuilder) {
    return id6628f59c761c8d4beb30bc461b4a9033a596bf14d71905302fee3144bffe677.NewDomainNameReferencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DomainNameReferencesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.domains.item.domainNameReferences.item collection
func (m *DomainItemRequestBuilder) DomainNameReferencesById(id string)(*i4eeff4aea5c653c30cf74b43987a1c46ab8d4723858e572edb5f60ff78363be0.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i4eeff4aea5c653c30cf74b43987a1c46ab8d4723858e572edb5f60ff78363be0.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FederationConfiguration the federationConfiguration property
func (m *DomainItemRequestBuilder) FederationConfiguration()(*ic9cf916b096ab628ac94150eded9fad2da688dcaced99eb8551954b4daabff85.FederationConfigurationRequestBuilder) {
    return ic9cf916b096ab628ac94150eded9fad2da688dcaced99eb8551954b4daabff85.NewFederationConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FederationConfigurationById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.domains.item.federationConfiguration.item collection
func (m *DomainItemRequestBuilder) FederationConfigurationById(id string)(*ibe4d7e9a1d86de76f2fb6d99d1444ab26d987f71d4d6f5054eebc0aebb2d6687.InternalDomainFederationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["internalDomainFederation%2Did"] = id
    }
    return ibe4d7e9a1d86de76f2fb6d99d1444ab26d987f71d4d6f5054eebc0aebb2d6687.NewInternalDomainFederationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ForceDelete the forceDelete property
func (m *DomainItemRequestBuilder) ForceDelete()(*i1881cafcb69c2a03b070eab982c3f40a50f879626e4f5a433adf3d53764e612f.ForceDeleteRequestBuilder) {
    return i1881cafcb69c2a03b070eab982c3f40a50f879626e4f5a433adf3d53764e612f.NewForceDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get retrieve the properties and relationships of domain object.
func (m *DomainItemRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Domainable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler retrieve the properties and relationships of domain object.
func (m *DomainItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DomainItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Domainable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDomainFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Domainable), nil
}
// Patch update the properties of domain object.
func (m *DomainItemRequestBuilder) Patch(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Domainable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the properties of domain object.
func (m *DomainItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Domainable, requestConfiguration *DomainItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ServiceConfigurationRecords the serviceConfigurationRecords property
func (m *DomainItemRequestBuilder) ServiceConfigurationRecords()(*if6c517fa3faff0b607727add22bace8751787df881a7f48b830e8ee8854493e1.ServiceConfigurationRecordsRequestBuilder) {
    return if6c517fa3faff0b607727add22bace8751787df881a7f48b830e8ee8854493e1.NewServiceConfigurationRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServiceConfigurationRecordsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.domains.item.serviceConfigurationRecords.item collection
func (m *DomainItemRequestBuilder) ServiceConfigurationRecordsById(id string)(*if238474bb023ee1ae7b1e221337d6649adb30bb8a532c530686588c4d8f149e4.DomainDnsRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domainDnsRecord%2Did"] = id
    }
    return if238474bb023ee1ae7b1e221337d6649adb30bb8a532c530686588c4d8f149e4.NewDomainDnsRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// VerificationDnsRecords the verificationDnsRecords property
func (m *DomainItemRequestBuilder) VerificationDnsRecords()(*icf5b8e1a8b1c94340467eae996f377e8f09f9616367b7d5309238180fac55b36.VerificationDnsRecordsRequestBuilder) {
    return icf5b8e1a8b1c94340467eae996f377e8f09f9616367b7d5309238180fac55b36.NewVerificationDnsRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VerificationDnsRecordsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.domains.item.verificationDnsRecords.item collection
func (m *DomainItemRequestBuilder) VerificationDnsRecordsById(id string)(*i920c7da1a3bc0cc91d66e6345a870d4afe1c3ff7eadb815aed8f74ac233de19e.DomainDnsRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domainDnsRecord%2Did"] = id
    }
    return i920c7da1a3bc0cc91d66e6345a870d4afe1c3ff7eadb815aed8f74ac233de19e.NewDomainDnsRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Verify the verify property
func (m *DomainItemRequestBuilder) Verify()(*iacad0a6834b5884a8d686d9e187dfb93e979024e3471c842fc86db3dd765352a.VerifyRequestBuilder) {
    return iacad0a6834b5884a8d686d9e187dfb93e979024e3471c842fc86db3dd765352a.NewVerifyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
