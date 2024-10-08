package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type EducationAssignmentPointsGrade struct {
    EducationAssignmentGrade
}
// EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points composed type wrapper for classes float32, ReferenceNumeric, string
type EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewEducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points instantiates a new EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points and sets the default values.
func NewEducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points()(*EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points) {
    m := &EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateEducationAssignmentPointsGrade_EducationAssignmentPointsGrade_pointsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEducationAssignmentPointsGrade_EducationAssignmentPointsGrade_pointsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewEducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetFloat(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points) GetFloat()(*float32) {
    val, err := m.GetBackingStore().Get("float")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points) GetReferenceNumeric()(*ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetFloat() != nil {
        err := writer.WriteFloat32Value("", m.GetFloat())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_pointsable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
// NewEducationAssignmentPointsGrade instantiates a new EducationAssignmentPointsGrade and sets the default values.
func NewEducationAssignmentPointsGrade()(*EducationAssignmentPointsGrade) {
    m := &EducationAssignmentPointsGrade{
        EducationAssignmentGrade: *NewEducationAssignmentGrade(),
    }
    odataTypeValue := "#microsoft.graph.educationAssignmentPointsGrade"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationAssignmentPointsGradeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEducationAssignmentPointsGradeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationAssignmentPointsGrade(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EducationAssignmentPointsGrade) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationAssignmentGrade.GetFieldDeserializers()
    res["points"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationAssignmentPointsGrade_EducationAssignmentPointsGrade_pointsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPoints(val.(*EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_points))
        }
        return nil
    }
    return res
}
// GetPoints gets the points property value. Number of points a teacher is giving this submission object.
// returns a EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_pointsable when successful
func (m *EducationAssignmentPointsGrade) GetPoints()(EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_pointsable) {
    val, err := m.GetBackingStore().Get("points")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_pointsable)
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
        err = writer.WriteObjectValue("points", m.GetPoints())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPoints sets the points property value. Number of points a teacher is giving this submission object.
func (m *EducationAssignmentPointsGrade) SetPoints(value EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_pointsable)() {
    err := m.GetBackingStore().Set("points", value)
    if err != nil {
        panic(err)
    }
}
type EducationAssignmentPointsGradeable interface {
    EducationAssignmentGradeable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPoints()(EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_pointsable)
    SetPoints(value EducationAssignmentPointsGrade_EducationAssignmentPointsGrade_pointsable)()
}
