package delta

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// Delta 
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
    // The collection of open extensions defined for the contact. Nullable.
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
// NewDelta instantiates a new delta and sets the default values.
func NewDelta()(*Delta) {
    m := &Delta{
        OutlookItem: *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOutlookItem(),
    }
    return m
}
// GetAssistantName gets the assistantName property value. The name of the contact's assistant.
func (m *Delta) GetAssistantName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assistantName
    }
}
// GetBirthday gets the birthday property value. The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Delta) GetBirthday()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.birthday
    }
}
// GetBusinessAddress gets the businessAddress property value. The contact's business address.
func (m *Delta) GetBusinessAddress()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.businessAddress
    }
}
// GetBusinessHomePage gets the businessHomePage property value. The business home page of the contact.
func (m *Delta) GetBusinessHomePage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.businessHomePage
    }
}
// GetBusinessPhones gets the businessPhones property value. The contact's business phone numbers.
func (m *Delta) GetBusinessPhones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.businessPhones
    }
}
// GetChildren gets the children property value. The names of the contact's children.
func (m *Delta) GetChildren()([]string) {
    if m == nil {
        return nil
    } else {
        return m.children
    }
}
// GetCompanyName gets the companyName property value. The name of the contact's company.
func (m *Delta) GetCompanyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.companyName
    }
}
// GetDepartment gets the department property value. The contact's department.
func (m *Delta) GetDepartment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.department
    }
}
// GetDisplayName gets the displayName property value. The contact's display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation.
func (m *Delta) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEmailAddresses gets the emailAddresses property value. The contact's email addresses.
func (m *Delta) GetEmailAddresses()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EmailAddress) {
    if m == nil {
        return nil
    } else {
        return m.emailAddresses
    }
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the contact. Nullable.
func (m *Delta) GetExtensions()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// GetFileAs gets the fileAs property value. The name the contact is filed under.
func (m *Delta) GetFileAs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileAs
    }
}
// GetGeneration gets the generation property value. The contact's generation.
func (m *Delta) GetGeneration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.generation
    }
}
// GetGivenName gets the givenName property value. The contact's given name.
func (m *Delta) GetGivenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.givenName
    }
}
// GetHomeAddress gets the homeAddress property value. The contact's home address.
func (m *Delta) GetHomeAddress()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.homeAddress
    }
}
// GetHomePhones gets the homePhones property value. The contact's home phone numbers.
func (m *Delta) GetHomePhones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.homePhones
    }
}
// GetImAddresses gets the imAddresses property value. 
func (m *Delta) GetImAddresses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.imAddresses
    }
}
// GetInitials gets the initials property value. 
func (m *Delta) GetInitials()(*string) {
    if m == nil {
        return nil
    } else {
        return m.initials
    }
}
// GetJobTitle gets the jobTitle property value. 
func (m *Delta) GetJobTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jobTitle
    }
}
// GetManager gets the manager property value. 
func (m *Delta) GetManager()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manager
    }
}
// GetMiddleName gets the middleName property value. 
func (m *Delta) GetMiddleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.middleName
    }
}
// GetMobilePhone gets the mobilePhone property value. 
func (m *Delta) GetMobilePhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mobilePhone
    }
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
func (m *Delta) GetMultiValueExtendedProperties()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MultiValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// GetNickName gets the nickName property value. 
func (m *Delta) GetNickName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nickName
    }
}
// GetOfficeLocation gets the officeLocation property value. 
func (m *Delta) GetOfficeLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.officeLocation
    }
}
// GetOtherAddress gets the otherAddress property value. 
func (m *Delta) GetOtherAddress()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.otherAddress
    }
}
// GetParentFolderId gets the parentFolderId property value. 
func (m *Delta) GetParentFolderId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentFolderId
    }
}
// GetPersonalNotes gets the personalNotes property value. 
func (m *Delta) GetPersonalNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.personalNotes
    }
}
// GetPhoto gets the photo property value. Optional contact picture. You can get or set a photo for a contact.
func (m *Delta) GetPhoto()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ProfilePhoto) {
    if m == nil {
        return nil
    } else {
        return m.photo
    }
}
// GetProfession gets the profession property value. 
func (m *Delta) GetProfession()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profession
    }
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the contact. Read-only. Nullable.
func (m *Delta) GetSingleValueExtendedProperties()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SingleValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
}
// GetSpouseName gets the spouseName property value. 
func (m *Delta) GetSpouseName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.spouseName
    }
}
// GetSurname gets the surname property value. 
func (m *Delta) GetSurname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.surname
    }
}
// GetTitle gets the title property value. 
func (m *Delta) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// GetYomiCompanyName gets the yomiCompanyName property value. 
func (m *Delta) GetYomiCompanyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yomiCompanyName
    }
}
// GetYomiGivenName gets the yomiGivenName property value. 
func (m *Delta) GetYomiGivenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yomiGivenName
    }
}
// GetYomiSurname gets the yomiSurname property value. 
func (m *Delta) GetYomiSurname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yomiSurname
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Delta) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["assistantName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssistantName(val)
        }
        return nil
    }
    res["birthday"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBirthday(val)
        }
        return nil
    }
    res["businessAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPhysicalAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessAddress(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress))
        }
        return nil
    }
    res["businessHomePage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessHomePage(val)
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
    res["children"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetChildren(res)
        }
        return nil
    }
    res["companyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyName(val)
        }
        return nil
    }
    res["department"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepartment(val)
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
    res["emailAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEmailAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EmailAddress, len(val))
            for i, v := range val {
                res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EmailAddress))
            }
            m.SetEmailAddresses(res)
        }
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewExtension() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Extension, len(val))
            for i, v := range val {
                res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Extension))
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["fileAs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileAs(val)
        }
        return nil
    }
    res["generation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeneration(val)
        }
        return nil
    }
    res["givenName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGivenName(val)
        }
        return nil
    }
    res["homeAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPhysicalAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeAddress(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress))
        }
        return nil
    }
    res["homePhones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetHomePhones(res)
        }
        return nil
    }
    res["imAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetImAddresses(res)
        }
        return nil
    }
    res["initials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitials(val)
        }
        return nil
    }
    res["jobTitle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobTitle(val)
        }
        return nil
    }
    res["manager"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManager(val)
        }
        return nil
    }
    res["middleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiddleName(val)
        }
        return nil
    }
    res["mobilePhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobilePhone(val)
        }
        return nil
    }
    res["multiValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewMultiValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MultiValueLegacyExtendedProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MultiValueLegacyExtendedProperty))
            }
            m.SetMultiValueExtendedProperties(res)
        }
        return nil
    }
    res["nickName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNickName(val)
        }
        return nil
    }
    res["officeLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeLocation(val)
        }
        return nil
    }
    res["otherAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPhysicalAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOtherAddress(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress))
        }
        return nil
    }
    res["parentFolderId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentFolderId(val)
        }
        return nil
    }
    res["personalNotes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalNotes(val)
        }
        return nil
    }
    res["photo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewProfilePhoto() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoto(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ProfilePhoto))
        }
        return nil
    }
    res["profession"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfession(val)
        }
        return nil
    }
    res["singleValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewSingleValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SingleValueLegacyExtendedProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SingleValueLegacyExtendedProperty))
            }
            m.SetSingleValueExtendedProperties(res)
        }
        return nil
    }
    res["spouseName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpouseName(val)
        }
        return nil
    }
    res["surname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSurname(val)
        }
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["yomiCompanyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYomiCompanyName(val)
        }
        return nil
    }
    res["yomiGivenName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYomiGivenName(val)
        }
        return nil
    }
    res["yomiSurname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYomiSurname(val)
        }
        return nil
    }
    return res
}
func (m *Delta) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAssistantName sets the assistantName property value. The name of the contact's assistant.
func (m *Delta) SetAssistantName(value *string)() {
    if m != nil {
        m.assistantName = value
    }
}
// SetBirthday sets the birthday property value. The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Delta) SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.birthday = value
    }
}
// SetBusinessAddress sets the businessAddress property value. The contact's business address.
func (m *Delta) SetBusinessAddress(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress)() {
    if m != nil {
        m.businessAddress = value
    }
}
// SetBusinessHomePage sets the businessHomePage property value. The business home page of the contact.
func (m *Delta) SetBusinessHomePage(value *string)() {
    if m != nil {
        m.businessHomePage = value
    }
}
// SetBusinessPhones sets the businessPhones property value. The contact's business phone numbers.
func (m *Delta) SetBusinessPhones(value []string)() {
    if m != nil {
        m.businessPhones = value
    }
}
// SetChildren sets the children property value. The names of the contact's children.
func (m *Delta) SetChildren(value []string)() {
    if m != nil {
        m.children = value
    }
}
// SetCompanyName sets the companyName property value. The name of the contact's company.
func (m *Delta) SetCompanyName(value *string)() {
    if m != nil {
        m.companyName = value
    }
}
// SetDepartment sets the department property value. The contact's department.
func (m *Delta) SetDepartment(value *string)() {
    if m != nil {
        m.department = value
    }
}
// SetDisplayName sets the displayName property value. The contact's display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation.
func (m *Delta) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEmailAddresses sets the emailAddresses property value. The contact's email addresses.
func (m *Delta) SetEmailAddresses(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EmailAddress)() {
    if m != nil {
        m.emailAddresses = value
    }
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the contact. Nullable.
func (m *Delta) SetExtensions(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Extension)() {
    if m != nil {
        m.extensions = value
    }
}
// SetFileAs sets the fileAs property value. The name the contact is filed under.
func (m *Delta) SetFileAs(value *string)() {
    if m != nil {
        m.fileAs = value
    }
}
// SetGeneration sets the generation property value. The contact's generation.
func (m *Delta) SetGeneration(value *string)() {
    if m != nil {
        m.generation = value
    }
}
// SetGivenName sets the givenName property value. The contact's given name.
func (m *Delta) SetGivenName(value *string)() {
    if m != nil {
        m.givenName = value
    }
}
// SetHomeAddress sets the homeAddress property value. The contact's home address.
func (m *Delta) SetHomeAddress(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress)() {
    if m != nil {
        m.homeAddress = value
    }
}
// SetHomePhones sets the homePhones property value. The contact's home phone numbers.
func (m *Delta) SetHomePhones(value []string)() {
    if m != nil {
        m.homePhones = value
    }
}
// SetImAddresses sets the imAddresses property value. 
func (m *Delta) SetImAddresses(value []string)() {
    if m != nil {
        m.imAddresses = value
    }
}
// SetInitials sets the initials property value. 
func (m *Delta) SetInitials(value *string)() {
    if m != nil {
        m.initials = value
    }
}
// SetJobTitle sets the jobTitle property value. 
func (m *Delta) SetJobTitle(value *string)() {
    if m != nil {
        m.jobTitle = value
    }
}
// SetManager sets the manager property value. 
func (m *Delta) SetManager(value *string)() {
    if m != nil {
        m.manager = value
    }
}
// SetMiddleName sets the middleName property value. 
func (m *Delta) SetMiddleName(value *string)() {
    if m != nil {
        m.middleName = value
    }
}
// SetMobilePhone sets the mobilePhone property value. 
func (m *Delta) SetMobilePhone(value *string)() {
    if m != nil {
        m.mobilePhone = value
    }
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
func (m *Delta) SetMultiValueExtendedProperties(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MultiValueLegacyExtendedProperty)() {
    if m != nil {
        m.multiValueExtendedProperties = value
    }
}
// SetNickName sets the nickName property value. 
func (m *Delta) SetNickName(value *string)() {
    if m != nil {
        m.nickName = value
    }
}
// SetOfficeLocation sets the officeLocation property value. 
func (m *Delta) SetOfficeLocation(value *string)() {
    if m != nil {
        m.officeLocation = value
    }
}
// SetOtherAddress sets the otherAddress property value. 
func (m *Delta) SetOtherAddress(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PhysicalAddress)() {
    if m != nil {
        m.otherAddress = value
    }
}
// SetParentFolderId sets the parentFolderId property value. 
func (m *Delta) SetParentFolderId(value *string)() {
    if m != nil {
        m.parentFolderId = value
    }
}
// SetPersonalNotes sets the personalNotes property value. 
func (m *Delta) SetPersonalNotes(value *string)() {
    if m != nil {
        m.personalNotes = value
    }
}
// SetPhoto sets the photo property value. Optional contact picture. You can get or set a photo for a contact.
func (m *Delta) SetPhoto(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ProfilePhoto)() {
    if m != nil {
        m.photo = value
    }
}
// SetProfession sets the profession property value. 
func (m *Delta) SetProfession(value *string)() {
    if m != nil {
        m.profession = value
    }
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the contact. Read-only. Nullable.
func (m *Delta) SetSingleValueExtendedProperties(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SingleValueLegacyExtendedProperty)() {
    if m != nil {
        m.singleValueExtendedProperties = value
    }
}
// SetSpouseName sets the spouseName property value. 
func (m *Delta) SetSpouseName(value *string)() {
    if m != nil {
        m.spouseName = value
    }
}
// SetSurname sets the surname property value. 
func (m *Delta) SetSurname(value *string)() {
    if m != nil {
        m.surname = value
    }
}
// SetTitle sets the title property value. 
func (m *Delta) SetTitle(value *string)() {
    if m != nil {
        m.title = value
    }
}
// SetYomiCompanyName sets the yomiCompanyName property value. 
func (m *Delta) SetYomiCompanyName(value *string)() {
    if m != nil {
        m.yomiCompanyName = value
    }
}
// SetYomiGivenName sets the yomiGivenName property value. 
func (m *Delta) SetYomiGivenName(value *string)() {
    if m != nil {
        m.yomiGivenName = value
    }
}
// SetYomiSurname sets the yomiSurname property value. 
func (m *Delta) SetYomiSurname(value *string)() {
    if m != nil {
        m.yomiSurname = value
    }
}
