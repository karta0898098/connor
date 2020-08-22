connor
=====
connor為基礎程式產生器，方便快速產出一套基礎設施框架，可以讓使用者專注開發業務邏輯。可以選擇http engine(gin,echo)
目前預設的為echo，gin框架依然可以使用，不過還欠缺自動生成route的功能

配置輸入使用了viper，所以可以選擇使用yaml或是toml配置文檔，預設使用toml

也提供了新增controller，以及根據user輸入的model來決定是否產生
repository或service層級CRUD程式碼。

本程式產生的程式碼，都用上了uber fx的依賴注入，所以不需要自動實體化controller,service,repository等

創建指令:
```
connor init -n project-name
```

新增controller:
```
connor add controller -n controllerName
```

根據Model新增CRUD:

輸入--repo會新增repository 輸入--srv會新增service
```
connor add entity -f ./pkg/model/xxxx.go --repo --srv
```



