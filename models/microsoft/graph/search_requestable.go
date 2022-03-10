package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SearchRequestable 
type SearchRequestable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAggregationFilters()([]string)
    GetAggregations()([]AggregationOptionable)
    GetContentSources()([]string)
    GetEnableTopResults()(*bool)
    GetEntityTypes()([]EntityType)
    GetFields()([]string)
    GetFrom()(*int32)
    GetQuery()(SearchQueryable)
    GetQueryAlterationOptions()(SearchAlterationOptionsable)
    GetResultTemplateOptions()(ResultTemplateOptionable)
    GetSize()(*int32)
    GetSortProperties()([]SortPropertyable)
    SetAggregationFilters(value []string)()
    SetAggregations(value []AggregationOptionable)()
    SetContentSources(value []string)()
    SetEnableTopResults(value *bool)()
    SetEntityTypes(value []EntityType)()
    SetFields(value []string)()
    SetFrom(value *int32)()
    SetQuery(value SearchQueryable)()
    SetQueryAlterationOptions(value SearchAlterationOptionsable)()
    SetResultTemplateOptions(value ResultTemplateOptionable)()
    SetSize(value *int32)()
    SetSortProperties(value []SortPropertyable)()
}
