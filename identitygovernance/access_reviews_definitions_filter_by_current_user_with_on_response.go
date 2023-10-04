package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewsDefinitionsFilterByCurrentUserWithOnResponse 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type AccessReviewsDefinitionsFilterByCurrentUserWithOnResponse struct {
    AccessReviewsDefinitionsFilterByCurrentUserWithOnGetResponse
}
// NewAccessReviewsDefinitionsFilterByCurrentUserWithOnResponse instantiates a new AccessReviewsDefinitionsFilterByCurrentUserWithOnResponse and sets the default values.
func NewAccessReviewsDefinitionsFilterByCurrentUserWithOnResponse()(*AccessReviewsDefinitionsFilterByCurrentUserWithOnResponse) {
    m := &AccessReviewsDefinitionsFilterByCurrentUserWithOnResponse{
        AccessReviewsDefinitionsFilterByCurrentUserWithOnGetResponse: *NewAccessReviewsDefinitionsFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateAccessReviewsDefinitionsFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewsDefinitionsFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewsDefinitionsFilterByCurrentUserWithOnResponse(), nil
}
// AccessReviewsDefinitionsFilterByCurrentUserWithOnResponseable 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type AccessReviewsDefinitionsFilterByCurrentUserWithOnResponseable interface {
    AccessReviewsDefinitionsFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
