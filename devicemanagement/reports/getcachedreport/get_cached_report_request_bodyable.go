package getcachedreport

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GetCachedReportRequestBodyable 
type GetCachedReportRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetGroupBy()([]string)
    GetId()(*string)
    GetOrderBy()([]string)
    GetSearch()(*string)
    GetSelect()([]string)
    GetSkip()(*int32)
    GetTop()(*int32)
    SetGroupBy(value []string)()
    SetId(value *string)()
    SetOrderBy(value []string)()
    SetSearch(value *string)()
    SetSelect(value []string)()
    SetSkip(value *int32)()
    SetTop(value *int32)()
}
