package eliotlibs

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"
)

type MigrationInterface interface {
	SaveMigration(version string) error
	ValidateMigrate(version string) (bool, error)
}

func RunMigrate(mInterface MigrationInterface, version string, migration func() error) error {
	exist, err := mInterface.ValidateMigrate(version)
	if err != nil {
		return err
	}

	if exist {
		return err
	}

	err = migration()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SAVING MIGRATION %s", version)
	return mInterface.SaveMigration(version)
}

func RunMigrations(mInterface MigrationInterface, migrations ...func() error) error {
	for _, migration := range migrations {
		version, err := GetNameOfFunction(migration)
		if err != nil {
			return err
		}

		err = RunMigrate(mInterface, version, migration)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetNameOfFunction(fn func() error) (string, error) {
	ptr := reflect.ValueOf(fn).Pointer()
	funcInfo := runtime.FuncForPC(ptr)
	var version string
	if funcInfo == nil {
		return "", fmt.Errorf("no se pudo obtener información de la función")
	}
	fullName := funcInfo.Name()

	lastDot := strings.LastIndex(fullName, ".")
	if lastDot != -1 {
		version = fullName[lastDot+1:]
	} else {
		version = fullName
	}
	return version, nil
}
