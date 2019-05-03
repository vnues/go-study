package main
import (
	"fmt"
	"testing"
)
func TestPrint(t *testing.T){
	 res := Print1to20()
     fmt.Println("hello")
	 if res != 210{
		t.Errorf("error")
	   }
}

func Print1to20() int{
	 return 210
}


func testPrint(t *testing.T){
	 fmt.Println("test")
}