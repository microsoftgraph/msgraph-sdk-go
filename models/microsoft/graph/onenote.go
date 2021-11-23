package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// onenote 
type Onenote struct {
    Entity
    // The collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
    notebooks []Notebook;
    // The status of OneNote operations. Getting an operations collection is not supported, but you can get the status of long-running operations if the Operation-Location header is returned in the response. Read-only. Nullable.
    operations []OnenoteOperation;
    // The pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
    pages []OnenotePage;
    // The image and other file resources in OneNote pages. Getting a resources collection is not supported, but you can get the binary content of a specific resource. Read-only. Nullable.
    resources []OnenoteResource;
    // The section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
    sectionGroups []SectionGroup;
    // The sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
    sections []OnenoteSection;
}
// NewOnenote instantiates a new onenote and sets the default values.
func NewOnenote()(*Onenote) {
    m := &Onenote{
        Entity: *NewEntity(),
    }
    return m
}
// GetNotebooks gets the notebooks property value. The collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
func (m *Onenote) GetNotebooks()([]Notebook) {
    if m == nil {
        return nil
    } else {
        return m.notebooks
    }
}
// GetOperations gets the operations property value. The status of OneNote operations. Getting an operations collection is not supported, but you can get the status of long-running operations if the Operation-Location header is returned in the response. Read-only. Nullable.
func (m *Onenote) GetOperations()([]OnenoteOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// GetPages gets the pages property value. The pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *Onenote) GetPages()([]OnenotePage) {
    if m == nil {
        return nil
    } else {
        return m.pages
    }
}
// GetResources gets the resources property value. The image and other file resources in OneNote pages. Getting a resources collection is not supported, but you can get the binary content of a specific resource. Read-only. Nullable.
func (m *Onenote) GetResources()([]OnenoteResource) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
// GetSectionGroups gets the sectionGroups property value. The section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *Onenote) GetSectionGroups()([]SectionGroup) {
    if m == nil {
        return nil
    } else {
        return m.sectionGroups
    }
}
// GetSections gets the sections property value. The sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *Onenote) GetSections()([]OnenoteSection) {
    if m == nil {
        return nil
    } else {
        return m.sections
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Onenote) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["notebooks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNotebook() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Notebook, len(val))
            for i, v := range val {
                res[i] = *(v.(*Notebook))
            }
            m.SetNotebooks(res)
        }
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnenoteOperation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnenoteOperation, len(val))
            for i, v := range val {
                res[i] = *(v.(*OnenoteOperation))
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["pages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnenotePage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnenotePage, len(val))
            for i, v := range val {
                res[i] = *(v.(*OnenotePage))
            }
            m.SetPages(res)
        }
        return nil
    }
    res["resources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnenoteResource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnenoteResource, len(val))
            for i, v := range val {
                res[i] = *(v.(*OnenoteResource))
            }
            m.SetResources(res)
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
    return res
}
func (m *Onenote) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Onenote) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNotebooks()))
        for i, v := range m.GetNotebooks() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("notebooks", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("resources", cast)
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
    return nil
}
// SetNotebooks sets the notebooks property value. The collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
func (m *Onenote) SetNotebooks(value []Notebook)() {
    m.notebooks = value
}
// SetOperations sets the operations property value. The status of OneNote operations. Getting an operations collection is not supported, but you can get the status of long-running operations if the Operation-Location header is returned in the response. Read-only. Nullable.
func (m *Onenote) SetOperations(value []OnenoteOperation)() {
    m.operations = value
}
// SetPages sets the pages property value. The pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *Onenote) SetPages(value []OnenotePage)() {
    m.pages = value
}
// SetResources sets the resources property value. The image and other file resources in OneNote pages. Getting a resources collection is not supported, but you can get the binary content of a specific resource. Read-only. Nullable.
func (m *Onenote) SetResources(value []OnenoteResource)() {
    m.resources = value
}
// SetSectionGroups sets the sectionGroups property value. The section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *Onenote) SetSectionGroups(value []SectionGroup)() {
    m.sectionGroups = value
}
// SetSections sets the sections property value. The sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *Onenote) SetSections(value []OnenoteSection)() {
    m.sections = value
}
