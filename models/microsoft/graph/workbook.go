package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Workbook provides operations to manage the collection of drive entities.
type Workbook struct {
    Entity
    // 
    application WorkbookApplicationable;
    // 
    comments []WorkbookCommentable;
    // 
    functions WorkbookFunctionsable;
    // Represents a collection of workbooks scoped named items (named ranges and constants). Read-only.
    names []WorkbookNamedItemable;
    // The status of workbook operations. Getting an operation collection is not supported, but you can get the status of a long-running operation if the Location header is returned in the response. Read-only.
    operations []WorkbookOperationable;
    // Represents a collection of tables associated with the workbook. Read-only.
    tables []WorkbookTableable;
    // Represents a collection of worksheets associated with the workbook. Read-only.
    worksheets []WorkbookWorksheetable;
}
// NewWorkbook instantiates a new workbook and sets the default values.
func NewWorkbook()(*Workbook) {
    m := &Workbook{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbook(), nil
}
// GetApplication gets the application property value. 
func (m *Workbook) GetApplication()(WorkbookApplicationable) {
    if m == nil {
        return nil
    } else {
        return m.application
    }
}
// GetComments gets the comments property value. 
func (m *Workbook) GetComments()([]WorkbookCommentable) {
    if m == nil {
        return nil
    } else {
        return m.comments
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Workbook) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["application"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val.(WorkbookApplicationable))
        }
        return nil
    }
    res["comments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookCommentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookCommentable, len(val))
            for i, v := range val {
                res[i] = v.(WorkbookCommentable)
            }
            m.SetComments(res)
        }
        return nil
    }
    res["functions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookFunctionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFunctions(val.(WorkbookFunctionsable))
        }
        return nil
    }
    res["names"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookNamedItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookNamedItemable, len(val))
            for i, v := range val {
                res[i] = v.(WorkbookNamedItemable)
            }
            m.SetNames(res)
        }
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookOperationable, len(val))
            for i, v := range val {
                res[i] = v.(WorkbookOperationable)
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["tables"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookTableFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookTableable, len(val))
            for i, v := range val {
                res[i] = v.(WorkbookTableable)
            }
            m.SetTables(res)
        }
        return nil
    }
    res["worksheets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookWorksheetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookWorksheetable, len(val))
            for i, v := range val {
                res[i] = v.(WorkbookWorksheetable)
            }
            m.SetWorksheets(res)
        }
        return nil
    }
    return res
}
// GetFunctions gets the functions property value. 
func (m *Workbook) GetFunctions()(WorkbookFunctionsable) {
    if m == nil {
        return nil
    } else {
        return m.functions
    }
}
// GetNames gets the names property value. Represents a collection of workbooks scoped named items (named ranges and constants). Read-only.
func (m *Workbook) GetNames()([]WorkbookNamedItemable) {
    if m == nil {
        return nil
    } else {
        return m.names
    }
}
// GetOperations gets the operations property value. The status of workbook operations. Getting an operation collection is not supported, but you can get the status of a long-running operation if the Location header is returned in the response. Read-only.
func (m *Workbook) GetOperations()([]WorkbookOperationable) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// GetTables gets the tables property value. Represents a collection of tables associated with the workbook. Read-only.
func (m *Workbook) GetTables()([]WorkbookTableable) {
    if m == nil {
        return nil
    } else {
        return m.tables
    }
}
// GetWorksheets gets the worksheets property value. Represents a collection of worksheets associated with the workbook. Read-only.
func (m *Workbook) GetWorksheets()([]WorkbookWorksheetable) {
    if m == nil {
        return nil
    } else {
        return m.worksheets
    }
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
    if m.GetComments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetComments()))
        for i, v := range m.GetComments() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetNames() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNames()))
        for i, v := range m.GetNames() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("names", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTables() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTables()))
        for i, v := range m.GetTables() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tables", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorksheets() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWorksheets()))
        for i, v := range m.GetWorksheets() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("worksheets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplication sets the application property value. 
func (m *Workbook) SetApplication(value WorkbookApplicationable)() {
    if m != nil {
        m.application = value
    }
}
// SetComments sets the comments property value. 
func (m *Workbook) SetComments(value []WorkbookCommentable)() {
    if m != nil {
        m.comments = value
    }
}
// SetFunctions sets the functions property value. 
func (m *Workbook) SetFunctions(value WorkbookFunctionsable)() {
    if m != nil {
        m.functions = value
    }
}
// SetNames sets the names property value. Represents a collection of workbooks scoped named items (named ranges and constants). Read-only.
func (m *Workbook) SetNames(value []WorkbookNamedItemable)() {
    if m != nil {
        m.names = value
    }
}
// SetOperations sets the operations property value. The status of workbook operations. Getting an operation collection is not supported, but you can get the status of a long-running operation if the Location header is returned in the response. Read-only.
func (m *Workbook) SetOperations(value []WorkbookOperationable)() {
    if m != nil {
        m.operations = value
    }
}
// SetTables sets the tables property value. Represents a collection of tables associated with the workbook. Read-only.
func (m *Workbook) SetTables(value []WorkbookTableable)() {
    if m != nil {
        m.tables = value
    }
}
// SetWorksheets sets the worksheets property value. Represents a collection of worksheets associated with the workbook. Read-only.
func (m *Workbook) SetWorksheets(value []WorkbookWorksheetable)() {
    if m != nil {
        m.worksheets = value
    }
}
