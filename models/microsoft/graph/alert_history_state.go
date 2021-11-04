package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AlertHistoryState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The Application ID of the calling application that submitted an update (PATCH) to the alert. The appId should be extracted from the auth token and not entered manually by the calling application.
    appId *string;
    // UPN of user the alert was assigned to (note: alert.assignedTo only stores the last value/UPN).
    assignedTo *string;
    // Comment entered by signed-in user.
    comments []string;
    // Analyst feedback on the alert in this update. Possible values are: unknown, truePositive, falsePositive, benignPositive.
    feedback *AlertFeedback;
    // Alert status value (if updated). Possible values are: unknown, newAlert, inProgress, resolved, dismissed.
    status *AlertStatus;
    // Date and time of the alert update. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    updatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // UPN of the signed-in user that updated the alert (taken from the bearer token - if in user/delegated auth mode).
    user *string;
}
// Instantiates a new alertHistoryState and sets the default values.
func NewAlertHistoryState()(*AlertHistoryState) {
    m := &AlertHistoryState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlertHistoryState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the appId property value. The Application ID of the calling application that submitted an update (PATCH) to the alert. The appId should be extracted from the auth token and not entered manually by the calling application.
func (m *AlertHistoryState) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// Gets the assignedTo property value. UPN of user the alert was assigned to (note: alert.assignedTo only stores the last value/UPN).
func (m *AlertHistoryState) GetAssignedTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignedTo
    }
}
// Gets the comments property value. Comment entered by signed-in user.
func (m *AlertHistoryState) GetComments()([]string) {
    if m == nil {
        return nil
    } else {
        return m.comments
    }
}
// Gets the feedback property value. Analyst feedback on the alert in this update. Possible values are: unknown, truePositive, falsePositive, benignPositive.
func (m *AlertHistoryState) GetFeedback()(*AlertFeedback) {
    if m == nil {
        return nil
    } else {
        return m.feedback
    }
}
// Gets the status property value. Alert status value (if updated). Possible values are: unknown, newAlert, inProgress, resolved, dismissed.
func (m *AlertHistoryState) GetStatus()(*AlertStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the updatedDateTime property value. Date and time of the alert update. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AlertHistoryState) GetUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.updatedDateTime
    }
}
// Gets the user property value. UPN of the signed-in user that updated the alert (taken from the bearer token - if in user/delegated auth mode).
func (m *AlertHistoryState) GetUser()(*string) {
    if m == nil {
        return nil
    } else {
        return m.user
    }
}
// The deserialization information for the current model
func (m *AlertHistoryState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppId(val)
        return nil
    }
    res["assignedTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAssignedTo(val)
        return nil
    }
    res["comments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetComments(res)
        return nil
    }
    res["feedback"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertFeedback)
        if err != nil {
            return err
        }
        cast := val.(AlertFeedback)
        m.SetFeedback(&cast)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertStatus)
        if err != nil {
            return err
        }
        cast := val.(AlertStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["updatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetUpdatedDateTime(val)
        return nil
    }
    res["user"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUser(val)
        return nil
    }
    return res
}
func (m *AlertHistoryState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AlertHistoryState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("assignedTo", m.GetAssignedTo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("comments", m.GetComments())
        if err != nil {
            return err
        }
    }
    if m.GetFeedback() != nil {
        cast := m.GetFeedback().String()
        err := writer.WriteStringValue("feedback", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updatedDateTime", m.GetUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("user", m.GetUser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AlertHistoryState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the appId property value. The Application ID of the calling application that submitted an update (PATCH) to the alert. The appId should be extracted from the auth token and not entered manually by the calling application.
// Parameters:
//  - value : Value to set for the appId property.
func (m *AlertHistoryState) SetAppId(value *string)() {
    m.appId = value
}
// Sets the assignedTo property value. UPN of user the alert was assigned to (note: alert.assignedTo only stores the last value/UPN).
// Parameters:
//  - value : Value to set for the assignedTo property.
func (m *AlertHistoryState) SetAssignedTo(value *string)() {
    m.assignedTo = value
}
// Sets the comments property value. Comment entered by signed-in user.
// Parameters:
//  - value : Value to set for the comments property.
func (m *AlertHistoryState) SetComments(value []string)() {
    m.comments = value
}
// Sets the feedback property value. Analyst feedback on the alert in this update. Possible values are: unknown, truePositive, falsePositive, benignPositive.
// Parameters:
//  - value : Value to set for the feedback property.
func (m *AlertHistoryState) SetFeedback(value *AlertFeedback)() {
    m.feedback = value
}
// Sets the status property value. Alert status value (if updated). Possible values are: unknown, newAlert, inProgress, resolved, dismissed.
// Parameters:
//  - value : Value to set for the status property.
func (m *AlertHistoryState) SetStatus(value *AlertStatus)() {
    m.status = value
}
// Sets the updatedDateTime property value. Date and time of the alert update. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the updatedDateTime property.
func (m *AlertHistoryState) SetUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedDateTime = value
}
// Sets the user property value. UPN of the signed-in user that updated the alert (taken from the bearer token - if in user/delegated auth mode).
// Parameters:
//  - value : Value to set for the user property.
func (m *AlertHistoryState) SetUser(value *string)() {
    m.user = value
}
