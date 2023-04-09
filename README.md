# Blog

- ## 前端文件相关

  - **接口地址：/zipfile/upload**

    **功能：上传前端文件**

    方法：**POST**

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

  - **接口地址：/zipfile/show/mine**

    **功能：用户获取自己上传的前端页面列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少个前端文件，默认值为20）。

    返回值：返回一个files和total，其中files是一个file数组，每个file包含id,user_id,title,content,create_at,updated_at,res_short,res_long。total表示返回文件的总数。

  - **接口地址：/zipfile/show/others/:id**

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

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出前端文件的id（即:id部分），在Body	中，raw格式提供json包含title，content，res_long（可选），res_short（可选）表示修改后的信息。

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

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出前端文件的id（即:id部分）。在Body	中，raw格式提供json包含comment，res_long（可选），res_short（可选）表示评论的信息。

    返回值：返回创建成功信息

  - **接口地址：/comment/:id**

    **功能：更新评论**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出评论的id（即:id部分）。在Body	中，raw格式提供json包含comment，res_long（可选），res_short（可选）表示评论的信息。

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
    
  - **接口地址：/comment/pagelist/mine**

    **功能：用户查看自己的评论列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇评论，默认值为20）。

    返回值：一个comments和total，其中comments是comment的数组，comment中包含了id、user_id、file_id、content、res_long、res_short、created_at、updated_at。

  - **接口地址：/comment/pagelist/others/:id**

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

  - **接口地址：/article/label/show/:id**

    **功能：查看文章标签**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的文章的id（即:id部分）

    返回值：一个labels，labels是一个string数组，表示每一个标签

  - **接口地址：/article/label/create/:id**

    **功能：创建文章标签**

    **方法：POST**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要创建的文章的id（即:id部分），在在Body	中，raw格式提供json包含一个label。

    返回值：返回设置成功

  - **接口地址：/article/label/delete/:id**

    **功能：删除文章标签**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要删除的文章的id（即:id部分），在Body	中，在Params处提供label。

    返回值：返回删除成功

  - **接口地址：/file/label/show/:id**

    **功能：查看前端文件标签**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的前端文件的id（即:id部分）

    返回参数：一个labels，labels是一个string数组，表示每一个标签

  - **接口地址：/file/label/create/:id**

    **功能：创建前端文件标签**

    **方法：POST**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要创建的前端文件的id（即:id部分），在在Body	中，raw格式提供json包含一个label。

    返回值：返回设置成功

  - **接口地址：/file/label/delete/:id**

    **功能：删除前端文件标签**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要删除的前端文件的id（即:id部分），在Params处提供label

    返回值：返回删除成功

  - **接口地址：/postlabel/show/:id**

    **功能：查看帖子标签**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的帖子的id（即:id部分）

    返回值：一个labels，labels是一个string数组，表示每一个标签

  - **接口地址：/post/label/create/:id**

    **功能：创建帖子标签**

    **方法：POST**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要创建的帖子的id（即:id部分），在在Body	中，raw格式提供json包含一个label。

    返回值：返回设置成功

  - **接口地址：/post/label/delete/:id**

    **功能：删除帖子标签**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要删除的帖子的id（即:id部分），在在Body	中，在Params处提供label

    返回值：返回删除成功

  - **接口地址：/user/label/show/:id**

    **功能：查看用户标签**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的用户的id（即:id部分）

    返回值：一个labels，labels是一个string数组，表示每一个标签

  - **接口地址：/user/label/create**

    **功能：创建用户标签**

    **方法：POST**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在Body	中，raw格式提供json包含一个label。

    返回值：返回设置成功

  - **接口地址：/user/label/delete/:id**

    **功能：删除指定的用户的标签**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要删除的用户的id（即:id部分），在Params处提供label。

    返回值：返回删除成功

  - **接口地址：/user/label/delete**

    **功能：删除用户的标签**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在Body	中，raw格式提供json包含一个label。

    返回值：返回删除成功

