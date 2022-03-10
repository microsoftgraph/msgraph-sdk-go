package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BookingServiceable 
type BookingServiceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAdditionalInformation()(*string)
    GetCustomQuestions()([]BookingQuestionAssignmentable)
    GetDefaultDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetDefaultLocation()(Locationable)
    GetDefaultPrice()(*float64)
    GetDefaultPriceType()(*BookingPriceType)
    GetDefaultReminders()([]BookingReminderable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsHiddenFromCustomers()(*bool)
    GetIsLocationOnline()(*bool)
    GetMaximumAttendeesCount()(*int32)
    GetNotes()(*string)
    GetPostBuffer()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetPreBuffer()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetSchedulingPolicy()(BookingSchedulingPolicyable)
    GetSmsNotificationsEnabled()(*bool)
    GetStaffMemberIds()([]string)
    GetWebUrl()(*string)
    SetAdditionalInformation(value *string)()
    SetCustomQuestions(value []BookingQuestionAssignmentable)()
    SetDefaultDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetDefaultLocation(value Locationable)()
    SetDefaultPrice(value *float64)()
    SetDefaultPriceType(value *BookingPriceType)()
    SetDefaultReminders(value []BookingReminderable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsHiddenFromCustomers(value *bool)()
    SetIsLocationOnline(value *bool)()
    SetMaximumAttendeesCount(value *int32)()
    SetNotes(value *string)()
    SetPostBuffer(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetPreBuffer(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetSchedulingPolicy(value BookingSchedulingPolicyable)()
    SetSmsNotificationsEnabled(value *bool)()
    SetStaffMemberIds(value []string)()
    SetWebUrl(value *string)()
}
