package error

type LongUrlAlreadyExist struct {
	LongUrl string
}

func (e LongUrlAlreadyExist) Error() string {
	return "Long url already exist"
}

type ShortUrlAlreadyExist struct {
	ShortUrl string
}

func (e ShortUrlAlreadyExist) Error() string {
	return "Short url already exist"
}
