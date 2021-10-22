package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/callrecords"
)

type NetworkInfo struct {
    additionalData map[string]interface{};
    bandwidthLowEventRatio *float32;
    basicServiceSetIdentifier *string;
    connectionType *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.NetworkConnectionType;
    delayEventRatio *float32;
    dnsSuffix *string;
    ipAddress *string;
    linkSpeed *int64;
    macAddress *string;
    port *int32;
    receivedQualityEventRatio *float32;
    reflexiveIPAddress *string;
    relayIPAddress *string;
    relayPort *int32;
    sentQualityEventRatio *float32;
    subnet *string;
    wifiBand *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiBand;
    wifiBatteryCharge *int32;
    wifiChannel *int32;
    wifiMicrosoftDriver *string;
    wifiMicrosoftDriverVersion *string;
    wifiRadioType *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiRadioType;
    wifiSignalStrength *int32;
    wifiVendorDriver *string;
    wifiVendorDriverVersion *string;
}
func NewNetworkInfo()(*NetworkInfo) {
    m := &NetworkInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *NetworkInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *NetworkInfo) GetBandwidthLowEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.bandwidthLowEventRatio
    }
}
func (m *NetworkInfo) GetBasicServiceSetIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.basicServiceSetIdentifier
    }
}
func (m *NetworkInfo) GetConnectionType()(*i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.NetworkConnectionType) {
    if m == nil {
        return nil
    } else {
        return m.connectionType
    }
}
func (m *NetworkInfo) GetDelayEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.delayEventRatio
    }
}
func (m *NetworkInfo) GetDnsSuffix()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dnsSuffix
    }
}
func (m *NetworkInfo) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
func (m *NetworkInfo) GetLinkSpeed()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.linkSpeed
    }
}
func (m *NetworkInfo) GetMacAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.macAddress
    }
}
func (m *NetworkInfo) GetPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.port
    }
}
func (m *NetworkInfo) GetReceivedQualityEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.receivedQualityEventRatio
    }
}
func (m *NetworkInfo) GetReflexiveIPAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reflexiveIPAddress
    }
}
func (m *NetworkInfo) GetRelayIPAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.relayIPAddress
    }
}
func (m *NetworkInfo) GetRelayPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.relayPort
    }
}
func (m *NetworkInfo) GetSentQualityEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.sentQualityEventRatio
    }
}
func (m *NetworkInfo) GetSubnet()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subnet
    }
}
func (m *NetworkInfo) GetWifiBand()(*i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiBand) {
    if m == nil {
        return nil
    } else {
        return m.wifiBand
    }
}
func (m *NetworkInfo) GetWifiBatteryCharge()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.wifiBatteryCharge
    }
}
func (m *NetworkInfo) GetWifiChannel()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.wifiChannel
    }
}
func (m *NetworkInfo) GetWifiMicrosoftDriver()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiMicrosoftDriver
    }
}
func (m *NetworkInfo) GetWifiMicrosoftDriverVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiMicrosoftDriverVersion
    }
}
func (m *NetworkInfo) GetWifiRadioType()(*i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiRadioType) {
    if m == nil {
        return nil
    } else {
        return m.wifiRadioType
    }
}
func (m *NetworkInfo) GetWifiSignalStrength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.wifiSignalStrength
    }
}
func (m *NetworkInfo) GetWifiVendorDriver()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiVendorDriver
    }
}
func (m *NetworkInfo) GetWifiVendorDriverVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiVendorDriverVersion
    }
}
func (m *NetworkInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bandwidthLowEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetBandwidthLowEventRatio(val)
        return nil
    }
    res["basicServiceSetIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBasicServiceSetIdentifier(val)
        return nil
    }
    res["connectionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.ParseNetworkConnectionType)
        if err != nil {
            return err
        }
        cast := val.(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.NetworkConnectionType)
        m.SetConnectionType(&cast)
        return nil
    }
    res["delayEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetDelayEventRatio(val)
        return nil
    }
    res["dnsSuffix"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDnsSuffix(val)
        return nil
    }
    res["ipAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIpAddress(val)
        return nil
    }
    res["linkSpeed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetLinkSpeed(val)
        return nil
    }
    res["macAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMacAddress(val)
        return nil
    }
    res["port"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPort(val)
        return nil
    }
    res["receivedQualityEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetReceivedQualityEventRatio(val)
        return nil
    }
    res["reflexiveIPAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReflexiveIPAddress(val)
        return nil
    }
    res["relayIPAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRelayIPAddress(val)
        return nil
    }
    res["relayPort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRelayPort(val)
        return nil
    }
    res["sentQualityEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetSentQualityEventRatio(val)
        return nil
    }
    res["subnet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubnet(val)
        return nil
    }
    res["wifiBand"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.ParseWifiBand)
        if err != nil {
            return err
        }
        cast := val.(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiBand)
        m.SetWifiBand(&cast)
        return nil
    }
    res["wifiBatteryCharge"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWifiBatteryCharge(val)
        return nil
    }
    res["wifiChannel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWifiChannel(val)
        return nil
    }
    res["wifiMicrosoftDriver"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWifiMicrosoftDriver(val)
        return nil
    }
    res["wifiMicrosoftDriverVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWifiMicrosoftDriverVersion(val)
        return nil
    }
    res["wifiRadioType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.ParseWifiRadioType)
        if err != nil {
            return err
        }
        cast := val.(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiRadioType)
        m.SetWifiRadioType(&cast)
        return nil
    }
    res["wifiSignalStrength"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWifiSignalStrength(val)
        return nil
    }
    res["wifiVendorDriver"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWifiVendorDriver(val)
        return nil
    }
    res["wifiVendorDriverVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWifiVendorDriverVersion(val)
        return nil
    }
    return res
}
func (m *NetworkInfo) IsNil()(bool) {
    return m == nil
}
func (m *NetworkInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteFloat32Value("bandwidthLowEventRatio", m.GetBandwidthLowEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("basicServiceSetIdentifier", m.GetBasicServiceSetIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetConnectionType() != nil {
        cast := m.GetConnectionType().String()
        err := writer.WriteStringValue("connectionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("delayEventRatio", m.GetDelayEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dnsSuffix", m.GetDnsSuffix())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("linkSpeed", m.GetLinkSpeed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("macAddress", m.GetMacAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("receivedQualityEventRatio", m.GetReceivedQualityEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reflexiveIPAddress", m.GetReflexiveIPAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("relayIPAddress", m.GetRelayIPAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("relayPort", m.GetRelayPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("sentQualityEventRatio", m.GetSentQualityEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subnet", m.GetSubnet())
        if err != nil {
            return err
        }
    }
    if m.GetWifiBand() != nil {
        cast := m.GetWifiBand().String()
        err := writer.WriteStringValue("wifiBand", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("wifiBatteryCharge", m.GetWifiBatteryCharge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("wifiChannel", m.GetWifiChannel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiMicrosoftDriver", m.GetWifiMicrosoftDriver())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiMicrosoftDriverVersion", m.GetWifiMicrosoftDriverVersion())
        if err != nil {
            return err
        }
    }
    if m.GetWifiRadioType() != nil {
        cast := m.GetWifiRadioType().String()
        err := writer.WriteStringValue("wifiRadioType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("wifiSignalStrength", m.GetWifiSignalStrength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiVendorDriver", m.GetWifiVendorDriver())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiVendorDriverVersion", m.GetWifiVendorDriverVersion())
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
func (m *NetworkInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *NetworkInfo) SetBandwidthLowEventRatio(value *float32)() {
    m.bandwidthLowEventRatio = value
}
func (m *NetworkInfo) SetBasicServiceSetIdentifier(value *string)() {
    m.basicServiceSetIdentifier = value
}
func (m *NetworkInfo) SetConnectionType(value *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.NetworkConnectionType)() {
    m.connectionType = value
}
func (m *NetworkInfo) SetDelayEventRatio(value *float32)() {
    m.delayEventRatio = value
}
func (m *NetworkInfo) SetDnsSuffix(value *string)() {
    m.dnsSuffix = value
}
func (m *NetworkInfo) SetIpAddress(value *string)() {
    m.ipAddress = value
}
func (m *NetworkInfo) SetLinkSpeed(value *int64)() {
    m.linkSpeed = value
}
func (m *NetworkInfo) SetMacAddress(value *string)() {
    m.macAddress = value
}
func (m *NetworkInfo) SetPort(value *int32)() {
    m.port = value
}
func (m *NetworkInfo) SetReceivedQualityEventRatio(value *float32)() {
    m.receivedQualityEventRatio = value
}
func (m *NetworkInfo) SetReflexiveIPAddress(value *string)() {
    m.reflexiveIPAddress = value
}
func (m *NetworkInfo) SetRelayIPAddress(value *string)() {
    m.relayIPAddress = value
}
func (m *NetworkInfo) SetRelayPort(value *int32)() {
    m.relayPort = value
}
func (m *NetworkInfo) SetSentQualityEventRatio(value *float32)() {
    m.sentQualityEventRatio = value
}
func (m *NetworkInfo) SetSubnet(value *string)() {
    m.subnet = value
}
func (m *NetworkInfo) SetWifiBand(value *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiBand)() {
    m.wifiBand = value
}
func (m *NetworkInfo) SetWifiBatteryCharge(value *int32)() {
    m.wifiBatteryCharge = value
}
func (m *NetworkInfo) SetWifiChannel(value *int32)() {
    m.wifiChannel = value
}
func (m *NetworkInfo) SetWifiMicrosoftDriver(value *string)() {
    m.wifiMicrosoftDriver = value
}
func (m *NetworkInfo) SetWifiMicrosoftDriverVersion(value *string)() {
    m.wifiMicrosoftDriverVersion = value
}
func (m *NetworkInfo) SetWifiRadioType(value *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiRadioType)() {
    m.wifiRadioType = value
}
func (m *NetworkInfo) SetWifiSignalStrength(value *int32)() {
    m.wifiSignalStrength = value
}
func (m *NetworkInfo) SetWifiVendorDriver(value *string)() {
    m.wifiVendorDriver = value
}
func (m *NetworkInfo) SetWifiVendorDriverVersion(value *string)() {
    m.wifiVendorDriverVersion = value
}
