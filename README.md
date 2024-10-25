# strong-defi-backend

This is GIN strong-defi-backend.


# DEPLOY
```
 前期部署命令：
    1、创建go项目，设置GOROOT GOPATH GO MOdule:environment设置：GOPROXY=https://goproxy.cn,direct
    2、 如果使用的是vscode要执行如下命令
        1、go mod init {project name}
        2、go mod tidy  清理未使用的依赖，添加未使用的依赖
    3、引入gin：go get -u github.com/gin-gonic/gin@v1.10.0
    4、引入gorm：go get -u gorm.io/gorm@v1.25.12
    5、引入mysql：go get -u gorm.io/driver/mysql@v1.5.7
    6、引入redis：go get -u github.com/go-redis/redis@v6.15.9   暂时还不需要
    7、引入go-ehtereum：go get -u github.com/ethereum/go-ethereum@v1.14.9
    8、引入go-rpc：go get github.com/ethereum/go-ethereum/rpc@v1.14.9版本，这里的版本，跟上面go-ethereum要一致，一定要引入这个包，不然连不上链节点
    9、引入lumberjack：go get gopkg.in/natefinch/lumberjack.v2,在 Go 语言中，lumberjack 包是一个日志轮转库，允许程序开发者轻松管理日志文件。它的主要功能包括：
        1、设置最大文件大小: 一旦日志文件达到指定的大小，新的日志将会被写入到一个新的文件。
        2、保留备份: 允许设定保留的旧日志文件的最大数量，以便进行管理。
        3、设置最大保留天数: 允许设定旧日志文件的保留天数，超出天数的日志文件会被自动删除。
        4、压缩功能: 可以选择是否对旧日志文件进行压缩，以节省存储空间。
     10、引入viper：go get github.com/spf13/viper，Viper是Go应用程序的完整配置解决方案，它支持:读取JSON, TOML, YAML, HCL, envfile和Java属性配置文件
     11、引入jwt：go get github.com/golang-jwt/jwt  
     12、引入fastjson：go get github.com/valyala/fastjson


合约转换go命令：
    1、通过hardhat进行编辑，注意，只需要复制abi那部分json即可
    2、abigen --abi={abi的绝对路径} --contract=contract --out={输出的文件名，可以加上路径}
    3、abigen --abi=abi/SCHStake.abi --pkg=contract --out=contract/schStake/SCHStake.go
    4、abigen --abi=abi/SCTToken_abi.abi --pkg=contract --out=contract/scToken/SCTToken.go
```
# 项目结构
```
common包：放入一写通用的代码，如枚举等，response状态码等
config包：配置文件，配置文件放在config文件夹下，通过viper读取
router：路由包相当于java的controller，封装了接口，通过gin框架实现
model包：封装了数据库的model，专门放置一些数据库表的model ，或者一些其他的model，通过gorm实现
req包：封装了请求参数，所以请求的实体创建都放在这里
resp包：封装了返回参数，所以返回的实体创建都放在这里
service包：封装了业务逻辑
utils包：自己写的一些工具类放入这里
resources：放入一些配置文件

```