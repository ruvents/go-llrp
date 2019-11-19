package llrp

// Keepalive generates Keepalive message.
func Keepalive(messageID uint32) []byte {
	var data = []interface{}{
		uint16(KeepaliveHeader), // Rsvd+Ver+Type=62 (KEEPALIVE)
		uint32(10),              // Length
		messageID,               // ID
	}
	return Pack(data)
}

// KeepaliveAck generates KeepaliveAck message.
func KeepaliveAck(messageID uint32) []byte {
	var data = []interface{}{
		uint16(KeepaliveAckHeader), // Rsvd+Ver+Type=62 (KEEPALIVE)
		uint32(10),                 // Length
		messageID,                  // ID
	}
	return Pack(data)
}

// ReaderEventNotification generates ReaderEventNotification message.
func ReaderEventNotification(messageID uint32, currentTime uint64) []byte {
	readerEventNotificationData := ReaderEventNotificationData(currentTime)
	readerEventNotificationLength :=
		len(readerEventNotificationData) + 10 // Rsvd+Ver+Type+Length+ID->80bits=10bytes
	var data = []interface{}{
		uint16(ReaderEventNotificationHeader), // Rsvd+Ver+Type=63 (READER_EVENT_NOTIFICATION)
		uint32(readerEventNotificationLength), // Length
		messageID, // ID
		readerEventNotificationData,
	}
	return Pack(data)
}

// SetReaderConfig generates SetReaderConfig message.
func SetReaderConfig(messageID uint32) []byte {
	keepaliveSpec := KeepaliveSpec()
	setReaderConfigLength :=
		len(keepaliveSpec) + 11 // Rsvd+Ver+Type+Length+ID+R+Rsvd->88bits=11bytes
	var data = []interface{}{
		uint16(SetReaderConfigHeader), // Rsvd+Ver+Type=3 (SET_READER_CONFIG)
		uint32(setReaderConfigLength), // Length
		messageID,                     // ID
		uint8(0),                      // RestoreFactorySetting(no=0)+Rsvd
		keepaliveSpec,
	}
	return Pack(data)
}

// SetReaderConfig generates SetReaderConfig message.
func SetEnableMode(messageID uint32) []byte {
	var data = []interface{}{
		uint16(ImpinjEnableCutomMessageHeader),
		uint32(19), // Length
		messageID,                     // ID
		uint32(25882),
		uint8(21),
		uint32(0),
	}
	return Pack(data)
}

// SetReaderConfigResponse generates SetReaderConfigResponse message.
func SetReaderConfigResponse(messageID uint32) []byte {
	llrpStatus := Status()
	setReaderConfigResponseLength :=
		len(llrpStatus) + 10 // Rsvd+Ver+Type+Length+ID+R+Rsvd->80bits=10bytes
	var data = []interface{}{
		uint16(SetReaderConfigResponseHeader), // Rsvd+Ver+Type=13 (SET_READER_CONFIG_RESPONSE)
		uint32(setReaderConfigResponseLength), // Length
		messageID, // ID
		llrpStatus,
	}
	return Pack(data)
}

//GetReaderCapability :
func GetReaderCapability(messageID uint32) []byte {
	getReaderCapabilityLength := 1 + 10
	var data = []interface{}{
		uint16(GetReaderCapabilityHeader),
		uint32(getReaderCapabilityLength),
		messageID,
		uint8(0), //all capabilities
	}
	return Pack(data)
}

//GetReaderCapabilityResponse :
func GetReaderCapabilityResponse(messageID uint32) []byte {

	llrpStatus := Status()
	generalCapabilites := GeneralDeviceCapabilities()
	llrpCapabilities := LlrpCapabilities()
	c1g2llrpCapabilities := C1G2llrpCapabilities()
	reguCapabilitles := ReguCapabilities()
	length := 2 + 4 + 4 + len(llrpStatus) + len(generalCapabilites) + len(llrpCapabilities) + len(reguCapabilitles) + len(c1g2llrpCapabilities)
	var data = []interface{}{
		uint16(GetReaderCapabilityResponseHeader),
		uint32(length),
		uint32(messageID),
		llrpStatus,
		generalCapabilites,
		llrpCapabilities,
		reguCapabilitles,
		// uint8(0),
		// uint8(0),
		// uint8(0),
		c1g2llrpCapabilities,
	}
	return Pack(data)
}

