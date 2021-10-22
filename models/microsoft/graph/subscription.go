package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Subscription struct {
    Entity
    applicationId *string;
    changeType *string;
    clientState *string;
    creatorId *string;
    encryptionCertificate *string;
    encryptionCertificateId *string;
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    includeResourceData *bool;
    latestSupportedTlsVersion *string;
    lifecycleNotificationUrl *string;
    notificationQueryOptions *string;
    notificationUrl *string;
    resource *string;
}
func NewSubscription()(*Subscription) {
    m := &Subscription{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Subscription) GetApplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationId
    }
}
func (m *Subscription) GetChangeType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.changeType
    }
}
func (m *Subscription) GetClientState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientState
    }
}
func (m *Subscription) GetCreatorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.creatorId
    }
}
func (m *Subscription) GetEncryptionCertificate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.encryptionCertificate
    }
}
func (m *Subscription) GetEncryptionCertificateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.encryptionCertificateId
    }
}
func (m *Subscription) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
func (m *Subscription) GetIncludeResourceData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.includeResourceData
    }
}
func (m *Subscription) GetLatestSupportedTlsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.latestSupportedTlsVersion
    }
}
func (m *Subscription) GetLifecycleNotificationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lifecycleNotificationUrl
    }
}
func (m *Subscription) GetNotificationQueryOptions()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationQueryOptions
    }
}
func (m *Subscription) GetNotificationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationUrl
    }
}
func (m *Subscription) GetResource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
func (m *Subscription) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApplicationId(val)
        return nil
    }
    res["changeType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetChangeType(val)
        return nil
    }
    res["clientState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClientState(val)
        return nil
    }
    res["creatorId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCreatorId(val)
        return nil
    }
    res["encryptionCertificate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEncryptionCertificate(val)
        return nil
    }
    res["encryptionCertificateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEncryptionCertificateId(val)
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpirationDateTime(val)
        return nil
    }
    res["includeResourceData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIncludeResourceData(val)
        return nil
    }
    res["latestSupportedTlsVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLatestSupportedTlsVersion(val)
        return nil
    }
    res["lifecycleNotificationUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLifecycleNotificationUrl(val)
        return nil
    }
    res["notificationQueryOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotificationQueryOptions(val)
        return nil
    }
    res["notificationUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotificationUrl(val)
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResource(val)
        return nil
    }
    return res
}
func (m *Subscription) IsNil()(bool) {
    return m == nil
}
func (m *Subscription) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("applicationId", m.GetApplicationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("changeType", m.GetChangeType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientState", m.GetClientState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("creatorId", m.GetCreatorId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("encryptionCertificate", m.GetEncryptionCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("encryptionCertificateId", m.GetEncryptionCertificateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("includeResourceData", m.GetIncludeResourceData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("latestSupportedTlsVersion", m.GetLatestSupportedTlsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lifecycleNotificationUrl", m.GetLifecycleNotificationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationQueryOptions", m.GetNotificationQueryOptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationUrl", m.GetNotificationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Subscription) SetApplicationId(value *string)() {
    m.applicationId = value
}
func (m *Subscription) SetChangeType(value *string)() {
    m.changeType = value
}
func (m *Subscription) SetClientState(value *string)() {
    m.clientState = value
}
func (m *Subscription) SetCreatorId(value *string)() {
    m.creatorId = value
}
func (m *Subscription) SetEncryptionCertificate(value *string)() {
    m.encryptionCertificate = value
}
func (m *Subscription) SetEncryptionCertificateId(value *string)() {
    m.encryptionCertificateId = value
}
func (m *Subscription) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
func (m *Subscription) SetIncludeResourceData(value *bool)() {
    m.includeResourceData = value
}
func (m *Subscription) SetLatestSupportedTlsVersion(value *string)() {
    m.latestSupportedTlsVersion = value
}
func (m *Subscription) SetLifecycleNotificationUrl(value *string)() {
    m.lifecycleNotificationUrl = value
}
func (m *Subscription) SetNotificationQueryOptions(value *string)() {
    m.notificationQueryOptions = value
}
func (m *Subscription) SetNotificationUrl(value *string)() {
    m.notificationUrl = value
}
func (m *Subscription) SetResource(value *string)() {
    m.resource = value
}
