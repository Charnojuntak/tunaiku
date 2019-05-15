package domain

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Charnojuntak/tunaiku/user/repository"
)

const (
	ErrKTPNumberNotValid = "ktp number isn't valid"
)

func AddUserData(user *repository.User) error {
	isValid := validateKtp(user.KTP)
	if !isValid {
		return errors.New(ErrKTPNumberNotValid)
	}

	err := repository.InsertUserData(user)
	if err != nil {
		return err
	}

	return nil
}

func validateKtp(ktpNumber string) bool {

	provCode, _ := strconv.Atoi(fmt.Sprintf("%s%s", ktpNumber[0], ktpNumber[1]))
	cityCode, _ := strconv.Atoi(fmt.Sprintf("%s%s", ktpNumber[2], ktpNumber[3]))
	districtCode, _ := strconv.Atoi(fmt.Sprintf("%s%s", ktpNumber[4], ktpNumber[5]))

	if checkDisrict(districtCode) && checkCity(cityCode) && checkProvince(provCode) {
		return true
	}

	return false
}

func checkDisrict(input int) bool {
	district := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68}
	for _, val := range district {
		if val == input {
			return true
		}
	}
	return false
}

func checkCity(input int) bool {
	city := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 71, 72, 73, 74, 75, 76, 77, 78, 79}
	for _, val := range city {
		if val == input {
			return true
		}
	}
	return false
}

func checkProvince(input int) bool {

	provinsi := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 31, 32, 33, 34, 35, 36, 51, 52, 53, 61, 62, 63, 64, 71, 73, 74, 75, 76, 81, 82, 94, 91, 72, 65}
	for _, val := range provinsi {
		if val == input {
			return true
		}
	}
	return false
}
