// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable instead.
type ItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse struct {
    ItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse
}
// NewItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse instantiates a new ItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse and sets the default values.
func NewItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse()(*ItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse) {
    m := &ItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse{
        ItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse: *NewItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse(),
    }
    return m
}
// CreateItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable instead.
type ItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseable interface {
    ItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
