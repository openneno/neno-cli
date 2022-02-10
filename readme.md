## neno-terminal

在终端中输入neno笔记

### 使用方式
#### 1. 设置环境变量

`githubToken`变量，并将其值设置为github的token

`githubRepo`变量，并将其值设置为保存笔记的github的仓库名称

`githubUsername`变量，并将其值设置为github的用户名
#### 2. 在终端中输入

快速记录一条笔记

```bash
neno "记录一条笔记"
```

`neno`无参数打开进入交互模式

```bash
neno> add "记录一条笔记"
```

使用vim记录多行笔记
```bash
neno> edit
```
退出交互模式
```bash
neno> exit
```



