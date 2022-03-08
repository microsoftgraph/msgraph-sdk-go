package createorget

import (
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// CreateOrGetResponseable 
type CreateOrGetResponseable interface {
    Parsable
    AdditionalDataHolder
    GetOnlineMeeting()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnlineMeetingable)
    SetOnlineMeeting(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnlineMeetingable)()
}
