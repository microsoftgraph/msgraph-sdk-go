package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
    // The collection of open extensions defined for the organization. Read-only. Nullable.
    extensions []Extension;
    // Not nullable.
    marketingNotificationEmails []string;
    // Mobile device management authority. Possible values are: unknown, intune, sccm, office365.
    mobileDeviceManagementAuthority *MdmAuthority;
    // The time and date at which the tenant was last synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced. Nullable. null if this object has never been synced from an on-premises directory (default).
    onPremisesSyncEnabled *bool;
    // Postal code of the address for the organization.
    postalCode *string;
    // The preferred language for the organization. Should follow ISO 639-1 Code; for example, en.
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
// Instantiates a new organization and sets the default values.
func NewOrganization()(*Organization) {
    m := &Organization{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// Gets the assignedPlans property value. The collection of service plans associated with the tenant. Not nullable.
func (m *Organization) GetAssignedPlans()([]AssignedPlan) {
    if m == nil {
        return nil
    } else {
        return m.assignedPlans
    }
}
// Gets the branding property value. 
func (m *Organization) GetBranding()(*OrganizationalBranding) {
    if m == nil {
        return nil
    } else {
        return m.branding
    }
}
// Gets the businessPhones property value. Telephone number for the organization. Although this is a string collection, only one number can be set for this property.
func (m *Organization) GetBusinessPhones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.businessPhones
    }
}
// Gets the certificateBasedAuthConfiguration property value. Navigation property to manage certificate-based authentication configuration. Only a single instance of certificateBasedAuthConfiguration can be created in the collection.
func (m *Organization) GetCertificateBasedAuthConfiguration()([]CertificateBasedAuthConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.certificateBasedAuthConfiguration
    }
}
// Gets the city property value. City name of the address for the organization.
func (m *Organization) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
// Gets the country property value. Country/region name of the address for the organization.
func (m *Organization) GetCountry()(*string) {
    if m == nil {
        return nil
    } else {
        return m.country
    }
}
// Gets the countryLetterCode property value. Country or region abbreviation for the organization in ISO 3166-2 format.
func (m *Organization) GetCountryLetterCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryLetterCode
    }
}
// Gets the createdDateTime property value. Timestamp of when the organization was created. The value cannot be modified and is automatically populated when the organization is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Organization) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the displayName property value. The display name for the tenant.
func (m *Organization) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the extensions property value. The collection of open extensions defined for the organization. Read-only. Nullable.
func (m *Organization) GetExtensions()([]Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// Gets the marketingNotificationEmails property value. Not nullable.
func (m *Organization) GetMarketingNotificationEmails()([]string) {
    if m == nil {
        return nil
    } else {
        return m.marketingNotificationEmails
    }
}
// Gets the mobileDeviceManagementAuthority property value. Mobile device management authority. Possible values are: unknown, intune, sccm, office365.
func (m *Organization) GetMobileDeviceManagementAuthority()(*MdmAuthority) {
    if m == nil {
        return nil
    } else {
        return m.mobileDeviceManagementAuthority
    }
}
// Gets the onPremisesLastSyncDateTime property value. The time and date at which the tenant was last synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Organization) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesLastSyncDateTime
    }
}
// Gets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced. Nullable. null if this object has never been synced from an on-premises directory (default).
func (m *Organization) GetOnPremisesSyncEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSyncEnabled
    }
}
// Gets the postalCode property value. Postal code of the address for the organization.
func (m *Organization) GetPostalCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postalCode
    }
}
// Gets the preferredLanguage property value. The preferred language for the organization. Should follow ISO 639-1 Code; for example, en.
func (m *Organization) GetPreferredLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredLanguage
    }
}
// Gets the privacyProfile property value. The privacy profile of an organization.
func (m *Organization) GetPrivacyProfile()(*PrivacyProfile) {
    if m == nil {
        return nil
    } else {
        return m.privacyProfile
    }
}
// Gets the provisionedPlans property value. Not nullable.
func (m *Organization) GetProvisionedPlans()([]ProvisionedPlan) {
    if m == nil {
        return nil
    } else {
        return m.provisionedPlans
    }
}
// Gets the securityComplianceNotificationMails property value. 
func (m *Organization) GetSecurityComplianceNotificationMails()([]string) {
    if m == nil {
        return nil
    } else {
        return m.securityComplianceNotificationMails
    }
}
// Gets the securityComplianceNotificationPhones property value. 
func (m *Organization) GetSecurityComplianceNotificationPhones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.securityComplianceNotificationPhones
    }
}
// Gets the state property value. State name of the address for the organization.
func (m *Organization) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Gets the street property value. Street name of the address for organization.
func (m *Organization) GetStreet()(*string) {
    if m == nil {
        return nil
    } else {
        return m.street
    }
}
// Gets the technicalNotificationMails property value. Not nullable.
func (m *Organization) GetTechnicalNotificationMails()([]string) {
    if m == nil {
        return nil
    } else {
        return m.technicalNotificationMails
    }
}
// Gets the tenantType property value. 
func (m *Organization) GetTenantType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantType
    }
}
// Gets the verifiedDomains property value. The collection of domains associated with this tenant. Not nullable.
func (m *Organization) GetVerifiedDomains()([]VerifiedDomain) {
    if m == nil {
        return nil
    } else {
        return m.verifiedDomains
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Organization) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
        err = writer.WriteCollectionOfStringValues("businessPhones", m.GetBusinessPhones())
        if err != nil {
            return err
        }
    }
    {
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
    {
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
    {
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
    {
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
    {
        err = writer.WriteCollectionOfStringValues("securityComplianceNotificationMails", m.GetSecurityComplianceNotificationMails())
        if err != nil {
            return err
        }
    }
    {
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
    {
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
    {
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
// Sets the assignedPlans property value. The collection of service plans associated with the tenant. Not nullable.
// Parameters:
//  - value : Value to set for the assignedPlans property.
func (m *Organization) SetAssignedPlans(value []AssignedPlan)() {
    m.assignedPlans = value
}
// Sets the branding property value. 
// Parameters:
//  - value : Value to set for the branding property.
func (m *Organization) SetBranding(value *OrganizationalBranding)() {
    m.branding = value
}
// Sets the businessPhones property value. Telephone number for the organization. Although this is a string collection, only one number can be set for this property.
// Parameters:
//  - value : Value to set for the businessPhones property.
func (m *Organization) SetBusinessPhones(value []string)() {
    m.businessPhones = value
}
// Sets the certificateBasedAuthConfiguration property value. Navigation property to manage certificate-based authentication configuration. Only a single instance of certificateBasedAuthConfiguration can be created in the collection.
// Parameters:
//  - value : Value to set for the certificateBasedAuthConfiguration property.
func (m *Organization) SetCertificateBasedAuthConfiguration(value []CertificateBasedAuthConfiguration)() {
    m.certificateBasedAuthConfiguration = value
}
// Sets the city property value. City name of the address for the organization.
// Parameters:
//  - value : Value to set for the city property.
func (m *Organization) SetCity(value *string)() {
    m.city = value
}
// Sets the country property value. Country/region name of the address for the organization.
// Parameters:
//  - value : Value to set for the country property.
func (m *Organization) SetCountry(value *string)() {
    m.country = value
}
// Sets the countryLetterCode property value. Country or region abbreviation for the organization in ISO 3166-2 format.
// Parameters:
//  - value : Value to set for the countryLetterCode property.
func (m *Organization) SetCountryLetterCode(value *string)() {
    m.countryLetterCode = value
}
// Sets the createdDateTime property value. Timestamp of when the organization was created. The value cannot be modified and is automatically populated when the organization is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *Organization) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the displayName property value. The display name for the tenant.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Organization) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the extensions property value. The collection of open extensions defined for the organization. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the extensions property.
func (m *Organization) SetExtensions(value []Extension)() {
    m.extensions = value
}
// Sets the marketingNotificationEmails property value. Not nullable.
// Parameters:
//  - value : Value to set for the marketingNotificationEmails property.
func (m *Organization) SetMarketingNotificationEmails(value []string)() {
    m.marketingNotificationEmails = value
}
// Sets the mobileDeviceManagementAuthority property value. Mobile device management authority. Possible values are: unknown, intune, sccm, office365.
// Parameters:
//  - value : Value to set for the mobileDeviceManagementAuthority property.
func (m *Organization) SetMobileDeviceManagementAuthority(value *MdmAuthority)() {
    m.mobileDeviceManagementAuthority = value
}
// Sets the onPremisesLastSyncDateTime property value. The time and date at which the tenant was last synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// Parameters:
//  - value : Value to set for the onPremisesLastSyncDateTime property.
func (m *Organization) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onPremisesLastSyncDateTime = value
}
// Sets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced. Nullable. null if this object has never been synced from an on-premises directory (default).
// Parameters:
//  - value : Value to set for the onPremisesSyncEnabled property.
func (m *Organization) SetOnPremisesSyncEnabled(value *bool)() {
    m.onPremisesSyncEnabled = value
}
// Sets the postalCode property value. Postal code of the address for the organization.
// Parameters:
//  - value : Value to set for the postalCode property.
func (m *Organization) SetPostalCode(value *string)() {
    m.postalCode = value
}
// Sets the preferredLanguage property value. The preferred language for the organization. Should follow ISO 639-1 Code; for example, en.
// Parameters:
//  - value : Value to set for the preferredLanguage property.
func (m *Organization) SetPreferredLanguage(value *string)() {
    m.preferredLanguage = value
}
// Sets the privacyProfile property value. The privacy profile of an organization.
// Parameters:
//  - value : Value to set for the privacyProfile property.
func (m *Organization) SetPrivacyProfile(value *PrivacyProfile)() {
    m.privacyProfile = value
}
// Sets the provisionedPlans property value. Not nullable.
// Parameters:
//  - value : Value to set for the provisionedPlans property.
func (m *Organization) SetProvisionedPlans(value []ProvisionedPlan)() {
    m.provisionedPlans = value
}
// Sets the securityComplianceNotificationMails property value. 
// Parameters:
//  - value : Value to set for the securityComplianceNotificationMails property.
func (m *Organization) SetSecurityComplianceNotificationMails(value []string)() {
    m.securityComplianceNotificationMails = value
}
// Sets the securityComplianceNotificationPhones property value. 
// Parameters:
//  - value : Value to set for the securityComplianceNotificationPhones property.
func (m *Organization) SetSecurityComplianceNotificationPhones(value []string)() {
    m.securityComplianceNotificationPhones = value
}
// Sets the state property value. State name of the address for the organization.
// Parameters:
//  - value : Value to set for the state property.
func (m *Organization) SetState(value *string)() {
    m.state = value
}
// Sets the street property value. Street name of the address for organization.
// Parameters:
//  - value : Value to set for the street property.
func (m *Organization) SetStreet(value *string)() {
    m.street = value
}
// Sets the technicalNotificationMails property value. Not nullable.
// Parameters:
//  - value : Value to set for the technicalNotificationMails property.
func (m *Organization) SetTechnicalNotificationMails(value []string)() {
    m.technicalNotificationMails = value
}
// Sets the tenantType property value. 
// Parameters:
//  - value : Value to set for the tenantType property.
func (m *Organization) SetTenantType(value *string)() {
    m.tenantType = value
}
// Sets the verifiedDomains property value. The collection of domains associated with this tenant. Not nullable.
// Parameters:
//  - value : Value to set for the verifiedDomains property.
func (m *Organization) SetVerifiedDomains(value []VerifiedDomain)() {
    m.verifiedDomains = value
}
