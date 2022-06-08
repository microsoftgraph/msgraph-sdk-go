package serviceprincipal

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i25852b515437e2773bfe8d0911e3481aaac04387aabd557c441288be197318fc "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/removepassword"
    i40b5b39e4f061c0147f9e71785113865789288b867a3857c246cecd335bd9ff5 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/addpassword"
    i48df8b89582e5db9044a5f3682def810456c066ae780f703cf726cdfd091bdd1 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/checkmemberobjects"
    i6beb7573c30ef065625ea3604d8287c8e5ef3ef80b925612d8c866eec86fcaea "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/removekey"
    ib5563dee452d5777eccdb54d4cb5953b2ef6103d41971f81b45d97ec940a8962 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/getmembergroups"
    ic947284e38c7a58bf8683b05b5898a948c746719cd61ffaee38ed97a23779780 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/addkey"
    idd0e4416488b88d71f46d014fe6574195767e988fa98beb22934fe3e0b955ace "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/restore"
    iec8ed9f214b7766db3e97a28a8dadae4268411c28a8937cf2776110204f30a0b "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/getmemberobjects"
    ied26e6322c7a99c1f7634cebab31124be5f60e58ec700c058b0bcdca5f38ee37 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/addtokensigningcertificate"
    ifc6bca81c297182a1af544e148f9bd3275cffbd38388919e8d5446dd942c53bc "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/checkmembergroups"
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
func (m *ServicePrincipalRequestBuilder) AddKey()(*ic947284e38c7a58bf8683b05b5898a948c746719cd61ffaee38ed97a23779780.AddKeyRequestBuilder) {
    return ic947284e38c7a58bf8683b05b5898a948c746719cd61ffaee38ed97a23779780.NewAddKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddPassword the addPassword property
func (m *ServicePrincipalRequestBuilder) AddPassword()(*i40b5b39e4f061c0147f9e71785113865789288b867a3857c246cecd335bd9ff5.AddPasswordRequestBuilder) {
    return i40b5b39e4f061c0147f9e71785113865789288b867a3857c246cecd335bd9ff5.NewAddPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddTokenSigningCertificate the addTokenSigningCertificate property
func (m *ServicePrincipalRequestBuilder) AddTokenSigningCertificate()(*ied26e6322c7a99c1f7634cebab31124be5f60e58ec700c058b0bcdca5f38ee37.AddTokenSigningCertificateRequestBuilder) {
    return ied26e6322c7a99c1f7634cebab31124be5f60e58ec700c058b0bcdca5f38ee37.NewAddTokenSigningCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *ServicePrincipalRequestBuilder) CheckMemberGroups()(*ifc6bca81c297182a1af544e148f9bd3275cffbd38388919e8d5446dd942c53bc.CheckMemberGroupsRequestBuilder) {
    return ifc6bca81c297182a1af544e148f9bd3275cffbd38388919e8d5446dd942c53bc.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *ServicePrincipalRequestBuilder) CheckMemberObjects()(*i48df8b89582e5db9044a5f3682def810456c066ae780f703cf726cdfd091bdd1.CheckMemberObjectsRequestBuilder) {
    return i48df8b89582e5db9044a5f3682def810456c066ae780f703cf726cdfd091bdd1.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewServicePrincipalRequestBuilderInternal instantiates a new ServicePrincipalRequestBuilder and sets the default values.
func NewServicePrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalRequestBuilder) {
    m := &ServicePrincipalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/createdObjects/{directoryObject%2Did}/microsoft.graph.servicePrincipal{?%24select,%24expand}";
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
func (m *ServicePrincipalRequestBuilder) GetMemberGroups()(*ib5563dee452d5777eccdb54d4cb5953b2ef6103d41971f81b45d97ec940a8962.GetMemberGroupsRequestBuilder) {
    return ib5563dee452d5777eccdb54d4cb5953b2ef6103d41971f81b45d97ec940a8962.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *ServicePrincipalRequestBuilder) GetMemberObjects()(*iec8ed9f214b7766db3e97a28a8dadae4268411c28a8937cf2776110204f30a0b.GetMemberObjectsRequestBuilder) {
    return iec8ed9f214b7766db3e97a28a8dadae4268411c28a8937cf2776110204f30a0b.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ServicePrincipalRequestBuilder) RemoveKey()(*i6beb7573c30ef065625ea3604d8287c8e5ef3ef80b925612d8c866eec86fcaea.RemoveKeyRequestBuilder) {
    return i6beb7573c30ef065625ea3604d8287c8e5ef3ef80b925612d8c866eec86fcaea.NewRemoveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemovePassword the removePassword property
func (m *ServicePrincipalRequestBuilder) RemovePassword()(*i25852b515437e2773bfe8d0911e3481aaac04387aabd557c441288be197318fc.RemovePasswordRequestBuilder) {
    return i25852b515437e2773bfe8d0911e3481aaac04387aabd557c441288be197318fc.NewRemovePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *ServicePrincipalRequestBuilder) Restore()(*idd0e4416488b88d71f46d014fe6574195767e988fa98beb22934fe3e0b955ace.RestoreRequestBuilder) {
    return idd0e4416488b88d71f46d014fe6574195767e988fa98beb22934fe3e0b955ace.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
