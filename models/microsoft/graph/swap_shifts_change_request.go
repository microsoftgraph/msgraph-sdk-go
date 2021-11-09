package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SwapShiftsChangeRequest struct {
    OfferShiftRequest
    // ShiftId for the recipient user with whom the request is to swap.
    recipientShiftId *string;
}
// Instantiates a new swapShiftsChangeRequest and sets the default values.
func NewSwapShiftsChangeRequest()(*SwapShiftsChangeRequest) {
    m := &SwapShiftsChangeRequest{
        OfferShiftRequest: *NewOfferShiftRequest(),
    }
    return m
}
// Gets the recipientShiftId property value. ShiftId for the recipient user with whom the request is to swap.
func (m *SwapShiftsChangeRequest) GetRecipientShiftId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recipientShiftId
    }
}
// The deserialization information for the current model
func (m *SwapShiftsChangeRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OfferShiftRequest.GetFieldDeserializers()
    res["recipientShiftId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientShiftId(val)
        }
        return nil
    }
    return res
}
func (m *SwapShiftsChangeRequest) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SwapShiftsChangeRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OfferShiftRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("recipientShiftId", m.GetRecipientShiftId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the recipientShiftId property value. ShiftId for the recipient user with whom the request is to swap.
// Parameters:
//  - value : Value to set for the recipientShiftId property.
func (m *SwapShiftsChangeRequest) SetRecipientShiftId(value *string)() {
    m.recipientShiftId = value
}
