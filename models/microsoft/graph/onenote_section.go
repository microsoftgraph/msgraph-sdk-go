package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnenoteSection provides operations to manage the drive singleton.
type OnenoteSection struct {
    OnenoteEntityHierarchyModel
    // Indicates whether this is the user's default section. Read-only.
    isDefault *bool;
    // Links for opening the section. The oneNoteClientURL link opens the section in the OneNote native client if it's installed. The oneNoteWebURL link opens the section in OneNote on the web.
    links SectionLinksable;
    // The collection of pages in the section.  Read-only. Nullable.
    pages []OnenotePageable;
    // The pages endpoint where you can get details for all the pages in the section. Read-only.
    pagesUrl *string;
    // The notebook that contains the section.  Read-only.
    parentNotebook Notebookable;
    // The section group that contains the section.  Read-only.
    parentSectionGroup SectionGroupable;
}
// NewOnenoteSection instantiates a new onenoteSection and sets the default values.
func NewOnenoteSection()(*OnenoteSection) {
    m := &OnenoteSection{
        OnenoteEntityHierarchyModel: *NewOnenoteEntityHierarchyModel(),
    }
    return m
}
// CreateOnenoteSectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenoteSectionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOnenoteSection(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenoteSection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["links"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateSectionLinksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(SectionLinksable))
        }
        return nil
    }
    res["pages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnenotePageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnenotePageable, len(val))
            for i, v := range val {
                res[i] = v.(OnenotePageable)
            }
            m.SetPages(res)
        }
        return nil
    }
    res["pagesUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPagesUrl(val)
        }
        return nil
    }
    res["parentNotebook"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateNotebookFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentNotebook(val.(Notebookable))
        }
        return nil
    }
    res["parentSectionGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateSectionGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentSectionGroup(val.(SectionGroupable))
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. Indicates whether this is the user's default section. Read-only.
func (m *OnenoteSection) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// GetLinks gets the links property value. Links for opening the section. The oneNoteClientURL link opens the section in the OneNote native client if it's installed. The oneNoteWebURL link opens the section in OneNote on the web.
func (m *OnenoteSection) GetLinks()(SectionLinksable) {
    if m == nil {
        return nil
    } else {
        return m.links
    }
}
// GetPages gets the pages property value. The collection of pages in the section.  Read-only. Nullable.
func (m *OnenoteSection) GetPages()([]OnenotePageable) {
    if m == nil {
        return nil
    } else {
        return m.pages
    }
}
// GetPagesUrl gets the pagesUrl property value. The pages endpoint where you can get details for all the pages in the section. Read-only.
func (m *OnenoteSection) GetPagesUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.pagesUrl
    }
}
// GetParentNotebook gets the parentNotebook property value. The notebook that contains the section.  Read-only.
func (m *OnenoteSection) GetParentNotebook()(Notebookable) {
    if m == nil {
        return nil
    } else {
        return m.parentNotebook
    }
}
// GetParentSectionGroup gets the parentSectionGroup property value. The section group that contains the section.  Read-only.
func (m *OnenoteSection) GetParentSectionGroup()(SectionGroupable) {
    if m == nil {
        return nil
    } else {
        return m.parentSectionGroup
    }
}
func (m *OnenoteSection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetPages() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPages()))
        for i, v := range m.GetPages() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetIsDefault sets the isDefault property value. Indicates whether this is the user's default section. Read-only.
func (m *OnenoteSection) SetIsDefault(value *bool)() {
    if m != nil {
        m.isDefault = value
    }
}
// SetLinks sets the links property value. Links for opening the section. The oneNoteClientURL link opens the section in the OneNote native client if it's installed. The oneNoteWebURL link opens the section in OneNote on the web.
func (m *OnenoteSection) SetLinks(value SectionLinksable)() {
    if m != nil {
        m.links = value
    }
}
// SetPages sets the pages property value. The collection of pages in the section.  Read-only. Nullable.
func (m *OnenoteSection) SetPages(value []OnenotePageable)() {
    if m != nil {
        m.pages = value
    }
}
// SetPagesUrl sets the pagesUrl property value. The pages endpoint where you can get details for all the pages in the section. Read-only.
func (m *OnenoteSection) SetPagesUrl(value *string)() {
    if m != nil {
        m.pagesUrl = value
    }
}
// SetParentNotebook sets the parentNotebook property value. The notebook that contains the section.  Read-only.
func (m *OnenoteSection) SetParentNotebook(value Notebookable)() {
    if m != nil {
        m.parentNotebook = value
    }
}
// SetParentSectionGroup sets the parentSectionGroup property value. The section group that contains the section.  Read-only.
func (m *OnenoteSection) SetParentSectionGroup(value SectionGroupable)() {
    if m != nil {
        m.parentSectionGroup = value
    }
}
