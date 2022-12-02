package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationMeAssignmentsDeltaResponse provides operations to call the delta method.
type EducationMeAssignmentsDeltaResponse struct {
    BaseDeltaFunctionResponse
    // The value property
    value []EducationAssignmentable
}
// NewEducationMeAssignmentsDeltaResponse instantiates a new EducationMeAssignmentsDeltaResponse and sets the default values.
func NewEducationMeAssignmentsDeltaResponse()(*EducationMeAssignmentsDeltaResponse) {
    m := &EducationMeAssignmentsDeltaResponse{
        BaseDeltaFunctionResponse: *NewBaseDeltaFunctionResponse(),
    }
    return m
}
// CreateEducationMeAssignmentsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationMeAssignmentsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationMeAssignmentsDeltaResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationMeAssignmentsDeltaResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseDeltaFunctionResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(EducationAssignmentable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *EducationMeAssignmentsDeltaResponse) GetValue()([]EducationAssignmentable) {
    return m.value
}
// Serialize serializes information the current object
func (m *EducationMeAssignmentsDeltaResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseDeltaFunctionResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *EducationMeAssignmentsDeltaResponse) SetValue(value []EducationAssignmentable)() {
    m.value = value
}
