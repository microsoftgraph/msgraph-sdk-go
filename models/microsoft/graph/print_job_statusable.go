package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintJobStatusable 
type PrintJobStatusable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetDescription()(*string)
    GetDetails()([]PrintJobStateDetail)
    GetIsAcquiredByPrinter()(*bool)
    GetState()(*PrintJobProcessingState)
    SetDescription(value *string)()
    SetDetails(value []PrintJobStateDetail)()
    SetIsAcquiredByPrinter(value *bool)()
    SetState(value *PrintJobProcessingState)()
}
