package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagedEBook struct {
    Entity
    // The list of assignments for this eBook.
    assignments []ManagedEBookAssignment;
    // The date and time when the eBook file was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Description.
    description *string;
    // The list of installation states for this eBook.
    deviceStates []DeviceInstallState;
    // Name of the eBook.
    displayName *string;
    // The more information Url.
    informationUrl *string;
    // Mobile App Install Summary.
    installSummary *EBookInstallSummary;
    // Book cover.
    largeCover *MimeContent;
    // The date and time when the eBook was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The privacy statement Url.
    privacyInformationUrl *string;
    // The date and time when the eBook was published.
    publishedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Publisher.
    publisher *string;
    // The list of installation states for this eBook.
    userStateSummary []UserInstallStateSummary;
}
// Instantiates a new managedEBook and sets the default values.
func NewManagedEBook()(*ManagedEBook) {
    m := &ManagedEBook{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. The list of assignments for this eBook.
func (m *ManagedEBook) GetAssignments()([]ManagedEBookAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the createdDateTime property value. The date and time when the eBook file was created.
func (m *ManagedEBook) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Description.
func (m *ManagedEBook) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the deviceStates property value. The list of installation states for this eBook.
func (m *ManagedEBook) GetDeviceStates()([]DeviceInstallState) {
    if m == nil {
        return nil
    } else {
        return m.deviceStates
    }
}
// Gets the displayName property value. Name of the eBook.
func (m *ManagedEBook) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the informationUrl property value. The more information Url.
func (m *ManagedEBook) GetInformationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.informationUrl
    }
}
// Gets the installSummary property value. Mobile App Install Summary.
func (m *ManagedEBook) GetInstallSummary()(*EBookInstallSummary) {
    if m == nil {
        return nil
    } else {
        return m.installSummary
    }
}
// Gets the largeCover property value. Book cover.
func (m *ManagedEBook) GetLargeCover()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.largeCover
    }
}
// Gets the lastModifiedDateTime property value. The date and time when the eBook was last modified.
func (m *ManagedEBook) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the privacyInformationUrl property value. The privacy statement Url.
func (m *ManagedEBook) GetPrivacyInformationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.privacyInformationUrl
    }
}
// Gets the publishedDateTime property value. The date and time when the eBook was published.
func (m *ManagedEBook) GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.publishedDateTime
    }
}
// Gets the publisher property value. Publisher.
func (m *ManagedEBook) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// Gets the userStateSummary property value. The list of installation states for this eBook.
func (m *ManagedEBook) GetUserStateSummary()([]UserInstallStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.userStateSummary
    }
}
// The deserialization information for the current model
func (m *ManagedEBook) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedEBookAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedEBookAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedEBookAssignment))
            }
            m.SetAssignments(res)
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
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["deviceStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceInstallState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceInstallState, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceInstallState))
            }
            m.SetDeviceStates(res)
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
    res["informationUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInformationUrl(val)
        }
        return nil
    }
    res["installSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEBookInstallSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallSummary(val.(*EBookInstallSummary))
        }
        return nil
    }
    res["largeCover"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMimeContent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLargeCover(val.(*MimeContent))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["privacyInformationUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyInformationUrl(val)
        }
        return nil
    }
    res["publishedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedDateTime(val)
        }
        return nil
    }
    res["publisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["userStateSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserInstallStateSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserInstallStateSummary, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserInstallStateSummary))
            }
            m.SetUserStateSummary(res)
        }
        return nil
    }
    return res
}
func (m *ManagedEBook) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagedEBook) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
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
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceStates()))
        for i, v := range m.GetDeviceStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStates", cast)
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
        err = writer.WriteStringValue("informationUrl", m.GetInformationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("installSummary", m.GetInstallSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("largeCover", m.GetLargeCover())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("privacyInformationUrl", m.GetPrivacyInformationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("publishedDateTime", m.GetPublishedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserStateSummary()))
        for i, v := range m.GetUserStateSummary() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userStateSummary", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignments property value. The list of assignments for this eBook.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *ManagedEBook) SetAssignments(value []ManagedEBookAssignment)() {
    m.assignments = value
}
// Sets the createdDateTime property value. The date and time when the eBook file was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *ManagedEBook) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Description.
// Parameters:
//  - value : Value to set for the description property.
func (m *ManagedEBook) SetDescription(value *string)() {
    m.description = value
}
// Sets the deviceStates property value. The list of installation states for this eBook.
// Parameters:
//  - value : Value to set for the deviceStates property.
func (m *ManagedEBook) SetDeviceStates(value []DeviceInstallState)() {
    m.deviceStates = value
}
// Sets the displayName property value. Name of the eBook.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ManagedEBook) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the informationUrl property value. The more information Url.
// Parameters:
//  - value : Value to set for the informationUrl property.
func (m *ManagedEBook) SetInformationUrl(value *string)() {
    m.informationUrl = value
}
// Sets the installSummary property value. Mobile App Install Summary.
// Parameters:
//  - value : Value to set for the installSummary property.
func (m *ManagedEBook) SetInstallSummary(value *EBookInstallSummary)() {
    m.installSummary = value
}
// Sets the largeCover property value. Book cover.
// Parameters:
//  - value : Value to set for the largeCover property.
func (m *ManagedEBook) SetLargeCover(value *MimeContent)() {
    m.largeCover = value
}
// Sets the lastModifiedDateTime property value. The date and time when the eBook was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *ManagedEBook) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the privacyInformationUrl property value. The privacy statement Url.
// Parameters:
//  - value : Value to set for the privacyInformationUrl property.
func (m *ManagedEBook) SetPrivacyInformationUrl(value *string)() {
    m.privacyInformationUrl = value
}
// Sets the publishedDateTime property value. The date and time when the eBook was published.
// Parameters:
//  - value : Value to set for the publishedDateTime property.
func (m *ManagedEBook) SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.publishedDateTime = value
}
// Sets the publisher property value. Publisher.
// Parameters:
//  - value : Value to set for the publisher property.
func (m *ManagedEBook) SetPublisher(value *string)() {
    m.publisher = value
}
// Sets the userStateSummary property value. The list of installation states for this eBook.
// Parameters:
//  - value : Value to set for the userStateSummary property.
func (m *ManagedEBook) SetUserStateSummary(value []UserInstallStateSummary)() {
    m.userStateSummary = value
}
