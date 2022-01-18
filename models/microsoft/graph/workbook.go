package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Workbook 
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
    // The status of Workbook operations. Getting an operation collection is not supported, but you can get the status of a long-running operation if the Location header is returned in the response. Read-only. Nullable.
    operations []WorkbookOperation;
    // Represents a collection of tables associated with the workbook. Read-only.
    tables []WorkbookTable;
    // Represents a collection of worksheets associated with the workbook. Read-only.
    worksheets []WorkbookWorksheet;
}
// NewWorkbook instantiates a new workbook and sets the default values.
func NewWorkbook()(*Workbook) {
    m := &Workbook{
        Entity: *NewEntity(),
    }
    return m
}
// GetApplication gets the application property value. 
func (m *Workbook) GetApplication()(*WorkbookApplication) {
    if m == nil {
        return nil
    } else {
        return m.application
    }
}
// GetComments gets the comments property value. 
func (m *Workbook) GetComments()([]WorkbookComment) {
    if m == nil {
        return nil
    } else {
        return m.comments
    }
}
// GetFunctions gets the functions property value. 
func (m *Workbook) GetFunctions()(*WorkbookFunctions) {
    if m == nil {
        return nil
    } else {
        return m.functions
    }
}
// GetNames gets the names property value. Represents a collection of workbooks scoped named items (named ranges and constants). Read-only.
func (m *Workbook) GetNames()([]WorkbookNamedItem) {
    if m == nil {
        return nil
    } else {
        return m.names
    }
}
// GetOperations gets the operations property value. The status of Workbook operations. Getting an operation collection is not supported, but you can get the status of a long-running operation if the Location header is returned in the response. Read-only. Nullable.
func (m *Workbook) GetOperations()([]WorkbookOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// GetTables gets the tables property value. Represents a collection of tables associated with the workbook. Read-only.
func (m *Workbook) GetTables()([]WorkbookTable) {
    if m == nil {
        return nil
    } else {
        return m.tables
    }
}
// GetWorksheets gets the worksheets property value. Represents a collection of worksheets associated with the workbook. Read-only.
func (m *Workbook) GetWorksheets()([]WorkbookWorksheet) {
    if m == nil {
        return nil
    } else {
        return m.worksheets
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetApplication sets the application property value. 
func (m *Workbook) SetApplication(value *WorkbookApplication)() {
    if m != nil {
        m.application = value
    }
}
// SetComments sets the comments property value. 
func (m *Workbook) SetComments(value []WorkbookComment)() {
    if m != nil {
        m.comments = value
    }
}
// SetFunctions sets the functions property value. 
func (m *Workbook) SetFunctions(value *WorkbookFunctions)() {
    if m != nil {
        m.functions = value
    }
}
// SetNames sets the names property value. Represents a collection of workbooks scoped named items (named ranges and constants). Read-only.
func (m *Workbook) SetNames(value []WorkbookNamedItem)() {
    if m != nil {
        m.names = value
    }
}
// SetOperations sets the operations property value. The status of Workbook operations. Getting an operation collection is not supported, but you can get the status of a long-running operation if the Location header is returned in the response. Read-only. Nullable.
func (m *Workbook) SetOperations(value []WorkbookOperation)() {
    if m != nil {
        m.operations = value
    }
}
// SetTables sets the tables property value. Represents a collection of tables associated with the workbook. Read-only.
func (m *Workbook) SetTables(value []WorkbookTable)() {
    if m != nil {
        m.tables = value
    }
}
// SetWorksheets sets the worksheets property value. Represents a collection of worksheets associated with the workbook. Read-only.
func (m *Workbook) SetWorksheets(value []WorkbookWorksheet)() {
    if m != nil {
        m.worksheets = value
    }
}
