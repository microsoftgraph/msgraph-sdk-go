package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SectionGroupable 
type SectionGroupable interface {
    OnenoteEntityHierarchyModelable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetParentNotebook()(Notebookable)
    GetParentSectionGroup()(SectionGroupable)
    GetSectionGroups()([]SectionGroupable)
    GetSectionGroupsUrl()(*string)
    GetSections()([]OnenoteSectionable)
    GetSectionsUrl()(*string)
    SetParentNotebook(value Notebookable)()
    SetParentSectionGroup(value SectionGroupable)()
    SetSectionGroups(value []SectionGroupable)()
    SetSectionGroupsUrl(value *string)()
    SetSections(value []OnenoteSectionable)()
    SetSectionsUrl(value *string)()
}
