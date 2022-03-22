package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookPivotTable 
type WorkbookPivotTable struct {
    Entity
    // Name of the PivotTable.
    name *string;
    // The worksheet containing the current PivotTable. Read-only.
    worksheet WorkbookWorksheetable;
}
// NewWorkbookPivotTable instantiates a new workbookPivotTable and sets the default values.
func NewWorkbookPivotTable()(*WorkbookPivotTable) {
    m := &WorkbookPivotTable{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookPivotTableFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookPivotTableFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookPivotTable(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookPivotTable) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
// GetName gets the name property value. Name of the PivotTable.
func (m *WorkbookPivotTable) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetWorksheet gets the worksheet property value. The worksheet containing the current PivotTable. Read-only.
func (m *WorkbookPivotTable) GetWorksheet()(WorkbookWorksheetable) {
    if m == nil {
        return nil
    } else {
        return m.worksheet
    }
}
// Serialize serializes information the current object
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
// SetName sets the name property value. Name of the PivotTable.
func (m *WorkbookPivotTable) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetWorksheet sets the worksheet property value. The worksheet containing the current PivotTable. Read-only.
func (m *WorkbookPivotTable) SetWorksheet(value WorkbookWorksheetable)() {
    if m != nil {
        m.worksheet = value
    }
}
