# Blog

- ## 前端文件相关

  - **接口地址：/zipfile/upload**

    **功能：上传前端文件**

    **方法：POST**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在Body中给用form-data格式给出file（文件类型，必须为zip压缩文件）。

    返回值：成功时返回file的一些基本信息

  - **接口地址：/zipfile/download/:id**

    **功能：下载前端文件**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出前端文件的id（即:id部分）

    返回值：返回对应的zip文件

  - **接口地址：/zipfile/delete/:id**

    **功能：删除前端文件**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出前端文件的id（即:id部分）

    返回值：返回删除成功消息

  - **接口地址：/zipfile/showlist**

    **功能：获取前端文件列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少个前端文件，默认值为20）。

    返回值：返回一个files和total，其中files是一个file数组，每个file包含id,user_id,title,content,create_at,updated_at,res_short,res_long。total表示返回文件的总数。

  - **接口地址：/zipfile/showmine**

    **功能：用户获取自己上传的前端页面列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少个前端文件，默认值为20）。

    返回值：返回一个files和total，其中files是一个file数组，每个file包含id,user_id,title,content,create_at,updated_at,res_short,res_long。total表示返回文件的总数。

  - **接口地址：/zipfile/showothers/:id**

    **功能：获取某一用户上传的前端页面列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址部分给出指定用户的id。在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少个前端文件，默认值为20）。

    返回值：返回一个files和total，其中files是一个file数组，每个file包含id,user_id,title,content,create_at,updated_at,res_short,res_long。total表示返回文件的总数。

  - **接口地址：/zipfile/show/:id**

    **功能：查询某一前端文件消息**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出前端文件的id（即:id部分）

    返回值：返回被查询文件的信息file，file包含id,user_id,title,content,create_at,updated_at,res_short,res_long。

  - **接口地址：/zipfile/update/:id**

    **功能：更新前端文件描述信息**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出前端文件的id（即:id部分），在Body	中，raw格式提供json包含title，remark，res_long（可选），res_short（可选）表示修改后的信息。

    返回值：返回一个file，包含更新后的前端文件的描述信息

  - **接口地址：/zipfile/img/create/:id**

    **功能：上传前端文件描述图片**

    **方法：POST**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在Body中给用form-data格式给出file（文件类型，必须为图片）。在接口地址中给出前端文件的id（即:id部分）

    返回值：给出一个fileImg，包含了id,user_id,file_id,create_at,updated_at

  - **接口地址：/zipfile/img/delete/:id**

    **功能：删除前端描述图片**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出前端描述图片的id（即:id部分）

    返回值：返回删除成功信息

  - **接口地址：/zipfile/img/show/:id**

    **功能：用户查看前端描述图片**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出前端文件的id（即:id部分）

    返回值：返回一个fileImgs和total，其中fileImgs为fileImg数组，total表示图片的数量。其中fileImg中包含id、user_id、file_id、created_at、updated_at

  - **接口地址：/zipfile/choose/:id**

    **功能：用户选择前端文件作为博客主页**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出前端文件的id（即:id部分）。

    返回值：返回选择成功信息

  - **接口地址：/comment/:id**

    **功能：创建评论**

    **方法：POST**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出前端文件的id（即:id部分）。在Body	中，raw格式提供json包含remark，res_long（可选），res_short（可选）表示评论的信息。

    返回值：返回创建成功信息

  - **接口地址：/comment/:id**

    **功能：更新评论**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出评论的id（即:id部分）。在Body	中，raw格式提供json包含remark，res_long（可选），res_short（可选）表示评论的信息。

    返回值：返回一个更新后的comment，其中包含了id、user_id、file_id、content、res_long、res_short、created_at、updated_at

  - **接口地址：/comment/:id**

    **功能：查看评论**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出评论的id（即:id部分）。

    返回值：返回一个comment，其中包含了id、user_id、file_id、content、res_long、res_short、created_at、updated_at

  - **接口地址：/comment/:id**

    **功能：删除评论**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出评论的id（即:id部分）。

    返回值：返回删除成功信息

  - **接口地址：/comment/pagelist/:id**

    **功能：查看评论的列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出评论的id（即:id部分）。在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇评论，默认值为20）。

    返回值：一个comments和total，其中comments是comment的数组，comment中包含了id、user_id、file_id、content、res_long、res_short、created_at、updated_at
    
  - **接口地址：/comment/pagelistmine**

    **功能：用户查看自己的评论列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇评论，默认值为20）。

    返回值：一个comments和total，其中comments是comment的数组，comment中包含了id、user_id、file_id、content、res_long、res_short、created_at、updated_at。

  - **接口地址：/comment/pagelistothers/:id**

    **功能：查看某一用户的评论列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出评论的id（即:id部分）。在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇评论，默认值为20）。

    返回值：一个comments和total，其中comments是comment的数组，comment中包含了id、user_id、file_id、content、res_long、res_short、created_at、updated_at

