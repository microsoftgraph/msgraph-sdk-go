package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Subscription 
type Subscription struct {
    Entity
    // Identifier of the application used to create the subscription. Read-only.
    applicationId *string;
    // Required. Indicates the type of change in the subscribed resource that will raise a change notification. The supported values are: created, updated, deleted. Multiple values can be combined using a comma-separated list. Note:  Drive root item and list change notifications support only the updated changeType. User and group change notifications support updated and deleted changeType.
    changeType *string;
    // Optional. Specifies the value of the clientState property sent by the service in each change notification. The maximum length is 255 characters. The client can check that the change notification came from the service by comparing the value of the clientState property sent with the subscription with the value of the clientState property received with each change notification.
    clientState *string;
    // Identifier of the user or service principal that created the subscription. If the app used delegated permissions to create the subscription, this field contains the ID of the signed-in user the app called on behalf of. If the app used application permissions, this field contains the ID of the service principal corresponding to the app. Read-only.
    creatorId *string;
    // A base64-encoded representation of a certificate with a public key used to encrypt resource data in change notifications. Optional but required when includeResourceData is true.
    encryptionCertificate *string;
    // Optional. A custom app-provided identifier to help identify the certificate needed to decrypt resource data. Required when includeResourceData is true.
    encryptionCertificateId *string;
    // Required. Specifies the date and time when the webhook subscription expires. The time is in UTC, and can be an amount of time from subscription creation that varies for the resource subscribed to. For the maximum supported subscription length of time, see the table below.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Optional. When set to true, change notifications include resource data (such as content of a chat message).
    includeResourceData *bool;
    // Specifies the latest version of Transport Layer Security (TLS) that the notification endpoint, specified by notificationUrl, supports. The possible values are: v1_0, v1_1, v1_2, v1_3. For subscribers whose notification endpoint supports a version lower than the currently recommended version (TLS 1.2), specifying this property by a set timeline allows them to temporarily use their deprecated version of TLS before completing their upgrade to TLS 1.2. For these subscribers, not setting this property per the timeline would result in subscription operations failing. For subscribers whose notification endpoint already supports TLS 1.2, setting this property is optional. In such cases, Microsoft Graph defaults the property to v1_2.
    latestSupportedTlsVersion *string;
    // Optional. The URL of the endpoint that receives lifecycle notifications, including subscriptionRemoved and missed notifications. This URL must make use of the HTTPS protocol.
    lifecycleNotificationUrl *string;
    // OData query options for specifying the value for the targeting resource. Clients receive notifications when the resource reaches the state matching the query options provided here. With this new property in the subscription creation payload along with all existing properties, Webhooks will deliver notifications whenever a resource reaches the desired state mentioned in the notificationQueryOptions property. For example, when the print job is completed or when a print job resource isFetchable property value becomes true etc.
    notificationQueryOptions *string;
    // Required. The URL of the endpoint that receives the change notifications. This URL must make use of the HTTPS protocol.
    notificationUrl *string;
    // 
    notificationUrlAppId *string;
    // Required. Specifies the resource that will be monitored for changes. Do not include the base URL (https://graph.microsoft.com/beta/). See the possible resource path values for each supported resource.
    resource *string;
}
// NewSubscription instantiates a new subscription and sets the default values.
func NewSubscription()(*Subscription) {
    m := &Subscription{
        Entity: *NewEntity(),
    }
    return m
}
// GetApplicationId gets the applicationId property value. Identifier of the application used to create the subscription. Read-only.
func (m *Subscription) GetApplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationId
    }
}
// GetChangeType gets the changeType property value. Required. Indicates the type of change in the subscribed resource that will raise a change notification. The supported values are: created, updated, deleted. Multiple values can be combined using a comma-separated list. Note:  Drive root item and list change notifications support only the updated changeType. User and group change notifications support updated and deleted changeType.
func (m *Subscription) GetChangeType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.changeType
    }
}
// GetClientState gets the clientState property value. Optional. Specifies the value of the clientState property sent by the service in each change notification. The maximum length is 255 characters. The client can check that the change notification came from the service by comparing the value of the clientState property sent with the subscription with the value of the clientState property received with each change notification.
func (m *Subscription) GetClientState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientState
    }
}
// GetCreatorId gets the creatorId property value. Identifier of the user or service principal that created the subscription. If the app used delegated permissions to create the subscription, this field contains the ID of the signed-in user the app called on behalf of. If the app used application permissions, this field contains the ID of the service principal corresponding to the app. Read-only.
func (m *Subscription) GetCreatorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.creatorId
    }
}
// GetEncryptionCertificate gets the encryptionCertificate property value. A base64-encoded representation of a certificate with a public key used to encrypt resource data in change notifications. Optional but required when includeResourceData is true.
func (m *Subscription) GetEncryptionCertificate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.encryptionCertificate
    }
}
// GetEncryptionCertificateId gets the encryptionCertificateId property value. Optional. A custom app-provided identifier to help identify the certificate needed to decrypt resource data. Required when includeResourceData is true.
func (m *Subscription) GetEncryptionCertificateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.encryptionCertificateId
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. Required. Specifies the date and time when the webhook subscription expires. The time is in UTC, and can be an amount of time from subscription creation that varies for the resource subscribed to. For the maximum supported subscription length of time, see the table below.
func (m *Subscription) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetIncludeResourceData gets the includeResourceData property value. Optional. When set to true, change notifications include resource data (such as content of a chat message).
func (m *Subscription) GetIncludeResourceData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.includeResourceData
    }
}
// GetLatestSupportedTlsVersion gets the latestSupportedTlsVersion property value. Specifies the latest version of Transport Layer Security (TLS) that the notification endpoint, specified by notificationUrl, supports. The possible values are: v1_0, v1_1, v1_2, v1_3. For subscribers whose notification endpoint supports a version lower than the currently recommended version (TLS 1.2), specifying this property by a set timeline allows them to temporarily use their deprecated version of TLS before completing their upgrade to TLS 1.2. For these subscribers, not setting this property per the timeline would result in subscription operations failing. For subscribers whose notification endpoint already supports TLS 1.2, setting this property is optional. In such cases, Microsoft Graph defaults the property to v1_2.
func (m *Subscription) GetLatestSupportedTlsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.latestSupportedTlsVersion
    }
}
// GetLifecycleNotificationUrl gets the lifecycleNotificationUrl property value. Optional. The URL of the endpoint that receives lifecycle notifications, including subscriptionRemoved and missed notifications. This URL must make use of the HTTPS protocol.
func (m *Subscription) GetLifecycleNotificationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lifecycleNotificationUrl
    }
}
// GetNotificationQueryOptions gets the notificationQueryOptions property value. OData query options for specifying the value for the targeting resource. Clients receive notifications when the resource reaches the state matching the query options provided here. With this new property in the subscription creation payload along with all existing properties, Webhooks will deliver notifications whenever a resource reaches the desired state mentioned in the notificationQueryOptions property. For example, when the print job is completed or when a print job resource isFetchable property value becomes true etc.
func (m *Subscription) GetNotificationQueryOptions()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationQueryOptions
    }
}
// GetNotificationUrl gets the notificationUrl property value. Required. The URL of the endpoint that receives the change notifications. This URL must make use of the HTTPS protocol.
func (m *Subscription) GetNotificationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationUrl
    }
}
// GetNotificationUrlAppId gets the notificationUrlAppId property value. 
func (m *Subscription) GetNotificationUrlAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationUrlAppId
    }
}
// GetResource gets the resource property value. Required. Specifies the resource that will be monitored for changes. Do not include the base URL (https://graph.microsoft.com/beta/). See the possible resource path values for each supported resource.
func (m *Subscription) GetResource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Subscription) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationId(val)
        }
        return nil
    }
    res["changeType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeType(val)
        }
        return nil
    }
    res["clientState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientState(val)
        }
        return nil
    }
    res["creatorId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatorId(val)
        }
        return nil
    }
    res["encryptionCertificate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionCertificate(val)
        }
        return nil
    }
    res["encryptionCertificateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionCertificateId(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["includeResourceData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeResourceData(val)
        }
        return nil
    }
    res["latestSupportedTlsVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatestSupportedTlsVersion(val)
        }
        return nil
    }
    res["lifecycleNotificationUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLifecycleNotificationUrl(val)
        }
        return nil
    }
    res["notificationQueryOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationQueryOptions(val)
        }
        return nil
    }
    res["notificationUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationUrl(val)
        }
        return nil
    }
    res["notificationUrlAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationUrlAppId(val)
        }
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val)
        }
        return nil
    }
    return res
}
func (m *Subscription) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        err = writer.WriteStringValue("notificationUrlAppId", m.GetNotificationUrlAppId())
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
// SetApplicationId sets the applicationId property value. Identifier of the application used to create the subscription. Read-only.
func (m *Subscription) SetApplicationId(value *string)() {
    if m != nil {
        m.applicationId = value
    }
}
// SetChangeType sets the changeType property value. Required. Indicates the type of change in the subscribed resource that will raise a change notification. The supported values are: created, updated, deleted. Multiple values can be combined using a comma-separated list. Note:  Drive root item and list change notifications support only the updated changeType. User and group change notifications support updated and deleted changeType.
func (m *Subscription) SetChangeType(value *string)() {
    if m != nil {
        m.changeType = value
    }
}
// SetClientState sets the clientState property value. Optional. Specifies the value of the clientState property sent by the service in each change notification. The maximum length is 255 characters. The client can check that the change notification came from the service by comparing the value of the clientState property sent with the subscription with the value of the clientState property received with each change notification.
func (m *Subscription) SetClientState(value *string)() {
    if m != nil {
        m.clientState = value
    }
}
// SetCreatorId sets the creatorId property value. Identifier of the user or service principal that created the subscription. If the app used delegated permissions to create the subscription, this field contains the ID of the signed-in user the app called on behalf of. If the app used application permissions, this field contains the ID of the service principal corresponding to the app. Read-only.
func (m *Subscription) SetCreatorId(value *string)() {
    if m != nil {
        m.creatorId = value
    }
}
// SetEncryptionCertificate sets the encryptionCertificate property value. A base64-encoded representation of a certificate with a public key used to encrypt resource data in change notifications. Optional but required when includeResourceData is true.
func (m *Subscription) SetEncryptionCertificate(value *string)() {
    if m != nil {
        m.encryptionCertificate = value
    }
}
// SetEncryptionCertificateId sets the encryptionCertificateId property value. Optional. A custom app-provided identifier to help identify the certificate needed to decrypt resource data. Required when includeResourceData is true.
func (m *Subscription) SetEncryptionCertificateId(value *string)() {
    if m != nil {
        m.encryptionCertificateId = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. Required. Specifies the date and time when the webhook subscription expires. The time is in UTC, and can be an amount of time from subscription creation that varies for the resource subscribed to. For the maximum supported subscription length of time, see the table below.
func (m *Subscription) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetIncludeResourceData sets the includeResourceData property value. Optional. When set to true, change notifications include resource data (such as content of a chat message).
func (m *Subscription) SetIncludeResourceData(value *bool)() {
    if m != nil {
        m.includeResourceData = value
    }
}
// SetLatestSupportedTlsVersion sets the latestSupportedTlsVersion property value. Specifies the latest version of Transport Layer Security (TLS) that the notification endpoint, specified by notificationUrl, supports. The possible values are: v1_0, v1_1, v1_2, v1_3. For subscribers whose notification endpoint supports a version lower than the currently recommended version (TLS 1.2), specifying this property by a set timeline allows them to temporarily use their deprecated version of TLS before completing their upgrade to TLS 1.2. For these subscribers, not setting this property per the timeline would result in subscription operations failing. For subscribers whose notification endpoint already supports TLS 1.2, setting this property is optional. In such cases, Microsoft Graph defaults the property to v1_2.
func (m *Subscription) SetLatestSupportedTlsVersion(value *string)() {
    if m != nil {
        m.latestSupportedTlsVersion = value
    }
}
// SetLifecycleNotificationUrl sets the lifecycleNotificationUrl property value. Optional. The URL of the endpoint that receives lifecycle notifications, including subscriptionRemoved and missed notifications. This URL must make use of the HTTPS protocol.
func (m *Subscription) SetLifecycleNotificationUrl(value *string)() {
    if m != nil {
        m.lifecycleNotificationUrl = value
    }
}
// SetNotificationQueryOptions sets the notificationQueryOptions property value. OData query options for specifying the value for the targeting resource. Clients receive notifications when the resource reaches the state matching the query options provided here. With this new property in the subscription creation payload along with all existing properties, Webhooks will deliver notifications whenever a resource reaches the desired state mentioned in the notificationQueryOptions property. For example, when the print job is completed or when a print job resource isFetchable property value becomes true etc.
func (m *Subscription) SetNotificationQueryOptions(value *string)() {
    if m != nil {
        m.notificationQueryOptions = value
    }
}
// SetNotificationUrl sets the notificationUrl property value. Required. The URL of the endpoint that receives the change notifications. This URL must make use of the HTTPS protocol.
func (m *Subscription) SetNotificationUrl(value *string)() {
    if m != nil {
        m.notificationUrl = value
    }
}
// SetNotificationUrlAppId sets the notificationUrlAppId property value. 
func (m *Subscription) SetNotificationUrlAppId(value *string)() {
    if m != nil {
        m.notificationUrlAppId = value
    }
}
// SetResource sets the resource property value. Required. Specifies the resource that will be monitored for changes. Do not include the base URL (https://graph.microsoft.com/beta/). See the possible resource path values for each supported resource.
func (m *Subscription) SetResource(value *string)() {
    if m != nil {
        m.resource = value
    }
}
