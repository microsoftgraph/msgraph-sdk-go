package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Calendarable 
type Calendarable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowedOnlineMeetingProviders()([]OnlineMeetingProviderType)
    GetCalendarPermissions()([]CalendarPermissionable)
    GetCalendarView()([]Eventable)
    GetCanEdit()(*bool)
    GetCanShare()(*bool)
    GetCanViewPrivateItems()(*bool)
    GetChangeKey()(*string)
    GetColor()(*CalendarColor)
    GetDefaultOnlineMeetingProvider()(*OnlineMeetingProviderType)
    GetEvents()([]Eventable)
    GetHexColor()(*string)
    GetIsDefaultCalendar()(*bool)
    GetIsRemovable()(*bool)
    GetIsTallyingResponses()(*bool)
    GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable)
    GetName()(*string)
    GetOwner()(EmailAddressable)
    GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable)
    SetAllowedOnlineMeetingProviders(value []OnlineMeetingProviderType)()
    SetCalendarPermissions(value []CalendarPermissionable)()
    SetCalendarView(value []Eventable)()
    SetCanEdit(value *bool)()
    SetCanShare(value *bool)()
    SetCanViewPrivateItems(value *bool)()
    SetChangeKey(value *string)()
    SetColor(value *CalendarColor)()
    SetDefaultOnlineMeetingProvider(value *OnlineMeetingProviderType)()
    SetEvents(value []Eventable)()
    SetHexColor(value *string)()
    SetIsDefaultCalendar(value *bool)()
    SetIsRemovable(value *bool)()
    SetIsTallyingResponses(value *bool)()
    SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)()
    SetName(value *string)()
    SetOwner(value EmailAddressable)()
    SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)()
}
