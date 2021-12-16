package customer

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	ErrName            = errors.New("name is required")
	ErrEmail           = errors.New("email is required")
	ErrEmailInvalid    = errors.New("email is invalid")
	ErrDocument        = errors.New("document is required")
	ErrDocumentInvalid = errors.New("document is invalid")
	ErrAddressStreet   = errors.New("address street is required")
	ErrAddressNumber   = errors.New("address number is required")
	ErrAddressCity     = errors.New("address city is required")
	ErrAddressState    = errors.New("address state is required")
	ErrAddressPostCode = errors.New("address post code is required")

	// EmailRgx see https://html.spec.whatwg.org/#valid-e-mail-address
	EmailRgx = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// Validate checks whether the Customer payload has all required fields
func (c *Customer) Validate() []error {
	var errs []error

	if c.Name == "" {
		errs = append(errs, ErrName)
	}

	if c.Email == "" {
		errs = append(errs, ErrEmail)
	} else if !EmailRgx.MatchString(c.Email) {
		errs = append(errs, ErrEmailInvalid)
	}

	if c.Document == "" {
		errs = append(errs, ErrDocument)
	} else if !isCPF(c.Document) {
		errs = append(errs, ErrDocumentInvalid)
	}

	if err := c.Address.Validate(); err != nil {
		errs = append(errs, err...)
	}

	return errs
}

// Validate checks whether the Address payload has all required fields
func (a *Address) Validate() []error {
	var errs []error

	if a.Street == "" {
		errs = append(errs, ErrAddressStreet)
	}

	if a.Number == "" {
		errs = append(errs, ErrAddressNumber)
	}

	if a.City == "" {
		errs = append(errs, ErrAddressCity)
	}

	if a.State == "" {
		errs = append(errs, ErrAddressState)
	}

	if a.PostCode == "" {
		errs = append(errs, ErrAddressPostCode)
	}

	return errs
}

// isCPF checks whether it's a valid CPF
// see https://dicasdeprogramacao.com.br/algoritmo-para-validar-cpf/
func isCPF(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	for i := 0; i < 10; i++ {
		invalid := strings.Repeat(fmt.Sprintf("%d", i), 11)
		if cpf == invalid {
			return false
		}
	}

	if ok := digitCheckerCPF(cpf, 9); !ok {
		return false
	}

	if ok := digitCheckerCPF(cpf, 10); !ok {
		return false
	}

	return true
}

// digitCheckerCPF checks whether the calculation of the digits
// match with verification digit
func digitCheckerCPF(cpf string, digs int) bool {
	var total int
	seqAt, digits := digs+1, cpf[:digs]

	for _, dig := range digits {
		total += int(dig-'0') * seqAt
		seqAt -= 1
	}

	checker := (total * 10) % 11
	if checker == 10 {
		checker = 0
	}

	verifier := int(cpf[digs] - '0')
	return checker == verifier
}
