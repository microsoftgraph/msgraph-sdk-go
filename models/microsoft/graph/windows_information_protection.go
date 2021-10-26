package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WindowsInformationProtection struct {
    ManagedAppPolicy
    // Navigation property to list of security groups targeted for policy.
    assignments []TargetedManagedAppPolicyAssignment;
    // Specifies whether to allow Azure RMS encryption for WIP
    azureRightsManagementServicesAllowed *bool;
    // Specifies a recovery certificate that can be used for data recovery of encrypted files. This is the same as the data recovery agent(DRA) certificate for encrypting file system(EFS)
    dataRecoveryCertificate *WindowsInformationProtectionDataRecoveryCertificate;
    // WIP enforcement level.See the Enum definition for supported values. Possible values are: noProtection, encryptAndAuditOnly, encryptAuditAndPrompt, encryptAuditAndBlock.
    enforcementLevel *WindowsInformationProtectionEnforcementLevel;
    // Primary enterprise domain
    enterpriseDomain *string;
    // This is the comma-separated list of internal proxy servers. For example, '157.54.14.28, 157.54.11.118, 10.202.14.167, 157.53.14.163, 157.69.210.59'. These proxies have been configured by the admin to connect to specific resources on the Internet. They are considered to be enterprise network locations. The proxies are only leveraged in configuring the EnterpriseProxiedDomains policy to force traffic to the matched domains through these proxies
    enterpriseInternalProxyServers []WindowsInformationProtectionResourceCollection;
    // Sets the enterprise IP ranges that define the computers in the enterprise network. Data that comes from those computers will be considered part of the enterprise and protected. These locations will be considered a safe destination for enterprise data to be shared to
    enterpriseIPRanges []WindowsInformationProtectionIPRangeCollection;
    // Boolean value that tells the client to accept the configured list and not to use heuristics to attempt to find other subnets. Default is false
    enterpriseIPRangesAreAuthoritative *bool;
    // This is the list of domains that comprise the boundaries of the enterprise. Data from one of these domains that is sent to a device will be considered enterprise data and protected These locations will be considered a safe destination for enterprise data to be shared to
    enterpriseNetworkDomainNames []WindowsInformationProtectionResourceCollection;
    // List of enterprise domains to be protected
    enterpriseProtectedDomainNames []WindowsInformationProtectionResourceCollection;
    // Contains a list of Enterprise resource domains hosted in the cloud that need to be protected. Connections to these resources are considered enterprise data. If a proxy is paired with a cloud resource, traffic to the cloud resource will be routed through the enterprise network via the denoted proxy server (on Port 80). A proxy server used for this purpose must also be configured using the EnterpriseInternalProxyServers policy
    enterpriseProxiedDomains []WindowsInformationProtectionProxiedDomainCollection;
    // This is a list of proxy servers. Any server not on this list is considered non-enterprise
    enterpriseProxyServers []WindowsInformationProtectionResourceCollection;
    // Boolean value that tells the client to accept the configured list of proxies and not try to detect other work proxies. Default is false
    enterpriseProxyServersAreAuthoritative *bool;
    // Another way to input exempt apps through xml files
    exemptAppLockerFiles []WindowsInformationProtectionAppLockerFile;
    // Exempt applications can also access enterprise data, but the data handled by those applications are not protected. This is because some critical enterprise applications may have compatibility problems with encrypted data.
    exemptApps []WindowsInformationProtectionApp;
    // Determines whether overlays are added to icons for WIP protected files in Explorer and enterprise only app tiles in the Start menu. Starting in Windows 10, version 1703 this setting also configures the visibility of the WIP icon in the title bar of a WIP-protected app
    iconsVisible *bool;
    // This switch is for the Windows Search Indexer, to allow or disallow indexing of items
    indexingEncryptedStoresOrItemsBlocked *bool;
    // Indicates if the policy is deployed to any inclusion groups or not.
    isAssigned *bool;
    // List of domain names that can used for work or personal resource
    neutralDomainResources []WindowsInformationProtectionResourceCollection;
    // Another way to input protected apps through xml files
    protectedAppLockerFiles []WindowsInformationProtectionAppLockerFile;
    // Protected applications can access enterprise data and the data handled by those applications are protected with encryption
    protectedApps []WindowsInformationProtectionApp;
    // Specifies whether the protection under lock feature (also known as encrypt under pin) should be configured
    protectionUnderLockConfigRequired *bool;
    // This policy controls whether to revoke the WIP keys when a device unenrolls from the management service. If set to 1 (Don't revoke keys), the keys will not be revoked and the user will continue to have access to protected files after unenrollment. If the keys are not revoked, there will be no revoked file cleanup subsequently.
    revokeOnUnenrollDisabled *bool;
    // TemplateID GUID to use for RMS encryption. The RMS template allows the IT admin to configure the details about who has access to RMS-protected file and how long they have access
    rightsManagementServicesTemplateId *string;
    // Specifies a list of file extensions, so that files with these extensions are encrypted when copying from an SMB share within the corporate boundary
    smbAutoEncryptedFileExtensions []WindowsInformationProtectionResourceCollection;
}
// Instantiates a new windowsInformationProtection and sets the default values.
func NewWindowsInformationProtection()(*WindowsInformationProtection) {
    m := &WindowsInformationProtection{
        ManagedAppPolicy: *NewManagedAppPolicy(),
    }
    return m
}
// Gets the assignments property value. Navigation property to list of security groups targeted for policy.
func (m *WindowsInformationProtection) GetAssignments()([]TargetedManagedAppPolicyAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the azureRightsManagementServicesAllowed property value. Specifies whether to allow Azure RMS encryption for WIP
func (m *WindowsInformationProtection) GetAzureRightsManagementServicesAllowed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.azureRightsManagementServicesAllowed
    }
}
// Gets the dataRecoveryCertificate property value. Specifies a recovery certificate that can be used for data recovery of encrypted files. This is the same as the data recovery agent(DRA) certificate for encrypting file system(EFS)
func (m *WindowsInformationProtection) GetDataRecoveryCertificate()(*WindowsInformationProtectionDataRecoveryCertificate) {
    if m == nil {
        return nil
    } else {
        return m.dataRecoveryCertificate
    }
}
// Gets the enforcementLevel property value. WIP enforcement level.See the Enum definition for supported values. Possible values are: noProtection, encryptAndAuditOnly, encryptAuditAndPrompt, encryptAuditAndBlock.
func (m *WindowsInformationProtection) GetEnforcementLevel()(*WindowsInformationProtectionEnforcementLevel) {
    if m == nil {
        return nil
    } else {
        return m.enforcementLevel
    }
}
// Gets the enterpriseDomain property value. Primary enterprise domain
func (m *WindowsInformationProtection) GetEnterpriseDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseDomain
    }
}
// Gets the enterpriseInternalProxyServers property value. This is the comma-separated list of internal proxy servers. For example, '157.54.14.28, 157.54.11.118, 10.202.14.167, 157.53.14.163, 157.69.210.59'. These proxies have been configured by the admin to connect to specific resources on the Internet. They are considered to be enterprise network locations. The proxies are only leveraged in configuring the EnterpriseProxiedDomains policy to force traffic to the matched domains through these proxies
func (m *WindowsInformationProtection) GetEnterpriseInternalProxyServers()([]WindowsInformationProtectionResourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseInternalProxyServers
    }
}
// Gets the enterpriseIPRanges property value. Sets the enterprise IP ranges that define the computers in the enterprise network. Data that comes from those computers will be considered part of the enterprise and protected. These locations will be considered a safe destination for enterprise data to be shared to
func (m *WindowsInformationProtection) GetEnterpriseIPRanges()([]WindowsInformationProtectionIPRangeCollection) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseIPRanges
    }
}
// Gets the enterpriseIPRangesAreAuthoritative property value. Boolean value that tells the client to accept the configured list and not to use heuristics to attempt to find other subnets. Default is false
func (m *WindowsInformationProtection) GetEnterpriseIPRangesAreAuthoritative()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseIPRangesAreAuthoritative
    }
}
// Gets the enterpriseNetworkDomainNames property value. This is the list of domains that comprise the boundaries of the enterprise. Data from one of these domains that is sent to a device will be considered enterprise data and protected These locations will be considered a safe destination for enterprise data to be shared to
func (m *WindowsInformationProtection) GetEnterpriseNetworkDomainNames()([]WindowsInformationProtectionResourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseNetworkDomainNames
    }
}
// Gets the enterpriseProtectedDomainNames property value. List of enterprise domains to be protected
func (m *WindowsInformationProtection) GetEnterpriseProtectedDomainNames()([]WindowsInformationProtectionResourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseProtectedDomainNames
    }
}
// Gets the enterpriseProxiedDomains property value. Contains a list of Enterprise resource domains hosted in the cloud that need to be protected. Connections to these resources are considered enterprise data. If a proxy is paired with a cloud resource, traffic to the cloud resource will be routed through the enterprise network via the denoted proxy server (on Port 80). A proxy server used for this purpose must also be configured using the EnterpriseInternalProxyServers policy
func (m *WindowsInformationProtection) GetEnterpriseProxiedDomains()([]WindowsInformationProtectionProxiedDomainCollection) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseProxiedDomains
    }
}
// Gets the enterpriseProxyServers property value. This is a list of proxy servers. Any server not on this list is considered non-enterprise
func (m *WindowsInformationProtection) GetEnterpriseProxyServers()([]WindowsInformationProtectionResourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseProxyServers
    }
}
// Gets the enterpriseProxyServersAreAuthoritative property value. Boolean value that tells the client to accept the configured list of proxies and not try to detect other work proxies. Default is false
func (m *WindowsInformationProtection) GetEnterpriseProxyServersAreAuthoritative()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseProxyServersAreAuthoritative
    }
}
// Gets the exemptAppLockerFiles property value. Another way to input exempt apps through xml files
func (m *WindowsInformationProtection) GetExemptAppLockerFiles()([]WindowsInformationProtectionAppLockerFile) {
    if m == nil {
        return nil
    } else {
        return m.exemptAppLockerFiles
    }
}
// Gets the exemptApps property value. Exempt applications can also access enterprise data, but the data handled by those applications are not protected. This is because some critical enterprise applications may have compatibility problems with encrypted data.
func (m *WindowsInformationProtection) GetExemptApps()([]WindowsInformationProtectionApp) {
    if m == nil {
        return nil
    } else {
        return m.exemptApps
    }
}
// Gets the iconsVisible property value. Determines whether overlays are added to icons for WIP protected files in Explorer and enterprise only app tiles in the Start menu. Starting in Windows 10, version 1703 this setting also configures the visibility of the WIP icon in the title bar of a WIP-protected app
func (m *WindowsInformationProtection) GetIconsVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iconsVisible
    }
}
// Gets the indexingEncryptedStoresOrItemsBlocked property value. This switch is for the Windows Search Indexer, to allow or disallow indexing of items
func (m *WindowsInformationProtection) GetIndexingEncryptedStoresOrItemsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.indexingEncryptedStoresOrItemsBlocked
    }
}
// Gets the isAssigned property value. Indicates if the policy is deployed to any inclusion groups or not.
func (m *WindowsInformationProtection) GetIsAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAssigned
    }
}
// Gets the neutralDomainResources property value. List of domain names that can used for work or personal resource
func (m *WindowsInformationProtection) GetNeutralDomainResources()([]WindowsInformationProtectionResourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.neutralDomainResources
    }
}
// Gets the protectedAppLockerFiles property value. Another way to input protected apps through xml files
func (m *WindowsInformationProtection) GetProtectedAppLockerFiles()([]WindowsInformationProtectionAppLockerFile) {
    if m == nil {
        return nil
    } else {
        return m.protectedAppLockerFiles
    }
}
// Gets the protectedApps property value. Protected applications can access enterprise data and the data handled by those applications are protected with encryption
func (m *WindowsInformationProtection) GetProtectedApps()([]WindowsInformationProtectionApp) {
    if m == nil {
        return nil
    } else {
        return m.protectedApps
    }
}
// Gets the protectionUnderLockConfigRequired property value. Specifies whether the protection under lock feature (also known as encrypt under pin) should be configured
func (m *WindowsInformationProtection) GetProtectionUnderLockConfigRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.protectionUnderLockConfigRequired
    }
}
// Gets the revokeOnUnenrollDisabled property value. This policy controls whether to revoke the WIP keys when a device unenrolls from the management service. If set to 1 (Don't revoke keys), the keys will not be revoked and the user will continue to have access to protected files after unenrollment. If the keys are not revoked, there will be no revoked file cleanup subsequently.
func (m *WindowsInformationProtection) GetRevokeOnUnenrollDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.revokeOnUnenrollDisabled
    }
}
// Gets the rightsManagementServicesTemplateId property value. TemplateID GUID to use for RMS encryption. The RMS template allows the IT admin to configure the details about who has access to RMS-protected file and how long they have access
func (m *WindowsInformationProtection) GetRightsManagementServicesTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rightsManagementServicesTemplateId
    }
}
// Gets the smbAutoEncryptedFileExtensions property value. Specifies a list of file extensions, so that files with these extensions are encrypted when copying from an SMB share within the corporate boundary
func (m *WindowsInformationProtection) GetSmbAutoEncryptedFileExtensions()([]WindowsInformationProtectionResourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.smbAutoEncryptedFileExtensions
    }
}
// The deserialization information for the current model
func (m *WindowsInformationProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ManagedAppPolicy.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTargetedManagedAppPolicyAssignment() })
        if err != nil {
            return err
        }
        res := make([]TargetedManagedAppPolicyAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*TargetedManagedAppPolicyAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
    res["azureRightsManagementServicesAllowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAzureRightsManagementServicesAllowed(val)
        return nil
    }
    res["dataRecoveryCertificate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionDataRecoveryCertificate() })
        if err != nil {
            return err
        }
        m.SetDataRecoveryCertificate(val.(*WindowsInformationProtectionDataRecoveryCertificate))
        return nil
    }
    res["enforcementLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsInformationProtectionEnforcementLevel)
        if err != nil {
            return err
        }
        cast := val.(WindowsInformationProtectionEnforcementLevel)
        m.SetEnforcementLevel(&cast)
        return nil
    }
    res["enterpriseDomain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEnterpriseDomain(val)
        return nil
    }
    res["enterpriseInternalProxyServers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionResourceCollection() })
        if err != nil {
            return err
        }
        res := make([]WindowsInformationProtectionResourceCollection, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsInformationProtectionResourceCollection))
        }
        m.SetEnterpriseInternalProxyServers(res)
        return nil
    }
    res["enterpriseIPRanges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionIPRangeCollection() })
        if err != nil {
            return err
        }
        res := make([]WindowsInformationProtectionIPRangeCollection, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsInformationProtectionIPRangeCollection))
        }
        m.SetEnterpriseIPRanges(res)
        return nil
    }
    res["enterpriseIPRangesAreAuthoritative"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnterpriseIPRangesAreAuthoritative(val)
        return nil
    }
    res["enterpriseNetworkDomainNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionResourceCollection() })
        if err != nil {
            return err
        }
        res := make([]WindowsInformationProtectionResourceCollection, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsInformationProtectionResourceCollection))
        }
        m.SetEnterpriseNetworkDomainNames(res)
        return nil
    }
    res["enterpriseProtectedDomainNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionResourceCollection() })
        if err != nil {
            return err
        }
        res := make([]WindowsInformationProtectionResourceCollection, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsInformationProtectionResourceCollection))
        }
        m.SetEnterpriseProtectedDomainNames(res)
        return nil
    }
    res["enterpriseProxiedDomains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionProxiedDomainCollection() })
        if err != nil {
            return err
        }
        res := make([]WindowsInformationProtectionProxiedDomainCollection, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsInformationProtectionProxiedDomainCollection))
        }
        m.SetEnterpriseProxiedDomains(res)
        return nil
    }
    res["enterpriseProxyServers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionResourceCollection() })
        if err != nil {
            return err
        }
        res := make([]WindowsInformationProtectionResourceCollection, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsInformationProtectionResourceCollection))
        }
        m.SetEnterpriseProxyServers(res)
        return nil
    }
    res["enterpriseProxyServersAreAuthoritative"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnterpriseProxyServersAreAuthoritative(val)
        return nil
    }
    res["exemptAppLockerFiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionAppLockerFile() })
        if err != nil {
            return err
        }
        res := make([]WindowsInformationProtectionAppLockerFile, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsInformationProtectionAppLockerFile))
        }
        m.SetExemptAppLockerFiles(res)
        return nil
    }
    res["exemptApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionApp() })
        if err != nil {
            return err
        }
        res := make([]WindowsInformationProtectionApp, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsInformationProtectionApp))
        }
        m.SetExemptApps(res)
        return nil
    }
    res["iconsVisible"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIconsVisible(val)
        return nil
    }
    res["indexingEncryptedStoresOrItemsBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIndexingEncryptedStoresOrItemsBlocked(val)
        return nil
    }
    res["isAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAssigned(val)
        return nil
    }
    res["neutralDomainResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionResourceCollection() })
        if err != nil {
            return err
        }
        res := make([]WindowsInformationProtectionResourceCollection, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsInformationProtectionResourceCollection))
        }
        m.SetNeutralDomainResources(res)
        return nil
    }
    res["protectedAppLockerFiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionAppLockerFile() })
        if err != nil {
            return err
        }
        res := make([]WindowsInformationProtectionAppLockerFile, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsInformationProtectionAppLockerFile))
        }
        m.SetProtectedAppLockerFiles(res)
        return nil
    }
    res["protectedApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionApp() })
        if err != nil {
            return err
        }
        res := make([]WindowsInformationProtectionApp, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsInformationProtectionApp))
        }
        m.SetProtectedApps(res)
        return nil
    }
    res["protectionUnderLockConfigRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetProtectionUnderLockConfigRequired(val)
        return nil
    }
    res["revokeOnUnenrollDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRevokeOnUnenrollDisabled(val)
        return nil
    }
    res["rightsManagementServicesTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRightsManagementServicesTemplateId(val)
        return nil
    }
    res["smbAutoEncryptedFileExtensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionResourceCollection() })
        if err != nil {
            return err
        }
        res := make([]WindowsInformationProtectionResourceCollection, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsInformationProtectionResourceCollection))
        }
        m.SetSmbAutoEncryptedFileExtensions(res)
        return nil
    }
    return res
}
func (m *WindowsInformationProtection) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WindowsInformationProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ManagedAppPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("azureRightsManagementServicesAllowed", m.GetAzureRightsManagementServicesAllowed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("dataRecoveryCertificate", m.GetDataRecoveryCertificate())
        if err != nil {
            return err
        }
    }
    if m.GetEnforcementLevel() != nil {
        cast := m.GetEnforcementLevel().String()
        err = writer.WriteStringValue("enforcementLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("enterpriseDomain", m.GetEnterpriseDomain())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEnterpriseInternalProxyServers()))
        for i, v := range m.GetEnterpriseInternalProxyServers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("enterpriseInternalProxyServers", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEnterpriseIPRanges()))
        for i, v := range m.GetEnterpriseIPRanges() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("enterpriseIPRanges", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enterpriseIPRangesAreAuthoritative", m.GetEnterpriseIPRangesAreAuthoritative())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEnterpriseNetworkDomainNames()))
        for i, v := range m.GetEnterpriseNetworkDomainNames() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("enterpriseNetworkDomainNames", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEnterpriseProtectedDomainNames()))
        for i, v := range m.GetEnterpriseProtectedDomainNames() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("enterpriseProtectedDomainNames", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEnterpriseProxiedDomains()))
        for i, v := range m.GetEnterpriseProxiedDomains() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("enterpriseProxiedDomains", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEnterpriseProxyServers()))
        for i, v := range m.GetEnterpriseProxyServers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("enterpriseProxyServers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enterpriseProxyServersAreAuthoritative", m.GetEnterpriseProxyServersAreAuthoritative())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExemptAppLockerFiles()))
        for i, v := range m.GetExemptAppLockerFiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("exemptAppLockerFiles", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExemptApps()))
        for i, v := range m.GetExemptApps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("exemptApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iconsVisible", m.GetIconsVisible())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("indexingEncryptedStoresOrItemsBlocked", m.GetIndexingEncryptedStoresOrItemsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAssigned", m.GetIsAssigned())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNeutralDomainResources()))
        for i, v := range m.GetNeutralDomainResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("neutralDomainResources", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProtectedAppLockerFiles()))
        for i, v := range m.GetProtectedAppLockerFiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("protectedAppLockerFiles", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProtectedApps()))
        for i, v := range m.GetProtectedApps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("protectedApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("protectionUnderLockConfigRequired", m.GetProtectionUnderLockConfigRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("revokeOnUnenrollDisabled", m.GetRevokeOnUnenrollDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("rightsManagementServicesTemplateId", m.GetRightsManagementServicesTemplateId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSmbAutoEncryptedFileExtensions()))
        for i, v := range m.GetSmbAutoEncryptedFileExtensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("smbAutoEncryptedFileExtensions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignments property value. Navigation property to list of security groups targeted for policy.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *WindowsInformationProtection) SetAssignments(value []TargetedManagedAppPolicyAssignment)() {
    m.assignments = value
}
// Sets the azureRightsManagementServicesAllowed property value. Specifies whether to allow Azure RMS encryption for WIP
// Parameters:
//  - value : Value to set for the azureRightsManagementServicesAllowed property.
func (m *WindowsInformationProtection) SetAzureRightsManagementServicesAllowed(value *bool)() {
    m.azureRightsManagementServicesAllowed = value
}
// Sets the dataRecoveryCertificate property value. Specifies a recovery certificate that can be used for data recovery of encrypted files. This is the same as the data recovery agent(DRA) certificate for encrypting file system(EFS)
// Parameters:
//  - value : Value to set for the dataRecoveryCertificate property.
func (m *WindowsInformationProtection) SetDataRecoveryCertificate(value *WindowsInformationProtectionDataRecoveryCertificate)() {
    m.dataRecoveryCertificate = value
}
// Sets the enforcementLevel property value. WIP enforcement level.See the Enum definition for supported values. Possible values are: noProtection, encryptAndAuditOnly, encryptAuditAndPrompt, encryptAuditAndBlock.
// Parameters:
//  - value : Value to set for the enforcementLevel property.
func (m *WindowsInformationProtection) SetEnforcementLevel(value *WindowsInformationProtectionEnforcementLevel)() {
    m.enforcementLevel = value
}
// Sets the enterpriseDomain property value. Primary enterprise domain
// Parameters:
//  - value : Value to set for the enterpriseDomain property.
func (m *WindowsInformationProtection) SetEnterpriseDomain(value *string)() {
    m.enterpriseDomain = value
}
// Sets the enterpriseInternalProxyServers property value. This is the comma-separated list of internal proxy servers. For example, '157.54.14.28, 157.54.11.118, 10.202.14.167, 157.53.14.163, 157.69.210.59'. These proxies have been configured by the admin to connect to specific resources on the Internet. They are considered to be enterprise network locations. The proxies are only leveraged in configuring the EnterpriseProxiedDomains policy to force traffic to the matched domains through these proxies
// Parameters:
//  - value : Value to set for the enterpriseInternalProxyServers property.
func (m *WindowsInformationProtection) SetEnterpriseInternalProxyServers(value []WindowsInformationProtectionResourceCollection)() {
    m.enterpriseInternalProxyServers = value
}
// Sets the enterpriseIPRanges property value. Sets the enterprise IP ranges that define the computers in the enterprise network. Data that comes from those computers will be considered part of the enterprise and protected. These locations will be considered a safe destination for enterprise data to be shared to
// Parameters:
//  - value : Value to set for the enterpriseIPRanges property.
func (m *WindowsInformationProtection) SetEnterpriseIPRanges(value []WindowsInformationProtectionIPRangeCollection)() {
    m.enterpriseIPRanges = value
}
// Sets the enterpriseIPRangesAreAuthoritative property value. Boolean value that tells the client to accept the configured list and not to use heuristics to attempt to find other subnets. Default is false
// Parameters:
//  - value : Value to set for the enterpriseIPRangesAreAuthoritative property.
func (m *WindowsInformationProtection) SetEnterpriseIPRangesAreAuthoritative(value *bool)() {
    m.enterpriseIPRangesAreAuthoritative = value
}
// Sets the enterpriseNetworkDomainNames property value. This is the list of domains that comprise the boundaries of the enterprise. Data from one of these domains that is sent to a device will be considered enterprise data and protected These locations will be considered a safe destination for enterprise data to be shared to
// Parameters:
//  - value : Value to set for the enterpriseNetworkDomainNames property.
func (m *WindowsInformationProtection) SetEnterpriseNetworkDomainNames(value []WindowsInformationProtectionResourceCollection)() {
    m.enterpriseNetworkDomainNames = value
}
// Sets the enterpriseProtectedDomainNames property value. List of enterprise domains to be protected
// Parameters:
//  - value : Value to set for the enterpriseProtectedDomainNames property.
func (m *WindowsInformationProtection) SetEnterpriseProtectedDomainNames(value []WindowsInformationProtectionResourceCollection)() {
    m.enterpriseProtectedDomainNames = value
}
// Sets the enterpriseProxiedDomains property value. Contains a list of Enterprise resource domains hosted in the cloud that need to be protected. Connections to these resources are considered enterprise data. If a proxy is paired with a cloud resource, traffic to the cloud resource will be routed through the enterprise network via the denoted proxy server (on Port 80). A proxy server used for this purpose must also be configured using the EnterpriseInternalProxyServers policy
// Parameters:
//  - value : Value to set for the enterpriseProxiedDomains property.
func (m *WindowsInformationProtection) SetEnterpriseProxiedDomains(value []WindowsInformationProtectionProxiedDomainCollection)() {
    m.enterpriseProxiedDomains = value
}
// Sets the enterpriseProxyServers property value. This is a list of proxy servers. Any server not on this list is considered non-enterprise
// Parameters:
//  - value : Value to set for the enterpriseProxyServers property.
func (m *WindowsInformationProtection) SetEnterpriseProxyServers(value []WindowsInformationProtectionResourceCollection)() {
    m.enterpriseProxyServers = value
}
// Sets the enterpriseProxyServersAreAuthoritative property value. Boolean value that tells the client to accept the configured list of proxies and not try to detect other work proxies. Default is false
// Parameters:
//  - value : Value to set for the enterpriseProxyServersAreAuthoritative property.
func (m *WindowsInformationProtection) SetEnterpriseProxyServersAreAuthoritative(value *bool)() {
    m.enterpriseProxyServersAreAuthoritative = value
}
// Sets the exemptAppLockerFiles property value. Another way to input exempt apps through xml files
// Parameters:
//  - value : Value to set for the exemptAppLockerFiles property.
func (m *WindowsInformationProtection) SetExemptAppLockerFiles(value []WindowsInformationProtectionAppLockerFile)() {
    m.exemptAppLockerFiles = value
}
// Sets the exemptApps property value. Exempt applications can also access enterprise data, but the data handled by those applications are not protected. This is because some critical enterprise applications may have compatibility problems with encrypted data.
// Parameters:
//  - value : Value to set for the exemptApps property.
func (m *WindowsInformationProtection) SetExemptApps(value []WindowsInformationProtectionApp)() {
    m.exemptApps = value
}
// Sets the iconsVisible property value. Determines whether overlays are added to icons for WIP protected files in Explorer and enterprise only app tiles in the Start menu. Starting in Windows 10, version 1703 this setting also configures the visibility of the WIP icon in the title bar of a WIP-protected app
// Parameters:
//  - value : Value to set for the iconsVisible property.
func (m *WindowsInformationProtection) SetIconsVisible(value *bool)() {
    m.iconsVisible = value
}
// Sets the indexingEncryptedStoresOrItemsBlocked property value. This switch is for the Windows Search Indexer, to allow or disallow indexing of items
// Parameters:
//  - value : Value to set for the indexingEncryptedStoresOrItemsBlocked property.
func (m *WindowsInformationProtection) SetIndexingEncryptedStoresOrItemsBlocked(value *bool)() {
    m.indexingEncryptedStoresOrItemsBlocked = value
}
// Sets the isAssigned property value. Indicates if the policy is deployed to any inclusion groups or not.
// Parameters:
//  - value : Value to set for the isAssigned property.
func (m *WindowsInformationProtection) SetIsAssigned(value *bool)() {
    m.isAssigned = value
}
// Sets the neutralDomainResources property value. List of domain names that can used for work or personal resource
// Parameters:
//  - value : Value to set for the neutralDomainResources property.
func (m *WindowsInformationProtection) SetNeutralDomainResources(value []WindowsInformationProtectionResourceCollection)() {
    m.neutralDomainResources = value
}
// Sets the protectedAppLockerFiles property value. Another way to input protected apps through xml files
// Parameters:
//  - value : Value to set for the protectedAppLockerFiles property.
func (m *WindowsInformationProtection) SetProtectedAppLockerFiles(value []WindowsInformationProtectionAppLockerFile)() {
    m.protectedAppLockerFiles = value
}
// Sets the protectedApps property value. Protected applications can access enterprise data and the data handled by those applications are protected with encryption
// Parameters:
//  - value : Value to set for the protectedApps property.
func (m *WindowsInformationProtection) SetProtectedApps(value []WindowsInformationProtectionApp)() {
    m.protectedApps = value
}
// Sets the protectionUnderLockConfigRequired property value. Specifies whether the protection under lock feature (also known as encrypt under pin) should be configured
// Parameters:
//  - value : Value to set for the protectionUnderLockConfigRequired property.
func (m *WindowsInformationProtection) SetProtectionUnderLockConfigRequired(value *bool)() {
    m.protectionUnderLockConfigRequired = value
}
// Sets the revokeOnUnenrollDisabled property value. This policy controls whether to revoke the WIP keys when a device unenrolls from the management service. If set to 1 (Don't revoke keys), the keys will not be revoked and the user will continue to have access to protected files after unenrollment. If the keys are not revoked, there will be no revoked file cleanup subsequently.
// Parameters:
//  - value : Value to set for the revokeOnUnenrollDisabled property.
func (m *WindowsInformationProtection) SetRevokeOnUnenrollDisabled(value *bool)() {
    m.revokeOnUnenrollDisabled = value
}
// Sets the rightsManagementServicesTemplateId property value. TemplateID GUID to use for RMS encryption. The RMS template allows the IT admin to configure the details about who has access to RMS-protected file and how long they have access
// Parameters:
//  - value : Value to set for the rightsManagementServicesTemplateId property.
func (m *WindowsInformationProtection) SetRightsManagementServicesTemplateId(value *string)() {
    m.rightsManagementServicesTemplateId = value
}
// Sets the smbAutoEncryptedFileExtensions property value. Specifies a list of file extensions, so that files with these extensions are encrypted when copying from an SMB share within the corporate boundary
// Parameters:
//  - value : Value to set for the smbAutoEncryptedFileExtensions property.
func (m *WindowsInformationProtection) SetSmbAutoEncryptedFileExtensions(value []WindowsInformationProtectionResourceCollection)() {
    m.smbAutoEncryptedFileExtensions = value
}