- ## 后台管理相关

  等级：

  1=>正常使用QA

  2=>可以上传前端文件

  3=>

  4=>

  5=>拥有删除他人发表的文章、帖子、跟帖、前端文件、评论、标签的权限

  - **接口地址：/system/permission/:id/:level**

    **功能：设置用户权限等级**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要设置的用户的id（即:id部分），出需要设置的等级的id（即:level部分）。注意，用户只能赋予其它用户低于自己的用户等级，用户不能更改高于或等于自己用户等级的用户。

    返回值：返回设置成功信息

  - **接口地址：/system/permission**

    **功能：查看当前用户权限等级**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。

    返回值：一个level，其值表示当前用户的权限等级

- ## 标签相关

  - **接口地址：/articlelabel/show/:id**

    **功能：查看文章标签**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的文章的id（即:id部分）

    返回值：一个labels，labels是一个string数组，表示每一个标签

  - **接口地址：/articlelabel/create/:id**

    **功能：创建文章标签**

    **方法：POST**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要创建的文章的id（即:id部分），在在Body	中，raw格式提供json包含一个label。

    返回值：返回设置成功

  - **接口地址：/articlelabel/delete/:id**

    **功能：删除文章标签**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要删除的文章的id（即:id部分），在Body	中，raw格式提供json包含一个label。

    返回值：返回删除成功

  - **接口地址：/filelabel/show/:id**

    **功能：查看前端文件标签**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的前端文件的id（即:id部分）

    返回参数：一个labels，labels是一个string数组，表示每一个标签

  - **接口地址：/filelabel/create/:id**

    **功能：创建前端文件标签**

    **方法：POST**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要创建的前端文件的id（即:id部分），在在Body	中，raw格式提供json包含一个label。

    返回值：返回设置成功

  - **接口地址：/filelabel/delete/:id**

    **功能：删除前端文件标签**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要删除的前端文件的id（即:id部分），在Body	中，raw格式提供json包含一个label。

    返回值：返回删除成功

  - **接口地址：/postlabel/show/:id**

    **功能：查看帖子标签**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的帖子的id（即:id部分）

    返回值：一个labels，labels是一个string数组，表示每一个标签

  - **接口地址：/postlabel/create/:id**

    **功能：创建帖子标签**

    **方法：POST**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要创建的帖子的id（即:id部分），在在Body	中，raw格式提供json包含一个label。

    返回值：返回设置成功

  - **接口地址：/postlabel/delete/:id**

    **功能：删除帖子标签**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要删除的帖子的id（即:id部分），在在Body	中，raw格式提供json包含一个label。

    返回值：返回删除成功

  - **接口地址：/userlabel/show/:id**

    **功能：查看用户标签**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的用户的id（即:id部分）

    返回值：一个labels，labels是一个string数组，表示每一个标签

  - **接口地址：/userlabel/create**

    **功能：创建用户标签**

    **方法：POST**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在Body	中，raw格式提供json包含一个label。

    返回值：返回设置成功

  - **接口地址：/userlabel/delete/:id**

    **功能：删除指定的用户的标签**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要删除的用户的id（即:id部分），在Body	中，raw格式提供json包含一个label。

    返回值：返回删除成功

  - **接口地址：/userlabel/delete**

    **功能：删除用户的标签**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在Body	中，raw格式提供json包含一个label。

    返回值：返回删除成功

