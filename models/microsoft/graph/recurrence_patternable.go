package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RecurrencePatternable 
type RecurrencePatternable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetDayOfMonth()(*int32)
    GetDaysOfWeek()([]DayOfWeek)
    GetFirstDayOfWeek()(*DayOfWeek)
    GetIndex()(*WeekIndex)
    GetInterval()(*int32)
    GetMonth()(*int32)
    GetType()(*RecurrencePatternType)
    SetDayOfMonth(value *int32)()
    SetDaysOfWeek(value []DayOfWeek)()
    SetFirstDayOfWeek(value *DayOfWeek)()
    SetIndex(value *WeekIndex)()
    SetInterval(value *int32)()
    SetMonth(value *int32)()
    SetType(value *RecurrencePatternType)()
}
