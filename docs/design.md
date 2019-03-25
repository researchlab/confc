# design 

## confc 覆盖

### Feature

`输入`: 一份template

`输出`: 多份生成配置

 配置生成环境，

 配置是否覆盖， 当存在同一文件则覆盖， 则不覆盖;

 判断是否存在同名文件， 存在则重命名 配置名.时间戳.工具名.后缀,  生成新的配置；

 配置文件通过环境前缀生成;

 命令 可以读取flag指令， 也可以读取环境变量， 生成时， 需要显示环境生成列表， 是否覆盖,

 配置生成路径， 分相对路径， 和绝对路径，

 默认目录为执行路径下 新建文件夹config生成配置文件，

 配置文件模板解析类型; json, yaml, toml;

 显示当前路径，

 环境生成，

 环境配置生成文件，

 环境配置另存文件,

 显示当前路径的相对路径,


加载模板 {配置模板路径名称} 以它的后缀为解析前置; --tpl ==  (模板前缀为合成前缀)

生成配置缓存,

加载已存在的配置文件,

生成配置新配置文件,  {生成配置文件路径名称} 保存策略{default 保存原有的生成 加时间戳 } -- 加选项直接替换; -- dist (path/filename) ==  , --policy==off (default:on)

--path

--生成名称为模板名称;

confc --env test,local,prod,15, --tpls==指定文件夹 --tpl==指定一个模板文件 --dist (path/filename)

confc --tpl --dist --replace=off (default)

env_tpl_name.suffix

env_tpl_name.json, conf,

env file-folder,

default: config/env

--dist:

--type json,yaml,yml,toml (合成前缀) 没有指定 默认以模板后缀为前缀;

--tmpl 判断文件 或文件夹

CONFC_TMPL=(相对/非相对路径)  --tmpl (当前路径 *.confc文件) [数组]

CONFC_ENV= (环境以逗号隔离) --env (dev)

CONFC_DIST=默认生成路径 --dist (config)

CONFC_TYPE=解析类型;  --type (json)

用户自定义配置,

环境配置,

默认配置,
