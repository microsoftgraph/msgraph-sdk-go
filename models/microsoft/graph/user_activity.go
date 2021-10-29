package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new userActivity and sets the default values.
func NewUserActivity()(*UserActivity) {
    m := &UserActivity{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activationUrl property value. Required. URL used to launch the activity in the best native experience represented by the appId. Might launch a web-based app if no native app exists.
func (m *UserActivity) GetActivationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activationUrl
    }
}
// Gets the activitySourceHost property value. Required. URL for the domain representing the cross-platform identity mapping for the app. Mapping is stored either as a JSON file hosted on the domain or configurable via Windows Dev Center. The JSON file is named cross-platform-app-identifiers and is hosted at root of your HTTPS domain, either at the top level domain or include a sub domain. For example: https://contoso.com or https://myapp.contoso.com but NOT https://myapp.contoso.com/somepath. You must have a unique file and domain (or sub domain) per cross-platform app identity. For example, a separate file and domain is needed for Word vs. PowerPoint.
func (m *UserActivity) GetActivitySourceHost()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activitySourceHost
    }
}
// Gets the appActivityId property value. Required. The unique activity ID in the context of the app - supplied by caller and immutable thereafter.
func (m *UserActivity) GetAppActivityId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appActivityId
    }
}
// Gets the appDisplayName property value. Optional. Short text description of the app used to generate the activity for use in cases when the app is not installed on the user’s local device.
func (m *UserActivity) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// Gets the contentInfo property value. Optional. A custom piece of data - JSON-LD extensible description of content according to schema.org syntax.
func (m *UserActivity) GetContentInfo()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.contentInfo
    }
}
// Gets the contentUrl property value. Optional. Used in the event the content can be rendered outside of a native or web-based app experience (for example, a pointer to an item in an RSS feed).
func (m *UserActivity) GetContentUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentUrl
    }
}
// Gets the createdDateTime property value. Set by the server. DateTime in UTC when the object was created on the server.
func (m *UserActivity) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the expirationDateTime property value. Set by the server. DateTime in UTC when the object expired on the server.
func (m *UserActivity) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the fallbackUrl property value. Optional. URL used to launch the activity in a web-based app, if available.
func (m *UserActivity) GetFallbackUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fallbackUrl
    }
}
// Gets the historyItems property value. Optional. NavigationProperty/Containment; navigation property to the activity's historyItems.
func (m *UserActivity) GetHistoryItems()([]ActivityHistoryItem) {
    if m == nil {
        return nil
    } else {
        return m.historyItems
    }
}
// Gets the lastModifiedDateTime property value. Set by the server. DateTime in UTC when the object was modified on the server.
func (m *UserActivity) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the status property value. Set by the server. A status code used to identify valid objects. Values: active, updated, deleted, ignored.
func (m *UserActivity) GetStatus()(*Status) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the userTimezone property value. Optional. The timezone in which the user's device used to generate the activity was located at activity creation time; values supplied as Olson IDs in order to support cross-platform representation.
func (m *UserActivity) GetUserTimezone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userTimezone
    }
}
// Gets the visualElements property value. 
func (m *UserActivity) GetVisualElements()(*VisualInfo) {
    if m == nil {
        return nil
    } else {
        return m.visualElements
    }
}
// The deserialization information for the current model
func (m *UserActivity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activationUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActivationUrl(val)
        return nil
    }
    res["activitySourceHost"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActivitySourceHost(val)
        return nil
    }
    res["appActivityId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppActivityId(val)
        return nil
    }
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppDisplayName(val)
        return nil
    }
    res["contentInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetContentInfo(val.(*Json))
        return nil
    }
    res["contentUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContentUrl(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
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
    res["fallbackUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFallbackUrl(val)
        return nil
    }
    res["historyItems"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewActivityHistoryItem() })
        if err != nil {
            return err
        }
        res := make([]ActivityHistoryItem, len(val))
        for i, v := range val {
            res[i] = *(v.(*ActivityHistoryItem))
        }
        m.SetHistoryItems(res)
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
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseStatus)
        if err != nil {
            return err
        }
        cast := val.(Status)
        m.SetStatus(&cast)
        return nil
    }
    res["userTimezone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserTimezone(val)
        return nil
    }
    res["visualElements"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVisualInfo() })
        if err != nil {
            return err
        }
        m.SetVisualElements(val.(*VisualInfo))
        return nil
    }
    return res
}
func (m *UserActivity) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the activationUrl property value. Required. URL used to launch the activity in the best native experience represented by the appId. Might launch a web-based app if no native app exists.
// Parameters:
//  - value : Value to set for the activationUrl property.
func (m *UserActivity) SetActivationUrl(value *string)() {
    m.activationUrl = value
}
// Sets the activitySourceHost property value. Required. URL for the domain representing the cross-platform identity mapping for the app. Mapping is stored either as a JSON file hosted on the domain or configurable via Windows Dev Center. The JSON file is named cross-platform-app-identifiers and is hosted at root of your HTTPS domain, either at the top level domain or include a sub domain. For example: https://contoso.com or https://myapp.contoso.com but NOT https://myapp.contoso.com/somepath. You must have a unique file and domain (or sub domain) per cross-platform app identity. For example, a separate file and domain is needed for Word vs. PowerPoint.
// Parameters:
//  - value : Value to set for the activitySourceHost property.
func (m *UserActivity) SetActivitySourceHost(value *string)() {
    m.activitySourceHost = value
}
// Sets the appActivityId property value. Required. The unique activity ID in the context of the app - supplied by caller and immutable thereafter.
// Parameters:
//  - value : Value to set for the appActivityId property.
func (m *UserActivity) SetAppActivityId(value *string)() {
    m.appActivityId = value
}
// Sets the appDisplayName property value. Optional. Short text description of the app used to generate the activity for use in cases when the app is not installed on the user’s local device.
// Parameters:
//  - value : Value to set for the appDisplayName property.
func (m *UserActivity) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// Sets the contentInfo property value. Optional. A custom piece of data - JSON-LD extensible description of content according to schema.org syntax.
// Parameters:
//  - value : Value to set for the contentInfo property.
func (m *UserActivity) SetContentInfo(value *Json)() {
    m.contentInfo = value
}
// Sets the contentUrl property value. Optional. Used in the event the content can be rendered outside of a native or web-based app experience (for example, a pointer to an item in an RSS feed).
// Parameters:
//  - value : Value to set for the contentUrl property.
func (m *UserActivity) SetContentUrl(value *string)() {
    m.contentUrl = value
}
// Sets the createdDateTime property value. Set by the server. DateTime in UTC when the object was created on the server.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *UserActivity) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the expirationDateTime property value. Set by the server. DateTime in UTC when the object expired on the server.
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *UserActivity) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the fallbackUrl property value. Optional. URL used to launch the activity in a web-based app, if available.
// Parameters:
//  - value : Value to set for the fallbackUrl property.
func (m *UserActivity) SetFallbackUrl(value *string)() {
    m.fallbackUrl = value
}
// Sets the historyItems property value. Optional. NavigationProperty/Containment; navigation property to the activity's historyItems.
// Parameters:
//  - value : Value to set for the historyItems property.
func (m *UserActivity) SetHistoryItems(value []ActivityHistoryItem)() {
    m.historyItems = value
}
// Sets the lastModifiedDateTime property value. Set by the server. DateTime in UTC when the object was modified on the server.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *UserActivity) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the status property value. Set by the server. A status code used to identify valid objects. Values: active, updated, deleted, ignored.
// Parameters:
//  - value : Value to set for the status property.
func (m *UserActivity) SetStatus(value *Status)() {
    m.status = value
}
// Sets the userTimezone property value. Optional. The timezone in which the user's device used to generate the activity was located at activity creation time; values supplied as Olson IDs in order to support cross-platform representation.
// Parameters:
//  - value : Value to set for the userTimezone property.
func (m *UserActivity) SetUserTimezone(value *string)() {
    m.userTimezone = value
}
// Sets the visualElements property value. 
// Parameters:
//  - value : Value to set for the visualElements property.
func (m *UserActivity) SetVisualElements(value *VisualInfo)() {
    m.visualElements = value
}
