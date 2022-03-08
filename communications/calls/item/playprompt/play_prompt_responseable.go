package playprompt

import (
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// PlayPromptResponseable 
type PlayPromptResponseable interface {
    Parsable
    AdditionalDataHolder
    GetPlayPromptOperation()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlayPromptOperationable)
    SetPlayPromptOperation(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlayPromptOperationable)()
}
