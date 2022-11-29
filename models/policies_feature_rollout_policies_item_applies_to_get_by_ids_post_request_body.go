package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBody provides operations to call the getByIds method.
type PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The ids property
    ids []string
    // The types property
    types []string
}
// NewPoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBody instantiates a new PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBody and sets the default values.
func NewPoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBody()(*PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBody) {
    m := &PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ids"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetIds)
    res["types"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetTypes)
    return res
}
// GetIds gets the ids property value. The ids property
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBody) GetIds()([]string) {
    return m.ids
}
// GetTypes gets the types property value. The types property
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBody) GetTypes()([]string) {
    return m.types
}
// Serialize serializes information the current object
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIds() != nil {
        err := writer.WriteCollectionOfStringValues("ids", m.GetIds())
        if err != nil {
            return err
        }
    }
    if m.GetTypes() != nil {
        err := writer.WriteCollectionOfStringValues("types", m.GetTypes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIds sets the ids property value. The ids property
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBody) SetIds(value []string)() {
    m.ids = value
}
// SetTypes sets the types property value. The types property
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBody) SetTypes(value []string)() {
    m.types = value
}
