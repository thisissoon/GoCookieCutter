package version

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

var (
	Version      string
	Timestamp    string
	GitCommit    string
	GitTreeState string
)

// BuildTime returns the build timestamp
func BuildTime() time.Time {
	ts, _ := strconv.ParseInt(Timestamp, 10, 64)
	return time.Unix(ts, 0).UTC()
}

// Write writes version info to an io.Writer
func Write(w io.Writer) {
	fmt.Fprintln(w, "Version:", Version)
	fmt.Fprintln(w, "Build Time:", BuildTime().Format(time.RFC1123))
	fmt.Fprintln(w, "Git Commit:", GitCommit)
	fmt.Fprintln(w, "Git Tree State:", GitTreeState)
}
