package waitlist

type messages struct {
	JoinEmailAlreadyExists    string
	JoinFailed                string
	LeaveFailed               string
	LeaveTokenNotFound        string
	CheckFailed               string
	WaitlistJoinedMailSubject string
	WaitlistLeavedMailSubject string
}

var I18nMessages = messages{
	JoinEmailAlreadyExists:    "waitlist_join_email_already_exists",
	JoinFailed:                "waitlist_join_failed",
	LeaveFailed:               "waitlist_leave_failed",
	LeaveTokenNotFound:        "waitlist_leave_token_not_found",
	CheckFailed:               "waitlist_check_failed",
	WaitlistJoinedMailSubject: "waitlist_joined_mail_subject",
	WaitlistLeavedMailSubject: "waitlist_leaved_mail_subject",
}
