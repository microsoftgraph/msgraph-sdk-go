package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserConsentRequest provides operations to manage the identityGovernance singleton.
type UserConsentRequest struct {
    Request
    // Approval decisions associated with a request.
    approval Approvalable;
    // The user's justification for requiring access to the app. Supports $filter (eq only) and $orderby.
    reason *string;
}
// NewUserConsentRequest instantiates a new userConsentRequest and sets the default values.
func NewUserConsentRequest()(*UserConsentRequest) {
    m := &UserConsentRequest{
        Request: *NewRequest(),
    }
    return m
}
// CreateUserConsentRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserConsentRequestFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUserConsentRequest(), nil
}
// GetApproval gets the approval property value. Approval decisions associated with a request.
func (m *UserConsentRequest) GetApproval()(Approvalable) {
    if m == nil {
        return nil
    } else {
        return m.approval
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserConsentRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Request.GetFieldDeserializers()
    res["approval"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateApprovalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApproval(val.(Approvalable))
        }
        return nil
    }
    res["reason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    return res
}
// GetReason gets the reason property value. The user's justification for requiring access to the app. Supports $filter (eq only) and $orderby.
func (m *UserConsentRequest) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
func (m *UserConsentRequest) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserConsentRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Request.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("approval", m.GetApproval())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reason", m.GetReason())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApproval sets the approval property value. Approval decisions associated with a request.
func (m *UserConsentRequest) SetApproval(value Approvalable)() {
    if m != nil {
        m.approval = value
    }
}
// SetReason sets the reason property value. The user's justification for requiring access to the app. Supports $filter (eq only) and $orderby.
func (m *UserConsentRequest) SetReason(value *string)() {
    if m != nil {
        m.reason = value
    }
}
