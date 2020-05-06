# 简介
前几天老爸过生日，全家人都忘记了，然后老爸抑郁了。电话中很明显的能感受到忧伤......
事后，我对此事进行了反思和复盘。
然后私下里，我跟老哥合计：干脆我来开发一个微信公众号吧！让它来自动提醒

# 开发资料和文档
- [微信公众平台](https://mp.weixin.qq.com/)
- [微信官方文档](https://developers.weixin.qq.com/doc/)
- [腾讯客服](https://kf.qq.com/product/weixinmp.html#hid=hot_faq)

# 后台模块划分
| 序号 | 服务名 | 描述 |
| ------ | ------ | ------ |
| 1 | AccessTokenService | token服务 |
| 2 | API-Proxy | 功能调用, 即把微信官方SDK再包一层供应用调用 |

# 官方推荐的架构
![建议参考框架](https://github.com/wltos/project/blob/feature/weixin/assets/20200506_01.jpg?raw=true)

# 注意事项
- 个人公众号不能api的方式设置菜单，但可以后台手工编辑
- 个人公众号不能认证(2015年之后的)，但可以申请测试账号试用高级功能