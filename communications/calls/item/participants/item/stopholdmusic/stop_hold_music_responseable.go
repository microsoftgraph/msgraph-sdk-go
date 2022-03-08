package stopholdmusic

import (
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// StopHoldMusicResponseable 
type StopHoldMusicResponseable interface {
    AdditionalDataHolder
    Parsable
    GetStopHoldMusicOperation()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.StopHoldMusicOperationable)
    SetStopHoldMusicOperation(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.StopHoldMusicOperationable)()
}
