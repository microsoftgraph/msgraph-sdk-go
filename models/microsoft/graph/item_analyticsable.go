package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemAnalyticsable 
type ItemAnalyticsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllTime()(ItemActivityStatable)
    GetItemActivityStats()([]ItemActivityStatable)
    GetLastSevenDays()(ItemActivityStatable)
    SetAllTime(value ItemActivityStatable)()
    SetItemActivityStats(value []ItemActivityStatable)()
    SetLastSevenDays(value ItemActivityStatable)()
}
