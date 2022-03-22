package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedEBook 
type ManagedEBook struct {
    Entity
    // The list of assignments for this eBook.
    assignments []ManagedEBookAssignmentable;
    // The date and time when the eBook file was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Description.
    description *string;
    // The list of installation states for this eBook.
    deviceStates []DeviceInstallStateable;
    // Name of the eBook.
    displayName *string;
    // The more information Url.
    informationUrl *string;
    // Mobile App Install Summary.
    installSummary EBookInstallSummaryable;
    // Book cover.
    largeCover MimeContentable;
    // The date and time when the eBook was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The privacy statement Url.
    privacyInformationUrl *string;
    // The date and time when the eBook was published.
    publishedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Publisher.
    publisher *string;
    // The list of installation states for this eBook.
    userStateSummary []UserInstallStateSummaryable;
}
// NewManagedEBook instantiates a new managedEBook and sets the default values.
func NewManagedEBook()(*ManagedEBook) {
    m := &ManagedEBook{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedEBookFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedEBookFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewManagedEBook(), nil
}
// GetAssignments gets the assignments property value. The list of assignments for this eBook.
func (m *ManagedEBook) GetAssignments()([]ManagedEBookAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the eBook file was created.
func (m *ManagedEBook) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. Description.
func (m *ManagedEBook) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDeviceStates gets the deviceStates property value. The list of installation states for this eBook.
func (m *ManagedEBook) GetDeviceStates()([]DeviceInstallStateable) {
    if m == nil {
        return nil
    } else {
        return m.deviceStates
    }
}
// GetDisplayName gets the displayName property value. Name of the eBook.
func (m *ManagedEBook) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedEBook) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedEBookAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedEBookAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedEBookAssignmentable)
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
        val, err := n.GetCollectionOfObjectValues(CreateDeviceInstallStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceInstallStateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceInstallStateable)
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
        val, err := n.GetObjectValue(CreateEBookInstallSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallSummary(val.(EBookInstallSummaryable))
        }
        return nil
    }
    res["largeCover"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateMimeContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLargeCover(val.(MimeContentable))
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
        val, err := n.GetCollectionOfObjectValues(CreateUserInstallStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserInstallStateSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(UserInstallStateSummaryable)
            }
            m.SetUserStateSummary(res)
        }
        return nil
    }
    return res
}
// GetInformationUrl gets the informationUrl property value. The more information Url.
func (m *ManagedEBook) GetInformationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.informationUrl
    }
}
// GetInstallSummary gets the installSummary property value. Mobile App Install Summary.
func (m *ManagedEBook) GetInstallSummary()(EBookInstallSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.installSummary
    }
}
// GetLargeCover gets the largeCover property value. Book cover.
func (m *ManagedEBook) GetLargeCover()(MimeContentable) {
    if m == nil {
        return nil
    } else {
        return m.largeCover
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the eBook was last modified.
func (m *ManagedEBook) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetPrivacyInformationUrl gets the privacyInformationUrl property value. The privacy statement Url.
func (m *ManagedEBook) GetPrivacyInformationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.privacyInformationUrl
    }
}
// GetPublishedDateTime gets the publishedDateTime property value. The date and time when the eBook was published.
func (m *ManagedEBook) GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.publishedDateTime
    }
}
// GetPublisher gets the publisher property value. Publisher.
func (m *ManagedEBook) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// GetUserStateSummary gets the userStateSummary property value. The list of installation states for this eBook.
func (m *ManagedEBook) GetUserStateSummary()([]UserInstallStateSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.userStateSummary
    }
}
// Serialize serializes information the current object
func (m *ManagedEBook) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetDeviceStates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceStates()))
        for i, v := range m.GetDeviceStates() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetUserStateSummary() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserStateSummary()))
        for i, v := range m.GetUserStateSummary() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userStateSummary", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of assignments for this eBook.
func (m *ManagedEBook) SetAssignments(value []ManagedEBookAssignmentable)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the eBook file was created.
func (m *ManagedEBook) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. Description.
func (m *ManagedEBook) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDeviceStates sets the deviceStates property value. The list of installation states for this eBook.
func (m *ManagedEBook) SetDeviceStates(value []DeviceInstallStateable)() {
    if m != nil {
        m.deviceStates = value
    }
}
// SetDisplayName sets the displayName property value. Name of the eBook.
func (m *ManagedEBook) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetInformationUrl sets the informationUrl property value. The more information Url.
func (m *ManagedEBook) SetInformationUrl(value *string)() {
    if m != nil {
        m.informationUrl = value
    }
}
// SetInstallSummary sets the installSummary property value. Mobile App Install Summary.
func (m *ManagedEBook) SetInstallSummary(value EBookInstallSummaryable)() {
    if m != nil {
        m.installSummary = value
    }
}
// SetLargeCover sets the largeCover property value. Book cover.
func (m *ManagedEBook) SetLargeCover(value MimeContentable)() {
    if m != nil {
        m.largeCover = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the eBook was last modified.
func (m *ManagedEBook) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetPrivacyInformationUrl sets the privacyInformationUrl property value. The privacy statement Url.
func (m *ManagedEBook) SetPrivacyInformationUrl(value *string)() {
    if m != nil {
        m.privacyInformationUrl = value
    }
}
// SetPublishedDateTime sets the publishedDateTime property value. The date and time when the eBook was published.
func (m *ManagedEBook) SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.publishedDateTime = value
    }
}
// SetPublisher sets the publisher property value. Publisher.
func (m *ManagedEBook) SetPublisher(value *string)() {
    if m != nil {
        m.publisher = value
    }
}
// SetUserStateSummary sets the userStateSummary property value. The list of installation states for this eBook.
func (m *ManagedEBook) SetUserStateSummary(value []UserInstallStateSummaryable)() {
    if m != nil {
        m.userStateSummary = value
    }
}
