package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Organization 
type Organization struct {
    DirectoryObject
    // The collection of service plans associated with the tenant. Not nullable.
    assignedPlans []AssignedPlan;
    // 
    branding *OrganizationalBranding;
    // Telephone number for the organization. Although this is a string collection, only one number can be set for this property.
    businessPhones []string;
    // Navigation property to manage certificate-based authentication configuration. Only a single instance of certificateBasedAuthConfiguration can be created in the collection.
    certificateBasedAuthConfiguration []CertificateBasedAuthConfiguration;
    // City name of the address for the organization.
    city *string;
    // Country/region name of the address for the organization.
    country *string;
    // Country or region abbreviation for the organization in ISO 3166-2 format.
    countryLetterCode *string;
    // Timestamp of when the organization was created. The value cannot be modified and is automatically populated when the organization is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The display name for the tenant.
    displayName *string;
    // The collection of open extensions defined for the organization resource. Nullable.
    extensions []Extension;
    // Not nullable.
    marketingNotificationEmails []string;
    // Mobile device management authority. Possible values are: unknown, intune, sccm, office365.
    mobileDeviceManagementAuthority *MdmAuthority;
    // The time and date at which the tenant was last synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; Nullable. null if this object has never been synced from an on-premises directory (default).
    onPremisesSyncEnabled *bool;
    // Postal code of the address for the organization.
    postalCode *string;
    // The preferred language for the organization. Should follow ISO 639-1 Code; for example en.
    preferredLanguage *string;
    // The privacy profile of an organization.
    privacyProfile *PrivacyProfile;
    // Not nullable.
    provisionedPlans []ProvisionedPlan;
    // 
    securityComplianceNotificationMails []string;
    // 
    securityComplianceNotificationPhones []string;
    // State name of the address for the organization.
    state *string;
    // Street name of the address for organization.
    street *string;
    // Not nullable.
    technicalNotificationMails []string;
    // 
    tenantType *string;
    // The collection of domains associated with this tenant. Not nullable.
    verifiedDomains []VerifiedDomain;
}
// NewOrganization instantiates a new organization and sets the default values.
func NewOrganization()(*Organization) {
    m := &Organization{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// GetAssignedPlans gets the assignedPlans property value. The collection of service plans associated with the tenant. Not nullable.
func (m *Organization) GetAssignedPlans()([]AssignedPlan) {
    if m == nil {
        return nil
    } else {
        return m.assignedPlans
    }
}
// GetBranding gets the branding property value. 
func (m *Organization) GetBranding()(*OrganizationalBranding) {
    if m == nil {
        return nil
    } else {
        return m.branding
    }
}
// GetBusinessPhones gets the businessPhones property value. Telephone number for the organization. Although this is a string collection, only one number can be set for this property.
func (m *Organization) GetBusinessPhones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.businessPhones
    }
}
// GetCertificateBasedAuthConfiguration gets the certificateBasedAuthConfiguration property value. Navigation property to manage certificate-based authentication configuration. Only a single instance of certificateBasedAuthConfiguration can be created in the collection.
func (m *Organization) GetCertificateBasedAuthConfiguration()([]CertificateBasedAuthConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.certificateBasedAuthConfiguration
    }
}
// GetCity gets the city property value. City name of the address for the organization.
func (m *Organization) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
// GetCountry gets the country property value. Country/region name of the address for the organization.
func (m *Organization) GetCountry()(*string) {
    if m == nil {
        return nil
    } else {
        return m.country
    }
}
// GetCountryLetterCode gets the countryLetterCode property value. Country or region abbreviation for the organization in ISO 3166-2 format.
func (m *Organization) GetCountryLetterCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryLetterCode
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Timestamp of when the organization was created. The value cannot be modified and is automatically populated when the organization is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Organization) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDisplayName gets the displayName property value. The display name for the tenant.
func (m *Organization) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the organization resource. Nullable.
func (m *Organization) GetExtensions()([]Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// GetMarketingNotificationEmails gets the marketingNotificationEmails property value. Not nullable.
func (m *Organization) GetMarketingNotificationEmails()([]string) {
    if m == nil {
        return nil
    } else {
        return m.marketingNotificationEmails
    }
}
// GetMobileDeviceManagementAuthority gets the mobileDeviceManagementAuthority property value. Mobile device management authority. Possible values are: unknown, intune, sccm, office365.
func (m *Organization) GetMobileDeviceManagementAuthority()(*MdmAuthority) {
    if m == nil {
        return nil
    } else {
        return m.mobileDeviceManagementAuthority
    }
}
// GetOnPremisesLastSyncDateTime gets the onPremisesLastSyncDateTime property value. The time and date at which the tenant was last synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Organization) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesLastSyncDateTime
    }
}
// GetOnPremisesSyncEnabled gets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; Nullable. null if this object has never been synced from an on-premises directory (default).
func (m *Organization) GetOnPremisesSyncEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSyncEnabled
    }
}
// GetPostalCode gets the postalCode property value. Postal code of the address for the organization.
func (m *Organization) GetPostalCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postalCode
    }
}
// GetPreferredLanguage gets the preferredLanguage property value. The preferred language for the organization. Should follow ISO 639-1 Code; for example en.
func (m *Organization) GetPreferredLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredLanguage
    }
}
// GetPrivacyProfile gets the privacyProfile property value. The privacy profile of an organization.
func (m *Organization) GetPrivacyProfile()(*PrivacyProfile) {
    if m == nil {
        return nil
    } else {
        return m.privacyProfile
    }
}
// GetProvisionedPlans gets the provisionedPlans property value. Not nullable.
func (m *Organization) GetProvisionedPlans()([]ProvisionedPlan) {
    if m == nil {
        return nil
    } else {
        return m.provisionedPlans
    }
}
// GetSecurityComplianceNotificationMails gets the securityComplianceNotificationMails property value. 
func (m *Organization) GetSecurityComplianceNotificationMails()([]string) {
    if m == nil {
        return nil
    } else {
        return m.securityComplianceNotificationMails
    }
}
// GetSecurityComplianceNotificationPhones gets the securityComplianceNotificationPhones property value. 
func (m *Organization) GetSecurityComplianceNotificationPhones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.securityComplianceNotificationPhones
    }
}
// GetState gets the state property value. State name of the address for the organization.
func (m *Organization) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetStreet gets the street property value. Street name of the address for organization.
func (m *Organization) GetStreet()(*string) {
    if m == nil {
        return nil
    } else {
        return m.street
    }
}
// GetTechnicalNotificationMails gets the technicalNotificationMails property value. Not nullable.
func (m *Organization) GetTechnicalNotificationMails()([]string) {
    if m == nil {
        return nil
    } else {
        return m.technicalNotificationMails
    }
}
// GetTenantType gets the tenantType property value. 
func (m *Organization) GetTenantType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantType
    }
}
// GetVerifiedDomains gets the verifiedDomains property value. The collection of domains associated with this tenant. Not nullable.
func (m *Organization) GetVerifiedDomains()([]VerifiedDomain) {
    if m == nil {
        return nil
    } else {
        return m.verifiedDomains
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Organization) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["assignedPlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignedPlan() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedPlan, len(val))
            for i, v := range val {
                res[i] = *(v.(*AssignedPlan))
            }
            m.SetAssignedPlans(res)
        }
        return nil
    }
    res["branding"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOrganizationalBranding() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBranding(val.(*OrganizationalBranding))
        }
        return nil
    }
    res["businessPhones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetBusinessPhones(res)
        }
        return nil
    }
    res["certificateBasedAuthConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCertificateBasedAuthConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CertificateBasedAuthConfiguration, len(val))
            for i, v := range val {
                res[i] = *(v.(*CertificateBasedAuthConfiguration))
            }
            m.SetCertificateBasedAuthConfiguration(res)
        }
        return nil
    }
    res["city"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["country"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountry(val)
        }
        return nil
    }
    res["countryLetterCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryLetterCode(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtension() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extension, len(val))
            for i, v := range val {
                res[i] = *(v.(*Extension))
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["marketingNotificationEmails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMarketingNotificationEmails(res)
        }
        return nil
    }
    res["mobileDeviceManagementAuthority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMdmAuthority)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(MdmAuthority)
            m.SetMobileDeviceManagementAuthority(&cast)
        }
        return nil
    }
    res["onPremisesLastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesLastSyncDateTime(val)
        }
        return nil
    }
    res["onPremisesSyncEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSyncEnabled(val)
        }
        return nil
    }
    res["postalCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostalCode(val)
        }
        return nil
    }
    res["preferredLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredLanguage(val)
        }
        return nil
    }
    res["privacyProfile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivacyProfile() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyProfile(val.(*PrivacyProfile))
        }
        return nil
    }
    res["provisionedPlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisionedPlan() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProvisionedPlan, len(val))
            for i, v := range val {
                res[i] = *(v.(*ProvisionedPlan))
            }
            m.SetProvisionedPlans(res)
        }
        return nil
    }
    res["securityComplianceNotificationMails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSecurityComplianceNotificationMails(res)
        }
        return nil
    }
    res["securityComplianceNotificationPhones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSecurityComplianceNotificationPhones(res)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["street"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStreet(val)
        }
        return nil
    }
    res["technicalNotificationMails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTechnicalNotificationMails(res)
        }
        return nil
    }
    res["tenantType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantType(val)
        }
        return nil
    }
    res["verifiedDomains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVerifiedDomain() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VerifiedDomain, len(val))
            for i, v := range val {
                res[i] = *(v.(*VerifiedDomain))
            }
            m.SetVerifiedDomains(res)
        }
        return nil
    }
    return res
}
func (m *Organization) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Organization) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignedPlans() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignedPlans()))
        for i, v := range m.GetAssignedPlans() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignedPlans", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("branding", m.GetBranding())
        if err != nil {
            return err
        }
    }
    if m.GetBusinessPhones() != nil {
        err = writer.WriteCollectionOfStringValues("businessPhones", m.GetBusinessPhones())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateBasedAuthConfiguration() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCertificateBasedAuthConfiguration()))
        for i, v := range m.GetCertificateBasedAuthConfiguration() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("certificateBasedAuthConfiguration", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("country", m.GetCountry())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("countryLetterCode", m.GetCountryLetterCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMarketingNotificationEmails() != nil {
        err = writer.WriteCollectionOfStringValues("marketingNotificationEmails", m.GetMarketingNotificationEmails())
        if err != nil {
            return err
        }
    }
    if m.GetMobileDeviceManagementAuthority() != nil {
        cast := m.GetMobileDeviceManagementAuthority().String()
        err = writer.WriteStringValue("mobileDeviceManagementAuthority", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("onPremisesLastSyncDateTime", m.GetOnPremisesLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onPremisesSyncEnabled", m.GetOnPremisesSyncEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("postalCode", m.GetPostalCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredLanguage", m.GetPreferredLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("privacyProfile", m.GetPrivacyProfile())
        if err != nil {
            return err
        }
    }
    if m.GetProvisionedPlans() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProvisionedPlans()))
        for i, v := range m.GetProvisionedPlans() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("provisionedPlans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecurityComplianceNotificationMails() != nil {
        err = writer.WriteCollectionOfStringValues("securityComplianceNotificationMails", m.GetSecurityComplianceNotificationMails())
        if err != nil {
            return err
        }
    }
    if m.GetSecurityComplianceNotificationPhones() != nil {
        err = writer.WriteCollectionOfStringValues("securityComplianceNotificationPhones", m.GetSecurityComplianceNotificationPhones())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("street", m.GetStreet())
        if err != nil {
            return err
        }
    }
    if m.GetTechnicalNotificationMails() != nil {
        err = writer.WriteCollectionOfStringValues("technicalNotificationMails", m.GetTechnicalNotificationMails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantType", m.GetTenantType())
        if err != nil {
            return err
        }
    }
    if m.GetVerifiedDomains() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetVerifiedDomains()))
        for i, v := range m.GetVerifiedDomains() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("verifiedDomains", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedPlans sets the assignedPlans property value. The collection of service plans associated with the tenant. Not nullable.
func (m *Organization) SetAssignedPlans(value []AssignedPlan)() {
    if m != nil {
        m.assignedPlans = value
    }
}
// SetBranding sets the branding property value. 
func (m *Organization) SetBranding(value *OrganizationalBranding)() {
    if m != nil {
        m.branding = value
    }
}
// SetBusinessPhones sets the businessPhones property value. Telephone number for the organization. Although this is a string collection, only one number can be set for this property.
func (m *Organization) SetBusinessPhones(value []string)() {
    if m != nil {
        m.businessPhones = value
    }
}
// SetCertificateBasedAuthConfiguration sets the certificateBasedAuthConfiguration property value. Navigation property to manage certificate-based authentication configuration. Only a single instance of certificateBasedAuthConfiguration can be created in the collection.
func (m *Organization) SetCertificateBasedAuthConfiguration(value []CertificateBasedAuthConfiguration)() {
    if m != nil {
        m.certificateBasedAuthConfiguration = value
    }
}
// SetCity sets the city property value. City name of the address for the organization.
func (m *Organization) SetCity(value *string)() {
    if m != nil {
        m.city = value
    }
}
// SetCountry sets the country property value. Country/region name of the address for the organization.
func (m *Organization) SetCountry(value *string)() {
    if m != nil {
        m.country = value
    }
}
// SetCountryLetterCode sets the countryLetterCode property value. Country or region abbreviation for the organization in ISO 3166-2 format.
func (m *Organization) SetCountryLetterCode(value *string)() {
    if m != nil {
        m.countryLetterCode = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Timestamp of when the organization was created. The value cannot be modified and is automatically populated when the organization is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Organization) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the tenant.
func (m *Organization) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the organization resource. Nullable.
func (m *Organization) SetExtensions(value []Extension)() {
    if m != nil {
        m.extensions = value
    }
}
// SetMarketingNotificationEmails sets the marketingNotificationEmails property value. Not nullable.
func (m *Organization) SetMarketingNotificationEmails(value []string)() {
    if m != nil {
        m.marketingNotificationEmails = value
    }
}
// SetMobileDeviceManagementAuthority sets the mobileDeviceManagementAuthority property value. Mobile device management authority. Possible values are: unknown, intune, sccm, office365.
func (m *Organization) SetMobileDeviceManagementAuthority(value *MdmAuthority)() {
    if m != nil {
        m.mobileDeviceManagementAuthority = value
    }
}
// SetOnPremisesLastSyncDateTime sets the onPremisesLastSyncDateTime property value. The time and date at which the tenant was last synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Organization) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.onPremisesLastSyncDateTime = value
    }
}
// SetOnPremisesSyncEnabled sets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; Nullable. null if this object has never been synced from an on-premises directory (default).
func (m *Organization) SetOnPremisesSyncEnabled(value *bool)() {
    if m != nil {
        m.onPremisesSyncEnabled = value
    }
}
// SetPostalCode sets the postalCode property value. Postal code of the address for the organization.
func (m *Organization) SetPostalCode(value *string)() {
    if m != nil {
        m.postalCode = value
    }
}
// SetPreferredLanguage sets the preferredLanguage property value. The preferred language for the organization. Should follow ISO 639-1 Code; for example en.
func (m *Organization) SetPreferredLanguage(value *string)() {
    if m != nil {
        m.preferredLanguage = value
    }
}
// SetPrivacyProfile sets the privacyProfile property value. The privacy profile of an organization.
func (m *Organization) SetPrivacyProfile(value *PrivacyProfile)() {
    if m != nil {
        m.privacyProfile = value
    }
}
// SetProvisionedPlans sets the provisionedPlans property value. Not nullable.
func (m *Organization) SetProvisionedPlans(value []ProvisionedPlan)() {
    if m != nil {
        m.provisionedPlans = value
    }
}
// SetSecurityComplianceNotificationMails sets the securityComplianceNotificationMails property value. 
func (m *Organization) SetSecurityComplianceNotificationMails(value []string)() {
    if m != nil {
        m.securityComplianceNotificationMails = value
    }
}
// SetSecurityComplianceNotificationPhones sets the securityComplianceNotificationPhones property value. 
func (m *Organization) SetSecurityComplianceNotificationPhones(value []string)() {
    if m != nil {
        m.securityComplianceNotificationPhones = value
    }
}
// SetState sets the state property value. State name of the address for the organization.
func (m *Organization) SetState(value *string)() {
    if m != nil {
        m.state = value
    }
}
// SetStreet sets the street property value. Street name of the address for organization.
func (m *Organization) SetStreet(value *string)() {
    if m != nil {
        m.street = value
    }
}
// SetTechnicalNotificationMails sets the technicalNotificationMails property value. Not nullable.
func (m *Organization) SetTechnicalNotificationMails(value []string)() {
    if m != nil {
        m.technicalNotificationMails = value
    }
}
// SetTenantType sets the tenantType property value. 
func (m *Organization) SetTenantType(value *string)() {
    if m != nil {
        m.tenantType = value
    }
}
// SetVerifiedDomains sets the verifiedDomains property value. The collection of domains associated with this tenant. Not nullable.
func (m *Organization) SetVerifiedDomains(value []VerifiedDomain)() {
    if m != nil {
        m.verifiedDomains = value
    }
}
