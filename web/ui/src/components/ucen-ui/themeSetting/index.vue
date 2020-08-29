<template>
  <div class="theme-setting">
    <el-dropdown>
      <span class="el-dropdown-link">
        {{appCurrent || "App"}}
      </span>
      <el-dropdown-menu slot="dropdown">
        <el-dropdown-item v-for="(item, i) in appList" @click.native="changeApp(item.app_name)">{{item.app_name}}</el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>

    <el-dropdown style="margin-left: 10px;">
      <span class="el-dropdown-link">
        设置<i class="el-icon-arrow-down el-icon--right"></i>
      </span>
      <el-dropdown-menu slot="dropdown">
        <el-dropdown-item @click.native="logout()">退出</el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>

  </div>
</template>

<script>
  import {AppListCurrent, LoginToken} from '../../../enum/enum'
  import {getAppList} from '../../../api/app'

  export default {
    name: 'ThemeSetting',
    props: {},
    data() {
      return {
        colorList: [
          {
            key: '莓红', color: '#C45A65'
          },
          {
            key: '玫瑰粉', color: '#f8b37f'
          },
          {
            key: '葵扇黄', color: '#e8b004'
          },
          {
            key: '蛙绿', color: '#45b787'
          },
          {
            key: '碧青', color: '#5cb3cc'
          },
          {
            key: '宝石蓝', color: '#2486b9'
          },
          {
            key: '洋葱紫', color: '#a8456b'
          }
        ],
        dialogTableVisible: false,
        appList: [],
        appCurrent: "",
      }
    },
    created() {
      this.getAppList()
    },
    methods: {
      openSetting() {
        this.dialogTableVisible = true
      },
      changeTheme(theme) {
        this.$store.dispatch('settings/toggleTheme', theme)
      },
      changePrimaryColor(color) {
        this.$store.dispatch('settings/togglePrimaryColor', color)
      },
      getAppList() {
        getAppList({}).then(response => {
          this.appList = response.data
          let currentAppLocal = localStorage.getItem(AppListCurrent) || ""
          console.log(currentAppLocal)
          if (currentAppLocal) {
            this.appCurrent = currentAppLocal || "无应用"
          } else {
            this.appCurrent = response.data[0]["app_name"] || "无应用"
          }
        })

      },
      changeApp(appName) {
        localStorage.setItem(AppListCurrent, appName)
        this.$router.go(0)
      },
      logout() {
        localStorage.setItem(LoginToken, "")
        localStorage.setItem(AppListCurrent, "")
        this.$router.push("/login");
      }
    }
  }
</script>

<style lang="scss" scoped>
@import "index";

.el-dropdown-link {
  cursor: pointer;
  color: #17202A;
}
.el-icon-arrow-down {
  font-size: 12px;
}
</style>
