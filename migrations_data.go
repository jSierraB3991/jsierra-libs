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

func RunMigrate(mInterface MigrationInterface, version string, migration func() error) {
	exist, err := mInterface.ValidateMigrate(version)
	FinsihApp(err)

	if exist {
		return
	}

	err = migration()
	FinsihApp(err)
	log.Printf("SAVING MIGRATION %s", version)
	err = mInterface.SaveMigration(version)
	FinsihApp(err)
}

func RunMigrations(mInterface MigrationInterface, migrations ...func() error) {
	for _, migration := range migrations {
		version, err := GetNameOfFunction(migration)
		FinsihApp(err)

		RunMigrate(mInterface, version, migration)
	}
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
