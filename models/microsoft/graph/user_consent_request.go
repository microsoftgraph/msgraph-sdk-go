package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserConsentRequest 
type UserConsentRequest struct {
    Request
    // Approval decisions associated with a request.
    approval *Approval;
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
// GetApproval gets the approval property value. Approval decisions associated with a request.
func (m *UserConsentRequest) GetApproval()(*Approval) {
    if m == nil {
        return nil
    } else {
        return m.approval
    }
}
// GetReason gets the reason property value. The user's justification for requiring access to the app. Supports $filter (eq only) and $orderby.
func (m *UserConsentRequest) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserConsentRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Request.GetFieldDeserializers()
    res["approval"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApproval() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApproval(val.(*Approval))
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
func (m *UserConsentRequest) SetApproval(value *Approval)() {
    m.approval = value
}
// SetReason sets the reason property value. The user's justification for requiring access to the app. Supports $filter (eq only) and $orderby.
func (m *UserConsentRequest) SetReason(value *string)() {
    m.reason = value
}
