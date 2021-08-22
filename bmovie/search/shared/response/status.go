package response

import "net/http"

type status int

const (
	SUCCESS status = iota
	SYSTEM_ERROR
	DUPLICATE_DATA
	DATA_NOT_EXIST
	BIND_ERROR
	RUNTIME_ERROR
	DATE_NOT_VALID
	VENDOR_SHUTDOWN
	METHOD_ARGUMENTS_NOT_VALID
	TOO_MANY_REQUEST
	BAD_REQUEST
	UNAUTHORIZE
)

var respCode = map[status]string{
	SUCCESS:                    "00000",
	SYSTEM_ERROR:               "00001",
	DUPLICATE_DATA:             "00002",
	DATA_NOT_EXIST:             "00003",
	BIND_ERROR:                 "00004",
	RUNTIME_ERROR:              "00005",
	DATE_NOT_VALID:             "00006",
	VENDOR_SHUTDOWN:            "00007",
	METHOD_ARGUMENTS_NOT_VALID: "00008",
	TOO_MANY_REQUEST:           "00009",
	BAD_REQUEST:                "000010",
	UNAUTHORIZE:                "000011",
}

var respStatus = map[status]string{
	SUCCESS:                    "SUCCESS",
	SYSTEM_ERROR:               "SYSTEM_ERROR",
	DUPLICATE_DATA:             "DUPLICATE_DATA",
	DATA_NOT_EXIST:             "DATA_NOT_EXIST",
	BIND_ERROR:                 "BIND_ERROR",
	RUNTIME_ERROR:              "RUNTIME_ERROR",
	DATE_NOT_VALID:             "DATE_NOT_VALID",
	VENDOR_SHUTDOWN:            "VENDOR_SHUTDOWN",
	METHOD_ARGUMENTS_NOT_VALID: "METHOD_ARGUMENTS_NOT_VALID",
	TOO_MANY_REQUEST:           "TOO_MANY_REQUEST",
	BAD_REQUEST:                "BAD_REQUEST",
	UNAUTHORIZE:                "UNAUTHORIZE",
}

var respMessage = map[status]string{
	SUCCESS:                    "Success",
	SYSTEM_ERROR:               "Contact Our Team",
	DUPLICATE_DATA:             "Contact Our Team",
	DATA_NOT_EXIST:             "Contact Our Team",
	BIND_ERROR:                 "Contact Our Team",
	RUNTIME_ERROR:              "Contact Our Team",
	DATE_NOT_VALID:             "Contact Our Team",
	VENDOR_SHUTDOWN:            "Contact Our Team",
	METHOD_ARGUMENTS_NOT_VALID: "Contact Our Team",
	TOO_MANY_REQUEST:           "Contact Our Team",
	BAD_REQUEST:                "Bad Request",
	UNAUTHORIZE:                "Unauthorized",
}

var respHttpStatus = map[status]int{
	SUCCESS:                    http.StatusOK,
	SYSTEM_ERROR:               http.StatusInternalServerError,
	DUPLICATE_DATA:             http.StatusInternalServerError,
	DATA_NOT_EXIST:             http.StatusNotFound,
	BIND_ERROR:                 http.StatusBadRequest,
	RUNTIME_ERROR:              http.StatusInternalServerError,
	DATE_NOT_VALID:             http.StatusInternalServerError,
	VENDOR_SHUTDOWN:            http.StatusInternalServerError,
	METHOD_ARGUMENTS_NOT_VALID: http.StatusInternalServerError,
	TOO_MANY_REQUEST:           http.StatusTooManyRequests,
	BAD_REQUEST:                http.StatusBadRequest,
	UNAUTHORIZE:                http.StatusUnauthorized,
}

func (stat status) GetCode() string {
	return respCode[stat]
}

func (stat status) GetStatus() string {
	return respStatus[stat]
}

func (stat status) GetMessage() string {
	return respMessage[stat]
}

func (stat status) GetHttpStatus() int {
	return respHttpStatus[stat]
}
