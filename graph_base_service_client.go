package msgraphsdkgo

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e "github.com/microsoftgraph/msgraph-sdk-go/builders"
)

// GraphBaseServiceClient the main entry point of the SDK, exposes the configuration and the fluent API.
type GraphBaseServiceClient struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Admin provides operations to manage the admin singleton.
func (m *GraphBaseServiceClient) Admin()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.AdminRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewAdminRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgreementAcceptances provides operations to manage the collection of agreementAcceptance entities.
func (m *GraphBaseServiceClient) AgreementAcceptances()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.AgreementAcceptancesRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewAgreementAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgreementAcceptancesById provides operations to manage the collection of agreementAcceptance entities.
func (m *GraphBaseServiceClient) AgreementAcceptancesById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.AgreementAcceptancesAgreementAcceptanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewAgreementAcceptancesAgreementAcceptanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Agreements provides operations to manage the collection of agreement entities.
func (m *GraphBaseServiceClient) Agreements()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.AgreementsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewAgreementsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgreementsById provides operations to manage the collection of agreement entities.
func (m *GraphBaseServiceClient) AgreementsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.AgreementsAgreementItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreement%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewAgreementsAgreementItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AppCatalogs provides operations to manage the appCatalogs singleton.
func (m *GraphBaseServiceClient) AppCatalogs()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.AppCatalogsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewAppCatalogsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Applications provides operations to manage the collection of application entities.
func (m *GraphBaseServiceClient) Applications()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ApplicationsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewApplicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplicationsById provides operations to manage the collection of application entities.
func (m *GraphBaseServiceClient) ApplicationsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ApplicationsApplicationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["application%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewApplicationsApplicationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ApplicationTemplates provides operations to manage the collection of applicationTemplate entities.
func (m *GraphBaseServiceClient) ApplicationTemplates()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ApplicationTemplatesRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewApplicationTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplicationTemplatesById provides operations to manage the collection of applicationTemplate entities.
func (m *GraphBaseServiceClient) ApplicationTemplatesById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ApplicationTemplatesApplicationTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["applicationTemplate%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewApplicationTemplatesApplicationTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuditLogs provides operations to manage the auditLogRoot singleton.
func (m *GraphBaseServiceClient) AuditLogs()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.AuditLogsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewAuditLogsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationMethodConfigurations provides operations to manage the collection of authenticationMethodConfiguration entities.
func (m *GraphBaseServiceClient) AuthenticationMethodConfigurations()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.AuthenticationMethodConfigurationsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewAuthenticationMethodConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationMethodConfigurationsById provides operations to manage the collection of authenticationMethodConfiguration entities.
func (m *GraphBaseServiceClient) AuthenticationMethodConfigurationsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationMethodConfiguration%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuthenticationMethodsPolicy provides operations to manage the authenticationMethodsPolicy singleton.
func (m *GraphBaseServiceClient) AuthenticationMethodsPolicy()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.AuthenticationMethodsPolicyRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewAuthenticationMethodsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Branding provides operations to manage the organizationalBranding singleton.
func (m *GraphBaseServiceClient) Branding()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.BrandingRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewBrandingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificateBasedAuthConfiguration provides operations to manage the collection of certificateBasedAuthConfiguration entities.
func (m *GraphBaseServiceClient) CertificateBasedAuthConfiguration()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.CertificateBasedAuthConfigurationRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewCertificateBasedAuthConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificateBasedAuthConfigurationById provides operations to manage the collection of certificateBasedAuthConfiguration entities.
func (m *GraphBaseServiceClient) CertificateBasedAuthConfigurationById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["certificateBasedAuthConfiguration%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewCertificateBasedAuthConfigurationCertificateBasedAuthConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Chats provides operations to manage the collection of chat entities.
func (m *GraphBaseServiceClient) Chats()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ChatsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewChatsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChatsById provides operations to manage the collection of chat entities.
func (m *GraphBaseServiceClient) ChatsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ChatsChatItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chat%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewChatsChatItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Communications provides operations to manage the cloudCommunications singleton.
func (m *GraphBaseServiceClient) Communications()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.CommunicationsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewCommunicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Compliance provides operations to manage the compliance singleton.
func (m *GraphBaseServiceClient) Compliance()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ComplianceRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewComplianceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Connections provides operations to manage the collection of externalConnection entities.
func (m *GraphBaseServiceClient) Connections()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ConnectionsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewConnectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectionsById provides operations to manage the collection of externalConnection entities.
func (m *GraphBaseServiceClient) ConnectionsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ConnectionsExternalConnectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalConnection%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewConnectionsExternalConnectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewGraphBaseServiceClient instantiates a new GraphBaseServiceClient and sets the default values.
func NewGraphBaseServiceClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GraphBaseServiceClient) {
    m := &GraphBaseServiceClient{
    }
    m.pathParameters = make(map[string]string);
    m.urlTemplate = "{+baseurl}";
    m.requestAdapter = requestAdapter;
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    if m.requestAdapter.GetBaseUrl() == "" {
        m.requestAdapter.SetBaseUrl("https://graph.microsoft.com/v1.0")
    }
    return m
}
// Contacts provides operations to manage the collection of orgContact entities.
func (m *GraphBaseServiceClient) Contacts()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ContactsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactsById provides operations to manage the collection of orgContact entities.
func (m *GraphBaseServiceClient) ContactsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ContactsOrgContactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["orgContact%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewContactsOrgContactItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Contracts provides operations to manage the collection of contract entities.
func (m *GraphBaseServiceClient) Contracts()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ContractsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewContractsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContractsById provides operations to manage the collection of contract entities.
func (m *GraphBaseServiceClient) ContractsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ContractsContractItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contract%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewContractsContractItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DataPolicyOperations provides operations to manage the collection of dataPolicyOperation entities.
func (m *GraphBaseServiceClient) DataPolicyOperations()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DataPolicyOperationsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDataPolicyOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DataPolicyOperationsById provides operations to manage the collection of dataPolicyOperation entities.
func (m *GraphBaseServiceClient) DataPolicyOperationsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DataPolicyOperationsDataPolicyOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataPolicyOperation%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDataPolicyOperationsDataPolicyOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceAppManagement provides operations to manage the deviceAppManagement singleton.
func (m *GraphBaseServiceClient) DeviceAppManagement()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DeviceAppManagementRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDeviceAppManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceManagement provides operations to manage the deviceManagement singleton.
func (m *GraphBaseServiceClient) DeviceManagement()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DeviceManagementRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDeviceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Devices provides operations to manage the collection of device entities.
func (m *GraphBaseServiceClient) Devices()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DevicesRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DevicesById provides operations to manage the collection of device entities.
func (m *GraphBaseServiceClient) DevicesById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DevicesDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["device%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDevicesDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Directory provides operations to manage the directory singleton.
func (m *GraphBaseServiceClient) Directory()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DirectoryRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDirectoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryObjects provides operations to manage the collection of directoryObject entities.
func (m *GraphBaseServiceClient) DirectoryObjects()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DirectoryObjectsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDirectoryObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryObjectsById provides operations to manage the collection of directoryObject entities.
func (m *GraphBaseServiceClient) DirectoryObjectsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DirectoryObjectsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDirectoryObjectsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DirectoryRoles provides operations to manage the collection of directoryRole entities.
func (m *GraphBaseServiceClient) DirectoryRoles()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DirectoryRolesRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDirectoryRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryRolesById provides operations to manage the collection of directoryRole entities.
func (m *GraphBaseServiceClient) DirectoryRolesById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DirectoryRolesDirectoryRoleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryRole%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDirectoryRolesDirectoryRoleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DirectoryRoleTemplates provides operations to manage the collection of directoryRoleTemplate entities.
func (m *GraphBaseServiceClient) DirectoryRoleTemplates()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DirectoryRoleTemplatesRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDirectoryRoleTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryRoleTemplatesById provides operations to manage the collection of directoryRoleTemplate entities.
func (m *GraphBaseServiceClient) DirectoryRoleTemplatesById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DirectoryRoleTemplatesDirectoryRoleTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryRoleTemplate%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDirectoryRoleTemplatesDirectoryRoleTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DomainDnsRecords provides operations to manage the collection of domainDnsRecord entities.
func (m *GraphBaseServiceClient) DomainDnsRecords()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DomainDnsRecordsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDomainDnsRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DomainDnsRecordsById provides operations to manage the collection of domainDnsRecord entities.
func (m *GraphBaseServiceClient) DomainDnsRecordsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DomainDnsRecordsDomainDnsRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domainDnsRecord%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDomainDnsRecordsDomainDnsRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Domains provides operations to manage the collection of domain entities.
func (m *GraphBaseServiceClient) Domains()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DomainsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDomainsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DomainsById provides operations to manage the collection of domain entities.
func (m *GraphBaseServiceClient) DomainsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DomainsDomainItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domain%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDomainsDomainItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Drive provides operations to manage the drive singleton.
func (m *GraphBaseServiceClient) Drive()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DriveRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives provides operations to manage the collection of drive entities.
func (m *GraphBaseServiceClient) Drives()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DrivesRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById provides operations to manage the collection of drive entities.
func (m *GraphBaseServiceClient) DrivesById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.DrivesDriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewDrivesDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Education provides operations to manage the educationRoot singleton.
func (m *GraphBaseServiceClient) Education()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.EducationRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewEducationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// External provides operations to manage the external singleton.
func (m *GraphBaseServiceClient) External()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ExternalRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewExternalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupLifecyclePolicies provides operations to manage the collection of groupLifecyclePolicy entities.
func (m *GraphBaseServiceClient) GroupLifecyclePolicies()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.GroupLifecyclePoliciesRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewGroupLifecyclePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupLifecyclePoliciesById provides operations to manage the collection of groupLifecyclePolicy entities.
func (m *GraphBaseServiceClient) GroupLifecyclePoliciesById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.GroupLifecyclePoliciesGroupLifecyclePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupLifecyclePolicy%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewGroupLifecyclePoliciesGroupLifecyclePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Groups provides operations to manage the collection of group entities.
func (m *GraphBaseServiceClient) Groups()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.GroupsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupsById provides operations to manage the collection of group entities.
func (m *GraphBaseServiceClient) GroupsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.GroupsGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewGroupsGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GroupSettings provides operations to manage the collection of groupSetting entities.
func (m *GraphBaseServiceClient) GroupSettings()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.GroupSettingsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewGroupSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupSettingsById provides operations to manage the collection of groupSetting entities.
func (m *GraphBaseServiceClient) GroupSettingsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.GroupSettingsGroupSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupSetting%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewGroupSettingsGroupSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GroupSettingTemplates provides operations to manage the collection of groupSettingTemplate entities.
func (m *GraphBaseServiceClient) GroupSettingTemplates()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.GroupSettingTemplatesRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewGroupSettingTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupSettingTemplatesById provides operations to manage the collection of groupSettingTemplate entities.
func (m *GraphBaseServiceClient) GroupSettingTemplatesById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.GroupSettingTemplatesGroupSettingTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupSettingTemplate%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewGroupSettingTemplatesGroupSettingTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Identity provides operations to manage the identityContainer singleton.
func (m *GraphBaseServiceClient) Identity()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.IdentityRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewIdentityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityGovernance provides operations to manage the identityGovernance singleton.
func (m *GraphBaseServiceClient) IdentityGovernance()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.IdentityGovernanceRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewIdentityGovernanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityProtection provides operations to manage the identityProtectionRoot singleton.
func (m *GraphBaseServiceClient) IdentityProtection()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.IdentityProtectionRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewIdentityProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityProviders provides operations to manage the collection of identityProvider entities.
func (m *GraphBaseServiceClient) IdentityProviders()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.IdentityProvidersRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityProvidersById provides operations to manage the collection of identityProvider entities.
func (m *GraphBaseServiceClient) IdentityProvidersById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.IdentityProvidersIdentityProviderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProvider%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewIdentityProvidersIdentityProviderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InformationProtection provides operations to manage the informationProtection singleton.
func (m *GraphBaseServiceClient) InformationProtection()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.InformationProtectionRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Invitations provides operations to manage the collection of invitation entities.
func (m *GraphBaseServiceClient) Invitations()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.InvitationsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewInvitationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InvitationsById provides operations to manage the collection of invitation entities.
func (m *GraphBaseServiceClient) InvitationsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.InvitationsInvitationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["invitation%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewInvitationsInvitationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Localizations provides operations to manage the collection of organizationalBrandingLocalization entities.
func (m *GraphBaseServiceClient) Localizations()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.LocalizationsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewLocalizationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocalizationsById provides operations to manage the collection of organizationalBrandingLocalization entities.
func (m *GraphBaseServiceClient) LocalizationsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.LocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["organizationalBrandingLocalization%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Me provides operations to manage the user singleton.
func (m *GraphBaseServiceClient) Me()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.UsersUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["user%2Did"] = "me"
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewUsersUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Oauth2PermissionGrants provides operations to manage the collection of oAuth2PermissionGrant entities.
func (m *GraphBaseServiceClient) Oauth2PermissionGrants()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.Oauth2PermissionGrantsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Oauth2PermissionGrantsById provides operations to manage the collection of oAuth2PermissionGrant entities.
func (m *GraphBaseServiceClient) Oauth2PermissionGrantsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.Oauth2PermissionGrantsOAuth2PermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["oAuth2PermissionGrant%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewOauth2PermissionGrantsOAuth2PermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Organization provides operations to manage the collection of organization entities.
func (m *GraphBaseServiceClient) Organization()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.OrganizationRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewOrganizationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrganizationById provides operations to manage the collection of organization entities.
func (m *GraphBaseServiceClient) OrganizationById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.OrganizationOrganizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["organization%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewOrganizationOrganizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PermissionGrants provides operations to manage the collection of resourceSpecificPermissionGrant entities.
func (m *GraphBaseServiceClient) PermissionGrants()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.PermissionGrantsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById provides operations to manage the collection of resourceSpecificPermissionGrant entities.
func (m *GraphBaseServiceClient) PermissionGrantsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.PermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Places provides operations to manage the collection of place entities.
func (m *GraphBaseServiceClient) Places()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.PlacesRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewPlacesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PlacesById provides operations to manage the collection of place entities.
func (m *GraphBaseServiceClient) PlacesById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.PlacesPlaceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["place%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewPlacesPlaceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Planner provides operations to manage the planner singleton.
func (m *GraphBaseServiceClient) Planner()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.PlannerRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Policies provides operations to manage the policyRoot singleton.
func (m *GraphBaseServiceClient) Policies()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.PoliciesRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Print provides operations to manage the print singleton.
func (m *GraphBaseServiceClient) Print()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.PrintRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewPrintRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Privacy provides operations to manage the privacy singleton.
func (m *GraphBaseServiceClient) Privacy()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.PrivacyRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewPrivacyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reports provides operations to manage the reportRoot singleton.
func (m *GraphBaseServiceClient) Reports()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ReportsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleManagement provides operations to manage the roleManagement singleton.
func (m *GraphBaseServiceClient) RoleManagement()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.RoleManagementRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewRoleManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchemaExtensions provides operations to manage the collection of schemaExtension entities.
func (m *GraphBaseServiceClient) SchemaExtensions()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.SchemaExtensionsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewSchemaExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchemaExtensionsById provides operations to manage the collection of schemaExtension entities.
func (m *GraphBaseServiceClient) SchemaExtensionsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.SchemaExtensionsSchemaExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schemaExtension%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewSchemaExtensionsSchemaExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ScopedRoleMemberships provides operations to manage the collection of scopedRoleMembership entities.
func (m *GraphBaseServiceClient) ScopedRoleMemberships()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ScopedRoleMembershipsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewScopedRoleMembershipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMembershipsById provides operations to manage the collection of scopedRoleMembership entities.
func (m *GraphBaseServiceClient) ScopedRoleMembershipsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ScopedRoleMembershipsScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewScopedRoleMembershipsScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Search provides operations to manage the searchEntity singleton.
func (m *GraphBaseServiceClient) Search()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.SearchRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewSearchRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Security provides operations to manage the security singleton.
func (m *GraphBaseServiceClient) Security()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.SecurityRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewSecurityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipals provides operations to manage the collection of servicePrincipal entities.
func (m *GraphBaseServiceClient) ServicePrincipals()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ServicePrincipalsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewServicePrincipalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipalsById provides operations to manage the collection of servicePrincipal entities.
func (m *GraphBaseServiceClient) ServicePrincipalsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.ServicePrincipalsServicePrincipalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["servicePrincipal%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewServicePrincipalsServicePrincipalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Shares provides operations to manage the collection of sharedDriveItem entities.
func (m *GraphBaseServiceClient) Shares()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.SharesRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharesById provides operations to manage the collection of sharedDriveItem entities.
func (m *GraphBaseServiceClient) SharesById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.SharesSharedDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedDriveItem%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewSharesSharedDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites provides operations to manage the collection of site entities.
func (m *GraphBaseServiceClient) Sites()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.SitesRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById provides operations to manage the collection of site entities.
func (m *GraphBaseServiceClient) SitesById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.SitesSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewSitesSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Solutions provides operations to manage the solutionsRoot singleton.
func (m *GraphBaseServiceClient) Solutions()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.SolutionsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewSolutionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribedSkus provides operations to manage the collection of subscribedSku entities.
func (m *GraphBaseServiceClient) SubscribedSkus()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.SubscribedSkusRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewSubscribedSkusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribedSkusById provides operations to manage the collection of subscribedSku entities.
func (m *GraphBaseServiceClient) SubscribedSkusById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.SubscribedSkusSubscribedSkuItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscribedSku%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewSubscribedSkusSubscribedSkuItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Subscriptions provides operations to manage the collection of subscription entities.
func (m *GraphBaseServiceClient) Subscriptions()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.SubscriptionsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById provides operations to manage the collection of subscription entities.
func (m *GraphBaseServiceClient) SubscriptionsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.SubscriptionsSubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewSubscriptionsSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Teams provides operations to manage the collection of team entities.
func (m *GraphBaseServiceClient) Teams()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.TeamsRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TeamsById provides operations to manage the collection of team entities.
func (m *GraphBaseServiceClient) TeamsById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.TeamsTeamItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["team%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewTeamsTeamItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TeamsTemplates provides operations to manage the collection of teamsTemplate entities.
func (m *GraphBaseServiceClient) TeamsTemplates()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.TeamsTemplatesRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewTeamsTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TeamsTemplatesById provides operations to manage the collection of teamsTemplate entities.
func (m *GraphBaseServiceClient) TeamsTemplatesById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.TeamsTemplatesTeamsTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTemplate%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewTeamsTemplatesTeamsTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Teamwork provides operations to manage the teamwork singleton.
func (m *GraphBaseServiceClient) Teamwork()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.TeamworkRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewTeamworkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Users provides operations to manage the collection of user entities.
func (m *GraphBaseServiceClient) Users()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.UsersRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsersById provides operations to manage the collection of user entities.
func (m *GraphBaseServiceClient) UsersById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.UsersUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewUsersUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Workbooks provides operations to manage the collection of driveItem entities.
func (m *GraphBaseServiceClient) Workbooks()(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.WorkbooksRequestBuilder) {
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewWorkbooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WorkbooksById provides operations to manage the collection of driveItem entities.
func (m *GraphBaseServiceClient) WorkbooksById(id string)(*ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.WorkbooksDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return ib12d3019d2395479aa856f09c37532c7d6064c28a6bf48bee1e80bf96f14cb5e.NewWorkbooksDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
