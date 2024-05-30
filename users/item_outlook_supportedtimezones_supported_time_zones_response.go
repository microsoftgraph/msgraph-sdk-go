package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemOutlookSupportedtimezonesSupportedTimeZonesGetResponseable instead.
type ItemOutlookSupportedtimezonesSupportedTimeZonesResponse struct {
    ItemOutlookSupportedtimezonesSupportedTimeZonesGetResponse
}
// NewItemOutlookSupportedtimezonesSupportedTimeZonesResponse instantiates a new ItemOutlookSupportedtimezonesSupportedTimeZonesResponse and sets the default values.
func NewItemOutlookSupportedtimezonesSupportedTimeZonesResponse()(*ItemOutlookSupportedtimezonesSupportedTimeZonesResponse) {
    m := &ItemOutlookSupportedtimezonesSupportedTimeZonesResponse{
        ItemOutlookSupportedtimezonesSupportedTimeZonesGetResponse: *NewItemOutlookSupportedtimezonesSupportedTimeZonesGetResponse(),
    }
    return m
}
// CreateItemOutlookSupportedtimezonesSupportedTimeZonesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemOutlookSupportedtimezonesSupportedTimeZonesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOutlookSupportedtimezonesSupportedTimeZonesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemOutlookSupportedtimezonesSupportedTimeZonesGetResponseable instead.
type ItemOutlookSupportedtimezonesSupportedTimeZonesResponseable interface {
    ItemOutlookSupportedtimezonesSupportedTimeZonesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
