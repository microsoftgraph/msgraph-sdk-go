package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OnenotePage struct {
    OnenoteEntitySchemaObjectModel
    // The page's HTML content.
    content []byte;
    // The URL for the page's HTML content.  Read-only.
    contentUrl *string;
    // The unique identifier of the application that created the page. Read-only.
    createdByAppId *string;
    // The date and time when the page was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The indentation level of the page. Read-only.
    level *int32;
    // Links for opening the page. The oneNoteClientURL link opens the page in the OneNote native client if it 's installed. The oneNoteWebUrl link opens the page in OneNote on the web. Read-only.
    links *PageLinks;
    // The order of the page within its parent section. Read-only.
    order *int32;
    // The notebook that contains the page.  Read-only.
    parentNotebook *Notebook;
    // The section that contains the page. Read-only.
    parentSection *OnenoteSection;
    // The title of the page.
    title *string;
    // 
    userTags []string;
}
// Instantiates a new onenotePage and sets the default values.
func NewOnenotePage()(*OnenotePage) {
    m := &OnenotePage{
        OnenoteEntitySchemaObjectModel: *NewOnenoteEntitySchemaObjectModel(),
    }
    return m
}
// Gets the content property value. The page's HTML content.
func (m *OnenotePage) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// Gets the contentUrl property value. The URL for the page's HTML content.  Read-only.
func (m *OnenotePage) GetContentUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentUrl
    }
}
// Gets the createdByAppId property value. The unique identifier of the application that created the page. Read-only.
func (m *OnenotePage) GetCreatedByAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdByAppId
    }
}
// Gets the lastModifiedDateTime property value. The date and time when the page was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *OnenotePage) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the level property value. The indentation level of the page. Read-only.
func (m *OnenotePage) GetLevel()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.level
    }
}
// Gets the links property value. Links for opening the page. The oneNoteClientURL link opens the page in the OneNote native client if it 's installed. The oneNoteWebUrl link opens the page in OneNote on the web. Read-only.
func (m *OnenotePage) GetLinks()(*PageLinks) {
    if m == nil {
        return nil
    } else {
        return m.links
    }
}
// Gets the order property value. The order of the page within its parent section. Read-only.
func (m *OnenotePage) GetOrder()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.order
    }
}
// Gets the parentNotebook property value. The notebook that contains the page.  Read-only.
func (m *OnenotePage) GetParentNotebook()(*Notebook) {
    if m == nil {
        return nil
    } else {
        return m.parentNotebook
    }
}
// Gets the parentSection property value. The section that contains the page. Read-only.
func (m *OnenotePage) GetParentSection()(*OnenoteSection) {
    if m == nil {
        return nil
    } else {
        return m.parentSection
    }
}
// Gets the title property value. The title of the page.
func (m *OnenotePage) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// Gets the userTags property value. 
func (m *OnenotePage) GetUserTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.userTags
    }
}
// The deserialization information for the current model
func (m *OnenotePage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OnenoteEntitySchemaObjectModel.GetFieldDeserializers()
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetContent(val)
        return nil
    }
    res["contentUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContentUrl(val)
        return nil
    }
    res["createdByAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCreatedByAppId(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["level"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetLevel(val)
        return nil
    }
    res["links"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPageLinks() })
        if err != nil {
            return err
        }
        m.SetLinks(val.(*PageLinks))
        return nil
    }
    res["order"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetOrder(val)
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
    res["parentSection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnenoteSection() })
        if err != nil {
            return err
        }
        m.SetParentSection(val.(*OnenoteSection))
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTitle(val)
        return nil
    }
    res["userTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetUserTags(res)
        return nil
    }
    return res
}
func (m *OnenotePage) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OnenotePage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OnenoteEntitySchemaObjectModel.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentUrl", m.GetContentUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdByAppId", m.GetCreatedByAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("level", m.GetLevel())
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
        err = writer.WriteInt32Value("order", m.GetOrder())
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
        err = writer.WriteObjectValue("parentSection", m.GetParentSection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("userTags", m.GetUserTags())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the content property value. The page's HTML content.
// Parameters:
//  - value : Value to set for the content property.
func (m *OnenotePage) SetContent(value []byte)() {
    m.content = value
}
// Sets the contentUrl property value. The URL for the page's HTML content.  Read-only.
// Parameters:
//  - value : Value to set for the contentUrl property.
func (m *OnenotePage) SetContentUrl(value *string)() {
    m.contentUrl = value
}
// Sets the createdByAppId property value. The unique identifier of the application that created the page. Read-only.
// Parameters:
//  - value : Value to set for the createdByAppId property.
func (m *OnenotePage) SetCreatedByAppId(value *string)() {
    m.createdByAppId = value
}
// Sets the lastModifiedDateTime property value. The date and time when the page was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *OnenotePage) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the level property value. The indentation level of the page. Read-only.
// Parameters:
//  - value : Value to set for the level property.
func (m *OnenotePage) SetLevel(value *int32)() {
    m.level = value
}
// Sets the links property value. Links for opening the page. The oneNoteClientURL link opens the page in the OneNote native client if it 's installed. The oneNoteWebUrl link opens the page in OneNote on the web. Read-only.
// Parameters:
//  - value : Value to set for the links property.
func (m *OnenotePage) SetLinks(value *PageLinks)() {
    m.links = value
}
// Sets the order property value. The order of the page within its parent section. Read-only.
// Parameters:
//  - value : Value to set for the order property.
func (m *OnenotePage) SetOrder(value *int32)() {
    m.order = value
}
// Sets the parentNotebook property value. The notebook that contains the page.  Read-only.
// Parameters:
//  - value : Value to set for the parentNotebook property.
func (m *OnenotePage) SetParentNotebook(value *Notebook)() {
    m.parentNotebook = value
}
// Sets the parentSection property value. The section that contains the page. Read-only.
// Parameters:
//  - value : Value to set for the parentSection property.
func (m *OnenotePage) SetParentSection(value *OnenoteSection)() {
    m.parentSection = value
}
// Sets the title property value. The title of the page.
// Parameters:
//  - value : Value to set for the title property.
func (m *OnenotePage) SetTitle(value *string)() {
    m.title = value
}
// Sets the userTags property value. 
// Parameters:
//  - value : Value to set for the userTags property.
func (m *OnenotePage) SetUserTags(value []string)() {
    m.userTags = value
}
