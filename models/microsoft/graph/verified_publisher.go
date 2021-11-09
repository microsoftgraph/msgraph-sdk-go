package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type VerifiedPublisher struct {
    // The timestamp when the verified publisher was first added or most recently updated.
    addedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The verified publisher name from the app publisher's Partner Center account.
    displayName *string;
    // The ID of the verified publisher from the app publisher's Partner Center account.
    verifiedPublisherId *string;
}
// Instantiates a new verifiedPublisher and sets the default values.
func NewVerifiedPublisher()(*VerifiedPublisher) {
    m := &VerifiedPublisher{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the addedDateTime property value. The timestamp when the verified publisher was first added or most recently updated.
func (m *VerifiedPublisher) GetAddedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.addedDateTime
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VerifiedPublisher) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the displayName property value. The verified publisher name from the app publisher's Partner Center account.
func (m *VerifiedPublisher) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the verifiedPublisherId property value. The ID of the verified publisher from the app publisher's Partner Center account.
func (m *VerifiedPublisher) GetVerifiedPublisherId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.verifiedPublisherId
    }
}
// The deserialization information for the current model
func (m *VerifiedPublisher) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["addedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddedDateTime(val)
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
    res["verifiedPublisherId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedPublisherId(val)
        }
        return nil
    }
    return res
}
func (m *VerifiedPublisher) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *VerifiedPublisher) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("addedDateTime", m.GetAddedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("verifiedPublisherId", m.GetVerifiedPublisherId())
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
// Sets the addedDateTime property value. The timestamp when the verified publisher was first added or most recently updated.
// Parameters:
//  - value : Value to set for the addedDateTime property.
func (m *VerifiedPublisher) SetAddedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.addedDateTime = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *VerifiedPublisher) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the displayName property value. The verified publisher name from the app publisher's Partner Center account.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *VerifiedPublisher) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the verifiedPublisherId property value. The ID of the verified publisher from the app publisher's Partner Center account.
// Parameters:
//  - value : Value to set for the verifiedPublisherId property.
func (m *VerifiedPublisher) SetVerifiedPublisherId(value *string)() {
    m.verifiedPublisherId = value
}
