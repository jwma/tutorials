## echo

非常简单的一个应用，你跟它说什么，它就会返回什么。

## API

|名称|echo|
|----|----|
|Description|你跟它说什么，它就会返回什么|
|Path|`/echo`|
|Method|POST|
|Request|`{"content": "hello world!"}`|
|Response|`{"echo": "hello world!"}`|

|名称|version|
|----|----|
|Description|查看当前 API 版本|
|Path|`/version`|
|Method|Any|
|Request|-|
|Response|`{"version": "v0.0.1"}`|
