package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type LocationConstraint struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The client requests the service to include in the response a meeting location for the meeting. If this is true and all the resources are busy, findMeetingTimes will not return any meeting time suggestions. If this is false and all the resources are busy, findMeetingTimes would still look for meeting times without locations.
    isRequired *bool;
    // Constraint information for one or more locations that the client requests for the meeting.
    locations []LocationConstraintItem;
    // The client requests the service to suggest one or more meeting locations.
    suggestLocation *bool;
}
// Instantiates a new locationConstraint and sets the default values.
func NewLocationConstraint()(*LocationConstraint) {
    m := &LocationConstraint{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LocationConstraint) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the isRequired property value. The client requests the service to include in the response a meeting location for the meeting. If this is true and all the resources are busy, findMeetingTimes will not return any meeting time suggestions. If this is false and all the resources are busy, findMeetingTimes would still look for meeting times without locations.
func (m *LocationConstraint) GetIsRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRequired
    }
}
// Gets the locations property value. Constraint information for one or more locations that the client requests for the meeting.
func (m *LocationConstraint) GetLocations()([]LocationConstraintItem) {
    if m == nil {
        return nil
    } else {
        return m.locations
    }
}
// Gets the suggestLocation property value. The client requests the service to suggest one or more meeting locations.
func (m *LocationConstraint) GetSuggestLocation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.suggestLocation
    }
}
// The deserialization information for the current model
func (m *LocationConstraint) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRequired(val)
        }
        return nil
    }
    res["locations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocationConstraintItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LocationConstraintItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*LocationConstraintItem))
            }
            m.SetLocations(res)
        }
        return nil
    }
    res["suggestLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuggestLocation(val)
        }
        return nil
    }
    return res
}
func (m *LocationConstraint) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *LocationConstraint) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isRequired", m.GetIsRequired())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLocations()))
        for i, v := range m.GetLocations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("locations", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("suggestLocation", m.GetSuggestLocation())
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
func (m *LocationConstraint) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the isRequired property value. The client requests the service to include in the response a meeting location for the meeting. If this is true and all the resources are busy, findMeetingTimes will not return any meeting time suggestions. If this is false and all the resources are busy, findMeetingTimes would still look for meeting times without locations.
// Parameters:
//  - value : Value to set for the isRequired property.
func (m *LocationConstraint) SetIsRequired(value *bool)() {
    m.isRequired = value
}
// Sets the locations property value. Constraint information for one or more locations that the client requests for the meeting.
// Parameters:
//  - value : Value to set for the locations property.
func (m *LocationConstraint) SetLocations(value []LocationConstraintItem)() {
    m.locations = value
}
// Sets the suggestLocation property value. The client requests the service to suggest one or more meeting locations.
// Parameters:
//  - value : Value to set for the suggestLocation property.
func (m *LocationConstraint) SetSuggestLocation(value *bool)() {
    m.suggestLocation = value
}
