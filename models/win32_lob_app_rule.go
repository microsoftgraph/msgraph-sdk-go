package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppRule a base complex type to store the detection or requirement rule data for a Win32 LOB app.
type Win32LobAppRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The rule type indicating the purpose of the rule. Possible values are: detection, requirement.
    ruleType *Win32LobAppRuleType
    // The type property
    type_escaped *string
}
// NewWin32LobAppRule instantiates a new win32LobAppRule and sets the default values.
func NewWin32LobAppRule()(*Win32LobAppRule) {
    m := &Win32LobAppRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWin32LobAppRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.win32LobAppFileSystemRule":
                        return NewWin32LobAppFileSystemRule(), nil
                    case "#microsoft.graph.win32LobAppPowerShellScriptRule":
                        return NewWin32LobAppPowerShellScriptRule(), nil
                    case "#microsoft.graph.win32LobAppProductCodeRule":
                        return NewWin32LobAppProductCodeRule(), nil
                    case "#microsoft.graph.win32LobAppRegistryRule":
                        return NewWin32LobAppRegistryRule(), nil
                }
            }
        }
    }
    return NewWin32LobAppRule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Win32LobAppRule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ruleType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppRuleType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleType(val.(*Win32LobAppRuleType))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetRuleType gets the ruleType property value. The rule type indicating the purpose of the rule. Possible values are: detection, requirement.
func (m *Win32LobAppRule) GetRuleType()(*Win32LobAppRuleType) {
    if m == nil {
        return nil
    } else {
        return m.ruleType
    }
}
// GetType gets the type property value. The type property
func (m *Win32LobAppRule) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *Win32LobAppRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRuleType() != nil {
        cast := (*m.GetRuleType()).String()
        err := writer.WriteStringValue("ruleType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *Win32LobAppRule) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRuleType sets the ruleType property value. The rule type indicating the purpose of the rule. Possible values are: detection, requirement.
func (m *Win32LobAppRule) SetRuleType(value *Win32LobAppRuleType)() {
    if m != nil {
        m.ruleType = value
    }
}
// SetType sets the type property value. The type property
func (m *Win32LobAppRule) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
