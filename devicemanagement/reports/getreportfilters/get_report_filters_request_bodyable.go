package getreportfilters

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GetReportFiltersRequestBodyable 
type GetReportFiltersRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetFilter()(*string)
    GetGroupBy()([]string)
    GetName()(*string)
    GetOrderBy()([]string)
    GetSearch()(*string)
    GetSelect()([]string)
    GetSessionId()(*string)
    GetSkip()(*int32)
    GetTop()(*int32)
    SetFilter(value *string)()
    SetGroupBy(value []string)()
    SetName(value *string)()
    SetOrderBy(value []string)()
    SetSearch(value *string)()
    SetSelect(value []string)()
    SetSessionId(value *string)()
    SetSkip(value *int32)()
    SetTop(value *int32)()
}
