package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AccessReviewScheduleSettings struct {
    additionalData map[string]interface{};
    applyActions []AccessReviewApplyAction;
    autoApplyDecisionsEnabled *bool;
    defaultDecision *string;
    defaultDecisionEnabled *bool;
    instanceDurationInDays *int32;
    justificationRequiredOnApproval *bool;
    mailNotificationsEnabled *bool;
    recommendationsEnabled *bool;
    recurrence *PatternedRecurrence;
    reminderNotificationsEnabled *bool;
}
func NewAccessReviewScheduleSettings()(*AccessReviewScheduleSettings) {
    m := &AccessReviewScheduleSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AccessReviewScheduleSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AccessReviewScheduleSettings) GetApplyActions()([]AccessReviewApplyAction) {
    if m == nil {
        return nil
    } else {
        return m.applyActions
    }
}
func (m *AccessReviewScheduleSettings) GetAutoApplyDecisionsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autoApplyDecisionsEnabled
    }
}
func (m *AccessReviewScheduleSettings) GetDefaultDecision()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultDecision
    }
}
func (m *AccessReviewScheduleSettings) GetDefaultDecisionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defaultDecisionEnabled
    }
}
func (m *AccessReviewScheduleSettings) GetInstanceDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.instanceDurationInDays
    }
}
func (m *AccessReviewScheduleSettings) GetJustificationRequiredOnApproval()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.justificationRequiredOnApproval
    }
}
func (m *AccessReviewScheduleSettings) GetMailNotificationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.mailNotificationsEnabled
    }
}
func (m *AccessReviewScheduleSettings) GetRecommendationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.recommendationsEnabled
    }
}
func (m *AccessReviewScheduleSettings) GetRecurrence()(*PatternedRecurrence) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
func (m *AccessReviewScheduleSettings) GetReminderNotificationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.reminderNotificationsEnabled
    }
}
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
func (m *AccessReviewScheduleSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AccessReviewScheduleSettings) SetApplyActions(value []AccessReviewApplyAction)() {
    m.applyActions = value
}
func (m *AccessReviewScheduleSettings) SetAutoApplyDecisionsEnabled(value *bool)() {
    m.autoApplyDecisionsEnabled = value
}
func (m *AccessReviewScheduleSettings) SetDefaultDecision(value *string)() {
    m.defaultDecision = value
}
func (m *AccessReviewScheduleSettings) SetDefaultDecisionEnabled(value *bool)() {
    m.defaultDecisionEnabled = value
}
func (m *AccessReviewScheduleSettings) SetInstanceDurationInDays(value *int32)() {
    m.instanceDurationInDays = value
}
func (m *AccessReviewScheduleSettings) SetJustificationRequiredOnApproval(value *bool)() {
    m.justificationRequiredOnApproval = value
}
func (m *AccessReviewScheduleSettings) SetMailNotificationsEnabled(value *bool)() {
    m.mailNotificationsEnabled = value
}
func (m *AccessReviewScheduleSettings) SetRecommendationsEnabled(value *bool)() {
    m.recommendationsEnabled = value
}
func (m *AccessReviewScheduleSettings) SetRecurrence(value *PatternedRecurrence)() {
    m.recurrence = value
}
func (m *AccessReviewScheduleSettings) SetReminderNotificationsEnabled(value *bool)() {
    m.reminderNotificationsEnabled = value
}
