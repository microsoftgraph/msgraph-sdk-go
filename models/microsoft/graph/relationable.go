package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/termstore"
)

// Relationable 
type Relationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetFromTerm()(Termable)
    GetRelationship()(*id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.RelationType)
    GetSet()(Setable)
    GetToTerm()(Termable)
    SetFromTerm(value Termable)()
    SetRelationship(value *id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.RelationType)()
    SetSet(value Setable)()
    SetToTerm(value Termable)()
}
