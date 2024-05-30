package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GetpresencesbyuseridGetPresencesByUserIdPostResponseable instead.
type GetpresencesbyuseridGetPresencesByUserIdResponse struct {
    GetpresencesbyuseridGetPresencesByUserIdPostResponse
}
// NewGetpresencesbyuseridGetPresencesByUserIdResponse instantiates a new GetpresencesbyuseridGetPresencesByUserIdResponse and sets the default values.
func NewGetpresencesbyuseridGetPresencesByUserIdResponse()(*GetpresencesbyuseridGetPresencesByUserIdResponse) {
    m := &GetpresencesbyuseridGetPresencesByUserIdResponse{
        GetpresencesbyuseridGetPresencesByUserIdPostResponse: *NewGetpresencesbyuseridGetPresencesByUserIdPostResponse(),
    }
    return m
}
// CreateGetpresencesbyuseridGetPresencesByUserIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetpresencesbyuseridGetPresencesByUserIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetpresencesbyuseridGetPresencesByUserIdResponse(), nil
}
// Deprecated: This class is obsolete. Use GetpresencesbyuseridGetPresencesByUserIdPostResponseable instead.
type GetpresencesbyuseridGetPresencesByUserIdResponseable interface {
    GetpresencesbyuseridGetPresencesByUserIdPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
