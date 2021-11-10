package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type NotificationMessageTemplate struct {
    Entity
    // The Message Template Branding Options. Branding is defined in the Intune Admin Console. Possible values are: none, includeCompanyLogo, includeCompanyName, includeContactInformation.
    brandingOptions *NotificationTemplateBrandingOptions;
    // The default locale to fallback onto when the requested locale is not available.
    defaultLocale *string;
    // Display name for the Notification Message Template.
    displayName *string;
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The list of localized messages for this Notification Message Template.
    localizedNotificationMessages []LocalizedNotificationMessage;
}
// Instantiates a new notificationMessageTemplate and sets the default values.
func NewNotificationMessageTemplate()(*NotificationMessageTemplate) {
    m := &NotificationMessageTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the brandingOptions property value. The Message Template Branding Options. Branding is defined in the Intune Admin Console. Possible values are: none, includeCompanyLogo, includeCompanyName, includeContactInformation.
func (m *NotificationMessageTemplate) GetBrandingOptions()(*NotificationTemplateBrandingOptions) {
    if m == nil {
        return nil
    } else {
        return m.brandingOptions
    }
}
// Gets the defaultLocale property value. The default locale to fallback onto when the requested locale is not available.
func (m *NotificationMessageTemplate) GetDefaultLocale()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultLocale
    }
}
// Gets the displayName property value. Display name for the Notification Message Template.
func (m *NotificationMessageTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *NotificationMessageTemplate) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the localizedNotificationMessages property value. The list of localized messages for this Notification Message Template.
func (m *NotificationMessageTemplate) GetLocalizedNotificationMessages()([]LocalizedNotificationMessage) {
    if m == nil {
        return nil
    } else {
        return m.localizedNotificationMessages
    }
}
// The deserialization information for the current model
func (m *NotificationMessageTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["brandingOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseNotificationTemplateBrandingOptions)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(NotificationTemplateBrandingOptions)
            m.SetBrandingOptions(&cast)
        }
        return nil
    }
    res["defaultLocale"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLocale(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["localizedNotificationMessages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocalizedNotificationMessage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LocalizedNotificationMessage, len(val))
            for i, v := range val {
                res[i] = *(v.(*LocalizedNotificationMessage))
            }
            m.SetLocalizedNotificationMessages(res)
        }
        return nil
    }
    return res
}
func (m *NotificationMessageTemplate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *NotificationMessageTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBrandingOptions() != nil {
        cast := m.GetBrandingOptions().String()
        err = writer.WriteStringValue("brandingOptions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultLocale", m.GetDefaultLocale())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLocalizedNotificationMessages()))
        for i, v := range m.GetLocalizedNotificationMessages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("localizedNotificationMessages", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the brandingOptions property value. The Message Template Branding Options. Branding is defined in the Intune Admin Console. Possible values are: none, includeCompanyLogo, includeCompanyName, includeContactInformation.
// Parameters:
//  - value : Value to set for the brandingOptions property.
func (m *NotificationMessageTemplate) SetBrandingOptions(value *NotificationTemplateBrandingOptions)() {
    m.brandingOptions = value
}
// Sets the defaultLocale property value. The default locale to fallback onto when the requested locale is not available.
// Parameters:
//  - value : Value to set for the defaultLocale property.
func (m *NotificationMessageTemplate) SetDefaultLocale(value *string)() {
    m.defaultLocale = value
}
// Sets the displayName property value. Display name for the Notification Message Template.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *NotificationMessageTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastModifiedDateTime property value. DateTime the object was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *NotificationMessageTemplate) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the localizedNotificationMessages property value. The list of localized messages for this Notification Message Template.
// Parameters:
//  - value : Value to set for the localizedNotificationMessages property.
func (m *NotificationMessageTemplate) SetLocalizedNotificationMessages(value []LocalizedNotificationMessage)() {
    m.localizedNotificationMessages = value
}
