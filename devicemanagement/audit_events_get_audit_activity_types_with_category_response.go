package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditEventsGetAuditActivityTypesWithCategoryResponse 
// Deprecated: This class is obsolete. Use getAuditActivityTypesWithCategoryGetResponse instead.
type AuditEventsGetAuditActivityTypesWithCategoryResponse struct {
    AuditEventsGetAuditActivityTypesWithCategoryGetResponse
}
// NewAuditEventsGetAuditActivityTypesWithCategoryResponse instantiates a new AuditEventsGetAuditActivityTypesWithCategoryResponse and sets the default values.
func NewAuditEventsGetAuditActivityTypesWithCategoryResponse()(*AuditEventsGetAuditActivityTypesWithCategoryResponse) {
    m := &AuditEventsGetAuditActivityTypesWithCategoryResponse{
        AuditEventsGetAuditActivityTypesWithCategoryGetResponse: *NewAuditEventsGetAuditActivityTypesWithCategoryGetResponse(),
    }
    return m
}
// CreateAuditEventsGetAuditActivityTypesWithCategoryResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditEventsGetAuditActivityTypesWithCategoryResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditEventsGetAuditActivityTypesWithCategoryResponse(), nil
}
// AuditEventsGetAuditActivityTypesWithCategoryResponseable 
// Deprecated: This class is obsolete. Use getAuditActivityTypesWithCategoryGetResponse instead.
type AuditEventsGetAuditActivityTypesWithCategoryResponseable interface {
    AuditEventsGetAuditActivityTypesWithCategoryGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
