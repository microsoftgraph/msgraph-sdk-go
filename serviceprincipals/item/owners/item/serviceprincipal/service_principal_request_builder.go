package serviceprincipal

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i22a64dd507be97787386a9ae23c4c7fa6a312df0dade9295bed96f76b124e586 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/getmemberobjects"
    i2a8eaa17f8da534bb025830839c2a5f91c76b1196a17850d0490f590d122650d "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/removepassword"
    i44800f594854e49d936f895cd330bacc8c254b5a0fdb831ed81d1b9c1b302935 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/checkmembergroups"
    i92167059256cfb735e694bac86ce2d78af12e2e2645897f22c0b2b6cffaaabcd "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/checkmemberobjects"
    i97245861c6bbe225b6764f74720f613338e50b7a4aa6bc99b54e9ed919d93426 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/addtokensigningcertificate"
    i9d3c59051cecc122dfb373b2996382e07abc2debbae844bee243192173092174 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/addkey"
    i9fa35bb9d3baffc4e88ac15c50c6e8f46662af2d454f9775f18a848c1818f104 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/addpassword"
    ia05a2b92508af75df3b1cd1c2729a8a70b75546e857acd00c1ce6b4acf100340 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/removekey"
    ib6891fb67758c9209766c71c12f57b1c7b8334d2bc8b17fae72872ae11311aac "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/restore"
    if3ae72be51cf76f6ab2d3f7494f79f2844b65bc0cda38c556c92c4ea8592367e "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/getmembergroups"
)

// ServicePrincipalRequestBuilder casts the previous resource to servicePrincipal.
type ServicePrincipalRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServicePrincipalRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
type ServicePrincipalRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServicePrincipalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalRequestBuilderGetQueryParameters
}
// AddKey the addKey property
func (m *ServicePrincipalRequestBuilder) AddKey()(*i9d3c59051cecc122dfb373b2996382e07abc2debbae844bee243192173092174.AddKeyRequestBuilder) {
    return i9d3c59051cecc122dfb373b2996382e07abc2debbae844bee243192173092174.NewAddKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddPassword the addPassword property
func (m *ServicePrincipalRequestBuilder) AddPassword()(*i9fa35bb9d3baffc4e88ac15c50c6e8f46662af2d454f9775f18a848c1818f104.AddPasswordRequestBuilder) {
    return i9fa35bb9d3baffc4e88ac15c50c6e8f46662af2d454f9775f18a848c1818f104.NewAddPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddTokenSigningCertificate the addTokenSigningCertificate property
func (m *ServicePrincipalRequestBuilder) AddTokenSigningCertificate()(*i97245861c6bbe225b6764f74720f613338e50b7a4aa6bc99b54e9ed919d93426.AddTokenSigningCertificateRequestBuilder) {
    return i97245861c6bbe225b6764f74720f613338e50b7a4aa6bc99b54e9ed919d93426.NewAddTokenSigningCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *ServicePrincipalRequestBuilder) CheckMemberGroups()(*i44800f594854e49d936f895cd330bacc8c254b5a0fdb831ed81d1b9c1b302935.CheckMemberGroupsRequestBuilder) {
    return i44800f594854e49d936f895cd330bacc8c254b5a0fdb831ed81d1b9c1b302935.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *ServicePrincipalRequestBuilder) CheckMemberObjects()(*i92167059256cfb735e694bac86ce2d78af12e2e2645897f22c0b2b6cffaaabcd.CheckMemberObjectsRequestBuilder) {
    return i92167059256cfb735e694bac86ce2d78af12e2e2645897f22c0b2b6cffaaabcd.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewServicePrincipalRequestBuilderInternal instantiates a new ServicePrincipalRequestBuilder and sets the default values.
func NewServicePrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalRequestBuilder) {
    m := &ServicePrincipalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/owners/{directoryObject%2Did}/microsoft.graph.servicePrincipal{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServicePrincipalRequestBuilder instantiates a new ServicePrincipalRequestBuilder and sets the default values.
func NewServicePrincipalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
func (m *ServicePrincipalRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
func (m *ServicePrincipalRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ServicePrincipalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
func (m *ServicePrincipalRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *ServicePrincipalRequestBuilder) GetMemberGroups()(*if3ae72be51cf76f6ab2d3f7494f79f2844b65bc0cda38c556c92c4ea8592367e.GetMemberGroupsRequestBuilder) {
    return if3ae72be51cf76f6ab2d3f7494f79f2844b65bc0cda38c556c92c4ea8592367e.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *ServicePrincipalRequestBuilder) GetMemberObjects()(*i22a64dd507be97787386a9ae23c4c7fa6a312df0dade9295bed96f76b124e586.GetMemberObjectsRequestBuilder) {
    return i22a64dd507be97787386a9ae23c4c7fa6a312df0dade9295bed96f76b124e586.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
func (m *ServicePrincipalRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ServicePrincipalRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServicePrincipalFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable), nil
}
// RemoveKey the removeKey property
func (m *ServicePrincipalRequestBuilder) RemoveKey()(*ia05a2b92508af75df3b1cd1c2729a8a70b75546e857acd00c1ce6b4acf100340.RemoveKeyRequestBuilder) {
    return ia05a2b92508af75df3b1cd1c2729a8a70b75546e857acd00c1ce6b4acf100340.NewRemoveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemovePassword the removePassword property
func (m *ServicePrincipalRequestBuilder) RemovePassword()(*i2a8eaa17f8da534bb025830839c2a5f91c76b1196a17850d0490f590d122650d.RemovePasswordRequestBuilder) {
    return i2a8eaa17f8da534bb025830839c2a5f91c76b1196a17850d0490f590d122650d.NewRemovePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *ServicePrincipalRequestBuilder) Restore()(*ib6891fb67758c9209766c71c12f57b1c7b8334d2bc8b17fae72872ae11311aac.RestoreRequestBuilder) {
    return ib6891fb67758c9209766c71c12f57b1c7b8334d2bc8b17fae72872ae11311aac.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
