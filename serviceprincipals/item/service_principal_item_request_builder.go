package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i003add1c72d92bd47a9ee37d3f334a2b1149222e48ce785c0b07d0542f89302f "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/removepassword"
    i077b2e62603ccb12d8cc7601297f5d1a834bbe55d05e304176e9f9ad509fc426 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/checkmembergroups"
    i1f042d332a3da52b9a4d53ae9423654b32107cbe9b637d9ca95d07e583140364 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/createdobjects"
    i2c910a3074da68d087d5c6fc96ab8be3cf4e54ef4ec7985acd4e132e52d42137 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/getmemberobjects"
    i2ebeefd4c13ab730e7f8f02313f1b93261a541ed22898061b1c09f058126615a "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/tokenissuancepolicies"
    i3c2ed9adbc44109667bcb7b30ffd0c6b2bb97d62b8e12c57b863c71af15e6903 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/removekey"
    i3f559eeb0a32c852d52064d33f878279ea63214977aec7d5605657ac813f843e "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/addkey"
    i495446d97cfd23af0607d58fb1c29e62eeec78e42cbd747f401374edf193b049 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/addpassword"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i60f4aa5edd1e35c0b5a11cc9ce9afe3a692e2aef2b419b5d6bb6c9bd8c0c2561 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/transitivememberof"
    i6265425e2a2080afde3979740ad79538958bd2d2ab7ce003f7d0a624b06989bc "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/checkmemberobjects"
    i93153ff64f35a046ae51130f6dbc11813af572bf8610f8142e2bebbb77269190 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/approleassignments"
    ia161e203c4876e46996f4ce25080c383320b3f90c23f41a4644fabc41a122c2c "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/owners"
    ia329fbf2e4e7682fe90bdfc6bf4cc749c189487202d60eefa0fdec3b62b0b7c3 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/ownedobjects"
    ia47ff52d5020404c6eaea0bc5dbd4bbafaa1e0c73dbc6d4dcfee263b2245b53e "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/tokenlifetimepolicies"
    iaa2be345892fc1350e44205c5f24bef008513f89d2497eabd4c16620df674828 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/getmembergroups"
    ib8feba2d910c8485220ccc6a35f2363ef36da3466e5a53c33911b796b751a405 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/addtokensigningcertificate"
    ic2b775c554f6f0689bed0deb6e610608281614f70275cada61da8274360068d0 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/endpoints"
    icc1091ecf66b1de7080efaeb2e59cdf1074f3b116fba7066b71a0d6cd9c57a53 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/claimsmappingpolicies"
    id692483c2e6fa4444b9d250ecaa7a6bea16dbc3cd7fa20c05fb69b2f7b69ad46 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/approleassignedto"
    ida4a28e902893200b019f06f91d4d3d9d5d69f185e4f2b9b706d4abd32593768 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/memberof"
    idbd9734414b90f65880420caeaf7c9aad49278527e828f8795e9b8bc0c65d8a5 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/oauth2permissiongrants"
    ie49af58d75366149082d7cc96496e031433add00a7bde619d68ae8a646ef55e2 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/homerealmdiscoverypolicies"
    iedd8dd60eed89931552d2c873e7c568c2a30c87dee999cc992446db26f8803fc "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/delegatedpermissionclassifications"
    if93b2c8d7b19d0341ed3f43075bc082dd6d4c57bb7cafaa0f2161fce000d6f1d "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/restore"
    i14cf3ba6d0e45059a525d05353e514742297a11654926e06d365c73081a6a0dd "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/tokenlifetimepolicies/item"
    i2e9c412d7776dbbef6b295067f85a1c21e3fd42b4d7846e1174dfd55eb8d8811 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/createdobjects/item"
    i38c5edd99af7364922d88fc2b49de45f320d11c5fa36e8234c12f1d863cbea73 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/delegatedpermissionclassifications/item"
    i4c7a31e042d98d59ede758ea63655f65992075acf7610c7db5c90df699c7c9e2 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/tokenissuancepolicies/item"
    i56e2c3bd7faf41b227bd7eb65043ee578bb4d1dc5df7c395a2d64d2ed4af7fce "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/memberof/item"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i80395ca3a6452268f5c4d3e92f315d7c92e78fef9911bf32e289ea47d6968415 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/approleassignments/item"
    i83568fd90c133e34217ad0e8f925c320323fb270d334badf13c19b1a8ac14daa "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/oauth2permissiongrants/item"
    i9fe3c2224826dcb0cde6e1f06bdfb308132901afdafc77fa24614124022a183b "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/ownedobjects/item"
    iae6f1167638ef01ba227e18af033b58e9203edca74b4641453bc1cec1de3aa91 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/endpoints/item"
    ib95837107c1a1d541878efdf1d0902a7c67b80de5fc97d771d55c403ab043621 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/transitivememberof/item"
    id569966c2c6119078f6226e2ed4c7c42d8bd056f52f7651bd3fafcc36ffb7d0c "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/claimsmappingpolicies/item"
    ie10051a8bcf74fee28b7346347c913cf4c35b13dc5337d6b1877c19ead0529ae "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/owners/item"
    ie33736e058581033dd3335478f67d4920269f0937f554299dea9b4a420f14037 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/homerealmdiscoverypolicies/item"
    ieade744732b95d81e11bac4d5146b5ac0bfeb407228f97be47815b5b44c77ff1 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/approleassignedto/item"
)

