# GoRealTimeChatWithVUE
这是一个我用于学习go语言的仓库，该仓库包含了一个使用Go语言编写的实时聊天项目，其功能包含私聊、群聊、表情包消息（群聊不可用）、文件语音消息（暂时有点问题）等基本的IM功能。学习的项目的地址为：https://github.com/IM-Tools/api-service，非常感谢该作者的分享。


# 如何运行？
## GO后端配置项

### 修改Config

首先打开 GoRealTimeChat/Configs/configs.yaml 文件

以下是配置文件中各个配置项的介绍：
- **Server**：服务器配置项
  - **Name**：服务器名称，此处为 'Im-Services'
  - **Listen**：服务器监听的端口号，此处为 ':8000'（gin http服务端口）
  - **Mode**：服务器运行模式，可选值为 'debug', 'release', 'test'，此处为 'debug'
  - **Env**：服务器环境，可选值为 'local', 'production', 'testing'，此处为 'local'
  - **Lang**：服务器语言设置，此处为 'zh'（中文）
  - **CoroutinePoll**：协程池数量，表示同时启动的协程数量，此处为 100000
  - **Node**：当前服务集群节点，此处为 'localhost:9505'
  - **ServiceOpen**：是否开启集群服务，此处为 false
  - **GrpcListen**：gRPC端口号，此处为 ':8002'
  - **FilePath**：文件保存路径，此处为 "H:\PROJECT\goSTU\GoRealTimeChat\asserts\UploadFile"

- **MySQL**：MySQL数据库配置项
  - **host**：MySQL主机地址，此处为 "127.0.0.1"
  - **port**：MySQL端口号，此处为 3306
  - **database**：数据库名称，此处为 "grtc"
  - **username**：数据库用户名，此处为 "root"
  - **password**：数据库密码，此处为 ******
  - **charset**：数据库字符集，此处为 "utf8mb4"

- **Log**：日志配置项
  - **level**：日志级别，可选值为 'debug', 'info', 'error'，此处为 'debug'
  - **type**：日志文件类型，可选值为 'single'（独立的文件）、'daily'（按日期每天一个），此处为 'daily'
  - **filename**：日志文件路径，此处为 "H:\PROJECT\goSTU\GoRealTimeChat\Storage\Logs\logs.log"
  - **maxSize**：日志文件最大大小（单位：M），此处为 64
  - **maxBackup**：最多保存的日志文件数，0 表示不限制，此处为 30
  - **maxAge**：最多保存的日志文件天数，此处为 7
  - **compress**：是否压缩日志文件，此处为 false

- **JWT**：JWT配置项，用于身份验证和授权
  - **secret**：JWT密钥，此处为 'acb2ca5bc8d7fb2ef8f890f1be15d964'
  - **timeToLive**：JWT的过期时间，此处为 640000

- **Redis**：Redis数据库配置项
  - **host**：Redis主机地址，此处为 '127.0.0.1'
  - **port**：Redis端口号，此处为 6379
  - **password**：Redis密码，此处为空
  - **db**：Redis数据库编号，此处为 1
  - **poll**：连接池大小，默认为4倍CPU数，此处为 15
  - **conn**：最小空闲连接数，此处为 10

- **Mail**：邮件配置项，用于发送邮件
  - **driver**：邮件发送驱动，此处为 'smtp'
  - **host**：邮件服务器主机地址，此处为 'smtp.qq.com'
  - **name**：发件人邮箱地址，此处为 '2557780575@qq.com'
  - **password**：发件人邮箱密码，此处为 '********'
  - **port**：邮件服务器端口号，此处为 465
  - **encryption**：邮件服务器加密方式，此处为 'ssl'
  - **fromName**：发件人名称，此处为 'Im-Services'
  - **emailCodeSubject**：邮箱验证码邮件的主题，此处为 "欢迎使用～GoChat,这是一封邮箱验证码的邮件!"
  - **emailCodeHtmlTemplateFilePath**：邮箱验证码邮件的HTML模板文件路径

- **Nsq**：Nsq配置项，用于配置Nsq消息队列
  - `lookupHost`：Nsq的Lookup Host地址和端口号，此处为 `'127.0.0.1:4161'`
  - `nsqHost`：Nsq的主机地址和端口号，此处为 `'127.0.0.1:4150'`

- **GoBot**：GoBot配置项，用于配置机器人账号信息
  - `email`：机器人的邮箱地址，此处为 `'2557780575@qq.com'`
  - `password`：机器人的密码，此处为 `'123456'`
  - `name`：机器人的名称，此处为 `'机器人'`
  - `avatar`：机器人的头像URL，此处为 `'https://api.multiavatar.com/2557780575.png'`

- **QiNiu**：七牛云配置项，用于配置七牛云存储相关信息
  - `accessKey`：七牛云的Access Key，用于身份验证，此处为 `'、111111111111111111111111111'`
  - `secretKey`：七牛云的Secret Key，用于身份验证，此处为 `'11111111111111111111111111111111111111'`
  - `bucket`：七牛云的存储空间名称，此处为 `'grtc'`
  - `domain`：七牛云存储空间的域名，用于访问存储的文件，此处为 `'1111111111111111111111111'`


