package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserConsentRequest struct {
    Request
    // Approval decisions associated with a request.
    approval *Approval;
    // The user's justification for requiring access to the app. Supports $filter (eq only) and $orderby.
    reason *string;
}
// Instantiates a new userConsentRequest and sets the default values.
func NewUserConsentRequest()(*UserConsentRequest) {
    m := &UserConsentRequest{
        Request: *NewRequest(),
    }
    return m
}
// Gets the approval property value. Approval decisions associated with a request.
func (m *UserConsentRequest) GetApproval()(*Approval) {
    if m == nil {
        return nil
    } else {
        return m.approval
    }
}
// Gets the reason property value. The user's justification for requiring access to the app. Supports $filter (eq only) and $orderby.
func (m *UserConsentRequest) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the approval property value. Approval decisions associated with a request.
// Parameters:
//  - value : Value to set for the approval property.
func (m *UserConsentRequest) SetApproval(value *Approval)() {
    m.approval = value
}
// Sets the reason property value. The user's justification for requiring access to the app. Supports $filter (eq only) and $orderby.
// Parameters:
//  - value : Value to set for the reason property.
func (m *UserConsentRequest) SetReason(value *string)() {
    m.reason = value
}
