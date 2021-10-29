package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UriClickSecurityState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    clickAction *string;
    // 
    clickDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    id *string;
    // 
    sourceId *string;
    // 
    uriDomain *string;
    // 
    verdict *string;
}
// Instantiates a new uriClickSecurityState and sets the default values.
func NewUriClickSecurityState()(*UriClickSecurityState) {
    m := &UriClickSecurityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UriClickSecurityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the clickAction property value. 
func (m *UriClickSecurityState) GetClickAction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clickAction
    }
}
// Gets the clickDateTime property value. 
func (m *UriClickSecurityState) GetClickDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.clickDateTime
    }
}
// Gets the id property value. 
func (m *UriClickSecurityState) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the sourceId property value. 
func (m *UriClickSecurityState) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
// Gets the uriDomain property value. 
func (m *UriClickSecurityState) GetUriDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uriDomain
    }
}
// Gets the verdict property value. 
func (m *UriClickSecurityState) GetVerdict()(*string) {
    if m == nil {
        return nil
    } else {
        return m.verdict
    }
}
// The deserialization information for the current model
func (m *UriClickSecurityState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["clickAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClickAction(val)
        return nil
    }
    res["clickDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetClickDateTime(val)
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["sourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSourceId(val)
        return nil
    }
    res["uriDomain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUriDomain(val)
        return nil
    }
    res["verdict"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVerdict(val)
        return nil
    }
    return res
}
func (m *UriClickSecurityState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UriClickSecurityState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clickAction", m.GetClickAction())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("clickDateTime", m.GetClickDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceId", m.GetSourceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uriDomain", m.GetUriDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("verdict", m.GetVerdict())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *UriClickSecurityState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the clickAction property value. 
// Parameters:
//  - value : Value to set for the clickAction property.
func (m *UriClickSecurityState) SetClickAction(value *string)() {
    m.clickAction = value
}
// Sets the clickDateTime property value. 
// Parameters:
//  - value : Value to set for the clickDateTime property.
func (m *UriClickSecurityState) SetClickDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.clickDateTime = value
}
// Sets the id property value. 
// Parameters:
//  - value : Value to set for the id property.
func (m *UriClickSecurityState) SetId(value *string)() {
    m.id = value
}
// Sets the sourceId property value. 
// Parameters:
//  - value : Value to set for the sourceId property.
func (m *UriClickSecurityState) SetSourceId(value *string)() {
    m.sourceId = value
}
// Sets the uriDomain property value. 
// Parameters:
//  - value : Value to set for the uriDomain property.
func (m *UriClickSecurityState) SetUriDomain(value *string)() {
    m.uriDomain = value
}
// Sets the verdict property value. 
// Parameters:
//  - value : Value to set for the verdict property.
func (m *UriClickSecurityState) SetVerdict(value *string)() {
    m.verdict = value
}