- ## 点赞相关

  - **接口地址：/articlelike/show/:id**

    **功能：查看文章是否点赞**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的文章的id（即:id部分）

    返回值：flag，是一个bool值，表示是否点赞

  - **接口地址：/articlelike/create/:id**

    **功能：给文章点赞**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的文章的id（即:id部分）

    返回值：返回点赞成功信息

  - **接口地址：/articlelike/delete/:id**

    **功能：取消文章点赞**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的文章的id（即:id部分）

    返回值：返回删除成功信息

  - **接口地址：/articlelike/list/:id**

    **功能：查看文章点赞列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的文章的id（即:id部分）

    返回值：返回一个Liks和total，其中Likes表示一个string数组，其中的元素表示点赞的用户id，total表示点赞数。

  - **接口地址：/postlike/show/:id**

    **功能：查看帖子是否点赞**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的帖子的id（即:id部分）

    返回值：flag，是一个bool值，表示是否点赞

  - **接口地址：/postlike/create/:id**

    **功能：给帖子点赞**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的帖子的id（即:id部分）

    返回值：返回点赞成功信息

  - **接口地址：/postlike/delete/:id**

    **功能：取消帖子点赞**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的帖子的id（即:id部分）

    返回值：返回删除成功信息

  - **接口地址：/postlike/list/:id**

    **功能：查看帖子点赞列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的帖子的id（即:id部分）

    返回值：返回一个Liks和total，其中Likes表示一个string数组，其中的元素表示点赞的用户id，total表示点赞数。

  - **接口地址：/threadlike/show/:id**

    **功能：查看跟帖是否点赞**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的跟帖的id（即:id部分）

    返回值：flag，是一个bool值，表示是否点赞

  - **接口地址：/threadlike/create/:id**

    **功能：给跟帖点赞**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的跟帖的id（即:id部分）

    返回值：返回点赞成功信息

  - **接口地址：/threadlike/delete/:id**

    **功能：取消跟帖点赞**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的跟帖的id（即:id部分）

    返回值：返回删除成功信息

  - **接口地址：/threadlike/list/:id**

    **功能：查看跟帖点赞列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的跟帖的id（即:id部分）

    返回值：返回一个Liks和total，其中Likes表示一个string数组，其中的元素表示点赞的用户id，total表示点赞数。

  - **接口地址：/filelike/show/:id**

    **功能：查看前端文件是否点赞**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的前端文件的id（即:id部分）

    返回值：flag，是一个bool值，表示是否点赞

  - **接口地址：/filelike/create/:id**

    **功能：给前端文件点赞**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的前端文件的id（即:id部分）

    返回值：返回点赞成功信息

  - **接口地址：/filelike/delete/:id**

    **功能：取消前端文件点赞**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的前端文件的id（即:id部分）

    返回值：返回删除成功信息

  - **接口地址：/filelike/list/:id**
  
    **功能：查看前端文件点赞列表**
  
    **方法：GET**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的前端文件的id（即:id部分）
  
    返回值：返回一个Liks和total，其中Likes表示一个string数组，其中的元素表示点赞的用户id，total表示点赞数。
  
  - **接口地址：/commentlike/show/:id**
  
    **功能：查看评论是否点赞**
  
    **方法：GET**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的评论的id（即:id部分）
  
    返回值：flag，是一个bool值，表示是否点赞
  
  - **接口地址：/commentlike/create/:id**
  
    **功能：给评论点赞**
  
    **方法：PUT**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的评论的id（即:id部分）
  
    返回值：返回点赞成功信息
  
  - **接口地址：/commentlike/delete/:id**
  
    **功能：取消评论点赞**
  
    **方法：DELETE**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的评论的id（即:id部分）
  
    返回值：返回删除成功信息
  
  - **接口地址：/commentlike/list/:id**
  
    **功能：查看前端评论列表**
  
    **方法：GET**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的评论的id（即:id部分）
  
    返回值：返回一个Liks和total，其中Likes表示一个string数组，其中的元素表示点赞的用户id，total表示点赞数。
  
