package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RiskyUserHistoryItemable 
type RiskyUserHistoryItemable interface {
    RiskyUserable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActivity()(RiskUserActivityable)
    GetInitiatedBy()(*string)
    GetUserId()(*string)
    SetActivity(value RiskUserActivityable)()
    SetInitiatedBy(value *string)()
    SetUserId(value *string)()
}
