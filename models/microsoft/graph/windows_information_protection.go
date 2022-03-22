package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsInformationProtection 
type WindowsInformationProtection struct {
    ManagedAppPolicy
    // Navigation property to list of security groups targeted for policy.
    assignments []TargetedManagedAppPolicyAssignmentable;
    // Specifies whether to allow Azure RMS encryption for WIP
    azureRightsManagementServicesAllowed *bool;
    // Specifies a recovery certificate that can be used for data recovery of encrypted files. This is the same as the data recovery agent(DRA) certificate for encrypting file system(EFS)
    dataRecoveryCertificate WindowsInformationProtectionDataRecoveryCertificateable;
    // WIP enforcement level.See the Enum definition for supported values. Possible values are: noProtection, encryptAndAuditOnly, encryptAuditAndPrompt, encryptAuditAndBlock.
    enforcementLevel *WindowsInformationProtectionEnforcementLevel;
    // Primary enterprise domain
    enterpriseDomain *string;
    // This is the comma-separated list of internal proxy servers. For example, '157.54.14.28, 157.54.11.118, 10.202.14.167, 157.53.14.163, 157.69.210.59'. These proxies have been configured by the admin to connect to specific resources on the Internet. They are considered to be enterprise network locations. The proxies are only leveraged in configuring the EnterpriseProxiedDomains policy to force traffic to the matched domains through these proxies
    enterpriseInternalProxyServers []WindowsInformationProtectionResourceCollectionable;
    // Sets the enterprise IP ranges that define the computers in the enterprise network. Data that comes from those computers will be considered part of the enterprise and protected. These locations will be considered a safe destination for enterprise data to be shared to
    enterpriseIPRanges []WindowsInformationProtectionIPRangeCollectionable;
    // Boolean value that tells the client to accept the configured list and not to use heuristics to attempt to find other subnets. Default is false
    enterpriseIPRangesAreAuthoritative *bool;
    // This is the list of domains that comprise the boundaries of the enterprise. Data from one of these domains that is sent to a device will be considered enterprise data and protected These locations will be considered a safe destination for enterprise data to be shared to
    enterpriseNetworkDomainNames []WindowsInformationProtectionResourceCollectionable;
    // List of enterprise domains to be protected
    enterpriseProtectedDomainNames []WindowsInformationProtectionResourceCollectionable;
    // Contains a list of Enterprise resource domains hosted in the cloud that need to be protected. Connections to these resources are considered enterprise data. If a proxy is paired with a cloud resource, traffic to the cloud resource will be routed through the enterprise network via the denoted proxy server (on Port 80). A proxy server used for this purpose must also be configured using the EnterpriseInternalProxyServers policy
    enterpriseProxiedDomains []WindowsInformationProtectionProxiedDomainCollectionable;
    // This is a list of proxy servers. Any server not on this list is considered non-enterprise
    enterpriseProxyServers []WindowsInformationProtectionResourceCollectionable;
    // Boolean value that tells the client to accept the configured list of proxies and not try to detect other work proxies. Default is false
    enterpriseProxyServersAreAuthoritative *bool;
    // Another way to input exempt apps through xml files
    exemptAppLockerFiles []WindowsInformationProtectionAppLockerFileable;
    // Exempt applications can also access enterprise data, but the data handled by those applications are not protected. This is because some critical enterprise applications may have compatibility problems with encrypted data.
    exemptApps []WindowsInformationProtectionAppable;
    // Determines whether overlays are added to icons for WIP protected files in Explorer and enterprise only app tiles in the Start menu. Starting in Windows 10, version 1703 this setting also configures the visibility of the WIP icon in the title bar of a WIP-protected app
    iconsVisible *bool;
    // This switch is for the Windows Search Indexer, to allow or disallow indexing of items
    indexingEncryptedStoresOrItemsBlocked *bool;
    // Indicates if the policy is deployed to any inclusion groups or not.
    isAssigned *bool;
    // List of domain names that can used for work or personal resource
    neutralDomainResources []WindowsInformationProtectionResourceCollectionable;
    // Another way to input protected apps through xml files
    protectedAppLockerFiles []WindowsInformationProtectionAppLockerFileable;
    // Protected applications can access enterprise data and the data handled by those applications are protected with encryption
    protectedApps []WindowsInformationProtectionAppable;
    // Specifies whether the protection under lock feature (also known as encrypt under pin) should be configured
    protectionUnderLockConfigRequired *bool;
    // This policy controls whether to revoke the WIP keys when a device unenrolls from the management service. If set to 1 (Don't revoke keys), the keys will not be revoked and the user will continue to have access to protected files after unenrollment. If the keys are not revoked, there will be no revoked file cleanup subsequently.
    revokeOnUnenrollDisabled *bool;
    // TemplateID GUID to use for RMS encryption. The RMS template allows the IT admin to configure the details about who has access to RMS-protected file and how long they have access
    rightsManagementServicesTemplateId *string;
    // Specifies a list of file extensions, so that files with these extensions are encrypted when copying from an SMB share within the corporate boundary
    smbAutoEncryptedFileExtensions []WindowsInformationProtectionResourceCollectionable;
}
// NewWindowsInformationProtection instantiates a new windowsInformationProtection and sets the default values.
func NewWindowsInformationProtection()(*WindowsInformationProtection) {
    m := &WindowsInformationProtection{
        ManagedAppPolicy: *NewManagedAppPolicy(),
    }
    return m
}
// CreateWindowsInformationProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsInformationProtectionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWindowsInformationProtection(), nil
}
// GetAssignments gets the assignments property value. Navigation property to list of security groups targeted for policy.
func (m *WindowsInformationProtection) GetAssignments()([]TargetedManagedAppPolicyAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetAzureRightsManagementServicesAllowed gets the azureRightsManagementServicesAllowed property value. Specifies whether to allow Azure RMS encryption for WIP
func (m *WindowsInformationProtection) GetAzureRightsManagementServicesAllowed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.azureRightsManagementServicesAllowed
    }
}
// GetDataRecoveryCertificate gets the dataRecoveryCertificate property value. Specifies a recovery certificate that can be used for data recovery of encrypted files. This is the same as the data recovery agent(DRA) certificate for encrypting file system(EFS)
func (m *WindowsInformationProtection) GetDataRecoveryCertificate()(WindowsInformationProtectionDataRecoveryCertificateable) {
    if m == nil {
        return nil
    } else {
        return m.dataRecoveryCertificate
    }
}
// GetEnforcementLevel gets the enforcementLevel property value. WIP enforcement level.See the Enum definition for supported values. Possible values are: noProtection, encryptAndAuditOnly, encryptAuditAndPrompt, encryptAuditAndBlock.
func (m *WindowsInformationProtection) GetEnforcementLevel()(*WindowsInformationProtectionEnforcementLevel) {
    if m == nil {
        return nil
    } else {
        return m.enforcementLevel
    }
}
// GetEnterpriseDomain gets the enterpriseDomain property value. Primary enterprise domain
func (m *WindowsInformationProtection) GetEnterpriseDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseDomain
    }
}
// GetEnterpriseInternalProxyServers gets the enterpriseInternalProxyServers property value. This is the comma-separated list of internal proxy servers. For example, '157.54.14.28, 157.54.11.118, 10.202.14.167, 157.53.14.163, 157.69.210.59'. These proxies have been configured by the admin to connect to specific resources on the Internet. They are considered to be enterprise network locations. The proxies are only leveraged in configuring the EnterpriseProxiedDomains policy to force traffic to the matched domains through these proxies
func (m *WindowsInformationProtection) GetEnterpriseInternalProxyServers()([]WindowsInformationProtectionResourceCollectionable) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseInternalProxyServers
    }
}
// GetEnterpriseIPRanges gets the enterpriseIPRanges property value. Sets the enterprise IP ranges that define the computers in the enterprise network. Data that comes from those computers will be considered part of the enterprise and protected. These locations will be considered a safe destination for enterprise data to be shared to
func (m *WindowsInformationProtection) GetEnterpriseIPRanges()([]WindowsInformationProtectionIPRangeCollectionable) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseIPRanges
    }
}
// GetEnterpriseIPRangesAreAuthoritative gets the enterpriseIPRangesAreAuthoritative property value. Boolean value that tells the client to accept the configured list and not to use heuristics to attempt to find other subnets. Default is false
func (m *WindowsInformationProtection) GetEnterpriseIPRangesAreAuthoritative()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseIPRangesAreAuthoritative
    }
}
// GetEnterpriseNetworkDomainNames gets the enterpriseNetworkDomainNames property value. This is the list of domains that comprise the boundaries of the enterprise. Data from one of these domains that is sent to a device will be considered enterprise data and protected These locations will be considered a safe destination for enterprise data to be shared to
func (m *WindowsInformationProtection) GetEnterpriseNetworkDomainNames()([]WindowsInformationProtectionResourceCollectionable) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseNetworkDomainNames
    }
}
// GetEnterpriseProtectedDomainNames gets the enterpriseProtectedDomainNames property value. List of enterprise domains to be protected
func (m *WindowsInformationProtection) GetEnterpriseProtectedDomainNames()([]WindowsInformationProtectionResourceCollectionable) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseProtectedDomainNames
    }
}
// GetEnterpriseProxiedDomains gets the enterpriseProxiedDomains property value. Contains a list of Enterprise resource domains hosted in the cloud that need to be protected. Connections to these resources are considered enterprise data. If a proxy is paired with a cloud resource, traffic to the cloud resource will be routed through the enterprise network via the denoted proxy server (on Port 80). A proxy server used for this purpose must also be configured using the EnterpriseInternalProxyServers policy
func (m *WindowsInformationProtection) GetEnterpriseProxiedDomains()([]WindowsInformationProtectionProxiedDomainCollectionable) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseProxiedDomains
    }
}
// GetEnterpriseProxyServers gets the enterpriseProxyServers property value. This is a list of proxy servers. Any server not on this list is considered non-enterprise
func (m *WindowsInformationProtection) GetEnterpriseProxyServers()([]WindowsInformationProtectionResourceCollectionable) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseProxyServers
    }
}
// GetEnterpriseProxyServersAreAuthoritative gets the enterpriseProxyServersAreAuthoritative property value. Boolean value that tells the client to accept the configured list of proxies and not try to detect other work proxies. Default is false
func (m *WindowsInformationProtection) GetEnterpriseProxyServersAreAuthoritative()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseProxyServersAreAuthoritative
    }
}
// GetExemptAppLockerFiles gets the exemptAppLockerFiles property value. Another way to input exempt apps through xml files
func (m *WindowsInformationProtection) GetExemptAppLockerFiles()([]WindowsInformationProtectionAppLockerFileable) {
    if m == nil {
        return nil
    } else {
        return m.exemptAppLockerFiles
    }
}
// GetExemptApps gets the exemptApps property value. Exempt applications can also access enterprise data, but the data handled by those applications are not protected. This is because some critical enterprise applications may have compatibility problems with encrypted data.
func (m *WindowsInformationProtection) GetExemptApps()([]WindowsInformationProtectionAppable) {
    if m == nil {
        return nil
    } else {
        return m.exemptApps
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ManagedAppPolicy.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTargetedManagedAppPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TargetedManagedAppPolicyAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(TargetedManagedAppPolicyAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["azureRightsManagementServicesAllowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureRightsManagementServicesAllowed(val)
        }
        return nil
    }
    res["dataRecoveryCertificate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsInformationProtectionDataRecoveryCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataRecoveryCertificate(val.(WindowsInformationProtectionDataRecoveryCertificateable))
        }
        return nil
    }
    res["enforcementLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsInformationProtectionEnforcementLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforcementLevel(val.(*WindowsInformationProtectionEnforcementLevel))
        }
        return nil
    }
    res["enterpriseDomain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseDomain(val)
        }
        return nil
    }
    res["enterpriseInternalProxyServers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionResourceCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionResourceCollectionable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionResourceCollectionable)
            }
            m.SetEnterpriseInternalProxyServers(res)
        }
        return nil
    }
    res["enterpriseIPRanges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionIPRangeCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionIPRangeCollectionable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionIPRangeCollectionable)
            }
            m.SetEnterpriseIPRanges(res)
        }
        return nil
    }
    res["enterpriseIPRangesAreAuthoritative"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseIPRangesAreAuthoritative(val)
        }
        return nil
    }
    res["enterpriseNetworkDomainNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionResourceCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionResourceCollectionable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionResourceCollectionable)
            }
            m.SetEnterpriseNetworkDomainNames(res)
        }
        return nil
    }
    res["enterpriseProtectedDomainNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionResourceCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionResourceCollectionable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionResourceCollectionable)
            }
            m.SetEnterpriseProtectedDomainNames(res)
        }
        return nil
    }
    res["enterpriseProxiedDomains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionProxiedDomainCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionProxiedDomainCollectionable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionProxiedDomainCollectionable)
            }
            m.SetEnterpriseProxiedDomains(res)
        }
        return nil
    }
    res["enterpriseProxyServers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionResourceCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionResourceCollectionable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionResourceCollectionable)
            }
            m.SetEnterpriseProxyServers(res)
        }
        return nil
    }
    res["enterpriseProxyServersAreAuthoritative"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseProxyServersAreAuthoritative(val)
        }
        return nil
    }
    res["exemptAppLockerFiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionAppLockerFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionAppLockerFileable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionAppLockerFileable)
            }
            m.SetExemptAppLockerFiles(res)
        }
        return nil
    }
    res["exemptApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionAppable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionAppable)
            }
            m.SetExemptApps(res)
        }
        return nil
    }
    res["iconsVisible"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIconsVisible(val)
        }
        return nil
    }
    res["indexingEncryptedStoresOrItemsBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndexingEncryptedStoresOrItemsBlocked(val)
        }
        return nil
    }
    res["isAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAssigned(val)
        }
        return nil
    }
    res["neutralDomainResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionResourceCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionResourceCollectionable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionResourceCollectionable)
            }
            m.SetNeutralDomainResources(res)
        }
        return nil
    }
    res["protectedAppLockerFiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionAppLockerFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionAppLockerFileable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionAppLockerFileable)
            }
            m.SetProtectedAppLockerFiles(res)
        }
        return nil
    }
    res["protectedApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionAppable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionAppable)
            }
            m.SetProtectedApps(res)
        }
        return nil
    }
    res["protectionUnderLockConfigRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionUnderLockConfigRequired(val)
        }
        return nil
    }
    res["revokeOnUnenrollDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRevokeOnUnenrollDisabled(val)
        }
        return nil
    }
    res["rightsManagementServicesTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRightsManagementServicesTemplateId(val)
        }
        return nil
    }
    res["smbAutoEncryptedFileExtensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionResourceCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionResourceCollectionable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionResourceCollectionable)
            }
            m.SetSmbAutoEncryptedFileExtensions(res)
        }
        return nil
    }
    return res
}
// GetIconsVisible gets the iconsVisible property value. Determines whether overlays are added to icons for WIP protected files in Explorer and enterprise only app tiles in the Start menu. Starting in Windows 10, version 1703 this setting also configures the visibility of the WIP icon in the title bar of a WIP-protected app
func (m *WindowsInformationProtection) GetIconsVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iconsVisible
    }
}
// GetIndexingEncryptedStoresOrItemsBlocked gets the indexingEncryptedStoresOrItemsBlocked property value. This switch is for the Windows Search Indexer, to allow or disallow indexing of items
func (m *WindowsInformationProtection) GetIndexingEncryptedStoresOrItemsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.indexingEncryptedStoresOrItemsBlocked
    }
}
// GetIsAssigned gets the isAssigned property value. Indicates if the policy is deployed to any inclusion groups or not.
func (m *WindowsInformationProtection) GetIsAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAssigned
    }
}
// GetNeutralDomainResources gets the neutralDomainResources property value. List of domain names that can used for work or personal resource
func (m *WindowsInformationProtection) GetNeutralDomainResources()([]WindowsInformationProtectionResourceCollectionable) {
    if m == nil {
        return nil
    } else {
        return m.neutralDomainResources
    }
}
// GetProtectedAppLockerFiles gets the protectedAppLockerFiles property value. Another way to input protected apps through xml files
func (m *WindowsInformationProtection) GetProtectedAppLockerFiles()([]WindowsInformationProtectionAppLockerFileable) {
    if m == nil {
        return nil
    } else {
        return m.protectedAppLockerFiles
    }
}
// GetProtectedApps gets the protectedApps property value. Protected applications can access enterprise data and the data handled by those applications are protected with encryption
func (m *WindowsInformationProtection) GetProtectedApps()([]WindowsInformationProtectionAppable) {
    if m == nil {
        return nil
    } else {
        return m.protectedApps
    }
}
// GetProtectionUnderLockConfigRequired gets the protectionUnderLockConfigRequired property value. Specifies whether the protection under lock feature (also known as encrypt under pin) should be configured
func (m *WindowsInformationProtection) GetProtectionUnderLockConfigRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.protectionUnderLockConfigRequired
    }
}
// GetRevokeOnUnenrollDisabled gets the revokeOnUnenrollDisabled property value. This policy controls whether to revoke the WIP keys when a device unenrolls from the management service. If set to 1 (Don't revoke keys), the keys will not be revoked and the user will continue to have access to protected files after unenrollment. If the keys are not revoked, there will be no revoked file cleanup subsequently.
func (m *WindowsInformationProtection) GetRevokeOnUnenrollDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.revokeOnUnenrollDisabled
    }
}
// GetRightsManagementServicesTemplateId gets the rightsManagementServicesTemplateId property value. TemplateID GUID to use for RMS encryption. The RMS template allows the IT admin to configure the details about who has access to RMS-protected file and how long they have access
func (m *WindowsInformationProtection) GetRightsManagementServicesTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rightsManagementServicesTemplateId
    }
}
// GetSmbAutoEncryptedFileExtensions gets the smbAutoEncryptedFileExtensions property value. Specifies a list of file extensions, so that files with these extensions are encrypted when copying from an SMB share within the corporate boundary
func (m *WindowsInformationProtection) GetSmbAutoEncryptedFileExtensions()([]WindowsInformationProtectionResourceCollectionable) {
    if m == nil {
        return nil
    } else {
        return m.smbAutoEncryptedFileExtensions
    }
}
// Serialize serializes information the current object
func (m *WindowsInformationProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ManagedAppPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
        cast := (*m.GetEnforcementLevel()).String()
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
    if m.GetEnterpriseInternalProxyServers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEnterpriseInternalProxyServers()))
        for i, v := range m.GetEnterpriseInternalProxyServers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("enterpriseInternalProxyServers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnterpriseIPRanges() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEnterpriseIPRanges()))
        for i, v := range m.GetEnterpriseIPRanges() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetEnterpriseNetworkDomainNames() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEnterpriseNetworkDomainNames()))
        for i, v := range m.GetEnterpriseNetworkDomainNames() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("enterpriseNetworkDomainNames", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnterpriseProtectedDomainNames() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEnterpriseProtectedDomainNames()))
        for i, v := range m.GetEnterpriseProtectedDomainNames() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("enterpriseProtectedDomainNames", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnterpriseProxiedDomains() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEnterpriseProxiedDomains()))
        for i, v := range m.GetEnterpriseProxiedDomains() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("enterpriseProxiedDomains", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnterpriseProxyServers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEnterpriseProxyServers()))
        for i, v := range m.GetEnterpriseProxyServers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetExemptAppLockerFiles() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExemptAppLockerFiles()))
        for i, v := range m.GetExemptAppLockerFiles() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("exemptAppLockerFiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExemptApps() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExemptApps()))
        for i, v := range m.GetExemptApps() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetNeutralDomainResources() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNeutralDomainResources()))
        for i, v := range m.GetNeutralDomainResources() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("neutralDomainResources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProtectedAppLockerFiles() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProtectedAppLockerFiles()))
        for i, v := range m.GetProtectedAppLockerFiles() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("protectedAppLockerFiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProtectedApps() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProtectedApps()))
        for i, v := range m.GetProtectedApps() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetSmbAutoEncryptedFileExtensions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSmbAutoEncryptedFileExtensions()))
        for i, v := range m.GetSmbAutoEncryptedFileExtensions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("smbAutoEncryptedFileExtensions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. Navigation property to list of security groups targeted for policy.
func (m *WindowsInformationProtection) SetAssignments(value []TargetedManagedAppPolicyAssignmentable)() {
    if m != nil {
        m.assignments = value
    }
}
// SetAzureRightsManagementServicesAllowed sets the azureRightsManagementServicesAllowed property value. Specifies whether to allow Azure RMS encryption for WIP
func (m *WindowsInformationProtection) SetAzureRightsManagementServicesAllowed(value *bool)() {
    if m != nil {
        m.azureRightsManagementServicesAllowed = value
    }
}
// SetDataRecoveryCertificate sets the dataRecoveryCertificate property value. Specifies a recovery certificate that can be used for data recovery of encrypted files. This is the same as the data recovery agent(DRA) certificate for encrypting file system(EFS)
func (m *WindowsInformationProtection) SetDataRecoveryCertificate(value WindowsInformationProtectionDataRecoveryCertificateable)() {
    if m != nil {
        m.dataRecoveryCertificate = value
    }
}
// SetEnforcementLevel sets the enforcementLevel property value. WIP enforcement level.See the Enum definition for supported values. Possible values are: noProtection, encryptAndAuditOnly, encryptAuditAndPrompt, encryptAuditAndBlock.
func (m *WindowsInformationProtection) SetEnforcementLevel(value *WindowsInformationProtectionEnforcementLevel)() {
    if m != nil {
        m.enforcementLevel = value
    }
}
// SetEnterpriseDomain sets the enterpriseDomain property value. Primary enterprise domain
func (m *WindowsInformationProtection) SetEnterpriseDomain(value *string)() {
    if m != nil {
        m.enterpriseDomain = value
    }
}
// SetEnterpriseInternalProxyServers sets the enterpriseInternalProxyServers property value. This is the comma-separated list of internal proxy servers. For example, '157.54.14.28, 157.54.11.118, 10.202.14.167, 157.53.14.163, 157.69.210.59'. These proxies have been configured by the admin to connect to specific resources on the Internet. They are considered to be enterprise network locations. The proxies are only leveraged in configuring the EnterpriseProxiedDomains policy to force traffic to the matched domains through these proxies
func (m *WindowsInformationProtection) SetEnterpriseInternalProxyServers(value []WindowsInformationProtectionResourceCollectionable)() {
    if m != nil {
        m.enterpriseInternalProxyServers = value
    }
}
// SetEnterpriseIPRanges sets the enterpriseIPRanges property value. Sets the enterprise IP ranges that define the computers in the enterprise network. Data that comes from those computers will be considered part of the enterprise and protected. These locations will be considered a safe destination for enterprise data to be shared to
func (m *WindowsInformationProtection) SetEnterpriseIPRanges(value []WindowsInformationProtectionIPRangeCollectionable)() {
    if m != nil {
        m.enterpriseIPRanges = value
    }
}
// SetEnterpriseIPRangesAreAuthoritative sets the enterpriseIPRangesAreAuthoritative property value. Boolean value that tells the client to accept the configured list and not to use heuristics to attempt to find other subnets. Default is false
func (m *WindowsInformationProtection) SetEnterpriseIPRangesAreAuthoritative(value *bool)() {
    if m != nil {
        m.enterpriseIPRangesAreAuthoritative = value
    }
}
// SetEnterpriseNetworkDomainNames sets the enterpriseNetworkDomainNames property value. This is the list of domains that comprise the boundaries of the enterprise. Data from one of these domains that is sent to a device will be considered enterprise data and protected These locations will be considered a safe destination for enterprise data to be shared to
func (m *WindowsInformationProtection) SetEnterpriseNetworkDomainNames(value []WindowsInformationProtectionResourceCollectionable)() {
    if m != nil {
        m.enterpriseNetworkDomainNames = value
    }
}
// SetEnterpriseProtectedDomainNames sets the enterpriseProtectedDomainNames property value. List of enterprise domains to be protected
func (m *WindowsInformationProtection) SetEnterpriseProtectedDomainNames(value []WindowsInformationProtectionResourceCollectionable)() {
    if m != nil {
        m.enterpriseProtectedDomainNames = value
    }
}
// SetEnterpriseProxiedDomains sets the enterpriseProxiedDomains property value. Contains a list of Enterprise resource domains hosted in the cloud that need to be protected. Connections to these resources are considered enterprise data. If a proxy is paired with a cloud resource, traffic to the cloud resource will be routed through the enterprise network via the denoted proxy server (on Port 80). A proxy server used for this purpose must also be configured using the EnterpriseInternalProxyServers policy
func (m *WindowsInformationProtection) SetEnterpriseProxiedDomains(value []WindowsInformationProtectionProxiedDomainCollectionable)() {
    if m != nil {
        m.enterpriseProxiedDomains = value
    }
}
// SetEnterpriseProxyServers sets the enterpriseProxyServers property value. This is a list of proxy servers. Any server not on this list is considered non-enterprise
func (m *WindowsInformationProtection) SetEnterpriseProxyServers(value []WindowsInformationProtectionResourceCollectionable)() {
    if m != nil {
        m.enterpriseProxyServers = value
    }
}
// SetEnterpriseProxyServersAreAuthoritative sets the enterpriseProxyServersAreAuthoritative property value. Boolean value that tells the client to accept the configured list of proxies and not try to detect other work proxies. Default is false
func (m *WindowsInformationProtection) SetEnterpriseProxyServersAreAuthoritative(value *bool)() {
    if m != nil {
        m.enterpriseProxyServersAreAuthoritative = value
    }
}
// SetExemptAppLockerFiles sets the exemptAppLockerFiles property value. Another way to input exempt apps through xml files
func (m *WindowsInformationProtection) SetExemptAppLockerFiles(value []WindowsInformationProtectionAppLockerFileable)() {
    if m != nil {
        m.exemptAppLockerFiles = value
    }
}
// SetExemptApps sets the exemptApps property value. Exempt applications can also access enterprise data, but the data handled by those applications are not protected. This is because some critical enterprise applications may have compatibility problems with encrypted data.
func (m *WindowsInformationProtection) SetExemptApps(value []WindowsInformationProtectionAppable)() {
    if m != nil {
        m.exemptApps = value
    }
}
// SetIconsVisible sets the iconsVisible property value. Determines whether overlays are added to icons for WIP protected files in Explorer and enterprise only app tiles in the Start menu. Starting in Windows 10, version 1703 this setting also configures the visibility of the WIP icon in the title bar of a WIP-protected app
func (m *WindowsInformationProtection) SetIconsVisible(value *bool)() {
    if m != nil {
        m.iconsVisible = value
    }
}
// SetIndexingEncryptedStoresOrItemsBlocked sets the indexingEncryptedStoresOrItemsBlocked property value. This switch is for the Windows Search Indexer, to allow or disallow indexing of items
func (m *WindowsInformationProtection) SetIndexingEncryptedStoresOrItemsBlocked(value *bool)() {
    if m != nil {
        m.indexingEncryptedStoresOrItemsBlocked = value
    }
}
// SetIsAssigned sets the isAssigned property value. Indicates if the policy is deployed to any inclusion groups or not.
func (m *WindowsInformationProtection) SetIsAssigned(value *bool)() {
    if m != nil {
        m.isAssigned = value
    }
}
// SetNeutralDomainResources sets the neutralDomainResources property value. List of domain names that can used for work or personal resource
func (m *WindowsInformationProtection) SetNeutralDomainResources(value []WindowsInformationProtectionResourceCollectionable)() {
    if m != nil {
        m.neutralDomainResources = value
    }
}
// SetProtectedAppLockerFiles sets the protectedAppLockerFiles property value. Another way to input protected apps through xml files
func (m *WindowsInformationProtection) SetProtectedAppLockerFiles(value []WindowsInformationProtectionAppLockerFileable)() {
    if m != nil {
        m.protectedAppLockerFiles = value
    }
}
// SetProtectedApps sets the protectedApps property value. Protected applications can access enterprise data and the data handled by those applications are protected with encryption
func (m *WindowsInformationProtection) SetProtectedApps(value []WindowsInformationProtectionAppable)() {
    if m != nil {
        m.protectedApps = value
    }
}
// SetProtectionUnderLockConfigRequired sets the protectionUnderLockConfigRequired property value. Specifies whether the protection under lock feature (also known as encrypt under pin) should be configured
func (m *WindowsInformationProtection) SetProtectionUnderLockConfigRequired(value *bool)() {
    if m != nil {
        m.protectionUnderLockConfigRequired = value
    }
}
// SetRevokeOnUnenrollDisabled sets the revokeOnUnenrollDisabled property value. This policy controls whether to revoke the WIP keys when a device unenrolls from the management service. If set to 1 (Don't revoke keys), the keys will not be revoked and the user will continue to have access to protected files after unenrollment. If the keys are not revoked, there will be no revoked file cleanup subsequently.
func (m *WindowsInformationProtection) SetRevokeOnUnenrollDisabled(value *bool)() {
    if m != nil {
        m.revokeOnUnenrollDisabled = value
    }
}
// SetRightsManagementServicesTemplateId sets the rightsManagementServicesTemplateId property value. TemplateID GUID to use for RMS encryption. The RMS template allows the IT admin to configure the details about who has access to RMS-protected file and how long they have access
func (m *WindowsInformationProtection) SetRightsManagementServicesTemplateId(value *string)() {
    if m != nil {
        m.rightsManagementServicesTemplateId = value
    }
}
// SetSmbAutoEncryptedFileExtensions sets the smbAutoEncryptedFileExtensions property value. Specifies a list of file extensions, so that files with these extensions are encrypted when copying from an SMB share within the corporate boundary
func (m *WindowsInformationProtection) SetSmbAutoEncryptedFileExtensions(value []WindowsInformationProtectionResourceCollectionable)() {
    if m != nil {
        m.smbAutoEncryptedFileExtensions = value
    }
}
