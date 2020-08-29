### 主题切换

>打开本系统就是一个现成的主题切换Demo

element-ui 升级为 2.0 之后官方文档提供了动态换肤的功能

>* 优点
>  -无需准备多套主题，可以自由动态换肤
>* 缺点
>  -自定义不够，只支持基础颜色的切换

本系统采用`webpack-theme-color-replacer`插件结合`element-theme-chalk`依赖达到动态切换主题，并且有良好的扩展性，对插件原理感兴趣的可以参见[webpack-theme-color-replacer](https://segmentfault.com/a/1190000016061608)这篇文章

### 动态切换主题的原理

`webpack-theme-color-replacer`插件的的原理可以简单的理解为颜色替换

在这之前我们先了解几个配置`config/settings.js`

![themechange.png](/help/pic/themechange.png)

这里我们把主题色和用户自定义颜色分开配置，用户自定义颜色为数组，默认亮色主题，当切换到暗色主题，暗色主题数组将替换亮色主题数组。主题色切换同理,切换时我们将主题色和自定义颜色数组整合达到切换的目的

>**注** 自定义颜色主题数组每一个元素都应该有其对应不变的含义，如亮色主题数组第一个颜色我们定为header文字的颜色，第二个颜色定位主背景色，第三个颜色定位边框色。在暗色主题颜色数组就分别是暗色主题下的header文字颜色 主背景色 边框色

### 怎样添加主题

如添加一个蓝色主题，我们首先在`config/settings.js`内紧随`darkThemeColors`数组后面添加一个`blueThemeColors`

```json
blueThemeColors: [
  // blue主题下对应的颜色
  '#xxxxxx',
  '#xxxxxx',
  '#xxxxxx'
]
```

### 怎样添加自定义颜色

所有的自定义颜色我们规定写在`@/styles/defines.scss`文件内,并且与`config/settings.js`默认主题下的颜色代码对应

`@/styles/defines.scss`代码如下

![themechange2.png](/help/pic/themechange2.png)

红框框出的颜色和`config/settings.js`亮色主题`lightThemeColors`（默认）自定义颜色代码对应，所以我们添加自定义颜色只需要在`defines.scss`文件里面添加颜色变量并在`config/settings.js`内的各主题颜色数组下面添加该主题下对应该颜色变量的值即可。

### 怎样切换主题

**切换主题**

```javascript
changeTheme(theme) {
  this.$store.dispatch('settings/toggleTheme', theme)
}
```

>`theme`的值为`config/settings.js`内`theme`属性的可能取值（`light|dark|blue...`）

**切换主题色**

```javascript
changePrimaryColor(color) {
  this.$store.dispatch('settings/togglePrimaryColor', color)
}
```

>`color`的值为任意颜色值（`HEX、RGB、HSL`）不支持`RGBA`和`HSLA`，建议用`HEX`,也就是以`#`开头的颜色码

**有些时候如果需要`RGBA`怎么处理？**

可以利用`Sass:RGB`颜色函数-`mix()`函数,因为`element-ui`官方用的是`sass`我们这里也选择`sass`作为框架的`CSS`扩展语言

在上面`@/styles/defines.scss`截图中可以看到我们用了`mix`函数来获取比默认文本颜色浅的其它颜色

```css
$--my-text-color85: mix($--my-text-color100, $--my-text-color100, 85%);
$--my-text-color65: mix($--my-text-color100, $--my-text-color100, 65%);
$--my-text-color45: mix($--my-text-color100, $--my-text-color100, 45%);
```

>举个简单的例子`$--my-text-color100`转成`rgba`颜色码为`rgba(5, 5, 5, 1)`调用mix函数`$--my-text-color85`返回色码为`rgba(5, 5, 5, .85)`
