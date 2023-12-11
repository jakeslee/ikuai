# IKuai SDK

自用 IKuai SDK，封装了部分爱快的接口。

### 依赖

```shell
go get -u https://github.com/jakeslee/ikuai
```

### 使用

```go
i := NewIKuai("http://10.10.1.253", "test", "test123", true, false)

login, err := i.Login()
if err != nil {
    log.Error(err)
}

result, err := client.ShowSysStat()
if err != nil {
    return
}

log.Printf("%+v", result)

```

