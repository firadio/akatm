package errorx

const (
	// 通用 1xxx
	CodeSuccess           int32 = 0
	CodeInternalError     int32 = 1001
	CodeParamError        int32 = 1002
	CodeUnauthorized      int32 = 1003
	CodeForbidden         int32 = 1004
	CodeNotFound          int32 = 1005
	CodeSignError         int32 = 1006
	CodeTimestampExpired  int32 = 1007
	CodeRateLimitExceeded int32 = 1008

	// 认证 2xxx
	CodeLoginFailed    int32 = 2001
	CodeCaptchaError   int32 = 2002
	CodeCaptchaExpired int32 = 2003
	CodeTokenInvalid   int32 = 2004
	CodeTokenExpired   int32 = 2005
	CodePasswordError  int32 = 2006
	CodeEmailCodeError int32 = 2007

	// 用户 3xxx
	CodeUserNotFound  int32 = 3001
	CodeUserExist     int32 = 3002
	CodeInviteInvalid int32 = 3003
	CodeInviteExpired int32 = 3004
	CodeInviteUsed    int32 = 3005

	// 客户 4xxx
	CodeCustomerNotFound int32 = 4001
	CodeCustomerExist    int32 = 4002

	// 账户 5xxx
	CodeAccountNotFound int32 = 5001
	CodeAccountExist    int32 = 5002
	CodeAccountFrozen   int32 = 5003

	// 钱包 6xxx
	CodeWalletNotFound      int32 = 6001
	CodeInsufficientBalance int32 = 6002
	CodeAmountInvalid       int32 = 6003

	// 提现 7xxx
	CodeWithdrawalNotFound   int32 = 7001
	CodeWithdrawalCanceled   int32 = 7002
	CodeWithdrawalProcessing int32 = 7003

	// 审核 8xxx
	CodeAuditNotFound  int32 = 8001
	CodeAuditProcessed int32 = 8002
)