//GetReaderConfigResponse :
func GetReaderConfigResponse(messageID uint32) []byte {
	llrpStatus := Status()
	//numOfAntennas := 52
	identification := GetReaderConfigResponseIdentification()
	length := 2 + 4 + 4 + len(llrpStatus) + len(identification) //+ numOfAntennas*9 + numOfAntennas*36
	var data = []interface{}{
		uint16(GetReaderConfigResponseHeader),
		uint32(length),
		messageID,
		llrpStatus,
		identification,
	}
	// x := Pack(data)
	// for i := 1; i <= numOfAntennas; i++ {
	// 	x = append(x, AntennaProperties(uint16(i))...)
	// }
	// for i := 1; i <= numOfAntennas; i++ {
	// 	x = append(x, AntennaConfiguration(uint16(i))...)
	// }
	return Pack(data)
}

//DeleteRospecResponse : Delete RoSpec Response
func DeleteAcessSpec(messageID uint32) []byte {
	var data = []interface{}{
		uint16(DeleteAccessSpecHeader),
		uint32(14), //length
		messageID,
		uint32(0),
	}
	return Pack(data)
}

//DeleteAccessSpecResponse : Delete Access Spec Response
func DeleteAccessSpecResponse(messageID uint32) []byte {
	llrpStatus := Status()
	var data = []interface{}{
		uint16(DeleteAccessSpecResponseHeader),
		uint32(18), //length
		messageID,
		llrpStatus,
	}
	return Pack(data)
}

//DeleteRospecResponse : Delete RoSpec Response
func DeleteRospec(messageID uint32) []byte {
	var data = []interface{}{
		uint16(DeleteRospecHeader),
		uint32(14), //length
		messageID,
		uint32(0),
	}
	return Pack(data)
}

//DeleteRospecResponse : Delete RoSpec Response
func DeleteRospecResponse(messageID uint32) []byte {
	llrpStatus := Status()
	var data = []interface{}{
		uint16(DeleteRospecResponseHeader),
		uint32(18), //length
		messageID,
		llrpStatus,
	}
	return Pack(data)
}


