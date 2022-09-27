package utilities

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func ExampleJob_Error() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("bad time")
	j := s.Jobs()[0]
	fmt.Println(j.Error())
	// Output:
	// the given time format is not supported
}
