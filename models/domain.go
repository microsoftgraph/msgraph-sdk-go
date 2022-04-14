package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Domain 
type Domain struct {
    Entity
    // Indicates the configured authentication type for the domain. The value is either Managed or Federated. Managed indicates a cloud managed domain where Azure AD performs user authentication. Federated indicates authentication is federated with an identity provider such as the tenant's on-premises Active Directory via Active Directory Federation Services. This property is read-only and is not nullable.
    authenticationType *string
    // This property is always null except when the verify action is used. When the verify action is used, a domain entity is returned in the response. The availabilityStatus property of the domain entity in the response is either AvailableImmediately or EmailVerifiedDomainTakeoverScheduled.
    availabilityStatus *string
    // Read-only, Nullable
    domainNameReferences []DirectoryObjectable
    // The value of the property is false if the DNS record management of the domain has been delegated to Microsoft 365. Otherwise, the value is true. Not nullable
    isAdminManaged *bool
    // true if this is the default domain that is used for user creation. There is only one default domain per company. Not nullable
    isDefault *bool
    // true if this is the initial domain created by Microsoft Online Services (companyname.onmicrosoft.com). There is only one initial domain per company. Not nullable
    isInitial *bool
    // true if the domain is a verified root domain. Otherwise, false if the domain is a subdomain or unverified. Not nullable
    isRoot *bool
    // true if the domain has completed domain ownership verification. Not nullable
    isVerified *bool
    // The manufacturer property
    manufacturer *string
    // The model property
    model *string
    // Specifies the number of days before a user receives notification that their password will expire. If the property is not set, a default value of 14 days will be used.
    passwordNotificationWindowInDays *int32
    // Specifies the length of time that a password is valid before it must be changed. If the property is not set, a default value of 90 days will be used.
    passwordValidityPeriodInDays *int32
    // DNS records the customer adds to the DNS zone file of the domain before the domain can be used by Microsoft Online services. Read-only, Nullable
    serviceConfigurationRecords []DomainDnsRecordable
    // Status of asynchronous operations scheduled for the domain.
    state DomainStateable
    // The capabilities assigned to the domain. Can include 0, 1 or more of following values: Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune. The values which you can add/remove using Graph API include: Email, OfficeCommunicationsOnline, Yammer. Not nullable
    supportedServices []string
    // DNS records that the customer adds to the DNS zone file of the domain before the customer can complete domain ownership verification with Azure AD. Read-only, Nullable
    verificationDnsRecords []DomainDnsRecordable
}
// NewDomain instantiates a new domain and sets the default values.
func NewDomain()(*Domain) {
    m := &Domain{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDomainFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDomainFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDomain(), nil
}
// GetAuthenticationType gets the authenticationType property value. Indicates the configured authentication type for the domain. The value is either Managed or Federated. Managed indicates a cloud managed domain where Azure AD performs user authentication. Federated indicates authentication is federated with an identity provider such as the tenant's on-premises Active Directory via Active Directory Federation Services. This property is read-only and is not nullable.
func (m *Domain) GetAuthenticationType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationType
    }
}
// GetAvailabilityStatus gets the availabilityStatus property value. This property is always null except when the verify action is used. When the verify action is used, a domain entity is returned in the response. The availabilityStatus property of the domain entity in the response is either AvailableImmediately or EmailVerifiedDomainTakeoverScheduled.
func (m *Domain) GetAvailabilityStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.availabilityStatus
    }
}
// GetDomainNameReferences gets the domainNameReferences property value. Read-only, Nullable
func (m *Domain) GetDomainNameReferences()([]DirectoryObjectable) {
    if m == nil {
        return nil
    } else {
        return m.domainNameReferences
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Domain) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationType(val)
        }
        return nil
    }
    res["availabilityStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityStatus(val)
        }
        return nil
    }
    res["domainNameReferences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetDomainNameReferences(res)
        }
        return nil
    }
    res["isAdminManaged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAdminManaged(val)
        }
        return nil
    }
    res["isDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["isInitial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInitial(val)
        }
        return nil
    }
    res["isRoot"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRoot(val)
        }
        return nil
    }
    res["isVerified"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsVerified(val)
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["passwordNotificationWindowInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordNotificationWindowInDays(val)
        }
        return nil
    }
    res["passwordValidityPeriodInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordValidityPeriodInDays(val)
        }
        return nil
    }
    res["serviceConfigurationRecords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDomainDnsRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DomainDnsRecordable, len(val))
            for i, v := range val {
                res[i] = v.(DomainDnsRecordable)
            }
            m.SetServiceConfigurationRecords(res)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDomainStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(DomainStateable))
        }
        return nil
    }
    res["supportedServices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSupportedServices(res)
        }
        return nil
    }
    res["verificationDnsRecords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDomainDnsRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DomainDnsRecordable, len(val))
            for i, v := range val {
                res[i] = v.(DomainDnsRecordable)
            }
            m.SetVerificationDnsRecords(res)
        }
        return nil
    }
    return res
}
// GetIsAdminManaged gets the isAdminManaged property value. The value of the property is false if the DNS record management of the domain has been delegated to Microsoft 365. Otherwise, the value is true. Not nullable
func (m *Domain) GetIsAdminManaged()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAdminManaged
    }
}
// GetIsDefault gets the isDefault property value. true if this is the default domain that is used for user creation. There is only one default domain per company. Not nullable
func (m *Domain) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// GetIsInitial gets the isInitial property value. true if this is the initial domain created by Microsoft Online Services (companyname.onmicrosoft.com). There is only one initial domain per company. Not nullable
func (m *Domain) GetIsInitial()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInitial
    }
}
// GetIsRoot gets the isRoot property value. true if the domain is a verified root domain. Otherwise, false if the domain is a subdomain or unverified. Not nullable
func (m *Domain) GetIsRoot()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRoot
    }
}
// GetIsVerified gets the isVerified property value. true if the domain has completed domain ownership verification. Not nullable
func (m *Domain) GetIsVerified()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isVerified
    }
}
// GetManufacturer gets the manufacturer property value. The manufacturer property
func (m *Domain) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// GetModel gets the model property value. The model property
func (m *Domain) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// GetPasswordNotificationWindowInDays gets the passwordNotificationWindowInDays property value. Specifies the number of days before a user receives notification that their password will expire. If the property is not set, a default value of 14 days will be used.
func (m *Domain) GetPasswordNotificationWindowInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordNotificationWindowInDays
    }
}
// GetPasswordValidityPeriodInDays gets the passwordValidityPeriodInDays property value. Specifies the length of time that a password is valid before it must be changed. If the property is not set, a default value of 90 days will be used.
func (m *Domain) GetPasswordValidityPeriodInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordValidityPeriodInDays
    }
}
// GetServiceConfigurationRecords gets the serviceConfigurationRecords property value. DNS records the customer adds to the DNS zone file of the domain before the domain can be used by Microsoft Online services. Read-only, Nullable
func (m *Domain) GetServiceConfigurationRecords()([]DomainDnsRecordable) {
    if m == nil {
        return nil
    } else {
        return m.serviceConfigurationRecords
    }
}
// GetState gets the state property value. Status of asynchronous operations scheduled for the domain.
func (m *Domain) GetState()(DomainStateable) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetSupportedServices gets the supportedServices property value. The capabilities assigned to the domain. Can include 0, 1 or more of following values: Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune. The values which you can add/remove using Graph API include: Email, OfficeCommunicationsOnline, Yammer. Not nullable
func (m *Domain) GetSupportedServices()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedServices
    }
}
// GetVerificationDnsRecords gets the verificationDnsRecords property value. DNS records that the customer adds to the DNS zone file of the domain before the customer can complete domain ownership verification with Azure AD. Read-only, Nullable
func (m *Domain) GetVerificationDnsRecords()([]DomainDnsRecordable) {
    if m == nil {
        return nil
    } else {
        return m.verificationDnsRecords
    }
}
// Serialize serializes information the current object
func (m *Domain) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetDomainNameReferences() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDomainNameReferences()))
        for i, v := range m.GetDomainNameReferences() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetServiceConfigurationRecords() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServiceConfigurationRecords()))
        for i, v := range m.GetServiceConfigurationRecords() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetSupportedServices() != nil {
        err = writer.WriteCollectionOfStringValues("supportedServices", m.GetSupportedServices())
        if err != nil {
            return err
        }
    }
    if m.GetVerificationDnsRecords() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVerificationDnsRecords()))
        for i, v := range m.GetVerificationDnsRecords() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("verificationDnsRecords", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationType sets the authenticationType property value. Indicates the configured authentication type for the domain. The value is either Managed or Federated. Managed indicates a cloud managed domain where Azure AD performs user authentication. Federated indicates authentication is federated with an identity provider such as the tenant's on-premises Active Directory via Active Directory Federation Services. This property is read-only and is not nullable.
func (m *Domain) SetAuthenticationType(value *string)() {
    if m != nil {
        m.authenticationType = value
    }
}
// SetAvailabilityStatus sets the availabilityStatus property value. This property is always null except when the verify action is used. When the verify action is used, a domain entity is returned in the response. The availabilityStatus property of the domain entity in the response is either AvailableImmediately or EmailVerifiedDomainTakeoverScheduled.
func (m *Domain) SetAvailabilityStatus(value *string)() {
    if m != nil {
        m.availabilityStatus = value
    }
}
// SetDomainNameReferences sets the domainNameReferences property value. Read-only, Nullable
func (m *Domain) SetDomainNameReferences(value []DirectoryObjectable)() {
    if m != nil {
        m.domainNameReferences = value
    }
}
// SetIsAdminManaged sets the isAdminManaged property value. The value of the property is false if the DNS record management of the domain has been delegated to Microsoft 365. Otherwise, the value is true. Not nullable
func (m *Domain) SetIsAdminManaged(value *bool)() {
    if m != nil {
        m.isAdminManaged = value
    }
}
// SetIsDefault sets the isDefault property value. true if this is the default domain that is used for user creation. There is only one default domain per company. Not nullable
func (m *Domain) SetIsDefault(value *bool)() {
    if m != nil {
        m.isDefault = value
    }
}
// SetIsInitial sets the isInitial property value. true if this is the initial domain created by Microsoft Online Services (companyname.onmicrosoft.com). There is only one initial domain per company. Not nullable
func (m *Domain) SetIsInitial(value *bool)() {
    if m != nil {
        m.isInitial = value
    }
}
// SetIsRoot sets the isRoot property value. true if the domain is a verified root domain. Otherwise, false if the domain is a subdomain or unverified. Not nullable
func (m *Domain) SetIsRoot(value *bool)() {
    if m != nil {
        m.isRoot = value
    }
}
// SetIsVerified sets the isVerified property value. true if the domain has completed domain ownership verification. Not nullable
func (m *Domain) SetIsVerified(value *bool)() {
    if m != nil {
        m.isVerified = value
    }
}
// SetManufacturer sets the manufacturer property value. The manufacturer property
func (m *Domain) SetManufacturer(value *string)() {
    if m != nil {
        m.manufacturer = value
    }
}
// SetModel sets the model property value. The model property
func (m *Domain) SetModel(value *string)() {
    if m != nil {
        m.model = value
    }
}
// SetPasswordNotificationWindowInDays sets the passwordNotificationWindowInDays property value. Specifies the number of days before a user receives notification that their password will expire. If the property is not set, a default value of 14 days will be used.
func (m *Domain) SetPasswordNotificationWindowInDays(value *int32)() {
    if m != nil {
        m.passwordNotificationWindowInDays = value
    }
}
// SetPasswordValidityPeriodInDays sets the passwordValidityPeriodInDays property value. Specifies the length of time that a password is valid before it must be changed. If the property is not set, a default value of 90 days will be used.
func (m *Domain) SetPasswordValidityPeriodInDays(value *int32)() {
    if m != nil {
        m.passwordValidityPeriodInDays = value
    }
}
// SetServiceConfigurationRecords sets the serviceConfigurationRecords property value. DNS records the customer adds to the DNS zone file of the domain before the domain can be used by Microsoft Online services. Read-only, Nullable
func (m *Domain) SetServiceConfigurationRecords(value []DomainDnsRecordable)() {
    if m != nil {
        m.serviceConfigurationRecords = value
    }
}
// SetState sets the state property value. Status of asynchronous operations scheduled for the domain.
func (m *Domain) SetState(value DomainStateable)() {
    if m != nil {
        m.state = value
    }
}
// SetSupportedServices sets the supportedServices property value. The capabilities assigned to the domain. Can include 0, 1 or more of following values: Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune. The values which you can add/remove using Graph API include: Email, OfficeCommunicationsOnline, Yammer. Not nullable
func (m *Domain) SetSupportedServices(value []string)() {
    if m != nil {
        m.supportedServices = value
    }
}
// SetVerificationDnsRecords sets the verificationDnsRecords property value. DNS records that the customer adds to the DNS zone file of the domain before the customer can complete domain ownership verification with Azure AD. Read-only, Nullable
func (m *Domain) SetVerificationDnsRecords(value []DomainDnsRecordable)() {
    if m != nil {
        m.verificationDnsRecords = value
    }
}
