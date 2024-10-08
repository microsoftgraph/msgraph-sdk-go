package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type WorkbookRangeFont struct {
    Entity
}
// WorkbookRangeFont_WorkbookRangeFont_size composed type wrapper for classes float64, ReferenceNumeric, string
type WorkbookRangeFont_WorkbookRangeFont_size struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWorkbookRangeFont_WorkbookRangeFont_size instantiates a new WorkbookRangeFont_WorkbookRangeFont_size and sets the default values.
func NewWorkbookRangeFont_WorkbookRangeFont_size()(*WorkbookRangeFont_WorkbookRangeFont_size) {
    m := &WorkbookRangeFont_WorkbookRangeFont_size{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateWorkbookRangeFont_WorkbookRangeFont_sizeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookRangeFont_WorkbookRangeFont_sizeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWorkbookRangeFont_WorkbookRangeFont_size()
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
func (m *WorkbookRangeFont_WorkbookRangeFont_size) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *WorkbookRangeFont_WorkbookRangeFont_size) GetDouble()(*float64) {
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
func (m *WorkbookRangeFont_WorkbookRangeFont_size) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *WorkbookRangeFont_WorkbookRangeFont_size) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *WorkbookRangeFont_WorkbookRangeFont_size) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *WorkbookRangeFont_WorkbookRangeFont_size) GetString()(*string) {
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
func (m *WorkbookRangeFont_WorkbookRangeFont_size) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *WorkbookRangeFont_WorkbookRangeFont_size) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *WorkbookRangeFont_WorkbookRangeFont_size) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *WorkbookRangeFont_WorkbookRangeFont_size) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *WorkbookRangeFont_WorkbookRangeFont_size) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type WorkbookRangeFont_WorkbookRangeFont_sizeable interface {
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
// NewWorkbookRangeFont instantiates a new WorkbookRangeFont and sets the default values.
func NewWorkbookRangeFont()(*WorkbookRangeFont) {
    m := &WorkbookRangeFont{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookRangeFontFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookRangeFontFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookRangeFont(), nil
}
// GetBold gets the bold property value. Inidicates whether the font is bold.
// returns a *bool when successful
func (m *WorkbookRangeFont) GetBold()(*bool) {
    val, err := m.GetBackingStore().Get("bold")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetColor gets the color property value. The HTML color code representation of the text color. For example, #FF0000 represents the color red.
// returns a *string when successful
func (m *WorkbookRangeFont) GetColor()(*string) {
    val, err := m.GetBackingStore().Get("color")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookRangeFont) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBold(val)
        }
        return nil
    }
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
        }
        return nil
    }
    res["italic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItalic(val)
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
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookRangeFont_WorkbookRangeFont_sizeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val.(*WorkbookRangeFont_WorkbookRangeFont_size))
        }
        return nil
    }
    res["underline"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnderline(val)
        }
        return nil
    }
    return res
}
// GetItalic gets the italic property value. Inidicates whether the font is italic.
// returns a *bool when successful
func (m *WorkbookRangeFont) GetItalic()(*bool) {
    val, err := m.GetBackingStore().Get("italic")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetName gets the name property value. The font name. For example, 'Calibri'.
// returns a *string when successful
func (m *WorkbookRangeFont) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSize gets the size property value. The font size.
// returns a WorkbookRangeFont_WorkbookRangeFont_sizeable when successful
func (m *WorkbookRangeFont) GetSize()(WorkbookRangeFont_WorkbookRangeFont_sizeable) {
    val, err := m.GetBackingStore().Get("size")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookRangeFont_WorkbookRangeFont_sizeable)
    }
    return nil
}
// GetUnderline gets the underline property value. The type of underlining applied to the font. The possible values are: None, Single, Double, SingleAccountant, DoubleAccountant.
// returns a *string when successful
func (m *WorkbookRangeFont) GetUnderline()(*string) {
    val, err := m.GetBackingStore().Get("underline")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WorkbookRangeFont) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("bold", m.GetBold())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("italic", m.GetItalic())
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
    {
        err = writer.WriteObjectValue("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("underline", m.GetUnderline())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBold sets the bold property value. Inidicates whether the font is bold.
func (m *WorkbookRangeFont) SetBold(value *bool)() {
    err := m.GetBackingStore().Set("bold", value)
    if err != nil {
        panic(err)
    }
}
// SetColor sets the color property value. The HTML color code representation of the text color. For example, #FF0000 represents the color red.
func (m *WorkbookRangeFont) SetColor(value *string)() {
    err := m.GetBackingStore().Set("color", value)
    if err != nil {
        panic(err)
    }
}
// SetItalic sets the italic property value. Inidicates whether the font is italic.
func (m *WorkbookRangeFont) SetItalic(value *bool)() {
    err := m.GetBackingStore().Set("italic", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The font name. For example, 'Calibri'.
func (m *WorkbookRangeFont) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetSize sets the size property value. The font size.
func (m *WorkbookRangeFont) SetSize(value WorkbookRangeFont_WorkbookRangeFont_sizeable)() {
    err := m.GetBackingStore().Set("size", value)
    if err != nil {
        panic(err)
    }
}
// SetUnderline sets the underline property value. The type of underlining applied to the font. The possible values are: None, Single, Double, SingleAccountant, DoubleAccountant.
func (m *WorkbookRangeFont) SetUnderline(value *string)() {
    err := m.GetBackingStore().Set("underline", value)
    if err != nil {
        panic(err)
    }
}
type WorkbookRangeFontable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBold()(*bool)
    GetColor()(*string)
    GetItalic()(*bool)
    GetName()(*string)
    GetSize()(WorkbookRangeFont_WorkbookRangeFont_sizeable)
    GetUnderline()(*string)
    SetBold(value *bool)()
    SetColor(value *string)()
    SetItalic(value *bool)()
    SetName(value *string)()
    SetSize(value WorkbookRangeFont_WorkbookRangeFont_sizeable)()
    SetUnderline(value *string)()
}
