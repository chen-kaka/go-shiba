package exception

const (
	// All kinds of error code definitions
	
	//success
	SUCCESS = 0
	
	// client error code
	PARAM_ERROR = 10001
	PARAM_MISSING = 10002
	PARAM_NOT_ALLOWED = 10003
	PASSWORD_ERROR = 10004
	LOGIN_FAILED = 10005
	
	// server error code
	SERVICE_ERROR = 20001
	DATABASE_ERROR = 20002
	MIDDLEWARE_ERROR = 20003
	DATA_ERROR = 20004
	UNKNOWN_ERROR = 20005
	
	// interface error code
	INTERFACE_ERROR = 30001
	INTERFACE_PENDING = 30002
	
	//define other error code here
	
)
