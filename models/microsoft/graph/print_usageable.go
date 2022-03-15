package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintUsageable 
type PrintUsageable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCompletedBlackAndWhiteJobCount()(*int32)
    GetCompletedColorJobCount()(*int32)
    GetIncompleteJobCount()(*int32)
    GetUsageDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    SetCompletedBlackAndWhiteJobCount(value *int32)()
    SetCompletedColorJobCount(value *int32)()
    SetIncompleteJobCount(value *int32)()
    SetUsageDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
}
