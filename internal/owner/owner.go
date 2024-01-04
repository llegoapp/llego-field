package owner

import "fields/pkg/apperror"
import "golang.org/x/crypto/bcrypt"

type Owner struct {
	Id          int
	Password    Password
	Name        Name
	PhoneNumber PhoneNumber
}

func New(id int, name, password string) (*Owner, error) {
	hashedPassword, err := Password(password).Hash()
	if err != nil {
		return nil, err
	}

	return &Owner{
		Id:       id,
		Name:     Name(name),
		Password: Password(hashedPassword),
	}, nil
}

func (o *Owner) Vaildate() error {
	if err := o.Name.Validate(); err != nil {
		return err
	}
	return nil
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

type Name string

func (n Name) Validate() error {
	if n == "" {
		return apperror.NewValidationError("Name can't be empty")
	}
	return nil
}