- ## 收藏相关

  - **接口地址：/articlefavorite/show/:id**

    **功能：查看文章是否收藏**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的文章的id（即:id部分）

    返回值：flag，是一个bool值，表示是否收藏
    
  - **接口地址：/articlefavorite/creat/:id**

    **功能：收藏文章**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要收藏的文章的id（即:id部分）

    返回值：返回收藏成功信息

  - **接口地址：/articlefavorite/delete/:id**

    **功能：取消收藏文章**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要取消收藏的文章的id（即:id部分）

    返回值：返回取消收藏成功信息

  - **接口地址：/articlefavorite/list/:id**

    **功能：查看文章的收藏列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看收藏列表的文章的id（即:id部分）

    返回值：返回一个Favorites和一个total，其中Favorites为string数组，每个元素为收藏文章的用户id，total表示收藏数

  - **接口地址：/articlefavorite/userlist**

    **功能：查看用户的文章收藏列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。

    返回值：返回一个Favorites和一个total，其中Favorites为string数组，每个元素为用户收藏的文章的id，total表示用户收藏的文章数

  - **接口地址：/postfavorite/show/:id**

    **功能：查看帖子是否收藏**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的帖子的id（即:id部分）

    返回值：flag，是一个bool值，表示是否收藏

  - **接口地址：/postfavorite/creat/:id**

    **功能：收藏帖子**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要收藏的帖子的id（即:id部分）

    返回值：返回收藏成功信息

  - **接口地址：/postfavorite/delete/:id**

    **功能：取消收藏帖子**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要取消收藏的帖子的id（即:id部分）

    返回值：返回取消收藏成功信息

  - **接口地址：/postfavorite/list/:id**

    **功能：查看帖子的收藏列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看收藏列表的帖子的id（即:id部分）

    返回值：返回一个Favorites和一个total，其中Favorites为string数组，每个元素为收藏帖子的用户id，total表示收藏数

  - **接口地址：/postfavorite/userlist**

    **功能：查看用户的帖子收藏列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。

    返回值：返回一个Favorites和一个total，其中Favorites为string数组，每个元素为用户收藏的帖子的id，total表示用户收藏的帖子数

  - **接口地址：/threadfavorite/show/:id**

    **功能：查看跟帖是否收藏**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的跟帖的id（即:id部分）

    返回值：flag，是一个bool值，表示是否收藏

  - **接口地址：/threadfavorite/creat/:id**

    **功能：收藏跟帖**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要收藏的跟帖的id（即:id部分）

    返回值：返回收藏成功信息

  - **接口地址：/threadfavorite/delete/:id**

    **功能：取消收藏跟帖**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要取消收藏的跟帖的id（即:id部分）

    返回值：返回取消收藏成功信息

  - **接口地址：/threadfavorite/list/:id**

    **功能：查看跟帖的收藏列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看收藏列表的跟帖的id（即:id部分）

    返回值：返回一个Favorites和一个total，其中Favorites为string数组，每个元素为收藏跟帖的用户id，total表示收藏数

  - **接口地址：/threadfavorite/userlist**

    **功能：查看用户的跟帖收藏列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。

    返回值：返回一个Favorites和一个total，其中Favorites为string数组，每个元素为用户收藏的跟帖的id，total表示用户收藏的跟帖数

  - **接口地址：/zipfilefavorite/show/:id**

    **功能：查看前端文件是否收藏**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的前端文件的id（即:id部分）

    返回值：flag，是一个bool值，表示是否收藏

  - **接口地址：/zipfilefavorite/creat/:id**

    **功能：收藏前端文件**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要收藏的前端文件的id（即:id部分）

    返回值：返回收藏成功信息

  - **接口地址：/zipfilefavorite/delete/:id**

    **功能：取消收藏前端文件**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要取消收藏的前端文件的id（即:id部分）

    返回值：返回取消收藏成功信息

  - **接口地址：/zipfilefavorite/list/:id**

    **功能：查看前端文件的收藏列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看收藏列表的前端文件的id（即:id部分）

    返回值：返回一个Favorites和一个total，其中Favorites为string数组，每个元素为收藏跟帖的用户id，total表示收藏数

  - **接口地址：/zipfilefavorite/userlist**

    **功能：查看用户的前端文件收藏列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。

    返回值：返回一个Favorites和一个total，其中Favorites为string数组，每个元素为用户收藏的前端文件的id，total表示用户收藏的前端文件数

- ## 好友相关

  - **接口地址：/friend/show**

    **功能：查看好友列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。

    返回值：返回一个friends， 其为一个string数组，每个元素表示了好友的用户id

  - **接口地址：/friend/show/applied**

    **功能：查看用户正在被申请，暂未通过的用户列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。

    返回值：返回一个friends， 其为一个string数组，每个元素表示了发送好友申请的用户id

  - **接口地址：/friend/show/applying**

    **功能：查看用户正在申请，暂未通过好友的用户列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。

    返回值：返回一个friends， 其为一个string数组，每个元素表示了被发送好友申请的用户id

  - **接口地址：/friend/applying/:id**

    **功能：发送好友申请**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出发送申请的用户的id（即:id部分）。

    返回值：返回申请成功信息

  - **接口地址：/friend/applied/:id**

    **功能：接受好友申请**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出发送申请的用户的id（即:id部分）。

    返回值：返回接收成功信息

  - **接口地址：/friend/delete/:id**

    **功能：删除好友**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要删除的用户的id（即:id部分）。

    返回值：返回删除成功信息
    
  - **接口地址：/friend/articles**

    **功能：获取好友圈文章**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇文章，默认值为20）。

    返回值：返回articles和total，其中articles是article数组，article表示一篇文章的信息。total表示总计有多少篇文章。

  - **接口地址：/friend/posts**

    **功能：获取好友圈帖子**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇帖子，默认值为20）。

    返回值：返回posts和total，其中posts是post数组，post表示一篇帖子的信息。total表示总计有多少篇帖子。

  - **接口地址：/friend/zipfiles**

    **功能：获取好友圈前端文件**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少个前端文件，默认值为20）。

    返回值：返回zipfiles和total，其中zipfiles是zipfile数组，zipfile表示一个前端文件的信息。total表示总计有多少个前端文件。

