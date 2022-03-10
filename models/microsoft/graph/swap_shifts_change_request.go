package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SwapShiftsChangeRequest provides operations to manage the educationRoot singleton.
type SwapShiftsChangeRequest struct {
    OfferShiftRequest
    // ShiftId for the recipient user with whom the request is to swap.
    recipientShiftId *string;
}
// NewSwapShiftsChangeRequest instantiates a new swapShiftsChangeRequest and sets the default values.
func NewSwapShiftsChangeRequest()(*SwapShiftsChangeRequest) {
    m := &SwapShiftsChangeRequest{
        OfferShiftRequest: *NewOfferShiftRequest(),
    }
    return m
}
// CreateSwapShiftsChangeRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSwapShiftsChangeRequestFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSwapShiftsChangeRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
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
// GetRecipientShiftId gets the recipientShiftId property value. ShiftId for the recipient user with whom the request is to swap.
func (m *SwapShiftsChangeRequest) GetRecipientShiftId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recipientShiftId
    }
}
func (m *SwapShiftsChangeRequest) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetRecipientShiftId sets the recipientShiftId property value. ShiftId for the recipient user with whom the request is to swap.
func (m *SwapShiftsChangeRequest) SetRecipientShiftId(value *string)() {
    if m != nil {
        m.recipientShiftId = value
    }
}
