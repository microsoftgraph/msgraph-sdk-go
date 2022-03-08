package synclicenses

import (
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// SyncLicensesResponseable 
type SyncLicensesResponseable interface {
    Parsable
    AdditionalDataHolder
    GetVppToken()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.VppTokenable)
    SetVppToken(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.VppTokenable)()
}
