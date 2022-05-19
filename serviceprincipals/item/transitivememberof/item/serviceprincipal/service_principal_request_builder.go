package serviceprincipal

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i15214a5ceb9022e6340299f8f7777f9086fcbdf5bbc8fc86a0b98a13ec225168 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/checkmembergroups"
    i4ae16959c30e9e5fba2564f005675b0cc5cc8093380e742fbc95692543e0e481 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/getmemberobjects"
    i525c87ebb0b68d3501dc35f292d23babba0b4ff1acb4ec099988bd4ca2cb1f20 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/addkey"
    i79ddfa1559c27c72f9be4b9a19f01567f057f587dd4e2c2f55eee35674676e0c "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/removepassword"
    i8b7bae45514c483d2b3fc1b8036324aa48db36ebe939765cecf1c99172d4a700 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/getmembergroups"
    ia1ea009ec4bb953cf8410dfd445f78498202a2ce6d5f62a3e6e4f726415880d1 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/restore"
    ic81bad984fa75cd551b10ad7826907af2f3f7d543d0cbc975e7c979bbd2f7651 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/addtokensigningcertificate"
    ie6471bbfddf24835a2019e3b814e2641fa6a3c80e97935d3cf0842583280c8ce "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/checkmemberobjects"
    iee0f70de7f17254e3db854472526a27998af03f529cf353d08930a52bb15fb96 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/addpassword"
    iffca471eeeb8191c83c7370cfa7f1a9e9e59879978cb2cc7ab9e496a31e531b4 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/removekey"
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
func (m *ServicePrincipalRequestBuilder) AddKey()(*i525c87ebb0b68d3501dc35f292d23babba0b4ff1acb4ec099988bd4ca2cb1f20.AddKeyRequestBuilder) {
    return i525c87ebb0b68d3501dc35f292d23babba0b4ff1acb4ec099988bd4ca2cb1f20.NewAddKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddPassword the addPassword property
func (m *ServicePrincipalRequestBuilder) AddPassword()(*iee0f70de7f17254e3db854472526a27998af03f529cf353d08930a52bb15fb96.AddPasswordRequestBuilder) {
    return iee0f70de7f17254e3db854472526a27998af03f529cf353d08930a52bb15fb96.NewAddPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddTokenSigningCertificate the addTokenSigningCertificate property
func (m *ServicePrincipalRequestBuilder) AddTokenSigningCertificate()(*ic81bad984fa75cd551b10ad7826907af2f3f7d543d0cbc975e7c979bbd2f7651.AddTokenSigningCertificateRequestBuilder) {
    return ic81bad984fa75cd551b10ad7826907af2f3f7d543d0cbc975e7c979bbd2f7651.NewAddTokenSigningCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *ServicePrincipalRequestBuilder) CheckMemberGroups()(*i15214a5ceb9022e6340299f8f7777f9086fcbdf5bbc8fc86a0b98a13ec225168.CheckMemberGroupsRequestBuilder) {
    return i15214a5ceb9022e6340299f8f7777f9086fcbdf5bbc8fc86a0b98a13ec225168.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *ServicePrincipalRequestBuilder) CheckMemberObjects()(*ie6471bbfddf24835a2019e3b814e2641fa6a3c80e97935d3cf0842583280c8ce.CheckMemberObjectsRequestBuilder) {
    return ie6471bbfddf24835a2019e3b814e2641fa6a3c80e97935d3cf0842583280c8ce.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewServicePrincipalRequestBuilderInternal instantiates a new ServicePrincipalRequestBuilder and sets the default values.
func NewServicePrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalRequestBuilder) {
    m := &ServicePrincipalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.servicePrincipal{?%24select,%24expand}";
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
func (m *ServicePrincipalRequestBuilder) GetMemberGroups()(*i8b7bae45514c483d2b3fc1b8036324aa48db36ebe939765cecf1c99172d4a700.GetMemberGroupsRequestBuilder) {
    return i8b7bae45514c483d2b3fc1b8036324aa48db36ebe939765cecf1c99172d4a700.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *ServicePrincipalRequestBuilder) GetMemberObjects()(*i4ae16959c30e9e5fba2564f005675b0cc5cc8093380e742fbc95692543e0e481.GetMemberObjectsRequestBuilder) {
    return i4ae16959c30e9e5fba2564f005675b0cc5cc8093380e742fbc95692543e0e481.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ServicePrincipalRequestBuilder) RemoveKey()(*iffca471eeeb8191c83c7370cfa7f1a9e9e59879978cb2cc7ab9e496a31e531b4.RemoveKeyRequestBuilder) {
    return iffca471eeeb8191c83c7370cfa7f1a9e9e59879978cb2cc7ab9e496a31e531b4.NewRemoveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemovePassword the removePassword property
func (m *ServicePrincipalRequestBuilder) RemovePassword()(*i79ddfa1559c27c72f9be4b9a19f01567f057f587dd4e2c2f55eee35674676e0c.RemovePasswordRequestBuilder) {
    return i79ddfa1559c27c72f9be4b9a19f01567f057f587dd4e2c2f55eee35674676e0c.NewRemovePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *ServicePrincipalRequestBuilder) Restore()(*ia1ea009ec4bb953cf8410dfd445f78498202a2ce6d5f62a3e6e4f726415880d1.RestoreRequestBuilder) {
    return ia1ea009ec4bb953cf8410dfd445f78498202a2ce6d5f62a3e6e4f726415880d1.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
