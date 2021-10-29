package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ApprovalStage struct {
    Entity
    // Indicates whether the stage is assigned to the calling user to review. Read-only.
    assignedToMe *bool;
    // The label provided by the policy creator to identify an approval stage. Read-only.
    displayName *string;
    // The justification associated with the approval stage decision.
    justification *string;
    // The identifier of the reviewer. Read-only.
    reviewedBy *Identity;
    // The date and time when a decision was recorded. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    reviewedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The result of this approval record. Possible values include: NotReviewed, Approved, Denied.
    reviewResult *string;
    // The stage status. Possible values: InProgress, Initializing, Completed, Expired. Read-only.
    status *string;
}
// Instantiates a new approvalStage and sets the default values.
func NewApprovalStage()(*ApprovalStage) {
    m := &ApprovalStage{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignedToMe property value. Indicates whether the stage is assigned to the calling user to review. Read-only.
func (m *ApprovalStage) GetAssignedToMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.assignedToMe
    }
}
// Gets the displayName property value. The label provided by the policy creator to identify an approval stage. Read-only.
func (m *ApprovalStage) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the justification property value. The justification associated with the approval stage decision.
func (m *ApprovalStage) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
// Gets the reviewedBy property value. The identifier of the reviewer. Read-only.
func (m *ApprovalStage) GetReviewedBy()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.reviewedBy
    }
}
// Gets the reviewedDateTime property value. The date and time when a decision was recorded. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *ApprovalStage) GetReviewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.reviewedDateTime
    }
}
// Gets the reviewResult property value. The result of this approval record. Possible values include: NotReviewed, Approved, Denied.
func (m *ApprovalStage) GetReviewResult()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reviewResult
    }
}
// Gets the status property value. The stage status. Possible values: InProgress, Initializing, Completed, Expired. Read-only.
func (m *ApprovalStage) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *ApprovalStage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedToMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAssignedToMe(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["justification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJustification(val)
        return nil
    }
    res["reviewedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentity() })
        if err != nil {
            return err
        }
        m.SetReviewedBy(val.(*Identity))
        return nil
    }
    res["reviewedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetReviewedDateTime(val)
        return nil
    }
    res["reviewResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReviewResult(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    return res
}
func (m *ApprovalStage) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ApprovalStage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("assignedToMe", m.GetAssignedToMe())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reviewedBy", m.GetReviewedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("reviewedDateTime", m.GetReviewedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reviewResult", m.GetReviewResult())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignedToMe property value. Indicates whether the stage is assigned to the calling user to review. Read-only.
// Parameters:
//  - value : Value to set for the assignedToMe property.
func (m *ApprovalStage) SetAssignedToMe(value *bool)() {
    m.assignedToMe = value
}
// Sets the displayName property value. The label provided by the policy creator to identify an approval stage. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ApprovalStage) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the justification property value. The justification associated with the approval stage decision.
// Parameters:
//  - value : Value to set for the justification property.
func (m *ApprovalStage) SetJustification(value *string)() {
    m.justification = value
}
// Sets the reviewedBy property value. The identifier of the reviewer. Read-only.
// Parameters:
//  - value : Value to set for the reviewedBy property.
func (m *ApprovalStage) SetReviewedBy(value *Identity)() {
    m.reviewedBy = value
}
// Sets the reviewedDateTime property value. The date and time when a decision was recorded. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// Parameters:
//  - value : Value to set for the reviewedDateTime property.
func (m *ApprovalStage) SetReviewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reviewedDateTime = value
}
// Sets the reviewResult property value. The result of this approval record. Possible values include: NotReviewed, Approved, Denied.
// Parameters:
//  - value : Value to set for the reviewResult property.
func (m *ApprovalStage) SetReviewResult(value *string)() {
    m.reviewResult = value
}
// Sets the status property value. The stage status. Possible values: InProgress, Initializing, Completed, Expired. Read-only.
// Parameters:
//  - value : Value to set for the status property.
func (m *ApprovalStage) SetStatus(value *string)() {
    m.status = value
}
