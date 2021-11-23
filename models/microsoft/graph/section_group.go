package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SectionGroup 
type SectionGroup struct {
    OnenoteEntityHierarchyModel
    // The notebook that contains the section group. Read-only.
    parentNotebook *Notebook;
    // The section group that contains the section group. Read-only.
    parentSectionGroup *SectionGroup;
    // The section groups in the section. Read-only. Nullable.
    sectionGroups []SectionGroup;
    // The URL for the sectionGroups navigation property, which returns all the section groups in the section group. Read-only.
    sectionGroupsUrl *string;
    // The sections in the section group. Read-only. Nullable.
    sections []OnenoteSection;
    // The URL for the sections navigation property, which returns all the sections in the section group. Read-only.
    sectionsUrl *string;
}
// NewSectionGroup instantiates a new sectionGroup and sets the default values.
func NewSectionGroup()(*SectionGroup) {
    m := &SectionGroup{
        OnenoteEntityHierarchyModel: *NewOnenoteEntityHierarchyModel(),
    }
    return m
}
// GetParentNotebook gets the parentNotebook property value. The notebook that contains the section group. Read-only.
func (m *SectionGroup) GetParentNotebook()(*Notebook) {
    if m == nil {
        return nil
    } else {
        return m.parentNotebook
    }
}
// GetParentSectionGroup gets the parentSectionGroup property value. The section group that contains the section group. Read-only.
func (m *SectionGroup) GetParentSectionGroup()(*SectionGroup) {
    if m == nil {
        return nil
    } else {
        return m.parentSectionGroup
    }
}
// GetSectionGroups gets the sectionGroups property value. The section groups in the section. Read-only. Nullable.
func (m *SectionGroup) GetSectionGroups()([]SectionGroup) {
    if m == nil {
        return nil
    } else {
        return m.sectionGroups
    }
}
// GetSectionGroupsUrl gets the sectionGroupsUrl property value. The URL for the sectionGroups navigation property, which returns all the section groups in the section group. Read-only.
func (m *SectionGroup) GetSectionGroupsUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sectionGroupsUrl
    }
}
// GetSections gets the sections property value. The sections in the section group. Read-only. Nullable.
func (m *SectionGroup) GetSections()([]OnenoteSection) {
    if m == nil {
        return nil
    } else {
        return m.sections
    }
}
// GetSectionsUrl gets the sectionsUrl property value. The URL for the sections navigation property, which returns all the sections in the section group. Read-only.
func (m *SectionGroup) GetSectionsUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sectionsUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SectionGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OnenoteEntityHierarchyModel.GetFieldDeserializers()
    res["parentNotebook"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNotebook() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentNotebook(val.(*Notebook))
        }
        return nil
    }
    res["parentSectionGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSectionGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentSectionGroup(val.(*SectionGroup))
        }
        return nil
    }
    res["sectionGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSectionGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SectionGroup, len(val))
            for i, v := range val {
                res[i] = *(v.(*SectionGroup))
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
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnenoteSection() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnenoteSection, len(val))
            for i, v := range val {
                res[i] = *(v.(*OnenoteSection))
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
    return res
}
func (m *SectionGroup) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SectionGroup) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OnenoteEntityHierarchyModel.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("parentNotebook", m.GetParentNotebook())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parentSectionGroup", m.GetParentSectionGroup())
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
    return nil
}
// SetParentNotebook sets the parentNotebook property value. The notebook that contains the section group. Read-only.
func (m *SectionGroup) SetParentNotebook(value *Notebook)() {
    m.parentNotebook = value
}
// SetParentSectionGroup sets the parentSectionGroup property value. The section group that contains the section group. Read-only.
func (m *SectionGroup) SetParentSectionGroup(value *SectionGroup)() {
    m.parentSectionGroup = value
}
// SetSectionGroups sets the sectionGroups property value. The section groups in the section. Read-only. Nullable.
func (m *SectionGroup) SetSectionGroups(value []SectionGroup)() {
    m.sectionGroups = value
}
// SetSectionGroupsUrl sets the sectionGroupsUrl property value. The URL for the sectionGroups navigation property, which returns all the section groups in the section group. Read-only.
func (m *SectionGroup) SetSectionGroupsUrl(value *string)() {
    m.sectionGroupsUrl = value
}
// SetSections sets the sections property value. The sections in the section group. Read-only. Nullable.
func (m *SectionGroup) SetSections(value []OnenoteSection)() {
    m.sections = value
}
// SetSectionsUrl sets the sectionsUrl property value. The URL for the sections navigation property, which returns all the sections in the section group. Read-only.
func (m *SectionGroup) SetSectionsUrl(value *string)() {
    m.sectionsUrl = value
}
