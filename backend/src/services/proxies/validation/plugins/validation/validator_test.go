package validation

import (
	"mock/plugins"
	"testing"
	"utils"
)

func TestValidEmail_True(t *testing.T) {
	for _, email := range plugins.ValidEmails {
		utils.AssertTrue(ValidEmail(email), t)
	}
}

func TestValidEmail_False(t *testing.T) {
	for _, email := range plugins.InvalidEmails {
		utils.AssertFalse(ValidEmail(email), t)
	}
}

func TestValidNickname_True(t *testing.T) {
	for _, nickname := range plugins.ValidNicknames {
		utils.AssertTrue(ValidNickname(nickname), t)
	}
}

func TestValidNickname_False(t *testing.T) {
	for _, nickname := range plugins.InValidNicknames {
		utils.AssertFalse(ValidNickname(nickname), t)
	}
}

func TestValidLongitude_True(t *testing.T) {
	for _, l := range plugins.ValidLongitudes {
		utils.AssertTrue(ValidLongitude(l), t)
	}
}

func TestValidLongitude_False(t *testing.T) {
	for _, l := range plugins.InvalidLongitudes {
		utils.AssertFalse(ValidLongitude(l), t)
	}
}

func TestValidLatitude_True(t *testing.T) {
	for _, l := range plugins.ValidLatitudes {
		utils.AssertTrue(ValidLatitude(l), t)
	}
}

func TestValidLatitude_False(t *testing.T) {
	for _, l := range plugins.InvalidLatitudes {
		utils.AssertFalse(ValidLatitude(l), t)
	}
}

func TestValidWholePositiveNumber_True(t *testing.T) {
	for _, n := range plugins.ValidWholePositiveNumbers {
		utils.AssertTrue(ValidWholePositiveNumber(n), t)
	}
}

func TestValidWholePositiveNumber_False(t *testing.T) {
	for _, n := range plugins.InvalidWholePositiveNumbers {
		utils.AssertFalse(ValidWholePositiveNumber(n), t)
	}
}

func TestValidPassword_True(t *testing.T) {
	for _, p := range plugins.ValidPasswords {
		utils.AssertTrue(ValidPassword(p), t)
	}
}

func TestValidPassword_False(t *testing.T) {
	for _, p := range plugins.InvalidPasswords {
		utils.AssertFalse(ValidPassword(p), t)
	}
}

func TestValidTitle_True(t *testing.T) {
	for _, title := range plugins.ValidTitles {
		utils.AssertTrue(ValidTitle(title), t)
	}
}

func TestValidTitle_False(t *testing.T) {
	for _, title := range plugins.InvalidTitles {
		utils.AssertFalse(ValidTitle(title), t)
	}
}

func TestValidDescription_True(t *testing.T) {
	for _, d := range plugins.ValidDescriptions {
		utils.AssertTrue(ValidDescription(d), t)
	}
}

func TestValidDescription_False(t *testing.T) {
	for _, d := range plugins.InvalidDescriptions {
		utils.AssertFalse(ValidDescription(d), t)
	}
}

func TestValidName_True(t *testing.T) {
	for _, n := range plugins.ValidNames {
		utils.AssertTrue(ValidName(n), t)
	}
}

func TestValidName_False(t *testing.T) {
	for _, n := range plugins.InvalidNames {
		utils.AssertFalse(ValidName(n), t)
	}
}

func TestValidDate_True(t *testing.T) {
	for _, d := range plugins.ValidDates {
		utils.AssertTrue(ValidDate(d), t)
	}
}

func TestValidDate_False(t *testing.T) {
	for _, d := range plugins.InvalidDates {
		utils.AssertFalse(ValidDate(d), t)
	}
}

func TestValidGender_True(t *testing.T) {
	for _, g := range plugins.ValidGenders {
		utils.AssertTrue(ValidGender(g), t)
	}
}

func TestValidGender_False(t *testing.T) {
	for _, g := range plugins.InvalidGenders {
		utils.AssertFalse(ValidGender(g), t)
	}
}