//AddRospec : Add ROSpec Response
func AddRospec(messageID uint32) []byte {
	var data = []interface{}{
		uint16(AddRospecHeader),
		uint32(441), //length
		messageID,
		/*RO  Spec */
		uint16(177),  //RO Spec
		uint16(431), //length
		uint32(123), //RO Spec ID
		uint8(0), //Priority
		uint8(0), //State
		/*RO Bound Spec */
		uint16(178),  //RO Bound Spec
		uint16(18),  //length
		uint16(179),  //RO Spec Start Triger
		uint16(5),  //length
		uint8(1), //RO Spec Start Triger Type
		uint16(182),  //RO Spec Stop Triger
		uint16(9),  //length
		uint8(0), //RO Spec Stop Triger Type
		uint32(0), //Duration triger value
		/*AI Spec */
		uint16(183), //AI Spec
		uint16(390),  //length
		uint16(4),  //Antena Count
		uint16(1),  //Antena ID
    uint16(2),  //Antena ID
		uint16(3),  //Antena ID
		uint16(4),  //Antena ID
		uint16(184), //AI Spec Stop Triger
		uint16(9),  //length
    uint8(0), //AI Spec Stop Triger Type
		uint32(0), //Duration triger value
		uint16(186), //AI Spec Invertory Parametr Spec ID
		uint16(367),  //length
		uint16(1234), //Invertory Parametr Spec ID
		uint8(1), //Protocol ID
		/*TLV Antena 1 */
		uint16(222), //Antena Configuration
		uint16(90),  //length
		uint16(1),  //Antena ID
		uint16(224), //RF Transmiter
		uint16(10),  //length
		uint16(0),  //HOP Table ID
		uint16(1),  //Channel Index
		uint16(81),  //Transmit Power
		uint16(330), //C1 G2 Ibventory Command
		uint16(74),  //length
		uint8(0),  //Tag Inventory State aware
		uint16(335), //RF Control
		uint16(8),  //length
		uint16(1000),  //Mode Index
		uint16(0),  //Tari
		uint16(336), //Sigulation Control
		uint16(11),  //length
		uint8(128),  //Session
		uint16(32),  //Tag Population
		uint32(0),  //Tag Transmit Time
		uint16(1023), //Inventory Search mode
		uint16(14),  //length
    uint32(25882),  //Tag Transmit Time
		uint32(23),  //Parametr type
		uint16(2),  //Inventory Search mode value
		uint16(1023), //Fixed Frequency List
		uint16(18),  //length
		uint32(25882),  //Tag Transmit Time
		uint32(26),  //Parametr type
		uint16(1),  //Fixed Frequency mode
		uint16(0),  //Reserverd
		uint16(0),  //Number Channels
		uint16(1023), //Low Dyte Cycle
		uint16(18),  //length
		uint32(25882),  //Tag Transmit Time
		uint32(28),  //Parametr type
		uint16(0),  //Low Dyte Cycle mode
		uint16(0),  //Empty field timeout
		uint16(0),  //Field Ping Timeout
		/*TLV Antena 2 */
		uint16(222), //Antena Configuration
		uint16(90),  //length
		uint16(2),  //Antena ID
		uint16(224), //RF Transmiter
		uint16(10),  //length
		uint16(0),  //HOP Table ID
		uint16(1),  //Channel Index
		uint16(81),  //Transmit Power
		uint16(330), //C1 G2 Ibventory Command
		uint16(74),  //length
		uint8(0),  //Tag Inventory State aware
		uint16(335), //RF Control
		uint16(8),  //length
		uint16(1000),  //Mode Index
		uint16(0),  //Tari
		uint16(336), //Sigulation Control
		uint16(11),  //length
		uint8(128),  //Session
		uint16(32),  //Tag Population
		uint32(0),  //Tag Transmit Time
		uint16(1023), //Inventory Search mode
		uint16(14),  //length
		uint32(25882),  //Tag Transmit Time
		uint32(23),  //Parametr type
		uint16(2),  //Inventory Search mode value
		uint16(1023), //Fixed Frequency List
		uint16(18),  //length
		uint32(25882),  //Tag Transmit Time
		uint32(26),  //Parametr type
		uint16(1),  //Fixed Frequency mode
		uint16(0),  //Reserverd
		uint16(0),  //Number Channels
		uint16(1023), //Low Dyte Cycle
		uint16(18),  //length
		uint32(25882),  //Tag Transmit Time
		uint32(28),  //Parametr type
		uint16(0),  //Low Dyte Cycle mode
		uint16(0),  //Empty field timeout
		uint16(0),  //Field Ping Timeout
		/*TLV Antena 3 */
		uint16(222), //Antena Configuration
		uint16(90),  //length
		uint16(3),  //Antena ID
		uint16(224), //RF Transmiter
		uint16(10),  //length
		uint16(0),  //HOP Table ID
		uint16(1),  //Channel Index
		uint16(81),  //Transmit Power
		uint16(330), //C1 G2 Ibventory Command
		uint16(74),  //length
		uint8(0),  //Tag Inventory State aware
		uint16(335), //RF Control
		uint16(8),  //length
		uint16(1000),  //Mode Index
		uint16(0),  //Tari
		uint16(336), //Sigulation Control
		uint16(11),  //length
		uint8(128),  //Session
		uint16(32),  //Tag Population
		uint32(0),  //Tag Transmit Time
		uint16(1023), //Inventory Search mode
		uint16(14),  //length
		uint32(25882),  //Tag Transmit Time
		uint32(23),  //Parametr type
		uint16(2),  //Inventory Search mode value
		uint16(1023), //Fixed Frequency List
		uint16(18),  //length
		uint32(25882),  //Tag Transmit Time
		uint32(26),  //Parametr type
		uint16(1),  //Fixed Frequency mode
		uint16(0),  //Reserverd
		uint16(0),  //Number Channels
		uint16(1023), //Low Dyte Cycle
		uint16(18),  //length
		uint32(25882),  //Tag Transmit Time
		uint32(28),  //Parametr type
		uint16(0),  //Low Dyte Cycle mode
		uint16(0),  //Empty field timeout
		uint16(0),  //Field Ping Timeout
		/*TLV Antena 4 */
		uint16(222), //Antena Configuration
		uint16(90),  //length
		uint16(4),  //Antena ID
		uint16(224), //RF Transmiter
		uint16(10),  //length
		uint16(0),  //HOP Table ID
		uint16(1),  //Channel Index
		uint16(81),  //Transmit Power
		uint16(330), //C1 G2 Ibventory Command
		uint16(74),  //length
		uint8(0),  //Tag Inventory State aware
		uint16(335), //RF Control
		uint16(8),  //length
		uint16(1000),  //Mode Index
		uint16(0),  //Tari
		uint16(336), //Sigulation Control
		uint16(11),  //length
		uint8(128),  //Session
		uint16(32),  //Tag Population
		uint32(0),  //Tag Transmit Time
		uint16(1023), //Inventory Search mode
		uint16(14),  //length
		uint32(25882),  //Tag Transmit Time
		uint32(23),  //Parametr type
		uint16(2),  //Inventory Search mode value
		uint16(1023), //Fixed Frequency List
		uint16(18),  //length
		uint32(25882),  //Tag Transmit Time
		uint32(26),  //Parametr type
		uint16(1),  //Fixed Frequency mode
		uint16(0),  //Reserverd
		uint16(0),  //Number Channels
		uint16(1023), //Low Dyte Cycle
		uint16(18),  //length
		uint32(25882),  //Tag Transmit Time
		uint32(28),  //Parametr type
		uint16(0),  //Low Dyte Cycle mode
		uint16(0),  //Empty field timeout
		uint16(0),  //Field Ping Timeout
		/*RO Report Spec */
		uint16(237), //RO Report Spec
		uint16(13),  //length
		uint8(2), //RO Report triger
		uint16(1),  //N
		uint16(238), //RO Report Spec
		uint16(6),  //length
		uint16(7744),  //Param
	}
	return Pack(data)
}
//AddRospecResponse : Add ROSpec Response
func AddRospecResponse(messageID uint32) []byte {
	llrpStatus := Status()
	var data = []interface{}{
		uint16(AddRospecResponseHeader),
		uint32(18), //length
		messageID,
		llrpStatus,
	}
	return Pack(data)
}

