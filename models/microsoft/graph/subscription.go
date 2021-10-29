package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Subscription struct {
    Entity
    // Identifier of the application used to create the subscription. Read-only.
    applicationId *string;
    // Required. Indicates the type of change in the subscribed resource that will raise a change notification. The supported values are: created, updated, deleted. Multiple values can be combined using a comma-separated list.Note: Drive root item and list change notifications support only the updated changeType. User and group change notifications support updated and deleted changeType.
    changeType *string;
    // Optional. Specifies the value of the clientState property sent by the service in each change notification. The maximum length is 128 characters. The client can check that the change notification came from the service by comparing the value of the clientState property sent with the subscription with the value of the clientState property received with each change notification.
    clientState *string;
    // Identifier of the user or service principal that created the subscription. If the app used delegated permissions to create the subscription, this field contains the id of the signed-in user the app called on behalf of. If the app used application permissions, this field contains the id of the service principal corresponding to the app. Read-only.
    creatorId *string;
    // A base64-encoded representation of a certificate with a public key used to encrypt resource data in change notifications. Optional. Required when includeResourceData is true.
    encryptionCertificate *string;
    // A custom app-provided identifier to help identify the certificate needed to decrypt resource data. Optional.
    encryptionCertificateId *string;
    // Required. Specifies the date and time when the webhook subscription expires. The time is in UTC, and can be an amount of time from subscription creation that varies for the resource subscribed to.  See the table below for maximum supported subscription length of time.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // When set to true, change notifications include resource data (such as content of a chat message). Optional.
    includeResourceData *bool;
    // Specifies the latest version of Transport Layer Security (TLS) that the notification endpoint, specified by notificationUrl, supports. The possible values are: v1_0, v1_1, v1_2, v1_3. For subscribers whose notification endpoint supports a version lower than the currently recommended version (TLS 1.2), specifying this property by a set timeline allows them to temporarily use their deprecated version of TLS before completing their upgrade to TLS 1.2. For these subscribers, not setting this property per the timeline would result in subscription operations failing. For subscribers whose notification endpoint already supports TLS 1.2, setting this property is optional. In such cases, Microsoft Graph defaults the property to v1_2.
    latestSupportedTlsVersion *string;
    // The URL of the endpoint that receives lifecycle notifications, including subscriptionRemoved and missed notifications. This URL must make use of the HTTPS protocol. Optional. Read more about how Outlook resources use lifecycle notifications.
    lifecycleNotificationUrl *string;
    // OData Query Options for specifying value for the targeting resource. Clients receive notifications when resource reaches the state matching the query options provided here. With this new property in the subscription creation payload along with all existing properties, Webhooks will deliver notifications whenever a resource reaches the desired state mentioned in the notificationQueryOptions property eg  when the print job is completed, when a print job resource isFetchable property value becomes true etc.
    notificationQueryOptions *string;
    // Required. The URL of the endpoint that will receive the change notifications. This URL must make use of the HTTPS protocol.
    notificationUrl *string;
    // 
    notificationUrlAppId *string;
    // Required. Specifies the resource that will be monitored for changes. Do not include the base URL (https://graph.microsoft.com/v1.0/). See the possible resource path values for each supported resource.
    resource *string;
}
// Instantiates a new subscription and sets the default values.
func NewSubscription()(*Subscription) {
    m := &Subscription{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the applicationId property value. Identifier of the application used to create the subscription. Read-only.
func (m *Subscription) GetApplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationId
    }
}
// Gets the changeType property value. Required. Indicates the type of change in the subscribed resource that will raise a change notification. The supported values are: created, updated, deleted. Multiple values can be combined using a comma-separated list.Note: Drive root item and list change notifications support only the updated changeType. User and group change notifications support updated and deleted changeType.
func (m *Subscription) GetChangeType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.changeType
    }
}
// Gets the clientState property value. Optional. Specifies the value of the clientState property sent by the service in each change notification. The maximum length is 128 characters. The client can check that the change notification came from the service by comparing the value of the clientState property sent with the subscription with the value of the clientState property received with each change notification.
func (m *Subscription) GetClientState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientState
    }
}
// Gets the creatorId property value. Identifier of the user or service principal that created the subscription. If the app used delegated permissions to create the subscription, this field contains the id of the signed-in user the app called on behalf of. If the app used application permissions, this field contains the id of the service principal corresponding to the app. Read-only.
func (m *Subscription) GetCreatorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.creatorId
    }
}
// Gets the encryptionCertificate property value. A base64-encoded representation of a certificate with a public key used to encrypt resource data in change notifications. Optional. Required when includeResourceData is true.
func (m *Subscription) GetEncryptionCertificate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.encryptionCertificate
    }
}
// Gets the encryptionCertificateId property value. A custom app-provided identifier to help identify the certificate needed to decrypt resource data. Optional.
func (m *Subscription) GetEncryptionCertificateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.encryptionCertificateId
    }
}
// Gets the expirationDateTime property value. Required. Specifies the date and time when the webhook subscription expires. The time is in UTC, and can be an amount of time from subscription creation that varies for the resource subscribed to.  See the table below for maximum supported subscription length of time.
func (m *Subscription) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the includeResourceData property value. When set to true, change notifications include resource data (such as content of a chat message). Optional.
func (m *Subscription) GetIncludeResourceData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.includeResourceData
    }
}
// Gets the latestSupportedTlsVersion property value. Specifies the latest version of Transport Layer Security (TLS) that the notification endpoint, specified by notificationUrl, supports. The possible values are: v1_0, v1_1, v1_2, v1_3. For subscribers whose notification endpoint supports a version lower than the currently recommended version (TLS 1.2), specifying this property by a set timeline allows them to temporarily use their deprecated version of TLS before completing their upgrade to TLS 1.2. For these subscribers, not setting this property per the timeline would result in subscription operations failing. For subscribers whose notification endpoint already supports TLS 1.2, setting this property is optional. In such cases, Microsoft Graph defaults the property to v1_2.
func (m *Subscription) GetLatestSupportedTlsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.latestSupportedTlsVersion
    }
}
// Gets the lifecycleNotificationUrl property value. The URL of the endpoint that receives lifecycle notifications, including subscriptionRemoved and missed notifications. This URL must make use of the HTTPS protocol. Optional. Read more about how Outlook resources use lifecycle notifications.
func (m *Subscription) GetLifecycleNotificationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lifecycleNotificationUrl
    }
}
// Gets the notificationQueryOptions property value. OData Query Options for specifying value for the targeting resource. Clients receive notifications when resource reaches the state matching the query options provided here. With this new property in the subscription creation payload along with all existing properties, Webhooks will deliver notifications whenever a resource reaches the desired state mentioned in the notificationQueryOptions property eg  when the print job is completed, when a print job resource isFetchable property value becomes true etc.
func (m *Subscription) GetNotificationQueryOptions()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationQueryOptions
    }
}
// Gets the notificationUrl property value. Required. The URL of the endpoint that will receive the change notifications. This URL must make use of the HTTPS protocol.
func (m *Subscription) GetNotificationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationUrl
    }
}
// Gets the notificationUrlAppId property value. 
func (m *Subscription) GetNotificationUrlAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationUrlAppId
    }
}
// Gets the resource property value. Required. Specifies the resource that will be monitored for changes. Do not include the base URL (https://graph.microsoft.com/v1.0/). See the possible resource path values for each supported resource.
func (m *Subscription) GetResource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// The deserialization information for the current model
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
    res["notificationUrlAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotificationUrlAppId(val)
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the applicationId property value. Identifier of the application used to create the subscription. Read-only.
// Parameters:
//  - value : Value to set for the applicationId property.
func (m *Subscription) SetApplicationId(value *string)() {
    m.applicationId = value
}
// Sets the changeType property value. Required. Indicates the type of change in the subscribed resource that will raise a change notification. The supported values are: created, updated, deleted. Multiple values can be combined using a comma-separated list.Note: Drive root item and list change notifications support only the updated changeType. User and group change notifications support updated and deleted changeType.
// Parameters:
//  - value : Value to set for the changeType property.
func (m *Subscription) SetChangeType(value *string)() {
    m.changeType = value
}
// Sets the clientState property value. Optional. Specifies the value of the clientState property sent by the service in each change notification. The maximum length is 128 characters. The client can check that the change notification came from the service by comparing the value of the clientState property sent with the subscription with the value of the clientState property received with each change notification.
// Parameters:
//  - value : Value to set for the clientState property.
func (m *Subscription) SetClientState(value *string)() {
    m.clientState = value
}
// Sets the creatorId property value. Identifier of the user or service principal that created the subscription. If the app used delegated permissions to create the subscription, this field contains the id of the signed-in user the app called on behalf of. If the app used application permissions, this field contains the id of the service principal corresponding to the app. Read-only.
// Parameters:
//  - value : Value to set for the creatorId property.
func (m *Subscription) SetCreatorId(value *string)() {
    m.creatorId = value
}
// Sets the encryptionCertificate property value. A base64-encoded representation of a certificate with a public key used to encrypt resource data in change notifications. Optional. Required when includeResourceData is true.
// Parameters:
//  - value : Value to set for the encryptionCertificate property.
func (m *Subscription) SetEncryptionCertificate(value *string)() {
    m.encryptionCertificate = value
}
// Sets the encryptionCertificateId property value. A custom app-provided identifier to help identify the certificate needed to decrypt resource data. Optional.
// Parameters:
//  - value : Value to set for the encryptionCertificateId property.
func (m *Subscription) SetEncryptionCertificateId(value *string)() {
    m.encryptionCertificateId = value
}
// Sets the expirationDateTime property value. Required. Specifies the date and time when the webhook subscription expires. The time is in UTC, and can be an amount of time from subscription creation that varies for the resource subscribed to.  See the table below for maximum supported subscription length of time.
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *Subscription) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the includeResourceData property value. When set to true, change notifications include resource data (such as content of a chat message). Optional.
// Parameters:
//  - value : Value to set for the includeResourceData property.
func (m *Subscription) SetIncludeResourceData(value *bool)() {
    m.includeResourceData = value
}
// Sets the latestSupportedTlsVersion property value. Specifies the latest version of Transport Layer Security (TLS) that the notification endpoint, specified by notificationUrl, supports. The possible values are: v1_0, v1_1, v1_2, v1_3. For subscribers whose notification endpoint supports a version lower than the currently recommended version (TLS 1.2), specifying this property by a set timeline allows them to temporarily use their deprecated version of TLS before completing their upgrade to TLS 1.2. For these subscribers, not setting this property per the timeline would result in subscription operations failing. For subscribers whose notification endpoint already supports TLS 1.2, setting this property is optional. In such cases, Microsoft Graph defaults the property to v1_2.
// Parameters:
//  - value : Value to set for the latestSupportedTlsVersion property.
func (m *Subscription) SetLatestSupportedTlsVersion(value *string)() {
    m.latestSupportedTlsVersion = value
}
// Sets the lifecycleNotificationUrl property value. The URL of the endpoint that receives lifecycle notifications, including subscriptionRemoved and missed notifications. This URL must make use of the HTTPS protocol. Optional. Read more about how Outlook resources use lifecycle notifications.
// Parameters:
//  - value : Value to set for the lifecycleNotificationUrl property.
func (m *Subscription) SetLifecycleNotificationUrl(value *string)() {
    m.lifecycleNotificationUrl = value
}
// Sets the notificationQueryOptions property value. OData Query Options for specifying value for the targeting resource. Clients receive notifications when resource reaches the state matching the query options provided here. With this new property in the subscription creation payload along with all existing properties, Webhooks will deliver notifications whenever a resource reaches the desired state mentioned in the notificationQueryOptions property eg  when the print job is completed, when a print job resource isFetchable property value becomes true etc.
// Parameters:
//  - value : Value to set for the notificationQueryOptions property.
func (m *Subscription) SetNotificationQueryOptions(value *string)() {
    m.notificationQueryOptions = value
}
// Sets the notificationUrl property value. Required. The URL of the endpoint that will receive the change notifications. This URL must make use of the HTTPS protocol.
// Parameters:
//  - value : Value to set for the notificationUrl property.
func (m *Subscription) SetNotificationUrl(value *string)() {
    m.notificationUrl = value
}
// Sets the notificationUrlAppId property value. 
// Parameters:
//  - value : Value to set for the notificationUrlAppId property.
func (m *Subscription) SetNotificationUrlAppId(value *string)() {
    m.notificationUrlAppId = value
}
// Sets the resource property value. Required. Specifies the resource that will be monitored for changes. Do not include the base URL (https://graph.microsoft.com/v1.0/). See the possible resource path values for each supported resource.
// Parameters:
//  - value : Value to set for the resource property.
func (m *Subscription) SetResource(value *string)() {
    m.resource = value
}
