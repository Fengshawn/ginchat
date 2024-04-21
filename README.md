## gochat
implement a local chat server with Vue+Gin+gorm based on open-sorce IM app

### API Introduce

| 接口                            | 功能       | 读/写 | 特点         | 备注               |
| ------------------------------- | ---------- | ----- | ------------ | ------------------ |
| /index                          | 登录页面   | 读     | 根本功能接口  |                    |
| /douyin/user/register/          | 注册接口   | 写    |              |                    |
| /douyin/user/login/             | 登陆接口   | 读    |              |                    |
| /douyin/user/                   | 用户信息   | 读    |              |                    |
| /douyin/publish/action/         | 发布视频   | 写    |              | 存储优化           |
| /douyin/publish/list/           | 发布列表   | 读    |              | 缓存优化           |
| /douyin/favorite/action/        | 点赞接口   | 写    | 操作数大     | 延迟处理           |
| /douyin/favorite/list/          | 点赞列表   | 读    |              | 缓存优化           |
| /douyin/comment/action/         | 评论接口   | 写    | 操作数大     | 延迟处理           |
| /douyin/comment/list/           | 评论列表   | 读    |              | 缓存优化           |
| /douyin/relation/action/        | 关系操作   | 写    | 操作数大     | 延迟处理           |
| /douyin/relation/follow/list/   | 关注列表   | 读    |              | 缓存优化           |
| /douyin/relation/follower/list/ | 粉丝列表   | 读    |              | 缓存优化           |
| /douyin/relation/friend/list/   | 朋友列表   | 读    |              | 缓存优化           |
| /douyin/message/chat/           | 聊天记录   | 读    |              | 存储优化，缓存优化 |
| /douyin/message/action/         | 消息操作   | 写    |              | 延迟处理           |

r.GET("/", service.GetIndex)
	r.GET("/index", service.GetIndex)
	r.GET("/toRegister", service.ToRegister)
	r.GET("/toChat", service.ToChat)
	r.POST("/searchFriends", service.SearchFriends)
	r.POST("/searchGroups", service.SearchGroups)

	r.GET("/user/getUserList", service.GetUserList)
	r.POST("/user/createUser", service.CreateUser)
	r.GET("/user/deleteUser", service.DeleteUser)
	r.POST("/user/findUserByNameAndPwd", service.FindUserByNameAndPwd)
	r.POST("/user/updateUser", service.UpdateUser)

	//发送消息
	r.GET("/user/sendMsg", service.SendMsg)
	r.GET("/user/sendUserMsg", service.SendUserMsg)
 
TODO
1. support sound and video communication
2. easy deployment
3. add chatbot user with Grok/LLaMA
