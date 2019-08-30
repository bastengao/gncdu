package scan

import "fmt"

const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
	PB = TB * 1024
)

func ToHumanSize(s int64) string {
	switch {
	case s < KB:
		return fmt.Sprintf("%d  B", s)
	case s < MB:
		return fmt.Sprintf("%.2f KB", float64(s)/float64(KB))
	case s < GB:
		return fmt.Sprintf("%.2f MB", float64(s)/float64(MB))
	case s < TB:
		return fmt.Sprintf("%.2f GB", float64(s)/float64(GB))
	case s < PB:
		return fmt.Sprintf("%.2f TB", float64(s)/float64(TB))
	default:
		return fmt.Sprintf("%d B", s)
	}
}
