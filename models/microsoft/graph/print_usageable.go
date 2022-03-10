package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintUsageable 
type PrintUsageable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCompletedBlackAndWhiteJobCount()(*int64)
    GetCompletedColorJobCount()(*int64)
    GetIncompleteJobCount()(*int64)
    GetUsageDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    SetCompletedBlackAndWhiteJobCount(value *int64)()
    SetCompletedColorJobCount(value *int64)()
    SetIncompleteJobCount(value *int64)()
    SetUsageDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
}
