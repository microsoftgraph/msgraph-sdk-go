package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserActivity 
type UserActivity struct {
    Entity
    // Required. URL used to launch the activity in the best native experience represented by the appId. Might launch a web-based app if no native app exists.
    activationUrl *string;
    // Required. URL for the domain representing the cross-platform identity mapping for the app. Mapping is stored either as a JSON file hosted on the domain or configurable via Windows Dev Center. The JSON file is named cross-platform-app-identifiers and is hosted at root of your HTTPS domain, either at the top level domain or include a sub domain. For example: https://contoso.com or https://myapp.contoso.com but NOT https://myapp.contoso.com/somepath. You must have a unique file and domain (or sub domain) per cross-platform app identity. For example, a separate file and domain is needed for Word vs. PowerPoint.
    activitySourceHost *string;
    // Required. The unique activity ID in the context of the app - supplied by caller and immutable thereafter.
    appActivityId *string;
    // Optional. Short text description of the app used to generate the activity for use in cases when the app is not installed on the user’s local device.
    appDisplayName *string;
    // Optional. A custom piece of data - JSON-LD extensible description of content according to schema.org syntax.
    contentInfo *Json;
    // Optional. Used in the event the content can be rendered outside of a native or web-based app experience (for example, a pointer to an item in an RSS feed).
    contentUrl *string;
    // Set by the server. DateTime in UTC when the object was created on the server.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Set by the server. DateTime in UTC when the object expired on the server.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Optional. URL used to launch the activity in a web-based app, if available.
    fallbackUrl *string;
    // Optional. NavigationProperty/Containment; navigation property to the activity's historyItems.
    historyItems []ActivityHistoryItem;
    // Set by the server. DateTime in UTC when the object was modified on the server.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Set by the server. A status code used to identify valid objects. Values: active, updated, deleted, ignored.
    status *Status;
    // Optional. The timezone in which the user's device used to generate the activity was located at activity creation time; values supplied as Olson IDs in order to support cross-platform representation.
    userTimezone *string;
    // 
    visualElements *VisualInfo;
}
// NewUserActivity instantiates a new userActivity and sets the default values.
func NewUserActivity()(*UserActivity) {
    m := &UserActivity{
        Entity: *NewEntity(),
    }
    return m
}
// GetActivationUrl gets the activationUrl property value. Required. URL used to launch the activity in the best native experience represented by the appId. Might launch a web-based app if no native app exists.
func (m *UserActivity) GetActivationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activationUrl
    }
}
// GetActivitySourceHost gets the activitySourceHost property value. Required. URL for the domain representing the cross-platform identity mapping for the app. Mapping is stored either as a JSON file hosted on the domain or configurable via Windows Dev Center. The JSON file is named cross-platform-app-identifiers and is hosted at root of your HTTPS domain, either at the top level domain or include a sub domain. For example: https://contoso.com or https://myapp.contoso.com but NOT https://myapp.contoso.com/somepath. You must have a unique file and domain (or sub domain) per cross-platform app identity. For example, a separate file and domain is needed for Word vs. PowerPoint.
func (m *UserActivity) GetActivitySourceHost()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activitySourceHost
    }
}
// GetAppActivityId gets the appActivityId property value. Required. The unique activity ID in the context of the app - supplied by caller and immutable thereafter.
func (m *UserActivity) GetAppActivityId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appActivityId
    }
}
// GetAppDisplayName gets the appDisplayName property value. Optional. Short text description of the app used to generate the activity for use in cases when the app is not installed on the user’s local device.
func (m *UserActivity) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// GetContentInfo gets the contentInfo property value. Optional. A custom piece of data - JSON-LD extensible description of content according to schema.org syntax.
func (m *UserActivity) GetContentInfo()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.contentInfo
    }
}
// GetContentUrl gets the contentUrl property value. Optional. Used in the event the content can be rendered outside of a native or web-based app experience (for example, a pointer to an item in an RSS feed).
func (m *UserActivity) GetContentUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentUrl
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Set by the server. DateTime in UTC when the object was created on the server.
func (m *UserActivity) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. Set by the server. DateTime in UTC when the object expired on the server.
func (m *UserActivity) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetFallbackUrl gets the fallbackUrl property value. Optional. URL used to launch the activity in a web-based app, if available.
func (m *UserActivity) GetFallbackUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fallbackUrl
    }
}
// GetHistoryItems gets the historyItems property value. Optional. NavigationProperty/Containment; navigation property to the activity's historyItems.
func (m *UserActivity) GetHistoryItems()([]ActivityHistoryItem) {
    if m == nil {
        return nil
    } else {
        return m.historyItems
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Set by the server. DateTime in UTC when the object was modified on the server.
func (m *UserActivity) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetStatus gets the status property value. Set by the server. A status code used to identify valid objects. Values: active, updated, deleted, ignored.
func (m *UserActivity) GetStatus()(*Status) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetUserTimezone gets the userTimezone property value. Optional. The timezone in which the user's device used to generate the activity was located at activity creation time; values supplied as Olson IDs in order to support cross-platform representation.
func (m *UserActivity) GetUserTimezone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userTimezone
    }
}
// GetVisualElements gets the visualElements property value. 
func (m *UserActivity) GetVisualElements()(*VisualInfo) {
    if m == nil {
        return nil
    } else {
        return m.visualElements
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserActivity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activationUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationUrl(val)
        }
        return nil
    }
    res["activitySourceHost"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivitySourceHost(val)
        }
        return nil
    }
    res["appActivityId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActivityId(val)
        }
        return nil
    }
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["contentInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentInfo(val.(*Json))
        }
        return nil
    }
    res["contentUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentUrl(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["fallbackUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFallbackUrl(val)
        }
        return nil
    }
    res["historyItems"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewActivityHistoryItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActivityHistoryItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*ActivityHistoryItem))
            }
            m.SetHistoryItems(res)
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
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(Status)
            m.SetStatus(&cast)
        }
        return nil
    }
    res["userTimezone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserTimezone(val)
        }
        return nil
    }
    res["visualElements"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVisualInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisualElements(val.(*VisualInfo))
        }
        return nil
    }
    return res
}
func (m *UserActivity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserActivity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("activationUrl", m.GetActivationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("activitySourceHost", m.GetActivitySourceHost())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appActivityId", m.GetAppActivityId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("contentInfo", m.GetContentInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentUrl", m.GetContentUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
        err = writer.WriteStringValue("fallbackUrl", m.GetFallbackUrl())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHistoryItems()))
        for i, v := range m.GetHistoryItems() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("historyItems", cast)
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
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userTimezone", m.GetUserTimezone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("visualElements", m.GetVisualElements())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivationUrl sets the activationUrl property value. Required. URL used to launch the activity in the best native experience represented by the appId. Might launch a web-based app if no native app exists.
func (m *UserActivity) SetActivationUrl(value *string)() {
    if m != nil {
        m.activationUrl = value
    }
}
// SetActivitySourceHost sets the activitySourceHost property value. Required. URL for the domain representing the cross-platform identity mapping for the app. Mapping is stored either as a JSON file hosted on the domain or configurable via Windows Dev Center. The JSON file is named cross-platform-app-identifiers and is hosted at root of your HTTPS domain, either at the top level domain or include a sub domain. For example: https://contoso.com or https://myapp.contoso.com but NOT https://myapp.contoso.com/somepath. You must have a unique file and domain (or sub domain) per cross-platform app identity. For example, a separate file and domain is needed for Word vs. PowerPoint.
func (m *UserActivity) SetActivitySourceHost(value *string)() {
    if m != nil {
        m.activitySourceHost = value
    }
}
// SetAppActivityId sets the appActivityId property value. Required. The unique activity ID in the context of the app - supplied by caller and immutable thereafter.
func (m *UserActivity) SetAppActivityId(value *string)() {
    if m != nil {
        m.appActivityId = value
    }
}
// SetAppDisplayName sets the appDisplayName property value. Optional. Short text description of the app used to generate the activity for use in cases when the app is not installed on the user’s local device.
func (m *UserActivity) SetAppDisplayName(value *string)() {
    if m != nil {
        m.appDisplayName = value
    }
}
// SetContentInfo sets the contentInfo property value. Optional. A custom piece of data - JSON-LD extensible description of content according to schema.org syntax.
func (m *UserActivity) SetContentInfo(value *Json)() {
    if m != nil {
        m.contentInfo = value
    }
}
// SetContentUrl sets the contentUrl property value. Optional. Used in the event the content can be rendered outside of a native or web-based app experience (for example, a pointer to an item in an RSS feed).
func (m *UserActivity) SetContentUrl(value *string)() {
    if m != nil {
        m.contentUrl = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Set by the server. DateTime in UTC when the object was created on the server.
func (m *UserActivity) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. Set by the server. DateTime in UTC when the object expired on the server.
func (m *UserActivity) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetFallbackUrl sets the fallbackUrl property value. Optional. URL used to launch the activity in a web-based app, if available.
func (m *UserActivity) SetFallbackUrl(value *string)() {
    if m != nil {
        m.fallbackUrl = value
    }
}
// SetHistoryItems sets the historyItems property value. Optional. NavigationProperty/Containment; navigation property to the activity's historyItems.
func (m *UserActivity) SetHistoryItems(value []ActivityHistoryItem)() {
    if m != nil {
        m.historyItems = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Set by the server. DateTime in UTC when the object was modified on the server.
func (m *UserActivity) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetStatus sets the status property value. Set by the server. A status code used to identify valid objects. Values: active, updated, deleted, ignored.
func (m *UserActivity) SetStatus(value *Status)() {
    if m != nil {
        m.status = value
    }
}
// SetUserTimezone sets the userTimezone property value. Optional. The timezone in which the user's device used to generate the activity was located at activity creation time; values supplied as Olson IDs in order to support cross-platform representation.
func (m *UserActivity) SetUserTimezone(value *string)() {
    if m != nil {
        m.userTimezone = value
    }
}
// SetVisualElements sets the visualElements property value. 
func (m *UserActivity) SetVisualElements(value *VisualInfo)() {
    if m != nil {
        m.visualElements = value
    }
}
