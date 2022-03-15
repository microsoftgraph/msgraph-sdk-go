package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServiceAnnouncementAttachment provides operations to manage the admin singleton.
type ServiceAnnouncementAttachment struct {
    Entity
    // The attachment content.
    content []byte;
    // 
    contentType *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    name *string;
    // 
    size *int32;
}
// NewServiceAnnouncementAttachment instantiates a new serviceAnnouncementAttachment and sets the default values.
func NewServiceAnnouncementAttachment()(*ServiceAnnouncementAttachment) {
    m := &ServiceAnnouncementAttachment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateServiceAnnouncementAttachmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceAnnouncementAttachmentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewServiceAnnouncementAttachment(), nil
}
// GetContent gets the content property value. The attachment content.
func (m *ServiceAnnouncementAttachment) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetContentType gets the contentType property value. 
func (m *ServiceAnnouncementAttachment) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceAnnouncementAttachment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.Get[]byteValue()
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
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["size"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *ServiceAnnouncementAttachment) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetName gets the name property value. 
func (m *ServiceAnnouncementAttachment) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetSize gets the size property value. 
func (m *ServiceAnnouncementAttachment) GetSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
func (m *ServiceAnnouncementAttachment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ServiceAnnouncementAttachment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
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
        err = writer.WriteStringValue("contentType", m.GetContentType())
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
// SetContent sets the content property value. The attachment content.
func (m *ServiceAnnouncementAttachment) SetContent(value []byte)() {
    if m != nil {
        m.content = value
    }
}
// SetContentType sets the contentType property value. 
func (m *ServiceAnnouncementAttachment) SetContentType(value *string)() {
    if m != nil {
        m.contentType = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *ServiceAnnouncementAttachment) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetName sets the name property value. 
func (m *ServiceAnnouncementAttachment) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetSize sets the size property value. 
func (m *ServiceAnnouncementAttachment) SetSize(value *int32)() {
    if m != nil {
        m.size = value
    }
}
