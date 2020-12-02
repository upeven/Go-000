---
title: week2学习总结
created: '2020-12-02T08:20:06.000Z'
modified: '2020-12-02T08:27:00.440Z'
---

# week2学习总结
### wrap error的要点：
 1、防止代码中出现大量日志信息
 2、每一层都对error进行wrap，记录堆栈信息，方便后面排查错误
 3、进行error判定（需要返回指针，而不是返回值）

Notes:因此在编写基础设施代码时，需要对error进行相应的wrap（记录堆栈信息），并返回给调用者

### panic
Go中的panic意味着非正常错误，不能要求调用者要进行revocer，需要自行处理，因此不要抛出panic。
