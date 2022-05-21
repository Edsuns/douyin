# Git 使用规范

## 开发流程

参考 [GitHub Flow](https://docs.github.com/en/get-started/quickstart/github-flow) ，有以下要求：

- 禁止直接向 main 分支提交 commit ，所有的 commit 应该提交到自己的开发分支，然后提 pull request 让管理审核后合并进 main 分支
- 开发分支的命名格式为 feat/name ，比如要开发一个注册功能，那么你应该在 feat/register 分支下开发和 commit
- 如果是新开发一个功能，无需提 issues ，直接新建分支进行开发，新功能开发完后提 pull request 即可

## Commit 信息（留言）

每次提交，commit message 都包括三个部分：header，body 和 footer。

```text
<type>(<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

其中，header 是必需的，body 和 footer 允许省略。任何一行都不得超过72个字符，这是为了避免太长了部分内容被隐藏影响阅读。

header 部分只有一行，包括三个字段：type（必需）、scope（可选）和 subject（必需）。subject 不要求句子首字母大写，内容简单扼要地陈述此次修改的内容即可。

type 用于说明 commit 的类别，只允许使用下面7个标识，全都是英文小写。

- feat：新功能 (feature)
- fix：修补 bug
- docs：文档 (documentation)
- style： 格式（不影响代码运行的变动）
- refactor：重构（即不是新增功能，也不是修复 bug 的代码变动）
- test：测试相关的改动
- chore：构建过程或辅助工具的变动

如果要写 body 或 footer 切记要空一行后再写。 
