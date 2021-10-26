package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookNamedItem struct {
    Entity
    // Represents the comment associated with this name.
    comment *string;
    // The name of the object. Read-only.
    name *string;
    // Indicates whether the name is scoped to the workbook or to a specific worksheet. Read-only.
    scope *string;
    // Indicates what type of reference is associated with the name. The possible values are: String, Integer, Double, Boolean, Range. Read-only.
    type_escaped *string;
    // Represents the formula that the name is defined to refer to. E.g. =Sheet14!$B$2:$H$12, =4.75, etc. Read-only.
    value *Json;
    // Specifies whether the object is visible or not.
    visible *bool;
    // Returns the worksheet on which the named item is scoped to. Available only if the item is scoped to the worksheet. Read-only.
    worksheet *WorkbookWorksheet;
}
// Instantiates a new workbookNamedItem and sets the default values.
func NewWorkbookNamedItem()(*WorkbookNamedItem) {
    m := &WorkbookNamedItem{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the comment property value. Represents the comment associated with this name.
func (m *WorkbookNamedItem) GetComment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.comment
    }
}
// Gets the name property value. The name of the object. Read-only.
func (m *WorkbookNamedItem) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the scope property value. Indicates whether the name is scoped to the workbook or to a specific worksheet. Read-only.
func (m *WorkbookNamedItem) GetScope()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// Gets the type_escaped property value. Indicates what type of reference is associated with the name. The possible values are: String, Integer, Double, Boolean, Range. Read-only.
func (m *WorkbookNamedItem) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Gets the value property value. Represents the formula that the name is defined to refer to. E.g. =Sheet14!$B$2:$H$12, =4.75, etc. Read-only.
func (m *WorkbookNamedItem) GetValue()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// Gets the visible property value. Specifies whether the object is visible or not.
func (m *WorkbookNamedItem) GetVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.visible
    }
}
// Gets the worksheet property value. Returns the worksheet on which the named item is scoped to. Available only if the item is scoped to the worksheet. Read-only.
func (m *WorkbookNamedItem) GetWorksheet()(*WorkbookWorksheet) {
    if m == nil {
        return nil
    } else {
        return m.worksheet
    }
}
// The deserialization information for the current model
func (m *WorkbookNamedItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["comment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetComment(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["scope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetScope(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetValue(val.(*Json))
        return nil
    }
    res["visible"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetVisible(val)
        return nil
    }
    res["worksheet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookWorksheet() })
        if err != nil {
            return err
        }
        m.SetWorksheet(val.(*WorkbookWorksheet))
        return nil
    }
    return res
}
func (m *WorkbookNamedItem) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookNamedItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("comment", m.GetComment())
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
        err = writer.WriteStringValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type_escaped", m.GetType_escaped())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("visible", m.GetVisible())
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
// Sets the comment property value. Represents the comment associated with this name.
// Parameters:
//  - value : Value to set for the comment property.
func (m *WorkbookNamedItem) SetComment(value *string)() {
    m.comment = value
}
// Sets the name property value. The name of the object. Read-only.
// Parameters:
//  - value : Value to set for the name property.
func (m *WorkbookNamedItem) SetName(value *string)() {
    m.name = value
}
// Sets the scope property value. Indicates whether the name is scoped to the workbook or to a specific worksheet. Read-only.
// Parameters:
//  - value : Value to set for the scope property.
func (m *WorkbookNamedItem) SetScope(value *string)() {
    m.scope = value
}
// Sets the type_escaped property value. Indicates what type of reference is associated with the name. The possible values are: String, Integer, Double, Boolean, Range. Read-only.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *WorkbookNamedItem) SetType_escaped(value *string)() {
    m.type_escaped = value
}
// Sets the value property value. Represents the formula that the name is defined to refer to. E.g. =Sheet14!$B$2:$H$12, =4.75, etc. Read-only.
// Parameters:
//  - value : Value to set for the value property.
func (m *WorkbookNamedItem) SetValue(value *Json)() {
    m.value = value
}
// Sets the visible property value. Specifies whether the object is visible or not.
// Parameters:
//  - value : Value to set for the visible property.
func (m *WorkbookNamedItem) SetVisible(value *bool)() {
    m.visible = value
}
// Sets the worksheet property value. Returns the worksheet on which the named item is scoped to. Available only if the item is scoped to the worksheet. Read-only.
// Parameters:
//  - value : Value to set for the worksheet property.
func (m *WorkbookNamedItem) SetWorksheet(value *WorkbookWorksheet)() {
    m.worksheet = value
}
