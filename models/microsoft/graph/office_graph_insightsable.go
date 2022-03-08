package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OfficeGraphInsightsable 
type OfficeGraphInsightsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetShared()([]SharedInsightable)
    GetTrending()([]Trendingable)
    GetUsed()([]UsedInsightable)
    SetShared(value []SharedInsightable)()
    SetTrending(value []Trendingable)()
    SetUsed(value []UsedInsightable)()
}
