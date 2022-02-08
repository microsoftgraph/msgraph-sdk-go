package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServiceUpdateMessage 
type ServiceUpdateMessage struct {
    ServiceAnnouncementBase
    // The expected deadline of the action for the message.
    actionRequiredByDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // A collection of serviceAnnouncementAttachments.
    attachments []ServiceAnnouncementAttachment;
    // The zip file that contains all attachments for a message.
    attachmentsArchive []byte;
    // 
    body *ItemBody;
    // The service message category. Possible values are: preventOrFixIssue, planForChange, stayInformed, unknownFutureValue.
    category *ServiceUpdateCategory;
    // Indicates whether the message has any attachment.
    hasAttachments *bool;
    // Indicates whether the message describes a major update for the service.
    isMajorChange *bool;
    // The affected services by the service message.
    services []string;
    // The severity of the service message. Possible values are: normal, high, critical, unknownFutureValue.
    severity *ServiceUpdateSeverity;
    // A collection of tags for the service message. Tags are provided by the service team/support team who post the message to tell whether this message contains privacy data, or whether this message is for a service new feature update, and so on.
    tags []string;
    // Represents user viewpoints data of the service message. This data includes message status such as whether the user has archived, read, or marked the message as favorite. This property is null when accessed with application permissions.
    viewPoint *ServiceUpdateMessageViewpoint;
}
// NewServiceUpdateMessage instantiates a new serviceUpdateMessage and sets the default values.
func NewServiceUpdateMessage()(*ServiceUpdateMessage) {
    m := &ServiceUpdateMessage{
        ServiceAnnouncementBase: *NewServiceAnnouncementBase(),
    }
    return m
}
// GetActionRequiredByDateTime gets the actionRequiredByDateTime property value. The expected deadline of the action for the message.
func (m *ServiceUpdateMessage) GetActionRequiredByDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.actionRequiredByDateTime
    }
}
// GetAttachments gets the attachments property value. A collection of serviceAnnouncementAttachments.
func (m *ServiceUpdateMessage) GetAttachments()([]ServiceAnnouncementAttachment) {
    if m == nil {
        return nil
    } else {
        return m.attachments
    }
}
// GetAttachmentsArchive gets the attachmentsArchive property value. The zip file that contains all attachments for a message.
func (m *ServiceUpdateMessage) GetAttachmentsArchive()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.attachmentsArchive
    }
}
// GetBody gets the body property value. 
func (m *ServiceUpdateMessage) GetBody()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.body
    }
}
// GetCategory gets the category property value. The service message category. Possible values are: preventOrFixIssue, planForChange, stayInformed, unknownFutureValue.
func (m *ServiceUpdateMessage) GetCategory()(*ServiceUpdateCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetHasAttachments gets the hasAttachments property value. Indicates whether the message has any attachment.
func (m *ServiceUpdateMessage) GetHasAttachments()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasAttachments
    }
}
// GetIsMajorChange gets the isMajorChange property value. Indicates whether the message describes a major update for the service.
func (m *ServiceUpdateMessage) GetIsMajorChange()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMajorChange
    }
}
// GetServices gets the services property value. The affected services by the service message.
func (m *ServiceUpdateMessage) GetServices()([]string) {
    if m == nil {
        return nil
    } else {
        return m.services
    }
}
// GetSeverity gets the severity property value. The severity of the service message. Possible values are: normal, high, critical, unknownFutureValue.
func (m *ServiceUpdateMessage) GetSeverity()(*ServiceUpdateSeverity) {
    if m == nil {
        return nil
    } else {
        return m.severity
    }
}
// GetTags gets the tags property value. A collection of tags for the service message. Tags are provided by the service team/support team who post the message to tell whether this message contains privacy data, or whether this message is for a service new feature update, and so on.
func (m *ServiceUpdateMessage) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// GetViewPoint gets the viewPoint property value. Represents user viewpoints data of the service message. This data includes message status such as whether the user has archived, read, or marked the message as favorite. This property is null when accessed with application permissions.
func (m *ServiceUpdateMessage) GetViewPoint()(*ServiceUpdateMessageViewpoint) {
    if m == nil {
        return nil
    } else {
        return m.viewPoint
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceUpdateMessage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ServiceAnnouncementBase.GetFieldDeserializers()
    res["actionRequiredByDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionRequiredByDateTime(val)
        }
        return nil
    }
    res["attachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceAnnouncementAttachment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceAnnouncementAttachment, len(val))
            for i, v := range val {
                res[i] = *(v.(*ServiceAnnouncementAttachment))
            }
            m.SetAttachments(res)
        }
        return nil
    }
    res["attachmentsArchive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachmentsArchive(val)
        }
        return nil
    }
    res["body"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val.(*ItemBody))
        }
        return nil
    }
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceUpdateCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*ServiceUpdateCategory))
        }
        return nil
    }
    res["hasAttachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasAttachments(val)
        }
        return nil
    }
    res["isMajorChange"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMajorChange(val)
        }
        return nil
    }
    res["services"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetServices(res)
        }
        return nil
    }
    res["severity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceUpdateSeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*ServiceUpdateSeverity))
        }
        return nil
    }
    res["tags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTags(res)
        }
        return nil
    }
    res["viewPoint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceUpdateMessageViewpoint() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewPoint(val.(*ServiceUpdateMessageViewpoint))
        }
        return nil
    }
    return res
}
func (m *ServiceUpdateMessage) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ServiceUpdateMessage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ServiceAnnouncementBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("actionRequiredByDateTime", m.GetActionRequiredByDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetAttachments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttachments()))
        for i, v := range m.GetAttachments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("attachments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("attachmentsArchive", m.GetAttachmentsArchive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasAttachments", m.GetHasAttachments())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMajorChange", m.GetIsMajorChange())
        if err != nil {
            return err
        }
    }
    if m.GetServices() != nil {
        err = writer.WriteCollectionOfStringValues("services", m.GetServices())
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err = writer.WriteStringValue("severity", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("viewPoint", m.GetViewPoint())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionRequiredByDateTime sets the actionRequiredByDateTime property value. The expected deadline of the action for the message.
func (m *ServiceUpdateMessage) SetActionRequiredByDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.actionRequiredByDateTime = value
    }
}
// SetAttachments sets the attachments property value. A collection of serviceAnnouncementAttachments.
func (m *ServiceUpdateMessage) SetAttachments(value []ServiceAnnouncementAttachment)() {
    if m != nil {
        m.attachments = value
    }
}
// SetAttachmentsArchive sets the attachmentsArchive property value. The zip file that contains all attachments for a message.
func (m *ServiceUpdateMessage) SetAttachmentsArchive(value []byte)() {
    if m != nil {
        m.attachmentsArchive = value
    }
}
// SetBody sets the body property value. 
func (m *ServiceUpdateMessage) SetBody(value *ItemBody)() {
    if m != nil {
        m.body = value
    }
}
// SetCategory sets the category property value. The service message category. Possible values are: preventOrFixIssue, planForChange, stayInformed, unknownFutureValue.
func (m *ServiceUpdateMessage) SetCategory(value *ServiceUpdateCategory)() {
    if m != nil {
        m.category = value
    }
}
// SetHasAttachments sets the hasAttachments property value. Indicates whether the message has any attachment.
func (m *ServiceUpdateMessage) SetHasAttachments(value *bool)() {
    if m != nil {
        m.hasAttachments = value
    }
}
// SetIsMajorChange sets the isMajorChange property value. Indicates whether the message describes a major update for the service.
func (m *ServiceUpdateMessage) SetIsMajorChange(value *bool)() {
    if m != nil {
        m.isMajorChange = value
    }
}
// SetServices sets the services property value. The affected services by the service message.
func (m *ServiceUpdateMessage) SetServices(value []string)() {
    if m != nil {
        m.services = value
    }
}
// SetSeverity sets the severity property value. The severity of the service message. Possible values are: normal, high, critical, unknownFutureValue.
func (m *ServiceUpdateMessage) SetSeverity(value *ServiceUpdateSeverity)() {
    if m != nil {
        m.severity = value
    }
}
// SetTags sets the tags property value. A collection of tags for the service message. Tags are provided by the service team/support team who post the message to tell whether this message contains privacy data, or whether this message is for a service new feature update, and so on.
func (m *ServiceUpdateMessage) SetTags(value []string)() {
    if m != nil {
        m.tags = value
    }
}
// SetViewPoint sets the viewPoint property value. Represents user viewpoints data of the service message. This data includes message status such as whether the user has archived, read, or marked the message as favorite. This property is null when accessed with application permissions.
func (m *ServiceUpdateMessage) SetViewPoint(value *ServiceUpdateMessageViewpoint)() {
    if m != nil {
        m.viewPoint = value
    }
}
