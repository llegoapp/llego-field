package owner

import "golang.org/x/crypto/bcrypt"

type Owner struct {
	Id          int
	Password    string
	Name        string
	PhoneNumber PhoneNumber
}

func New(id int, name, password string) (*Owner, error) {
	hashedPassword, err := Password(password).Hash()
	if err != nil {
		return nil, err
	}

	return &Owner{
		Id:       id,
		Name:     name,
		Password: hashedPassword,
	}, nil
}

type PhoneNumber struct {
	CountryCode string
	Number      string
}

type Password string

func (p Password) Hash() (string, error) {
	// Generate a hashed version of the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (p Password) Compare(password string) bool {
	// Compare the provided password with the stored hashed password
	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(password))
	return err == nil
}