- ## 点赞相关

  - **接口地址：/article/like/show/:id**

    **功能：查看文章是否点赞**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的文章的id（即:id部分）

    返回值：flag，是一个bool值，表示是否点赞

  - **接口地址：/article/like/create/:id**

    **功能：给文章点赞**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的文章的id（即:id部分）

    返回值：返回点赞成功信息

  - **接口地址：/article/like/delete/:id**

    **功能：取消文章点赞**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的文章的id（即:id部分）

    返回值：返回删除成功信息

  - **接口地址：/article/like/list/:id**

    **功能：查看文章点赞列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的文章的id（即:id部分）

    返回值：返回一个Liks和total，其中Likes表示一个string数组，其中的元素表示点赞的用户id，total表示点赞数。

  - **接口地址：/post/like/show/:id**

    **功能：查看帖子是否点赞**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的帖子的id（即:id部分）

    返回值：flag，是一个bool值，表示是否点赞

  - **接口地址：/post/like/create/:id**

    **功能：给帖子点赞**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的帖子的id（即:id部分）

    返回值：返回点赞成功信息

  - **接口地址：/post/like/delete/:id**

    **功能：取消帖子点赞**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的帖子的id（即:id部分）

    返回值：返回删除成功信息

  - **接口地址：/post/like/list/:id**

    **功能：查看帖子点赞列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的帖子的id（即:id部分）

    返回值：返回一个Liks和total，其中Likes表示一个string数组，其中的元素表示点赞的用户id，total表示点赞数。

  - **接口地址：/thread/like/show/:id**

    **功能：查看跟帖是否点赞**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的跟帖的id（即:id部分）

    返回值：flag，是一个bool值，表示是否点赞

  - **接口地址：/thread/like/create/:id**

    **功能：给跟帖点赞**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的跟帖的id（即:id部分）

    返回值：返回点赞成功信息

  - **接口地址：/thread/like/delete/:id**

    **功能：取消跟帖点赞**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的跟帖的id（即:id部分）

    返回值：返回删除成功信息

  - **接口地址：/thread/like/list/:id**

    **功能：查看跟帖点赞列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的跟帖的id（即:id部分）

    返回值：返回一个Liks和total，其中Likes表示一个string数组，其中的元素表示点赞的用户id，total表示点赞数。

  - **接口地址：/file/like/show/:id**

    **功能：查看前端文件是否点赞**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的前端文件的id（即:id部分）

    返回值：flag，是一个bool值，表示是否点赞

  - **接口地址：/file/like/create/:id**

    **功能：给前端文件点赞**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的前端文件的id（即:id部分）

    返回值：返回点赞成功信息

  - **接口地址：/file/like/delete/:id**

    **功能：取消前端文件点赞**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的前端文件的id（即:id部分）

    返回值：返回删除成功信息

  - **接口地址：/file/like/list/:id**
  
    **功能：查看前端文件点赞列表**
  
    **方法：GET**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的前端文件的id（即:id部分）
  
    返回值：返回一个Liks和total，其中Likes表示一个string数组，其中的元素表示点赞的用户id，total表示点赞数。
  
  - **接口地址：/comment/like/show/:id**
  
    **功能：查看评论是否点赞**
  
    **方法：GET**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的评论的id（即:id部分）
  
    返回值：flag，是一个bool值，表示是否点赞
  
  - **接口地址：/comment/like/create/:id**
  
    **功能：给评论点赞**
  
    **方法：PUT**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的评论的id（即:id部分）
  
    返回值：返回点赞成功信息
  
  - **接口地址：/comment/like/delete/:id**
  
    **功能：取消评论点赞**
  
    **方法：DELETE**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的评论的id（即:id部分）
  
    返回值：返回删除成功信息
  
  - **接口地址：/comment/like/list/:id**
  
    **功能：查看评论点赞列表**
  
    **方法：GET**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的评论的id（即:id部分）
  
    返回值：返回一个Liks和total，其中Likes表示一个string数组，其中的元素表示点赞的用户id，total表示点赞数。
  
