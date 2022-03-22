package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookNamedItem 
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
    value Jsonable;
    // Specifies whether the object is visible or not.
    visible *bool;
    // Returns the worksheet on which the named item is scoped to. Available only if the item is scoped to the worksheet. Read-only.
    worksheet WorkbookWorksheetable;
}
// NewWorkbookNamedItem instantiates a new workbookNamedItem and sets the default values.
func NewWorkbookNamedItem()(*WorkbookNamedItem) {
    m := &WorkbookNamedItem{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookNamedItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookNamedItemFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookNamedItem(), nil
}
// GetComment gets the comment property value. Represents the comment associated with this name.
func (m *WorkbookNamedItem) GetComment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.comment
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookNamedItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["comment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["scope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(Jsonable))
        }
        return nil
    }
    res["visible"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisible(val)
        }
        return nil
    }
    res["worksheet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
// GetName gets the name property value. The name of the object. Read-only.
func (m *WorkbookNamedItem) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetScope gets the scope property value. Indicates whether the name is scoped to the workbook or to a specific worksheet. Read-only.
func (m *WorkbookNamedItem) GetScope()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// GetType gets the type property value. Indicates what type of reference is associated with the name. The possible values are: String, Integer, Double, Boolean, Range. Read-only.
func (m *WorkbookNamedItem) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetValue gets the value property value. Represents the formula that the name is defined to refer to. E.g. =Sheet14!$B$2:$H$12, =4.75, etc. Read-only.
func (m *WorkbookNamedItem) GetValue()(Jsonable) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// GetVisible gets the visible property value. Specifies whether the object is visible or not.
func (m *WorkbookNamedItem) GetVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.visible
    }
}
// GetWorksheet gets the worksheet property value. Returns the worksheet on which the named item is scoped to. Available only if the item is scoped to the worksheet. Read-only.
func (m *WorkbookNamedItem) GetWorksheet()(WorkbookWorksheetable) {
    if m == nil {
        return nil
    } else {
        return m.worksheet
    }
}
// Serialize serializes information the current object
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
        err = writer.WriteStringValue("type", m.GetType())
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
// SetComment sets the comment property value. Represents the comment associated with this name.
func (m *WorkbookNamedItem) SetComment(value *string)() {
    if m != nil {
        m.comment = value
    }
}
// SetName sets the name property value. The name of the object. Read-only.
func (m *WorkbookNamedItem) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetScope sets the scope property value. Indicates whether the name is scoped to the workbook or to a specific worksheet. Read-only.
func (m *WorkbookNamedItem) SetScope(value *string)() {
    if m != nil {
        m.scope = value
    }
}
// SetType sets the type property value. Indicates what type of reference is associated with the name. The possible values are: String, Integer, Double, Boolean, Range. Read-only.
func (m *WorkbookNamedItem) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetValue sets the value property value. Represents the formula that the name is defined to refer to. E.g. =Sheet14!$B$2:$H$12, =4.75, etc. Read-only.
func (m *WorkbookNamedItem) SetValue(value Jsonable)() {
    if m != nil {
        m.value = value
    }
}
// SetVisible sets the visible property value. Specifies whether the object is visible or not.
func (m *WorkbookNamedItem) SetVisible(value *bool)() {
    if m != nil {
        m.visible = value
    }
}
// SetWorksheet sets the worksheet property value. Returns the worksheet on which the named item is scoped to. Available only if the item is scoped to the worksheet. Read-only.
func (m *WorkbookNamedItem) SetWorksheet(value WorkbookWorksheetable)() {
    if m != nil {
        m.worksheet = value
    }
}
