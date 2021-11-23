package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// mobileApp 
type MobileApp struct {
    Entity
    // The list of group assignments for this mobile app.
    assignments []MobileAppAssignment;
    // The list of categories for this app.
    categories []MobileAppCategory;
    // The date and time the app was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The description of the app.
    description *string;
    // The developer of the app.
    developer *string;
    // The admin provided or imported title of the app.
    displayName *string;
    // The more information Url.
    informationUrl *string;
    // The value indicating whether the app is marked as featured by the admin.
    isFeatured *bool;
    // The large icon, to be displayed in the app details and used for upload of the icon.
    largeIcon *MimeContent;
    // The date and time the app was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Notes for the app.
    notes *string;
    // The owner of the app.
    owner *string;
    // The privacy statement Url.
    privacyInformationUrl *string;
    // The publisher of the app.
    publisher *string;
    // The publishing state for the app. The app cannot be assigned unless the app is published. Possible values are: notPublished, processing, published.
    publishingState *MobileAppPublishingState;
}
// NewMobileApp instantiates a new mobileApp and sets the default values.
func NewMobileApp()(*MobileApp) {
    m := &MobileApp{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignments gets the assignments property value. The list of group assignments for this mobile app.
func (m *MobileApp) GetAssignments()([]MobileAppAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCategories gets the categories property value. The list of categories for this app.
func (m *MobileApp) GetCategories()([]MobileAppCategory) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the app was created.
func (m *MobileApp) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. The description of the app.
func (m *MobileApp) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDeveloper gets the developer property value. The developer of the app.
func (m *MobileApp) GetDeveloper()(*string) {
    if m == nil {
        return nil
    } else {
        return m.developer
    }
}
// GetDisplayName gets the displayName property value. The admin provided or imported title of the app.
func (m *MobileApp) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetInformationUrl gets the informationUrl property value. The more information Url.
func (m *MobileApp) GetInformationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.informationUrl
    }
}
// GetIsFeatured gets the isFeatured property value. The value indicating whether the app is marked as featured by the admin.
func (m *MobileApp) GetIsFeatured()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFeatured
    }
}
// GetLargeIcon gets the largeIcon property value. The large icon, to be displayed in the app details and used for upload of the icon.
func (m *MobileApp) GetLargeIcon()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.largeIcon
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the app was last modified.
func (m *MobileApp) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetNotes gets the notes property value. Notes for the app.
func (m *MobileApp) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// GetOwner gets the owner property value. The owner of the app.
func (m *MobileApp) GetOwner()(*string) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
// GetPrivacyInformationUrl gets the privacyInformationUrl property value. The privacy statement Url.
func (m *MobileApp) GetPrivacyInformationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.privacyInformationUrl
    }
}
// GetPublisher gets the publisher property value. The publisher of the app.
func (m *MobileApp) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// GetPublishingState gets the publishingState property value. The publishing state for the app. The app cannot be assigned unless the app is published. Possible values are: notPublished, processing, published.
func (m *MobileApp) GetPublishingState()(*MobileAppPublishingState) {
    if m == nil {
        return nil
    } else {
        return m.publishingState
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileApp) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*MobileAppAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppCategory, len(val))
            for i, v := range val {
                res[i] = *(v.(*MobileAppCategory))
            }
            m.SetCategories(res)
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
    res["developer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeveloper(val)
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
    res["isFeatured"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFeatured(val)
        }
        return nil
    }
    res["largeIcon"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMimeContent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLargeIcon(val.(*MimeContent))
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
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["owner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val)
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
    res["publishingState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppPublishingState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(MobileAppPublishingState)
            m.SetPublishingState(&cast)
        }
        return nil
    }
    return res
}
func (m *MobileApp) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MobileApp) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCategories()))
        for i, v := range m.GetCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("categories", cast)
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
        err = writer.WriteStringValue("developer", m.GetDeveloper())
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
        err = writer.WriteBoolValue("isFeatured", m.GetIsFeatured())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("largeIcon", m.GetLargeIcon())
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
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("owner", m.GetOwner())
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
        err = writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    if m.GetPublishingState() != nil {
        cast := m.GetPublishingState().String()
        err = writer.WriteStringValue("publishingState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of group assignments for this mobile app.
func (m *MobileApp) SetAssignments(value []MobileAppAssignment)() {
    m.assignments = value
}
// SetCategories sets the categories property value. The list of categories for this app.
func (m *MobileApp) SetCategories(value []MobileAppCategory)() {
    m.categories = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the app was created.
func (m *MobileApp) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description of the app.
func (m *MobileApp) SetDescription(value *string)() {
    m.description = value
}
// SetDeveloper sets the developer property value. The developer of the app.
func (m *MobileApp) SetDeveloper(value *string)() {
    m.developer = value
}
// SetDisplayName sets the displayName property value. The admin provided or imported title of the app.
func (m *MobileApp) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetInformationUrl sets the informationUrl property value. The more information Url.
func (m *MobileApp) SetInformationUrl(value *string)() {
    m.informationUrl = value
}
// SetIsFeatured sets the isFeatured property value. The value indicating whether the app is marked as featured by the admin.
func (m *MobileApp) SetIsFeatured(value *bool)() {
    m.isFeatured = value
}
// SetLargeIcon sets the largeIcon property value. The large icon, to be displayed in the app details and used for upload of the icon.
func (m *MobileApp) SetLargeIcon(value *MimeContent)() {
    m.largeIcon = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the app was last modified.
func (m *MobileApp) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNotes sets the notes property value. Notes for the app.
func (m *MobileApp) SetNotes(value *string)() {
    m.notes = value
}
// SetOwner sets the owner property value. The owner of the app.
func (m *MobileApp) SetOwner(value *string)() {
    m.owner = value
}
// SetPrivacyInformationUrl sets the privacyInformationUrl property value. The privacy statement Url.
func (m *MobileApp) SetPrivacyInformationUrl(value *string)() {
    m.privacyInformationUrl = value
}
// SetPublisher sets the publisher property value. The publisher of the app.
func (m *MobileApp) SetPublisher(value *string)() {
    m.publisher = value
}
// SetPublishingState sets the publishingState property value. The publishing state for the app. The app cannot be assigned unless the app is published. Possible values are: notPublished, processing, published.
func (m *MobileApp) SetPublishingState(value *MobileAppPublishingState)() {
    m.publishingState = value
}
