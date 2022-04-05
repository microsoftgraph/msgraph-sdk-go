package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageApprovalStage 
type AccessPackageApprovalStage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The number of days that a request can be pending a response before it is automatically denied.
    durationBeforeAutomaticDenial *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration;
    // If escalation is required, the time a request can be pending a response from a primary approver.
    durationBeforeEscalation *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration;
    // If escalation is enabled and the primary approvers do not respond before the escalation time, the escalationApprovers are the users who will be asked to approve requests.
    escalationApprovers []SubjectSetable;
    // The subjects, typically users, who are the fallback escalation approvers.
    fallbackEscalationApprovers []SubjectSetable;
    // The subjects, typically users, who are the fallback primary approvers.
    fallbackPrimaryApprovers []SubjectSetable;
    // Indicates whether the approver is required to provide a justification for approving a request.
    isApproverJustificationRequired *bool;
    // If true, then one or more escalationApprovers are configured in this approval stage.
    isEscalationEnabled *bool;
    // The subjects, typically users, who will be asked to approve requests. A collection of singleUser, groupMembers, requestorManager, internalSponsors or externalSponsors.
    primaryApprovers []SubjectSetable;
}
// NewAccessPackageApprovalStage instantiates a new accessPackageApprovalStage and sets the default values.
func NewAccessPackageApprovalStage()(*AccessPackageApprovalStage) {
    m := &AccessPackageApprovalStage{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAccessPackageApprovalStageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageApprovalStageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageApprovalStage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageApprovalStage) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDurationBeforeAutomaticDenial gets the durationBeforeAutomaticDenial property value. The number of days that a request can be pending a response before it is automatically denied.
func (m *AccessPackageApprovalStage) GetDurationBeforeAutomaticDenial()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.durationBeforeAutomaticDenial
    }
}
// GetDurationBeforeEscalation gets the durationBeforeEscalation property value. If escalation is required, the time a request can be pending a response from a primary approver.
func (m *AccessPackageApprovalStage) GetDurationBeforeEscalation()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.durationBeforeEscalation
    }
}
// GetEscalationApprovers gets the escalationApprovers property value. If escalation is enabled and the primary approvers do not respond before the escalation time, the escalationApprovers are the users who will be asked to approve requests.
func (m *AccessPackageApprovalStage) GetEscalationApprovers()([]SubjectSetable) {
    if m == nil {
        return nil
    } else {
        return m.escalationApprovers
    }
}
// GetFallbackEscalationApprovers gets the fallbackEscalationApprovers property value. The subjects, typically users, who are the fallback escalation approvers.
func (m *AccessPackageApprovalStage) GetFallbackEscalationApprovers()([]SubjectSetable) {
    if m == nil {
        return nil
    } else {
        return m.fallbackEscalationApprovers
    }
}
// GetFallbackPrimaryApprovers gets the fallbackPrimaryApprovers property value. The subjects, typically users, who are the fallback primary approvers.
func (m *AccessPackageApprovalStage) GetFallbackPrimaryApprovers()([]SubjectSetable) {
    if m == nil {
        return nil
    } else {
        return m.fallbackPrimaryApprovers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageApprovalStage) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["durationBeforeAutomaticDenial"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationBeforeAutomaticDenial(val)
        }
        return nil
    }
    res["durationBeforeEscalation"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationBeforeEscalation(val)
        }
        return nil
    }
    res["escalationApprovers"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectSetable, len(val))
            for i, v := range val {
                res[i] = v.(SubjectSetable)
            }
            m.SetEscalationApprovers(res)
        }
        return nil
    }
    res["fallbackEscalationApprovers"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectSetable, len(val))
            for i, v := range val {
                res[i] = v.(SubjectSetable)
            }
            m.SetFallbackEscalationApprovers(res)
        }
        return nil
    }
    res["fallbackPrimaryApprovers"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectSetable, len(val))
            for i, v := range val {
                res[i] = v.(SubjectSetable)
            }
            m.SetFallbackPrimaryApprovers(res)
        }
        return nil
    }
    res["isApproverJustificationRequired"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApproverJustificationRequired(val)
        }
        return nil
    }
    res["isEscalationEnabled"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEscalationEnabled(val)
        }
        return nil
    }
    res["primaryApprovers"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectSetable, len(val))
            for i, v := range val {
                res[i] = v.(SubjectSetable)
            }
            m.SetPrimaryApprovers(res)
        }
        return nil
    }
    return res
}
// GetIsApproverJustificationRequired gets the isApproverJustificationRequired property value. Indicates whether the approver is required to provide a justification for approving a request.
func (m *AccessPackageApprovalStage) GetIsApproverJustificationRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApproverJustificationRequired
    }
}
// GetIsEscalationEnabled gets the isEscalationEnabled property value. If true, then one or more escalationApprovers are configured in this approval stage.
func (m *AccessPackageApprovalStage) GetIsEscalationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEscalationEnabled
    }
}
// GetPrimaryApprovers gets the primaryApprovers property value. The subjects, typically users, who will be asked to approve requests. A collection of singleUser, groupMembers, requestorManager, internalSponsors or externalSponsors.
func (m *AccessPackageApprovalStage) GetPrimaryApprovers()([]SubjectSetable) {
    if m == nil {
        return nil
    } else {
        return m.primaryApprovers
    }
}
// Serialize serializes information the current object
func (m *AccessPackageApprovalStage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteISODurationValue("durationBeforeAutomaticDenial", m.GetDurationBeforeAutomaticDenial())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("durationBeforeEscalation", m.GetDurationBeforeEscalation())
        if err != nil {
            return err
        }
    }
    if m.GetEscalationApprovers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEscalationApprovers()))
        for i, v := range m.GetEscalationApprovers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("escalationApprovers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFallbackEscalationApprovers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFallbackEscalationApprovers()))
        for i, v := range m.GetFallbackEscalationApprovers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("fallbackEscalationApprovers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFallbackPrimaryApprovers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFallbackPrimaryApprovers()))
        for i, v := range m.GetFallbackPrimaryApprovers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("fallbackPrimaryApprovers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApproverJustificationRequired", m.GetIsApproverJustificationRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEscalationEnabled", m.GetIsEscalationEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetPrimaryApprovers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPrimaryApprovers()))
        for i, v := range m.GetPrimaryApprovers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("primaryApprovers", cast)
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
func (m *AccessPackageApprovalStage) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDurationBeforeAutomaticDenial sets the durationBeforeAutomaticDenial property value. The number of days that a request can be pending a response before it is automatically denied.
func (m *AccessPackageApprovalStage) SetDurationBeforeAutomaticDenial(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.durationBeforeAutomaticDenial = value
    }
}
// SetDurationBeforeEscalation sets the durationBeforeEscalation property value. If escalation is required, the time a request can be pending a response from a primary approver.
func (m *AccessPackageApprovalStage) SetDurationBeforeEscalation(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.durationBeforeEscalation = value
    }
}
// SetEscalationApprovers sets the escalationApprovers property value. If escalation is enabled and the primary approvers do not respond before the escalation time, the escalationApprovers are the users who will be asked to approve requests.
func (m *AccessPackageApprovalStage) SetEscalationApprovers(value []SubjectSetable)() {
    if m != nil {
        m.escalationApprovers = value
    }
}
// SetFallbackEscalationApprovers sets the fallbackEscalationApprovers property value. The subjects, typically users, who are the fallback escalation approvers.
func (m *AccessPackageApprovalStage) SetFallbackEscalationApprovers(value []SubjectSetable)() {
    if m != nil {
        m.fallbackEscalationApprovers = value
    }
}
// SetFallbackPrimaryApprovers sets the fallbackPrimaryApprovers property value. The subjects, typically users, who are the fallback primary approvers.
func (m *AccessPackageApprovalStage) SetFallbackPrimaryApprovers(value []SubjectSetable)() {
    if m != nil {
        m.fallbackPrimaryApprovers = value
    }
}
// SetIsApproverJustificationRequired sets the isApproverJustificationRequired property value. Indicates whether the approver is required to provide a justification for approving a request.
func (m *AccessPackageApprovalStage) SetIsApproverJustificationRequired(value *bool)() {
    if m != nil {
        m.isApproverJustificationRequired = value
    }
}
// SetIsEscalationEnabled sets the isEscalationEnabled property value. If true, then one or more escalationApprovers are configured in this approval stage.
func (m *AccessPackageApprovalStage) SetIsEscalationEnabled(value *bool)() {
    if m != nil {
        m.isEscalationEnabled = value
    }
}
// SetPrimaryApprovers sets the primaryApprovers property value. The subjects, typically users, who will be asked to approve requests. A collection of singleUser, groupMembers, requestorManager, internalSponsors or externalSponsors.
func (m *AccessPackageApprovalStage) SetPrimaryApprovers(value []SubjectSetable)() {
    if m != nil {
        m.primaryApprovers = value
    }
}
