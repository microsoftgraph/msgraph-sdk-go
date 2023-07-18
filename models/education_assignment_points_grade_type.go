package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationAssignmentPointsGradeType 
type EducationAssignmentPointsGradeType struct {
    EducationAssignmentGradeType
}
// NewEducationAssignmentPointsGradeType instantiates a new educationAssignmentPointsGradeType and sets the default values.
func NewEducationAssignmentPointsGradeType()(*EducationAssignmentPointsGradeType) {
    m := &EducationAssignmentPointsGradeType{
        EducationAssignmentGradeType: *NewEducationAssignmentGradeType(),
    }
    odataTypeValue := "#microsoft.graph.educationAssignmentPointsGradeType"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationAssignmentPointsGradeTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationAssignmentPointsGradeTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationAssignmentPointsGradeType(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationAssignmentPointsGradeType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationAssignmentGradeType.GetFieldDeserializers()
    res["maxPoints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxPoints(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetMaxPoints gets the maxPoints property value. Max points possible for this assignment.
func (m *EducationAssignmentPointsGradeType) GetMaxPoints()(*float32) {
    val, err := m.GetBackingStore().Get("maxPoints")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationAssignmentPointsGradeType) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EducationAssignmentPointsGradeType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationAssignmentGradeType.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat32Value("maxPoints", m.GetMaxPoints())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMaxPoints sets the maxPoints property value. Max points possible for this assignment.
func (m *EducationAssignmentPointsGradeType) SetMaxPoints(value *float32)() {
    err := m.GetBackingStore().Set("maxPoints", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationAssignmentPointsGradeType) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// EducationAssignmentPointsGradeTypeable 
type EducationAssignmentPointsGradeTypeable interface {
    EducationAssignmentGradeTypeable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaxPoints()(*float32)
    GetOdataType()(*string)
    SetMaxPoints(value *float32)()
    SetOdataType(value *string)()
}
