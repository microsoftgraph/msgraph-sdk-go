package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Workbookable 
type Workbookable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplication()(WorkbookApplicationable)
    GetComments()([]WorkbookCommentable)
    GetFunctions()(WorkbookFunctionsable)
    GetNames()([]WorkbookNamedItemable)
    GetOperations()([]WorkbookOperationable)
    GetTables()([]WorkbookTableable)
    GetWorksheets()([]WorkbookWorksheetable)
    SetApplication(value WorkbookApplicationable)()
    SetComments(value []WorkbookCommentable)()
    SetFunctions(value WorkbookFunctionsable)()
    SetNames(value []WorkbookNamedItemable)()
    SetOperations(value []WorkbookOperationable)()
    SetTables(value []WorkbookTableable)()
    SetWorksheets(value []WorkbookWorksheetable)()
}
