package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnenoteSectionable 
type OnenoteSectionable interface {
    OnenoteEntityHierarchyModelable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetIsDefault()(*bool)
    GetLinks()(SectionLinksable)
    GetPages()([]OnenotePageable)
    GetPagesUrl()(*string)
    GetParentNotebook()(Notebookable)
    GetParentSectionGroup()(SectionGroupable)
    SetIsDefault(value *bool)()
    SetLinks(value SectionLinksable)()
    SetPages(value []OnenotePageable)()
    SetPagesUrl(value *string)()
    SetParentNotebook(value Notebookable)()
    SetParentSectionGroup(value SectionGroupable)()
}