- ## 收藏相关

  - **接口地址：/article/favorite/show/:id**

    **功能：查看文章是否收藏**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的文章的id（即:id部分）

    返回值：flag，是一个bool值，表示是否收藏
    
  - **接口地址：/article/favorite/create/:id**

    **功能：收藏文章**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要收藏的文章的id（即:id部分）

    返回值：返回收藏成功信息

  - **接口地址：/article/favorite/delete/:id**

    **功能：取消收藏文章**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要取消收藏的文章的id（即:id部分）

    返回值：返回取消收藏成功信息

  - **接口地址：/article/favorite/list/:id**

    **功能：查看文章的收藏列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看收藏列表的文章的id（即:id部分）

    返回值：返回一个Favorites和一个total，其中Favorites为string数组，每个元素为收藏文章的用户id，total表示收藏数

  - **接口地址：/article/favorite/userlist**

    **功能：查看用户的文章收藏列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。

    返回值：返回一个Favorites和一个total，其中Favorites为string数组，每个元素为用户收藏的文章的id，total表示用户收藏的文章数

  - **接口地址：/post/favorite/show/:id**

    **功能：查看帖子是否收藏**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的帖子的id（即:id部分）

    返回值：flag，是一个bool值，表示是否收藏

  - **接口地址：/post/favorite/create/:id**

    **功能：收藏帖子**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要收藏的帖子的id（即:id部分）

    返回值：返回收藏成功信息

  - **接口地址：/post/favorite/delete/:id**

    **功能：取消收藏帖子**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要取消收藏的帖子的id（即:id部分）

    返回值：返回取消收藏成功信息

  - **接口地址：/post/favorite/list/:id**

    **功能：查看帖子的收藏列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看收藏列表的帖子的id（即:id部分）

    返回值：返回一个Favorites和一个total，其中Favorites为string数组，每个元素为收藏帖子的用户id，total表示收藏数

  - **接口地址：/postfavorite/userlist**

    **功能：查看用户的帖子收藏列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。

    返回值：返回一个Favorites和一个total，其中Favorites为string数组，每个元素为用户收藏的帖子的id，total表示用户收藏的帖子数

  - **接口地址：/thread/favorite/show/:id**

    **功能：查看跟帖是否收藏**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的跟帖的id（即:id部分）

    返回值：flag，是一个bool值，表示是否收藏

  - **接口地址：/thread/favorite/create/:id**

    **功能：收藏跟帖**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要收藏的跟帖的id（即:id部分）

    返回值：返回收藏成功信息

  - **接口地址：/thread/favorite/delete/:id**

    **功能：取消收藏跟帖**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要取消收藏的跟帖的id（即:id部分）

    返回值：返回取消收藏成功信息

  - **接口地址：/thread/favorite/list/:id**

    **功能：查看跟帖的收藏列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看收藏列表的跟帖的id（即:id部分）

    返回值：返回一个Favorites和一个total，其中Favorites为string数组，每个元素为收藏跟帖的用户id，total表示收藏数

  - **接口地址：/thread/favorite/userlist**

    **功能：查看用户的跟帖收藏列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。

    返回值：返回一个Favorites和一个total，其中Favorites为string数组，每个元素为用户收藏的跟帖的id，total表示用户收藏的跟帖数

  - **接口地址：/zipfile/favorite/show/:id**

    **功能：查看前端文件是否收藏**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看的前端文件的id（即:id部分）

    返回值：flag，是一个bool值，表示是否收藏

  - **接口地址：/zipfile/favorite/create/:id**

    **功能：收藏前端文件**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要收藏的前端文件的id（即:id部分）

    返回值：返回收藏成功信息

  - **接口地址：/zipfile/favorite/delete/:id**

    **功能：取消收藏前端文件**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要取消收藏的前端文件的id（即:id部分）

    返回值：返回取消收藏成功信息

  - **接口地址：/zipfile/favorite/list/:id**

    **功能：查看前端文件的收藏列表**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口地址中给出需要查看收藏列表的前端文件的id（即:id部分）

    返回值：返回一个Favorites和一个total，其中Favorites为string数组，每个元素为收藏跟帖的用户id，total表示收藏数

  - **接口地址：/zipfile/favorite/userlist**

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

  - **接口地址：/friend/refused/:id**

    **功能：拒绝好友申请**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出发送申请的用户的id（即:id部分）。

    返回值：返回拒绝成功信息

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
  
  - **接口地址：/visible/post/thread/:id**
  
    **功能：设置帖子是否可以跟帖**
  
    **方法：PUT**
  
    接受参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要设置跟帖等级的帖子id（即:id部分），Body部分，form-data类型，接收一个整形Level，要求不能大于4。
  
    返回值：返回设置成功信息
  
  - **接口地址：/visible/zipfile/comment/:id**
  
    **功能：设置前端文件是否可以评论**
  
    **方法：PUT**
  
    接受参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要设置评论等级的前端文件id（即:id部分），Body部分，form-data类型，接收一个整形Level，要求不能大于4。
  
    返回值：返回设置成功信息
  
  - **接口地址：/visible/post/thread/:id**
  
    **功能：查看帖子是否可以跟帖**
  
    **方法：GET**
  
    接受参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要设置可见等级的帖子id（即:id部分）
  
    返回值：返回一个Level值，表示帖子的可跟帖等级
  
  - **接口地址：/visible/zipfile/comment/:id**
  
    **功能：查看前端文件是否可以评论**
  
    **方法：GET**
  
    接受参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看评论等级的前端文件id（即:id部分）
  
    返回值：返回一个Level值，表示前端文件的可评论等级
  
  - **接口地址：/visible/zipfile/download/:id**
  
    **功能：设置前端文件下载等级**
  
    **方法：PUT**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要设置前端文件下载等级的前端文件id（即:id部分），Body部分，form-data类型，接收一个整形Level，要求不能大于4。
  
    返回值：返回设置成功信息
  
  - **接口地址：/visible/post/thread/can/:id**
  
    **功能：查看帖子是否可以跟帖**
  
    **方法：GET**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看的帖子的id（即:id部分）
  
    返回参数：返回一个flag，其为bool类型，表示当前用户是否可以跟帖
  
  - **接口地址：/visible/zipfile/can/:id**
  
    **功能：查看前端文件是否可以下载**
  
    **方法：GET**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看的前端文件id（即:id部分）
  
    返回参数：返回一个flag，其为bool类型，表示当前用户是否可以下载
  
  - **接口地址：/visible/zipfile/comment/can/:id**
  
    **功能：查看前端文件是否可以评论**
  
    **方法：GET**
  
    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看的前端文件id（即:id部分）
  
    返回参数：返回一个flag，其为bool类型，表示当前用户是否可以评论
  