修改各项配置项为你本身机器的各项配置，建议使用Docker进行配置

- docker 配置 mysql

    要通过命令行配置MySQL容器，你可以使用Docker命令行工具来创建和管理容器。以下是一些常用的Docker命令行配置MySQL容器的示例：

    1. 创建MySQL容器并设置密码：

    ```shell
    docker run -d --name mysql-container -e MYSQL_ROOT_PASSWORD=your_password mysql:latest
    ```

    这将创建一个名为`mysql-container`的MySQL容器，并设置root用户的密码为`your_password`。

    2. 创建MySQL容器并将数据持久化到主机：

    ```shell
    docker run -d --name mysql-container -e MYSQL_ROOT_PASSWORD=your_password -v /path/to/data:/var/lib/mysql mysql:latest
    ```

    这将创建一个名为`mysql-container`的MySQL容器，并将数据持久化到主机上的`/path/to/data`目录。你可以将`/path/to/data`替换为你希望将数据保存的主机目录。

    3. 创建MySQL容器并将端口映射到主机：

    ```shell
    docker run -d --name mysql-container -e MYSQL_ROOT_PASSWORD=your_password -p 3306:3306 mysql:latest
    ```

    这将创建一个名为`mysql-container`的MySQL容器，并将容器的3306端口映射到主机的3306端口。这样，你就可以通过主机的3306端口访问MySQL服务。

    请注意，以上示例中的`your_password`应该替换为你想要设置的实际密码。你还可以根据需要添加其他环境变量、端口映射和数据卷挂载等配置。



- docker 配置 redis

    要通过命令行配置Redis容器，你可以使用Docker命令行工具来创建和管理容器。以下是一些常用的Docker命令行配置Redis容器的示例：

    1. 创建Redis容器：

    ```shell
    docker run -d --name redis-container -p 6379:6379 redis:latest
    ```

    这将创建一个名为`redis-container`的Redis容器，并将容器的6379端口映射到主机的6379端口。

    2. 创建Redis容器并将数据持久化到主机：

    ```shell
    docker run -d --name redis-container -v /path/to/data:/data redis:latest
    ```

    这将创建一个名为`redis-container`的Redis容器，并将数据持久化到主机上的`/path/to/data`目录。你可以将`/path/to/data`替换为你希望将数据保存的主机目录。

    请注意，以上示例中使用的是Redis的最新版本镜像`redis:latest`，你也可以根据需要选择其他版本。你还可以根据需要添加其他配置，如设置密码和使用配置文件等。

- docker 配置 nsq

    要通过Docker命令行配置NSQ和NSQ的Lookup Host，请按照以下步骤进行操作：

    1. 创建NSQ容器并挂载配置文件：

    ```shell
    docker run -d --name nsq-container -v /path/to/config:/etc/nsq nsqio/nsq
    ```

    这将创建一个名为`nsq-container`的NSQ容器，并将主机上的配置文件`/path/to/config`挂载到容器内的`/etc/nsq`目录，以便使用自定义的NSQ配置。

    2. 编辑NSQ的配置文件：

    ```shell
    docker exec -it nsq-container vi /etc/nsq/nsqd.conf
    ```

    这将进入`nsq-container`容器，并使用`vi`编辑器打开NSQ的配置文件`/etc/nsq/nsqd.conf`。

    3. 在配置文件中添加NSQ的Lookup Host：

    在打开的配置文件中，找到以下行：

    ```conf
    # nsqd TCP address (e.g. localhost:4150)
    ```

    在该行的下方添加以下配置行：

    ```conf
    lookupd-tcp-address=lookup-host:4160
    ```

    将`lookup-host`替换为你要使用的Lookup Host的IP地址或域名。

    4. 保存并退出编辑器：

    按下`Esc`键，然后输入`:wq`保存并退出编辑器。

    5. 重启NSQ容器：

    ```shell
    docker restart nsq-container
    ```

    这将重启名为`nsq-container`的NSQ容器，使新的配置生效。

    请注意，上述步骤假设你已经安装了Docker，并且在主机上准备好了NSQ的配置文件。你需要将`/path/to/config`替换为你实际的配置文件路径。确保在配置文件中正确设置NSQ的Lookup Host等参数，以满足你的需求。

- 验证邮箱配置

    这个请自行百度

配置成功后创建一个mysql数据库，然后运行main.go就可以启动后端服务了。


## VUE前端配置项

### 配置请求地址
修改

    /root/STU/GOSTU/GoRealTimeChatWithVUE/grtcfrontend/src/api/request.ts

里面的 BASEURL 为你自己后端的相关接口就行。然后在主目录下运行命令 yarn install 等待安装完成后 再运行 yarn dev 即可启动调试。
