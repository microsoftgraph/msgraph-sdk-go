package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AuditeventsGetauditcategoriesGetAuditCategoriesGetResponseable instead.
type AuditeventsGetauditcategoriesGetAuditCategoriesResponse struct {
    AuditeventsGetauditcategoriesGetAuditCategoriesGetResponse
}
// NewAuditeventsGetauditcategoriesGetAuditCategoriesResponse instantiates a new AuditeventsGetauditcategoriesGetAuditCategoriesResponse and sets the default values.
func NewAuditeventsGetauditcategoriesGetAuditCategoriesResponse()(*AuditeventsGetauditcategoriesGetAuditCategoriesResponse) {
    m := &AuditeventsGetauditcategoriesGetAuditCategoriesResponse{
        AuditeventsGetauditcategoriesGetAuditCategoriesGetResponse: *NewAuditeventsGetauditcategoriesGetAuditCategoriesGetResponse(),
    }
    return m
}
// CreateAuditeventsGetauditcategoriesGetAuditCategoriesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuditeventsGetauditcategoriesGetAuditCategoriesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditeventsGetauditcategoriesGetAuditCategoriesResponse(), nil
}
// Deprecated: This class is obsolete. Use AuditeventsGetauditcategoriesGetAuditCategoriesGetResponseable instead.
type AuditeventsGetauditcategoriesGetAuditCategoriesResponseable interface {
    AuditeventsGetauditcategoriesGetAuditCategoriesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
