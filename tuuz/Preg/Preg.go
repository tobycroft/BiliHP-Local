package Preg

import (
	"fmt"
	"regexp"
)

func Exp(exp string) (*regexp.Regexp, error) {
	reg, err := regexp.Compile(exp)
	return reg, err
}

func Match(exp string, str string) (string, error) {
	reg, err := Exp(exp)
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	return reg.FindString(str), err
}

func MatchOwn(exp string, str *string) (string, error) {
	reg, err := Exp(exp)
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	*str = reg.FindString(*str)
	return "", err
}

func MatchAll(exp string, str string) ([]string, error) {
	reg, err := Exp(exp)
	if err != nil {
		fmt.Print(err)
		return []string{""}, err
	}
	return reg.FindAllString(str, -1), err
}

func IsMatched(exp string, str string) bool {
	reg, err := Exp(exp)
	if err != nil {
		return false
	}
	if reg.MatchString(str) {
		return true
	} else {
		return false
	}
}

func FilterOwn(exp string, str *string) (bool, error) {
	reg, err := Exp(exp)
	if err != nil {
		return true, err
	}

	*str = reg.ReplaceAllString(*str, "")
	return true, err
}

func Filter(exp string, str string) (string, error) {
	reg, err := Exp(exp)
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	return reg.ReplaceAllString(str, ""), err

}
