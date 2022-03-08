package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApprovalStage provides operations to manage the identityGovernance singleton.
type ApprovalStage struct {
    Entity
    // Indicates whether the stage is assigned to the calling user to review. Read-only.
    assignedToMe *bool;
    // The label provided by the policy creator to identify an approval stage. Read-only.
    displayName *string;
    // The justification associated with the approval stage decision.
    justification *string;
    // The identifier of the reviewer. Read-only.
    reviewedBy Identityable;
    // The date and time when a decision was recorded. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    reviewedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The result of this approval record. Possible values include: NotReviewed, Approved, Denied.
    reviewResult *string;
    // The stage status. Possible values: InProgress, Initializing, Completed, Expired. Read-only.
    status *string;
}
// NewApprovalStage instantiates a new approvalStage and sets the default values.
func NewApprovalStage()(*ApprovalStage) {
    m := &ApprovalStage{
        Entity: *NewEntity(),
    }
    return m
}
// CreateApprovalStageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApprovalStageFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewApprovalStage(), nil
}
// GetAssignedToMe gets the assignedToMe property value. Indicates whether the stage is assigned to the calling user to review. Read-only.
func (m *ApprovalStage) GetAssignedToMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.assignedToMe
    }
}
// GetDisplayName gets the displayName property value. The label provided by the policy creator to identify an approval stage. Read-only.
func (m *ApprovalStage) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApprovalStage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedToMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedToMe(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["justification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
        }
        return nil
    }
    res["reviewedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewedBy(val.(Identityable))
        }
        return nil
    }
    res["reviewedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewedDateTime(val)
        }
        return nil
    }
    res["reviewResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewResult(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetJustification gets the justification property value. The justification associated with the approval stage decision.
func (m *ApprovalStage) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
// GetReviewedBy gets the reviewedBy property value. The identifier of the reviewer. Read-only.
func (m *ApprovalStage) GetReviewedBy()(Identityable) {
    if m == nil {
        return nil
    } else {
        return m.reviewedBy
    }
}
// GetReviewedDateTime gets the reviewedDateTime property value. The date and time when a decision was recorded. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *ApprovalStage) GetReviewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.reviewedDateTime
    }
}
// GetReviewResult gets the reviewResult property value. The result of this approval record. Possible values include: NotReviewed, Approved, Denied.
func (m *ApprovalStage) GetReviewResult()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reviewResult
    }
}
// GetStatus gets the status property value. The stage status. Possible values: InProgress, Initializing, Completed, Expired. Read-only.
func (m *ApprovalStage) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *ApprovalStage) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAssignedToMe sets the assignedToMe property value. Indicates whether the stage is assigned to the calling user to review. Read-only.
func (m *ApprovalStage) SetAssignedToMe(value *bool)() {
    if m != nil {
        m.assignedToMe = value
    }
}
// SetDisplayName sets the displayName property value. The label provided by the policy creator to identify an approval stage. Read-only.
func (m *ApprovalStage) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetJustification sets the justification property value. The justification associated with the approval stage decision.
func (m *ApprovalStage) SetJustification(value *string)() {
    if m != nil {
        m.justification = value
    }
}
// SetReviewedBy sets the reviewedBy property value. The identifier of the reviewer. Read-only.
func (m *ApprovalStage) SetReviewedBy(value Identityable)() {
    if m != nil {
        m.reviewedBy = value
    }
}
// SetReviewedDateTime sets the reviewedDateTime property value. The date and time when a decision was recorded. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *ApprovalStage) SetReviewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.reviewedDateTime = value
    }
}
// SetReviewResult sets the reviewResult property value. The result of this approval record. Possible values include: NotReviewed, Approved, Denied.
func (m *ApprovalStage) SetReviewResult(value *string)() {
    if m != nil {
        m.reviewResult = value
    }
}
// SetStatus sets the status property value. The stage status. Possible values: InProgress, Initializing, Completed, Expired. Read-only.
func (m *ApprovalStage) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