//EnableRospec : Enabled Rospec Response
func EnableRospec(messageID uint32) []byte {
	var data = []interface{}{
		uint16(EnableRospecHeader),
		uint32(14), //length
		messageID,
		uint32(123),
	}
	return Pack(data)
}

//EnableRospecResponse : Enabled Rospec Response
func EnableRospecResponse(messageID uint32) []byte {
	llrpStatus := Status()
	var data = []interface{}{
		uint16(EnableRospecResponseHeader),
		uint32(18), //length
		messageID,
		llrpStatus,
	}
	return Pack(data)
}

//StartRospec : Enabled Rospec Response
func StartRospec(messageID uint32) []byte {
	var data = []interface{}{
		uint16(StartRospecHeader),
		uint32(14), //length
		messageID,
		uint32(123),
	}
	return Pack(data)
}

//ReceiveSensitivityEntries : Generates ReceiveSensitivityEntries used in General capabilities
func ReceiveSensitivityEntries(numOfAntennas int) []interface{} {
	var data = []interface{}{}
	for i := 1; i <= numOfAntennas; i++ {
		x := ReceiveSensitivityEntry(uint16(i))
		data = append(data, x)
	}
	return data
}

//ReceiveSensitivityEntry :
func ReceiveSensitivityEntry(id uint16) []byte {
	var data = []interface{}{
		uint16(139), //type
		uint16(8),   //length
		uint16(id),  //id
		uint16(11),  //receive sentitvitiy value
	}
	return Pack(data)
}

//GPIOCapabilities : Generates GPIO capabilities proeprty
func GPIOCapabilities() []byte {
	var data = []interface{}{
		uint16(141), //type
		uint16(8),   //length
		uint16(0),   //num of GPI port
		uint16(0),   //num of GPO port
	}
	return Pack(data)
}

//AntennaAirPortList :
func AntennaAirPortList(numOfAntennas int) []interface{} {
	var data = []interface{}{}
	for i := 1; i <= numOfAntennas; i++ {
		x := AntennaAirPort(uint16(i))
		data = append(data, x)
	}
	return data
}

//AntennaAirPort :
func AntennaAirPort(id uint16) []byte {
	var data = []interface{}{
		uint16(140), //type
		uint16(9),   //length
		uint16(id),
		uint16(1), //num of protocols
		uint8(1),  //protocol id : EPCGlobal Class 1 Gen 2
	}
	return Pack(data)
}
