package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	hash := HashPwd("123456")
	fmt.Println(hash)
}

func TestCheckPwd(t *testing.T) {
	// $2a$04$4WZ.eaYBWzhUbc/C39ajfeb2ZfyVPU0/fB6tfraixbVO1TfmKln72
	// $2a$04$M7tOkKqOXZwKw9jbti7IQe1UJqNwZWvlpBUXdib19ccdB4pVeOjo6

	ok := CheckPwd("$2a$04$4WZ.eaYBWzhUbc/C39ajfeb2ZfyVPU0/fB6tfraixbVO1TfmKln72", "1234561")
	fmt.Println(ok)

}
