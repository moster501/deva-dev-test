package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

func GetSumPower(pwType string) (int, error) {
	var sumResult int
	statement := fmt.Sprintf("SELECT SUM(%v) FROM deva_test", pwType)
	err := DBMySQL.QueryRow(statement).Scan(&sumResult)
	if err != nil {
		log.Println(err.Error())
		return sumResult, err
	}
	return sumResult, nil
}

func SetRandomData() error {
	for i := 0; i < 1000; i++ {
		statement := fmt.Sprintf("INSERT INTO deva_test (active_power,power_input) VALUES (%d,%d)", RandNumber(), RandNumber())
		result, err := DBMySQL.Exec(statement)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		result.RowsAffected()
	}
	return nil
}

func RandNumber() int {
	min := 1
	max := 1000
	rnum, trerr := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if trerr != nil {
		log.Println(trerr)
		return 0
	}
	return min + int(rnum.Int64())
}
