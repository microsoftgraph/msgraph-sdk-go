package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type WorkbookChart struct {
    Entity
}
// WorkbookChart_WorkbookChart_height composed type wrapper for classes float64, ReferenceNumeric, string
type WorkbookChart_WorkbookChart_height struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWorkbookChart_WorkbookChart_height instantiates a new WorkbookChart_WorkbookChart_height and sets the default values.
func NewWorkbookChart_WorkbookChart_height()(*WorkbookChart_WorkbookChart_height) {
    m := &WorkbookChart_WorkbookChart_height{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateWorkbookChart_WorkbookChart_heightFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookChart_WorkbookChart_heightFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWorkbookChart_WorkbookChart_height()
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
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
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
func (m *WorkbookChart_WorkbookChart_height) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *WorkbookChart_WorkbookChart_height) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookChart_WorkbookChart_height) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *WorkbookChart_WorkbookChart_height) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *WorkbookChart_WorkbookChart_height) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *WorkbookChart_WorkbookChart_height) GetString()(*string) {
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
func (m *WorkbookChart_WorkbookChart_height) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
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
func (m *WorkbookChart_WorkbookChart_height) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *WorkbookChart_WorkbookChart_height) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *WorkbookChart_WorkbookChart_height) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *WorkbookChart_WorkbookChart_height) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// WorkbookChart_WorkbookChart_left composed type wrapper for classes float64, ReferenceNumeric, string
type WorkbookChart_WorkbookChart_left struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWorkbookChart_WorkbookChart_left instantiates a new WorkbookChart_WorkbookChart_left and sets the default values.
func NewWorkbookChart_WorkbookChart_left()(*WorkbookChart_WorkbookChart_left) {
    m := &WorkbookChart_WorkbookChart_left{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateWorkbookChart_WorkbookChart_leftFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookChart_WorkbookChart_leftFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWorkbookChart_WorkbookChart_left()
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
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
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
func (m *WorkbookChart_WorkbookChart_left) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *WorkbookChart_WorkbookChart_left) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookChart_WorkbookChart_left) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *WorkbookChart_WorkbookChart_left) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *WorkbookChart_WorkbookChart_left) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *WorkbookChart_WorkbookChart_left) GetString()(*string) {
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
func (m *WorkbookChart_WorkbookChart_left) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
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
func (m *WorkbookChart_WorkbookChart_left) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *WorkbookChart_WorkbookChart_left) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *WorkbookChart_WorkbookChart_left) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *WorkbookChart_WorkbookChart_left) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// WorkbookChart_WorkbookChart_top composed type wrapper for classes float64, ReferenceNumeric, string
type WorkbookChart_WorkbookChart_top struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWorkbookChart_WorkbookChart_top instantiates a new WorkbookChart_WorkbookChart_top and sets the default values.
func NewWorkbookChart_WorkbookChart_top()(*WorkbookChart_WorkbookChart_top) {
    m := &WorkbookChart_WorkbookChart_top{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateWorkbookChart_WorkbookChart_topFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookChart_WorkbookChart_topFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWorkbookChart_WorkbookChart_top()
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
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
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
func (m *WorkbookChart_WorkbookChart_top) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *WorkbookChart_WorkbookChart_top) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookChart_WorkbookChart_top) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *WorkbookChart_WorkbookChart_top) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *WorkbookChart_WorkbookChart_top) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *WorkbookChart_WorkbookChart_top) GetString()(*string) {
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
func (m *WorkbookChart_WorkbookChart_top) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
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
func (m *WorkbookChart_WorkbookChart_top) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *WorkbookChart_WorkbookChart_top) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *WorkbookChart_WorkbookChart_top) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *WorkbookChart_WorkbookChart_top) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// WorkbookChart_WorkbookChart_width composed type wrapper for classes float64, ReferenceNumeric, string
type WorkbookChart_WorkbookChart_width struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWorkbookChart_WorkbookChart_width instantiates a new WorkbookChart_WorkbookChart_width and sets the default values.
func NewWorkbookChart_WorkbookChart_width()(*WorkbookChart_WorkbookChart_width) {
    m := &WorkbookChart_WorkbookChart_width{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateWorkbookChart_WorkbookChart_widthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookChart_WorkbookChart_widthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWorkbookChart_WorkbookChart_width()
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
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
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
func (m *WorkbookChart_WorkbookChart_width) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *WorkbookChart_WorkbookChart_width) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookChart_WorkbookChart_width) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *WorkbookChart_WorkbookChart_width) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *WorkbookChart_WorkbookChart_width) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *WorkbookChart_WorkbookChart_width) GetString()(*string) {
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
func (m *WorkbookChart_WorkbookChart_width) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
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
func (m *WorkbookChart_WorkbookChart_width) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *WorkbookChart_WorkbookChart_width) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *WorkbookChart_WorkbookChart_width) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *WorkbookChart_WorkbookChart_width) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type WorkbookChart_WorkbookChart_heightable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
type WorkbookChart_WorkbookChart_leftable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
type WorkbookChart_WorkbookChart_topable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
type WorkbookChart_WorkbookChart_widthable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
// NewWorkbookChart instantiates a new WorkbookChart and sets the default values.
func NewWorkbookChart()(*WorkbookChart) {
    m := &WorkbookChart{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookChartFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChart(), nil
}
// GetAxes gets the axes property value. Represents chart axes. Read-only.
// returns a WorkbookChartAxesable when successful
func (m *WorkbookChart) GetAxes()(WorkbookChartAxesable) {
    val, err := m.GetBackingStore().Get("axes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookChartAxesable)
    }
    return nil
}
// GetDataLabels gets the dataLabels property value. Represents the data labels on the chart. Read-only.
// returns a WorkbookChartDataLabelsable when successful
func (m *WorkbookChart) GetDataLabels()(WorkbookChartDataLabelsable) {
    val, err := m.GetBackingStore().Get("dataLabels")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookChartDataLabelsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookChart) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["axes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartAxesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAxes(val.(WorkbookChartAxesable))
        }
        return nil
    }
    res["dataLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartDataLabelsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataLabels(val.(WorkbookChartDataLabelsable))
        }
        return nil
    }
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartAreaFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(WorkbookChartAreaFormatable))
        }
        return nil
    }
    res["height"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChart_WorkbookChart_heightFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeight(val.(*WorkbookChart_WorkbookChart_height))
        }
        return nil
    }
    res["left"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChart_WorkbookChart_leftFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLeft(val.(*WorkbookChart_WorkbookChart_left))
        }
        return nil
    }
    res["legend"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartLegendFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLegend(val.(WorkbookChartLegendable))
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["series"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookChartSeriesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookChartSeriesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WorkbookChartSeriesable)
                }
            }
            m.SetSeries(res)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartTitleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val.(WorkbookChartTitleable))
        }
        return nil
    }
    res["top"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChart_WorkbookChart_topFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTop(val.(*WorkbookChart_WorkbookChart_top))
        }
        return nil
    }
    res["width"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChart_WorkbookChart_widthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWidth(val.(*WorkbookChart_WorkbookChart_width))
        }
        return nil
    }
    res["worksheet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookWorksheetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorksheet(val.(WorkbookWorksheetable))
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. Encapsulates the format properties for the chart area. Read-only.
// returns a WorkbookChartAreaFormatable when successful
func (m *WorkbookChart) GetFormat()(WorkbookChartAreaFormatable) {
    val, err := m.GetBackingStore().Get("format")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookChartAreaFormatable)
    }
    return nil
}
// GetHeight gets the height property value. Represents the height, in points, of the chart object.
// returns a WorkbookChart_WorkbookChart_heightable when successful
func (m *WorkbookChart) GetHeight()(WorkbookChart_WorkbookChart_heightable) {
    val, err := m.GetBackingStore().Get("height")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookChart_WorkbookChart_heightable)
    }
    return nil
}
// GetLeft gets the left property value. The distance, in points, from the left side of the chart to the worksheet origin.
// returns a WorkbookChart_WorkbookChart_leftable when successful
func (m *WorkbookChart) GetLeft()(WorkbookChart_WorkbookChart_leftable) {
    val, err := m.GetBackingStore().Get("left")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookChart_WorkbookChart_leftable)
    }
    return nil
}
// GetLegend gets the legend property value. Represents the legend for the chart. Read-only.
// returns a WorkbookChartLegendable when successful
func (m *WorkbookChart) GetLegend()(WorkbookChartLegendable) {
    val, err := m.GetBackingStore().Get("legend")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookChartLegendable)
    }
    return nil
}
// GetName gets the name property value. Represents the name of a chart object.
// returns a *string when successful
func (m *WorkbookChart) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSeries gets the series property value. Represents either a single series or collection of series in the chart. Read-only.
// returns a []WorkbookChartSeriesable when successful
func (m *WorkbookChart) GetSeries()([]WorkbookChartSeriesable) {
    val, err := m.GetBackingStore().Get("series")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WorkbookChartSeriesable)
    }
    return nil
}
// GetTitle gets the title property value. Represents the title of the specified chart, including the text, visibility, position and formatting of the title. Read-only.
// returns a WorkbookChartTitleable when successful
func (m *WorkbookChart) GetTitle()(WorkbookChartTitleable) {
    val, err := m.GetBackingStore().Get("title")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookChartTitleable)
    }
    return nil
}
// GetTop gets the top property value. Represents the distance, in points, from the top edge of the object to the top of row 1 (on a worksheet) or the top of the chart area (on a chart).
// returns a WorkbookChart_WorkbookChart_topable when successful
func (m *WorkbookChart) GetTop()(WorkbookChart_WorkbookChart_topable) {
    val, err := m.GetBackingStore().Get("top")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookChart_WorkbookChart_topable)
    }
    return nil
}
// GetWidth gets the width property value. Represents the width, in points, of the chart object.
// returns a WorkbookChart_WorkbookChart_widthable when successful
func (m *WorkbookChart) GetWidth()(WorkbookChart_WorkbookChart_widthable) {
    val, err := m.GetBackingStore().Get("width")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookChart_WorkbookChart_widthable)
    }
    return nil
}
// GetWorksheet gets the worksheet property value. The worksheet containing the current chart. Read-only.
// returns a WorkbookWorksheetable when successful
func (m *WorkbookChart) GetWorksheet()(WorkbookWorksheetable) {
    val, err := m.GetBackingStore().Get("worksheet")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookWorksheetable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WorkbookChart) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("axes", m.GetAxes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("dataLabels", m.GetDataLabels())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("height", m.GetHeight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("left", m.GetLeft())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("legend", m.GetLegend())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetSeries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSeries()))
        for i, v := range m.GetSeries() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("series", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("top", m.GetTop())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("width", m.GetWidth())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("worksheet", m.GetWorksheet())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAxes sets the axes property value. Represents chart axes. Read-only.
func (m *WorkbookChart) SetAxes(value WorkbookChartAxesable)() {
    err := m.GetBackingStore().Set("axes", value)
    if err != nil {
        panic(err)
    }
}
// SetDataLabels sets the dataLabels property value. Represents the data labels on the chart. Read-only.
func (m *WorkbookChart) SetDataLabels(value WorkbookChartDataLabelsable)() {
    err := m.GetBackingStore().Set("dataLabels", value)
    if err != nil {
        panic(err)
    }
}
// SetFormat sets the format property value. Encapsulates the format properties for the chart area. Read-only.
func (m *WorkbookChart) SetFormat(value WorkbookChartAreaFormatable)() {
    err := m.GetBackingStore().Set("format", value)
    if err != nil {
        panic(err)
    }
}
// SetHeight sets the height property value. Represents the height, in points, of the chart object.
func (m *WorkbookChart) SetHeight(value WorkbookChart_WorkbookChart_heightable)() {
    err := m.GetBackingStore().Set("height", value)
    if err != nil {
        panic(err)
    }
}
// SetLeft sets the left property value. The distance, in points, from the left side of the chart to the worksheet origin.
func (m *WorkbookChart) SetLeft(value WorkbookChart_WorkbookChart_leftable)() {
    err := m.GetBackingStore().Set("left", value)
    if err != nil {
        panic(err)
    }
}
// SetLegend sets the legend property value. Represents the legend for the chart. Read-only.
func (m *WorkbookChart) SetLegend(value WorkbookChartLegendable)() {
    err := m.GetBackingStore().Set("legend", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. Represents the name of a chart object.
func (m *WorkbookChart) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetSeries sets the series property value. Represents either a single series or collection of series in the chart. Read-only.
func (m *WorkbookChart) SetSeries(value []WorkbookChartSeriesable)() {
    err := m.GetBackingStore().Set("series", value)
    if err != nil {
        panic(err)
    }
}
// SetTitle sets the title property value. Represents the title of the specified chart, including the text, visibility, position and formatting of the title. Read-only.
func (m *WorkbookChart) SetTitle(value WorkbookChartTitleable)() {
    err := m.GetBackingStore().Set("title", value)
    if err != nil {
        panic(err)
    }
}
// SetTop sets the top property value. Represents the distance, in points, from the top edge of the object to the top of row 1 (on a worksheet) or the top of the chart area (on a chart).
func (m *WorkbookChart) SetTop(value WorkbookChart_WorkbookChart_topable)() {
    err := m.GetBackingStore().Set("top", value)
    if err != nil {
        panic(err)
    }
}
// SetWidth sets the width property value. Represents the width, in points, of the chart object.
func (m *WorkbookChart) SetWidth(value WorkbookChart_WorkbookChart_widthable)() {
    err := m.GetBackingStore().Set("width", value)
    if err != nil {
        panic(err)
    }
}
// SetWorksheet sets the worksheet property value. The worksheet containing the current chart. Read-only.
func (m *WorkbookChart) SetWorksheet(value WorkbookWorksheetable)() {
    err := m.GetBackingStore().Set("worksheet", value)
    if err != nil {
        panic(err)
    }
}
type WorkbookChartable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAxes()(WorkbookChartAxesable)
    GetDataLabels()(WorkbookChartDataLabelsable)
    GetFormat()(WorkbookChartAreaFormatable)
    GetHeight()(WorkbookChart_WorkbookChart_heightable)
    GetLeft()(WorkbookChart_WorkbookChart_leftable)
    GetLegend()(WorkbookChartLegendable)
    GetName()(*string)
    GetSeries()([]WorkbookChartSeriesable)
    GetTitle()(WorkbookChartTitleable)
    GetTop()(WorkbookChart_WorkbookChart_topable)
    GetWidth()(WorkbookChart_WorkbookChart_widthable)
    GetWorksheet()(WorkbookWorksheetable)
    SetAxes(value WorkbookChartAxesable)()
    SetDataLabels(value WorkbookChartDataLabelsable)()
    SetFormat(value WorkbookChartAreaFormatable)()
    SetHeight(value WorkbookChart_WorkbookChart_heightable)()
    SetLeft(value WorkbookChart_WorkbookChart_leftable)()
    SetLegend(value WorkbookChartLegendable)()
    SetName(value *string)()
    SetSeries(value []WorkbookChartSeriesable)()
    SetTitle(value WorkbookChartTitleable)()
    SetTop(value WorkbookChart_WorkbookChart_topable)()
    SetWidth(value WorkbookChart_WorkbookChart_widthable)()
    SetWorksheet(value WorkbookWorksheetable)()
}