- ## 热点信息相关

  - **接口地址：/hot/article/visit/:id**

    **功能：查看文章的游览次数**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看的文章id（即:id部分）

    返回参数：返回一个views，其为int类型，表示游览次数

  - **接口地址：/hot/post/visit/:id**

    **功能：查看帖子的游览次数**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看的帖子id（即:id部分）

    返回参数：返回一个views，其为int类型，表示游览次数

  - **接口地址：/hot/zipfile/visit/:id**

    **功能：查看前端文件的游览次数**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看的前端文件的id（即:id部分）

    返回参数：返回一个views，其为int类型，表示游览次数

  - **接口地址：/hot/zipfile/download/:id**

    **功能：查看前端文件的下载次数**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看的前端文件id（即:id部分）

    返回参数：返回一个downloads，其为int类型，表示下载次数

  - **接口地址：/hot/zipfile/use/:id**

    **功能：查看前端文件的使用人次**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看的前端文件id（即:id部分）

    返回参数：返回一个uses，其为int类型，表示使用人次

  - **接口地址：/hot/zipfile/comment/:id**

    **功能：查看前端文件的评论人次**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看的前端文件id（即:id部分）

    返回参数：返回一个comments，其为int类型，表示评论人次

  - **接口地址：/hot/post/thread/:id**

    **功能：查看帖子的跟帖人次**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址中给出需要查看的帖子id（即:id部分）

    返回参数：返回一个threads，其为int类型，表示跟帖人次

  - **接口地址：/hot/article**

    **功能：查看文章的热度排行**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇文章，默认值为20）。

    返回参数：返回articles和total，其中articles为string数组，每个元素表示一个article的id。注意，这个article的id不一定有查看权限。total表示文章的总数。

  - **接口地址：/hot/post**

    **功能：查看帖子的热度排行**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇帖子，默认值为20）。

    返回参数：返回posts和total，其中posts为string数组，每个元素表示一个post的id。注意，这个post的id不一定有查看权限。total表示帖子的总数。

  - **接口地址：/hot/zipfile**

    **功能：查看前端文件的热度排行**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇前端文件，默认值为20）。

    返回参数：返回zipfiles和total，其中zipfile为string数组，每个元素表示一个zipfile的id。注意，这个zipfile的id不一定有查看权限。total表示前端文件的总数。

  - **接口地址：/hot/user**

    **功能：查看用户的热度排行**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少个用户，默认值为20）。

    返回参数：返回users和total，其中users为string数组，每个元素表示一个user的id。total表示用户的总数。

  - **接口地址：/hot/user/level/:id**

    **功能：查看用户的热度等级**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，给在接口地址处给出需要查看的用户id

    返回参数：返回一个Level，为一个整数，表示用户的热度等级

  - **接口地址：/hot/user/level**

    **功能：用户查看自己的热度等级**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token

    返回参数：返回一个Level，为一个整数，表示用户的热度等级

  - **接口地址：/hot/user/powerpoint**

    **功能：用户查看自己的简报**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Params处提供start（起始时间）和end（终止时间）。

    返回参数：返回一个powerpoints，为powerpoint数组，每个powerpoint包含id,user_id,label,score,created_at，其中id表示powerpoint的id，user_id表示这条简报信息的所属者，label表示标签，score表示该标签的分数，created_at表示该消息的创建时间。

  - **接口地址：/hot/user/powerpoint/:id**

    **功能：用户查看指定用户的简报（需要权限等级3及以上）**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Params处提供start（起始时间）和end（终止时间），在接口处给出需要查看的用户的id。

    返回参数：返回一个powerpoints，为powerpoint数组，每个powerpoint包含id,user_id,label,score,created_at，其中id表示powerpoint的id，user_id表示这条简报信息的所属者，label表示标签，score表示该标签的分数，created_at表示该消息的创建时间。

  - **接口地址：/hot/article/recomment**

    **功能：用户查看文章推荐**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇文章，默认值为20）。

    返回参数：返回articles和total，其中articles为string数组，每个元素表示一个article的id。total表示文章的总数。

  - **接口地址：/hot/post/recomment**

    **功能：用户查看帖子推荐**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇帖子，默认值为20）。

    返回参数：返回posts和total，其中posts为string数组，每个元素表示一个post的id。total表示帖子的总数。

  - **接口地址：/hot/zipfile/recomment**

    **功能：用户查看前端文件推荐**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少个前端文件，默认值为20）。

    返回参数：返回zipfiles和total，其中zipfiles为string数组，每个元素表示一个zipfiles的id。total表示前端文件的总数。

  - **接口地址：/hot/user/recomment**

    **功能：用户查看用户推荐**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少个用户，默认值为20）。

    返回参数：返回users和total，其中users为string数组，每个元素表示一个user的id。total表示用户的总数。

