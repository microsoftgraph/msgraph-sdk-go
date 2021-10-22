package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Workbook struct {
    Entity
    application *WorkbookApplication;
    comments []WorkbookComment;
    functions *WorkbookFunctions;
    names []WorkbookNamedItem;
    operations []WorkbookOperation;
    tables []WorkbookTable;
    worksheets []WorkbookWorksheet;
}
func NewWorkbook()(*Workbook) {
    m := &Workbook{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Workbook) GetApplication()(*WorkbookApplication) {
    if m == nil {
        return nil
    } else {
        return m.application
    }
}
func (m *Workbook) GetComments()([]WorkbookComment) {
    if m == nil {
        return nil
    } else {
        return m.comments
    }
}
func (m *Workbook) GetFunctions()(*WorkbookFunctions) {
    if m == nil {
        return nil
    } else {
        return m.functions
    }
}
func (m *Workbook) GetNames()([]WorkbookNamedItem) {
    if m == nil {
        return nil
    } else {
        return m.names
    }
}
func (m *Workbook) GetOperations()([]WorkbookOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
func (m *Workbook) GetTables()([]WorkbookTable) {
    if m == nil {
        return nil
    } else {
        return m.tables
    }
}
func (m *Workbook) GetWorksheets()([]WorkbookWorksheet) {
    if m == nil {
        return nil
    } else {
        return m.worksheets
    }
}
func (m *Workbook) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["application"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookApplication() })
        if err != nil {
            return err
        }
        m.SetApplication(val.(*WorkbookApplication))
        return nil
    }
    res["comments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookComment() })
        if err != nil {
            return err
        }
        res := make([]WorkbookComment, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkbookComment))
        }
        m.SetComments(res)
        return nil
    }
    res["functions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookFunctions() })
        if err != nil {
            return err
        }
        m.SetFunctions(val.(*WorkbookFunctions))
        return nil
    }
    res["names"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookNamedItem() })
        if err != nil {
            return err
        }
        res := make([]WorkbookNamedItem, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkbookNamedItem))
        }
        m.SetNames(res)
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookOperation() })
        if err != nil {
            return err
        }
        res := make([]WorkbookOperation, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkbookOperation))
        }
        m.SetOperations(res)
        return nil
    }
    res["tables"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookTable() })
        if err != nil {
            return err
        }
        res := make([]WorkbookTable, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkbookTable))
        }
        m.SetTables(res)
        return nil
    }
    res["worksheets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookWorksheet() })
        if err != nil {
            return err
        }
        res := make([]WorkbookWorksheet, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkbookWorksheet))
        }
        m.SetWorksheets(res)
        return nil
    }
    return res
}
func (m *Workbook) IsNil()(bool) {
    return m == nil
}
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
func (m *Workbook) SetApplication(value *WorkbookApplication)() {
    m.application = value
}
func (m *Workbook) SetComments(value []WorkbookComment)() {
    m.comments = value
}
func (m *Workbook) SetFunctions(value *WorkbookFunctions)() {
    m.functions = value
}
func (m *Workbook) SetNames(value []WorkbookNamedItem)() {
    m.names = value
}
func (m *Workbook) SetOperations(value []WorkbookOperation)() {
    m.operations = value
}
func (m *Workbook) SetTables(value []WorkbookTable)() {
    m.tables = value
}
func (m *Workbook) SetWorksheets(value []WorkbookWorksheet)() {
    m.worksheets = value
}
