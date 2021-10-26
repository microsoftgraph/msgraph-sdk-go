package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type LocalizedNotificationMessage struct {
    Entity
    // Flag to indicate whether or not this is the default locale for language fallback. This flag can only be set. To unset, set this property to true on another Localized Notification Message.
    isDefault *bool;
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The Locale for which this message is destined.
    locale *string;
    // The Message Template content.
    messageTemplate *string;
    // The Message Template Subject.
    subject *string;
}
// Instantiates a new localizedNotificationMessage and sets the default values.
func NewLocalizedNotificationMessage()(*LocalizedNotificationMessage) {
    m := &LocalizedNotificationMessage{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the isDefault property value. Flag to indicate whether or not this is the default locale for language fallback. This flag can only be set. To unset, set this property to true on another Localized Notification Message.
func (m *LocalizedNotificationMessage) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// Gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *LocalizedNotificationMessage) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the locale property value. The Locale for which this message is destined.
func (m *LocalizedNotificationMessage) GetLocale()(*string) {
    if m == nil {
        return nil
    } else {
        return m.locale
    }
}
// Gets the messageTemplate property value. The Message Template content.
func (m *LocalizedNotificationMessage) GetMessageTemplate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.messageTemplate
    }
}
// Gets the subject property value. The Message Template Subject.
func (m *LocalizedNotificationMessage) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// The deserialization information for the current model
func (m *LocalizedNotificationMessage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefault(val)
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
    res["locale"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLocale(val)
        return nil
    }
    res["messageTemplate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessageTemplate(val)
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubject(val)
        return nil
    }
    return res
}
func (m *LocalizedNotificationMessage) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *LocalizedNotificationMessage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("locale", m.GetLocale())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("messageTemplate", m.GetMessageTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the isDefault property value. Flag to indicate whether or not this is the default locale for language fallback. This flag can only be set. To unset, set this property to true on another Localized Notification Message.
// Parameters:
//  - value : Value to set for the isDefault property.
func (m *LocalizedNotificationMessage) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// Sets the lastModifiedDateTime property value. DateTime the object was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *LocalizedNotificationMessage) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the locale property value. The Locale for which this message is destined.
// Parameters:
//  - value : Value to set for the locale property.
func (m *LocalizedNotificationMessage) SetLocale(value *string)() {
    m.locale = value
}
// Sets the messageTemplate property value. The Message Template content.
// Parameters:
//  - value : Value to set for the messageTemplate property.
func (m *LocalizedNotificationMessage) SetMessageTemplate(value *string)() {
    m.messageTemplate = value
}
// Sets the subject property value. The Message Template Subject.
// Parameters:
//  - value : Value to set for the subject property.
func (m *LocalizedNotificationMessage) SetSubject(value *string)() {
    m.subject = value
}
