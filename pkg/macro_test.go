package pkg

type PrivateEmailInformationTEST struct {
	Smtp_server string
	Port        int
	Sender      string
	Password    string
	Recver      string
}

func InitPrivateEmailInformationTEST() PrivateEmailInformationTEST {
	return PrivateEmailInformationTEST{
		Smtp_server: "smtp.foobar.com",
		Port:        25,
		Sender:      "foobar@foobar.com",
		Password:    "FoObAr123",
		Recver:      "barfoo@foobar.com",
	}
}
