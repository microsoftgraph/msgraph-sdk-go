package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SwapShiftsChangeRequest struct {
    OfferShiftRequest
    recipientShiftId *string;
}
func NewSwapShiftsChangeRequest()(*SwapShiftsChangeRequest) {
    m := &SwapShiftsChangeRequest{
        OfferShiftRequest: *NewOfferShiftRequest(),
    }
    return m
}
func (m *SwapShiftsChangeRequest) GetRecipientShiftId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recipientShiftId
    }
}
func (m *SwapShiftsChangeRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OfferShiftRequest.GetFieldDeserializers()
    res["recipientShiftId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRecipientShiftId(val)
        return nil
    }
    return res
}
func (m *SwapShiftsChangeRequest) IsNil()(bool) {
    return m == nil
}
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
func (m *SwapShiftsChangeRequest) SetRecipientShiftId(value *string)() {
    m.recipientShiftId = value
}
