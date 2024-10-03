package constant

import "time"

const (
	REGEX_EMAIL                          = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	BLANK                                = ""
	DEFAULT_INTEGER                      = 0
	MAX_TIME_TOKEN                       = time.Hour * 24
	ORDER_PAGE_SIZE                      = 10
	DEFAULT_PAGE_SIZE                    = 10
	MAX_CHECK_EXPIRING_CUSTOMER_PER_TIME = 100
)

const (
	DEFAULT_SHEET_NAME                  = "Good Infos"
	DEFAULT_SPREADSHEET_NAME            = "Shipments"
	DEFAULT_STATUS_COLUMN_NAME          = "status"
	DEFAULT_TRACKING_NUMBER_COLUMN_NAME = "tracking number"
	DEFAULT_LABEL_COLUMN_NAME           = "labels"
	DEFAULT_NOTE_COLUMN_NAME            = "note"
	HPW_CLIENT_ID                       = "X-HPW-Client-Id"
	HPW_SERVICE_TYPE                    = "X-HPW-Service-Type"

	DEFAULT_SHEET_ID           = 0
	DEFAULT_START_ROW_INDEX    = 0
	DEFAULT_END_ROW_INDEX      = 1
	DEFAULT_START_COLUMN_INDEX = 0
	DEFAULT_CELL_PADDING       = 30
)

const (
	OAUTH_DEFAULT_STATE = "state-token"
)

// Country code
const (
	COUNTRY_CODE_US = "US"
)

// Create shipment by hpw
const (
	STATE_CA       = "CA"
	MAX_CANDIDATES = 5
)

const GOOGLE_SHEET_LINK_FORMAT = "https://docs.google.com/spreadsheets/d/%s/edit"

// order
const (
	ORDER_SOURCE_HPW_BOT_GOOGLE_SHEET = "bot_google_sheet"
)

const (
	FORMAT_TIME_STANDARD                 = "02-01-2006 15:04:05"
	FORMAT_TIME_IMPORT_EXPIRING_CUSTOMER = "1/2/2006 15:04"
	FORMAT_DATABASE_TIME                 = "2006-01-02 15:04:05"
	FORMAT_DATE_ONLY                     = "02/01/2006"
)

const (
	STATUS_SUCCESS = "status-success"
	STATUS_FAILED  = "status-failed"
)

const DASHBOARD_DATA_KEY = "dashboard-layout"

const DEFAULT_NORMAL_AGENT_ID = 1

const LENGTH_PRODUCT_IMPORT_TEMPLATE = 14

const LANGUAGE_SUFFIX = "language.json"

const DEFAULT_RECOMMEND_PRODUCT_COUNT = 8

const (
	SELECTED_ALL = "ALL"
)

const (
	OPENAI_ROLE_USER                   = "user"
	OPENAI_ROLE_ASSISTANT              = "assistant"
	OPENAI_NUMBER_OF_MESSAGE_IN_THREAD = 2
	OPENAI_EXPIRED_TIME_OF_THREAD      = time.Hour * 2
)

const (
	ZALO_EVENT_ANONYMOUS_SEND_TEXT = "anonymous_send_text"
	ZALO_EVENT_USER_SEND_TEXT      = "user_send_text"
)

const (
	MAX_TIME_CHECK_EXPIRING_CUSTOMER = time.Hour * 24 * 5
)

const (
	UPLOAD_FILE_LOG_PATH = "upload_file_log"
)

const (
	EXPIRE_DATE_TODAY       = "expire-date-today"
	EXPIRE_DATE_EXPIRED     = "expire-date-expired"
	EXPIRE_DATE_NOT_EXPIRED = "expire-date-not-expired"
)

const MAX_EXPIRING_CUSTOMER = 2000
