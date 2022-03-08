package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SharedInsightable 
type SharedInsightable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetLastShared()(SharingDetailable)
    GetLastSharedMethod()(Entityable)
    GetResource()(Entityable)
    GetResourceReference()(ResourceReferenceable)
    GetResourceVisualization()(ResourceVisualizationable)
    GetSharingHistory()([]SharingDetailable)
    SetLastShared(value SharingDetailable)()
    SetLastSharedMethod(value Entityable)()
    SetResource(value Entityable)()
    SetResourceReference(value ResourceReferenceable)()
    SetResourceVisualization(value ResourceVisualizationable)()
    SetSharingHistory(value []SharingDetailable)()
}
