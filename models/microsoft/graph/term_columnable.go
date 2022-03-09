package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/termstore"
)

// TermColumnable 
type TermColumnable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowMultipleValues()(*bool)
    GetParentTerm()(id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.Termable)
    GetShowFullyQualifiedName()(*bool)
    GetTermSet()(id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.Setable)
    SetAllowMultipleValues(value *bool)()
    SetParentTerm(value id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.Termable)()
    SetShowFullyQualifiedName(value *bool)()
    SetTermSet(value id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.Setable)()
}
