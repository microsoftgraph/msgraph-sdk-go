package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Domain struct {
    Entity
    // Indicates the configured authentication type for the domain. The value is either Managed or Federated. Managed indicates a cloud managed domain where Azure AD performs user authentication. Federated indicates authentication is federated with an identity provider such as the tenant's on-premises Active Directory via Active Directory Federation Services. This property is read-only and is not nullable.
    authenticationType *string;
    // This property is always null except when the verify action is used. When the verify action is used, a domain entity is returned in the response. The availabilityStatus property of the domain entity in the response is either AvailableImmediately or EmailVerifiedDomainTakeoverScheduled.
    availabilityStatus *string;
    // Read-only, Nullable
    domainNameReferences []DirectoryObject;
    // The value of the property is false if the DNS record management of the domain has been delegated to Microsoft 365. Otherwise, the value is true. Not nullable
    isAdminManaged *bool;
    // true if this is the default domain that is used for user creation. There is only one default domain per company. Not nullable
    isDefault *bool;
    // true if this is the initial domain created by Microsoft Online Services (companyname.onmicrosoft.com). There is only one initial domain per company. Not nullable
    isInitial *bool;
    // true if the domain is a verified root domain. Otherwise, false if the domain is a subdomain or unverified. Not nullable
    isRoot *bool;
    // true if the domain has completed domain ownership verification. Not nullable
    isVerified *bool;
    // 
    manufacturer *string;
    // 
    model *string;
    // Specifies the number of days before a user receives notification that their password will expire. If the property is not set, a default value of 14 days will be used.
    passwordNotificationWindowInDays *int32;
    // Specifies the length of time that a password is valid before it must be changed. If the property is not set, a default value of 90 days will be used.
    passwordValidityPeriodInDays *int32;
    // DNS records the customer adds to the DNS zone file of the domain before the domain can be used by Microsoft Online services. Read-only, Nullable
    serviceConfigurationRecords []DomainDnsRecord;
    // Status of asynchronous operations scheduled for the domain.
    state *DomainState;
    // The capabilities assigned to the domain. Can include 0, 1 or more of following values: Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune. The values which you can add/remove using Graph API include: Email, OfficeCommunicationsOnline, Yammer. Not nullable
    supportedServices []string;
    // DNS records that the customer adds to the DNS zone file of the domain before the customer can complete domain ownership verification with Azure AD. Read-only, Nullable
    verificationDnsRecords []DomainDnsRecord;
}
// Instantiates a new domain and sets the default values.
func NewDomain()(*Domain) {
    m := &Domain{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the authenticationType property value. Indicates the configured authentication type for the domain. The value is either Managed or Federated. Managed indicates a cloud managed domain where Azure AD performs user authentication. Federated indicates authentication is federated with an identity provider such as the tenant's on-premises Active Directory via Active Directory Federation Services. This property is read-only and is not nullable.
func (m *Domain) GetAuthenticationType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationType
    }
}
// Gets the availabilityStatus property value. This property is always null except when the verify action is used. When the verify action is used, a domain entity is returned in the response. The availabilityStatus property of the domain entity in the response is either AvailableImmediately or EmailVerifiedDomainTakeoverScheduled.
func (m *Domain) GetAvailabilityStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.availabilityStatus
    }
}
// Gets the domainNameReferences property value. Read-only, Nullable
func (m *Domain) GetDomainNameReferences()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.domainNameReferences
    }
}
// Gets the isAdminManaged property value. The value of the property is false if the DNS record management of the domain has been delegated to Microsoft 365. Otherwise, the value is true. Not nullable
func (m *Domain) GetIsAdminManaged()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAdminManaged
    }
}
// Gets the isDefault property value. true if this is the default domain that is used for user creation. There is only one default domain per company. Not nullable
func (m *Domain) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// Gets the isInitial property value. true if this is the initial domain created by Microsoft Online Services (companyname.onmicrosoft.com). There is only one initial domain per company. Not nullable
func (m *Domain) GetIsInitial()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInitial
    }
}
// Gets the isRoot property value. true if the domain is a verified root domain. Otherwise, false if the domain is a subdomain or unverified. Not nullable
func (m *Domain) GetIsRoot()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRoot
    }
}
// Gets the isVerified property value. true if the domain has completed domain ownership verification. Not nullable
func (m *Domain) GetIsVerified()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isVerified
    }
}
// Gets the manufacturer property value. 
func (m *Domain) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// Gets the model property value. 
func (m *Domain) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// Gets the passwordNotificationWindowInDays property value. Specifies the number of days before a user receives notification that their password will expire. If the property is not set, a default value of 14 days will be used.
func (m *Domain) GetPasswordNotificationWindowInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordNotificationWindowInDays
    }
}
// Gets the passwordValidityPeriodInDays property value. Specifies the length of time that a password is valid before it must be changed. If the property is not set, a default value of 90 days will be used.
func (m *Domain) GetPasswordValidityPeriodInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordValidityPeriodInDays
    }
}
// Gets the serviceConfigurationRecords property value. DNS records the customer adds to the DNS zone file of the domain before the domain can be used by Microsoft Online services. Read-only, Nullable
func (m *Domain) GetServiceConfigurationRecords()([]DomainDnsRecord) {
    if m == nil {
        return nil
    } else {
        return m.serviceConfigurationRecords
    }
}
// Gets the state property value. Status of asynchronous operations scheduled for the domain.
func (m *Domain) GetState()(*DomainState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Gets the supportedServices property value. The capabilities assigned to the domain. Can include 0, 1 or more of following values: Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune. The values which you can add/remove using Graph API include: Email, OfficeCommunicationsOnline, Yammer. Not nullable
func (m *Domain) GetSupportedServices()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedServices
    }
}
// Gets the verificationDnsRecords property value. DNS records that the customer adds to the DNS zone file of the domain before the customer can complete domain ownership verification with Azure AD. Read-only, Nullable
func (m *Domain) GetVerificationDnsRecords()([]DomainDnsRecord) {
    if m == nil {
        return nil
    } else {
        return m.verificationDnsRecords
    }
}
// The deserialization information for the current model
func (m *Domain) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAuthenticationType(val)
        return nil
    }
    res["availabilityStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAvailabilityStatus(val)
        return nil
    }
    res["domainNameReferences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetDomainNameReferences(res)
        return nil
    }
    res["isAdminManaged"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAdminManaged(val)
        return nil
    }
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefault(val)
        return nil
    }
    res["isInitial"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsInitial(val)
        return nil
    }
    res["isRoot"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRoot(val)
        return nil
    }
    res["isVerified"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsVerified(val)
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManufacturer(val)
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetModel(val)
        return nil
    }
    res["passwordNotificationWindowInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPasswordNotificationWindowInDays(val)
        return nil
    }
    res["passwordValidityPeriodInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPasswordValidityPeriodInDays(val)
        return nil
    }
    res["serviceConfigurationRecords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDomainDnsRecord() })
        if err != nil {
            return err
        }
        res := make([]DomainDnsRecord, len(val))
        for i, v := range val {
            res[i] = *(v.(*DomainDnsRecord))
        }
        m.SetServiceConfigurationRecords(res)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDomainState() })
        if err != nil {
            return err
        }
        m.SetState(val.(*DomainState))
        return nil
    }
    res["supportedServices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetSupportedServices(res)
        return nil
    }
    res["verificationDnsRecords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDomainDnsRecord() })
        if err != nil {
            return err
        }
        res := make([]DomainDnsRecord, len(val))
        for i, v := range val {
            res[i] = *(v.(*DomainDnsRecord))
        }
        m.SetVerificationDnsRecords(res)
        return nil
    }
    return res
}
func (m *Domain) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Domain) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("authenticationType", m.GetAuthenticationType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("availabilityStatus", m.GetAvailabilityStatus())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDomainNameReferences()))
        for i, v := range m.GetDomainNameReferences() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("domainNameReferences", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAdminManaged", m.GetIsAdminManaged())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isInitial", m.GetIsInitial())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRoot", m.GetIsRoot())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isVerified", m.GetIsVerified())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordNotificationWindowInDays", m.GetPasswordNotificationWindowInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordValidityPeriodInDays", m.GetPasswordValidityPeriodInDays())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetServiceConfigurationRecords()))
        for i, v := range m.GetServiceConfigurationRecords() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("serviceConfigurationRecords", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("supportedServices", m.GetSupportedServices())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetVerificationDnsRecords()))
        for i, v := range m.GetVerificationDnsRecords() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("verificationDnsRecords", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the authenticationType property value. Indicates the configured authentication type for the domain. The value is either Managed or Federated. Managed indicates a cloud managed domain where Azure AD performs user authentication. Federated indicates authentication is federated with an identity provider such as the tenant's on-premises Active Directory via Active Directory Federation Services. This property is read-only and is not nullable.
// Parameters:
//  - value : Value to set for the authenticationType property.
func (m *Domain) SetAuthenticationType(value *string)() {
    m.authenticationType = value
}
// Sets the availabilityStatus property value. This property is always null except when the verify action is used. When the verify action is used, a domain entity is returned in the response. The availabilityStatus property of the domain entity in the response is either AvailableImmediately or EmailVerifiedDomainTakeoverScheduled.
// Parameters:
//  - value : Value to set for the availabilityStatus property.
func (m *Domain) SetAvailabilityStatus(value *string)() {
    m.availabilityStatus = value
}
// Sets the domainNameReferences property value. Read-only, Nullable
// Parameters:
//  - value : Value to set for the domainNameReferences property.
func (m *Domain) SetDomainNameReferences(value []DirectoryObject)() {
    m.domainNameReferences = value
}
// Sets the isAdminManaged property value. The value of the property is false if the DNS record management of the domain has been delegated to Microsoft 365. Otherwise, the value is true. Not nullable
// Parameters:
//  - value : Value to set for the isAdminManaged property.
func (m *Domain) SetIsAdminManaged(value *bool)() {
    m.isAdminManaged = value
}
// Sets the isDefault property value. true if this is the default domain that is used for user creation. There is only one default domain per company. Not nullable
// Parameters:
//  - value : Value to set for the isDefault property.
func (m *Domain) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// Sets the isInitial property value. true if this is the initial domain created by Microsoft Online Services (companyname.onmicrosoft.com). There is only one initial domain per company. Not nullable
// Parameters:
//  - value : Value to set for the isInitial property.
func (m *Domain) SetIsInitial(value *bool)() {
    m.isInitial = value
}
// Sets the isRoot property value. true if the domain is a verified root domain. Otherwise, false if the domain is a subdomain or unverified. Not nullable
// Parameters:
//  - value : Value to set for the isRoot property.
func (m *Domain) SetIsRoot(value *bool)() {
    m.isRoot = value
}
// Sets the isVerified property value. true if the domain has completed domain ownership verification. Not nullable
// Parameters:
//  - value : Value to set for the isVerified property.
func (m *Domain) SetIsVerified(value *bool)() {
    m.isVerified = value
}
// Sets the manufacturer property value. 
// Parameters:
//  - value : Value to set for the manufacturer property.
func (m *Domain) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// Sets the model property value. 
// Parameters:
//  - value : Value to set for the model property.
func (m *Domain) SetModel(value *string)() {
    m.model = value
}
// Sets the passwordNotificationWindowInDays property value. Specifies the number of days before a user receives notification that their password will expire. If the property is not set, a default value of 14 days will be used.
// Parameters:
//  - value : Value to set for the passwordNotificationWindowInDays property.
func (m *Domain) SetPasswordNotificationWindowInDays(value *int32)() {
    m.passwordNotificationWindowInDays = value
}
// Sets the passwordValidityPeriodInDays property value. Specifies the length of time that a password is valid before it must be changed. If the property is not set, a default value of 90 days will be used.
// Parameters:
//  - value : Value to set for the passwordValidityPeriodInDays property.
func (m *Domain) SetPasswordValidityPeriodInDays(value *int32)() {
    m.passwordValidityPeriodInDays = value
}
// Sets the serviceConfigurationRecords property value. DNS records the customer adds to the DNS zone file of the domain before the domain can be used by Microsoft Online services. Read-only, Nullable
// Parameters:
//  - value : Value to set for the serviceConfigurationRecords property.
func (m *Domain) SetServiceConfigurationRecords(value []DomainDnsRecord)() {
    m.serviceConfigurationRecords = value
}
// Sets the state property value. Status of asynchronous operations scheduled for the domain.
// Parameters:
//  - value : Value to set for the state property.
func (m *Domain) SetState(value *DomainState)() {
    m.state = value
}
// Sets the supportedServices property value. The capabilities assigned to the domain. Can include 0, 1 or more of following values: Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune. The values which you can add/remove using Graph API include: Email, OfficeCommunicationsOnline, Yammer. Not nullable
// Parameters:
//  - value : Value to set for the supportedServices property.
func (m *Domain) SetSupportedServices(value []string)() {
    m.supportedServices = value
}
// Sets the verificationDnsRecords property value. DNS records that the customer adds to the DNS zone file of the domain before the customer can complete domain ownership verification with Azure AD. Read-only, Nullable
// Parameters:
//  - value : Value to set for the verificationDnsRecords property.
func (m *Domain) SetVerificationDnsRecords(value []DomainDnsRecord)() {
    m.verificationDnsRecords = value
}