// ServicePrincipalItemRequestBuilder provides operations to manage the collection of servicePrincipal entities.
type ServicePrincipalItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ServicePrincipalItemRequestBuilderDeleteOptions options for Delete
type ServicePrincipalItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ServicePrincipalItemRequestBuilderGetOptions options for Get
type ServicePrincipalItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ServicePrincipalItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ServicePrincipalItemRequestBuilderGetQueryParameters get entity from servicePrincipals by key
type ServicePrincipalItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ServicePrincipalItemRequestBuilderPatchOptions options for Patch
type ServicePrincipalItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServicePrincipalable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ServicePrincipalItemRequestBuilder) AddKey()(*i3f559eeb0a32c852d52064d33f878279ea63214977aec7d5605657ac813f843e.AddKeyRequestBuilder) {
    return i3f559eeb0a32c852d52064d33f878279ea63214977aec7d5605657ac813f843e.NewAddKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) AddPassword()(*i495446d97cfd23af0607d58fb1c29e62eeec78e42cbd747f401374edf193b049.AddPasswordRequestBuilder) {
    return i495446d97cfd23af0607d58fb1c29e62eeec78e42cbd747f401374edf193b049.NewAddPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) AddTokenSigningCertificate()(*ib8feba2d910c8485220ccc6a35f2363ef36da3466e5a53c33911b796b751a405.AddTokenSigningCertificateRequestBuilder) {
    return ib8feba2d910c8485220ccc6a35f2363ef36da3466e5a53c33911b796b751a405.NewAddTokenSigningCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) AppRoleAssignedTo()(*id692483c2e6fa4444b9d250ecaa7a6bea16dbc3cd7fa20c05fb69b2f7b69ad46.AppRoleAssignedToRequestBuilder) {
    return id692483c2e6fa4444b9d250ecaa7a6bea16dbc3cd7fa20c05fb69b2f7b69ad46.NewAppRoleAssignedToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignedToById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.servicePrincipals.item.appRoleAssignedTo.item collection
func (m *ServicePrincipalItemRequestBuilder) AppRoleAssignedToById(id string)(*ieade744732b95d81e11bac4d5146b5ac0bfeb407228f97be47815b5b44c77ff1.AppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment_id"] = id
    }
    return ieade744732b95d81e11bac4d5146b5ac0bfeb407228f97be47815b5b44c77ff1.NewAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) AppRoleAssignments()(*i93153ff64f35a046ae51130f6dbc11813af572bf8610f8142e2bebbb77269190.AppRoleAssignmentsRequestBuilder) {
    return i93153ff64f35a046ae51130f6dbc11813af572bf8610f8142e2bebbb77269190.NewAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.servicePrincipals.item.appRoleAssignments.item collection
func (m *ServicePrincipalItemRequestBuilder) AppRoleAssignmentsById(id string)(*i80395ca3a6452268f5c4d3e92f315d7c92e78fef9911bf32e289ea47d6968415.AppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment_id"] = id
    }
    return i80395ca3a6452268f5c4d3e92f315d7c92e78fef9911bf32e289ea47d6968415.NewAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) CheckMemberGroups()(*i077b2e62603ccb12d8cc7601297f5d1a834bbe55d05e304176e9f9ad509fc426.CheckMemberGroupsRequestBuilder) {
    return i077b2e62603ccb12d8cc7601297f5d1a834bbe55d05e304176e9f9ad509fc426.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) CheckMemberObjects()(*i6265425e2a2080afde3979740ad79538958bd2d2ab7ce003f7d0a624b06989bc.CheckMemberObjectsRequestBuilder) {
    return i6265425e2a2080afde3979740ad79538958bd2d2ab7ce003f7d0a624b06989bc.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) ClaimsMappingPolicies()(*icc1091ecf66b1de7080efaeb2e59cdf1074f3b116fba7066b71a0d6cd9c57a53.ClaimsMappingPoliciesRequestBuilder) {
    return icc1091ecf66b1de7080efaeb2e59cdf1074f3b116fba7066b71a0d6cd9c57a53.NewClaimsMappingPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClaimsMappingPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.servicePrincipals.item.claimsMappingPolicies.item collection
func (m *ServicePrincipalItemRequestBuilder) ClaimsMappingPoliciesById(id string)(*id569966c2c6119078f6226e2ed4c7c42d8bd056f52f7651bd3fafcc36ffb7d0c.ClaimsMappingPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["claimsMappingPolicy_id"] = id
    }
    return id569966c2c6119078f6226e2ed4c7c42d8bd056f52f7651bd3fafcc36ffb7d0c.NewClaimsMappingPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewServicePrincipalItemRequestBuilderInternal instantiates a new ServicePrincipalItemRequestBuilder and sets the default values.
func NewServicePrincipalItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServicePrincipalItemRequestBuilder) {
    m := &ServicePrincipalItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServicePrincipalItemRequestBuilder instantiates a new ServicePrincipalItemRequestBuilder and sets the default values.
func NewServicePrincipalItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServicePrincipalItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from servicePrincipals
func (m *ServicePrincipalItemRequestBuilder) CreateDeleteRequestInformation(options *ServicePrincipalItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ServicePrincipalItemRequestBuilder) CreatedObjects()(*i1f042d332a3da52b9a4d53ae9423654b32107cbe9b637d9ca95d07e583140364.CreatedObjectsRequestBuilder) {
    return i1f042d332a3da52b9a4d53ae9423654b32107cbe9b637d9ca95d07e583140364.NewCreatedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatedObjectsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.servicePrincipals.item.createdObjects.item collection
func (m *ServicePrincipalItemRequestBuilder) CreatedObjectsById(id string)(*i2e9c412d7776dbbef6b295067f85a1c21e3fd42b4d7846e1174dfd55eb8d8811.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i2e9c412d7776dbbef6b295067f85a1c21e3fd42b4d7846e1174dfd55eb8d8811.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateGetRequestInformation get entity from servicePrincipals by key
func (m *ServicePrincipalItemRequestBuilder) CreateGetRequestInformation(options *ServicePrincipalItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in servicePrincipals
func (m *ServicePrincipalItemRequestBuilder) CreatePatchRequestInformation(options *ServicePrincipalItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ServicePrincipalItemRequestBuilder) DelegatedPermissionClassifications()(*iedd8dd60eed89931552d2c873e7c568c2a30c87dee999cc992446db26f8803fc.DelegatedPermissionClassificationsRequestBuilder) {
    return iedd8dd60eed89931552d2c873e7c568c2a30c87dee999cc992446db26f8803fc.NewDelegatedPermissionClassificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DelegatedPermissionClassificationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.servicePrincipals.item.delegatedPermissionClassifications.item collection
func (m *ServicePrincipalItemRequestBuilder) DelegatedPermissionClassificationsById(id string)(*i38c5edd99af7364922d88fc2b49de45f320d11c5fa36e8234c12f1d863cbea73.DelegatedPermissionClassificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedPermissionClassification_id"] = id
    }
    return i38c5edd99af7364922d88fc2b49de45f320d11c5fa36e8234c12f1d863cbea73.NewDelegatedPermissionClassificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete entity from servicePrincipals
func (m *ServicePrincipalItemRequestBuilder) Delete(options *ServicePrincipalItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ServicePrincipalItemRequestBuilder) Endpoints()(*ic2b775c554f6f0689bed0deb6e610608281614f70275cada61da8274360068d0.EndpointsRequestBuilder) {
    return ic2b775c554f6f0689bed0deb6e610608281614f70275cada61da8274360068d0.NewEndpointsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EndpointsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.servicePrincipals.item.endpoints.item collection
func (m *ServicePrincipalItemRequestBuilder) EndpointsById(id string)(*iae6f1167638ef01ba227e18af033b58e9203edca74b4641453bc1cec1de3aa91.EndpointItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["endpoint_id"] = id
    }
    return iae6f1167638ef01ba227e18af033b58e9203edca74b4641453bc1cec1de3aa91.NewEndpointItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from servicePrincipals by key
func (m *ServicePrincipalItemRequestBuilder) Get(options *ServicePrincipalItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServicePrincipalable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateServicePrincipalFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServicePrincipalable), nil
}
func (m *ServicePrincipalItemRequestBuilder) GetMemberGroups()(*iaa2be345892fc1350e44205c5f24bef008513f89d2497eabd4c16620df674828.GetMemberGroupsRequestBuilder) {
    return iaa2be345892fc1350e44205c5f24bef008513f89d2497eabd4c16620df674828.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) GetMemberObjects()(*i2c910a3074da68d087d5c6fc96ab8be3cf4e54ef4ec7985acd4e132e52d42137.GetMemberObjectsRequestBuilder) {
    return i2c910a3074da68d087d5c6fc96ab8be3cf4e54ef4ec7985acd4e132e52d42137.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) HomeRealmDiscoveryPolicies()(*ie49af58d75366149082d7cc96496e031433add00a7bde619d68ae8a646ef55e2.HomeRealmDiscoveryPoliciesRequestBuilder) {
    return ie49af58d75366149082d7cc96496e031433add00a7bde619d68ae8a646ef55e2.NewHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HomeRealmDiscoveryPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.servicePrincipals.item.homeRealmDiscoveryPolicies.item collection
func (m *ServicePrincipalItemRequestBuilder) HomeRealmDiscoveryPoliciesById(id string)(*ie33736e058581033dd3335478f67d4920269f0937f554299dea9b4a420f14037.HomeRealmDiscoveryPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["homeRealmDiscoveryPolicy_id"] = id
    }
    return ie33736e058581033dd3335478f67d4920269f0937f554299dea9b4a420f14037.NewHomeRealmDiscoveryPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) MemberOf()(*ida4a28e902893200b019f06f91d4d3d9d5d69f185e4f2b9b706d4abd32593768.MemberOfRequestBuilder) {
    return ida4a28e902893200b019f06f91d4d3d9d5d69f185e4f2b9b706d4abd32593768.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.servicePrincipals.item.memberOf.item collection
func (m *ServicePrincipalItemRequestBuilder) MemberOfById(id string)(*i56e2c3bd7faf41b227bd7eb65043ee578bb4d1dc5df7c395a2d64d2ed4af7fce.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i56e2c3bd7faf41b227bd7eb65043ee578bb4d1dc5df7c395a2d64d2ed4af7fce.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) Oauth2PermissionGrants()(*idbd9734414b90f65880420caeaf7c9aad49278527e828f8795e9b8bc0c65d8a5.Oauth2PermissionGrantsRequestBuilder) {
    return idbd9734414b90f65880420caeaf7c9aad49278527e828f8795e9b8bc0c65d8a5.NewOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Oauth2PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.servicePrincipals.item.oauth2PermissionGrants.item collection
func (m *ServicePrincipalItemRequestBuilder) Oauth2PermissionGrantsById(id string)(*i83568fd90c133e34217ad0e8f925c320323fb270d334badf13c19b1a8ac14daa.OAuth2PermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["oAuth2PermissionGrant_id"] = id
    }
    return i83568fd90c133e34217ad0e8f925c320323fb270d334badf13c19b1a8ac14daa.NewOAuth2PermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) OwnedObjects()(*ia329fbf2e4e7682fe90bdfc6bf4cc749c189487202d60eefa0fdec3b62b0b7c3.OwnedObjectsRequestBuilder) {
    return ia329fbf2e4e7682fe90bdfc6bf4cc749c189487202d60eefa0fdec3b62b0b7c3.NewOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedObjectsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.servicePrincipals.item.ownedObjects.item collection
func (m *ServicePrincipalItemRequestBuilder) OwnedObjectsById(id string)(*i9fe3c2224826dcb0cde6e1f06bdfb308132901afdafc77fa24614124022a183b.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i9fe3c2224826dcb0cde6e1f06bdfb308132901afdafc77fa24614124022a183b.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) Owners()(*ia161e203c4876e46996f4ce25080c383320b3f90c23f41a4644fabc41a122c2c.OwnersRequestBuilder) {
    return ia161e203c4876e46996f4ce25080c383320b3f90c23f41a4644fabc41a122c2c.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.servicePrincipals.item.owners.item collection
func (m *ServicePrincipalItemRequestBuilder) OwnersById(id string)(*ie10051a8bcf74fee28b7346347c913cf4c35b13dc5337d6b1877c19ead0529ae.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return ie10051a8bcf74fee28b7346347c913cf4c35b13dc5337d6b1877c19ead0529ae.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in servicePrincipals
func (m *ServicePrincipalItemRequestBuilder) Patch(options *ServicePrincipalItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ServicePrincipalItemRequestBuilder) RemoveKey()(*i3c2ed9adbc44109667bcb7b30ffd0c6b2bb97d62b8e12c57b863c71af15e6903.RemoveKeyRequestBuilder) {
    return i3c2ed9adbc44109667bcb7b30ffd0c6b2bb97d62b8e12c57b863c71af15e6903.NewRemoveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) RemovePassword()(*i003add1c72d92bd47a9ee37d3f334a2b1149222e48ce785c0b07d0542f89302f.RemovePasswordRequestBuilder) {
    return i003add1c72d92bd47a9ee37d3f334a2b1149222e48ce785c0b07d0542f89302f.NewRemovePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) Restore()(*if93b2c8d7b19d0341ed3f43075bc082dd6d4c57bb7cafaa0f2161fce000d6f1d.RestoreRequestBuilder) {
    return if93b2c8d7b19d0341ed3f43075bc082dd6d4c57bb7cafaa0f2161fce000d6f1d.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) TokenIssuancePolicies()(*i2ebeefd4c13ab730e7f8f02313f1b93261a541ed22898061b1c09f058126615a.TokenIssuancePoliciesRequestBuilder) {
    return i2ebeefd4c13ab730e7f8f02313f1b93261a541ed22898061b1c09f058126615a.NewTokenIssuancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenIssuancePoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.servicePrincipals.item.tokenIssuancePolicies.item collection
func (m *ServicePrincipalItemRequestBuilder) TokenIssuancePoliciesById(id string)(*i4c7a31e042d98d59ede758ea63655f65992075acf7610c7db5c90df699c7c9e2.TokenIssuancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenIssuancePolicy_id"] = id
    }
    return i4c7a31e042d98d59ede758ea63655f65992075acf7610c7db5c90df699c7c9e2.NewTokenIssuancePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) TokenLifetimePolicies()(*ia47ff52d5020404c6eaea0bc5dbd4bbafaa1e0c73dbc6d4dcfee263b2245b53e.TokenLifetimePoliciesRequestBuilder) {
    return ia47ff52d5020404c6eaea0bc5dbd4bbafaa1e0c73dbc6d4dcfee263b2245b53e.NewTokenLifetimePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenLifetimePoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.servicePrincipals.item.tokenLifetimePolicies.item collection
func (m *ServicePrincipalItemRequestBuilder) TokenLifetimePoliciesById(id string)(*i14cf3ba6d0e45059a525d05353e514742297a11654926e06d365c73081a6a0dd.TokenLifetimePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenLifetimePolicy_id"] = id
    }
    return i14cf3ba6d0e45059a525d05353e514742297a11654926e06d365c73081a6a0dd.NewTokenLifetimePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServicePrincipalItemRequestBuilder) TransitiveMemberOf()(*i60f4aa5edd1e35c0b5a11cc9ce9afe3a692e2aef2b419b5d6bb6c9bd8c0c2561.TransitiveMemberOfRequestBuilder) {
    return i60f4aa5edd1e35c0b5a11cc9ce9afe3a692e2aef2b419b5d6bb6c9bd8c0c2561.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.servicePrincipals.item.transitiveMemberOf.item collection
func (m *ServicePrincipalItemRequestBuilder) TransitiveMemberOfById(id string)(*ib95837107c1a1d541878efdf1d0902a7c67b80de5fc97d771d55c403ab043621.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return ib95837107c1a1d541878efdf1d0902a7c67b80de5fc97d771d55c403ab043621.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
