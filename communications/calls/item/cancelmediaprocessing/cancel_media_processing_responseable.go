package cancelmediaprocessing

import (
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// CancelMediaProcessingResponseable 
type CancelMediaProcessingResponseable interface {
    Parsable
    AdditionalDataHolder
    GetCancelMediaProcessingOperation()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CancelMediaProcessingOperationable)
    SetCancelMediaProcessingOperation(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CancelMediaProcessingOperationable)()
}
