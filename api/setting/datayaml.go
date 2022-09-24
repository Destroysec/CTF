package config

type Data_Config struct {
	Debug       bool
	Senderemail struct {
		Email    string
		Password string
	}
	Database struct {
		Url          string
		Collection   string
		Recollection string
	}
	Pay struct {
		Numberphone string
	}
}
