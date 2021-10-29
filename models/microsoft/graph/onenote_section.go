package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OnenoteSection struct {
    OnenoteEntityHierarchyModel
    // Indicates whether this is the user's default section. Read-only.
    isDefault *bool;
    // Links for opening the section. The oneNoteClientURL link opens the section in the OneNote native client if it's installed. The oneNoteWebURL link opens the section in OneNote on the web.
    links *SectionLinks;
    // The collection of pages in the section.  Read-only. Nullable.
    pages []OnenotePage;
    // The pages endpoint where you can get details for all the pages in the section. Read-only.
    pagesUrl *string;
    // The notebook that contains the section.  Read-only.
    parentNotebook *Notebook;
    // The section group that contains the section.  Read-only.
    parentSectionGroup *SectionGroup;
}
// Instantiates a new onenoteSection and sets the default values.
func NewOnenoteSection()(*OnenoteSection) {
    m := &OnenoteSection{
        OnenoteEntityHierarchyModel: *NewOnenoteEntityHierarchyModel(),
    }
    return m
}
// Gets the isDefault property value. Indicates whether this is the user's default section. Read-only.
func (m *OnenoteSection) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// Gets the links property value. Links for opening the section. The oneNoteClientURL link opens the section in the OneNote native client if it's installed. The oneNoteWebURL link opens the section in OneNote on the web.
func (m *OnenoteSection) GetLinks()(*SectionLinks) {
    if m == nil {
        return nil
    } else {
        return m.links
    }
}
// Gets the pages property value. The collection of pages in the section.  Read-only. Nullable.
func (m *OnenoteSection) GetPages()([]OnenotePage) {
    if m == nil {
        return nil
    } else {
        return m.pages
    }
}
// Gets the pagesUrl property value. The pages endpoint where you can get details for all the pages in the section. Read-only.
func (m *OnenoteSection) GetPagesUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.pagesUrl
    }
}
// Gets the parentNotebook property value. The notebook that contains the section.  Read-only.
func (m *OnenoteSection) GetParentNotebook()(*Notebook) {
    if m == nil {
        return nil
    } else {
        return m.parentNotebook
    }
}
// Gets the parentSectionGroup property value. The section group that contains the section.  Read-only.
func (m *OnenoteSection) GetParentSectionGroup()(*SectionGroup) {
    if m == nil {
        return nil
    } else {
        return m.parentSectionGroup
    }
}
// The deserialization information for the current model
func (m *OnenoteSection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OnenoteEntityHierarchyModel.GetFieldDeserializers()
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefault(val)
        return nil
    }
    res["links"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSectionLinks() })
        if err != nil {
            return err
        }
        m.SetLinks(val.(*SectionLinks))
        return nil
    }
    res["pages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnenotePage() })
        if err != nil {
            return err
        }
        res := make([]OnenotePage, len(val))
        for i, v := range val {
            res[i] = *(v.(*OnenotePage))
        }
        m.SetPages(res)
        return nil
    }
    res["pagesUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPagesUrl(val)
        return nil
    }
    res["parentNotebook"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNotebook() })
        if err != nil {
            return err
        }
        m.SetParentNotebook(val.(*Notebook))
        return nil
    }
    res["parentSectionGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSectionGroup() })
        if err != nil {
            return err
        }
        m.SetParentSectionGroup(val.(*SectionGroup))
        return nil
    }
    return res
}
func (m *OnenoteSection) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OnenoteSection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("links", m.GetLinks())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPages()))
        for i, v := range m.GetPages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("pages", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("pagesUrl", m.GetPagesUrl())
        if err != nil {
            return err
        }
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
    return nil
}
// Sets the isDefault property value. Indicates whether this is the user's default section. Read-only.
// Parameters:
//  - value : Value to set for the isDefault property.
func (m *OnenoteSection) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// Sets the links property value. Links for opening the section. The oneNoteClientURL link opens the section in the OneNote native client if it's installed. The oneNoteWebURL link opens the section in OneNote on the web.
// Parameters:
//  - value : Value to set for the links property.
func (m *OnenoteSection) SetLinks(value *SectionLinks)() {
    m.links = value
}
// Sets the pages property value. The collection of pages in the section.  Read-only. Nullable.
// Parameters:
//  - value : Value to set for the pages property.
func (m *OnenoteSection) SetPages(value []OnenotePage)() {
    m.pages = value
}
// Sets the pagesUrl property value. The pages endpoint where you can get details for all the pages in the section. Read-only.
// Parameters:
//  - value : Value to set for the pagesUrl property.
func (m *OnenoteSection) SetPagesUrl(value *string)() {
    m.pagesUrl = value
}
// Sets the parentNotebook property value. The notebook that contains the section.  Read-only.
// Parameters:
//  - value : Value to set for the parentNotebook property.
func (m *OnenoteSection) SetParentNotebook(value *Notebook)() {
    m.parentNotebook = value
}
// Sets the parentSectionGroup property value. The section group that contains the section.  Read-only.
// Parameters:
//  - value : Value to set for the parentSectionGroup property.
func (m *OnenoteSection) SetParentSectionGroup(value *SectionGroup)() {
    m.parentSectionGroup = value
}
