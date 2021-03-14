# malltmp

> zeromall temp micro service

## Product Service

开发步骤：

- 根据已有的 `sql` -> `model`

- 再开始定义 `.api file` 中的 `PortalProductDetail` handler

问题：

1. 如果先生成 `model` ，但是同时 `.api` 中也需要相同的 `model struct` 组合成为 `resp struct` 。目前 `goctl` 还不支持这种声明转换（但是在 `model` 和 `.api type` 同时声明一样的 `struct` ，本身这样的设计就很多余。）
