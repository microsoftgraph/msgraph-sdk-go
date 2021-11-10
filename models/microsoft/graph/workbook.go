package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Workbook struct {
    Entity
    // 
    application *WorkbookApplication;
    // 
    comments []WorkbookComment;
    // 
    functions *WorkbookFunctions;
    // Represents a collection of workbooks scoped named items (named ranges and constants). Read-only.
    names []WorkbookNamedItem;
    // The status of workbook operations. Getting an operation collection is not supported, but you can get the status of a long-running operation if the Location header is returned in the response. Read-only.
    operations []WorkbookOperation;
    // Represents a collection of tables associated with the workbook. Read-only.
    tables []WorkbookTable;
    // Represents a collection of worksheets associated with the workbook. Read-only.
    worksheets []WorkbookWorksheet;
}
// Instantiates a new workbook and sets the default values.
func NewWorkbook()(*Workbook) {
    m := &Workbook{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the application property value. 
func (m *Workbook) GetApplication()(*WorkbookApplication) {
    if m == nil {
        return nil
    } else {
        return m.application
    }
}
// Gets the comments property value. 
func (m *Workbook) GetComments()([]WorkbookComment) {
    if m == nil {
        return nil
    } else {
        return m.comments
    }
}
// Gets the functions property value. 
func (m *Workbook) GetFunctions()(*WorkbookFunctions) {
    if m == nil {
        return nil
    } else {
        return m.functions
    }
}
// Gets the names property value. Represents a collection of workbooks scoped named items (named ranges and constants). Read-only.
func (m *Workbook) GetNames()([]WorkbookNamedItem) {
    if m == nil {
        return nil
    } else {
        return m.names
    }
}
// Gets the operations property value. The status of workbook operations. Getting an operation collection is not supported, but you can get the status of a long-running operation if the Location header is returned in the response. Read-only.
func (m *Workbook) GetOperations()([]WorkbookOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// Gets the tables property value. Represents a collection of tables associated with the workbook. Read-only.
func (m *Workbook) GetTables()([]WorkbookTable) {
    if m == nil {
        return nil
    } else {
        return m.tables
    }
}
// Gets the worksheets property value. Represents a collection of worksheets associated with the workbook. Read-only.
func (m *Workbook) GetWorksheets()([]WorkbookWorksheet) {
    if m == nil {
        return nil
    } else {
        return m.worksheets
    }
}
// The deserialization information for the current model
func (m *Workbook) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["application"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookApplication() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val.(*WorkbookApplication))
        }
        return nil
    }
    res["comments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookComment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookComment, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkbookComment))
            }
            m.SetComments(res)
        }
        return nil
    }
    res["functions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookFunctions() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFunctions(val.(*WorkbookFunctions))
        }
        return nil
    }
    res["names"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookNamedItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookNamedItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkbookNamedItem))
            }
            m.SetNames(res)
        }
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookOperation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookOperation, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkbookOperation))
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["tables"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookTable() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookTable, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkbookTable))
            }
            m.SetTables(res)
        }
        return nil
    }
    res["worksheets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookWorksheet() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookWorksheet, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkbookWorksheet))
            }
            m.SetWorksheets(res)
        }
        return nil
    }
    return res
}
func (m *Workbook) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Workbook) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetComments()))
        for i, v := range m.GetComments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("comments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("functions", m.GetFunctions())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNames()))
        for i, v := range m.GetNames() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("names", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTables()))
        for i, v := range m.GetTables() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tables", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWorksheets()))
        for i, v := range m.GetWorksheets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("worksheets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the application property value. 
// Parameters:
//  - value : Value to set for the application property.
func (m *Workbook) SetApplication(value *WorkbookApplication)() {
    m.application = value
}
// Sets the comments property value. 
// Parameters:
//  - value : Value to set for the comments property.
func (m *Workbook) SetComments(value []WorkbookComment)() {
    m.comments = value
}
// Sets the functions property value. 
// Parameters:
//  - value : Value to set for the functions property.
func (m *Workbook) SetFunctions(value *WorkbookFunctions)() {
    m.functions = value
}
// Sets the names property value. Represents a collection of workbooks scoped named items (named ranges and constants). Read-only.
// Parameters:
//  - value : Value to set for the names property.
func (m *Workbook) SetNames(value []WorkbookNamedItem)() {
    m.names = value
}
// Sets the operations property value. The status of workbook operations. Getting an operation collection is not supported, but you can get the status of a long-running operation if the Location header is returned in the response. Read-only.
// Parameters:
//  - value : Value to set for the operations property.
func (m *Workbook) SetOperations(value []WorkbookOperation)() {
    m.operations = value
}
// Sets the tables property value. Represents a collection of tables associated with the workbook. Read-only.
// Parameters:
//  - value : Value to set for the tables property.
func (m *Workbook) SetTables(value []WorkbookTable)() {
    m.tables = value
}
// Sets the worksheets property value. Represents a collection of worksheets associated with the workbook. Read-only.
// Parameters:
//  - value : Value to set for the worksheets property.
func (m *Workbook) SetWorksheets(value []WorkbookWorksheet)() {
    m.worksheets = value
}
