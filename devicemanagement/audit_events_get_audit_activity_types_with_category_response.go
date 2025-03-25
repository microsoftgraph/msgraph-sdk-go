package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AuditEventsGetAuditActivityTypesWithCategorygetResponseable instead.
type AuditEventsGetAuditActivityTypesWithCategoryResponse struct {
    AuditEventsGetAuditActivityTypesWithCategorygetResponse
}
// NewAuditEventsGetAuditActivityTypesWithCategoryResponse instantiates a new AuditEventsGetAuditActivityTypesWithCategoryResponse and sets the default values.
func NewAuditEventsGetAuditActivityTypesWithCategoryResponse()(*AuditEventsGetAuditActivityTypesWithCategoryResponse) {
    m := &AuditEventsGetAuditActivityTypesWithCategoryResponse{
        AuditEventsGetAuditActivityTypesWithCategorygetResponse: *NewAuditEventsGetAuditActivityTypesWithCategorygetResponse(),
    }
    return m
}
// CreateAuditEventsGetAuditActivityTypesWithCategoryResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuditEventsGetAuditActivityTypesWithCategoryResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditEventsGetAuditActivityTypesWithCategoryResponse(), nil
}
// Deprecated: This class is obsolete. Use AuditEventsGetAuditActivityTypesWithCategorygetResponseable instead.
type AuditEventsGetAuditActivityTypesWithCategoryResponseable interface {
    AuditEventsGetAuditActivityTypesWithCategorygetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
