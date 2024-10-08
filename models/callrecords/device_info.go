package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type DeviceInfo struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// DeviceInfo_DeviceInfo_captureNotFunctioningEventRatio composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type DeviceInfo_DeviceInfo_captureNotFunctioningEventRatio struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceInfo_DeviceInfo_captureNotFunctioningEventRatio instantiates a new DeviceInfo_DeviceInfo_captureNotFunctioningEventRatio and sets the default values.
func NewDeviceInfo_DeviceInfo_captureNotFunctioningEventRatio()(*DeviceInfo_DeviceInfo_captureNotFunctioningEventRatio) {
    m := &DeviceInfo_DeviceInfo_captureNotFunctioningEventRatio{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateDeviceInfo_DeviceInfo_captureNotFunctioningEventRatioFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceInfo_DeviceInfo_captureNotFunctioningEventRatioFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeviceInfo_DeviceInfo_captureNotFunctioningEventRatio()
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
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
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
func (m *DeviceInfo_DeviceInfo_captureNotFunctioningEventRatio) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceInfo_DeviceInfo_captureNotFunctioningEventRatio) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *DeviceInfo_DeviceInfo_captureNotFunctioningEventRatio) GetFloat()(*float32) {
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
func (m *DeviceInfo_DeviceInfo_captureNotFunctioningEventRatio) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *DeviceInfo_DeviceInfo_captureNotFunctioningEventRatio) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *DeviceInfo_DeviceInfo_captureNotFunctioningEventRatio) GetString()(*string) {
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
func (m *DeviceInfo_DeviceInfo_captureNotFunctioningEventRatio) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceInfo_DeviceInfo_captureNotFunctioningEventRatio) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *DeviceInfo_DeviceInfo_captureNotFunctioningEventRatio) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *DeviceInfo_DeviceInfo_captureNotFunctioningEventRatio) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *DeviceInfo_DeviceInfo_captureNotFunctioningEventRatio) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// DeviceInfo_DeviceInfo_cpuInsufficentEventRatio composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type DeviceInfo_DeviceInfo_cpuInsufficentEventRatio struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceInfo_DeviceInfo_cpuInsufficentEventRatio instantiates a new DeviceInfo_DeviceInfo_cpuInsufficentEventRatio and sets the default values.
func NewDeviceInfo_DeviceInfo_cpuInsufficentEventRatio()(*DeviceInfo_DeviceInfo_cpuInsufficentEventRatio) {
    m := &DeviceInfo_DeviceInfo_cpuInsufficentEventRatio{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateDeviceInfo_DeviceInfo_cpuInsufficentEventRatioFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceInfo_DeviceInfo_cpuInsufficentEventRatioFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeviceInfo_DeviceInfo_cpuInsufficentEventRatio()
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
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
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
func (m *DeviceInfo_DeviceInfo_cpuInsufficentEventRatio) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceInfo_DeviceInfo_cpuInsufficentEventRatio) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *DeviceInfo_DeviceInfo_cpuInsufficentEventRatio) GetFloat()(*float32) {
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
func (m *DeviceInfo_DeviceInfo_cpuInsufficentEventRatio) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *DeviceInfo_DeviceInfo_cpuInsufficentEventRatio) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *DeviceInfo_DeviceInfo_cpuInsufficentEventRatio) GetString()(*string) {
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
func (m *DeviceInfo_DeviceInfo_cpuInsufficentEventRatio) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceInfo_DeviceInfo_cpuInsufficentEventRatio) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *DeviceInfo_DeviceInfo_cpuInsufficentEventRatio) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *DeviceInfo_DeviceInfo_cpuInsufficentEventRatio) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *DeviceInfo_DeviceInfo_cpuInsufficentEventRatio) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// DeviceInfo_DeviceInfo_deviceClippingEventRatio composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type DeviceInfo_DeviceInfo_deviceClippingEventRatio struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceInfo_DeviceInfo_deviceClippingEventRatio instantiates a new DeviceInfo_DeviceInfo_deviceClippingEventRatio and sets the default values.
func NewDeviceInfo_DeviceInfo_deviceClippingEventRatio()(*DeviceInfo_DeviceInfo_deviceClippingEventRatio) {
    m := &DeviceInfo_DeviceInfo_deviceClippingEventRatio{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateDeviceInfo_DeviceInfo_deviceClippingEventRatioFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceInfo_DeviceInfo_deviceClippingEventRatioFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeviceInfo_DeviceInfo_deviceClippingEventRatio()
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
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
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
func (m *DeviceInfo_DeviceInfo_deviceClippingEventRatio) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceInfo_DeviceInfo_deviceClippingEventRatio) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *DeviceInfo_DeviceInfo_deviceClippingEventRatio) GetFloat()(*float32) {
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
func (m *DeviceInfo_DeviceInfo_deviceClippingEventRatio) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *DeviceInfo_DeviceInfo_deviceClippingEventRatio) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *DeviceInfo_DeviceInfo_deviceClippingEventRatio) GetString()(*string) {
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
func (m *DeviceInfo_DeviceInfo_deviceClippingEventRatio) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceInfo_DeviceInfo_deviceClippingEventRatio) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *DeviceInfo_DeviceInfo_deviceClippingEventRatio) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *DeviceInfo_DeviceInfo_deviceClippingEventRatio) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *DeviceInfo_DeviceInfo_deviceClippingEventRatio) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// DeviceInfo_DeviceInfo_deviceGlitchEventRatio composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type DeviceInfo_DeviceInfo_deviceGlitchEventRatio struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceInfo_DeviceInfo_deviceGlitchEventRatio instantiates a new DeviceInfo_DeviceInfo_deviceGlitchEventRatio and sets the default values.
func NewDeviceInfo_DeviceInfo_deviceGlitchEventRatio()(*DeviceInfo_DeviceInfo_deviceGlitchEventRatio) {
    m := &DeviceInfo_DeviceInfo_deviceGlitchEventRatio{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateDeviceInfo_DeviceInfo_deviceGlitchEventRatioFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceInfo_DeviceInfo_deviceGlitchEventRatioFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeviceInfo_DeviceInfo_deviceGlitchEventRatio()
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
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
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
func (m *DeviceInfo_DeviceInfo_deviceGlitchEventRatio) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceInfo_DeviceInfo_deviceGlitchEventRatio) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *DeviceInfo_DeviceInfo_deviceGlitchEventRatio) GetFloat()(*float32) {
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
func (m *DeviceInfo_DeviceInfo_deviceGlitchEventRatio) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *DeviceInfo_DeviceInfo_deviceGlitchEventRatio) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *DeviceInfo_DeviceInfo_deviceGlitchEventRatio) GetString()(*string) {
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
func (m *DeviceInfo_DeviceInfo_deviceGlitchEventRatio) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceInfo_DeviceInfo_deviceGlitchEventRatio) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *DeviceInfo_DeviceInfo_deviceGlitchEventRatio) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *DeviceInfo_DeviceInfo_deviceGlitchEventRatio) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *DeviceInfo_DeviceInfo_deviceGlitchEventRatio) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare instantiates a new DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare and sets the default values.
func NewDeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare()(*DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare) {
    m := &DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateDeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquareFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquareFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare()
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
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
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
func (m *DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare) GetFloat()(*float32) {
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
func (m *DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare) GetString()(*string) {
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
func (m *DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// DeviceInfo_DeviceInfo_lowSpeechLevelEventRatio composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type DeviceInfo_DeviceInfo_lowSpeechLevelEventRatio struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceInfo_DeviceInfo_lowSpeechLevelEventRatio instantiates a new DeviceInfo_DeviceInfo_lowSpeechLevelEventRatio and sets the default values.
func NewDeviceInfo_DeviceInfo_lowSpeechLevelEventRatio()(*DeviceInfo_DeviceInfo_lowSpeechLevelEventRatio) {
    m := &DeviceInfo_DeviceInfo_lowSpeechLevelEventRatio{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateDeviceInfo_DeviceInfo_lowSpeechLevelEventRatioFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceInfo_DeviceInfo_lowSpeechLevelEventRatioFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeviceInfo_DeviceInfo_lowSpeechLevelEventRatio()
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
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
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
func (m *DeviceInfo_DeviceInfo_lowSpeechLevelEventRatio) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceInfo_DeviceInfo_lowSpeechLevelEventRatio) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *DeviceInfo_DeviceInfo_lowSpeechLevelEventRatio) GetFloat()(*float32) {
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
func (m *DeviceInfo_DeviceInfo_lowSpeechLevelEventRatio) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *DeviceInfo_DeviceInfo_lowSpeechLevelEventRatio) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *DeviceInfo_DeviceInfo_lowSpeechLevelEventRatio) GetString()(*string) {
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
func (m *DeviceInfo_DeviceInfo_lowSpeechLevelEventRatio) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceInfo_DeviceInfo_lowSpeechLevelEventRatio) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *DeviceInfo_DeviceInfo_lowSpeechLevelEventRatio) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *DeviceInfo_DeviceInfo_lowSpeechLevelEventRatio) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *DeviceInfo_DeviceInfo_lowSpeechLevelEventRatio) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio instantiates a new DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio and sets the default values.
func NewDeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio()(*DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio) {
    m := &DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateDeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatioFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatioFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio()
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
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
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
func (m *DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio) GetFloat()(*float32) {
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
func (m *DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio) GetString()(*string) {
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
func (m *DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// DeviceInfo_DeviceInfo_micGlitchRate composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type DeviceInfo_DeviceInfo_micGlitchRate struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceInfo_DeviceInfo_micGlitchRate instantiates a new DeviceInfo_DeviceInfo_micGlitchRate and sets the default values.
func NewDeviceInfo_DeviceInfo_micGlitchRate()(*DeviceInfo_DeviceInfo_micGlitchRate) {
    m := &DeviceInfo_DeviceInfo_micGlitchRate{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateDeviceInfo_DeviceInfo_micGlitchRateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceInfo_DeviceInfo_micGlitchRateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeviceInfo_DeviceInfo_micGlitchRate()
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
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
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
func (m *DeviceInfo_DeviceInfo_micGlitchRate) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceInfo_DeviceInfo_micGlitchRate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *DeviceInfo_DeviceInfo_micGlitchRate) GetFloat()(*float32) {
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
func (m *DeviceInfo_DeviceInfo_micGlitchRate) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *DeviceInfo_DeviceInfo_micGlitchRate) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *DeviceInfo_DeviceInfo_micGlitchRate) GetString()(*string) {
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
func (m *DeviceInfo_DeviceInfo_micGlitchRate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceInfo_DeviceInfo_micGlitchRate) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *DeviceInfo_DeviceInfo_micGlitchRate) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *DeviceInfo_DeviceInfo_micGlitchRate) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *DeviceInfo_DeviceInfo_micGlitchRate) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// DeviceInfo_DeviceInfo_renderMuteEventRatio composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type DeviceInfo_DeviceInfo_renderMuteEventRatio struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceInfo_DeviceInfo_renderMuteEventRatio instantiates a new DeviceInfo_DeviceInfo_renderMuteEventRatio and sets the default values.
func NewDeviceInfo_DeviceInfo_renderMuteEventRatio()(*DeviceInfo_DeviceInfo_renderMuteEventRatio) {
    m := &DeviceInfo_DeviceInfo_renderMuteEventRatio{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateDeviceInfo_DeviceInfo_renderMuteEventRatioFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceInfo_DeviceInfo_renderMuteEventRatioFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeviceInfo_DeviceInfo_renderMuteEventRatio()
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
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
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
func (m *DeviceInfo_DeviceInfo_renderMuteEventRatio) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceInfo_DeviceInfo_renderMuteEventRatio) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *DeviceInfo_DeviceInfo_renderMuteEventRatio) GetFloat()(*float32) {
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
func (m *DeviceInfo_DeviceInfo_renderMuteEventRatio) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *DeviceInfo_DeviceInfo_renderMuteEventRatio) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *DeviceInfo_DeviceInfo_renderMuteEventRatio) GetString()(*string) {
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
func (m *DeviceInfo_DeviceInfo_renderMuteEventRatio) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceInfo_DeviceInfo_renderMuteEventRatio) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *DeviceInfo_DeviceInfo_renderMuteEventRatio) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *DeviceInfo_DeviceInfo_renderMuteEventRatio) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *DeviceInfo_DeviceInfo_renderMuteEventRatio) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// DeviceInfo_DeviceInfo_renderNotFunctioningEventRatio composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type DeviceInfo_DeviceInfo_renderNotFunctioningEventRatio struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceInfo_DeviceInfo_renderNotFunctioningEventRatio instantiates a new DeviceInfo_DeviceInfo_renderNotFunctioningEventRatio and sets the default values.
func NewDeviceInfo_DeviceInfo_renderNotFunctioningEventRatio()(*DeviceInfo_DeviceInfo_renderNotFunctioningEventRatio) {
    m := &DeviceInfo_DeviceInfo_renderNotFunctioningEventRatio{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateDeviceInfo_DeviceInfo_renderNotFunctioningEventRatioFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceInfo_DeviceInfo_renderNotFunctioningEventRatioFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeviceInfo_DeviceInfo_renderNotFunctioningEventRatio()
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
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
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
func (m *DeviceInfo_DeviceInfo_renderNotFunctioningEventRatio) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceInfo_DeviceInfo_renderNotFunctioningEventRatio) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *DeviceInfo_DeviceInfo_renderNotFunctioningEventRatio) GetFloat()(*float32) {
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
func (m *DeviceInfo_DeviceInfo_renderNotFunctioningEventRatio) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *DeviceInfo_DeviceInfo_renderNotFunctioningEventRatio) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *DeviceInfo_DeviceInfo_renderNotFunctioningEventRatio) GetString()(*string) {
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
func (m *DeviceInfo_DeviceInfo_renderNotFunctioningEventRatio) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceInfo_DeviceInfo_renderNotFunctioningEventRatio) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *DeviceInfo_DeviceInfo_renderNotFunctioningEventRatio) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *DeviceInfo_DeviceInfo_renderNotFunctioningEventRatio) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *DeviceInfo_DeviceInfo_renderNotFunctioningEventRatio) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// DeviceInfo_DeviceInfo_renderZeroVolumeEventRatio composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type DeviceInfo_DeviceInfo_renderZeroVolumeEventRatio struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceInfo_DeviceInfo_renderZeroVolumeEventRatio instantiates a new DeviceInfo_DeviceInfo_renderZeroVolumeEventRatio and sets the default values.
func NewDeviceInfo_DeviceInfo_renderZeroVolumeEventRatio()(*DeviceInfo_DeviceInfo_renderZeroVolumeEventRatio) {
    m := &DeviceInfo_DeviceInfo_renderZeroVolumeEventRatio{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateDeviceInfo_DeviceInfo_renderZeroVolumeEventRatioFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceInfo_DeviceInfo_renderZeroVolumeEventRatioFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeviceInfo_DeviceInfo_renderZeroVolumeEventRatio()
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
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
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
func (m *DeviceInfo_DeviceInfo_renderZeroVolumeEventRatio) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceInfo_DeviceInfo_renderZeroVolumeEventRatio) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *DeviceInfo_DeviceInfo_renderZeroVolumeEventRatio) GetFloat()(*float32) {
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
func (m *DeviceInfo_DeviceInfo_renderZeroVolumeEventRatio) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *DeviceInfo_DeviceInfo_renderZeroVolumeEventRatio) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *DeviceInfo_DeviceInfo_renderZeroVolumeEventRatio) GetString()(*string) {
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
func (m *DeviceInfo_DeviceInfo_renderZeroVolumeEventRatio) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceInfo_DeviceInfo_renderZeroVolumeEventRatio) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *DeviceInfo_DeviceInfo_renderZeroVolumeEventRatio) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *DeviceInfo_DeviceInfo_renderZeroVolumeEventRatio) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *DeviceInfo_DeviceInfo_renderZeroVolumeEventRatio) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// DeviceInfo_DeviceInfo_speakerGlitchRate composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type DeviceInfo_DeviceInfo_speakerGlitchRate struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceInfo_DeviceInfo_speakerGlitchRate instantiates a new DeviceInfo_DeviceInfo_speakerGlitchRate and sets the default values.
func NewDeviceInfo_DeviceInfo_speakerGlitchRate()(*DeviceInfo_DeviceInfo_speakerGlitchRate) {
    m := &DeviceInfo_DeviceInfo_speakerGlitchRate{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateDeviceInfo_DeviceInfo_speakerGlitchRateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceInfo_DeviceInfo_speakerGlitchRateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeviceInfo_DeviceInfo_speakerGlitchRate()
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
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
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
func (m *DeviceInfo_DeviceInfo_speakerGlitchRate) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceInfo_DeviceInfo_speakerGlitchRate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *DeviceInfo_DeviceInfo_speakerGlitchRate) GetFloat()(*float32) {
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
func (m *DeviceInfo_DeviceInfo_speakerGlitchRate) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *DeviceInfo_DeviceInfo_speakerGlitchRate) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *DeviceInfo_DeviceInfo_speakerGlitchRate) GetString()(*string) {
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
func (m *DeviceInfo_DeviceInfo_speakerGlitchRate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceInfo_DeviceInfo_speakerGlitchRate) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *DeviceInfo_DeviceInfo_speakerGlitchRate) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *DeviceInfo_DeviceInfo_speakerGlitchRate) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *DeviceInfo_DeviceInfo_speakerGlitchRate) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type DeviceInfo_DeviceInfo_captureNotFunctioningEventRatioable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type DeviceInfo_DeviceInfo_cpuInsufficentEventRatioable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type DeviceInfo_DeviceInfo_deviceClippingEventRatioable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type DeviceInfo_DeviceInfo_deviceGlitchEventRatioable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquareable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type DeviceInfo_DeviceInfo_lowSpeechLevelEventRatioable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatioable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type DeviceInfo_DeviceInfo_micGlitchRateable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type DeviceInfo_DeviceInfo_renderMuteEventRatioable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type DeviceInfo_DeviceInfo_renderNotFunctioningEventRatioable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type DeviceInfo_DeviceInfo_renderZeroVolumeEventRatioable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type DeviceInfo_DeviceInfo_speakerGlitchRateable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
// NewDeviceInfo instantiates a new DeviceInfo and sets the default values.
func NewDeviceInfo()(*DeviceInfo) {
    m := &DeviceInfo{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeviceInfo) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *DeviceInfo) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCaptureDeviceDriver gets the captureDeviceDriver property value. Name of the capture device driver used by the media endpoint.
