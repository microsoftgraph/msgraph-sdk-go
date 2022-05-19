package serviceprincipal

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i075f221fba45a757d37d33b29b5508ab1b3ec78e9280f7854fc3fba7369e50a1 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/removekey"
    i0df47f80429ed175c691d1d6efac0bb7e73b04638dc2c65e0bb031950aa0efe2 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/removepassword"
    i1fbf3128759790a7042e80ed5f593310d3e051edf51767d16dd84bce5f247728 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/addtokensigningcertificate"
    i22d890fd9d70f1cc626c58a43b72a284731df5338d1cf0e73a3f7f8f63257bb0 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/getmemberobjects"
    i639298ac58827770954176a42d3d76e11939ffed8656328ac36dd0f454c18514 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/addpassword"
    i7eaeb6247f9cd4df0f44ea56b037dc3059012141b3053cb86ebc28e1746baa57 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/restore"
    ie74bfee78c38cc003d740457379f20d0d936b05843d03bd9372de4d2cb5524ac "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/checkmemberobjects"
    iec2e95e157eaa85c98204b031d84df60fcac1d9abc08ced49b1344654f4e7b7a "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/getmembergroups"
    if7fa1dafa8effc9cce2c26d88b1ef5aaf341f1a2363329f7d6d52ff757f38ce9 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/addkey"
    ifc745ca605b8fb697bc3068e11a78661264ce8602500109a87680a863961a896 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/checkmembergroups"
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
func (m *ServicePrincipalRequestBuilder) AddKey()(*if7fa1dafa8effc9cce2c26d88b1ef5aaf341f1a2363329f7d6d52ff757f38ce9.AddKeyRequestBuilder) {
    return if7fa1dafa8effc9cce2c26d88b1ef5aaf341f1a2363329f7d6d52ff757f38ce9.NewAddKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddPassword the addPassword property
func (m *ServicePrincipalRequestBuilder) AddPassword()(*i639298ac58827770954176a42d3d76e11939ffed8656328ac36dd0f454c18514.AddPasswordRequestBuilder) {
    return i639298ac58827770954176a42d3d76e11939ffed8656328ac36dd0f454c18514.NewAddPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddTokenSigningCertificate the addTokenSigningCertificate property
func (m *ServicePrincipalRequestBuilder) AddTokenSigningCertificate()(*i1fbf3128759790a7042e80ed5f593310d3e051edf51767d16dd84bce5f247728.AddTokenSigningCertificateRequestBuilder) {
    return i1fbf3128759790a7042e80ed5f593310d3e051edf51767d16dd84bce5f247728.NewAddTokenSigningCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *ServicePrincipalRequestBuilder) CheckMemberGroups()(*ifc745ca605b8fb697bc3068e11a78661264ce8602500109a87680a863961a896.CheckMemberGroupsRequestBuilder) {
    return ifc745ca605b8fb697bc3068e11a78661264ce8602500109a87680a863961a896.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *ServicePrincipalRequestBuilder) CheckMemberObjects()(*ie74bfee78c38cc003d740457379f20d0d936b05843d03bd9372de4d2cb5524ac.CheckMemberObjectsRequestBuilder) {
    return ie74bfee78c38cc003d740457379f20d0d936b05843d03bd9372de4d2cb5524ac.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewServicePrincipalRequestBuilderInternal instantiates a new ServicePrincipalRequestBuilder and sets the default values.
func NewServicePrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalRequestBuilder) {
    m := &ServicePrincipalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/memberOf/{directoryObject%2Did}/microsoft.graph.servicePrincipal{?%24select,%24expand}";
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
func (m *ServicePrincipalRequestBuilder) GetMemberGroups()(*iec2e95e157eaa85c98204b031d84df60fcac1d9abc08ced49b1344654f4e7b7a.GetMemberGroupsRequestBuilder) {
    return iec2e95e157eaa85c98204b031d84df60fcac1d9abc08ced49b1344654f4e7b7a.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *ServicePrincipalRequestBuilder) GetMemberObjects()(*i22d890fd9d70f1cc626c58a43b72a284731df5338d1cf0e73a3f7f8f63257bb0.GetMemberObjectsRequestBuilder) {
    return i22d890fd9d70f1cc626c58a43b72a284731df5338d1cf0e73a3f7f8f63257bb0.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ServicePrincipalRequestBuilder) RemoveKey()(*i075f221fba45a757d37d33b29b5508ab1b3ec78e9280f7854fc3fba7369e50a1.RemoveKeyRequestBuilder) {
    return i075f221fba45a757d37d33b29b5508ab1b3ec78e9280f7854fc3fba7369e50a1.NewRemoveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemovePassword the removePassword property
func (m *ServicePrincipalRequestBuilder) RemovePassword()(*i0df47f80429ed175c691d1d6efac0bb7e73b04638dc2c65e0bb031950aa0efe2.RemovePasswordRequestBuilder) {
    return i0df47f80429ed175c691d1d6efac0bb7e73b04638dc2c65e0bb031950aa0efe2.NewRemovePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *ServicePrincipalRequestBuilder) Restore()(*i7eaeb6247f9cd4df0f44ea56b037dc3059012141b3053cb86ebc28e1746baa57.RestoreRequestBuilder) {
    return i7eaeb6247f9cd4df0f44ea56b037dc3059012141b3053cb86ebc28e1746baa57.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