- ## 游览历史记录相关

  - **接口地址：/history/article/:id**

    **功能：设置文章历史记录**

    **方法：POST**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址处给出文章的id。

    返回参数：返回创建成功消息。

  - **接口地址：/history/post/:id**

    **功能：设置帖子历史记录**

    **方法：POST**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址处给出帖子的id。

    返回参数：返回创建成功消息。

  - **接口地址：/history/zipfile/:id**

    **功能：设置前端文件历史记录**

    **方法：POST**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口地址处给出前端文件的id。

    返回参数：返回创建成功消息。

  - **接口地址：/history/article**

    **功能：查看文章历史记录**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在Params处提供start（起始时间）和end（终止时间）。

    返回参数：返回articleHistorys，其为articleHistory数组，每个articleHistory包含id,user_id,article_id,created_id，其中user_id表示游览文章历史所属者作者id，article_id表示所游览文章的id，created_id表示游览的时间戳。

  - **接口地址：/history/post**

    **功能：查看帖子历史记录**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在Params处提供start（起始时间）和end（终止时间）。

    返回参数：返回postHistorys，其为postHistory数组，每个postHistory包含id,user_id,post_id,created_id，其中user_id表示游览帖子历史所属者作者id，post_id表示所游览帖子的id，created_id表示游览的时间戳。

  - **接口地址：/history/zipfile**

    **功能：查看前端文件历史记录**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在Params处提供start（起始时间）和end（终止时间）。

    返回参数：返回zipfileHistorys，其为zipfileHistory数组，每个zipfileHistory包含id,user_id,zipfile_id,created_id，其中user_id表示游览前端文件历史所属者作者id，zipfile_id表示所游览前端文件的id，created_id表示游览的时间戳。

  - **接口地址：/history/article/:id**

    **功能：删除文章历史记录**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口处提供需要删除的文章历史游览记录id。

    返回参数：返回删除成功消息。

  - **接口地址：/history/post/:id**

    **功能：删除帖子历史记录**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口处提供需要删除的帖子历史游览记录id。

    返回参数：返回删除成功消息。

  - **接口地址：/history/zipfile/:id**

    **功能：删除前端文件历史记录**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口处提供需要删除的前端文件历史游览记录id。

    返回参数：返回删除成功消息。

  - **接口地址：/history/article/all**

    **功能：清空文章历史记录**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。

    返回参数：返回删除成功消息。

  - **接口地址：/history/post/all**

    **功能：清空帖子历史记录**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。

    返回参数：返回删除成功消息。

  - **接口地址：/history/zipfile/all**

    **功能：清空前端文件历史记录**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。

    返回参数：返回删除成功消息。

