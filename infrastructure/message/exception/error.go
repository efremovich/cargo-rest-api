package exception

import (
	"errors"
)

// Common errors.
var (
	// ErrorTextNoRecordInsertedToRedis is an error representing there is no record inserted to redis.
	ErrorTextNoRecordInsertedToRedis = errors.New("api.msg.error.common.no_record_inserted_to_redis")

	// ErrorTextInternalServerError is an error representing internal server error.
	ErrorTextInternalServerError = errors.New("api.msg.error.common.internal_server_error")

	// ErrorTextAnErrorOccurred is an error representing an error occurred.
	ErrorTextAnErrorOccurred = errors.New("api.msg.error.common.an_error_occurred")

	// ErrorTextUnauthorized is an error representing unauthorized request.
	ErrorTextUnauthorized = errors.New("api.msg.error.common.unauthorized")

	// ErrorTextForbidden is an error representing forbidden request.
	ErrorTextForbidden = errors.New("api.msg.error.common.forbidden")

	// ErrorTextBadRequest is an error representing bad request.
	ErrorTextBadRequest = errors.New("api.msg.error.common.bad_request")

	// ErrorTextUnprocessableEntity is an error representing unprocessable entity.
	ErrorTextUnprocessableEntity = errors.New("api.msg.error.common.unprocessable_entity")

	// ErrorTextNotFound is an error representing request not found.
	ErrorTextNotFound = errors.New("api.msg.error.common.not_found")

	// ErrorTextFileTooLarge is an error representing that received file size too large.
	ErrorTextFileTooLarge = errors.New("api.msg.error.common.file_too_large")

	// ErrorTextInvalidPrivateKey is an error representing invalid private key.
	ErrorTextInvalidPrivateKey = errors.New("api.msg.error.common.invalid_private_key")

	// ErrorTextInvalidPublicKey is an error representing invalid public key.
	ErrorTextInvalidPublicKey = errors.New("api.msg.error.common.invalid_public_key")

	// ErrorTextRefreshTokenIsExpired is an error representing refresh token is expired.
	ErrorTextRefreshTokenIsExpired = errors.New("api.msg.error.common.refresh_token_expired")

	// ErrorTextPerPage is an error representing request per page over the limit.
	ErrorTextPerPage = errors.New("api.msg.error.common.per_page")
)

// Errors for document
var (
	// ErrorTextDocumentNotFound is an error representing document not found in database.
	ErrorTextDocumentNotFound = errors.New("api.msg.error.document.not_found")

	// ErrorTextDocumentInvalidUUID is an error representing UUID not found in database.
	ErrorTextDocumentInvalidUUID = errors.New("api.msg.error.document.invalid_uuid")
)

// Errors for role.
var (
	// ErrorTextRoleNotFound is an error representing role not found in database.
	ErrorTextRoleNotFound = errors.New("api.msg.error.role.not_found")

	// ErrorTextRoleInvalidUUID is an error representing UUID not found in database.
	ErrorTextRoleInvalidUUID = errors.New("api.msg.error.role.invalid_uuid")
)

// Errors fot tour.
var (
	// ErrorTextTourSlugAlreadyExists is an error representing slug is already exists in database.
	ErrorTextTourSlugAlreadyExists = errors.New("api.msg.error.tour.slug_already_exists")
)

// Errors for user.
var (
	// ErrorTextUserNotFound is an error representing user not found in database.
	ErrorTextUserNotFound = errors.New("api.msg.error.user.not_found")

	// ErrorTextUserInvalidUUID is an error representing UUID not found in database.
	ErrorTextUserInvalidUUID = errors.New("api.msg.error.user.invalid_uuid")

	// ErrorTextUserInvalidPassword is an error representing hashed password not match with stored in database.
	ErrorTextUserInvalidPassword = errors.New("api.msg.error.user.invalid_password")

	// ErrorTextUserInvalidUsernameAndPassword is an error representing hashed password not match with stored in database.
	ErrorTextUserInvalidUsernameAndPassword = errors.New("api.msg.error.user.invalid_email_and_password")

	// ErrorTextUserEmailNotRegistered is an error representing email already is not exists in database.
	ErrorTextUserEmailNotRegistered = errors.New("api.msg.error.user.email_not_registered")

	// ErrorTextUserPhoneNotRegistered is an error representing email already is not exists in database.
	ErrorTextUserPhoneNotRegistered = errors.New("api.msg.error.user.phone_not_registered")

	// ErrorTextUserEmailAlreadyTaken is an error representing email already exists in database.
	ErrorTextUserEmailAlreadyTaken = errors.New("api.msg.error.user.email_already_taken")

	// ErrorTextUserPhoneAlreadyTaken is an error representing phone already exists in database.
	ErrorTextUserPhoneAlreadyTaken = errors.New("api.msg.error.user.phone_already_taken")

	// ErrorTextUserPreferenceInvalidUUID is an error representing UUID not found in database.
	ErrorTextUserPreferenceInvalidUUID = errors.New("api.msg.error.user.preference.invalid_uuid")

	// ErrorTextUserForgotPasswordTokenNotFound is an error representing Token not found in database.
	ErrorTextUserForgotPasswordTokenNotFound = errors.New("api.msg.error.user.forgot_password.token_not_found")
)

