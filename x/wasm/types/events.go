package types

const (
	// WasmModuleEventType is stored with any contract TX that returns non empty EventAttributes
	WasmModuleEventType = "wasm"
	// CustomContractEventPrefix contracts can create custom events. To not mix them with other system events they got the `wasm-` prefix.
	CustomContractEventPrefix = "wasm-"

	EventTypeStoreCode              = "store_code"
	EventTypeInstantiate            = "instantiate"
	EventTypeExecute                = "execute"
	EventTypeMigrate                = "migrate"
	EventTypePinCode                = "pin_code"
	EventTypeUnpinCode              = "unpin_code"
	EventTypeSudo                   = "sudo"
	EventTypeReply                  = "reply"
	EventTypeGovContractResult      = "gov_contract_result"
	EventTypeUpdateContractAdmin    = "update_contract_admin"
	EventTypeUpdateCodeAccessConfig = "update_code_access_config"
)

// event attributes returned from contract execution
const (
	AttributeReservedPrefix = "_"

	AttributeKeyContractAddr        = "_contract_address"
	AttributeKeyCodeID              = "code_id"
	AttributeKeyChecksum            = "code_checksum"
	AttributeKeyResultDataHex       = "result"
	AttributeKeyRequiredCapability  = "required_capability"
	AttributeKeyNewAdmin            = "new_admin_address"
	AttributeKeyCodePermission      = "code_permission"
	AttributeKeyAuthorizedAddresses = "authorized_addresses"
)
