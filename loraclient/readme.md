lora-app-server将权限分为3个层级: organization(组织) -> application(应用) -> device(设备).

组织即为通常意义上的用户, 而ta所管理的用户也是organization级别的.

组织下分application, 每个app理论上表示某一类相同属性的设备. 比如所有的烟感, 所有的温度传感器等.

org与user一同提供服务, 并且只需要新增操作. 封装初始化过程.

然后是gateway和sensor, 同样只提供增删改的功能即可.

尽量与数据库保持低耦合.