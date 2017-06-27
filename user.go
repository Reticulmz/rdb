package rdb

// User is an user on Ripple.
type User struct {
	ID           int           `json:"id"`
	Username     string        `json:"username"`
	Password     string        `json:"password" gorm:"column:password_md5"`
	Salt         string        `json:"salt"`
	Email        string        `json:"email"`
	RegisteredOn UnixTimestamp `json:"registered_on" gorm:"column:register_datetime"`
	Privileges   int           `json:"privileges"`
	DonorExpire  UnixTimestamp `json:"donor_expire"`
	BannedOn     UnixTimestamp `json:"banned_on" gorm:"column:ban_datetime"`
	AQN          bool          `json:"aqn" gorm:"column:aqn"`
	// Technically a row, but since it's LONGTEXT we don't want to fetch it for every SELECT query.
	// Notes string `json:"notes"`
	LatestActivity  UnixTimestamp `json:"latest_activity"`
	SilenceEnd      UnixTimestamp `json:"silence_end"`
	SilenceReason   string        `json:"silence_reason"`
	PasswordVersion int           `json:"password_version"`
	UsernameSafe    string        `json:"username_safe"`
	Flags           int           `json:"flags"`
}

// GetUser retrieves an user knowing its ID.
func GetUser(id int) *User {
	var u User
	if DB.First(&u, id).RecordNotFound() {
		return nil
	}
	return &u
}

// GetUserByUsername retrieves an user knowing its username.
func GetUserByUsername(s string) *User {
	var u User
	if DB.Where("username = ?", s).First(&u).RecordNotFound() {
		return nil
	}
	return &u
}

// GetUserByUsernameSafe retrieves an user knowing its username_safe.
func GetUserByUsernameSafe(s string) *User {
	var u User
	if DB.Where("username_safe = ?", s).First(&u).RecordNotFound() {
		return nil
	}
	return &u
}

// GetUserByEmail retrieves an user knowing its email.
func GetUserByEmail(s string) *User {
	var u User
	if DB.Where("email = ?", s).First(&u).RecordNotFound() {
		return nil
	}
	return &u
}
