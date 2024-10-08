package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type EducationAssignmentPointsGradeType struct {
    EducationAssignmentGradeType
}
// EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints composed type wrapper for classes float32, ReferenceNumeric, string
type EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewEducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints instantiates a new EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints and sets the default values.
func NewEducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints()(*EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints) {
    m := &EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateEducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPointsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPointsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewEducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints()
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
func (m *EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints) GetFloat()(*float32) {
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
func (m *EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints) GetString()(*string) {
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
func (m *EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPointsable interface {
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
// NewEducationAssignmentPointsGradeType instantiates a new EducationAssignmentPointsGradeType and sets the default values.
func NewEducationAssignmentPointsGradeType()(*EducationAssignmentPointsGradeType) {
    m := &EducationAssignmentPointsGradeType{
        EducationAssignmentGradeType: *NewEducationAssignmentGradeType(),
    }
    odataTypeValue := "#microsoft.graph.educationAssignmentPointsGradeType"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationAssignmentPointsGradeTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEducationAssignmentPointsGradeTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationAssignmentPointsGradeType(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EducationAssignmentPointsGradeType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationAssignmentGradeType.GetFieldDeserializers()
    res["maxPoints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPointsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxPoints(val.(*EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPoints))
        }
        return nil
    }
    return res
}
// GetMaxPoints gets the maxPoints property value. Max points possible for this assignment.
// returns a EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPointsable when successful
func (m *EducationAssignmentPointsGradeType) GetMaxPoints()(EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPointsable) {
    val, err := m.GetBackingStore().Get("maxPoints")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPointsable)
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
        err = writer.WriteObjectValue("maxPoints", m.GetMaxPoints())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMaxPoints sets the maxPoints property value. Max points possible for this assignment.
func (m *EducationAssignmentPointsGradeType) SetMaxPoints(value EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPointsable)() {
    err := m.GetBackingStore().Set("maxPoints", value)
    if err != nil {
        panic(err)
    }
}
type EducationAssignmentPointsGradeTypeable interface {
    EducationAssignmentGradeTypeable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaxPoints()(EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPointsable)
    SetMaxPoints(value EducationAssignmentPointsGradeType_EducationAssignmentPointsGradeType_maxPointsable)()
}
