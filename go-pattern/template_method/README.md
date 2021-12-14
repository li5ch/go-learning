# 模版方法模式

>抽象类里定义好算法的执行步骤和具体算法，以及可能发生变化的算法定义为抽象方法。不同的子类继承该抽象类，并实现父类的抽象方法。

模版方法模式使用继承机制，把通用步骤和通用方法放到父类中，把具体实现延迟到子类中实现。使得实现符合开闭原则。

如实例代码中通用步骤在父类中实现（`准备`、`下载`、`保存`、`收尾`）下载和保存的具体实现留到子类中，并且提供 `保存`方法的默认实现。

因为Golang不提供继承机制，需要使用匿名组合模拟实现继承。

此处需要注意：因为父类需要调用子类方法，所以子类需要匿名组合父类的同时，父类需要持有子类的引用。

- 不变：Run方法里的抽奖步骤 -> 被继承复用
- 变：不同场景下 -> 被具体实现
