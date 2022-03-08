package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExcludeTarget provides operations to manage the authenticationMethodsPolicy singleton.
type ExcludeTarget struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The object identifier of an Azure Active Directory user or group.
    id *string;
    // The type of the authentication method target. Possible values are: user, group, unknownFutureValue.
    targetType *AuthenticationMethodTargetType;
}
// NewExcludeTarget instantiates a new excludeTarget and sets the default values.
func NewExcludeTarget()(*ExcludeTarget) {
    m := &ExcludeTarget{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateExcludeTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExcludeTargetFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewExcludeTarget(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExcludeTarget) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExcludeTarget) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["targetType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodTargetType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetType(val.(*AuthenticationMethodTargetType))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The object identifier of an Azure Active Directory user or group.
func (m *ExcludeTarget) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetTargetType gets the targetType property value. The type of the authentication method target. Possible values are: user, group, unknownFutureValue.
func (m *ExcludeTarget) GetTargetType()(*AuthenticationMethodTargetType) {
    if m == nil {
        return nil
    } else {
        return m.targetType
    }
}
func (m *ExcludeTarget) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExcludeTarget) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetTargetType() != nil {
        cast := (*m.GetTargetType()).String()
        err := writer.WriteStringValue("targetType", &cast)
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
func (m *ExcludeTarget) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetId sets the id property value. The object identifier of an Azure Active Directory user or group.
func (m *ExcludeTarget) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetTargetType sets the targetType property value. The type of the authentication method target. Possible values are: user, group, unknownFutureValue.
func (m *ExcludeTarget) SetTargetType(value *AuthenticationMethodTargetType)() {
    if m != nil {
        m.targetType = value
    }
}
