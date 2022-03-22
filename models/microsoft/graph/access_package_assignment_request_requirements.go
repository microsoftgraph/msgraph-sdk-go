package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentRequestRequirements 
type AccessPackageAssignmentRequestRequirements struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether the requestor is allowed to set a custom schedule.
    allowCustomAssignmentSchedule *bool;
    // Indicates whether a request to add must be approved by an approver.
    isApprovalRequiredForAdd *bool;
    // Indicates whether a request to update must be approved by an approver.
    isApprovalRequiredForUpdate *bool;
    // The description of the policy that the user is trying to request access using.
    policyDescription *string;
    // The display name of the policy that the user is trying to request access using.
    policyDisplayName *string;
    // The identifier of the policy that these requirements are associated with. This identifier can be used when creating a new assignment request.
    policyId *string;
    // Schedule restrictions enforced, if any.
    schedule EntitlementManagementScheduleable;
}
// NewAccessPackageAssignmentRequestRequirements instantiates a new accessPackageAssignmentRequestRequirements and sets the default values.
func NewAccessPackageAssignmentRequestRequirements()(*AccessPackageAssignmentRequestRequirements) {
    m := &AccessPackageAssignmentRequestRequirements{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAccessPackageAssignmentRequestRequirementsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentRequestRequirementsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAccessPackageAssignmentRequestRequirements(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageAssignmentRequestRequirements) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowCustomAssignmentSchedule gets the allowCustomAssignmentSchedule property value. Indicates whether the requestor is allowed to set a custom schedule.
func (m *AccessPackageAssignmentRequestRequirements) GetAllowCustomAssignmentSchedule()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowCustomAssignmentSchedule
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentRequestRequirements) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowCustomAssignmentSchedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowCustomAssignmentSchedule(val)
        }
        return nil
    }
    res["isApprovalRequiredForAdd"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApprovalRequiredForAdd(val)
        }
        return nil
    }
    res["isApprovalRequiredForUpdate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApprovalRequiredForUpdate(val)
        }
        return nil
    }
    res["policyDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyDescription(val)
        }
        return nil
    }
    res["policyDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyDisplayName(val)
        }
        return nil
    }
    res["policyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyId(val)
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
// GetIsApprovalRequiredForAdd gets the isApprovalRequiredForAdd property value. Indicates whether a request to add must be approved by an approver.
func (m *AccessPackageAssignmentRequestRequirements) GetIsApprovalRequiredForAdd()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApprovalRequiredForAdd
    }
}
// GetIsApprovalRequiredForUpdate gets the isApprovalRequiredForUpdate property value. Indicates whether a request to update must be approved by an approver.
func (m *AccessPackageAssignmentRequestRequirements) GetIsApprovalRequiredForUpdate()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApprovalRequiredForUpdate
    }
}
// GetPolicyDescription gets the policyDescription property value. The description of the policy that the user is trying to request access using.
func (m *AccessPackageAssignmentRequestRequirements) GetPolicyDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyDescription
    }
}
// GetPolicyDisplayName gets the policyDisplayName property value. The display name of the policy that the user is trying to request access using.
func (m *AccessPackageAssignmentRequestRequirements) GetPolicyDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyDisplayName
    }
}
// GetPolicyId gets the policyId property value. The identifier of the policy that these requirements are associated with. This identifier can be used when creating a new assignment request.
func (m *AccessPackageAssignmentRequestRequirements) GetPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyId
    }
}
// GetSchedule gets the schedule property value. Schedule restrictions enforced, if any.
func (m *AccessPackageAssignmentRequestRequirements) GetSchedule()(EntitlementManagementScheduleable) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentRequestRequirements) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowCustomAssignmentSchedule", m.GetAllowCustomAssignmentSchedule())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApprovalRequiredForAdd", m.GetIsApprovalRequiredForAdd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApprovalRequiredForUpdate", m.GetIsApprovalRequiredForUpdate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyDescription", m.GetPolicyDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyDisplayName", m.GetPolicyDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyId", m.GetPolicyId())
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
func (m *AccessPackageAssignmentRequestRequirements) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowCustomAssignmentSchedule sets the allowCustomAssignmentSchedule property value. Indicates whether the requestor is allowed to set a custom schedule.
func (m *AccessPackageAssignmentRequestRequirements) SetAllowCustomAssignmentSchedule(value *bool)() {
    if m != nil {
        m.allowCustomAssignmentSchedule = value
    }
}
// SetIsApprovalRequiredForAdd sets the isApprovalRequiredForAdd property value. Indicates whether a request to add must be approved by an approver.
func (m *AccessPackageAssignmentRequestRequirements) SetIsApprovalRequiredForAdd(value *bool)() {
    if m != nil {
        m.isApprovalRequiredForAdd = value
    }
}
// SetIsApprovalRequiredForUpdate sets the isApprovalRequiredForUpdate property value. Indicates whether a request to update must be approved by an approver.
func (m *AccessPackageAssignmentRequestRequirements) SetIsApprovalRequiredForUpdate(value *bool)() {
    if m != nil {
        m.isApprovalRequiredForUpdate = value
    }
}
// SetPolicyDescription sets the policyDescription property value. The description of the policy that the user is trying to request access using.
func (m *AccessPackageAssignmentRequestRequirements) SetPolicyDescription(value *string)() {
    if m != nil {
        m.policyDescription = value
    }
}
// SetPolicyDisplayName sets the policyDisplayName property value. The display name of the policy that the user is trying to request access using.
func (m *AccessPackageAssignmentRequestRequirements) SetPolicyDisplayName(value *string)() {
    if m != nil {
        m.policyDisplayName = value
    }
}
// SetPolicyId sets the policyId property value. The identifier of the policy that these requirements are associated with. This identifier can be used when creating a new assignment request.
func (m *AccessPackageAssignmentRequestRequirements) SetPolicyId(value *string)() {
    if m != nil {
        m.policyId = value
    }
}
// SetSchedule sets the schedule property value. Schedule restrictions enforced, if any.
func (m *AccessPackageAssignmentRequestRequirements) SetSchedule(value EntitlementManagementScheduleable)() {
    if m != nil {
        m.schedule = value
    }
}
