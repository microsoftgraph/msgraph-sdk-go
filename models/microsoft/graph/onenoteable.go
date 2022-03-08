package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Onenoteable 
type Onenoteable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetNotebooks()([]Notebookable)
    GetOperations()([]OnenoteOperationable)
    GetPages()([]OnenotePageable)
    GetResources()([]OnenoteResourceable)
    GetSectionGroups()([]SectionGroupable)
    GetSections()([]OnenoteSectionable)
    SetNotebooks(value []Notebookable)()
    SetOperations(value []OnenoteOperationable)()
    SetPages(value []OnenotePageable)()
    SetResources(value []OnenoteResourceable)()
    SetSectionGroups(value []SectionGroupable)()
    SetSections(value []OnenoteSectionable)()
}