- ## 搜索相关

  - **接口地址：/search/article/:text**

    **功能：按文本搜索文章**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇文章，默认值为20）。

    返回参数：返回articles和total，其中articles为article数组，每个article中包含id,user_id,content,create_at,updated_at,res_short,res_long。total表示文章搜索结果的总数。

  - **接口地址：/search/post/:text**

    **功能：按文本搜索帖子**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇帖子，默认值为20）。

    返回参数：返回posts和total，其中posts为post数组，每个post中包含id,user_id,content,create_at,updated_at,res_short,res_long。total表示帖子搜索结果的总数。

  - **接口地址：/search/zipfile/:text**

    **功能：按文本搜索前端文件**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少个前端文件，默认值为20）。

    返回参数：返回zipfiles和total，其中zipfiles为zipfile数组，每个zipfile中包含id,user_id,title,content,create_at,updated_at,res_short,res_long。total表示前端文件搜索结果的总数。

  - **接口地址：/search/article/inter/:text**

    **功能：按文本和标签交集搜索文章**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇文章，默认值为20），labels数组，labels表示搜索包含的标签。

    返回参数：返回articles和total，其中articles为article数组，每个article中包含id,user_id,content,create_at,updated_at,res_short,res_long。total表示文章搜索结果的总数。

  - **接口地址：/search/post/inter/:text**

    **功能：按文本和标签交集搜索帖子**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇帖子，默认值为20），labels数组，labels表示搜索包含的标签。

    返回参数：返回posts和total，其中posts为post数组，每个post中包含id,user_id,content,create_at,updated_at,res_short,res_long。total表示帖子搜索结果的总数。

  - **接口地址：/search/zipfile/inter/:text**

    **功能：按文本和标签交集搜索前端文件**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少个前端文件，默认值为20），labels数组，labels表示搜索包含的标签。

    返回参数：返回zipfiles和total，其中zipfiles为zipfile数组，每个zipfile中包含id,user_id,title,content,create_at,updated_at,res_short,res_long。total表示前端文件搜索结果的总数。

  - **接口地址：/search/article/union/:text**

    **功能：按文本和标签并集搜索文章**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇文章，默认值为20），labels数组，labels表示搜索包含的标签。

    返回参数：返回articles和total，其中articles为article数组，每个article中包含id,user_id,content,create_at,updated_at,res_short,res_long。total表示文章搜索结果的总数。

  - **接口地址：/search/post/union/:text**

    **功能：按文本和标签并集搜索帖子**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇帖子，默认值为20），labels数组，labels表示搜索包含的标签。

    返回参数：返回posts和total，其中posts为post数组，每个post中包含id,user_id,content,create_at,updated_at,res_short,res_long。total表示帖子搜索结果的总数。

  - **接口地址：/search/zipfile/union/:text**

    **功能：按文本和标签并集搜索前端文件**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少个前端文件，默认值为20），labels数组，labels表示搜索包含的标签。

    返回参数：返回zipfiles和total，其中zipfiles为zipfile数组，每个zipfile中包含id,user_id,title,content,create_at,updated_at,res_short,res_long。total表示前端文件搜索结果的总数。
    
  - **接口地址：/search/article/:text/:id**

    **功能：按文本搜索指定用户的文章**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text和用户id，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇文章，默认值为20）。

    返回参数：返回articles和total，其中articles为article数组，每个article中包含id,user_id,content,create_at,updated_at,res_short,res_long。total表示文章搜索结果的总数。

  - **接口地址：/search/post/:text/:id**

    **功能：按文本搜索指定用户的帖子**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text和用户id，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇帖子，默认值为20）。

    返回参数：返回posts和total，其中posts为post数组，每个post中包含id,user_id,content,create_at,updated_at,res_short,res_long。total表示帖子搜索结果的总数。

  - **接口地址：/search/zipfile/:text/:id**

    **功能：按文本搜索指定用户的帖子**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text和用户id，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少个前端文件，默认值为20）。

    返回参数：返回zipfiles和total，其中zipfiles为zipfile数组，每个zipfile中包含id,user_id,title,content,create_at,updated_at,res_short,res_long。total表示前端文件搜索结果的总数。

  - **接口地址：/search/article/inter/:text/:id**

    **功能：按文本和标签交集搜索指定用户的文章**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text和用户id，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇文章，默认值为20），labels数组，labels表示搜索包含的标签。

    返回参数：返回articles和total，其中articles为article数组，每个article中包含id,user_id,content,create_at,updated_at,res_short,res_long。total表示文章搜索结果的总数。

  - **接口地址：/search/post/inter/:text/:id**

    **功能：按文本和标签交集搜索指定用户的帖子**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text和用户id，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇帖子，默认值为20），labels数组，labels表示搜索包含的标签。

    返回参数：返回posts和total，其中posts为post数组，每个post中包含id,user_id,content,create_at,updated_at,res_short,res_long。total表示帖子搜索结果的总数。

  - **接口地址：/search/zipfile/inter/:text/:id**

    **功能：按文本和标签交集搜索指定用户的前端文件**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text和用户id，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少个前端文件，默认值为20），labels数组，labels表示搜索包含的标签。

    返回参数：返回zipfiles和total，其中zipfiles为zipfile数组，每个zipfile中包含id,user_id,title,content,create_at,updated_at,res_short,res_long。total表示前端文件搜索结果的总数。

  - **接口地址：/search/article/union/:text/:id**

    **功能：按文本和标签并集搜索指定用户的文章**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text和用户id，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇文章，默认值为20），labels数组，labels表示搜索包含的标签。

    返回参数：返回articles和total，其中articles为article数组，每个article中包含id,user_id,content,create_at,updated_at,res_short,res_long。total表示文章搜索结果的总数。

  - **接口地址：/search/post/union/:text/:id**

    **功能：按文本和标签并集搜索指定用户的帖子**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text和用户id，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇帖子，默认值为20），labels数组，labels表示搜索包含的标签。

    返回参数：返回posts和total，其中posts为post数组，每个post中包含id,user_id,content,create_at,updated_at,res_short,res_long。total表示帖子搜索结果的总数。

  - **接口地址：/search/zipfile/union/:text/:id**

    **功能：按文本和标签并集搜索指定用户的前端文件**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在接口处给出需要搜索的文本信息text和用户id，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少个前端文件，默认值为20），labels数组，labels表示搜索包含的标签。

    返回参数：返回zipfiles和total，其中zipfiles为zipfile数组，每个zipfile中包含id,user_id,title,content,create_at,updated_at,res_short,res_long。total表示前端文件搜索结果的总数。

