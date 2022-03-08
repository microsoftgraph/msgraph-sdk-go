package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookOperationable 
type WorkbookOperationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetError()(WorkbookOperationErrorable)
    GetResourceLocation()(*string)
    GetStatus()(*WorkbookOperationStatus)
    SetError(value WorkbookOperationErrorable)()
    SetResourceLocation(value *string)()
    SetStatus(value *WorkbookOperationStatus)()
}
