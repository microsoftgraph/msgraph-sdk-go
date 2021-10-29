package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessReviewScheduleSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Optional field. Describes the  actions to take once a review is complete. There are two types that are currently supported: removeAccessApplyAction (default) and disableAndDeleteUserApplyAction. Field only needs to be specified in the case of disableAndDeleteUserApplyAction.
    applyActions []AccessReviewApplyAction;
    // Indicates whether decisions are automatically applied. When set to false, an admin must apply the decisions manually once the reviewer completes the access review. When set to true, decisions are applied automatically after the access review instance duration ends, whether or not the reviewers have responded. Default value is false.
    autoApplyDecisionsEnabled *bool;
    // Decision chosen if defaultDecisionEnabled is true. Can be one of Approve, Deny, or Recommendation.
    defaultDecision *string;
    // Indicates whether the default decision is enabled or disabled when reviewers do not respond. Default value is false.
    defaultDecisionEnabled *bool;
    // Duration of each recurrence of review (accessReviewInstance) in number of days.
    instanceDurationInDays *int32;
    // Indicates whether reviewers are required to provide justification with their decision. Default value is false.
    justificationRequiredOnApproval *bool;
    // Indicates whether emails are enabled or disabled. Default value is false.
    mailNotificationsEnabled *bool;
    // Indicates whether decision recommendations are enabled or disabled.
    recommendationsEnabled *bool;
    // Detailed settings for recurrence using the standard Outlook recurrence object. Only weekly and absoluteMonthly on recurrencePattern are supported. Use the property startDate on recurrenceRange to determine the day the review starts.
    recurrence *PatternedRecurrence;
    // Indicates whether reminders are enabled or disabled. Default value is false.
    reminderNotificationsEnabled *bool;
}
// Instantiates a new accessReviewScheduleSettings and sets the default values.
func NewAccessReviewScheduleSettings()(*AccessReviewScheduleSettings) {
    m := &AccessReviewScheduleSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewScheduleSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the applyActions property value. Optional field. Describes the  actions to take once a review is complete. There are two types that are currently supported: removeAccessApplyAction (default) and disableAndDeleteUserApplyAction. Field only needs to be specified in the case of disableAndDeleteUserApplyAction.
func (m *AccessReviewScheduleSettings) GetApplyActions()([]AccessReviewApplyAction) {
    if m == nil {
        return nil
    } else {
        return m.applyActions
    }
}
// Gets the autoApplyDecisionsEnabled property value. Indicates whether decisions are automatically applied. When set to false, an admin must apply the decisions manually once the reviewer completes the access review. When set to true, decisions are applied automatically after the access review instance duration ends, whether or not the reviewers have responded. Default value is false.
func (m *AccessReviewScheduleSettings) GetAutoApplyDecisionsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autoApplyDecisionsEnabled
    }
}
// Gets the defaultDecision property value. Decision chosen if defaultDecisionEnabled is true. Can be one of Approve, Deny, or Recommendation.
func (m *AccessReviewScheduleSettings) GetDefaultDecision()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultDecision
    }
}
// Gets the defaultDecisionEnabled property value. Indicates whether the default decision is enabled or disabled when reviewers do not respond. Default value is false.
func (m *AccessReviewScheduleSettings) GetDefaultDecisionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defaultDecisionEnabled
    }
}
// Gets the instanceDurationInDays property value. Duration of each recurrence of review (accessReviewInstance) in number of days.
func (m *AccessReviewScheduleSettings) GetInstanceDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.instanceDurationInDays
    }
}
// Gets the justificationRequiredOnApproval property value. Indicates whether reviewers are required to provide justification with their decision. Default value is false.
func (m *AccessReviewScheduleSettings) GetJustificationRequiredOnApproval()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.justificationRequiredOnApproval
    }
}
// Gets the mailNotificationsEnabled property value. Indicates whether emails are enabled or disabled. Default value is false.
func (m *AccessReviewScheduleSettings) GetMailNotificationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.mailNotificationsEnabled
    }
}
// Gets the recommendationsEnabled property value. Indicates whether decision recommendations are enabled or disabled.
func (m *AccessReviewScheduleSettings) GetRecommendationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.recommendationsEnabled
    }
}
// Gets the recurrence property value. Detailed settings for recurrence using the standard Outlook recurrence object. Only weekly and absoluteMonthly on recurrencePattern are supported. Use the property startDate on recurrenceRange to determine the day the review starts.
func (m *AccessReviewScheduleSettings) GetRecurrence()(*PatternedRecurrence) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
// Gets the reminderNotificationsEnabled property value. Indicates whether reminders are enabled or disabled. Default value is false.
func (m *AccessReviewScheduleSettings) GetReminderNotificationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.reminderNotificationsEnabled
    }
}
// The deserialization information for the current model
func (m *AccessReviewScheduleSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applyActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewApplyAction() })
        if err != nil {
            return err
        }
        res := make([]AccessReviewApplyAction, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessReviewApplyAction))
        }
        m.SetApplyActions(res)
        return nil
    }
    res["autoApplyDecisionsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAutoApplyDecisionsEnabled(val)
        return nil
    }
    res["defaultDecision"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefaultDecision(val)
        return nil
    }
    res["defaultDecisionEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDefaultDecisionEnabled(val)
        return nil
    }
    res["instanceDurationInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetInstanceDurationInDays(val)
        return nil
    }
    res["justificationRequiredOnApproval"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetJustificationRequiredOnApproval(val)
        return nil
    }
    res["mailNotificationsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetMailNotificationsEnabled(val)
        return nil
    }
    res["recommendationsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRecommendationsEnabled(val)
        return nil
    }
    res["recurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPatternedRecurrence() })
        if err != nil {
            return err
        }
        m.SetRecurrence(val.(*PatternedRecurrence))
        return nil
    }
    res["reminderNotificationsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetReminderNotificationsEnabled(val)
        return nil
    }
    return res
}
func (m *AccessReviewScheduleSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AccessReviewScheduleSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetApplyActions()))
        for i, v := range m.GetApplyActions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("applyActions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("autoApplyDecisionsEnabled", m.GetAutoApplyDecisionsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("defaultDecision", m.GetDefaultDecision())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("defaultDecisionEnabled", m.GetDefaultDecisionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("instanceDurationInDays", m.GetInstanceDurationInDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("justificationRequiredOnApproval", m.GetJustificationRequiredOnApproval())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("mailNotificationsEnabled", m.GetMailNotificationsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("recommendationsEnabled", m.GetRecommendationsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("recurrence", m.GetRecurrence())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("reminderNotificationsEnabled", m.GetReminderNotificationsEnabled())
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
func (m *AccessReviewScheduleSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the applyActions property value. Optional field. Describes the  actions to take once a review is complete. There are two types that are currently supported: removeAccessApplyAction (default) and disableAndDeleteUserApplyAction. Field only needs to be specified in the case of disableAndDeleteUserApplyAction.
// Parameters:
//  - value : Value to set for the applyActions property.
func (m *AccessReviewScheduleSettings) SetApplyActions(value []AccessReviewApplyAction)() {
    m.applyActions = value
}
// Sets the autoApplyDecisionsEnabled property value. Indicates whether decisions are automatically applied. When set to false, an admin must apply the decisions manually once the reviewer completes the access review. When set to true, decisions are applied automatically after the access review instance duration ends, whether or not the reviewers have responded. Default value is false.
// Parameters:
//  - value : Value to set for the autoApplyDecisionsEnabled property.
func (m *AccessReviewScheduleSettings) SetAutoApplyDecisionsEnabled(value *bool)() {
    m.autoApplyDecisionsEnabled = value
}
// Sets the defaultDecision property value. Decision chosen if defaultDecisionEnabled is true. Can be one of Approve, Deny, or Recommendation.
// Parameters:
//  - value : Value to set for the defaultDecision property.
func (m *AccessReviewScheduleSettings) SetDefaultDecision(value *string)() {
    m.defaultDecision = value
}
// Sets the defaultDecisionEnabled property value. Indicates whether the default decision is enabled or disabled when reviewers do not respond. Default value is false.
// Parameters:
//  - value : Value to set for the defaultDecisionEnabled property.
func (m *AccessReviewScheduleSettings) SetDefaultDecisionEnabled(value *bool)() {
    m.defaultDecisionEnabled = value
}
// Sets the instanceDurationInDays property value. Duration of each recurrence of review (accessReviewInstance) in number of days.
// Parameters:
//  - value : Value to set for the instanceDurationInDays property.
func (m *AccessReviewScheduleSettings) SetInstanceDurationInDays(value *int32)() {
    m.instanceDurationInDays = value
}
// Sets the justificationRequiredOnApproval property value. Indicates whether reviewers are required to provide justification with their decision. Default value is false.
// Parameters:
//  - value : Value to set for the justificationRequiredOnApproval property.
func (m *AccessReviewScheduleSettings) SetJustificationRequiredOnApproval(value *bool)() {
    m.justificationRequiredOnApproval = value
}
// Sets the mailNotificationsEnabled property value. Indicates whether emails are enabled or disabled. Default value is false.
// Parameters:
//  - value : Value to set for the mailNotificationsEnabled property.
func (m *AccessReviewScheduleSettings) SetMailNotificationsEnabled(value *bool)() {
    m.mailNotificationsEnabled = value
}
// Sets the recommendationsEnabled property value. Indicates whether decision recommendations are enabled or disabled.
// Parameters:
//  - value : Value to set for the recommendationsEnabled property.
func (m *AccessReviewScheduleSettings) SetRecommendationsEnabled(value *bool)() {
    m.recommendationsEnabled = value
}
// Sets the recurrence property value. Detailed settings for recurrence using the standard Outlook recurrence object. Only weekly and absoluteMonthly on recurrencePattern are supported. Use the property startDate on recurrenceRange to determine the day the review starts.
// Parameters:
//  - value : Value to set for the recurrence property.
func (m *AccessReviewScheduleSettings) SetRecurrence(value *PatternedRecurrence)() {
    m.recurrence = value
}
// Sets the reminderNotificationsEnabled property value. Indicates whether reminders are enabled or disabled. Default value is false.
// Parameters:
//  - value : Value to set for the reminderNotificationsEnabled property.
func (m *AccessReviewScheduleSettings) SetReminderNotificationsEnabled(value *bool)() {
    m.reminderNotificationsEnabled = value
}
