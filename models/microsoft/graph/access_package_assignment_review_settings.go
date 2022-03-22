package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentReviewSettings 
type AccessPackageAssignmentReviewSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    expirationBehavior *AccessReviewExpirationBehavior;
    // 
    fallbackReviewers []SubjectSetable;
    // 
    isEnabled *bool;
    // 
    isRecommendationEnabled *bool;
    // 
    isReviewerJustificationRequired *bool;
    // 
    isSelfReview *bool;
    // 
    primaryReviewers []SubjectSetable;
    // 
    schedule EntitlementManagementScheduleable;
}
// NewAccessPackageAssignmentReviewSettings instantiates a new accessPackageAssignmentReviewSettings and sets the default values.
func NewAccessPackageAssignmentReviewSettings()(*AccessPackageAssignmentReviewSettings) {
    m := &AccessPackageAssignmentReviewSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAccessPackageAssignmentReviewSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentReviewSettingsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAccessPackageAssignmentReviewSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageAssignmentReviewSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExpirationBehavior gets the expirationBehavior property value. 
func (m *AccessPackageAssignmentReviewSettings) GetExpirationBehavior()(*AccessReviewExpirationBehavior) {
    if m == nil {
        return nil
    } else {
        return m.expirationBehavior
    }
}
// GetFallbackReviewers gets the fallbackReviewers property value. 
func (m *AccessPackageAssignmentReviewSettings) GetFallbackReviewers()([]SubjectSetable) {
    if m == nil {
        return nil
    } else {
        return m.fallbackReviewers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentReviewSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["expirationBehavior"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessReviewExpirationBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationBehavior(val.(*AccessReviewExpirationBehavior))
        }
        return nil
    }
    res["fallbackReviewers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectSetable, len(val))
            for i, v := range val {
                res[i] = v.(SubjectSetable)
            }
            m.SetFallbackReviewers(res)
        }
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["isRecommendationEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRecommendationEnabled(val)
        }
        return nil
    }
    res["isReviewerJustificationRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReviewerJustificationRequired(val)
        }
        return nil
    }
    res["isSelfReview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSelfReview(val)
        }
        return nil
    }
    res["primaryReviewers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectSetable, len(val))
            for i, v := range val {
                res[i] = v.(SubjectSetable)
            }
            m.SetPrimaryReviewers(res)
        }
        return nil
    }
    res["schedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntitlementManagementScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(EntitlementManagementScheduleable))
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. 
func (m *AccessPackageAssignmentReviewSettings) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetIsRecommendationEnabled gets the isRecommendationEnabled property value. 
func (m *AccessPackageAssignmentReviewSettings) GetIsRecommendationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRecommendationEnabled
    }
}
// GetIsReviewerJustificationRequired gets the isReviewerJustificationRequired property value. 
func (m *AccessPackageAssignmentReviewSettings) GetIsReviewerJustificationRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReviewerJustificationRequired
    }
}
// GetIsSelfReview gets the isSelfReview property value. 
func (m *AccessPackageAssignmentReviewSettings) GetIsSelfReview()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSelfReview
    }
}
// GetPrimaryReviewers gets the primaryReviewers property value. 
func (m *AccessPackageAssignmentReviewSettings) GetPrimaryReviewers()([]SubjectSetable) {
    if m == nil {
        return nil
    } else {
        return m.primaryReviewers
    }
}
// GetSchedule gets the schedule property value. 
func (m *AccessPackageAssignmentReviewSettings) GetSchedule()(EntitlementManagementScheduleable) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentReviewSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetExpirationBehavior() != nil {
        cast := (*m.GetExpirationBehavior()).String()
        err := writer.WriteStringValue("expirationBehavior", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFallbackReviewers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFallbackReviewers()))
        for i, v := range m.GetFallbackReviewers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("fallbackReviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRecommendationEnabled", m.GetIsRecommendationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isReviewerJustificationRequired", m.GetIsReviewerJustificationRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSelfReview", m.GetIsSelfReview())
        if err != nil {
            return err
        }
    }
    if m.GetPrimaryReviewers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPrimaryReviewers()))
        for i, v := range m.GetPrimaryReviewers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("primaryReviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("schedule", m.GetSchedule())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageAssignmentReviewSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExpirationBehavior sets the expirationBehavior property value. 
func (m *AccessPackageAssignmentReviewSettings) SetExpirationBehavior(value *AccessReviewExpirationBehavior)() {
    if m != nil {
        m.expirationBehavior = value
    }
}
// SetFallbackReviewers sets the fallbackReviewers property value. 
func (m *AccessPackageAssignmentReviewSettings) SetFallbackReviewers(value []SubjectSetable)() {
    if m != nil {
        m.fallbackReviewers = value
    }
}
// SetIsEnabled sets the isEnabled property value. 
func (m *AccessPackageAssignmentReviewSettings) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetIsRecommendationEnabled sets the isRecommendationEnabled property value. 
func (m *AccessPackageAssignmentReviewSettings) SetIsRecommendationEnabled(value *bool)() {
    if m != nil {
        m.isRecommendationEnabled = value
    }
}
// SetIsReviewerJustificationRequired sets the isReviewerJustificationRequired property value. 
func (m *AccessPackageAssignmentReviewSettings) SetIsReviewerJustificationRequired(value *bool)() {
    if m != nil {
        m.isReviewerJustificationRequired = value
    }
}
// SetIsSelfReview sets the isSelfReview property value. 
func (m *AccessPackageAssignmentReviewSettings) SetIsSelfReview(value *bool)() {
    if m != nil {
        m.isSelfReview = value
    }
}
// SetPrimaryReviewers sets the primaryReviewers property value. 
func (m *AccessPackageAssignmentReviewSettings) SetPrimaryReviewers(value []SubjectSetable)() {
    if m != nil {
        m.primaryReviewers = value
    }
}
// SetSchedule sets the schedule property value. 
func (m *AccessPackageAssignmentReviewSettings) SetSchedule(value EntitlementManagementScheduleable)() {
    if m != nil {
        m.schedule = value
    }
}
