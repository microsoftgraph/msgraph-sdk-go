package policies

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FeatureRolloutPoliciesItemAppliesToGetByIdsResponse 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type FeatureRolloutPoliciesItemAppliesToGetByIdsResponse struct {
    FeatureRolloutPoliciesItemAppliesToGetByIdsPostResponse
}
// NewFeatureRolloutPoliciesItemAppliesToGetByIdsResponse instantiates a new FeatureRolloutPoliciesItemAppliesToGetByIdsResponse and sets the default values.
func NewFeatureRolloutPoliciesItemAppliesToGetByIdsResponse()(*FeatureRolloutPoliciesItemAppliesToGetByIdsResponse) {
    m := &FeatureRolloutPoliciesItemAppliesToGetByIdsResponse{
        FeatureRolloutPoliciesItemAppliesToGetByIdsPostResponse: *NewFeatureRolloutPoliciesItemAppliesToGetByIdsPostResponse(),
    }
    return m
}
// CreateFeatureRolloutPoliciesItemAppliesToGetByIdsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFeatureRolloutPoliciesItemAppliesToGetByIdsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFeatureRolloutPoliciesItemAppliesToGetByIdsResponse(), nil
}
// FeatureRolloutPoliciesItemAppliesToGetByIdsResponseable 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type FeatureRolloutPoliciesItemAppliesToGetByIdsResponseable interface {
    FeatureRolloutPoliciesItemAppliesToGetByIdsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
