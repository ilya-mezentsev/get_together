package plugins

import "strings"

var (
	ValidEmails = []string{
		"mail@ya.ru", "hello.world@gmail.com", "wtf@mail.ru",
		"my_mail@yandex.ru", "12-13@ya.ru", "hey-1@mail.ru",
	}
	InvalidEmails = []string{
		"123", "just_do_it!", "hello@world", "@ya.ru", "mail@.ru",
		"wtf@gmail.", "",
	}
	ValidNicknames = []string{
		"mather_fucker", "hello-world", "valid$nick", "killer228",
	}
	InValidNicknames = []string{
		"hello*world", "   ", "too_long_nickname", "nea<", "invalid^too",
		"!!", "hello.", "it's me", "how r u?", "u&me",
	}
	ValidLongitudes = []float64{ // -180 <= longitude <= 180
		0.0, -180, 180, 99.2,
	}
	InvalidLongitudes = []float64{
		-181, 181, 359, 460, -190,
	}
	ValidLatitudes = []float64 { // -90 <= latitude <= 90
		0.0, -90, 90, 89, 55.0,
	}
	InvalidLatitudes = []float64 {
		-91, 92, -180, 355, -322,
	}
	ValidWholePositiveNumbers = []float64{
		1, 15, 28, 305, 14500,
	}
	InvalidWholePositiveNumbers = []float64{
		0, -1, -30, -500,
	}
	ValidPasswords = []string{
		"mYStRoNg*PwD12", "hello&world", "another*strong$password",
		"p%w%d_password", "good-password",
	}
	InvalidPasswords = []string{
		"short", "bad|char", "one=more", "one+two+three", "1/20000000",
		"very_very_very_long_password_invalid", "", "                 ",
	}
	ValidTitles = []string{
		"This is the title", "Заголовок встречи", "Привет мир 21",
		"Встреча 42", "28 friends", "title_with_underscore", "what_$",
		"Вы готовы, дети?", "Саша - п*др (шутка)",
	}
	InvalidTitles = []string{
		"", "        ", "1", "bad_<<", "hello/world", "bad^title",
		"плохой&заголовок",
	}
	ValidDescriptions = []string{
		"Some simple text. Can contain underscores _ (for i.e.) or numbers 1, 2",
		"Большое описание встречи. Содержит, по мимо всего прочего, интересные символы вроде: *()?$%",
	}
	InvalidDescriptions = []string{
		"too short", strings.Repeat("very long", 200), "here we are<",
	}
	ValidNames = []string{
		"Иван Иванов", "John Doe", "Илья М", "tag1", "тег встречи",
	}
	InvalidNames = []string{
		"", "     ", "1", "bad_<<", "hello/world", "bad^name",
		"также невалиден!", strings.Repeat("too long", 10),
	}
	ValidDates = []string{
		"17-08-2016 15:11:06", "29-02-2020 12:00:01", "01-02-1970 00:00:01",
		"01-01-1900 12:30:45", "01-01-0001 12:10:31",
	}
	InvalidDates = []string{
		"", "17.08.2016 15:11:06", "29-02-2020", "1-2-1970 00:00:01",
		"01-01-1900 12:30", "01-01-0001 12:10:31 +0000",
	}
	ValidGenders = []string{
		"male", "female", "",
	}
	InvalidGenders = []string{
		"bad", "мальчик", "Саша",
	}
)
