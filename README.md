# Pseudo-Nezha-Agent

[fake-nezha-agent](https://github.com/dysf888/fake-nezha-agent)的改进版，支持更多数据伪装  

改自原版[nezha-agent](https://github.com/nezhahq/agent)

## Guide

| Option                    | Parameter Example | Describe                                               |
| ------------------------- | ----------------- | ------------------------------------------------------ |
| --pseudo-platform         | debian            | 修改系统名                                             |
| --pseudo-platform-version | 11.9              | 修改系统版本                                           |
| --pseudo-arch             | arm64             | 修改系统架构                                           |
| --pseudo-virt             | lxc               | 修改系统虚拟化                                         |
| --pseudo-cpu-model        | E5-2699v4         | 修改CPU型号*                                           |
| --pseudo-cpu-core         | 22                | 修改CPU核数*                                           |
| pseudo-mem-total          | 274877906944      | 修改内存大小(单位Byte)                                 |
| --pseudo-disk-total       | 175921860444160   | 修改硬盘大小(单位Byte)                                 |
| --pseudo-swap-total       | 2147483648        | 修改交换大小(单位Byte)                                 |
| --pseudo-ip               | 1.1.1.1           | 修改上报IP(仅后台可见)                                 |
| --pseudo-loc              | AQ                | 修改上报位置(使用 ISO 3166-1 alpha-2 所定义的国家代码) |
| --pseudo-boot-time        | 946684800         | 修改启动时间(使用 UnixTimeStamp)                       |
| --pseudo-version          | 0.15.11           | 修改Agent版本                                          |

*****:支持使用逗号","分隔多个参数，需要两个选项的数量相等  

(例：----pseudo-cpu-model "E5-2699v4,E5-2699v4" --pseudo-cpu-core "22,22")


## Contributors

<!--GAMFC_DELIMITER--><a href="https://github.com/naiba" title="naiba">
  <img src="https://avatars.githubusercontent.com/u/29243953?v=4" width="50;" alt="naiba"/>
</a>
<a href="https://github.com/zhangnew" title="zhangnew">
  <img src="https://avatars.githubusercontent.com/u/9146834?v=4" width="50;" alt="zhangnew"/>
</a>
<a href="https://github.com/Erope" title="卖女孩的小火柴">
  <img src="https://avatars.githubusercontent.com/u/44471469?v=4" width="50;" alt="卖女孩的小火柴"/>
</a><!--GAMFC_DELIMITER_END-->
