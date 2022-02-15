package instances

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/delta"
)

// InstancesResponse 
type InstancesResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    nextLink *string;
    // 
    value []ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec.Event;
}
// NewInstancesResponse instantiates a new instancesResponse and sets the default values.
func NewInstancesResponse()(*InstancesResponse) {
    m := &InstancesResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InstancesResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetNextLink gets the @odata.nextLink property value. 
func (m *InstancesResponse) GetNextLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nextLink
    }
}
// GetValue gets the value property value. 
func (m *InstancesResponse) GetValue()([]ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec.Event) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InstancesResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["@odata.nextLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextLink(val)
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec.NewEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec.Event, len(val))
            for i, v := range val {
                res[i] = *(v.(*ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec.Event))
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
func (m *InstancesResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *InstancesResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.nextLink", m.GetNextLink())
        if err != nil {
            return err
        }
    }
    if m.GetValue() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("value", cast)
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
func (m *InstancesResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetNextLink sets the @odata.nextLink property value. 
func (m *InstancesResponse) SetNextLink(value *string)() {
    if m != nil {
        m.nextLink = value
    }
}
// SetValue sets the value property value. 
func (m *InstancesResponse) SetValue(value []ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec.Event)() {
    if m != nil {
        m.value = value
    }
}