- ## 留言板相关

  - **接口地址：/guestbook**

    **功能：用户查看自己的留言板**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇留言，默认值为20）

    返回值：返回guestbooks和total，其中guestbooks为guestbook数组，每个guestbook中包含id,user_id,author,content,create_at。total表示留言的总数。

  - **接口地址：/guestbook/:id**

    **功能：查看指定用户的留言板**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Params处提供pageNum（表示第几页，默认值为1）和pageSize（表示一页多少篇留言，默认值为20），在接口地址处提供需要查找的用户的id

    返回值：返回guestbooks和total，其中guestbooks为guestbook数组，每个guestbook中包含id,user_id,author,content,create_at。total表示留言的总数。

  - **接口地址：/guestbook/:id**

    **功能：给某个用户留言**

    **方法：POST**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Body，raw格式给出json类型数据包含content。在接口地址处提供需要查找的用户的id。

    返回值：返回留言成功信息

  - **接口地址：/guestbook/:id**

    **功能：更新留言**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Body，raw格式给出json类型数据包含content。在接口处提供需要更新的留言的id

    返回值：返回留言更新成功信息

  - **接口地址：/guestbook/:id**

    **功能：更新留言**

    **方法：PUT**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Body，raw格式给出json类型数据包含content。在接口处提供需要更新的留言的id

    返回值：返回留言更新成功信息

  - **接口地址：/guestbook/:id**

    **功能：删除留言**

    **方法：DELETE**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token，在Body，raw格式给出json类型数据包含content。在接口处提供需要删除的留言的id

    返回值：返回留言删除成功信息

- ## 查找用户相关

  - **接口地址：/user/name/:id**

    **功能：通过用户名称搜索用户**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口处提供需要查询的用户名称，即/:id处

    返回值：返回一个user，包含被搜索用户的信息

  - **接口地址：/user/email/:id**

    **功能：通过用户的邮箱搜索用户**

    **方法：GET**

    接收参数：Authorization中的Bearer Token中提供注册、登录时给出的token。在接口处提供需要查询的用户邮箱，即/:id处

    返回值：返回一个user，包含被搜索用户的信息

