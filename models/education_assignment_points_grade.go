package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationAssignmentPointsGrade 
type EducationAssignmentPointsGrade struct {
    EducationAssignmentGrade
}
// NewEducationAssignmentPointsGrade instantiates a new educationAssignmentPointsGrade and sets the default values.
func NewEducationAssignmentPointsGrade()(*EducationAssignmentPointsGrade) {
    m := &EducationAssignmentPointsGrade{
        EducationAssignmentGrade: *NewEducationAssignmentGrade(),
    }
    odataTypeValue := "#microsoft.graph.educationAssignmentPointsGrade"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationAssignmentPointsGradeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationAssignmentPointsGradeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationAssignmentPointsGrade(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationAssignmentPointsGrade) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationAssignmentGrade.GetFieldDeserializers()
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
    res["points"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPoints(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationAssignmentPointsGrade) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPoints gets the points property value. Number of points a teacher is giving this submission object.
func (m *EducationAssignmentPointsGrade) GetPoints()(*float32) {
    val, err := m.GetBackingStore().Get("points")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EducationAssignmentPointsGrade) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationAssignmentGrade.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat32Value("points", m.GetPoints())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationAssignmentPointsGrade) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPoints sets the points property value. Number of points a teacher is giving this submission object.
func (m *EducationAssignmentPointsGrade) SetPoints(value *float32)() {
    err := m.GetBackingStore().Set("points", value)
    if err != nil {
        panic(err)
    }
}
// EducationAssignmentPointsGradeable 
type EducationAssignmentPointsGradeable interface {
    EducationAssignmentGradeable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetPoints()(*float32)
    SetOdataType(value *string)()
    SetPoints(value *float32)()
}
