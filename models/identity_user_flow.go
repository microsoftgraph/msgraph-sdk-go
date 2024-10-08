package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type IdentityUserFlow struct {
    Entity
}
// IdentityUserFlow_IdentityUserFlow_userFlowTypeVersion composed type wrapper for classes float32, ReferenceNumeric, string
type IdentityUserFlow_IdentityUserFlow_userFlowTypeVersion struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewIdentityUserFlow_IdentityUserFlow_userFlowTypeVersion instantiates a new IdentityUserFlow_IdentityUserFlow_userFlowTypeVersion and sets the default values.
func NewIdentityUserFlow_IdentityUserFlow_userFlowTypeVersion()(*IdentityUserFlow_IdentityUserFlow_userFlowTypeVersion) {
    m := &IdentityUserFlow_IdentityUserFlow_userFlowTypeVersion{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateIdentityUserFlow_IdentityUserFlow_userFlowTypeVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentityUserFlow_IdentityUserFlow_userFlowTypeVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewIdentityUserFlow_IdentityUserFlow_userFlowTypeVersion()
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
func (m *IdentityUserFlow_IdentityUserFlow_userFlowTypeVersion) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IdentityUserFlow_IdentityUserFlow_userFlowTypeVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *IdentityUserFlow_IdentityUserFlow_userFlowTypeVersion) GetFloat()(*float32) {
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
func (m *IdentityUserFlow_IdentityUserFlow_userFlowTypeVersion) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *IdentityUserFlow_IdentityUserFlow_userFlowTypeVersion) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *IdentityUserFlow_IdentityUserFlow_userFlowTypeVersion) GetString()(*string) {
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
func (m *IdentityUserFlow_IdentityUserFlow_userFlowTypeVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *IdentityUserFlow_IdentityUserFlow_userFlowTypeVersion) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *IdentityUserFlow_IdentityUserFlow_userFlowTypeVersion) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *IdentityUserFlow_IdentityUserFlow_userFlowTypeVersion) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *IdentityUserFlow_IdentityUserFlow_userFlowTypeVersion) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type IdentityUserFlow_IdentityUserFlow_userFlowTypeVersionable interface {
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
// NewIdentityUserFlow instantiates a new IdentityUserFlow and sets the default values.
func NewIdentityUserFlow()(*IdentityUserFlow) {
    m := &IdentityUserFlow{
        Entity: *NewEntity(),
    }
    return m
}
// CreateIdentityUserFlowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentityUserFlowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.b2xIdentityUserFlow":
                        return NewB2xIdentityUserFlow(), nil
                }
            }
        }
    }
    return NewIdentityUserFlow(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IdentityUserFlow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["userFlowType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserFlowType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserFlowType(val.(*UserFlowType))
        }
        return nil
    }
    res["userFlowTypeVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityUserFlow_IdentityUserFlow_userFlowTypeVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserFlowTypeVersion(val.(*IdentityUserFlow_IdentityUserFlow_userFlowTypeVersion))
        }
        return nil
    }
    return res
}
// GetUserFlowType gets the userFlowType property value. The userFlowType property
// returns a *UserFlowType when successful
func (m *IdentityUserFlow) GetUserFlowType()(*UserFlowType) {
    val, err := m.GetBackingStore().Get("userFlowType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserFlowType)
    }
    return nil
}
// GetUserFlowTypeVersion gets the userFlowTypeVersion property value. The userFlowTypeVersion property
// returns a IdentityUserFlow_IdentityUserFlow_userFlowTypeVersionable when successful
func (m *IdentityUserFlow) GetUserFlowTypeVersion()(IdentityUserFlow_IdentityUserFlow_userFlowTypeVersionable) {
    val, err := m.GetBackingStore().Get("userFlowTypeVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentityUserFlow_IdentityUserFlow_userFlowTypeVersionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IdentityUserFlow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUserFlowType() != nil {
        cast := (*m.GetUserFlowType()).String()
        err = writer.WriteStringValue("userFlowType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userFlowTypeVersion", m.GetUserFlowTypeVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserFlowType sets the userFlowType property value. The userFlowType property
func (m *IdentityUserFlow) SetUserFlowType(value *UserFlowType)() {
    err := m.GetBackingStore().Set("userFlowType", value)
    if err != nil {
        panic(err)
    }
}
// SetUserFlowTypeVersion sets the userFlowTypeVersion property value. The userFlowTypeVersion property
func (m *IdentityUserFlow) SetUserFlowTypeVersion(value IdentityUserFlow_IdentityUserFlow_userFlowTypeVersionable)() {
    err := m.GetBackingStore().Set("userFlowTypeVersion", value)
    if err != nil {
        panic(err)
    }
}
type IdentityUserFlowable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUserFlowType()(*UserFlowType)
    GetUserFlowTypeVersion()(IdentityUserFlow_IdentityUserFlow_userFlowTypeVersionable)
    SetUserFlowType(value *UserFlowType)()
    SetUserFlowTypeVersion(value IdentityUserFlow_IdentityUserFlow_userFlowTypeVersionable)()
}
