package delta

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// 
type Delta struct {
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OutlookItem
    // The name of the contact's assistant.
    assistantName *string;
    // The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    birthday *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The contact's business address.
    businessAddress *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress;
    // The business home page of the contact.
    businessHomePage *string;
    // The contact's business phone numbers.
    businessPhones []string;
    // The names of the contact's children.
    children []string;
    // The name of the contact's company.
    companyName *string;
    // The contact's department.
    department *string;
    // The contact's display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation.
    displayName *string;
    // The contact's email addresses.
    emailAddresses []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EmailAddress;
    // The collection of open extensions defined for the contact. Read-only. Nullable.
    extensions []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Extension;
    // The name the contact is filed under.
    fileAs *string;
    // The contact's generation.
    generation *string;
    // The contact's given name.
    givenName *string;
    // The contact's home address.
    homeAddress *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress;
    // The contact's home phone numbers.
    homePhones []string;
    // 
    imAddresses []string;
    // 
    initials *string;
    // 
    jobTitle *string;
    // 
    manager *string;
    // 
    middleName *string;
    // 
    mobilePhone *string;
    // The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
    multiValueExtendedProperties []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MultiValueLegacyExtendedProperty;
    // 
    nickName *string;
    // 
    officeLocation *string;
    // 
    otherAddress *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress;
    // 
    parentFolderId *string;
    // 
    personalNotes *string;
    // Optional contact picture. You can get or set a photo for a contact.
    photo *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ProfilePhoto;
    // 
    profession *string;
    // The collection of single-value extended properties defined for the contact. Read-only. Nullable.
    singleValueExtendedProperties []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SingleValueLegacyExtendedProperty;
    // 
    spouseName *string;
    // 
    surname *string;
    // 
    title *string;
    // 
    yomiCompanyName *string;
    // 
    yomiGivenName *string;
    // 
    yomiSurname *string;
}
// Instantiates a new delta and sets the default values.
func NewDelta()(*Delta) {
    m := &Delta{
        OutlookItem: *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOutlookItem(),
    }
    return m
}
// Gets the assistantName property value. The name of the contact's assistant.
func (m *Delta) GetAssistantName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assistantName
    }
}
// Gets the birthday property value. The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Delta) GetBirthday()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.birthday
    }
}
// Gets the businessAddress property value. The contact's business address.
func (m *Delta) GetBusinessAddress()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.businessAddress
    }
}
// Gets the businessHomePage property value. The business home page of the contact.
func (m *Delta) GetBusinessHomePage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.businessHomePage
    }
}
// Gets the businessPhones property value. The contact's business phone numbers.
func (m *Delta) GetBusinessPhones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.businessPhones
    }
}
// Gets the children property value. The names of the contact's children.
func (m *Delta) GetChildren()([]string) {
    if m == nil {
        return nil
    } else {
        return m.children
    }
}
// Gets the companyName property value. The name of the contact's company.
func (m *Delta) GetCompanyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.companyName
    }
}
// Gets the department property value. The contact's department.
func (m *Delta) GetDepartment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.department
    }
}
// Gets the displayName property value. The contact's display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation.
func (m *Delta) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the emailAddresses property value. The contact's email addresses.
func (m *Delta) GetEmailAddresses()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EmailAddress) {
    if m == nil {
        return nil
    } else {
        return m.emailAddresses
    }
}
// Gets the extensions property value. The collection of open extensions defined for the contact. Read-only. Nullable.
func (m *Delta) GetExtensions()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// Gets the fileAs property value. The name the contact is filed under.
func (m *Delta) GetFileAs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileAs
    }
}
// Gets the generation property value. The contact's generation.
func (m *Delta) GetGeneration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.generation
    }
}
// Gets the givenName property value. The contact's given name.
func (m *Delta) GetGivenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.givenName
    }
}
// Gets the homeAddress property value. The contact's home address.
func (m *Delta) GetHomeAddress()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.homeAddress
    }
}
// Gets the homePhones property value. The contact's home phone numbers.
func (m *Delta) GetHomePhones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.homePhones
    }
}
// Gets the imAddresses property value. 
func (m *Delta) GetImAddresses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.imAddresses
    }
}
// Gets the initials property value. 
func (m *Delta) GetInitials()(*string) {
    if m == nil {
        return nil
    } else {
        return m.initials
    }
}
// Gets the jobTitle property value. 
func (m *Delta) GetJobTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jobTitle
    }
}
// Gets the manager property value. 
func (m *Delta) GetManager()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manager
    }
}
// Gets the middleName property value. 
func (m *Delta) GetMiddleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.middleName
    }
}
// Gets the mobilePhone property value. 
func (m *Delta) GetMobilePhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mobilePhone
    }
}
// Gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
func (m *Delta) GetMultiValueExtendedProperties()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MultiValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// Gets the nickName property value. 
func (m *Delta) GetNickName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nickName
    }
}
// Gets the officeLocation property value. 
func (m *Delta) GetOfficeLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.officeLocation
    }
}
// Gets the otherAddress property value. 
func (m *Delta) GetOtherAddress()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.otherAddress
    }
}
// Gets the parentFolderId property value. 
func (m *Delta) GetParentFolderId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentFolderId
    }
}
// Gets the personalNotes property value. 
func (m *Delta) GetPersonalNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.personalNotes
    }
}
// Gets the photo property value. Optional contact picture. You can get or set a photo for a contact.
func (m *Delta) GetPhoto()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ProfilePhoto) {
    if m == nil {
        return nil
    } else {
        return m.photo
    }
}
// Gets the profession property value. 
func (m *Delta) GetProfession()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profession
    }
}
// Gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the contact. Read-only. Nullable.
func (m *Delta) GetSingleValueExtendedProperties()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SingleValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
}
// Gets the spouseName property value. 
func (m *Delta) GetSpouseName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.spouseName
    }
}
// Gets the surname property value. 
func (m *Delta) GetSurname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.surname
    }
}
// Gets the title property value. 
func (m *Delta) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// Gets the yomiCompanyName property value. 
func (m *Delta) GetYomiCompanyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yomiCompanyName
    }
}
// Gets the yomiGivenName property value. 
func (m *Delta) GetYomiGivenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yomiGivenName
    }
}
// Gets the yomiSurname property value. 
func (m *Delta) GetYomiSurname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yomiSurname
    }
}
// The deserialization information for the current model
func (m *Delta) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["assistantName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAssistantName(val)
        return nil
    }
    res["birthday"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetBirthday(val)
        return nil
    }
    res["businessAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPhysicalAddress() })
        if err != nil {
            return err
        }
        m.SetBusinessAddress(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress))
        return nil
    }
    res["businessHomePage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBusinessHomePage(val)
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
    res["children"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetChildren(res)
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
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["emailAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEmailAddress() })
        if err != nil {
            return err
        }
        res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EmailAddress, len(val))
        for i, v := range val {
            res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EmailAddress))
        }
        m.SetEmailAddresses(res)
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewExtension() })
        if err != nil {
            return err
        }
        res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Extension, len(val))
        for i, v := range val {
            res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Extension))
        }
        m.SetExtensions(res)
        return nil
    }
    res["fileAs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFileAs(val)
        return nil
    }
    res["generation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGeneration(val)
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
    res["homeAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPhysicalAddress() })
        if err != nil {
            return err
        }
        m.SetHomeAddress(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress))
        return nil
    }
    res["homePhones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetHomePhones(res)
        return nil
    }
    res["imAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetImAddresses(res)
        return nil
    }
    res["initials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInitials(val)
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
    res["manager"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManager(val)
        return nil
    }
    res["middleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMiddleName(val)
        return nil
    }
    res["mobilePhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMobilePhone(val)
        return nil
    }
    res["multiValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewMultiValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MultiValueLegacyExtendedProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MultiValueLegacyExtendedProperty))
        }
        m.SetMultiValueExtendedProperties(res)
        return nil
    }
    res["nickName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNickName(val)
        return nil
    }
    res["officeLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOfficeLocation(val)
        return nil
    }
    res["otherAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPhysicalAddress() })
        if err != nil {
            return err
        }
        m.SetOtherAddress(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress))
        return nil
    }
    res["parentFolderId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetParentFolderId(val)
        return nil
    }
    res["personalNotes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPersonalNotes(val)
        return nil
    }
    res["photo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewProfilePhoto() })
        if err != nil {
            return err
        }
        m.SetPhoto(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ProfilePhoto))
        return nil
    }
    res["profession"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProfession(val)
        return nil
    }
    res["singleValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewSingleValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SingleValueLegacyExtendedProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SingleValueLegacyExtendedProperty))
        }
        m.SetSingleValueExtendedProperties(res)
        return nil
    }
    res["spouseName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSpouseName(val)
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
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTitle(val)
        return nil
    }
    res["yomiCompanyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetYomiCompanyName(val)
        return nil
    }
    res["yomiGivenName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetYomiGivenName(val)
        return nil
    }
    res["yomiSurname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetYomiSurname(val)
        return nil
    }
    return res
}
func (m *Delta) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Delta) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OutlookItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assistantName", m.GetAssistantName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("birthday", m.GetBirthday())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("businessAddress", m.GetBusinessAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("businessHomePage", m.GetBusinessHomePage())
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
        err = writer.WriteCollectionOfStringValues("children", m.GetChildren())
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEmailAddresses()))
        for i, v := range m.GetEmailAddresses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("emailAddresses", cast)
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
        err = writer.WriteStringValue("fileAs", m.GetFileAs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("generation", m.GetGeneration())
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
        err = writer.WriteObjectValue("homeAddress", m.GetHomeAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("homePhones", m.GetHomePhones())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("imAddresses", m.GetImAddresses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initials", m.GetInitials())
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
        err = writer.WriteStringValue("manager", m.GetManager())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("middleName", m.GetMiddleName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mobilePhone", m.GetMobilePhone())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("nickName", m.GetNickName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("officeLocation", m.GetOfficeLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("otherAddress", m.GetOtherAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentFolderId", m.GetParentFolderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("personalNotes", m.GetPersonalNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("photo", m.GetPhoto())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("profession", m.GetProfession())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("spouseName", m.GetSpouseName())
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
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("yomiCompanyName", m.GetYomiCompanyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("yomiGivenName", m.GetYomiGivenName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("yomiSurname", m.GetYomiSurname())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assistantName property value. The name of the contact's assistant.
// Parameters:
//  - value : Value to set for the assistantName property.
func (m *Delta) SetAssistantName(value *string)() {
    m.assistantName = value
}
// Sets the birthday property value. The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the birthday property.
func (m *Delta) SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.birthday = value
}
// Sets the businessAddress property value. The contact's business address.
// Parameters:
//  - value : Value to set for the businessAddress property.
func (m *Delta) SetBusinessAddress(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress)() {
    m.businessAddress = value
}
// Sets the businessHomePage property value. The business home page of the contact.
// Parameters:
//  - value : Value to set for the businessHomePage property.
func (m *Delta) SetBusinessHomePage(value *string)() {
    m.businessHomePage = value
}
// Sets the businessPhones property value. The contact's business phone numbers.
// Parameters:
//  - value : Value to set for the businessPhones property.
func (m *Delta) SetBusinessPhones(value []string)() {
    m.businessPhones = value
}
// Sets the children property value. The names of the contact's children.
// Parameters:
//  - value : Value to set for the children property.
func (m *Delta) SetChildren(value []string)() {
    m.children = value
}
// Sets the companyName property value. The name of the contact's company.
// Parameters:
//  - value : Value to set for the companyName property.
func (m *Delta) SetCompanyName(value *string)() {
    m.companyName = value
}
// Sets the department property value. The contact's department.
// Parameters:
//  - value : Value to set for the department property.
func (m *Delta) SetDepartment(value *string)() {
    m.department = value
}
// Sets the displayName property value. The contact's display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Delta) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the emailAddresses property value. The contact's email addresses.
// Parameters:
//  - value : Value to set for the emailAddresses property.
func (m *Delta) SetEmailAddresses(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EmailAddress)() {
    m.emailAddresses = value
}
// Sets the extensions property value. The collection of open extensions defined for the contact. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the extensions property.
func (m *Delta) SetExtensions(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Extension)() {
    m.extensions = value
}
// Sets the fileAs property value. The name the contact is filed under.
// Parameters:
//  - value : Value to set for the fileAs property.
func (m *Delta) SetFileAs(value *string)() {
    m.fileAs = value
}
// Sets the generation property value. The contact's generation.
// Parameters:
//  - value : Value to set for the generation property.
func (m *Delta) SetGeneration(value *string)() {
    m.generation = value
}
// Sets the givenName property value. The contact's given name.
// Parameters:
//  - value : Value to set for the givenName property.
func (m *Delta) SetGivenName(value *string)() {
    m.givenName = value
}
// Sets the homeAddress property value. The contact's home address.
// Parameters:
//  - value : Value to set for the homeAddress property.
func (m *Delta) SetHomeAddress(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress)() {
    m.homeAddress = value
}
// Sets the homePhones property value. The contact's home phone numbers.
// Parameters:
//  - value : Value to set for the homePhones property.
func (m *Delta) SetHomePhones(value []string)() {
    m.homePhones = value
}
// Sets the imAddresses property value. 
// Parameters:
//  - value : Value to set for the imAddresses property.
func (m *Delta) SetImAddresses(value []string)() {
    m.imAddresses = value
}
// Sets the initials property value. 
// Parameters:
//  - value : Value to set for the initials property.
func (m *Delta) SetInitials(value *string)() {
    m.initials = value
}
// Sets the jobTitle property value. 
// Parameters:
//  - value : Value to set for the jobTitle property.
func (m *Delta) SetJobTitle(value *string)() {
    m.jobTitle = value
}
// Sets the manager property value. 
// Parameters:
//  - value : Value to set for the manager property.
func (m *Delta) SetManager(value *string)() {
    m.manager = value
}
// Sets the middleName property value. 
// Parameters:
//  - value : Value to set for the middleName property.
func (m *Delta) SetMiddleName(value *string)() {
    m.middleName = value
}
// Sets the mobilePhone property value. 
// Parameters:
//  - value : Value to set for the mobilePhone property.
func (m *Delta) SetMobilePhone(value *string)() {
    m.mobilePhone = value
}
// Sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the multiValueExtendedProperties property.
func (m *Delta) SetMultiValueExtendedProperties(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MultiValueLegacyExtendedProperty)() {
    m.multiValueExtendedProperties = value
}
// Sets the nickName property value. 
// Parameters:
//  - value : Value to set for the nickName property.
func (m *Delta) SetNickName(value *string)() {
    m.nickName = value
}
// Sets the officeLocation property value. 
// Parameters:
//  - value : Value to set for the officeLocation property.
func (m *Delta) SetOfficeLocation(value *string)() {
    m.officeLocation = value
}
// Sets the otherAddress property value. 
// Parameters:
//  - value : Value to set for the otherAddress property.
func (m *Delta) SetOtherAddress(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress)() {
    m.otherAddress = value
}
// Sets the parentFolderId property value. 
// Parameters:
//  - value : Value to set for the parentFolderId property.
func (m *Delta) SetParentFolderId(value *string)() {
    m.parentFolderId = value
}
// Sets the personalNotes property value. 
// Parameters:
//  - value : Value to set for the personalNotes property.
func (m *Delta) SetPersonalNotes(value *string)() {
    m.personalNotes = value
}
// Sets the photo property value. Optional contact picture. You can get or set a photo for a contact.
// Parameters:
//  - value : Value to set for the photo property.
func (m *Delta) SetPhoto(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ProfilePhoto)() {
    m.photo = value
}
// Sets the profession property value. 
// Parameters:
//  - value : Value to set for the profession property.
func (m *Delta) SetProfession(value *string)() {
    m.profession = value
}
// Sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the contact. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the singleValueExtendedProperties property.
func (m *Delta) SetSingleValueExtendedProperties(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SingleValueLegacyExtendedProperty)() {
    m.singleValueExtendedProperties = value
}
// Sets the spouseName property value. 
// Parameters:
//  - value : Value to set for the spouseName property.
func (m *Delta) SetSpouseName(value *string)() {
    m.spouseName = value
}
// Sets the surname property value. 
// Parameters:
//  - value : Value to set for the surname property.
func (m *Delta) SetSurname(value *string)() {
    m.surname = value
}
// Sets the title property value. 
// Parameters:
//  - value : Value to set for the title property.
func (m *Delta) SetTitle(value *string)() {
    m.title = value
}
// Sets the yomiCompanyName property value. 
// Parameters:
//  - value : Value to set for the yomiCompanyName property.
func (m *Delta) SetYomiCompanyName(value *string)() {
    m.yomiCompanyName = value
}
// Sets the yomiGivenName property value. 
// Parameters:
//  - value : Value to set for the yomiGivenName property.
func (m *Delta) SetYomiGivenName(value *string)() {
    m.yomiGivenName = value
}
// Sets the yomiSurname property value. 
// Parameters:
//  - value : Value to set for the yomiSurname property.
func (m *Delta) SetYomiSurname(value *string)() {
    m.yomiSurname = value
}
