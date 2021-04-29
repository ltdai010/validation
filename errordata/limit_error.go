package errordata

import (
	"errors"
)

var (
	SUCCESS             = errors.New("SUCCESS")
	UNSUCCESS           = errors.New("UNSUCCESS")
	ERR_UNKNOWN         = errors.New("UNKNOWN")
	BAD_REQUEST         = errors.New("BAD_REQUEST")
	EMPTY_DATA          = errors.New("EMPTY_DATA")
	ER_NOT_INT          = errors.New("ERR_INT")
	SESSION_EXPIRED     = errors.New("SESSION_EXPIRED")
	VERIFY_CODE_EXPIRED = errors.New("VERIFY_CODE_EXPIRED")
	EXISTED             = errors.New("EXISTED")
	NOT_ADMIN           = errors.New("NOT_ADMIN")
	NOT_FOUND           = errors.New("NOT_FOUND")
	NOT_PERMISSION      = errors.New("NOT_PERMISSION")
	UNAUTHORIZED        = errors.New("UNAUTHORIZED")
	ERR_LOGIN           = errors.New("LOGIN_ERROR")
	ERR_SYSTEM          = errors.New("SYSTEM_ERROR")
	NOT_EXISTED         = errors.New("DATA_NOT_EXISTED!")
	USER_ID_NOT_EXISTED = errors.New("USER_ID_DOES_NOT_EXIST")
	ERR_CHANGEPASS      = errors.New("CHANGE_PASSWORD_ERROR!")
	NOT_MORE            = errors.New("NO_MORE_DATA")
	EXISTED_POS_USER    = errors.New("EXISTED_POS_USER")
	LIMIT_CHARACTER     = errors.New("LIMIT_CHARACTER")
	CANNOT_EMPTY        = errors.New("CANNOT_EMPTY")
	ERR_SIG             = errors.New("SIG_ERROR")
	TIME_NOT_EXCEED     = errors.New("TIME_NOT_EXCEED")
	BANNED              = errors.New("BANNED")

	//
	LIMIT_WITHDRAWAL_ERROR                     = errors.New("LIMIT_WITHDRAWAL_ERROR")
	BLOCK_CHAIN_ADDRESS_NOT_EXIST_ERROR        = errors.New("BLOCK_CHAIN_ADDRESS_NOT_EXIST_ERROR")
	BLOCK_CHAIN_ADDRESS_INVALID                = errors.New("BLOCK_CHAIN_ADDRESS_INVALID")
	ITEM_NOT_EXIST                             = errors.New("ITEM_NOT_EXIST")
	ITEM_EXISTED                               = errors.New("ITEM_EXISTED")
	BALANCE_INSUFFICIENT_ERROR                 = errors.New("BALANCE_INSUFFICIENT_ERROR")
	UPDATE_BALANCE_FAILED                      = errors.New("UPDATE_BALANCE_FAILED")
	INVALID_ID                                 = errors.New("INVALID_ID")
	PUBKEY_INVALID_ERROR                       = errors.New("PUBKEY_INVALID_ERROR")
	LEVEL_SECURITY_ERROR                       = errors.New("LEVEL_SECURITY_ERROR")
	LIMIT_NOT_EXISTED                          = errors.New("LIMIT_NOT_EXISTED")
	USER_INFO_EMPTY                            = errors.New("USER_INFO_EMPTY")
	ADDRESS_POOL_EMPTY                         = errors.New("ADDRESS_POOL_EMPTY")
	COIN_NOT_EXIST                             = errors.New("COIN_NOT_EXIST")
	TRANSFER_FAILED                            = errors.New("TRANSFER_FAILED")
	PRIVATE_KEY_INVALID                        = errors.New("PRIVATE_KEY_INVALID")
	PUBKEY_NOT_EXIST_ERROR                     = errors.New("PUBKEY_NOT_EXIST_ERROR")
	ERROR_UNKNOWN                              = errors.New("ERROR_UNKNOWN")
	CONTRACT_INVALID                           = errors.New("CONTRACT_INVALID")
	DATABIGSET_NULL                            = errors.New("DATABIGSET_NULL")
	DATABIGSET_PUT_DATA_ERR                    = errors.New("DATABIGSET_PUT_DATA_ERR")
	CAN_NOT_EMPTY                              = errors.New("CAN_NOT_EMPTY")
	DATA_INVALID                               = errors.New("DATA_INVALID")
	THIS_LUCKY_PRIZE_HAVE_NO_PRIZE             = errors.New("THIS_LUCKY_PRIZE_HAVE_NO_PRIZE")
	ADDRESS_EXHAUSTED                          = errors.New("ADDRESS_EXHAUSTED")
	CLIENT_IS_NIL                              = errors.New("CLIENT_IS_NIL")
	ADDRESS_NOT_VALID                          = errors.New("ADDRESS_NOT_VALID")
	QR_TIMEOUT                                 = errors.New("QR_TIMEOUT")
	NEED_CONFIRM                               = errors.New("NEED_CONFIRM")
	STEP_INVALID                               = errors.New("STEP_INVALID")
	USER_IS_NOT_EXISTED                        = errors.New("USER_IS_NOT_EXISTED")
	ROLE_IS_NOT_EXISTED                        = errors.New("ROLE_IS_NOT_EXISTED")
	PERMISSION_IS_NOT_EXISTED                  = errors.New("PERMISSION_IS_NOT_EXISTED")
	NOT_VALID_CHAIN                            = errors.New("NOT_VALID_CHAIN")
	CACHE_INVALID                              = errors.New("CACHE_INVALID")
	GET_ID_ERROR                               = errors.New("GET_ID_ERROR")
	GET_INDEX_ERROR                            = errors.New("GET_INDEX_ERROR")
	CREATE_COIN_FAILED                         = errors.New("CREATE_COIN_FAILED")
	ROLE_NAME_IS_ALREADY_TAKEN                 = errors.New("ROLE_NAME_IS_ALREADY_TAKEN")
	RESULT_EMPTY                               = errors.New("RESULT_EMPTY")
	PERMISSION_NAME_AND_PATH_ARE_ALREADY_TAKEN = errors.New("PERMISSION_NAME_AND_PATH_ARE_ALREADY_TAKEN")
	UNMARSHAL_JSON                             = errors.New("UNMARSHAL_JSON")
	//MUST_BE_GREATER_THAN                       = func(agrs ...string) error { return errors.New(fmt.Sprintf("MUST_BE_GREATER_THAN_NUMBER_%s", agrs)) }
	//CAN_NOT_CONNECT_TO_SERVER_BACKEND          = func(agrs ...string) error {
	//	return errors.New(fmt.Sprintf("Can not connect to backend service: host: %s port: %s", agrs))
	//}
	// validation
	VALIDATION_ADDRESS_BTC                       = errors.New("VALIDATION_ADDRESS_BTC")
	VALIDATION_CREATE_ADMIN                      = errors.New("VALIDATION_CREATE_ADMIN")
	VALIDATION_MIGRATE                           = errors.New("VALIDATION_MIGRATE")
	VALIDATION_NOTIFICATION                      = errors.New("VALIDATION_MIGRATE")
	VALIDATION_CAMPAIGN                          = errors.New("VALIDATION_CAMPAIGN")
	VALIDATION_PUBKEY                            = errors.New("VALIDATION_PUBKEY")
	VALIDATION_PERMISSION                        = errors.New("VALIDATION_PERMISSION")
	VALIDATION_ROLENAME                          = errors.New("VALIDATION_ROLENAME")
	VALIDATION_ROLEPERMISSION                    = errors.New("VALIDATION_ROLEPERMISSION")
	VALIDATION_ROLE_USER                         = errors.New("VALIDATION_ROLE_USER")
	VALIDATION_ADMIN_TRANSFER_LIMIT_BODY_REQUEST = errors.New("ADMIN_TRANSFER_LIMIT_BODY_REQUEST")
	VALIDATION_LIMIT_WITHDRAWAL                  = errors.New("VALIDATION_LIMIT_WITHDRAWAL")
	VALIDATION_COIN_SETTING                      = errors.New("VALIDATION_COIN_SETTING")
	VALIDATION_COIN_REQUEST                      = errors.New("VALIDATION_COIN_REQUEST")
	VALIDATION_WITHDRAW                          = errors.New("VALIDATION_WITHDRAW")
	VALIDATION_CHAIN                             = errors.New("VALIDATION_CHAIN")
	VALIDATION_ADMIN_NOTIFICATION                = errors.New("VALIDATION_Admin_Notification")
	VALIDATION_COIN                              = errors.New("VALIDATION_COIN")
	VALIDATION_CRYPTO_EXCHANGE                   = errors.New("VALIDATION_CRYPTO_EXCHANGE")
	VALIDATION_ADMIN_REQUEST                     = errors.New("VALIDATION_ADMIN_REQUEST")
	VALIDATION_ADMIN_WALLET_REQUEST              = errors.New("VALIDATION_ADMIN_WALLET_REQUEST")
	// logic error
	OBJECT_MUST_BE_A_POINTER                         = errors.New("OBJECT_MUST_BE_A_POINTER")
	OBJECT_MUST_BE_A_STRUCT                          = errors.New("OBJECT_MUST_BE_A_STRUCT")
	OBJECT_MUST_BE_NOT_NIL                           = errors.New("OBJECT_MUST_BE_NOT_NIL")
	SCAN_VALUE_IS_NOT_INT64                          = errors.New("SCAN_VALUE_IS_NOT_INT64")
	COUNTER_WRONG_NUMBER                             = errors.New("COUNTER_WRONG_NUMBER")
	ARR_MUST_BE_AN_ARRAY_OR_A_SLICE                  = errors.New("ARR_MUST_BE_AN_ARRAY_OR_A_SLICE")
	THE_TYPE_OF_THE_ARR_AND_THE_OBJ_MUST_BE_THE_SAME = errors.New("THE_TYPE_OF_THE_ARR_AND_THE_OBJ_MUST_BE_THE_SAME")
	COULD_NOT_FIND_OBJ_IN_ARR                        = errors.New("COULD_NOT_FIND_OBJ_IN_ARR")
	KIND_OF_OBJECTS_DOES_NOT_MATCH_SLICE_TYPE        = errors.New("KIND_OF_OBJECTS_DOES_NOT_MATCH_SLICE_TYPE")
	CONTRACT_DID_NOT_MATCH                           = errors.New("CONTRACT_DID_NOT_MATCH")
	INDEX_MUST_BE_GREATER_THAN_OR_EQUAL_TO_0         = errors.New("INDEX_MUST_BE_GREATER_THAN_OR_EQUAL_TO_0")

	MapDescription = map[error]string{
		SUCCESS:                             "Success!",
		ERR_UNKNOWN:                         "Unknown error!",
		UNAUTHORIZED:                        "UnAuthorized",
		BAD_REQUEST:                         "Bad Request!",
		EMPTY_DATA:                          "Empty Data",
		ER_NOT_INT:                          "Param not int!",
		UNSUCCESS:                           "Unsuccess!",
		VERIFY_CODE_EXPIRED:                 "Verify Code Expired",
		SESSION_EXPIRED:                     "SessionExpired!",
		EXISTED:                             "Existed !",
		NOT_ADMIN:                           "Not Admin !",
		NOT_FOUND:                           "Not Found !",
		NOT_PERMISSION:                      "Not Permission !",
		ERR_LOGIN:                           "Wrong username, password. ",
		ERR_SYSTEM:                          "The system is having problems.",
		NOT_EXISTED:                         "Data Not Existed!",
		USER_ID_NOT_EXISTED:                 "User Id Data Not Existed!",
		ERR_CHANGEPASS:                      "Wrong username, password.",
		NOT_MORE:                            "No more data",
		EXISTED_POS_USER:                    "Existed position organization",
		ERR_SIG:                             "Sig error !",
		LIMIT_CHARACTER:                     "Limit character",
		CANNOT_EMPTY:                        "Data empty !",
		TIME_NOT_EXCEED:                     "Time not exceed!",
		BANNED:                              "Banned",
		LIMIT_WITHDRAWAL_ERROR:              "Withdraw exceed limit",
		PUBKEY_INVALID_ERROR:                "Pubkey Invalid",
		BLOCK_CHAIN_ADDRESS_NOT_EXIST_ERROR: "Block chain address not exist",
		PRIVATE_KEY_INVALID:                 "Private key invalid",
		BALANCE_INSUFFICIENT_ERROR:          "Balance insufficient",
		LEVEL_SECURITY_ERROR:                "Level security to low for access feature",
		INVALID_ID:                          "Id invalid",
		USER_INFO_EMPTY:                     "User info empty",
		ADDRESS_POOL_EMPTY:                  "Address poll is empty",
		COIN_NOT_EXIST:                      "Coin not exist",
		BLOCK_CHAIN_ADDRESS_INVALID:         "Block chain address invalid",
		PUBKEY_NOT_EXIST_ERROR:              "Pubkey not exist",
		ERROR_UNKNOWN:                       "Wrong Unknown.",
		CONTRACT_INVALID:                    "Databigset Null",
		DATABIGSET_NULL:                     "Databigset Null!",
		DATABIGSET_PUT_DATA_ERR:             "Databigset Put Data Err",
		CAN_NOT_EMPTY:                       "Can Not Empty",
		DATA_INVALID:                        "Data Invalid",
		THIS_LUCKY_PRIZE_HAVE_NO_PRIZE:      "This Lucky Prize Have No Prize",
		ADDRESS_EXHAUSTED:                   "Address Exhausted",
		CLIENT_IS_NIL:                       "Client Is Nil",
		ADDRESS_NOT_VALID:                   "Address Not Valid",
		QR_TIMEOUT:                          "Qr Timeout",
		NEED_CONFIRM:                        "Need Confirm",
		STEP_INVALID:                        "Step Invalid",
		USER_IS_NOT_EXISTED:                 "User Is Not Existed",
		ROLE_IS_NOT_EXISTED:                 "Role Is Not Existed",
		PERMISSION_IS_NOT_EXISTED:           "Permission Is Not Existed",
		NOT_VALID_CHAIN:                     "Not Valid Chain",
		CACHE_INVALID:                       "Cache Invalid",
		GET_ID_ERROR:                        "Get Id Error",
		GET_INDEX_ERROR:                     "Get Index Error",
		CREATE_COIN_FAILED:                  "Create Coin Failed",
		ROLE_NAME_IS_ALREADY_TAKEN:          "Role Name Is Already Taken",
		RESULT_EMPTY:                        "Result Empty",
		UNMARSHAL_JSON:                      "Unmarshal Json",
		PERMISSION_NAME_AND_PATH_ARE_ALREADY_TAKEN: "Permission Name And Path Are Already Taken",
		//MUST_BE_GREATER_THAN():                     "Must be greater than %s",
		/// validation
		VALIDATION_ADDRESS_BTC:                       ":path Address btc is invalid",
		VALIDATION_CREATE_ADMIN:                      "body Create body is invalid",
		VALIDATION_MIGRATE:                           "body post body is invalid GAS or ADDRESS",
		VALIDATION_NOTIFICATION:                      "body create body is invalid",
		VALIDATION_CAMPAIGN:                          "body create body is invalid",
		VALIDATION_PUBKEY:                            "Pubkey is invalid",
		VALIDATION_PERMISSION:                        "body permission is invalid",
		VALIDATION_ROLENAME:                          "body Role name is invalid",
		VALIDATION_ROLEPERMISSION:                    "body Role permission is invalid",
		VALIDATION_ROLE_USER:                         "body role user request is invalid",
		VALIDATION_ADMIN_TRANSFER_LIMIT_BODY_REQUEST: "body role admin transfer limit body is invalid",
		VALIDATION_LIMIT_WITHDRAWAL:                  "body limit withdraw body is invalid",
		VALIDATION_COIN_SETTING:                      "body coin setting body is invalid",
		VALIDATION_COIN_REQUEST:                      "body coin request body is invalid",
		VALIDATION_WITHDRAW:                          "body coin withdraw body is invalid",
		VALIDATION_CHAIN:                             "Chain is invalid",
		VALIDATION_ADMIN_NOTIFICATION:                "Admin notification invalid",
		VALIDATION_COIN:                              "coin/currency is invalid",
		VALIDATION_CRYPTO_EXCHANGE:                   "crypto exchange is invalid",
		VALIDATION_ADMIN_REQUEST:                     "body admin request is invalid",
		VALIDATION_ADMIN_WALLET_REQUEST:              "body admin wallet request is invalid",
		/// logic
		OBJECT_MUST_BE_A_POINTER:                         "Object must be a pointer",
		OBJECT_MUST_BE_A_STRUCT:                          "Object must be a struct",
		OBJECT_MUST_BE_NOT_NIL:                           "Object must be not nil",
		SCAN_VALUE_IS_NOT_INT64:                          "Scan value is not int64",
		COUNTER_WRONG_NUMBER:                             "Counter wrong number",
		ARR_MUST_BE_AN_ARRAY_OR_A_SLICE:                  "Array must be an array or a slice",
		THE_TYPE_OF_THE_ARR_AND_THE_OBJ_MUST_BE_THE_SAME: "The Type Of The Arr And The Obj Must Be The Same",
		COULD_NOT_FIND_OBJ_IN_ARR:                        "Could Not Find Obj In Arr",
		KIND_OF_OBJECTS_DOES_NOT_MATCH_SLICE_TYPE:        "Kind Of Objects Does Not Match Slice Type",
		CONTRACT_DID_NOT_MATCH:                           "Contract Did Not Match",
		INDEX_MUST_BE_GREATER_THAN_OR_EQUAL_TO_0:         "Index Must Be Greater Than Or Equal To 0",
	}

	MapErrorCode = map[error]int64{
		SUCCESS:                             200,
		UNSUCCESS:                           201,
		ER_NOT_INT:                          302,
		SESSION_EXPIRED:                     303,
		NOT_EXISTED:                         304,
		USER_ID_NOT_EXISTED:                 415,
		EXISTED:                             305,
		ERR_CHANGEPASS:                      306,
		NOT_ADMIN:                           307,
		NOT_FOUND:                           404,
		NOT_PERMISSION:                      403,
		NOT_MORE:                            309,
		EXISTED_POS_USER:                    310,
		CANNOT_EMPTY:                        311,
		LIMIT_CHARACTER:                     312,
		ERR_SIG:                             316,
		BAD_REQUEST:                         400,
		EMPTY_DATA:                          202,
		ERR_UNKNOWN:                         401,
		ERR_LOGIN:                           402,
		ERR_SYSTEM:                          500,
		UNAUTHORIZED:                        405,
		VERIFY_CODE_EXPIRED:                 406,
		TIME_NOT_EXCEED:                     408,
		BANNED:                              400,
		LIMIT_WITHDRAWAL_ERROR:              407,
		PUBKEY_INVALID_ERROR:                407,
		BLOCK_CHAIN_ADDRESS_NOT_EXIST_ERROR: 408,
		BALANCE_INSUFFICIENT_ERROR:          409,
		INVALID_ID:                          501,
		LEVEL_SECURITY_ERROR:                410,
		USER_INFO_EMPTY:                     303,
		ADDRESS_POOL_EMPTY:                  410,
		BLOCK_CHAIN_ADDRESS_INVALID:         411,
		COIN_NOT_EXIST:                      412,
		PRIVATE_KEY_INVALID:                 413,
		PUBKEY_NOT_EXIST_ERROR:              414,
		LIMIT_NOT_EXISTED:                   416,
		TRANSFER_FAILED:                     415,
		ERROR_UNKNOWN:                       401,
		CONTRACT_INVALID:                    417,
		DATABIGSET_NULL:                     418,
		DATABIGSET_PUT_DATA_ERR:             419,
		CAN_NOT_EMPTY:                       420,
		DATA_INVALID:                        311,
		THIS_LUCKY_PRIZE_HAVE_NO_PRIZE:      421,
		ADDRESS_EXHAUSTED:                   422,
		CLIENT_IS_NIL:                       423,
		ADDRESS_NOT_VALID:                   424,
		QR_TIMEOUT:                          425,
		NEED_CONFIRM:                        426,
		STEP_INVALID:                        427,
		USER_IS_NOT_EXISTED:                 428,
		ROLE_IS_NOT_EXISTED:                 429,
		PERMISSION_IS_NOT_EXISTED:           430,
		NOT_VALID_CHAIN:                     431,
		CACHE_INVALID:                       432,
		GET_ID_ERROR:                        433,
		GET_INDEX_ERROR:                     434,
		CREATE_COIN_FAILED:                  435,
		ROLE_NAME_IS_ALREADY_TAKEN:          436,
		RESULT_EMPTY:                        202,
		PERMISSION_NAME_AND_PATH_ARE_ALREADY_TAKEN: 437,
		//MUST_BE_GREATER_THAN():                     449,
		// logic
		OBJECT_MUST_BE_A_POINTER:                         438,
		OBJECT_MUST_BE_A_STRUCT:                          439,
		OBJECT_MUST_BE_NOT_NIL:                           440,
		SCAN_VALUE_IS_NOT_INT64:                          441,
		COUNTER_WRONG_NUMBER:                             442,
		ARR_MUST_BE_AN_ARRAY_OR_A_SLICE:                  443,
		THE_TYPE_OF_THE_ARR_AND_THE_OBJ_MUST_BE_THE_SAME: 444,
		COULD_NOT_FIND_OBJ_IN_ARR:                        445,
		KIND_OF_OBJECTS_DOES_NOT_MATCH_SLICE_TYPE:        446,
		CONTRACT_DID_NOT_MATCH:                           447,
		INDEX_MUST_BE_GREATER_THAN_OR_EQUAL_TO_0:         448,
		VALIDATION_ADDRESS_BTC:                           499,
		VALIDATION_MIGRATE:                               499,
		VALIDATION_NOTIFICATION:                          499,
		UNMARSHAL_JSON:                                   498,
		VALIDATION_CAMPAIGN:                              499,
		VALIDATION_PUBKEY:                                499,
		VALIDATION_PERMISSION:                            499,
		VALIDATION_ROLENAME:                              499,
		VALIDATION_ROLEPERMISSION:                        499,
		VALIDATION_ROLE_USER:                             499,
		VALIDATION_ADMIN_TRANSFER_LIMIT_BODY_REQUEST:     499,
		VALIDATION_LIMIT_WITHDRAWAL:                      499,
		VALIDATION_COIN_SETTING:                          499,
		VALIDATION_COIN_REQUEST:                          499,
		VALIDATION_WITHDRAW:                              499,
		VALIDATION_CHAIN:                                 499,
		VALIDATION_ADMIN_NOTIFICATION:                    499,
		VALIDATION_COIN:                                  499,
		VALIDATION_CRYPTO_EXCHANGE:                       499,
		VALIDATION_ADMIN_REQUEST:                         499,
		VALIDATION_ADMIN_WALLET_REQUEST:                  499,
	}
)

