package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CommsOperationable 
type CommsOperationable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetClientContext()(*string)
    GetResultInfo()(ResultInfoable)
    GetStatus()(*OperationStatus)
    SetClientContext(value *string)()
    SetResultInfo(value ResultInfoable)()
    SetStatus(value *OperationStatus)()
}
