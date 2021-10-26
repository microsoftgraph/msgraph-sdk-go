package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Calendar struct {
    Entity
    // Represent the online meeting service providers that can be used to create online meetings in this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
    allowedOnlineMeetingProviders []OnlineMeetingProviderType;
    // The permissions of the users with whom the calendar is shared.
    calendarPermissions []CalendarPermission;
    // The calendar view for the calendar. Navigation property. Read-only.
    calendarView []Event;
    // true if the user can write to the calendar, false otherwise. This property is true for the user who created the calendar. This property is also true for a user who has been shared a calendar and granted write access.
    canEdit *bool;
    // true if the user has the permission to share the calendar, false otherwise. Only the user who created the calendar can share it.
    canShare *bool;
    // true if the user can read calendar items that have been marked private, false otherwise.
    canViewPrivateItems *bool;
    // Identifies the version of the calendar object. Every time the calendar is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
    changeKey *string;
    // Specifies the color theme to distinguish the calendar from other calendars in a UI. The property values are: auto, lightBlue, lightGreen, lightOrange, lightGray, lightYellow, lightTeal, lightPink, lightBrown, lightRed, maxColor.
    color *CalendarColor;
    // The default online meeting provider for meetings sent from this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
    defaultOnlineMeetingProvider *OnlineMeetingProviderType;
    // The events in the calendar. Navigation property. Read-only.
    events []Event;
    // The calendar color, expressed in a hex color code of three hexadecimal values, each ranging from 00 to FF and representing the red, green, or blue components of the color in the RGB color space. If the user has never explicitly set a color for the calendar, this property is empty. Read-only.
    hexColor *string;
    // true if this is the default calendar where new events are created by default, false otherwise.
    isDefaultCalendar *bool;
    // Indicates whether this user calendar can be deleted from the user mailbox.
    isRemovable *bool;
    // Indicates whether this user calendar supports tracking of meeting responses. Only meeting invites sent from users' primary calendars support tracking of meeting responses.
    isTallyingResponses *bool;
    // The collection of multi-value extended properties defined for the calendar. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedProperty;
    // The calendar name.
    name *string;
    // If set, this represents the user who created or added the calendar. For a calendar that the user created or added, the owner property is set to the user. For a calendar shared with the user, the owner property is set to the person who shared that calendar with the user.
    owner *EmailAddress;
    // The collection of single-value extended properties defined for the calendar. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedProperty;
}
// Instantiates a new calendar and sets the default values.
func NewCalendar()(*Calendar) {
    m := &Calendar{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the allowedOnlineMeetingProviders property value. Represent the online meeting service providers that can be used to create online meetings in this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
func (m *Calendar) GetAllowedOnlineMeetingProviders()([]OnlineMeetingProviderType) {
    if m == nil {
        return nil
    } else {
        return m.allowedOnlineMeetingProviders
    }
}
// Gets the calendarPermissions property value. The permissions of the users with whom the calendar is shared.
func (m *Calendar) GetCalendarPermissions()([]CalendarPermission) {
    if m == nil {
        return nil
    } else {
        return m.calendarPermissions
    }
}
// Gets the calendarView property value. The calendar view for the calendar. Navigation property. Read-only.
func (m *Calendar) GetCalendarView()([]Event) {
    if m == nil {
        return nil
    } else {
        return m.calendarView
    }
}
// Gets the canEdit property value. true if the user can write to the calendar, false otherwise. This property is true for the user who created the calendar. This property is also true for a user who has been shared a calendar and granted write access.
func (m *Calendar) GetCanEdit()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.canEdit
    }
}
// Gets the canShare property value. true if the user has the permission to share the calendar, false otherwise. Only the user who created the calendar can share it.
func (m *Calendar) GetCanShare()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.canShare
    }
}
// Gets the canViewPrivateItems property value. true if the user can read calendar items that have been marked private, false otherwise.
func (m *Calendar) GetCanViewPrivateItems()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.canViewPrivateItems
    }
}
// Gets the changeKey property value. Identifies the version of the calendar object. Every time the calendar is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
func (m *Calendar) GetChangeKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.changeKey
    }
}
// Gets the color property value. Specifies the color theme to distinguish the calendar from other calendars in a UI. The property values are: auto, lightBlue, lightGreen, lightOrange, lightGray, lightYellow, lightTeal, lightPink, lightBrown, lightRed, maxColor.
func (m *Calendar) GetColor()(*CalendarColor) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
// Gets the defaultOnlineMeetingProvider property value. The default online meeting provider for meetings sent from this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
func (m *Calendar) GetDefaultOnlineMeetingProvider()(*OnlineMeetingProviderType) {
    if m == nil {
        return nil
    } else {
        return m.defaultOnlineMeetingProvider
    }
}
// Gets the events property value. The events in the calendar. Navigation property. Read-only.
func (m *Calendar) GetEvents()([]Event) {
    if m == nil {
        return nil
    } else {
        return m.events
    }
}
// Gets the hexColor property value. The calendar color, expressed in a hex color code of three hexadecimal values, each ranging from 00 to FF and representing the red, green, or blue components of the color in the RGB color space. If the user has never explicitly set a color for the calendar, this property is empty. Read-only.
func (m *Calendar) GetHexColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hexColor
    }
}
// Gets the isDefaultCalendar property value. true if this is the default calendar where new events are created by default, false otherwise.
func (m *Calendar) GetIsDefaultCalendar()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefaultCalendar
    }
}
// Gets the isRemovable property value. Indicates whether this user calendar can be deleted from the user mailbox.
func (m *Calendar) GetIsRemovable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRemovable
    }
}
// Gets the isTallyingResponses property value. Indicates whether this user calendar supports tracking of meeting responses. Only meeting invites sent from users' primary calendars support tracking of meeting responses.
func (m *Calendar) GetIsTallyingResponses()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isTallyingResponses
    }
}
// Gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the calendar. Read-only. Nullable.
func (m *Calendar) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// Gets the name property value. The calendar name.
func (m *Calendar) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the owner property value. If set, this represents the user who created or added the calendar. For a calendar that the user created or added, the owner property is set to the user. For a calendar shared with the user, the owner property is set to the person who shared that calendar with the user.
func (m *Calendar) GetOwner()(*EmailAddress) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
// Gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the calendar. Read-only. Nullable.
func (m *Calendar) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
}
// The deserialization information for the current model
func (m *Calendar) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedOnlineMeetingProviders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseOnlineMeetingProviderType)
        if err != nil {
            return err
        }
        res := make([]OnlineMeetingProviderType, len(val))
        for i, v := range val {
            res[i] = *(v.(*OnlineMeetingProviderType))
        }
        m.SetAllowedOnlineMeetingProviders(res)
        return nil
    }
    res["calendarPermissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCalendarPermission() })
        if err != nil {
            return err
        }
        res := make([]CalendarPermission, len(val))
        for i, v := range val {
            res[i] = *(v.(*CalendarPermission))
        }
        m.SetCalendarPermissions(res)
        return nil
    }
    res["calendarView"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEvent() })
        if err != nil {
            return err
        }
        res := make([]Event, len(val))
        for i, v := range val {
            res[i] = *(v.(*Event))
        }
        m.SetCalendarView(res)
        return nil
    }
    res["canEdit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetCanEdit(val)
        return nil
    }
    res["canShare"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetCanShare(val)
        return nil
    }
    res["canViewPrivateItems"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetCanViewPrivateItems(val)
        return nil
    }
    res["changeKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetChangeKey(val)
        return nil
    }
    res["color"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCalendarColor)
        if err != nil {
            return err
        }
        cast := val.(CalendarColor)
        m.SetColor(&cast)
        return nil
    }
    res["defaultOnlineMeetingProvider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnlineMeetingProviderType)
        if err != nil {
            return err
        }
        cast := val.(OnlineMeetingProviderType)
        m.SetDefaultOnlineMeetingProvider(&cast)
        return nil
    }
    res["events"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEvent() })
        if err != nil {
            return err
        }
        res := make([]Event, len(val))
        for i, v := range val {
            res[i] = *(v.(*Event))
        }
        m.SetEvents(res)
        return nil
    }
    res["hexColor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHexColor(val)
        return nil
    }
    res["isDefaultCalendar"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefaultCalendar(val)
        return nil
    }
    res["isRemovable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRemovable(val)
        return nil
    }
    res["isTallyingResponses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsTallyingResponses(val)
        return nil
    }
    res["multiValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMultiValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        res := make([]MultiValueLegacyExtendedProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*MultiValueLegacyExtendedProperty))
        }
        m.SetMultiValueExtendedProperties(res)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["owner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmailAddress() })
        if err != nil {
            return err
        }
        m.SetOwner(val.(*EmailAddress))
        return nil
    }
    res["singleValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSingleValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        res := make([]SingleValueLegacyExtendedProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*SingleValueLegacyExtendedProperty))
        }
        m.SetSingleValueExtendedProperties(res)
        return nil
    }
    return res
}
func (m *Calendar) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Calendar) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("allowedOnlineMeetingProviders", SerializeOnlineMeetingProviderType(m.GetAllowedOnlineMeetingProviders()))
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCalendarPermissions()))
        for i, v := range m.GetCalendarPermissions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("calendarPermissions", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCalendarView()))
        for i, v := range m.GetCalendarView() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := m.GetColor().String()
        err = writer.WriteStringValue("color", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefaultOnlineMeetingProvider() != nil {
        cast := m.GetDefaultOnlineMeetingProvider().String()
        err = writer.WriteStringValue("defaultOnlineMeetingProvider", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEvents()))
        for i, v := range m.GetEvents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    return nil
}
// Sets the allowedOnlineMeetingProviders property value. Represent the online meeting service providers that can be used to create online meetings in this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
// Parameters:
//  - value : Value to set for the allowedOnlineMeetingProviders property.
func (m *Calendar) SetAllowedOnlineMeetingProviders(value []OnlineMeetingProviderType)() {
    m.allowedOnlineMeetingProviders = value
}
// Sets the calendarPermissions property value. The permissions of the users with whom the calendar is shared.
// Parameters:
//  - value : Value to set for the calendarPermissions property.
func (m *Calendar) SetCalendarPermissions(value []CalendarPermission)() {
    m.calendarPermissions = value
}
// Sets the calendarView property value. The calendar view for the calendar. Navigation property. Read-only.
// Parameters:
//  - value : Value to set for the calendarView property.
func (m *Calendar) SetCalendarView(value []Event)() {
    m.calendarView = value
}
// Sets the canEdit property value. true if the user can write to the calendar, false otherwise. This property is true for the user who created the calendar. This property is also true for a user who has been shared a calendar and granted write access.
// Parameters:
//  - value : Value to set for the canEdit property.
func (m *Calendar) SetCanEdit(value *bool)() {
    m.canEdit = value
}
// Sets the canShare property value. true if the user has the permission to share the calendar, false otherwise. Only the user who created the calendar can share it.
// Parameters:
//  - value : Value to set for the canShare property.
func (m *Calendar) SetCanShare(value *bool)() {
    m.canShare = value
}
// Sets the canViewPrivateItems property value. true if the user can read calendar items that have been marked private, false otherwise.
// Parameters:
//  - value : Value to set for the canViewPrivateItems property.
func (m *Calendar) SetCanViewPrivateItems(value *bool)() {
    m.canViewPrivateItems = value
}
// Sets the changeKey property value. Identifies the version of the calendar object. Every time the calendar is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
// Parameters:
//  - value : Value to set for the changeKey property.
func (m *Calendar) SetChangeKey(value *string)() {
    m.changeKey = value
}
// Sets the color property value. Specifies the color theme to distinguish the calendar from other calendars in a UI. The property values are: auto, lightBlue, lightGreen, lightOrange, lightGray, lightYellow, lightTeal, lightPink, lightBrown, lightRed, maxColor.
// Parameters:
//  - value : Value to set for the color property.
func (m *Calendar) SetColor(value *CalendarColor)() {
    m.color = value
}
// Sets the defaultOnlineMeetingProvider property value. The default online meeting provider for meetings sent from this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
// Parameters:
//  - value : Value to set for the defaultOnlineMeetingProvider property.
func (m *Calendar) SetDefaultOnlineMeetingProvider(value *OnlineMeetingProviderType)() {
    m.defaultOnlineMeetingProvider = value
}
// Sets the events property value. The events in the calendar. Navigation property. Read-only.
// Parameters:
//  - value : Value to set for the events property.
func (m *Calendar) SetEvents(value []Event)() {
    m.events = value
}
// Sets the hexColor property value. The calendar color, expressed in a hex color code of three hexadecimal values, each ranging from 00 to FF and representing the red, green, or blue components of the color in the RGB color space. If the user has never explicitly set a color for the calendar, this property is empty. Read-only.
// Parameters:
//  - value : Value to set for the hexColor property.
func (m *Calendar) SetHexColor(value *string)() {
    m.hexColor = value
}
// Sets the isDefaultCalendar property value. true if this is the default calendar where new events are created by default, false otherwise.
// Parameters:
//  - value : Value to set for the isDefaultCalendar property.
func (m *Calendar) SetIsDefaultCalendar(value *bool)() {
    m.isDefaultCalendar = value
}
// Sets the isRemovable property value. Indicates whether this user calendar can be deleted from the user mailbox.
// Parameters:
//  - value : Value to set for the isRemovable property.
func (m *Calendar) SetIsRemovable(value *bool)() {
    m.isRemovable = value
}
// Sets the isTallyingResponses property value. Indicates whether this user calendar supports tracking of meeting responses. Only meeting invites sent from users' primary calendars support tracking of meeting responses.
// Parameters:
//  - value : Value to set for the isTallyingResponses property.
func (m *Calendar) SetIsTallyingResponses(value *bool)() {
    m.isTallyingResponses = value
}
// Sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the calendar. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the multiValueExtendedProperties property.
func (m *Calendar) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedProperty)() {
    m.multiValueExtendedProperties = value
}
// Sets the name property value. The calendar name.
// Parameters:
//  - value : Value to set for the name property.
func (m *Calendar) SetName(value *string)() {
    m.name = value
}
// Sets the owner property value. If set, this represents the user who created or added the calendar. For a calendar that the user created or added, the owner property is set to the user. For a calendar shared with the user, the owner property is set to the person who shared that calendar with the user.
// Parameters:
//  - value : Value to set for the owner property.
func (m *Calendar) SetOwner(value *EmailAddress)() {
    m.owner = value
}
// Sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the calendar. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the singleValueExtendedProperties property.
func (m *Calendar) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedProperty)() {
    m.singleValueExtendedProperties = value
}
