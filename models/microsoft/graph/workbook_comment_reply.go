package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// workbookCommentReply 
type WorkbookCommentReply struct {
    Entity
    // The content of a comment reply.
    content *string;
    // Indicates the type for the comment reply.
    contentType *string;
}
// NewWorkbookCommentReply instantiates a new workbookCommentReply and sets the default values.
func NewWorkbookCommentReply()(*WorkbookCommentReply) {
    m := &WorkbookCommentReply{
        Entity: *NewEntity(),
    }
    return m
}
// GetContent gets the content property value. The content of a comment reply.
func (m *WorkbookCommentReply) GetContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetContentType gets the contentType property value. Indicates the type for the comment reply.
func (m *WorkbookCommentReply) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookCommentReply) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    return res
}
func (m *WorkbookCommentReply) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookCommentReply) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content of a comment reply.
func (m *WorkbookCommentReply) SetContent(value *string)() {
    m.content = value
}
// SetContentType sets the contentType property value. Indicates the type for the comment reply.
func (m *WorkbookCommentReply) SetContentType(value *string)() {
    m.contentType = value
}
