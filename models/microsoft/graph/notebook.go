package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Notebook provides operations to manage the educationRoot singleton.
type Notebook struct {
    OnenoteEntityHierarchyModel
    // Indicates whether this is the user's default notebook. Read-only.
    isDefault *bool;
    // Indicates whether the notebook is shared. If true, the contents of the notebook can be seen by people other than the owner. Read-only.
    isShared *bool;
    // Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote native client if it's installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
    links NotebookLinksable;
    // The section groups in the notebook. Read-only. Nullable.
    sectionGroups []SectionGroupable;
    // The URL for the sectionGroups navigation property, which returns all the section groups in the notebook. Read-only.
    sectionGroupsUrl *string;
    // The sections in the notebook. Read-only. Nullable.
    sections []OnenoteSectionable;
    // The URL for the sections navigation property, which returns all the sections in the notebook. Read-only.
    sectionsUrl *string;
    // Possible values are: Owner, Contributor, Reader, None. Owner represents owner-level access to the notebook. Contributor represents read/write access to the notebook. Reader represents read-only access to the notebook. Read-only.
    userRole *OnenoteUserRole;
}
// NewNotebook instantiates a new notebook and sets the default values.
func NewNotebook()(*Notebook) {
    m := &Notebook{
        OnenoteEntityHierarchyModel: *NewOnenoteEntityHierarchyModel(),
    }
    return m
}
// CreateNotebookFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNotebookFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewNotebook(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Notebook) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OnenoteEntityHierarchyModel.GetFieldDeserializers()
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["isShared"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsShared(val)
        }
        return nil
    }
    res["links"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateNotebookLinksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(NotebookLinksable))
        }
        return nil
    }
    res["sectionGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSectionGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SectionGroupable, len(val))
            for i, v := range val {
                res[i] = v.(SectionGroupable)
            }
            m.SetSectionGroups(res)
        }
        return nil
    }
    res["sectionGroupsUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSectionGroupsUrl(val)
        }
        return nil
    }
    res["sections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnenoteSectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnenoteSectionable, len(val))
            for i, v := range val {
                res[i] = v.(OnenoteSectionable)
            }
            m.SetSections(res)
        }
        return nil
    }
    res["sectionsUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSectionsUrl(val)
        }
        return nil
    }
    res["userRole"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnenoteUserRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRole(val.(*OnenoteUserRole))
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. Indicates whether this is the user's default notebook. Read-only.
func (m *Notebook) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// GetIsShared gets the isShared property value. Indicates whether the notebook is shared. If true, the contents of the notebook can be seen by people other than the owner. Read-only.
func (m *Notebook) GetIsShared()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isShared
    }
}
// GetLinks gets the links property value. Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote native client if it's installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
func (m *Notebook) GetLinks()(NotebookLinksable) {
    if m == nil {
        return nil
    } else {
        return m.links
    }
}
// GetSectionGroups gets the sectionGroups property value. The section groups in the notebook. Read-only. Nullable.
func (m *Notebook) GetSectionGroups()([]SectionGroupable) {
    if m == nil {
        return nil
    } else {
        return m.sectionGroups
    }
}
// GetSectionGroupsUrl gets the sectionGroupsUrl property value. The URL for the sectionGroups navigation property, which returns all the section groups in the notebook. Read-only.
func (m *Notebook) GetSectionGroupsUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sectionGroupsUrl
    }
}
// GetSections gets the sections property value. The sections in the notebook. Read-only. Nullable.
func (m *Notebook) GetSections()([]OnenoteSectionable) {
    if m == nil {
        return nil
    } else {
        return m.sections
    }
}
// GetSectionsUrl gets the sectionsUrl property value. The URL for the sections navigation property, which returns all the sections in the notebook. Read-only.
func (m *Notebook) GetSectionsUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sectionsUrl
    }
}
// GetUserRole gets the userRole property value. Possible values are: Owner, Contributor, Reader, None. Owner represents owner-level access to the notebook. Contributor represents read/write access to the notebook. Reader represents read-only access to the notebook. Read-only.
func (m *Notebook) GetUserRole()(*OnenoteUserRole) {
    if m == nil {
        return nil
    } else {
        return m.userRole
    }
}
func (m *Notebook) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetSectionGroups() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSectionGroups()))
        for i, v := range m.GetSectionGroups() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetSections() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSections()))
        for i, v := range m.GetSections() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
        cast := (*m.GetUserRole()).String()
        err = writer.WriteStringValue("userRole", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsDefault sets the isDefault property value. Indicates whether this is the user's default notebook. Read-only.
func (m *Notebook) SetIsDefault(value *bool)() {
    if m != nil {
        m.isDefault = value
    }
}
// SetIsShared sets the isShared property value. Indicates whether the notebook is shared. If true, the contents of the notebook can be seen by people other than the owner. Read-only.
func (m *Notebook) SetIsShared(value *bool)() {
    if m != nil {
        m.isShared = value
    }
}
// SetLinks sets the links property value. Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote native client if it's installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
func (m *Notebook) SetLinks(value NotebookLinksable)() {
    if m != nil {
        m.links = value
    }
}
// SetSectionGroups sets the sectionGroups property value. The section groups in the notebook. Read-only. Nullable.
func (m *Notebook) SetSectionGroups(value []SectionGroupable)() {
    if m != nil {
        m.sectionGroups = value
    }
}
// SetSectionGroupsUrl sets the sectionGroupsUrl property value. The URL for the sectionGroups navigation property, which returns all the section groups in the notebook. Read-only.
func (m *Notebook) SetSectionGroupsUrl(value *string)() {
    if m != nil {
        m.sectionGroupsUrl = value
    }
}
// SetSections sets the sections property value. The sections in the notebook. Read-only. Nullable.
func (m *Notebook) SetSections(value []OnenoteSectionable)() {
    if m != nil {
        m.sections = value
    }
}
// SetSectionsUrl sets the sectionsUrl property value. The URL for the sections navigation property, which returns all the sections in the notebook. Read-only.
func (m *Notebook) SetSectionsUrl(value *string)() {
    if m != nil {
        m.sectionsUrl = value
    }
}
// SetUserRole sets the userRole property value. Possible values are: Owner, Contributor, Reader, None. Owner represents owner-level access to the notebook. Contributor represents read/write access to the notebook. Reader represents read-only access to the notebook. Read-only.
func (m *Notebook) SetUserRole(value *OnenoteUserRole)() {
    if m != nil {
        m.userRole = value
    }
}
