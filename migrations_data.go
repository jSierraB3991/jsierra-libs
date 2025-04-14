package eliotlibs

import (
	"log"
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
	for i, migration := range migrations {
		version := FillZeros(2, i+1)
		err := RunMigrate(mInterface, version, migration)
		if err != nil {
			return err
		}
	}
	return nil
}
