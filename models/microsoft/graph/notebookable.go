package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Notebookable 
type Notebookable interface {
    OnenoteEntityHierarchyModelable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetIsDefault()(*bool)
    GetIsShared()(*bool)
    GetLinks()(NotebookLinksable)
    GetSectionGroups()([]SectionGroupable)
    GetSectionGroupsUrl()(*string)
    GetSections()([]OnenoteSectionable)
    GetSectionsUrl()(*string)
    GetUserRole()(*OnenoteUserRole)
    SetIsDefault(value *bool)()
    SetIsShared(value *bool)()
    SetLinks(value NotebookLinksable)()
    SetSectionGroups(value []SectionGroupable)()
    SetSectionGroupsUrl(value *string)()
    SetSections(value []OnenoteSectionable)()
    SetSectionsUrl(value *string)()
    SetUserRole(value *OnenoteUserRole)()
}
