package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOutlookSupportedTimeZonesWithTimeZoneStandardResponse 
// Deprecated: This class is obsolete. Use supportedTimeZonesWithTimeZoneStandardGetResponse instead.
type ItemOutlookSupportedTimeZonesWithTimeZoneStandardResponse struct {
    ItemOutlookSupportedTimeZonesWithTimeZoneStandardGetResponse
}
// NewItemOutlookSupportedTimeZonesWithTimeZoneStandardResponse instantiates a new ItemOutlookSupportedTimeZonesWithTimeZoneStandardResponse and sets the default values.
func NewItemOutlookSupportedTimeZonesWithTimeZoneStandardResponse()(*ItemOutlookSupportedTimeZonesWithTimeZoneStandardResponse) {
    m := &ItemOutlookSupportedTimeZonesWithTimeZoneStandardResponse{
        ItemOutlookSupportedTimeZonesWithTimeZoneStandardGetResponse: *NewItemOutlookSupportedTimeZonesWithTimeZoneStandardGetResponse(),
    }
    return m
}
// CreateItemOutlookSupportedTimeZonesWithTimeZoneStandardResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOutlookSupportedTimeZonesWithTimeZoneStandardResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOutlookSupportedTimeZonesWithTimeZoneStandardResponse(), nil
}
// ItemOutlookSupportedTimeZonesWithTimeZoneStandardResponseable 
// Deprecated: This class is obsolete. Use supportedTimeZonesWithTimeZoneStandardGetResponse instead.
type ItemOutlookSupportedTimeZonesWithTimeZoneStandardResponseable interface {
    ItemOutlookSupportedTimeZonesWithTimeZoneStandardGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
