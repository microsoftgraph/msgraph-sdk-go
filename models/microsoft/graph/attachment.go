package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Attachment struct {
    Entity
    // The MIME type.
    contentType *string;
    // true if the attachment is an inline attachment; otherwise, false.
    isInline *bool;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The attachment's file name.
    name *string;
    // The length of the attachment in bytes.
    size *int32;
}
// Instantiates a new attachment and sets the default values.
func NewAttachment()(*Attachment) {
    m := &Attachment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the contentType property value. The MIME type.
func (m *Attachment) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// Gets the isInline property value. true if the attachment is an inline attachment; otherwise, false.
func (m *Attachment) GetIsInline()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInline
    }
}
// Gets the lastModifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Attachment) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the name property value. The attachment's file name.
func (m *Attachment) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the size property value. The length of the attachment in bytes.
func (m *Attachment) GetSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// The deserialization information for the current model
func (m *Attachment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContentType(val)
        return nil
    }
    res["isInline"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsInline(val)
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
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["size"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSize(val)
        return nil
    }
    return res
}
func (m *Attachment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Attachment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isInline", m.GetIsInline())
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
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the contentType property value. The MIME type.
// Parameters:
//  - value : Value to set for the contentType property.
func (m *Attachment) SetContentType(value *string)() {
    m.contentType = value
}
// Sets the isInline property value. true if the attachment is an inline attachment; otherwise, false.
// Parameters:
//  - value : Value to set for the isInline property.
func (m *Attachment) SetIsInline(value *bool)() {
    m.isInline = value
}
// Sets the lastModifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *Attachment) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the name property value. The attachment's file name.
// Parameters:
//  - value : Value to set for the name property.
func (m *Attachment) SetName(value *string)() {
    m.name = value
}
// Sets the size property value. The length of the attachment in bytes.
// Parameters:
//  - value : Value to set for the size property.
func (m *Attachment) SetSize(value *int32)() {
    m.size = value
}