- ## 隔离相关

  注意：隔离等级Level
  
  1=>公开，所有人可以操作
  
  2=>仅好友圈内可操作
  
  3=>仅自己可见、可操作
  
  - **接口地址：/visible/article/:id**
  
    **功能：设置文章可见等级**
  
    **方法：PUT**
  
    接受参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要设置可见等级的文章id（即:id部分），Body部分，form-data类型，接收一个整形Level，要求不能大于4。
  
    返回值：返回设置成功信息
  
  - **接口地址：/visible/post/:id**
  
    **功能：设置帖子可见等级**
  
    **方法：PUT**
  
    接受参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要设置可见等级的帖子id（即:id部分），Body部分，form-data类型，接收一个整形Level，要求不能大于4。
  
    返回值：返回设置成功信息
  
  - **接口地址：/visible/zipfile/:id**
  
    **功能：设置前端文件可见等级**
  
    **方法：PUT**
  
    接受参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要设置可见等级的前端文件id（即:id部分），Body部分，form-data类型，接收一个整形Level，要求不能大于4。
  
    返回值：返回设置成功信息
  
  - **接口地址：/visible/article/:id**
  
    **功能：查看文章可见等级**
  
    **方法：GET**
  
    接受参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看可见等级的文章id（即:id部分）
  
    返回值：返回一个Level值，表示文章的可见等级
  
  - **接口地址：/visible/post/:id**
  
    **功能：查看帖子可见等级**
  
    **方法：GET**
  
    接受参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看可见等级的帖子id（即:id部分）
  
    返回值：返回一个Level值，表示帖子的可见等级
  
  - **接口地址：/visible/zipfile/:id**
  
    **功能：查看前端文件可见等级**
  
    **方法：GET**
  
    接受参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看可见等级的前端文件id（即:id部分）
  
    返回值：返回一个Level值，表示前端文件的可见等级
  
  - **接口地址：/visible/postthread/:id**
  
    **功能：设置帖子是否可以跟帖**
  
    **方法：PUT**
  
    接受参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要设置跟帖等级的帖子id（即:id部分），Body部分，form-data类型，接收一个整形Level，要求不能大于4。
  
    返回值：返回设置成功信息
  
  - **接口地址：/visible/zipfilecomment/:id**
  
    **功能：设置前端文件是否可以评论**
  
    **方法：PUT**
  
    接受参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要设置评论等级的前端文件id（即:id部分），Body部分，form-data类型，接收一个整形Level，要求不能大于4。
  
    返回值：返回设置成功信息
  
  - **接口地址：/visible/postthread/:id**
  
    **功能：查看帖子是否可以跟帖**
  
    **方法：GET**
  
    接受参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要设置可见等级的帖子id（即:id部分）
  
    返回值：返回一个Level值，表示帖子的可跟帖等级
  
  - **接口地址：/visible/zipfilecomment/:id**
  
    **功能：查看前端文件是否可以评论**
  
    **方法：GET**
  
    接受参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看评论等级的前端文件id（即:id部分）
  
    返回值：返回一个Level值，表示前端文件的可评论等级
  
  - **接口地址：/zipfiledownload/:id**
  
    **功能：设置前端文件下载等级**
  
    **方法：PUT**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要设置前端文件下载等级的前端文件id（即:id部分），Body部分，form-data类型，接收一个整形Level，要求不能大于4。
  
    返回值：返回设置成功信息
  
  - **接口地址：/postthreadcan/:id**
  
    **功能：查看帖子是否可以跟帖**
  
    **方法：GET**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看的帖子的id（即:id部分）
  
    返回参数：返回一个flag，其为bool类型，表示当前用户是否可以跟帖
  
  - **接口地址：/zipfilecan/:id**
  
    **功能：查看前端文件是否可以下载**
  
    **方法：GET**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看的前端文件id（即:id部分）
  
    返回参数：返回一个flag，其为bool类型，表示当前用户是否可以下载
  
  - **接口地址：/zipfilecommentcan/:id**
  
    **功能：查看前端文件是否可以评论**
  
    **方法：GET**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看的前端文件id（即:id部分）
  
    返回参数：返回一个flag，其为bool类型，表示当前用户是否可以评论
  
