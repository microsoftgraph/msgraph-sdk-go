package policies

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FeatureRolloutPoliciesItemAppliesToDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type FeatureRolloutPoliciesItemAppliesToDeltaResponse struct {
    FeatureRolloutPoliciesItemAppliesToDeltaGetResponse
}
// NewFeatureRolloutPoliciesItemAppliesToDeltaResponse instantiates a new FeatureRolloutPoliciesItemAppliesToDeltaResponse and sets the default values.
func NewFeatureRolloutPoliciesItemAppliesToDeltaResponse()(*FeatureRolloutPoliciesItemAppliesToDeltaResponse) {
    m := &FeatureRolloutPoliciesItemAppliesToDeltaResponse{
        FeatureRolloutPoliciesItemAppliesToDeltaGetResponse: *NewFeatureRolloutPoliciesItemAppliesToDeltaGetResponse(),
    }
    return m
}
// CreateFeatureRolloutPoliciesItemAppliesToDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFeatureRolloutPoliciesItemAppliesToDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFeatureRolloutPoliciesItemAppliesToDeltaResponse(), nil
}
// FeatureRolloutPoliciesItemAppliesToDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type FeatureRolloutPoliciesItemAppliesToDeltaResponseable interface {
    FeatureRolloutPoliciesItemAppliesToDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
