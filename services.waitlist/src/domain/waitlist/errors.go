package waitlist

import "github.com/mixarchitecture/i18np"

type Errors interface {
	JoinEmailAlreadyExists() *i18np.Error
	JoinFailed() *i18np.Error
	LeaveFailed() *i18np.Error
	LeaveTokenNotFound() *i18np.Error
	CheckFailed() *i18np.Error
}

type waitlistErrors struct{}

func newWaitlistErrors() Errors {
	return &waitlistErrors{}
}

func (e *waitlistErrors) JoinEmailAlreadyExists() *i18np.Error {
	return i18np.NewError(I18nMessages.JoinEmailAlreadyExists)
}

func (e *waitlistErrors) JoinFailed() *i18np.Error {
	return i18np.NewError(I18nMessages.JoinFailed)
}

func (e *waitlistErrors) LeaveFailed() *i18np.Error {
	return i18np.NewError(I18nMessages.LeaveFailed)
}

func (e *waitlistErrors) LeaveTokenNotFound() *i18np.Error {
	return i18np.NewError(I18nMessages.LeaveTokenNotFound)
}

func (e *waitlistErrors) CheckFailed() *i18np.Error {
	return i18np.NewError(I18nMessages.CheckFailed)
}
