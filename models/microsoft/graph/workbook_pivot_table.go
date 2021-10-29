package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookPivotTable struct {
    Entity
    // Name of the PivotTable.
    name *string;
    // The worksheet containing the current PivotTable. Read-only.
    worksheet *WorkbookWorksheet;
}
// Instantiates a new workbookPivotTable and sets the default values.
func NewWorkbookPivotTable()(*WorkbookPivotTable) {
    m := &WorkbookPivotTable{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the name property value. Name of the PivotTable.
func (m *WorkbookPivotTable) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the worksheet property value. The worksheet containing the current PivotTable. Read-only.
func (m *WorkbookPivotTable) GetWorksheet()(*WorkbookWorksheet) {
    if m == nil {
        return nil
    } else {
        return m.worksheet
    }
}
// The deserialization information for the current model
func (m *WorkbookPivotTable) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
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
func (m *WorkbookPivotTable) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookPivotTable) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
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
// Sets the name property value. Name of the PivotTable.
// Parameters:
//  - value : Value to set for the name property.
func (m *WorkbookPivotTable) SetName(value *string)() {
    m.name = value
}
// Sets the worksheet property value. The worksheet containing the current PivotTable. Read-only.
// Parameters:
//  - value : Value to set for the worksheet property.
func (m *WorkbookPivotTable) SetWorksheet(value *WorkbookWorksheet)() {
    m.worksheet = value
}
