## gochat
implement a local chat server with Vue+Gin+gorm based on open-sorce IM app

### API Introduce

| 接口                            | 功能       | 读/写 | 特点         | 备注               |
| ------------------------------- | ---------- | ----- | ------------ | ------------------ |
| /index                          | 登录页面     | 读    |              |                    |
| /toRegister          		  | 注册页面     | 写    |              |                    |
| /toChat             		  | 进入聊天     | 读    | 根本功能接口  |                    |
| /searchFriends                  | 用户好友列表 |读      |              |                    |
| /searchGroups                   | 用户群聊列表 | 读    |             |            	   |
| /user/getUserList               | 用户列表     | 读    |              |            	   |
| /user/createUser                | 创建用户     | 写    |             |            	   |
| /user/deleteUser          	  | 删除用户     | 写    |              |            	   |
| /user/findUserByNameAndPwd      | 查找用户     | 写    |              | 根据用户名和密码查  |
| /user/updateUser                | 更新用户信息 | 写    |              |            	   |
| /user/sendMsg                   | 发送消息     | 写    |              |            	   |
| /user/sendUserMsg               | 发送用户消息  | 写    |              |            	   |

 
TODO
1. support sound and video communication
2. easy deployment
3. add chatbot user with Grok/LLaMA
