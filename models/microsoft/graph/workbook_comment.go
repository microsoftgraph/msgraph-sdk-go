package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookComment 
type WorkbookComment struct {
    Entity
    // The content of comment.
    content *string;
    // Indicates the type for the comment.
    contentType *string;
    // Read-only. Nullable.
    replies []WorkbookCommentReply;
}
// NewWorkbookComment instantiates a new workbookComment and sets the default values.
func NewWorkbookComment()(*WorkbookComment) {
    m := &WorkbookComment{
        Entity: *NewEntity(),
    }
    return m
}
// GetContent gets the content property value. The content of comment.
func (m *WorkbookComment) GetContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetContentType gets the contentType property value. Indicates the type for the comment.
func (m *WorkbookComment) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// GetReplies gets the replies property value. Read-only. Nullable.
func (m *WorkbookComment) GetReplies()([]WorkbookCommentReply) {
    if m == nil {
        return nil
    } else {
        return m.replies
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookComment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["replies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookCommentReply() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookCommentReply, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkbookCommentReply))
            }
            m.SetReplies(res)
        }
        return nil
    }
    return res
}
func (m *WorkbookComment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookComment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReplies()))
        for i, v := range m.GetReplies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("replies", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content of comment.
func (m *WorkbookComment) SetContent(value *string)() {
    if m != nil {
        m.content = value
    }
}
// SetContentType sets the contentType property value. Indicates the type for the comment.
func (m *WorkbookComment) SetContentType(value *string)() {
    if m != nil {
        m.contentType = value
    }
}
// SetReplies sets the replies property value. Read-only. Nullable.
func (m *WorkbookComment) SetReplies(value []WorkbookCommentReply)() {
    if m != nil {
        m.replies = value
    }
}
