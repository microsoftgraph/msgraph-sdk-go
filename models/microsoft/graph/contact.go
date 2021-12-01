package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Contact 
type Contact struct {
    OutlookItem
    // The name of the contact's assistant.
    assistantName *string;
    // The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    birthday *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The contact's business address.
    businessAddress *PhysicalAddress;
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
    emailAddresses []EmailAddress;
    // The collection of open extensions defined for the contact. Read-only. Nullable.
    extensions []Extension;
    // The name the contact is filed under.
    fileAs *string;
    // The contact's generation.
    generation *string;
    // The contact's given name.
    givenName *string;
    // The contact's home address.
    homeAddress *PhysicalAddress;
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
    multiValueExtendedProperties []MultiValueLegacyExtendedProperty;
    // 
    nickName *string;
    // 
    officeLocation *string;
    // 
    otherAddress *PhysicalAddress;
    // 
    parentFolderId *string;
    // 
    personalNotes *string;
    // Optional contact picture. You can get or set a photo for a contact.
    photo *ProfilePhoto;
    // 
    profession *string;
    // The collection of single-value extended properties defined for the contact. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedProperty;
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
// NewContact instantiates a new contact and sets the default values.
func NewContact()(*Contact) {
    m := &Contact{
        OutlookItem: *NewOutlookItem(),
    }
    return m
}
// GetAssistantName gets the assistantName property value. The name of the contact's assistant.
func (m *Contact) GetAssistantName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assistantName
    }
}
// GetBirthday gets the birthday property value. The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Contact) GetBirthday()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.birthday
    }
}
// GetBusinessAddress gets the businessAddress property value. The contact's business address.
func (m *Contact) GetBusinessAddress()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.businessAddress
    }
}
// GetBusinessHomePage gets the businessHomePage property value. The business home page of the contact.
func (m *Contact) GetBusinessHomePage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.businessHomePage
    }
}
// GetBusinessPhones gets the businessPhones property value. The contact's business phone numbers.
func (m *Contact) GetBusinessPhones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.businessPhones
    }
}
// GetChildren gets the children property value. The names of the contact's children.
func (m *Contact) GetChildren()([]string) {
    if m == nil {
        return nil
    } else {
        return m.children
    }
}
// GetCompanyName gets the companyName property value. The name of the contact's company.
func (m *Contact) GetCompanyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.companyName
    }
}
// GetDepartment gets the department property value. The contact's department.
func (m *Contact) GetDepartment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.department
    }
}
// GetDisplayName gets the displayName property value. The contact's display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation.
func (m *Contact) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEmailAddresses gets the emailAddresses property value. The contact's email addresses.
func (m *Contact) GetEmailAddresses()([]EmailAddress) {
    if m == nil {
        return nil
    } else {
        return m.emailAddresses
    }
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the contact. Read-only. Nullable.
func (m *Contact) GetExtensions()([]Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// GetFileAs gets the fileAs property value. The name the contact is filed under.
func (m *Contact) GetFileAs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileAs
    }
}
// GetGeneration gets the generation property value. The contact's generation.
func (m *Contact) GetGeneration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.generation
    }
}
// GetGivenName gets the givenName property value. The contact's given name.
func (m *Contact) GetGivenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.givenName
    }
}
// GetHomeAddress gets the homeAddress property value. The contact's home address.
func (m *Contact) GetHomeAddress()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.homeAddress
    }
}
// GetHomePhones gets the homePhones property value. The contact's home phone numbers.
func (m *Contact) GetHomePhones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.homePhones
    }
}
// GetImAddresses gets the imAddresses property value. 
func (m *Contact) GetImAddresses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.imAddresses
    }
}
// GetInitials gets the initials property value. 
func (m *Contact) GetInitials()(*string) {
    if m == nil {
        return nil
    } else {
        return m.initials
    }
}
// GetJobTitle gets the jobTitle property value. 
func (m *Contact) GetJobTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jobTitle
    }
}
// GetManager gets the manager property value. 
func (m *Contact) GetManager()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manager
    }
}
// GetMiddleName gets the middleName property value. 
func (m *Contact) GetMiddleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.middleName
    }
}
// GetMobilePhone gets the mobilePhone property value. 
func (m *Contact) GetMobilePhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mobilePhone
    }
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
func (m *Contact) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// GetNickName gets the nickName property value. 
func (m *Contact) GetNickName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nickName
    }
}
// GetOfficeLocation gets the officeLocation property value. 
func (m *Contact) GetOfficeLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.officeLocation
    }
}
// GetOtherAddress gets the otherAddress property value. 
func (m *Contact) GetOtherAddress()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.otherAddress
    }
}
// GetParentFolderId gets the parentFolderId property value. 
func (m *Contact) GetParentFolderId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentFolderId
    }
}
// GetPersonalNotes gets the personalNotes property value. 
func (m *Contact) GetPersonalNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.personalNotes
    }
}
// GetPhoto gets the photo property value. Optional contact picture. You can get or set a photo for a contact.
func (m *Contact) GetPhoto()(*ProfilePhoto) {
    if m == nil {
        return nil
    } else {
        return m.photo
    }
}
// GetProfession gets the profession property value. 
func (m *Contact) GetProfession()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profession
    }
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the contact. Read-only. Nullable.
func (m *Contact) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
}
// GetSpouseName gets the spouseName property value. 
func (m *Contact) GetSpouseName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.spouseName
    }
}
// GetSurname gets the surname property value. 
func (m *Contact) GetSurname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.surname
    }
}
// GetTitle gets the title property value. 
func (m *Contact) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// GetYomiCompanyName gets the yomiCompanyName property value. 
func (m *Contact) GetYomiCompanyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yomiCompanyName
    }
}
// GetYomiGivenName gets the yomiGivenName property value. 
func (m *Contact) GetYomiGivenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yomiGivenName
    }
}
// GetYomiSurname gets the yomiSurname property value. 
func (m *Contact) GetYomiSurname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yomiSurname
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Contact) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessAddress(val.(*PhysicalAddress))
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
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmailAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmailAddress, len(val))
            for i, v := range val {
                res[i] = *(v.(*EmailAddress))
            }
            m.SetEmailAddresses(res)
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeAddress(val.(*PhysicalAddress))
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
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMultiValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MultiValueLegacyExtendedProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*MultiValueLegacyExtendedProperty))
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOtherAddress(val.(*PhysicalAddress))
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProfilePhoto() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoto(val.(*ProfilePhoto))
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
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSingleValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SingleValueLegacyExtendedProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*SingleValueLegacyExtendedProperty))
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
func (m *Contact) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Contact) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *Contact) SetAssistantName(value *string)() {
    if m != nil {
        m.assistantName = value
    }
}
// SetBirthday sets the birthday property value. The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Contact) SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.birthday = value
    }
}
// SetBusinessAddress sets the businessAddress property value. The contact's business address.
func (m *Contact) SetBusinessAddress(value *PhysicalAddress)() {
    if m != nil {
        m.businessAddress = value
    }
}
// SetBusinessHomePage sets the businessHomePage property value. The business home page of the contact.
func (m *Contact) SetBusinessHomePage(value *string)() {
    if m != nil {
        m.businessHomePage = value
    }
}
// SetBusinessPhones sets the businessPhones property value. The contact's business phone numbers.
func (m *Contact) SetBusinessPhones(value []string)() {
    if m != nil {
        m.businessPhones = value
    }
}
// SetChildren sets the children property value. The names of the contact's children.
func (m *Contact) SetChildren(value []string)() {
    if m != nil {
        m.children = value
    }
}
// SetCompanyName sets the companyName property value. The name of the contact's company.
func (m *Contact) SetCompanyName(value *string)() {
    if m != nil {
        m.companyName = value
    }
}
// SetDepartment sets the department property value. The contact's department.
func (m *Contact) SetDepartment(value *string)() {
    if m != nil {
        m.department = value
    }
}
// SetDisplayName sets the displayName property value. The contact's display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation.
func (m *Contact) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEmailAddresses sets the emailAddresses property value. The contact's email addresses.
func (m *Contact) SetEmailAddresses(value []EmailAddress)() {
    if m != nil {
        m.emailAddresses = value
    }
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the contact. Read-only. Nullable.
func (m *Contact) SetExtensions(value []Extension)() {
    if m != nil {
        m.extensions = value
    }
}
// SetFileAs sets the fileAs property value. The name the contact is filed under.
func (m *Contact) SetFileAs(value *string)() {
    if m != nil {
        m.fileAs = value
    }
}
// SetGeneration sets the generation property value. The contact's generation.
func (m *Contact) SetGeneration(value *string)() {
    if m != nil {
        m.generation = value
    }
}
// SetGivenName sets the givenName property value. The contact's given name.
func (m *Contact) SetGivenName(value *string)() {
    if m != nil {
        m.givenName = value
    }
}
// SetHomeAddress sets the homeAddress property value. The contact's home address.
func (m *Contact) SetHomeAddress(value *PhysicalAddress)() {
    if m != nil {
        m.homeAddress = value
    }
}
// SetHomePhones sets the homePhones property value. The contact's home phone numbers.
func (m *Contact) SetHomePhones(value []string)() {
    if m != nil {
        m.homePhones = value
    }
}
// SetImAddresses sets the imAddresses property value. 
func (m *Contact) SetImAddresses(value []string)() {
    if m != nil {
        m.imAddresses = value
    }
}
// SetInitials sets the initials property value. 
func (m *Contact) SetInitials(value *string)() {
    if m != nil {
        m.initials = value
    }
}
// SetJobTitle sets the jobTitle property value. 
func (m *Contact) SetJobTitle(value *string)() {
    if m != nil {
        m.jobTitle = value
    }
}
// SetManager sets the manager property value. 
func (m *Contact) SetManager(value *string)() {
    if m != nil {
        m.manager = value
    }
}
// SetMiddleName sets the middleName property value. 
func (m *Contact) SetMiddleName(value *string)() {
    if m != nil {
        m.middleName = value
    }
}
// SetMobilePhone sets the mobilePhone property value. 
func (m *Contact) SetMobilePhone(value *string)() {
    if m != nil {
        m.mobilePhone = value
    }
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
func (m *Contact) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedProperty)() {
    if m != nil {
        m.multiValueExtendedProperties = value
    }
}
// SetNickName sets the nickName property value. 
func (m *Contact) SetNickName(value *string)() {
    if m != nil {
        m.nickName = value
    }
}
// SetOfficeLocation sets the officeLocation property value. 
func (m *Contact) SetOfficeLocation(value *string)() {
    if m != nil {
        m.officeLocation = value
    }
}
// SetOtherAddress sets the otherAddress property value. 
func (m *Contact) SetOtherAddress(value *PhysicalAddress)() {
    if m != nil {
        m.otherAddress = value
    }
}
// SetParentFolderId sets the parentFolderId property value. 
func (m *Contact) SetParentFolderId(value *string)() {
    if m != nil {
        m.parentFolderId = value
    }
}
// SetPersonalNotes sets the personalNotes property value. 
func (m *Contact) SetPersonalNotes(value *string)() {
    if m != nil {
        m.personalNotes = value
    }
}
// SetPhoto sets the photo property value. Optional contact picture. You can get or set a photo for a contact.
func (m *Contact) SetPhoto(value *ProfilePhoto)() {
    if m != nil {
        m.photo = value
    }
}
// SetProfession sets the profession property value. 
func (m *Contact) SetProfession(value *string)() {
    if m != nil {
        m.profession = value
    }
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the contact. Read-only. Nullable.
func (m *Contact) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedProperty)() {
    if m != nil {
        m.singleValueExtendedProperties = value
    }
}
// SetSpouseName sets the spouseName property value. 
func (m *Contact) SetSpouseName(value *string)() {
    if m != nil {
        m.spouseName = value
    }
}
// SetSurname sets the surname property value. 
func (m *Contact) SetSurname(value *string)() {
    if m != nil {
        m.surname = value
    }
}
// SetTitle sets the title property value. 
func (m *Contact) SetTitle(value *string)() {
    if m != nil {
        m.title = value
    }
}
// SetYomiCompanyName sets the yomiCompanyName property value. 
func (m *Contact) SetYomiCompanyName(value *string)() {
    if m != nil {
        m.yomiCompanyName = value
    }
}
// SetYomiGivenName sets the yomiGivenName property value. 
func (m *Contact) SetYomiGivenName(value *string)() {
    if m != nil {
        m.yomiGivenName = value
    }
}
// SetYomiSurname sets the yomiSurname property value. 
func (m *Contact) SetYomiSurname(value *string)() {
    if m != nil {
        m.yomiSurname = value
    }
}
