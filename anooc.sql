/*
Navicat MySQL Data Transfer

Source Server         : qa
Source Server Version : 50713
Source Host           : 139.196.16.67:3306
Source Database       : anooc

Target Server Type    : MYSQL
Target Server Version : 50713
File Encoding         : 65001

Date: 2016-10-08 15:37:18
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for advertisement
-- ----------------------------
DROP TABLE IF EXISTS `advertisement`;
CREATE TABLE `advertisement` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '广告名称',
  `code` varchar(255) NOT NULL DEFAULT '' COMMENT '广告内容代码(html、js等)',
  `source` varchar(20) NOT NULL DEFAULT '' COMMENT '广告来源，如 baidu_union，shiyanlou',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='广告表';

-- ----------------------------
-- Table structure for anote
-- ----------------------------
DROP TABLE IF EXISTS `anote`;
CREATE TABLE `anote` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `uid` int(11) DEFAULT '0' COMMENT '用户uid',
  `title` varchar(255) DEFAULT NULL COMMENT '文章标题',
  `content` text COMMENT '文章内容',
  `viewnum` int(11) DEFAULT '0' COMMENT '浏览数目',
  `commentnum` int(11) DEFAULT '0' COMMENT '评论数',
  `newslist_tpl` tinyint(1) DEFAULT '1' COMMENT '新闻列表模板：1无图标',
  `ctime` timestamp NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `mtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `domain` varchar(50) NOT NULL DEFAULT '' COMMENT '来源域名（不一定是顶级域名）',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '来源名称',
  `title` varchar(127) NOT NULL DEFAULT '' COMMENT '文章标题',
  `cover` varchar(127) NOT NULL DEFAULT '' COMMENT '图片封面',
  `author` varchar(1024) NOT NULL DEFAULT '' COMMENT '文章作者(可能带html)',
  `author_txt` varchar(127) NOT NULL DEFAULT '' COMMENT '文章作者(纯文本)',
  `lang` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '语言：0-中文；1-英文',
  `pub_date` varchar(20) NOT NULL DEFAULT '' COMMENT '发布时间',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '文章原始链接',
  `content` longtext NOT NULL COMMENT '正文(带html)',
  `txt` text NOT NULL COMMENT '正文(纯文本)',
  `tags` varchar(50) NOT NULL DEFAULT '' COMMENT '文章tag，逗号分隔',
  `css` varchar(255) NOT NULL DEFAULT '' COMMENT '需要额外引入的css样式',
  `viewnum` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '浏览数',
  `cmtnum` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '评论数',
  `likenum` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '赞数',
  `top` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '置顶，0否，1置顶',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态：0-初始抓取；1-已上线；2-下线(审核拒绝)',
  `op_user` varchar(20) NOT NULL DEFAULT '' COMMENT '操作人',
  `ctime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `url` (`url`),
  KEY `top` (`top`),
  KEY `domain` (`domain`),
  KEY `ctime` (`ctime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='网络文章聚合表';

-- ----------------------------
-- Table structure for authority
-- ----------------------------
DROP TABLE IF EXISTS `authority`;
CREATE TABLE `authority` (
  `aid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT '权限名',
  `menu1` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '所属一级菜单，本身为一级菜单，则为0',
  `menu2` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '所属二级菜单，本身为二级菜单，则为0',
  `route` varchar(128) NOT NULL DEFAULT '' COMMENT '路由（权限）',
  `op_user` varchar(20) NOT NULL COMMENT '操作人',
  `ctime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`aid`),
  KEY `route` (`route`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8 COMMENT='权限表，常驻内存';

-- ----------------------------
-- Table structure for bind_user
-- ----------------------------
DROP TABLE IF EXISTS `bind_user`;
CREATE TABLE `bind_user` (
  `uid` int(10) unsigned NOT NULL,
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '绑定的第三方类型',
  `email` varchar(128) NOT NULL DEFAULT '',
  `tuid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '第三方uid',
  `username` varchar(20) NOT NULL COMMENT '用户名',
  `token` varchar(50) NOT NULL COMMENT '第三方access_token',
  `refresh` varchar(50) NOT NULL COMMENT '第三方refresh_token',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='第三方绑定表';

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `cid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `objid` int(10) unsigned NOT NULL COMMENT '对象id，属主（评论给谁）',
  `objtype` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '类型,0-帖子;1-博文;2-资源;3-wiki;4-project',
  `content` text NOT NULL,
  `uid` int(10) unsigned NOT NULL COMMENT '回复者',
  `floor` int(10) unsigned NOT NULL COMMENT '第几楼',
  `flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '审核标识,0-未审核;1-已审核;2-审核删除;3-用户自己删除',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`cid`),
  UNIQUE KEY `objid` (`objid`,`objtype`,`floor`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='评论表（帖子回复、博客文章评论等，统一处理）';

-- ----------------------------
-- Table structure for crawl_rule
-- ----------------------------
DROP TABLE IF EXISTS `crawl_rule`;
CREATE TABLE `crawl_rule` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `domain` varchar(50) NOT NULL DEFAULT '' COMMENT '来源域名（不一定是顶级域名）',
  `subpath` varchar(20) NOT NULL DEFAULT '' COMMENT '域名下面紧接着的path（区别同一网站多个路径不同抓取规则）',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '来源名称',
  `lang` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '主要语言：0-中文;1-英文',
  `title` varchar(127) NOT NULL DEFAULT '' COMMENT '文章标题规则',
  `in_url` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '作者信息是否在url中；0-否;1-是(是的时候，author表示在url中的位置)',
  `author` varchar(127) NOT NULL DEFAULT '' COMMENT '文章作者规则',
  `pub_date` varchar(127) NOT NULL DEFAULT '' COMMENT '发布时间规则',
  `content` varchar(127) NOT NULL DEFAULT '' COMMENT '正文规则',
  `op_user` varchar(20) NOT NULL DEFAULT '' COMMENT '操作人',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `domain` (`domain`,`subpath`),
  KEY `ctime` (`ctime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='网站抓取规则表';

-- ----------------------------
-- Table structure for dynamic
-- ----------------------------
DROP TABLE IF EXISTS `dynamic`;
CREATE TABLE `dynamic` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `content` varchar(255) NOT NULL DEFAULT '' COMMENT '动态内容',
  `dmtype` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '类型：0-Go动态；1-本站动态',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '链接',
  `seq` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '顺序（越大越在前）',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `seq` (`seq`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='动态表';

-- ----------------------------
-- Table structure for favorites
-- ----------------------------
DROP TABLE IF EXISTS `favorites`;
CREATE TABLE `favorites` (
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户uid',
  `objtype` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '类型,0-帖子;1-博文;2-资源;3-wiki',
  `objid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '对象id，属主',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`,`objtype`,`objid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户收藏';

-- ----------------------------
-- Table structure for image
-- ----------------------------
DROP TABLE IF EXISTS `image`;
CREATE TABLE `image` (
  `pid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `md5` char(32) NOT NULL DEFAULT '' COMMENT '图片md5',
  `path` varchar(127) NOT NULL DEFAULT '' COMMENT '图片路径（不包括域名）',
  `size` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '图片的大小（字节）',
  `width` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '图片宽度',
  `height` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '图片高度',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`pid`),
  UNIQUE KEY `md5` (`md5`),
  KEY `created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='图片表';

-- ----------------------------
-- Table structure for likes
-- ----------------------------
DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes` (
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '喜欢人的uid',
  `objtype` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '类型,0-帖子;1-博文;2-资源;3-wiki',
  `objid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '对象id，属主',
  `flag` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1-喜欢；2-不喜欢（暂时不支持）',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`,`objtype`,`objid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='喜欢表';

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `content` text NOT NULL COMMENT '消息内容',
  `hasread` enum('未读','已读') NOT NULL DEFAULT '未读',
  `from` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '来自谁',
  `fdel` enum('未删','已删') NOT NULL DEFAULT '未删' COMMENT '发送方删除标识',
  `to` int(10) unsigned NOT NULL COMMENT '发给谁',
  `tdel` enum('未删','已删') NOT NULL DEFAULT '未删' COMMENT '接收方删除标识',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `to` (`to`),
  KEY `from` (`from`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='message 短消息（私信）';

-- ----------------------------
-- Table structure for morning_reading
-- ----------------------------
DROP TABLE IF EXISTS `morning_reading`;
CREATE TABLE `morning_reading` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `content` varchar(255) NOT NULL DEFAULT '' COMMENT '晨读内容',
  `rtype` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '晨读类别：0-Go技术晨读;1-综合技术晨读',
  `inner` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '本站文章id，如果外站文章，则为0',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '文章链接，本站文章时为空',
  `moreurls` varchar(1024) NOT NULL DEFAULT '' COMMENT '可能顺带推荐多篇文章；url逗号分隔',
  `clicknum` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '点击数',
  `username` varchar(20) NOT NULL DEFAULT '' COMMENT '发布人',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='技术晨读表';

-- ----------------------------
-- Table structure for open_project
-- ----------------------------
DROP TABLE IF EXISTS `open_project`;
CREATE TABLE `open_project` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '项目id',
  `name` varchar(127) NOT NULL DEFAULT '' COMMENT '项目名(软件名)，如 Docker',
  `category` varchar(127) NOT NULL DEFAULT '' COMMENT '项目类别，如 Linux 容器引擎',
  `uri` varchar(127) NOT NULL DEFAULT '' COMMENT '项目uri，访问使用（如/p/docker 中的 docker)',
  `home` varchar(127) NOT NULL DEFAULT '' COMMENT '项目首页',
  `doc` varchar(127) NOT NULL DEFAULT '' COMMENT '项目文档地址',
  `download` varchar(127) NOT NULL DEFAULT '' COMMENT '项目下载地址',
  `src` varchar(127) NOT NULL DEFAULT '' COMMENT '源码地址',
  `logo` varchar(127) NOT NULL DEFAULT '' COMMENT '项目logo',
  `desc` text NOT NULL COMMENT '项目描述',
  `repo` varchar(127) NOT NULL DEFAULT '' COMMENT '源码uri部分，方便repo widget插件使用',
  `author` varchar(127) NOT NULL DEFAULT '' COMMENT '作者',
  `licence` varchar(127) NOT NULL DEFAULT '' COMMENT '授权协议',
  `lang` varchar(127) NOT NULL DEFAULT '' COMMENT '开发语言',
  `os` varchar(127) NOT NULL DEFAULT '' COMMENT '操作系统（多个逗号分隔）',
  `tags` varchar(127) NOT NULL DEFAULT '' COMMENT 'tag，逗号分隔',
  `username` varchar(127) NOT NULL DEFAULT '' COMMENT '收录人',
  `viewnum` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '浏览数',
  `cmtnum` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '评论数',
  `likenum` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '赞数',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态：0-新建；1-已上线；2-下线(审核拒绝)',
  `ctime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '加入时间',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `uri` (`uri`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='开源项目';

-- ----------------------------
-- Table structure for page_ad
-- ----------------------------
DROP TABLE IF EXISTS `page_ad`;
CREATE TABLE `page_ad` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `ad_id` varchar(20) NOT NULL DEFAULT '' COMMENT '广告名称',
  `code` varchar(255) NOT NULL DEFAULT '' COMMENT '广告内容代码(html、js等)',
  `source` varchar(20) NOT NULL DEFAULT '' COMMENT '广告来源，如 baidu_union，shiyanlou',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='页面广告管理表';

-- ----------------------------
-- Table structure for resource
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL COMMENT '资源标题',
  `form` enum('只是链接','包括内容') DEFAULT NULL,
  `content` longtext NOT NULL COMMENT '资源内容',
  `url` varchar(150) NOT NULL COMMENT '链接url',
  `uid` int(10) unsigned NOT NULL COMMENT '作者',
  `catid` int(10) unsigned NOT NULL COMMENT '所属类别',
  `ctime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `url` (`url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='资源';

-- ----------------------------
-- Table structure for resource_category
-- ----------------------------
DROP TABLE IF EXISTS `resource_category`;
CREATE TABLE `resource_category` (
  `catid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL COMMENT '分类名',
  `intro` varchar(50) NOT NULL DEFAULT '' COMMENT '分类简介',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`catid`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='资源分类表';

-- ----------------------------
-- Table structure for resource_ex
-- ----------------------------
DROP TABLE IF EXISTS `resource_ex`;
CREATE TABLE `resource_ex` (
  `id` int(10) unsigned NOT NULL,
  `viewnum` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '浏览数',
  `cmtnum` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '回复数',
  `likenum` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '喜欢数',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='资源扩展表';

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `roleid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT '角色名',
  `op_user` varchar(20) NOT NULL DEFAULT '' COMMENT '操作人',
  `ctime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`roleid`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COMMENT='角色表，常驻内存';

-- ----------------------------
-- Table structure for role_authority
-- ----------------------------
DROP TABLE IF EXISTS `role_authority`;
CREATE TABLE `role_authority` (
  `roleid` int(10) unsigned NOT NULL,
  `aid` int(10) unsigned NOT NULL,
  `op_user` varchar(20) NOT NULL DEFAULT '' COMMENT '操作人',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`roleid`,`aid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色拥有的权限表';

-- ----------------------------
-- Table structure for search_stat
-- ----------------------------
DROP TABLE IF EXISTS `search_stat`;
CREATE TABLE `search_stat` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `keyword` varchar(127) NOT NULL DEFAULT '' COMMENT '搜索词',
  `times` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '次数',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`),
  KEY `times` (`times`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='搜索词统计';

-- ----------------------------
-- Table structure for system_message
-- ----------------------------
DROP TABLE IF EXISTS `system_message`;
CREATE TABLE `system_message` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `msgtype` tinyint(4) NOT NULL DEFAULT '0' COMMENT '系统消息类型',
  `hasread` enum('未读','已读') NOT NULL DEFAULT '未读',
  `to` int(10) unsigned NOT NULL COMMENT '发给谁',
  `ext` text NOT NULL COMMENT '额外信息',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `to` (`to`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='system_message 系统消息表';

-- ----------------------------
-- Table structure for topics
-- ----------------------------
DROP TABLE IF EXISTS `topics`;
CREATE TABLE `topics` (
  `tid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `nid` int(10) unsigned NOT NULL COMMENT '节点id',
  `uid` int(10) unsigned NOT NULL COMMENT '帖子作者',
  `lastreplyuid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后回复者',
  `lastreplytime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后回复时间',
  `flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '审核标识,0-未审核;1-已审核;2-审核删除;3-用户自己删除',
  `editor_uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后编辑人',
  `top` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '置顶，0否，1置顶',
  `ctime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`tid`),
  KEY `uid` (`uid`),
  KEY `nid` (`nid`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='帖子内容表';

-- ----------------------------
-- Table structure for topics_ex
-- ----------------------------
DROP TABLE IF EXISTS `topics_ex`;
CREATE TABLE `topics_ex` (
  `tid` int(10) unsigned NOT NULL,
  `view` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '浏览数',
  `reply` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '回复数',
  `like` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '喜欢数',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`tid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='帖子扩展表（计数）';

-- ----------------------------
-- Table structure for topics_node
-- ----------------------------
DROP TABLE IF EXISTS `topics_node`;
CREATE TABLE `topics_node` (
  `nid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父节点id，无父节点为0',
  `name` varchar(20) NOT NULL COMMENT '节点名',
  `intro` varchar(50) NOT NULL DEFAULT '' COMMENT '节点简介',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`nid`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8 COMMENT='帖子节点表';

-- ----------------------------
-- Table structure for user_active
-- ----------------------------
DROP TABLE IF EXISTS `user_active`;
CREATE TABLE `user_active` (
  `uid` int(10) unsigned NOT NULL,
  `email` varchar(128) NOT NULL,
  `username` varchar(20) NOT NULL COMMENT '用户名',
  `weight` smallint(6) NOT NULL DEFAULT '1' COMMENT '活跃度，越大越活跃',
  `avatar` varchar(128) NOT NULL DEFAULT '' COMMENT '头像(如果为空，则使用http://www.gravatar.com)',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='活跃用户表';

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info` (
  `uid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(128) NOT NULL DEFAULT '',
  `open` tinyint(4) NOT NULL DEFAULT '1' COMMENT '邮箱是否公开，默认公开',
  `username` varchar(20) NOT NULL COMMENT '用户名',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '姓名',
  `avatar` varchar(128) NOT NULL DEFAULT '' COMMENT '头像(如果为空，则使用http://www.gravatar.com)',
  `city` varchar(10) NOT NULL DEFAULT '' COMMENT '居住地',
  `company` varchar(64) NOT NULL DEFAULT '',
  `github` varchar(20) NOT NULL DEFAULT '',
  `weibo` varchar(20) NOT NULL DEFAULT '',
  `website` varchar(50) NOT NULL DEFAULT '' COMMENT '个人主页，博客',
  `monlog` varchar(140) NOT NULL DEFAULT '' COMMENT '个人状态，签名，独白',
  `introduce` text NOT NULL COMMENT '个人简介',
  `unsubscribe` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否退订本站邮件，0-否；1-是',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '用户账号状态。0-默认；1-已审核；2-拒绝；3-冻结；4-停号',
  `is_root` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否超级用户，不受权限控制：1-是',
  `ctime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='用户信息表';

-- ----------------------------
-- Table structure for user_login
-- ----------------------------
DROP TABLE IF EXISTS `user_login`;
CREATE TABLE `user_login` (
  `uid` int(10) unsigned NOT NULL,
  `email` varchar(128) NOT NULL DEFAULT '',
  `username` varchar(20) NOT NULL COMMENT '用户名',
  `passcode` char(12) NOT NULL DEFAULT '' COMMENT '加密随机数',
  `passwd` char(32) NOT NULL DEFAULT '' COMMENT 'md5密码',
  `login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次登录时间（主动登录或cookie登录）',
  PRIMARY KEY (`uid`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  KEY `logintime` (`login_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户登录表';

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
  `uid` int(10) unsigned NOT NULL,
  `roleid` int(10) unsigned NOT NULL,
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`,`roleid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户角色表（用户是什么角色，可以多个角色）';

-- ----------------------------
-- Table structure for users_info
-- ----------------------------
DROP TABLE IF EXISTS `users_info`;
CREATE TABLE `users_info` (
  `uid` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户信息',
  `nickname` varchar(255) DEFAULT NULL COMMENT '用户昵称',
  `qq` int(11) DEFAULT NULL COMMENT 'qq号码',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=10006 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for wiki
-- ----------------------------
DROP TABLE IF EXISTS `wiki`;
CREATE TABLE `wiki` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL COMMENT 'wiki标题',
  `content` longtext NOT NULL COMMENT 'wiki内容',
  `uri` varchar(50) NOT NULL COMMENT 'uri',
  `uid` int(10) unsigned NOT NULL COMMENT '作者',
  `cuid` varchar(100) NOT NULL DEFAULT '' COMMENT '贡献者',
  `ctime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uri` (`uri`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='wiki页';
SET FOREIGN_KEY_CHECKS=1;
