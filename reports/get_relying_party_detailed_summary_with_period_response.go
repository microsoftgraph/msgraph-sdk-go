package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GetRelyingPartyDetailedSummaryWithPeriodgetResponseable instead.
type GetRelyingPartyDetailedSummaryWithPeriodResponse struct {
    GetRelyingPartyDetailedSummaryWithPeriodgetResponse
}
// NewGetRelyingPartyDetailedSummaryWithPeriodResponse instantiates a new GetRelyingPartyDetailedSummaryWithPeriodResponse and sets the default values.
func NewGetRelyingPartyDetailedSummaryWithPeriodResponse()(*GetRelyingPartyDetailedSummaryWithPeriodResponse) {
    m := &GetRelyingPartyDetailedSummaryWithPeriodResponse{
        GetRelyingPartyDetailedSummaryWithPeriodgetResponse: *NewGetRelyingPartyDetailedSummaryWithPeriodgetResponse(),
    }
    return m
}
// CreateGetRelyingPartyDetailedSummaryWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetRelyingPartyDetailedSummaryWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetRelyingPartyDetailedSummaryWithPeriodResponse(), nil
}
// Deprecated: This class is obsolete. Use GetRelyingPartyDetailedSummaryWithPeriodgetResponseable instead.
type GetRelyingPartyDetailedSummaryWithPeriodResponseable interface {
    GetRelyingPartyDetailedSummaryWithPeriodgetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