// returns a *string when successful
func (m *DeviceInfo) GetCaptureDeviceDriver()(*string) {
    val, err := m.GetBackingStore().Get("captureDeviceDriver")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCaptureDeviceName gets the captureDeviceName property value. Name of the capture device used by the media endpoint.
// returns a *string when successful
func (m *DeviceInfo) GetCaptureDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("captureDeviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCaptureNotFunctioningEventRatio gets the captureNotFunctioningEventRatio property value. Fraction of the call that the media endpoint detected the capture device was not working properly.
// returns a DeviceInfo_DeviceInfo_captureNotFunctioningEventRatioable when successful
func (m *DeviceInfo) GetCaptureNotFunctioningEventRatio()(DeviceInfo_DeviceInfo_captureNotFunctioningEventRatioable) {
    val, err := m.GetBackingStore().Get("captureNotFunctioningEventRatio")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceInfo_DeviceInfo_captureNotFunctioningEventRatioable)
    }
    return nil
}
// GetCpuInsufficentEventRatio gets the cpuInsufficentEventRatio property value. Fraction of the call that the media endpoint detected the CPU resources available were insufficient and caused poor quality of the audio sent and received.
// returns a DeviceInfo_DeviceInfo_cpuInsufficentEventRatioable when successful
func (m *DeviceInfo) GetCpuInsufficentEventRatio()(DeviceInfo_DeviceInfo_cpuInsufficentEventRatioable) {
    val, err := m.GetBackingStore().Get("cpuInsufficentEventRatio")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceInfo_DeviceInfo_cpuInsufficentEventRatioable)
    }
    return nil
}
// GetDeviceClippingEventRatio gets the deviceClippingEventRatio property value. Fraction of the call that the media endpoint detected clipping in the captured audio that caused poor quality of the audio being sent.
// returns a DeviceInfo_DeviceInfo_deviceClippingEventRatioable when successful
func (m *DeviceInfo) GetDeviceClippingEventRatio()(DeviceInfo_DeviceInfo_deviceClippingEventRatioable) {
    val, err := m.GetBackingStore().Get("deviceClippingEventRatio")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceInfo_DeviceInfo_deviceClippingEventRatioable)
    }
    return nil
}
// GetDeviceGlitchEventRatio gets the deviceGlitchEventRatio property value. Fraction of the call that the media endpoint detected glitches or gaps in the audio played or captured that caused poor quality of the audio being sent or received.
// returns a DeviceInfo_DeviceInfo_deviceGlitchEventRatioable when successful
func (m *DeviceInfo) GetDeviceGlitchEventRatio()(DeviceInfo_DeviceInfo_deviceGlitchEventRatioable) {
    val, err := m.GetBackingStore().Get("deviceGlitchEventRatio")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceInfo_DeviceInfo_deviceGlitchEventRatioable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["captureDeviceDriver"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptureDeviceDriver(val)
        }
        return nil
    }
    res["captureDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptureDeviceName(val)
        }
        return nil
    }
    res["captureNotFunctioningEventRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceInfo_DeviceInfo_captureNotFunctioningEventRatioFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptureNotFunctioningEventRatio(val.(*DeviceInfo_DeviceInfo_captureNotFunctioningEventRatio))
        }
        return nil
    }
    res["cpuInsufficentEventRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceInfo_DeviceInfo_cpuInsufficentEventRatioFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuInsufficentEventRatio(val.(*DeviceInfo_DeviceInfo_cpuInsufficentEventRatio))
        }
        return nil
    }
    res["deviceClippingEventRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceInfo_DeviceInfo_deviceClippingEventRatioFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceClippingEventRatio(val.(*DeviceInfo_DeviceInfo_deviceClippingEventRatio))
        }
        return nil
    }
    res["deviceGlitchEventRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceInfo_DeviceInfo_deviceGlitchEventRatioFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGlitchEventRatio(val.(*DeviceInfo_DeviceInfo_deviceGlitchEventRatio))
        }
        return nil
    }
    res["howlingEventCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHowlingEventCount(val)
        }
        return nil
    }
    res["initialSignalLevelRootMeanSquare"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquareFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitialSignalLevelRootMeanSquare(val.(*DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquare))
        }
        return nil
    }
    res["lowSpeechLevelEventRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceInfo_DeviceInfo_lowSpeechLevelEventRatioFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLowSpeechLevelEventRatio(val.(*DeviceInfo_DeviceInfo_lowSpeechLevelEventRatio))
        }
        return nil
    }
    res["lowSpeechToNoiseEventRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatioFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLowSpeechToNoiseEventRatio(val.(*DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatio))
        }
        return nil
    }
    res["micGlitchRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceInfo_DeviceInfo_micGlitchRateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicGlitchRate(val.(*DeviceInfo_DeviceInfo_micGlitchRate))
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
    res["receivedNoiseLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedNoiseLevel(val)
        }
        return nil
    }
    res["receivedSignalLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedSignalLevel(val)
        }
        return nil
    }
    res["renderDeviceDriver"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenderDeviceDriver(val)
        }
        return nil
    }
    res["renderDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenderDeviceName(val)
        }
        return nil
    }
    res["renderMuteEventRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceInfo_DeviceInfo_renderMuteEventRatioFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenderMuteEventRatio(val.(*DeviceInfo_DeviceInfo_renderMuteEventRatio))
        }
        return nil
    }
    res["renderNotFunctioningEventRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceInfo_DeviceInfo_renderNotFunctioningEventRatioFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenderNotFunctioningEventRatio(val.(*DeviceInfo_DeviceInfo_renderNotFunctioningEventRatio))
        }
        return nil
    }
    res["renderZeroVolumeEventRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceInfo_DeviceInfo_renderZeroVolumeEventRatioFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenderZeroVolumeEventRatio(val.(*DeviceInfo_DeviceInfo_renderZeroVolumeEventRatio))
        }
        return nil
    }
    res["sentNoiseLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentNoiseLevel(val)
        }
        return nil
    }
    res["sentSignalLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentSignalLevel(val)
        }
        return nil
    }
    res["speakerGlitchRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceInfo_DeviceInfo_speakerGlitchRateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeakerGlitchRate(val.(*DeviceInfo_DeviceInfo_speakerGlitchRate))
        }
        return nil
    }
    return res
}
// GetHowlingEventCount gets the howlingEventCount property value. Number of times during the call that the media endpoint detected howling or screeching audio.
// returns a *int32 when successful
func (m *DeviceInfo) GetHowlingEventCount()(*int32) {
    val, err := m.GetBackingStore().Get("howlingEventCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetInitialSignalLevelRootMeanSquare gets the initialSignalLevelRootMeanSquare property value. The root mean square (RMS) of the incoming signal of up to the first 30 seconds of the call.
// returns a DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquareable when successful
func (m *DeviceInfo) GetInitialSignalLevelRootMeanSquare()(DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquareable) {
    val, err := m.GetBackingStore().Get("initialSignalLevelRootMeanSquare")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquareable)
    }
    return nil
}
// GetLowSpeechLevelEventRatio gets the lowSpeechLevelEventRatio property value. Fraction of the call that the media endpoint detected low speech level that caused poor quality of the audio being sent.
// returns a DeviceInfo_DeviceInfo_lowSpeechLevelEventRatioable when successful
func (m *DeviceInfo) GetLowSpeechLevelEventRatio()(DeviceInfo_DeviceInfo_lowSpeechLevelEventRatioable) {
    val, err := m.GetBackingStore().Get("lowSpeechLevelEventRatio")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceInfo_DeviceInfo_lowSpeechLevelEventRatioable)
    }
    return nil
}
// GetLowSpeechToNoiseEventRatio gets the lowSpeechToNoiseEventRatio property value. Fraction of the call that the media endpoint detected low speech to noise level that caused poor quality of the audio being sent.
// returns a DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatioable when successful
func (m *DeviceInfo) GetLowSpeechToNoiseEventRatio()(DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatioable) {
    val, err := m.GetBackingStore().Get("lowSpeechToNoiseEventRatio")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatioable)
    }
    return nil
}
// GetMicGlitchRate gets the micGlitchRate property value. Glitches per 5 minute interval for the media endpoint's microphone.
// returns a DeviceInfo_DeviceInfo_micGlitchRateable when successful
func (m *DeviceInfo) GetMicGlitchRate()(DeviceInfo_DeviceInfo_micGlitchRateable) {
    val, err := m.GetBackingStore().Get("micGlitchRate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceInfo_DeviceInfo_micGlitchRateable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DeviceInfo) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReceivedNoiseLevel gets the receivedNoiseLevel property value. Average energy level of received audio for audio classified as mono noise or left channel of stereo noise by the media endpoint.
// returns a *int32 when successful
func (m *DeviceInfo) GetReceivedNoiseLevel()(*int32) {
    val, err := m.GetBackingStore().Get("receivedNoiseLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetReceivedSignalLevel gets the receivedSignalLevel property value. Average energy level of received audio for audio classified as mono speech, or left channel of stereo speech by the media endpoint.
// returns a *int32 when successful
func (m *DeviceInfo) GetReceivedSignalLevel()(*int32) {
    val, err := m.GetBackingStore().Get("receivedSignalLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRenderDeviceDriver gets the renderDeviceDriver property value. Name of the render device driver used by the media endpoint.
// returns a *string when successful
func (m *DeviceInfo) GetRenderDeviceDriver()(*string) {
    val, err := m.GetBackingStore().Get("renderDeviceDriver")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRenderDeviceName gets the renderDeviceName property value. Name of the render device used by the media endpoint.
// returns a *string when successful
func (m *DeviceInfo) GetRenderDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("renderDeviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRenderMuteEventRatio gets the renderMuteEventRatio property value. Fraction of the call that media endpoint detected device render is muted.
// returns a DeviceInfo_DeviceInfo_renderMuteEventRatioable when successful
func (m *DeviceInfo) GetRenderMuteEventRatio()(DeviceInfo_DeviceInfo_renderMuteEventRatioable) {
    val, err := m.GetBackingStore().Get("renderMuteEventRatio")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceInfo_DeviceInfo_renderMuteEventRatioable)
    }
    return nil
}
// GetRenderNotFunctioningEventRatio gets the renderNotFunctioningEventRatio property value. Fraction of the call that the media endpoint detected the render device was not working properly.
// returns a DeviceInfo_DeviceInfo_renderNotFunctioningEventRatioable when successful
func (m *DeviceInfo) GetRenderNotFunctioningEventRatio()(DeviceInfo_DeviceInfo_renderNotFunctioningEventRatioable) {
    val, err := m.GetBackingStore().Get("renderNotFunctioningEventRatio")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceInfo_DeviceInfo_renderNotFunctioningEventRatioable)
    }
    return nil
}
// GetRenderZeroVolumeEventRatio gets the renderZeroVolumeEventRatio property value. Fraction of the call that media endpoint detected device render volume is set to 0.
// returns a DeviceInfo_DeviceInfo_renderZeroVolumeEventRatioable when successful
func (m *DeviceInfo) GetRenderZeroVolumeEventRatio()(DeviceInfo_DeviceInfo_renderZeroVolumeEventRatioable) {
    val, err := m.GetBackingStore().Get("renderZeroVolumeEventRatio")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceInfo_DeviceInfo_renderZeroVolumeEventRatioable)
    }
    return nil
}
// GetSentNoiseLevel gets the sentNoiseLevel property value. Average energy level of sent audio for audio classified as mono noise or left channel of stereo noise by the media endpoint.
// returns a *int32 when successful
func (m *DeviceInfo) GetSentNoiseLevel()(*int32) {
    val, err := m.GetBackingStore().Get("sentNoiseLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSentSignalLevel gets the sentSignalLevel property value. Average energy level of sent audio for audio classified as mono speech, or left channel of stereo speech by the media endpoint.
// returns a *int32 when successful
func (m *DeviceInfo) GetSentSignalLevel()(*int32) {
    val, err := m.GetBackingStore().Get("sentSignalLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSpeakerGlitchRate gets the speakerGlitchRate property value. Glitches per 5 minute internal for the media endpoint's loudspeaker.
// returns a DeviceInfo_DeviceInfo_speakerGlitchRateable when successful
func (m *DeviceInfo) GetSpeakerGlitchRate()(DeviceInfo_DeviceInfo_speakerGlitchRateable) {
    val, err := m.GetBackingStore().Get("speakerGlitchRate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceInfo_DeviceInfo_speakerGlitchRateable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("captureDeviceDriver", m.GetCaptureDeviceDriver())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("captureDeviceName", m.GetCaptureDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("captureNotFunctioningEventRatio", m.GetCaptureNotFunctioningEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("cpuInsufficentEventRatio", m.GetCpuInsufficentEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deviceClippingEventRatio", m.GetDeviceClippingEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deviceGlitchEventRatio", m.GetDeviceGlitchEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("howlingEventCount", m.GetHowlingEventCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("initialSignalLevelRootMeanSquare", m.GetInitialSignalLevelRootMeanSquare())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lowSpeechLevelEventRatio", m.GetLowSpeechLevelEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lowSpeechToNoiseEventRatio", m.GetLowSpeechToNoiseEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("micGlitchRate", m.GetMicGlitchRate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("receivedNoiseLevel", m.GetReceivedNoiseLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("receivedSignalLevel", m.GetReceivedSignalLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("renderDeviceDriver", m.GetRenderDeviceDriver())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("renderDeviceName", m.GetRenderDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("renderMuteEventRatio", m.GetRenderMuteEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("renderNotFunctioningEventRatio", m.GetRenderNotFunctioningEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("renderZeroVolumeEventRatio", m.GetRenderZeroVolumeEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("sentNoiseLevel", m.GetSentNoiseLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("sentSignalLevel", m.GetSentSignalLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("speakerGlitchRate", m.GetSpeakerGlitchRate())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceInfo) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DeviceInfo) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCaptureDeviceDriver sets the captureDeviceDriver property value. Name of the capture device driver used by the media endpoint.
func (m *DeviceInfo) SetCaptureDeviceDriver(value *string)() {
    err := m.GetBackingStore().Set("captureDeviceDriver", value)
    if err != nil {
        panic(err)
    }
}
// SetCaptureDeviceName sets the captureDeviceName property value. Name of the capture device used by the media endpoint.
func (m *DeviceInfo) SetCaptureDeviceName(value *string)() {
    err := m.GetBackingStore().Set("captureDeviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetCaptureNotFunctioningEventRatio sets the captureNotFunctioningEventRatio property value. Fraction of the call that the media endpoint detected the capture device was not working properly.
func (m *DeviceInfo) SetCaptureNotFunctioningEventRatio(value DeviceInfo_DeviceInfo_captureNotFunctioningEventRatioable)() {
    err := m.GetBackingStore().Set("captureNotFunctioningEventRatio", value)
    if err != nil {
        panic(err)
    }
}
// SetCpuInsufficentEventRatio sets the cpuInsufficentEventRatio property value. Fraction of the call that the media endpoint detected the CPU resources available were insufficient and caused poor quality of the audio sent and received.
func (m *DeviceInfo) SetCpuInsufficentEventRatio(value DeviceInfo_DeviceInfo_cpuInsufficentEventRatioable)() {
    err := m.GetBackingStore().Set("cpuInsufficentEventRatio", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceClippingEventRatio sets the deviceClippingEventRatio property value. Fraction of the call that the media endpoint detected clipping in the captured audio that caused poor quality of the audio being sent.
func (m *DeviceInfo) SetDeviceClippingEventRatio(value DeviceInfo_DeviceInfo_deviceClippingEventRatioable)() {
    err := m.GetBackingStore().Set("deviceClippingEventRatio", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceGlitchEventRatio sets the deviceGlitchEventRatio property value. Fraction of the call that the media endpoint detected glitches or gaps in the audio played or captured that caused poor quality of the audio being sent or received.
func (m *DeviceInfo) SetDeviceGlitchEventRatio(value DeviceInfo_DeviceInfo_deviceGlitchEventRatioable)() {
    err := m.GetBackingStore().Set("deviceGlitchEventRatio", value)
    if err != nil {
        panic(err)
    }
}
// SetHowlingEventCount sets the howlingEventCount property value. Number of times during the call that the media endpoint detected howling or screeching audio.
func (m *DeviceInfo) SetHowlingEventCount(value *int32)() {
    err := m.GetBackingStore().Set("howlingEventCount", value)
    if err != nil {
        panic(err)
    }
}
// SetInitialSignalLevelRootMeanSquare sets the initialSignalLevelRootMeanSquare property value. The root mean square (RMS) of the incoming signal of up to the first 30 seconds of the call.
func (m *DeviceInfo) SetInitialSignalLevelRootMeanSquare(value DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquareable)() {
    err := m.GetBackingStore().Set("initialSignalLevelRootMeanSquare", value)
    if err != nil {
        panic(err)
    }
}
// SetLowSpeechLevelEventRatio sets the lowSpeechLevelEventRatio property value. Fraction of the call that the media endpoint detected low speech level that caused poor quality of the audio being sent.
func (m *DeviceInfo) SetLowSpeechLevelEventRatio(value DeviceInfo_DeviceInfo_lowSpeechLevelEventRatioable)() {
    err := m.GetBackingStore().Set("lowSpeechLevelEventRatio", value)
    if err != nil {
        panic(err)
    }
}
// SetLowSpeechToNoiseEventRatio sets the lowSpeechToNoiseEventRatio property value. Fraction of the call that the media endpoint detected low speech to noise level that caused poor quality of the audio being sent.
func (m *DeviceInfo) SetLowSpeechToNoiseEventRatio(value DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatioable)() {
    err := m.GetBackingStore().Set("lowSpeechToNoiseEventRatio", value)
    if err != nil {
        panic(err)
    }
}
// SetMicGlitchRate sets the micGlitchRate property value. Glitches per 5 minute interval for the media endpoint's microphone.
func (m *DeviceInfo) SetMicGlitchRate(value DeviceInfo_DeviceInfo_micGlitchRateable)() {
    err := m.GetBackingStore().Set("micGlitchRate", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceInfo) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetReceivedNoiseLevel sets the receivedNoiseLevel property value. Average energy level of received audio for audio classified as mono noise or left channel of stereo noise by the media endpoint.
func (m *DeviceInfo) SetReceivedNoiseLevel(value *int32)() {
    err := m.GetBackingStore().Set("receivedNoiseLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetReceivedSignalLevel sets the receivedSignalLevel property value. Average energy level of received audio for audio classified as mono speech, or left channel of stereo speech by the media endpoint.
func (m *DeviceInfo) SetReceivedSignalLevel(value *int32)() {
    err := m.GetBackingStore().Set("receivedSignalLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetRenderDeviceDriver sets the renderDeviceDriver property value. Name of the render device driver used by the media endpoint.
func (m *DeviceInfo) SetRenderDeviceDriver(value *string)() {
    err := m.GetBackingStore().Set("renderDeviceDriver", value)
    if err != nil {
        panic(err)
    }
}
// SetRenderDeviceName sets the renderDeviceName property value. Name of the render device used by the media endpoint.
func (m *DeviceInfo) SetRenderDeviceName(value *string)() {
    err := m.GetBackingStore().Set("renderDeviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetRenderMuteEventRatio sets the renderMuteEventRatio property value. Fraction of the call that media endpoint detected device render is muted.
func (m *DeviceInfo) SetRenderMuteEventRatio(value DeviceInfo_DeviceInfo_renderMuteEventRatioable)() {
    err := m.GetBackingStore().Set("renderMuteEventRatio", value)
    if err != nil {
        panic(err)
    }
}
// SetRenderNotFunctioningEventRatio sets the renderNotFunctioningEventRatio property value. Fraction of the call that the media endpoint detected the render device was not working properly.
func (m *DeviceInfo) SetRenderNotFunctioningEventRatio(value DeviceInfo_DeviceInfo_renderNotFunctioningEventRatioable)() {
    err := m.GetBackingStore().Set("renderNotFunctioningEventRatio", value)
    if err != nil {
        panic(err)
    }
}
// SetRenderZeroVolumeEventRatio sets the renderZeroVolumeEventRatio property value. Fraction of the call that media endpoint detected device render volume is set to 0.
func (m *DeviceInfo) SetRenderZeroVolumeEventRatio(value DeviceInfo_DeviceInfo_renderZeroVolumeEventRatioable)() {
    err := m.GetBackingStore().Set("renderZeroVolumeEventRatio", value)
    if err != nil {
        panic(err)
    }
}
// SetSentNoiseLevel sets the sentNoiseLevel property value. Average energy level of sent audio for audio classified as mono noise or left channel of stereo noise by the media endpoint.
func (m *DeviceInfo) SetSentNoiseLevel(value *int32)() {
    err := m.GetBackingStore().Set("sentNoiseLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetSentSignalLevel sets the sentSignalLevel property value. Average energy level of sent audio for audio classified as mono speech, or left channel of stereo speech by the media endpoint.
func (m *DeviceInfo) SetSentSignalLevel(value *int32)() {
    err := m.GetBackingStore().Set("sentSignalLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetSpeakerGlitchRate sets the speakerGlitchRate property value. Glitches per 5 minute internal for the media endpoint's loudspeaker.
func (m *DeviceInfo) SetSpeakerGlitchRate(value DeviceInfo_DeviceInfo_speakerGlitchRateable)() {
    err := m.GetBackingStore().Set("speakerGlitchRate", value)
    if err != nil {
        panic(err)
    }
}
type DeviceInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCaptureDeviceDriver()(*string)
    GetCaptureDeviceName()(*string)
    GetCaptureNotFunctioningEventRatio()(DeviceInfo_DeviceInfo_captureNotFunctioningEventRatioable)
    GetCpuInsufficentEventRatio()(DeviceInfo_DeviceInfo_cpuInsufficentEventRatioable)
    GetDeviceClippingEventRatio()(DeviceInfo_DeviceInfo_deviceClippingEventRatioable)
    GetDeviceGlitchEventRatio()(DeviceInfo_DeviceInfo_deviceGlitchEventRatioable)
    GetHowlingEventCount()(*int32)
    GetInitialSignalLevelRootMeanSquare()(DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquareable)
    GetLowSpeechLevelEventRatio()(DeviceInfo_DeviceInfo_lowSpeechLevelEventRatioable)
    GetLowSpeechToNoiseEventRatio()(DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatioable)
    GetMicGlitchRate()(DeviceInfo_DeviceInfo_micGlitchRateable)
    GetOdataType()(*string)
    GetReceivedNoiseLevel()(*int32)
    GetReceivedSignalLevel()(*int32)
    GetRenderDeviceDriver()(*string)
    GetRenderDeviceName()(*string)
    GetRenderMuteEventRatio()(DeviceInfo_DeviceInfo_renderMuteEventRatioable)
    GetRenderNotFunctioningEventRatio()(DeviceInfo_DeviceInfo_renderNotFunctioningEventRatioable)
    GetRenderZeroVolumeEventRatio()(DeviceInfo_DeviceInfo_renderZeroVolumeEventRatioable)
    GetSentNoiseLevel()(*int32)
    GetSentSignalLevel()(*int32)
    GetSpeakerGlitchRate()(DeviceInfo_DeviceInfo_speakerGlitchRateable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCaptureDeviceDriver(value *string)()
    SetCaptureDeviceName(value *string)()
    SetCaptureNotFunctioningEventRatio(value DeviceInfo_DeviceInfo_captureNotFunctioningEventRatioable)()
    SetCpuInsufficentEventRatio(value DeviceInfo_DeviceInfo_cpuInsufficentEventRatioable)()
    SetDeviceClippingEventRatio(value DeviceInfo_DeviceInfo_deviceClippingEventRatioable)()
    SetDeviceGlitchEventRatio(value DeviceInfo_DeviceInfo_deviceGlitchEventRatioable)()
    SetHowlingEventCount(value *int32)()
    SetInitialSignalLevelRootMeanSquare(value DeviceInfo_DeviceInfo_initialSignalLevelRootMeanSquareable)()
    SetLowSpeechLevelEventRatio(value DeviceInfo_DeviceInfo_lowSpeechLevelEventRatioable)()
    SetLowSpeechToNoiseEventRatio(value DeviceInfo_DeviceInfo_lowSpeechToNoiseEventRatioable)()
    SetMicGlitchRate(value DeviceInfo_DeviceInfo_micGlitchRateable)()
    SetOdataType(value *string)()
    SetReceivedNoiseLevel(value *int32)()
    SetReceivedSignalLevel(value *int32)()
    SetRenderDeviceDriver(value *string)()
    SetRenderDeviceName(value *string)()
    SetRenderMuteEventRatio(value DeviceInfo_DeviceInfo_renderMuteEventRatioable)()
    SetRenderNotFunctioningEventRatio(value DeviceInfo_DeviceInfo_renderNotFunctioningEventRatioable)()
    SetRenderZeroVolumeEventRatio(value DeviceInfo_DeviceInfo_renderZeroVolumeEventRatioable)()
    SetSentNoiseLevel(value *int32)()
    SetSentSignalLevel(value *int32)()
    SetSpeakerGlitchRate(value DeviceInfo_DeviceInfo_speakerGlitchRateable)()
}
