# Pseudo-Nezha-Agent

[fake-nezha-agent](https://github.com/dysf888/fake-nezha-agent)的改进版，支持更多数据伪装  

改自原版[nezha-agent](https://github.com/nezhahq/agent)

## Guide

| Option                    | Parameter Example       | Data Type | Describe                                               |
| :------------------------ | ----------------------- | --------- | ------------------------------------------------------ |
| --pseudo-platform         | debian                  | string    | 修改系统名                                             |
| --pseudo-platform-version | 11.9                    | string    | 修改系统版本                                           |
| --pseudo-arch             | arm64                   | string    | 修改系统架构                                           |
| --pseudo-virt             | lxc                     | string    | 修改系统虚拟化                                         |
| --pseudo-cpu-model        | E5-2699v4               | string    | 修改CPU型号*                                           |
| --pseudo-cpu-core         | 22                      | string    | 修改CPU核数*                                           |
| --pseudo-mem-total        | 274877906944            | int       | 修改内存大小(单位Byte)                                 |
| --pseudo-disk-total       | 175921860444160         | int       | Int修改硬盘大小(单位Byte)                              |
| --pseudo-swap-total       | 2147483648              | int       | Int修改交换大小(单位Byte)                              |
| --pseudo-ip               | 1.1.1.1/2606:4700::1111 | string    | 修改上报IP(仅后台可见)                                 |
| --pseudo-loc              | KP                      | string    | 修改上报位置(使用 ISO 3166-1 alpha-2 所定义的国家代码) |
| --pseudo-boot-time        | 946684800               | int       | 修改启动时间(使用 UnixTimeStamp)                       |
| --pseudo-version          | 0.15.11                 | string    | 修改Agent版本                                          |
| --pseudo-cpu-used         | 20                      | string    | 修改CPU占用率**                                        |
| --pseudo-mem-used         | 50-80                   | string    | 修改内存占用率**                                       |
| --pseudo-swap-used        | 50,5                    | string    | 修改交换占用率**                                       |
| --pseudo-tcp-count        | 20000                   | int       | 修改TCP连接数                                          |
| --pseudo-udp-count        | 20000                   | int       | 修改UDP连接数                                          |
| --pseudo-netin-factor     | 10                      | int       | 下行流量倍率                                           |
| --pseudo-netout-factor    | 10                      | int       | 上行流量倍率                                           |

*:支持使用逗号","分隔多个参数，需要两个选项的数量相等  

(例：--pseudo-cpu-model "E5-2699v4,E5-2699v4" --pseudo-cpu-core "22,22")

**:拥有三种模式

固定值(0-100之间数值)

(例：--pseudo-cpu-used "20")

随机范围(最小值，最大值)

(例：--pseudo-mem-used "50-80")

正态分布(平均值，标准差)

(例：--pseudo-swap-used "50,5")


## Contributors

<!--GAMFC_DELIMITER--><a href="https://github.com/naiba" title="naiba">
  <img src="https://avatars.githubusercontent.com/u/29243953?v=4" width="50;" alt="naiba"/>
</a>
<a href="https://github.com/Redamancy2319" title="Redamancy2319">
  <img src="https://avatars.githubusercontent.com/u/101871681?v=4" width="50;" alt="Redamancy2319"/>
</a>
<a href="https://github.com/zhangnew" title="zhangnew">
  <img src="https://avatars.githubusercontent.com/u/9146834?v=4" width="50;" alt="zhangnew"/>
</a>
<a href="https://github.com/Erope" title="卖女孩的小火柴">
  <img src="https://avatars.githubusercontent.com/u/44471469?v=4" width="50;" alt="卖女孩的小火柴"/>
</a><!--GAMFC_DELIMITER_END-->