// Errors for storage.
var (
	// ErrorTextStorageCategoryNotFound is an error representing storage_category not found in database.
	ErrorTextStorageCategoryNotFound = errors.New("api.msg.error.storage.category.not_found")

	// ErrorTextStorageFileNotFound is an error representing storage_file not found in database.
	ErrorTextStorageFileNotFound = errors.New("api.msg.error.storage.file.not_found")

	// ErrorTextStorageUploadCannotOpenFile is an error representing uploaded file can not opened by system.
	ErrorTextStorageUploadCannotOpenFile = errors.New("api.msg.error.storage.file.cannot_open_file")

	// ErrorTextStorageUploadInvalidSize is an error representing uploaded file size is greater than allowed maximum size.
	ErrorTextStorageUploadInvalidSize = errors.New("api.msg.error.storage.file.invalid_file_size")

	// ErrorTextStorageUploadInvalidFileType is an error representing uploaded file has invalid file type.
	ErrorTextStorageUploadInvalidFileType = errors.New("api.msg.error.storage.file.invalid_file_type")
)

// Errors for sity.
var (
	// ErrorTextRoleNotFound is an error representing sity not found in database.
	ErrorTextSityNotFound = errors.New("api.msg.error.sity.not_found")

	// ErrorTextRoleInvalidUUID is an error representing UUID not found in database.
	ErrorTextSityInvalidUUID = errors.New("api.msg.error.sity.invalid_uuid")
)

// Errors for vehicle.
var (
	// ErrorTextRoleNotFound is an error representing vehicle not found in database.
	ErrorTextVehicleNotFound = errors.New("api.msg.error.vehicle.not_found")

	// ErrorTextRoleInvalidUUID is an error representing UUID not found in database.
	ErrorTextVehicleInvalidUUID = errors.New("api.msg.error.vehicle.invalid_uuid")
)

// Errors for passenger_type.
var (
	// ErrorTextRoleNotFound is an error representing passenger_type not found in database.
	ErrorTextPassengerTypeNotFound = errors.New("api.msg.error.passenger_type.not_found")

	// ErrorTextRoleInvalidUUID is an error representing UUID not found in database.
	ErrorTextPassengerTypeInvalidUUID = errors.New("api.msg.error.passenger_type.invalid_uuid")
)

// Errors for price.
var (
	// ErrorTextRoleNotFound is an error representing price not found in database.
	ErrorTextPriceNotFound = errors.New("api.msg.error.price.not_found")

	// ErrorTextRoleInvalidUUID is an error representing UUID not found in database.
	ErrorTextPriceInvalidUUID = errors.New("api.msg.error.price.invalid_uuid")
)

// Errors for passenger.
var (
	// ErrorTextRoleNotFound is an error representing passenger not found in database.
	ErrorTextPassengerNotFound = errors.New("api.msg.error.passenger.not_found")

	// ErrorTextRoleInvalidUUID is an error representing UUID not found in database.
	ErrorTextPassengerInvalidUUID = errors.New("api.msg.error.passenger.invalid_uuid")
)

// Errors for document_type.
var (
	// ErrorTextRoleNotFound is an error representing document_type not found in database.
	ErrorTextDocumentTypeNotFound = errors.New("api.msg.error.document_type.not_found")

	// ErrorTextRoleInvalidUUID is an error representing UUID not found in database.
	ErrorTextDocumentTypeInvalidUUID = errors.New("api.msg.error.document_type.invalid_uuid")
)

// Errors for regularity.
var (
	// ErrorTextRegularityTypeNotFound is an error representing regularity not found in database.
	ErrorTextRegularityTypeNotFound = errors.New("api.msg.error.regularity_type.not_found")

	// ErrorTextRegularityTypeInvalidUUID is an error representing UUID not found in database.
	ErrorTextRegularityTypeInvalidUUID = errors.New("api.msg.error.regularity_type.invalid_uuid")
)

// Errors for order_status_type.
var (
	// ErrorTextOrderStatusTypeNotFound is an error representing regularity not found in database.
	ErrorTextOrderStatusTypeNotFound = errors.New("api.msg.error.order_status_type.not_found")

	// ErrorTextOrderStatusTypeInvalidUUID is an error representing UUID not found in database.
	ErrorTextOrderStatusTypeInvalidUUID = errors.New("api.msg.error.order_status_type.invalid_uuid")
)

// Errors for driver.
var (
	// ErrorTextDriverNotFound is an error representing regularity not found in database.
	ErrorTextDriverNotFound = errors.New("api.msg.error.driver.not_found")

	// ErrorTextDriverInvalidUUID is an error representing UUID not found in database.
	ErrorTextDriverInvalidUUID = errors.New("api.msg.error.driver.invalid_uuid")
)

// Errors for route.
var (
	// ErrorTextRouteNotFound is an error representing regularity not found in database.
	ErrorTextRouteNotFound = errors.New("api.msg.error.route.not_found")

	// ErrorTextRouteInvalidUUID is an error representing UUID not found in database.
	ErrorTextRouteInvalidUUID = errors.New("api.msg.error.route.invalid_uuid")
)

// Errors for trip.
var (
	// ErrorTextTripNotFound is an error representing regularity not found in database.
	ErrorTextTripNotFound = errors.New("api.msg.error.trip.not_found")

	// ErrorTextTripInvalidUUID is an error representing UUID not found in database.
	ErrorTextTripInvalidUUID = errors.New("api.msg.error.trip.invalid_uuid")
)

// Errors for order.
var (
	// ErrorTextOrderNotFound is an error representing regularity not found in database.
	ErrorTextOrderNotFound = errors.New("api.msg.error.order.not_found")

	// ErrorTextOrderInvalidUUID is an error representing UUID not found in database.
	ErrorTextOrderInvalidUUID = errors.New("api.msg.error.order.invalid_uuid")
)

// Errors for payment.
var (
	// ErrorTextPaymentNotFound is an error representing regularity not found in database.
	ErrorTextPaymentNotFound = errors.New("api.msg.error.payment.not_found")

	// ErrorTextPaymentInvalidUUID is an error representing UUID not found in database.
	ErrorTextPaymentInvalidUUID = errors.New("api.msg.error.payment.invalid_uuid")
)
