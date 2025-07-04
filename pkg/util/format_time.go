package util

import (
	"database/sql/driver"
	"time"
)

const timeLayout = "2006-01-02 15:04:05"

type FormatTime time.Time

func (t FormatTime) MarshalJSON() ([]byte, error) {
	// 如果是0值
	if t.String() == "0001-01-01 00:00:00" {
		return []byte(`""`), nil
	}
	b := make([]byte, 0, len(timeLayout)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeLayout)
	b = append(b, '"')
	return b, nil
}

func (t *FormatTime) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = FormatTime(time.Time{})
		return
	}

	// 指定解析的格式
	now, err := time.Parse(`"`+timeLayout+`"`, string(data))
	*t = FormatTime(now)
	return
}

func (t FormatTime) ToTime() time.Time {
	return time.Time(t)
}

// 写入 mysql 时调用
func (t FormatTime) Value() (driver.Value, error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(timeLayout)), nil
}

// 检出 mysql 时调用
func (t *FormatTime) Scan(v any) error {
	// mysql 内部日期的格式可能是 2006-01-02 15:04:05 +0800 CST 格式，所以检出的时候还需要进行一次格式化
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = FormatTime(tTime)
	return nil
}

// 用于 fmt.Println 和后续验证场景
func (t FormatTime) String() string {
	return time.Time(t).Format(timeLayout)
}
