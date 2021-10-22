package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsInformationProtection struct {
    ManagedAppPolicy
    assignments []TargetedManagedAppPolicyAssignment;
    azureRightsManagementServicesAllowed *bool;
    dataRecoveryCertificate *WindowsInformationProtectionDataRecoveryCertificate;
    enforcementLevel *WindowsInformationProtectionEnforcementLevel;
    enterpriseDomain *string;
    enterpriseInternalProxyServers []WindowsInformationProtectionResourceCollection;
    enterpriseIPRanges []WindowsInformationProtectionIPRangeCollection;
    enterpriseIPRangesAreAuthoritative *bool;
    enterpriseNetworkDomainNames []WindowsInformationProtectionResourceCollection;
    enterpriseProtectedDomainNames []WindowsInformationProtectionResourceCollection;
    enterpriseProxiedDomains []WindowsInformationProtectionProxiedDomainCollection;
    enterpriseProxyServers []WindowsInformationProtectionResourceCollection;
    enterpriseProxyServersAreAuthoritative *bool;
    exemptAppLockerFiles []WindowsInformationProtectionAppLockerFile;
    exemptApps []WindowsInformationProtectionApp;
    iconsVisible *bool;
    indexingEncryptedStoresOrItemsBlocked *bool;
    isAssigned *bool;
    neutralDomainResources []WindowsInformationProtectionResourceCollection;
    protectedAppLockerFiles []WindowsInformationProtectionAppLockerFile;
    protectedApps []WindowsInformationProtectionApp;
    protectionUnderLockConfigRequired *bool;
    revokeOnUnenrollDisabled *bool;
    rightsManagementServicesTemplateId *string;
    smbAutoEncryptedFileExtensions []WindowsInformationProtectionResourceCollection;
}
func NewWindowsInformationProtection()(*WindowsInformationProtection) {
    m := &WindowsInformationProtection{
        ManagedAppPolicy: *NewManagedAppPolicy(),
    }
    return m
}
func (m *WindowsInformationProtection) GetAssignments()([]TargetedManagedAppPolicyAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
func (m *WindowsInformationProtection) GetAzureRightsManagementServicesAllowed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.azureRightsManagementServicesAllowed
    }
}
func (m *WindowsInformationProtection) GetDataRecoveryCertificate()(*WindowsInformationProtectionDataRecoveryCertificate) {
    if m == nil {
        return nil
    } else {
        return m.dataRecoveryCertificate
    }
}
func (m *WindowsInformationProtection) GetEnforcementLevel()(*WindowsInformationProtectionEnforcementLevel) {
    if m == nil {
        return nil
    } else {
        return m.enforcementLevel
    }
}
func (m *WindowsInformationProtection) GetEnterpriseDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseDomain
    }
}
func (m *WindowsInformationProtection) GetEnterpriseInternalProxyServers()([]WindowsInformationProtectionResourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseInternalProxyServers
    }
}
func (m *WindowsInformationProtection) GetEnterpriseIPRanges()([]WindowsInformationProtectionIPRangeCollection) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseIPRanges
    }
}
func (m *WindowsInformationProtection) GetEnterpriseIPRangesAreAuthoritative()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseIPRangesAreAuthoritative
    }
}
func (m *WindowsInformationProtection) GetEnterpriseNetworkDomainNames()([]WindowsInformationProtectionResourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseNetworkDomainNames
    }
}
func (m *WindowsInformationProtection) GetEnterpriseProtectedDomainNames()([]WindowsInformationProtectionResourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseProtectedDomainNames
    }
}
func (m *WindowsInformationProtection) GetEnterpriseProxiedDomains()([]WindowsInformationProtectionProxiedDomainCollection) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseProxiedDomains
    }
}
func (m *WindowsInformationProtection) GetEnterpriseProxyServers()([]WindowsInformationProtectionResourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseProxyServers
    }
}
func (m *WindowsInformationProtection) GetEnterpriseProxyServersAreAuthoritative()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseProxyServersAreAuthoritative
    }
}
func (m *WindowsInformationProtection) GetExemptAppLockerFiles()([]WindowsInformationProtectionAppLockerFile) {
    if m == nil {
        return nil
    } else {
        return m.exemptAppLockerFiles
    }
}
func (m *WindowsInformationProtection) GetExemptApps()([]WindowsInformationProtectionApp) {
    if m == nil {
        return nil
    } else {
        return m.exemptApps
    }
}
func (m *WindowsInformationProtection) GetIconsVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iconsVisible
    }
}
func (m *WindowsInformationProtection) GetIndexingEncryptedStoresOrItemsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.indexingEncryptedStoresOrItemsBlocked
    }
}
func (m *WindowsInformationProtection) GetIsAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAssigned
    }
}
func (m *WindowsInformationProtection) GetNeutralDomainResources()([]WindowsInformationProtectionResourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.neutralDomainResources
    }
}
func (m *WindowsInformationProtection) GetProtectedAppLockerFiles()([]WindowsInformationProtectionAppLockerFile) {
    if m == nil {
        return nil
    } else {
        return m.protectedAppLockerFiles
    }
}
func (m *WindowsInformationProtection) GetProtectedApps()([]WindowsInformationProtectionApp) {
    if m == nil {
        return nil
    } else {
        return m.protectedApps
    }
}
func (m *WindowsInformationProtection) GetProtectionUnderLockConfigRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.protectionUnderLockConfigRequired
    }
}
func (m *WindowsInformationProtection) GetRevokeOnUnenrollDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.revokeOnUnenrollDisabled
    }
}
func (m *WindowsInformationProtection) GetRightsManagementServicesTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rightsManagementServicesTemplateId
    }
}
func (m *WindowsInformationProtection) GetSmbAutoEncryptedFileExtensions()([]WindowsInformationProtectionResourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.smbAutoEncryptedFileExtensions
    }
}
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
func (m *WindowsInformationProtection) SetAssignments(value []TargetedManagedAppPolicyAssignment)() {
    m.assignments = value
}
func (m *WindowsInformationProtection) SetAzureRightsManagementServicesAllowed(value *bool)() {
    m.azureRightsManagementServicesAllowed = value
}
func (m *WindowsInformationProtection) SetDataRecoveryCertificate(value *WindowsInformationProtectionDataRecoveryCertificate)() {
    m.dataRecoveryCertificate = value
}
func (m *WindowsInformationProtection) SetEnforcementLevel(value *WindowsInformationProtectionEnforcementLevel)() {
    m.enforcementLevel = value
}
func (m *WindowsInformationProtection) SetEnterpriseDomain(value *string)() {
    m.enterpriseDomain = value
}
func (m *WindowsInformationProtection) SetEnterpriseInternalProxyServers(value []WindowsInformationProtectionResourceCollection)() {
    m.enterpriseInternalProxyServers = value
}
func (m *WindowsInformationProtection) SetEnterpriseIPRanges(value []WindowsInformationProtectionIPRangeCollection)() {
    m.enterpriseIPRanges = value
}
func (m *WindowsInformationProtection) SetEnterpriseIPRangesAreAuthoritative(value *bool)() {
    m.enterpriseIPRangesAreAuthoritative = value
}
func (m *WindowsInformationProtection) SetEnterpriseNetworkDomainNames(value []WindowsInformationProtectionResourceCollection)() {
    m.enterpriseNetworkDomainNames = value
}
func (m *WindowsInformationProtection) SetEnterpriseProtectedDomainNames(value []WindowsInformationProtectionResourceCollection)() {
    m.enterpriseProtectedDomainNames = value
}
func (m *WindowsInformationProtection) SetEnterpriseProxiedDomains(value []WindowsInformationProtectionProxiedDomainCollection)() {
    m.enterpriseProxiedDomains = value
}
func (m *WindowsInformationProtection) SetEnterpriseProxyServers(value []WindowsInformationProtectionResourceCollection)() {
    m.enterpriseProxyServers = value
}
func (m *WindowsInformationProtection) SetEnterpriseProxyServersAreAuthoritative(value *bool)() {
    m.enterpriseProxyServersAreAuthoritative = value
}
func (m *WindowsInformationProtection) SetExemptAppLockerFiles(value []WindowsInformationProtectionAppLockerFile)() {
    m.exemptAppLockerFiles = value
}
func (m *WindowsInformationProtection) SetExemptApps(value []WindowsInformationProtectionApp)() {
    m.exemptApps = value
}
func (m *WindowsInformationProtection) SetIconsVisible(value *bool)() {
    m.iconsVisible = value
}
func (m *WindowsInformationProtection) SetIndexingEncryptedStoresOrItemsBlocked(value *bool)() {
    m.indexingEncryptedStoresOrItemsBlocked = value
}
func (m *WindowsInformationProtection) SetIsAssigned(value *bool)() {
    m.isAssigned = value
}
func (m *WindowsInformationProtection) SetNeutralDomainResources(value []WindowsInformationProtectionResourceCollection)() {
    m.neutralDomainResources = value
}
func (m *WindowsInformationProtection) SetProtectedAppLockerFiles(value []WindowsInformationProtectionAppLockerFile)() {
    m.protectedAppLockerFiles = value
}
func (m *WindowsInformationProtection) SetProtectedApps(value []WindowsInformationProtectionApp)() {
    m.protectedApps = value
}
func (m *WindowsInformationProtection) SetProtectionUnderLockConfigRequired(value *bool)() {
    m.protectionUnderLockConfigRequired = value
}
func (m *WindowsInformationProtection) SetRevokeOnUnenrollDisabled(value *bool)() {
    m.revokeOnUnenrollDisabled = value
}
func (m *WindowsInformationProtection) SetRightsManagementServicesTemplateId(value *string)() {
    m.rightsManagementServicesTemplateId = value
}
func (m *WindowsInformationProtection) SetSmbAutoEncryptedFileExtensions(value []WindowsInformationProtectionResourceCollection)() {
    m.smbAutoEncryptedFileExtensions = value
}
