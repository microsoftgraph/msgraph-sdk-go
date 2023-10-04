package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOutlookSupportedTimeZonesResponse 
// Deprecated: This class is obsolete. Use supportedTimeZonesGetResponse instead.
type ItemOutlookSupportedTimeZonesResponse struct {
    ItemOutlookSupportedTimeZonesGetResponse
}
// NewItemOutlookSupportedTimeZonesResponse instantiates a new ItemOutlookSupportedTimeZonesResponse and sets the default values.
func NewItemOutlookSupportedTimeZonesResponse()(*ItemOutlookSupportedTimeZonesResponse) {
    m := &ItemOutlookSupportedTimeZonesResponse{
        ItemOutlookSupportedTimeZonesGetResponse: *NewItemOutlookSupportedTimeZonesGetResponse(),
    }
    return m
}
// CreateItemOutlookSupportedTimeZonesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOutlookSupportedTimeZonesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOutlookSupportedTimeZonesResponse(), nil
}
// ItemOutlookSupportedTimeZonesResponseable 
// Deprecated: This class is obsolete. Use supportedTimeZonesGetResponse instead.
type ItemOutlookSupportedTimeZonesResponseable interface {
    ItemOutlookSupportedTimeZonesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
