package enum

type UserStatus string

const (
	UserStatusNormal  UserStatus = "normal"
	UserStatusDisable UserStatus = "disable"
)

type UserGender string

const (
	UserGenderMale   = "male"
	UserGenderFeMale = "female"
)

var (
	userStatusMap = map[int]UserStatus{
		1: UserStatusNormal,
		2: UserStatusDisable,
	}

	userGenderMap = map[int]UserGender{
		1: UserGenderMale,
		2: UserGenderFeMale,
	}
)

func GetUserStatus(status int) UserStatus {
	return userStatusMap[status]
}

func GetUserGender(status int) UserGender {
	return userGenderMap[status]
}
