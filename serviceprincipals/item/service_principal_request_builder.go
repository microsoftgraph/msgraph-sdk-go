package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i003add1c72d92bd47a9ee37d3f334a2b1149222e48ce785c0b07d0542f89302f "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/removepassword"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
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
    ic2b775c554f6f0689bed0deb6e610608281614f70275cada61da8274360068d0 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/endpoints"
    icc1091ecf66b1de7080efaeb2e59cdf1074f3b116fba7066b71a0d6cd9c57a53 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/claimsmappingpolicies"
    id692483c2e6fa4444b9d250ecaa7a6bea16dbc3cd7fa20c05fb69b2f7b69ad46 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/approleassignedto"
    ida4a28e902893200b019f06f91d4d3d9d5d69f185e4f2b9b706d4abd32593768 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/memberof"
    idbd9734414b90f65880420caeaf7c9aad49278527e828f8795e9b8bc0c65d8a5 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/oauth2permissiongrants"
    ie49af58d75366149082d7cc96496e031433add00a7bde619d68ae8a646ef55e2 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/homerealmdiscoverypolicies"
    iedd8dd60eed89931552d2c873e7c568c2a30c87dee999cc992446db26f8803fc "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/delegatedpermissionclassifications"
    if93b2c8d7b19d0341ed3f43075bc082dd6d4c57bb7cafaa0f2161fce000d6f1d "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/restore"
    i38c5edd99af7364922d88fc2b49de45f320d11c5fa36e8234c12f1d863cbea73 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/delegatedpermissionclassifications/item"
    i80395ca3a6452268f5c4d3e92f315d7c92e78fef9911bf32e289ea47d6968415 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/approleassignments/item"
    iae6f1167638ef01ba227e18af033b58e9203edca74b4641453bc1cec1de3aa91 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/endpoints/item"
    ieade744732b95d81e11bac4d5146b5ac0bfeb407228f97be47815b5b44c77ff1 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/approleassignedto/item"
)

// Builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}
type ServicePrincipalRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ServicePrincipalRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ServicePrincipalRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ServicePrincipalRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get entity from servicePrincipals by key
type ServicePrincipalRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ServicePrincipalRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServicePrincipal;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ServicePrincipalRequestBuilder) AddKey()(*i3f559eeb0a32c852d52064d33f878279ea63214977aec7d5605657ac813f843e.AddKeyRequestBuilder) {
    return i3f559eeb0a32c852d52064d33f878279ea63214977aec7d5605657ac813f843e.NewAddKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) AddPassword()(*i495446d97cfd23af0607d58fb1c29e62eeec78e42cbd747f401374edf193b049.AddPasswordRequestBuilder) {
    return i495446d97cfd23af0607d58fb1c29e62eeec78e42cbd747f401374edf193b049.NewAddPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) AppRoleAssignedTo()(*id692483c2e6fa4444b9d250ecaa7a6bea16dbc3cd7fa20c05fb69b2f7b69ad46.AppRoleAssignedToRequestBuilder) {
    return id692483c2e6fa4444b9d250ecaa7a6bea16dbc3cd7fa20c05fb69b2f7b69ad46.NewAppRoleAssignedToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.servicePrincipals.item.appRoleAssignedTo.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ServicePrincipalRequestBuilder) AppRoleAssignedToById(id string)(*ieade744732b95d81e11bac4d5146b5ac0bfeb407228f97be47815b5b44c77ff1.AppRoleAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment_id"] = id
    }
    return ieade744732b95d81e11bac4d5146b5ac0bfeb407228f97be47815b5b44c77ff1.NewAppRoleAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) AppRoleAssignments()(*i93153ff64f35a046ae51130f6dbc11813af572bf8610f8142e2bebbb77269190.AppRoleAssignmentsRequestBuilder) {
    return i93153ff64f35a046ae51130f6dbc11813af572bf8610f8142e2bebbb77269190.NewAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.servicePrincipals.item.appRoleAssignments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ServicePrincipalRequestBuilder) AppRoleAssignmentsById(id string)(*i80395ca3a6452268f5c4d3e92f315d7c92e78fef9911bf32e289ea47d6968415.AppRoleAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment_id"] = id
    }
    return i80395ca3a6452268f5c4d3e92f315d7c92e78fef9911bf32e289ea47d6968415.NewAppRoleAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) CheckMemberGroups()(*i077b2e62603ccb12d8cc7601297f5d1a834bbe55d05e304176e9f9ad509fc426.CheckMemberGroupsRequestBuilder) {
    return i077b2e62603ccb12d8cc7601297f5d1a834bbe55d05e304176e9f9ad509fc426.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) CheckMemberObjects()(*i6265425e2a2080afde3979740ad79538958bd2d2ab7ce003f7d0a624b06989bc.CheckMemberObjectsRequestBuilder) {
    return i6265425e2a2080afde3979740ad79538958bd2d2ab7ce003f7d0a624b06989bc.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) ClaimsMappingPolicies()(*icc1091ecf66b1de7080efaeb2e59cdf1074f3b116fba7066b71a0d6cd9c57a53.ClaimsMappingPoliciesRequestBuilder) {
    return icc1091ecf66b1de7080efaeb2e59cdf1074f3b116fba7066b71a0d6cd9c57a53.NewClaimsMappingPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ServicePrincipalRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewServicePrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServicePrincipalRequestBuilder) {
    m := &ServicePrincipalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ServicePrincipalRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewServicePrincipalRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServicePrincipalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete entity from servicePrincipals
// Parameters:
//  - options : Options for the request
func (m *ServicePrincipalRequestBuilder) CreateDeleteRequestInformation(options *ServicePrincipalRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ServicePrincipalRequestBuilder) CreatedObjects()(*i1f042d332a3da52b9a4d53ae9423654b32107cbe9b637d9ca95d07e583140364.CreatedObjectsRequestBuilder) {
    return i1f042d332a3da52b9a4d53ae9423654b32107cbe9b637d9ca95d07e583140364.NewCreatedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get entity from servicePrincipals by key
// Parameters:
//  - options : Options for the request
func (m *ServicePrincipalRequestBuilder) CreateGetRequestInformation(options *ServicePrincipalRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Update entity in servicePrincipals
// Parameters:
//  - options : Options for the request
func (m *ServicePrincipalRequestBuilder) CreatePatchRequestInformation(options *ServicePrincipalRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ServicePrincipalRequestBuilder) DelegatedPermissionClassifications()(*iedd8dd60eed89931552d2c873e7c568c2a30c87dee999cc992446db26f8803fc.DelegatedPermissionClassificationsRequestBuilder) {
    return iedd8dd60eed89931552d2c873e7c568c2a30c87dee999cc992446db26f8803fc.NewDelegatedPermissionClassificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.servicePrincipals.item.delegatedPermissionClassifications.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ServicePrincipalRequestBuilder) DelegatedPermissionClassificationsById(id string)(*i38c5edd99af7364922d88fc2b49de45f320d11c5fa36e8234c12f1d863cbea73.DelegatedPermissionClassificationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedPermissionClassification_id"] = id
    }
    return i38c5edd99af7364922d88fc2b49de45f320d11c5fa36e8234c12f1d863cbea73.NewDelegatedPermissionClassificationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete entity from servicePrincipals
// Parameters:
//  - options : Options for the request
func (m *ServicePrincipalRequestBuilder) Delete(options *ServicePrincipalRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ServicePrincipalRequestBuilder) Endpoints()(*ic2b775c554f6f0689bed0deb6e610608281614f70275cada61da8274360068d0.EndpointsRequestBuilder) {
    return ic2b775c554f6f0689bed0deb6e610608281614f70275cada61da8274360068d0.NewEndpointsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.servicePrincipals.item.endpoints.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ServicePrincipalRequestBuilder) EndpointsById(id string)(*iae6f1167638ef01ba227e18af033b58e9203edca74b4641453bc1cec1de3aa91.EndpointRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["endpoint_id"] = id
    }
    return iae6f1167638ef01ba227e18af033b58e9203edca74b4641453bc1cec1de3aa91.NewEndpointRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get entity from servicePrincipals by key
// Parameters:
//  - options : Options for the request
func (m *ServicePrincipalRequestBuilder) Get(options *ServicePrincipalRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServicePrincipal, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewServicePrincipal() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServicePrincipal), nil
}
func (m *ServicePrincipalRequestBuilder) GetMemberGroups()(*iaa2be345892fc1350e44205c5f24bef008513f89d2497eabd4c16620df674828.GetMemberGroupsRequestBuilder) {
    return iaa2be345892fc1350e44205c5f24bef008513f89d2497eabd4c16620df674828.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) GetMemberObjects()(*i2c910a3074da68d087d5c6fc96ab8be3cf4e54ef4ec7985acd4e132e52d42137.GetMemberObjectsRequestBuilder) {
    return i2c910a3074da68d087d5c6fc96ab8be3cf4e54ef4ec7985acd4e132e52d42137.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) HomeRealmDiscoveryPolicies()(*ie49af58d75366149082d7cc96496e031433add00a7bde619d68ae8a646ef55e2.HomeRealmDiscoveryPoliciesRequestBuilder) {
    return ie49af58d75366149082d7cc96496e031433add00a7bde619d68ae8a646ef55e2.NewHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) MemberOf()(*ida4a28e902893200b019f06f91d4d3d9d5d69f185e4f2b9b706d4abd32593768.MemberOfRequestBuilder) {
    return ida4a28e902893200b019f06f91d4d3d9d5d69f185e4f2b9b706d4abd32593768.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) Oauth2PermissionGrants()(*idbd9734414b90f65880420caeaf7c9aad49278527e828f8795e9b8bc0c65d8a5.Oauth2PermissionGrantsRequestBuilder) {
    return idbd9734414b90f65880420caeaf7c9aad49278527e828f8795e9b8bc0c65d8a5.NewOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) OwnedObjects()(*ia329fbf2e4e7682fe90bdfc6bf4cc749c189487202d60eefa0fdec3b62b0b7c3.OwnedObjectsRequestBuilder) {
    return ia329fbf2e4e7682fe90bdfc6bf4cc749c189487202d60eefa0fdec3b62b0b7c3.NewOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) Owners()(*ia161e203c4876e46996f4ce25080c383320b3f90c23f41a4644fabc41a122c2c.OwnersRequestBuilder) {
    return ia161e203c4876e46996f4ce25080c383320b3f90c23f41a4644fabc41a122c2c.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Update entity in servicePrincipals
// Parameters:
//  - options : Options for the request
func (m *ServicePrincipalRequestBuilder) Patch(options *ServicePrincipalRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ServicePrincipalRequestBuilder) RemoveKey()(*i3c2ed9adbc44109667bcb7b30ffd0c6b2bb97d62b8e12c57b863c71af15e6903.RemoveKeyRequestBuilder) {
    return i3c2ed9adbc44109667bcb7b30ffd0c6b2bb97d62b8e12c57b863c71af15e6903.NewRemoveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) RemovePassword()(*i003add1c72d92bd47a9ee37d3f334a2b1149222e48ce785c0b07d0542f89302f.RemovePasswordRequestBuilder) {
    return i003add1c72d92bd47a9ee37d3f334a2b1149222e48ce785c0b07d0542f89302f.NewRemovePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) Restore()(*if93b2c8d7b19d0341ed3f43075bc082dd6d4c57bb7cafaa0f2161fce000d6f1d.RestoreRequestBuilder) {
    return if93b2c8d7b19d0341ed3f43075bc082dd6d4c57bb7cafaa0f2161fce000d6f1d.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) TokenIssuancePolicies()(*i2ebeefd4c13ab730e7f8f02313f1b93261a541ed22898061b1c09f058126615a.TokenIssuancePoliciesRequestBuilder) {
    return i2ebeefd4c13ab730e7f8f02313f1b93261a541ed22898061b1c09f058126615a.NewTokenIssuancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) TokenLifetimePolicies()(*ia47ff52d5020404c6eaea0bc5dbd4bbafaa1e0c73dbc6d4dcfee263b2245b53e.TokenLifetimePoliciesRequestBuilder) {
    return ia47ff52d5020404c6eaea0bc5dbd4bbafaa1e0c73dbc6d4dcfee263b2245b53e.NewTokenLifetimePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) TransitiveMemberOf()(*i60f4aa5edd1e35c0b5a11cc9ce9afe3a692e2aef2b419b5d6bb6c9bd8c0c2561.TransitiveMemberOfRequestBuilder) {
    return i60f4aa5edd1e35c0b5a11cc9ce9afe3a692e2aef2b419b5d6bb6c9bd8c0c2561.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
