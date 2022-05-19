package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Calendar casts the previous resource to user.
type Calendar struct {
    Entity
    // Represent the online meeting service providers that can be used to create online meetings in this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
    allowedOnlineMeetingProviders []string
    // The permissions of the users with whom the calendar is shared.
    calendarPermissions []CalendarPermissionable
    // The calendar view for the calendar. Navigation property. Read-only.
    calendarView []Eventable
    // true if the user can write to the calendar, false otherwise. This property is true for the user who created the calendar. This property is also true for a user who has been shared a calendar and granted write access, through an Outlook client or the corresponding calendarPermission resource. Read-only.
    canEdit *bool
    // true if the user has the permission to share the calendar, false otherwise. Only the user who created the calendar can share it. Read-only.
    canShare *bool
    // true if the user can read calendar items that have been marked private, false otherwise. This property is set through an Outlook client or the corresponding calendarPermission resource. Read-only.
    canViewPrivateItems *bool
    // Identifies the version of the calendar object. Every time the calendar is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
    changeKey *string
    // Specifies the color theme to distinguish the calendar from other calendars in a UI. The property values are: auto, lightBlue, lightGreen, lightOrange, lightGray, lightYellow, lightTeal, lightPink, lightBrown, lightRed, maxColor.
    color *CalendarColor
    // The default online meeting provider for meetings sent from this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
    defaultOnlineMeetingProvider *OnlineMeetingProviderType
    // The events in the calendar. Navigation property. Read-only.
    events []Eventable
    // The calendar color, expressed in a hex color code of three hexadecimal values, each ranging from 00 to FF and representing the red, green, or blue components of the color in the RGB color space. If the user has never explicitly set a color for the calendar, this property is  empty.
    hexColor *string
    // true if this is the default calendar where new events are created by default, false otherwise.
    isDefaultCalendar *bool
    // Indicates whether this user calendar can be deleted from the user mailbox.
    isRemovable *bool
    // Indicates whether this user calendar supports tracking of meeting responses. Only meeting invites sent from users' primary calendars support tracking of meeting responses.
    isTallyingResponses *bool
    // The collection of multi-value extended properties defined for the calendar. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable
    // The calendar name.
    name *string
    // If set, this represents the user who created or added the calendar. For a calendar that the user created or added, the owner property is set to the user. For a calendar shared with the user, the owner property is set to the person who shared that calendar with the user. Read-only.
    owner EmailAddressable
    // The collection of single-value extended properties defined for the calendar. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable
}
// NewCalendar instantiates a new calendar and sets the default values.
func NewCalendar()(*Calendar) {
    m := &Calendar{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCalendarFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCalendarFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCalendar(), nil
}
// GetAllowedOnlineMeetingProviders gets the allowedOnlineMeetingProviders property value. Represent the online meeting service providers that can be used to create online meetings in this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
func (m *Calendar) GetAllowedOnlineMeetingProviders()([]string) {
    if m == nil {
        return nil
    } else {
        return m.allowedOnlineMeetingProviders
    }
}
// GetCalendarPermissions gets the calendarPermissions property value. The permissions of the users with whom the calendar is shared.
func (m *Calendar) GetCalendarPermissions()([]CalendarPermissionable) {
    if m == nil {
        return nil
    } else {
        return m.calendarPermissions
    }
}
// GetCalendarView gets the calendarView property value. The calendar view for the calendar. Navigation property. Read-only.
func (m *Calendar) GetCalendarView()([]Eventable) {
    if m == nil {
        return nil
    } else {
        return m.calendarView
    }
}
// GetCanEdit gets the canEdit property value. true if the user can write to the calendar, false otherwise. This property is true for the user who created the calendar. This property is also true for a user who has been shared a calendar and granted write access, through an Outlook client or the corresponding calendarPermission resource. Read-only.
func (m *Calendar) GetCanEdit()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.canEdit
    }
}
// GetCanShare gets the canShare property value. true if the user has the permission to share the calendar, false otherwise. Only the user who created the calendar can share it. Read-only.
func (m *Calendar) GetCanShare()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.canShare
    }
}
// GetCanViewPrivateItems gets the canViewPrivateItems property value. true if the user can read calendar items that have been marked private, false otherwise. This property is set through an Outlook client or the corresponding calendarPermission resource. Read-only.
func (m *Calendar) GetCanViewPrivateItems()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.canViewPrivateItems
    }
}
// GetChangeKey gets the changeKey property value. Identifies the version of the calendar object. Every time the calendar is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
func (m *Calendar) GetChangeKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.changeKey
    }
}
// GetColor gets the color property value. Specifies the color theme to distinguish the calendar from other calendars in a UI. The property values are: auto, lightBlue, lightGreen, lightOrange, lightGray, lightYellow, lightTeal, lightPink, lightBrown, lightRed, maxColor.
func (m *Calendar) GetColor()(*CalendarColor) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
// GetDefaultOnlineMeetingProvider gets the defaultOnlineMeetingProvider property value. The default online meeting provider for meetings sent from this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
func (m *Calendar) GetDefaultOnlineMeetingProvider()(*OnlineMeetingProviderType) {
    if m == nil {
        return nil
    } else {
        return m.defaultOnlineMeetingProvider
    }
}
// GetEvents gets the events property value. The events in the calendar. Navigation property. Read-only.
func (m *Calendar) GetEvents()([]Eventable) {
    if m == nil {
        return nil
    } else {
        return m.events
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Calendar) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedOnlineMeetingProviders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedOnlineMeetingProviders(res)
        }
        return nil
    }
    res["calendarPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCalendarPermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CalendarPermissionable, len(val))
            for i, v := range val {
                res[i] = v.(CalendarPermissionable)
            }
            m.SetCalendarPermissions(res)
        }
        return nil
    }
    res["calendarView"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Eventable, len(val))
            for i, v := range val {
                res[i] = v.(Eventable)
            }
            m.SetCalendarView(res)
        }
        return nil
    }
    res["canEdit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanEdit(val)
        }
        return nil
    }
    res["canShare"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanShare(val)
        }
        return nil
    }
    res["canViewPrivateItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanViewPrivateItems(val)
        }
        return nil
    }
    res["changeKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeKey(val)
        }
        return nil
    }
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCalendarColor)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val.(*CalendarColor))
        }
        return nil
    }
    res["defaultOnlineMeetingProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnlineMeetingProviderType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultOnlineMeetingProvider(val.(*OnlineMeetingProviderType))
        }
        return nil
    }
    res["events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Eventable, len(val))
            for i, v := range val {
                res[i] = v.(Eventable)
            }
            m.SetEvents(res)
        }
        return nil
    }
    res["hexColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHexColor(val)
        }
        return nil
    }
    res["isDefaultCalendar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultCalendar(val)
        }
        return nil
    }
    res["isRemovable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRemovable(val)
        }
        return nil
    }
    res["isTallyingResponses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTallyingResponses(val)
        }
        return nil
    }
    res["multiValueExtendedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MultiValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(MultiValueLegacyExtendedPropertyable)
            }
            m.SetMultiValueExtendedProperties(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(EmailAddressable))
        }
        return nil
    }
    res["singleValueExtendedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SingleValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(SingleValueLegacyExtendedPropertyable)
            }
            m.SetSingleValueExtendedProperties(res)
        }
        return nil
    }
    return res
}
// GetHexColor gets the hexColor property value. The calendar color, expressed in a hex color code of three hexadecimal values, each ranging from 00 to FF and representing the red, green, or blue components of the color in the RGB color space. If the user has never explicitly set a color for the calendar, this property is  empty.
func (m *Calendar) GetHexColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hexColor
    }
}
// GetIsDefaultCalendar gets the isDefaultCalendar property value. true if this is the default calendar where new events are created by default, false otherwise.
func (m *Calendar) GetIsDefaultCalendar()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefaultCalendar
    }
}
// GetIsRemovable gets the isRemovable property value. Indicates whether this user calendar can be deleted from the user mailbox.
func (m *Calendar) GetIsRemovable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRemovable
    }
}
// GetIsTallyingResponses gets the isTallyingResponses property value. Indicates whether this user calendar supports tracking of meeting responses. Only meeting invites sent from users' primary calendars support tracking of meeting responses.
func (m *Calendar) GetIsTallyingResponses()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isTallyingResponses
    }
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the calendar. Read-only. Nullable.
func (m *Calendar) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// GetName gets the name property value. The calendar name.
func (m *Calendar) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetOwner gets the owner property value. If set, this represents the user who created or added the calendar. For a calendar that the user created or added, the owner property is set to the user. For a calendar shared with the user, the owner property is set to the person who shared that calendar with the user. Read-only.
func (m *Calendar) GetOwner()(EmailAddressable) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the calendar. Read-only. Nullable.
func (m *Calendar) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
}
// Serialize serializes information the current object
func (m *Calendar) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedOnlineMeetingProviders() != nil {
        err = writer.WriteCollectionOfStringValues("allowedOnlineMeetingProviders", m.GetAllowedOnlineMeetingProviders())
        if err != nil {
            return err
        }
    }
    if m.GetCalendarPermissions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCalendarPermissions()))
        for i, v := range m.GetCalendarPermissions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("calendarPermissions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCalendarView() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCalendarView()))
        for i, v := range m.GetCalendarView() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("calendarView", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("canEdit", m.GetCanEdit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("canShare", m.GetCanShare())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("canViewPrivateItems", m.GetCanViewPrivateItems())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("changeKey", m.GetChangeKey())
        if err != nil {
            return err
        }
    }
    if m.GetColor() != nil {
        cast := (*m.GetColor()).String()
        err = writer.WriteStringValue("color", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefaultOnlineMeetingProvider() != nil {
        cast := (*m.GetDefaultOnlineMeetingProvider()).String()
        err = writer.WriteStringValue("defaultOnlineMeetingProvider", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEvents()))
        for i, v := range m.GetEvents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("events", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("hexColor", m.GetHexColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefaultCalendar", m.GetIsDefaultCalendar())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRemovable", m.GetIsRemovable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isTallyingResponses", m.GetIsTallyingResponses())
        if err != nil {
            return err
        }
    }
    if m.GetMultiValueExtendedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    if m.GetSingleValueExtendedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedOnlineMeetingProviders sets the allowedOnlineMeetingProviders property value. Represent the online meeting service providers that can be used to create online meetings in this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
func (m *Calendar) SetAllowedOnlineMeetingProviders(value []string)() {
    if m != nil {
        m.allowedOnlineMeetingProviders = value
    }
}
// SetCalendarPermissions sets the calendarPermissions property value. The permissions of the users with whom the calendar is shared.
func (m *Calendar) SetCalendarPermissions(value []CalendarPermissionable)() {
    if m != nil {
        m.calendarPermissions = value
    }
}
// SetCalendarView sets the calendarView property value. The calendar view for the calendar. Navigation property. Read-only.
func (m *Calendar) SetCalendarView(value []Eventable)() {
    if m != nil {
        m.calendarView = value
    }
}
// SetCanEdit sets the canEdit property value. true if the user can write to the calendar, false otherwise. This property is true for the user who created the calendar. This property is also true for a user who has been shared a calendar and granted write access, through an Outlook client or the corresponding calendarPermission resource. Read-only.
func (m *Calendar) SetCanEdit(value *bool)() {
    if m != nil {
        m.canEdit = value
    }
}
// SetCanShare sets the canShare property value. true if the user has the permission to share the calendar, false otherwise. Only the user who created the calendar can share it. Read-only.
func (m *Calendar) SetCanShare(value *bool)() {
    if m != nil {
        m.canShare = value
    }
}
// SetCanViewPrivateItems sets the canViewPrivateItems property value. true if the user can read calendar items that have been marked private, false otherwise. This property is set through an Outlook client or the corresponding calendarPermission resource. Read-only.
func (m *Calendar) SetCanViewPrivateItems(value *bool)() {
    if m != nil {
        m.canViewPrivateItems = value
    }
}
// SetChangeKey sets the changeKey property value. Identifies the version of the calendar object. Every time the calendar is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
func (m *Calendar) SetChangeKey(value *string)() {
    if m != nil {
        m.changeKey = value
    }
}
// SetColor sets the color property value. Specifies the color theme to distinguish the calendar from other calendars in a UI. The property values are: auto, lightBlue, lightGreen, lightOrange, lightGray, lightYellow, lightTeal, lightPink, lightBrown, lightRed, maxColor.
func (m *Calendar) SetColor(value *CalendarColor)() {
    if m != nil {
        m.color = value
    }
}
// SetDefaultOnlineMeetingProvider sets the defaultOnlineMeetingProvider property value. The default online meeting provider for meetings sent from this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
func (m *Calendar) SetDefaultOnlineMeetingProvider(value *OnlineMeetingProviderType)() {
    if m != nil {
        m.defaultOnlineMeetingProvider = value
    }
}
// SetEvents sets the events property value. The events in the calendar. Navigation property. Read-only.
func (m *Calendar) SetEvents(value []Eventable)() {
    if m != nil {
        m.events = value
    }
}
// SetHexColor sets the hexColor property value. The calendar color, expressed in a hex color code of three hexadecimal values, each ranging from 00 to FF and representing the red, green, or blue components of the color in the RGB color space. If the user has never explicitly set a color for the calendar, this property is  empty.
func (m *Calendar) SetHexColor(value *string)() {
    if m != nil {
        m.hexColor = value
    }
}
// SetIsDefaultCalendar sets the isDefaultCalendar property value. true if this is the default calendar where new events are created by default, false otherwise.
func (m *Calendar) SetIsDefaultCalendar(value *bool)() {
    if m != nil {
        m.isDefaultCalendar = value
    }
}
// SetIsRemovable sets the isRemovable property value. Indicates whether this user calendar can be deleted from the user mailbox.
func (m *Calendar) SetIsRemovable(value *bool)() {
    if m != nil {
        m.isRemovable = value
    }
}
// SetIsTallyingResponses sets the isTallyingResponses property value. Indicates whether this user calendar supports tracking of meeting responses. Only meeting invites sent from users' primary calendars support tracking of meeting responses.
func (m *Calendar) SetIsTallyingResponses(value *bool)() {
    if m != nil {
        m.isTallyingResponses = value
    }
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the calendar. Read-only. Nullable.
func (m *Calendar) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    if m != nil {
        m.multiValueExtendedProperties = value
    }
}
// SetName sets the name property value. The calendar name.
func (m *Calendar) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetOwner sets the owner property value. If set, this represents the user who created or added the calendar. For a calendar that the user created or added, the owner property is set to the user. For a calendar shared with the user, the owner property is set to the person who shared that calendar with the user. Read-only.
func (m *Calendar) SetOwner(value EmailAddressable)() {
    if m != nil {
        m.owner = value
    }
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the calendar. Read-only. Nullable.
func (m *Calendar) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    if m != nil {
        m.singleValueExtendedProperties = value
    }
}
