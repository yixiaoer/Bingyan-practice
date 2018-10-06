# 使用文档

## 注意

* 仅支持ASCII不支持UTF-8，不支持中文输入
* 语句末可以不需要分号
* 变量名仅称支持字母和下划线，且字母区分大小写
* 可以直接输入数字或字符串，将直接输出
* 可以输入表达式进行计算

## 使用

* `Def`

  `Def <identifier> = <expression>;`

  * 用于声明

    Def x = 2;

    Def max = mul(a, b) {

      return a * b;
    };

    Def num_ = 5 * 5 / 10 + 18 - mul(2, 3);
    Def n_um = num;

* `Return`

  `Return <expression>;`

  * 用于返回值

* `If`,`Else`

  `if (<condition>) <consequence> else <alternative>`

  if (x == y) {
    return x;
  } else {
    return y;
  }

* `Func`

  函数的声明`Func <parameters> <block statement>`

  函数的调用`<expression>(<comma separated expressions>)`

* 运算符左边是数字时，左边的数字与运算符需要有间隔