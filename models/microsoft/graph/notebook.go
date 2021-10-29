package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Notebook struct {
    OnenoteEntityHierarchyModel
    // Indicates whether this is the user's default notebook. Read-only.
    isDefault *bool;
    // Indicates whether the notebook is shared. If true, the contents of the notebook can be seen by people other than the owner. Read-only.
    isShared *bool;
    // Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote native client if it's installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
    links *NotebookLinks;
    // The section groups in the notebook. Read-only. Nullable.
    sectionGroups []SectionGroup;
    // The URL for the sectionGroups navigation property, which returns all the section groups in the notebook. Read-only.
    sectionGroupsUrl *string;
    // The sections in the notebook. Read-only. Nullable.
    sections []OnenoteSection;
    // The URL for the sections navigation property, which returns all the sections in the notebook. Read-only.
    sectionsUrl *string;
    // Possible values are: Owner, Contributor, Reader, None. Owner represents owner-level access to the notebook. Contributor represents read/write access to the notebook. Reader represents read-only access to the notebook. Read-only.
    userRole *OnenoteUserRole;
}
// Instantiates a new notebook and sets the default values.
func NewNotebook()(*Notebook) {
    m := &Notebook{
        OnenoteEntityHierarchyModel: *NewOnenoteEntityHierarchyModel(),
    }
    return m
}
// Gets the isDefault property value. Indicates whether this is the user's default notebook. Read-only.
func (m *Notebook) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// Gets the isShared property value. Indicates whether the notebook is shared. If true, the contents of the notebook can be seen by people other than the owner. Read-only.
func (m *Notebook) GetIsShared()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isShared
    }
}
// Gets the links property value. Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote native client if it's installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
func (m *Notebook) GetLinks()(*NotebookLinks) {
    if m == nil {
        return nil
    } else {
        return m.links
    }
}
// Gets the sectionGroups property value. The section groups in the notebook. Read-only. Nullable.
func (m *Notebook) GetSectionGroups()([]SectionGroup) {
    if m == nil {
        return nil
    } else {
        return m.sectionGroups
    }
}
// Gets the sectionGroupsUrl property value. The URL for the sectionGroups navigation property, which returns all the section groups in the notebook. Read-only.
func (m *Notebook) GetSectionGroupsUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sectionGroupsUrl
    }
}
// Gets the sections property value. The sections in the notebook. Read-only. Nullable.
func (m *Notebook) GetSections()([]OnenoteSection) {
    if m == nil {
        return nil
    } else {
        return m.sections
    }
}
// Gets the sectionsUrl property value. The URL for the sections navigation property, which returns all the sections in the notebook. Read-only.
func (m *Notebook) GetSectionsUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sectionsUrl
    }
}
// Gets the userRole property value. Possible values are: Owner, Contributor, Reader, None. Owner represents owner-level access to the notebook. Contributor represents read/write access to the notebook. Reader represents read-only access to the notebook. Read-only.
func (m *Notebook) GetUserRole()(*OnenoteUserRole) {
    if m == nil {
        return nil
    } else {
        return m.userRole
    }
}
// The deserialization information for the current model
func (m *Notebook) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OnenoteEntityHierarchyModel.GetFieldDeserializers()
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefault(val)
        return nil
    }
    res["isShared"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsShared(val)
        return nil
    }
    res["links"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNotebookLinks() })
        if err != nil {
            return err
        }
        m.SetLinks(val.(*NotebookLinks))
        return nil
    }
    res["sectionGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSectionGroup() })
        if err != nil {
            return err
        }
        res := make([]SectionGroup, len(val))
        for i, v := range val {
            res[i] = *(v.(*SectionGroup))
        }
        m.SetSectionGroups(res)
        return nil
    }
    res["sectionGroupsUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSectionGroupsUrl(val)
        return nil
    }
    res["sections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnenoteSection() })
        if err != nil {
            return err
        }
        res := make([]OnenoteSection, len(val))
        for i, v := range val {
            res[i] = *(v.(*OnenoteSection))
        }
        m.SetSections(res)
        return nil
    }
    res["sectionsUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSectionsUrl(val)
        return nil
    }
    res["userRole"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnenoteUserRole)
        if err != nil {
            return err
        }
        cast := val.(OnenoteUserRole)
        m.SetUserRole(&cast)
        return nil
    }
    return res
}
func (m *Notebook) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Notebook) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OnenoteEntityHierarchyModel.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isShared", m.GetIsShared())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("links", m.GetLinks())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSectionGroups()))
        for i, v := range m.GetSectionGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("sectionGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sectionGroupsUrl", m.GetSectionGroupsUrl())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSections()))
        for i, v := range m.GetSections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("sections", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sectionsUrl", m.GetSectionsUrl())
        if err != nil {
            return err
        }
    }
    if m.GetUserRole() != nil {
        cast := m.GetUserRole().String()
        err = writer.WriteStringValue("userRole", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the isDefault property value. Indicates whether this is the user's default notebook. Read-only.
// Parameters:
//  - value : Value to set for the isDefault property.
func (m *Notebook) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// Sets the isShared property value. Indicates whether the notebook is shared. If true, the contents of the notebook can be seen by people other than the owner. Read-only.
// Parameters:
//  - value : Value to set for the isShared property.
func (m *Notebook) SetIsShared(value *bool)() {
    m.isShared = value
}
// Sets the links property value. Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote native client if it's installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
// Parameters:
//  - value : Value to set for the links property.
func (m *Notebook) SetLinks(value *NotebookLinks)() {
    m.links = value
}
// Sets the sectionGroups property value. The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the sectionGroups property.
func (m *Notebook) SetSectionGroups(value []SectionGroup)() {
    m.sectionGroups = value
}
// Sets the sectionGroupsUrl property value. The URL for the sectionGroups navigation property, which returns all the section groups in the notebook. Read-only.
// Parameters:
//  - value : Value to set for the sectionGroupsUrl property.
func (m *Notebook) SetSectionGroupsUrl(value *string)() {
    m.sectionGroupsUrl = value
}
// Sets the sections property value. The sections in the notebook. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the sections property.
func (m *Notebook) SetSections(value []OnenoteSection)() {
    m.sections = value
}
// Sets the sectionsUrl property value. The URL for the sections navigation property, which returns all the sections in the notebook. Read-only.
// Parameters:
//  - value : Value to set for the sectionsUrl property.
func (m *Notebook) SetSectionsUrl(value *string)() {
    m.sectionsUrl = value
}
// Sets the userRole property value. Possible values are: Owner, Contributor, Reader, None. Owner represents owner-level access to the notebook. Contributor represents read/write access to the notebook. Reader represents read-only access to the notebook. Read-only.
// Parameters:
//  - value : Value to set for the userRole property.
func (m *Notebook) SetUserRole(value *OnenoteUserRole)() {
    m.userRole = value
}
