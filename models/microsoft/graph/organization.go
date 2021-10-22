package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Organization struct {
    DirectoryObject
    assignedPlans []AssignedPlan;
    branding *OrganizationalBranding;
    businessPhones []string;
    certificateBasedAuthConfiguration []CertificateBasedAuthConfiguration;
    city *string;
    country *string;
    countryLetterCode *string;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    displayName *string;
    extensions []Extension;
    marketingNotificationEmails []string;
    mobileDeviceManagementAuthority *MdmAuthority;
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    onPremisesSyncEnabled *bool;
    postalCode *string;
    preferredLanguage *string;
    privacyProfile *PrivacyProfile;
    provisionedPlans []ProvisionedPlan;
    securityComplianceNotificationMails []string;
    securityComplianceNotificationPhones []string;
    state *string;
    street *string;
    technicalNotificationMails []string;
    tenantType *string;
    verifiedDomains []VerifiedDomain;
}
func NewOrganization()(*Organization) {
    m := &Organization{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
func (m *Organization) GetAssignedPlans()([]AssignedPlan) {
    if m == nil {
        return nil
    } else {
        return m.assignedPlans
    }
}
func (m *Organization) GetBranding()(*OrganizationalBranding) {
    if m == nil {
        return nil
    } else {
        return m.branding
    }
}
func (m *Organization) GetBusinessPhones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.businessPhones
    }
}
func (m *Organization) GetCertificateBasedAuthConfiguration()([]CertificateBasedAuthConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.certificateBasedAuthConfiguration
    }
}
func (m *Organization) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
func (m *Organization) GetCountry()(*string) {
    if m == nil {
        return nil
    } else {
        return m.country
    }
}
func (m *Organization) GetCountryLetterCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryLetterCode
    }
}
func (m *Organization) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *Organization) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *Organization) GetExtensions()([]Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
func (m *Organization) GetMarketingNotificationEmails()([]string) {
    if m == nil {
        return nil
    } else {
        return m.marketingNotificationEmails
    }
}
func (m *Organization) GetMobileDeviceManagementAuthority()(*MdmAuthority) {
    if m == nil {
        return nil
    } else {
        return m.mobileDeviceManagementAuthority
    }
}
func (m *Organization) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesLastSyncDateTime
    }
}
func (m *Organization) GetOnPremisesSyncEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSyncEnabled
    }
}
func (m *Organization) GetPostalCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postalCode
    }
}
func (m *Organization) GetPreferredLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredLanguage
    }
}
func (m *Organization) GetPrivacyProfile()(*PrivacyProfile) {
    if m == nil {
        return nil
    } else {
        return m.privacyProfile
    }
}
func (m *Organization) GetProvisionedPlans()([]ProvisionedPlan) {
    if m == nil {
        return nil
    } else {
        return m.provisionedPlans
    }
}
func (m *Organization) GetSecurityComplianceNotificationMails()([]string) {
    if m == nil {
        return nil
    } else {
        return m.securityComplianceNotificationMails
    }
}
func (m *Organization) GetSecurityComplianceNotificationPhones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.securityComplianceNotificationPhones
    }
}
func (m *Organization) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *Organization) GetStreet()(*string) {
    if m == nil {
        return nil
    } else {
        return m.street
    }
}
func (m *Organization) GetTechnicalNotificationMails()([]string) {
    if m == nil {
        return nil
    } else {
        return m.technicalNotificationMails
    }
}
func (m *Organization) GetTenantType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantType
    }
}
func (m *Organization) GetVerifiedDomains()([]VerifiedDomain) {
    if m == nil {
        return nil
    } else {
        return m.verifiedDomains
    }
}
func (m *Organization) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["assignedPlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignedPlan() })
        if err != nil {
            return err
        }
        res := make([]AssignedPlan, len(val))
        for i, v := range val {
            res[i] = *(v.(*AssignedPlan))
        }
        m.SetAssignedPlans(res)
        return nil
    }
    res["branding"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOrganizationalBranding() })
        if err != nil {
            return err
        }
        m.SetBranding(val.(*OrganizationalBranding))
        return nil
    }
    res["businessPhones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetBusinessPhones(res)
        return nil
    }
    res["certificateBasedAuthConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCertificateBasedAuthConfiguration() })
        if err != nil {
            return err
        }
        res := make([]CertificateBasedAuthConfiguration, len(val))
        for i, v := range val {
            res[i] = *(v.(*CertificateBasedAuthConfiguration))
        }
        m.SetCertificateBasedAuthConfiguration(res)
        return nil
    }
    res["city"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCity(val)
        return nil
    }
    res["country"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCountry(val)
        return nil
    }
    res["countryLetterCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCountryLetterCode(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtension() })
        if err != nil {
            return err
        }
        res := make([]Extension, len(val))
        for i, v := range val {
            res[i] = *(v.(*Extension))
        }
        m.SetExtensions(res)
        return nil
    }
    res["marketingNotificationEmails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetMarketingNotificationEmails(res)
        return nil
    }
    res["mobileDeviceManagementAuthority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMdmAuthority)
        if err != nil {
            return err
        }
        cast := val.(MdmAuthority)
        m.SetMobileDeviceManagementAuthority(&cast)
        return nil
    }
    res["onPremisesLastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetOnPremisesLastSyncDateTime(val)
        return nil
    }
    res["onPremisesSyncEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOnPremisesSyncEnabled(val)
        return nil
    }
    res["postalCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPostalCode(val)
        return nil
    }
    res["preferredLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPreferredLanguage(val)
        return nil
    }
    res["privacyProfile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivacyProfile() })
        if err != nil {
            return err
        }
        m.SetPrivacyProfile(val.(*PrivacyProfile))
        return nil
    }
    res["provisionedPlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisionedPlan() })
        if err != nil {
            return err
        }
        res := make([]ProvisionedPlan, len(val))
        for i, v := range val {
            res[i] = *(v.(*ProvisionedPlan))
        }
        m.SetProvisionedPlans(res)
        return nil
    }
    res["securityComplianceNotificationMails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSecurityComplianceNotificationMails(res)
        return nil
    }
    res["securityComplianceNotificationPhones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSecurityComplianceNotificationPhones(res)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetState(val)
        return nil
    }
    res["street"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStreet(val)
        return nil
    }
    res["technicalNotificationMails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetTechnicalNotificationMails(res)
        return nil
    }
    res["tenantType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantType(val)
        return nil
    }
    res["verifiedDomains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVerifiedDomain() })
        if err != nil {
            return err
        }
        res := make([]VerifiedDomain, len(val))
        for i, v := range val {
            res[i] = *(v.(*VerifiedDomain))
        }
        m.SetVerifiedDomains(res)
        return nil
    }
    return res
}
func (m *Organization) IsNil()(bool) {
    return m == nil
}
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
func (m *Organization) SetAssignedPlans(value []AssignedPlan)() {
    m.assignedPlans = value
}
func (m *Organization) SetBranding(value *OrganizationalBranding)() {
    m.branding = value
}
func (m *Organization) SetBusinessPhones(value []string)() {
    m.businessPhones = value
}
func (m *Organization) SetCertificateBasedAuthConfiguration(value []CertificateBasedAuthConfiguration)() {
    m.certificateBasedAuthConfiguration = value
}
func (m *Organization) SetCity(value *string)() {
    m.city = value
}
func (m *Organization) SetCountry(value *string)() {
    m.country = value
}
func (m *Organization) SetCountryLetterCode(value *string)() {
    m.countryLetterCode = value
}
func (m *Organization) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *Organization) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *Organization) SetExtensions(value []Extension)() {
    m.extensions = value
}
func (m *Organization) SetMarketingNotificationEmails(value []string)() {
    m.marketingNotificationEmails = value
}
func (m *Organization) SetMobileDeviceManagementAuthority(value *MdmAuthority)() {
    m.mobileDeviceManagementAuthority = value
}
func (m *Organization) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onPremisesLastSyncDateTime = value
}
func (m *Organization) SetOnPremisesSyncEnabled(value *bool)() {
    m.onPremisesSyncEnabled = value
}
func (m *Organization) SetPostalCode(value *string)() {
    m.postalCode = value
}
func (m *Organization) SetPreferredLanguage(value *string)() {
    m.preferredLanguage = value
}
func (m *Organization) SetPrivacyProfile(value *PrivacyProfile)() {
    m.privacyProfile = value
}
func (m *Organization) SetProvisionedPlans(value []ProvisionedPlan)() {
    m.provisionedPlans = value
}
func (m *Organization) SetSecurityComplianceNotificationMails(value []string)() {
    m.securityComplianceNotificationMails = value
}
func (m *Organization) SetSecurityComplianceNotificationPhones(value []string)() {
    m.securityComplianceNotificationPhones = value
}
func (m *Organization) SetState(value *string)() {
    m.state = value
}
func (m *Organization) SetStreet(value *string)() {
    m.street = value
}
func (m *Organization) SetTechnicalNotificationMails(value []string)() {
    m.technicalNotificationMails = value
}
func (m *Organization) SetTenantType(value *string)() {
    m.tenantType = value
}
func (m *Organization) SetVerifiedDomains(value []VerifiedDomain)() {
    m.verifiedDomains = value
}
