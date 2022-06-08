package serviceprincipal

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i051fcba8f706110628d7b985e32bf36a64509dcb9d160a902b80fa96119576ca "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/addtokensigningcertificate"
    i1c4da1b8595db0c7425870693356f5ba373174aecd92c971c559a91845af4158 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/restore"
    i265432709fa2facb5781efbd2843370b05d85afc47c4234ac1cb9b907e1ece49 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/addpassword"
    i79883eb42f12a2496a09e14ae3c3b366cac020f31c4e3f340bddba6ebf98f9ce "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/getmemberobjects"
    ia4701c0e48a018a6696158da4a9b95454966ec62f214de0e7c2691997dcb3d9f "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/removepassword"
    ic024e499513c513eea70d37a163ad6c797d9d5ee6a356982db00136ae8538a93 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/getmembergroups"
    ic6010d5e3780676c63a084b36d8bedb17d49dc4d525bedfeadd100c02819c462 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/checkmemberobjects"
    idaf41e902dbaf8a9c5a9390794ed288596ffb968010505e8665163b2af61aa8c "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/checkmembergroups"
    idf7e1ac94ba6c10b232840c83d1ba2f5b17b7f9582ecf9da6e2beb3224f1e839 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/addkey"
    ie8e65116caa1e6fc066ab23f9c71c42587406f77a50ba85f8f2bf44690e3f529 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/removekey"
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
func (m *ServicePrincipalRequestBuilder) AddKey()(*idf7e1ac94ba6c10b232840c83d1ba2f5b17b7f9582ecf9da6e2beb3224f1e839.AddKeyRequestBuilder) {
    return idf7e1ac94ba6c10b232840c83d1ba2f5b17b7f9582ecf9da6e2beb3224f1e839.NewAddKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddPassword the addPassword property
func (m *ServicePrincipalRequestBuilder) AddPassword()(*i265432709fa2facb5781efbd2843370b05d85afc47c4234ac1cb9b907e1ece49.AddPasswordRequestBuilder) {
    return i265432709fa2facb5781efbd2843370b05d85afc47c4234ac1cb9b907e1ece49.NewAddPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddTokenSigningCertificate the addTokenSigningCertificate property
func (m *ServicePrincipalRequestBuilder) AddTokenSigningCertificate()(*i051fcba8f706110628d7b985e32bf36a64509dcb9d160a902b80fa96119576ca.AddTokenSigningCertificateRequestBuilder) {
    return i051fcba8f706110628d7b985e32bf36a64509dcb9d160a902b80fa96119576ca.NewAddTokenSigningCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *ServicePrincipalRequestBuilder) CheckMemberGroups()(*idaf41e902dbaf8a9c5a9390794ed288596ffb968010505e8665163b2af61aa8c.CheckMemberGroupsRequestBuilder) {
    return idaf41e902dbaf8a9c5a9390794ed288596ffb968010505e8665163b2af61aa8c.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *ServicePrincipalRequestBuilder) CheckMemberObjects()(*ic6010d5e3780676c63a084b36d8bedb17d49dc4d525bedfeadd100c02819c462.CheckMemberObjectsRequestBuilder) {
    return ic6010d5e3780676c63a084b36d8bedb17d49dc4d525bedfeadd100c02819c462.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewServicePrincipalRequestBuilderInternal instantiates a new ServicePrincipalRequestBuilder and sets the default values.
func NewServicePrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalRequestBuilder) {
    m := &ServicePrincipalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/ownedObjects/{directoryObject%2Did}/microsoft.graph.servicePrincipal{?%24select,%24expand}";
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
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
func (m *ServicePrincipalRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *ServicePrincipalRequestBuilder) GetMemberGroups()(*ic024e499513c513eea70d37a163ad6c797d9d5ee6a356982db00136ae8538a93.GetMemberGroupsRequestBuilder) {
    return ic024e499513c513eea70d37a163ad6c797d9d5ee6a356982db00136ae8538a93.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *ServicePrincipalRequestBuilder) GetMemberObjects()(*i79883eb42f12a2496a09e14ae3c3b366cac020f31c4e3f340bddba6ebf98f9ce.GetMemberObjectsRequestBuilder) {
    return i79883eb42f12a2496a09e14ae3c3b366cac020f31c4e3f340bddba6ebf98f9ce.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ServicePrincipalRequestBuilder) RemoveKey()(*ie8e65116caa1e6fc066ab23f9c71c42587406f77a50ba85f8f2bf44690e3f529.RemoveKeyRequestBuilder) {
    return ie8e65116caa1e6fc066ab23f9c71c42587406f77a50ba85f8f2bf44690e3f529.NewRemoveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemovePassword the removePassword property
func (m *ServicePrincipalRequestBuilder) RemovePassword()(*ia4701c0e48a018a6696158da4a9b95454966ec62f214de0e7c2691997dcb3d9f.RemovePasswordRequestBuilder) {
    return ia4701c0e48a018a6696158da4a9b95454966ec62f214de0e7c2691997dcb3d9f.NewRemovePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *ServicePrincipalRequestBuilder) Restore()(*i1c4da1b8595db0c7425870693356f5ba373174aecd92c971c559a91845af4158.RestoreRequestBuilder) {
    return i1c4da1b8595db0c7425870693356f5ba373174aecd92c971c559a91845af4158.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
