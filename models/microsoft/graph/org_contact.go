package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OrgContact struct {
    DirectoryObject
    // Postal addresses for this organizational contact. For now a contact can only have one physical address.
    addresses []PhysicalOfficeAddress;
    // Name of the company that this organizational contact belong to. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    companyName *string;
    // The name for the department in which the contact works. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    department *string;
    // The contact's direct reports. (The users and contacts that have their manager property set to this contact.) Read-only. Nullable. Supports $expand.
    directReports []DirectoryObject;
    // Display name for this organizational contact. Supports $filter (eq, ne, NOT, ge, le, in, startsWith), $search, and $orderBy.
    displayName *string;
    // First name for this organizational contact. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    givenName *string;
    // Job title for this organizational contact. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    jobTitle *string;
    // The SMTP address for the contact, for example, 'jeff@contoso.onmicrosoft.com'. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    mail *string;
    // Email alias (portion of email address pre-pending the @ symbol) for this organizational contact. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    mailNickname *string;
    // The user or contact that is this contact's manager. Read-only. Supports $expand.
    manager *DirectoryObject;
    // Groups that this contact is a member of. Read-only. Nullable. Supports $expand.
    memberOf []DirectoryObject;
    // Date and time when this organizational contact was last synchronized from on-premises AD. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $filter (eq, ne, NOT, ge, le, in).
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // List of any synchronization provisioning errors for this organizational contact. Supports $filter (eq, NOT).
    onPremisesProvisioningErrors []OnPremisesProvisioningError;
    // true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced and now mastered in Exchange; null if this object has never been synced from an on-premises directory (default).
    onPremisesSyncEnabled *bool;
    // List of phones for this organizational contact. Phone types can be mobile, business, and businessFax. Only one of each type can ever be present in the collection. Supports $filter (eq, ne, NOT, in).
    phones []Phone;
    // For example: 'SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com'. The any operator is required for filter expressions on multi-valued properties. Supports $filter (eq, NOT, ge, le, startsWith).
    proxyAddresses []string;
    // Last name for this organizational contact. Supports $filter (eq, ne, NOT, ge, le, in, startsWith)
    surname *string;
    // 
    transitiveMemberOf []DirectoryObject;
}
// Instantiates a new orgContact and sets the default values.
func NewOrgContact()(*OrgContact) {
    m := &OrgContact{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// Gets the addresses property value. Postal addresses for this organizational contact. For now a contact can only have one physical address.
func (m *OrgContact) GetAddresses()([]PhysicalOfficeAddress) {
    if m == nil {
        return nil
    } else {
        return m.addresses
    }
}
// Gets the companyName property value. Name of the company that this organizational contact belong to. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *OrgContact) GetCompanyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.companyName
    }
}
// Gets the department property value. The name for the department in which the contact works. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *OrgContact) GetDepartment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.department
    }
}
// Gets the directReports property value. The contact's direct reports. (The users and contacts that have their manager property set to this contact.) Read-only. Nullable. Supports $expand.
func (m *OrgContact) GetDirectReports()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.directReports
    }
}
// Gets the displayName property value. Display name for this organizational contact. Supports $filter (eq, ne, NOT, ge, le, in, startsWith), $search, and $orderBy.
func (m *OrgContact) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the givenName property value. First name for this organizational contact. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *OrgContact) GetGivenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.givenName
    }
}
// Gets the jobTitle property value. Job title for this organizational contact. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *OrgContact) GetJobTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jobTitle
    }
}
// Gets the mail property value. The SMTP address for the contact, for example, 'jeff@contoso.onmicrosoft.com'. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *OrgContact) GetMail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mail
    }
}
// Gets the mailNickname property value. Email alias (portion of email address pre-pending the @ symbol) for this organizational contact. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *OrgContact) GetMailNickname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mailNickname
    }
}
// Gets the manager property value. The user or contact that is this contact's manager. Read-only. Supports $expand.
func (m *OrgContact) GetManager()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.manager
    }
}
// Gets the memberOf property value. Groups that this contact is a member of. Read-only. Nullable. Supports $expand.
func (m *OrgContact) GetMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.memberOf
    }
}
// Gets the onPremisesLastSyncDateTime property value. Date and time when this organizational contact was last synchronized from on-premises AD. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $filter (eq, ne, NOT, ge, le, in).
func (m *OrgContact) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesLastSyncDateTime
    }
}
// Gets the onPremisesProvisioningErrors property value. List of any synchronization provisioning errors for this organizational contact. Supports $filter (eq, NOT).
func (m *OrgContact) GetOnPremisesProvisioningErrors()([]OnPremisesProvisioningError) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesProvisioningErrors
    }
}
// Gets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced and now mastered in Exchange; null if this object has never been synced from an on-premises directory (default).
func (m *OrgContact) GetOnPremisesSyncEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSyncEnabled
    }
}
// Gets the phones property value. List of phones for this organizational contact. Phone types can be mobile, business, and businessFax. Only one of each type can ever be present in the collection. Supports $filter (eq, ne, NOT, in).
func (m *OrgContact) GetPhones()([]Phone) {
    if m == nil {
        return nil
    } else {
        return m.phones
    }
}
// Gets the proxyAddresses property value. For example: 'SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com'. The any operator is required for filter expressions on multi-valued properties. Supports $filter (eq, NOT, ge, le, startsWith).
func (m *OrgContact) GetProxyAddresses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.proxyAddresses
    }
}
// Gets the surname property value. Last name for this organizational contact. Supports $filter (eq, ne, NOT, ge, le, in, startsWith)
func (m *OrgContact) GetSurname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.surname
    }
}
// Gets the transitiveMemberOf property value. 
func (m *OrgContact) GetTransitiveMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.transitiveMemberOf
    }
}
// The deserialization information for the current model
func (m *OrgContact) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["addresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalOfficeAddress() })
        if err != nil {
            return err
        }
        res := make([]PhysicalOfficeAddress, len(val))
        for i, v := range val {
            res[i] = *(v.(*PhysicalOfficeAddress))
        }
        m.SetAddresses(res)
        return nil
    }
    res["companyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCompanyName(val)
        return nil
    }
    res["department"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDepartment(val)
        return nil
    }
    res["directReports"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetDirectReports(res)
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
    res["givenName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGivenName(val)
        return nil
    }
    res["jobTitle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJobTitle(val)
        return nil
    }
    res["mail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMail(val)
        return nil
    }
    res["mailNickname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMailNickname(val)
        return nil
    }
    res["manager"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        m.SetManager(val.(*DirectoryObject))
        return nil
    }
    res["memberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetMemberOf(res)
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
    res["onPremisesProvisioningErrors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesProvisioningError() })
        if err != nil {
            return err
        }
        res := make([]OnPremisesProvisioningError, len(val))
        for i, v := range val {
            res[i] = *(v.(*OnPremisesProvisioningError))
        }
        m.SetOnPremisesProvisioningErrors(res)
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
    res["phones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhone() })
        if err != nil {
            return err
        }
        res := make([]Phone, len(val))
        for i, v := range val {
            res[i] = *(v.(*Phone))
        }
        m.SetPhones(res)
        return nil
    }
    res["proxyAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetProxyAddresses(res)
        return nil
    }
    res["surname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSurname(val)
        return nil
    }
    res["transitiveMemberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetTransitiveMemberOf(res)
        return nil
    }
    return res
}
func (m *OrgContact) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OrgContact) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAddresses()))
        for i, v := range m.GetAddresses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("addresses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("companyName", m.GetCompanyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("department", m.GetDepartment())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDirectReports()))
        for i, v := range m.GetDirectReports() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("directReports", cast)
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
        err = writer.WriteStringValue("givenName", m.GetGivenName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("jobTitle", m.GetJobTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mail", m.GetMail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mailNickname", m.GetMailNickname())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("manager", m.GetManager())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMemberOf()))
        for i, v := range m.GetMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("memberOf", cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOnPremisesProvisioningErrors()))
        for i, v := range m.GetOnPremisesProvisioningErrors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("onPremisesProvisioningErrors", cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPhones()))
        for i, v := range m.GetPhones() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("phones", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("proxyAddresses", m.GetProxyAddresses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("surname", m.GetSurname())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTransitiveMemberOf()))
        for i, v := range m.GetTransitiveMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("transitiveMemberOf", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the addresses property value. Postal addresses for this organizational contact. For now a contact can only have one physical address.
// Parameters:
//  - value : Value to set for the addresses property.
func (m *OrgContact) SetAddresses(value []PhysicalOfficeAddress)() {
    m.addresses = value
}
// Sets the companyName property value. Name of the company that this organizational contact belong to. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the companyName property.
func (m *OrgContact) SetCompanyName(value *string)() {
    m.companyName = value
}
// Sets the department property value. The name for the department in which the contact works. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the department property.
func (m *OrgContact) SetDepartment(value *string)() {
    m.department = value
}
// Sets the directReports property value. The contact's direct reports. (The users and contacts that have their manager property set to this contact.) Read-only. Nullable. Supports $expand.
// Parameters:
//  - value : Value to set for the directReports property.
func (m *OrgContact) SetDirectReports(value []DirectoryObject)() {
    m.directReports = value
}
// Sets the displayName property value. Display name for this organizational contact. Supports $filter (eq, ne, NOT, ge, le, in, startsWith), $search, and $orderBy.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *OrgContact) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the givenName property value. First name for this organizational contact. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the givenName property.
func (m *OrgContact) SetGivenName(value *string)() {
    m.givenName = value
}
// Sets the jobTitle property value. Job title for this organizational contact. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the jobTitle property.
func (m *OrgContact) SetJobTitle(value *string)() {
    m.jobTitle = value
}
// Sets the mail property value. The SMTP address for the contact, for example, 'jeff@contoso.onmicrosoft.com'. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the mail property.
func (m *OrgContact) SetMail(value *string)() {
    m.mail = value
}
// Sets the mailNickname property value. Email alias (portion of email address pre-pending the @ symbol) for this organizational contact. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the mailNickname property.
func (m *OrgContact) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// Sets the manager property value. The user or contact that is this contact's manager. Read-only. Supports $expand.
// Parameters:
//  - value : Value to set for the manager property.
func (m *OrgContact) SetManager(value *DirectoryObject)() {
    m.manager = value
}
// Sets the memberOf property value. Groups that this contact is a member of. Read-only. Nullable. Supports $expand.
// Parameters:
//  - value : Value to set for the memberOf property.
func (m *OrgContact) SetMemberOf(value []DirectoryObject)() {
    m.memberOf = value
}
// Sets the onPremisesLastSyncDateTime property value. Date and time when this organizational contact was last synchronized from on-premises AD. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $filter (eq, ne, NOT, ge, le, in).
// Parameters:
//  - value : Value to set for the onPremisesLastSyncDateTime property.
func (m *OrgContact) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onPremisesLastSyncDateTime = value
}
// Sets the onPremisesProvisioningErrors property value. List of any synchronization provisioning errors for this organizational contact. Supports $filter (eq, NOT).
// Parameters:
//  - value : Value to set for the onPremisesProvisioningErrors property.
func (m *OrgContact) SetOnPremisesProvisioningErrors(value []OnPremisesProvisioningError)() {
    m.onPremisesProvisioningErrors = value
}
// Sets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced and now mastered in Exchange; null if this object has never been synced from an on-premises directory (default).
// Parameters:
//  - value : Value to set for the onPremisesSyncEnabled property.
func (m *OrgContact) SetOnPremisesSyncEnabled(value *bool)() {
    m.onPremisesSyncEnabled = value
}
// Sets the phones property value. List of phones for this organizational contact. Phone types can be mobile, business, and businessFax. Only one of each type can ever be present in the collection. Supports $filter (eq, ne, NOT, in).
// Parameters:
//  - value : Value to set for the phones property.
func (m *OrgContact) SetPhones(value []Phone)() {
    m.phones = value
}
// Sets the proxyAddresses property value. For example: 'SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com'. The any operator is required for filter expressions on multi-valued properties. Supports $filter (eq, NOT, ge, le, startsWith).
// Parameters:
//  - value : Value to set for the proxyAddresses property.
func (m *OrgContact) SetProxyAddresses(value []string)() {
    m.proxyAddresses = value
}
// Sets the surname property value. Last name for this organizational contact. Supports $filter (eq, ne, NOT, ge, le, in, startsWith)
// Parameters:
//  - value : Value to set for the surname property.
func (m *OrgContact) SetSurname(value *string)() {
    m.surname = value
}
// Sets the transitiveMemberOf property value. 
// Parameters:
//  - value : Value to set for the transitiveMemberOf property.
func (m *OrgContact) SetTransitiveMemberOf(value []DirectoryObject)() {
    m.transitiveMemberOf = value
}
