# 列出golang常用的计算时间
统一返回 time.Time


```go
// 获取本周第一天
MondayCurrentWeek() time.Time

// 获取本周最后一天
SundayCurrentWeek() time.Time

// 获取本月第一天
FirstDayCurrentMonth() time.Time

// 获取本月最后一天
LastDayCurrentMonth() time.Time

// 获取本年第一天
FirstDayCurrentYear() time.Time

// 获取本年最后一天
LastDayCurrentYear() time.Time
```
