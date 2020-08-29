<template>
  <div class="app-container">
    <el-button type="primary" @click.native="dialogFormVisible = true">添加应用</el-button>

    <el-table
      :data="app_list"
      style="width: 100%;margin-top: 20px;">
      <el-table-column
        prop="title"
        label="应用名"
        width="180">
      </el-table-column>
      <el-table-column
        prop="app_name"
        label="应用id"
        width="180">
      </el-table-column>
      <el-table-column
        prop="detail"
        label="明细"
        width="180">
      </el-table-column>
      <el-table-column
        prop="update_time"
        label="时间"
        width="180">
      </el-table-column>
      <el-table-column
        label="操作"
        width="100">
        <template slot-scope="scope">
          <el-button @click.native="" type="text" size="small">删除</el-button>
        </template>
      </el-table-column>
      <el-table-column
        label="变更应用">
        <template slot-scope="scope">
          <el-button type="success" @click.native="setCurrentApp(scope.row.app_name)">设置为当前应用</el-button>
        </template>
      </el-table-column>
    </el-table>


    <el-dialog title="新增应用" :visible.sync="dialogFormVisible">
      <el-form :model="appEditForm">
        <el-form-item label="应用中文名称" :label-width="formLabelWidth">
          <el-input v-model="appEditForm.title" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="应用英文名称" :label-width="formLabelWidth">
          <el-input v-model="appEditForm.app_name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="应用简要描述" :label-width="formLabelWidth">
          <el-input v-model="appEditForm.detail" autocomplete="off"></el-input>
        </el-form-item>

      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click.native="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click.native="saveApp()">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
  import {getAppList, saveApp} from '../../api/app'
  import {AppListCurrent} from '../../enum/enum'

  export default {
    data() {
      return {
        app_list: [],
        dialogFormVisible: false,
        formLabelWidth: '120px',
        appEditForm: {
          app_name: "",
          create_time: "",
          title: "",
          detail: "",
        },

      }
    },
    created() {
      this.getList()
    },
    methods: {
      getList() {
        getAppList({}).then(response => {
          if (response.ret === 1) {
            this.app_list = response.data
          }
        })
      },
      saveApp() {
        if (this.appEditForm.app_name === "" || this.appEditForm.title === "" || this.appEditForm.detail === "") {
          this.$message("应用内容不能为空")
          return false;
        }
        let params = this.appEditForm
        saveApp(params).then(response => {
          if (response.ret === 1) {
            this.dialogFormVisible = false
            this.getList()
          } else {
            this.$message(response.error.msg)
          }
        })
      },
      setCurrentApp(appName) {
        localStorage.setItem(AppListCurrent, appName)
        this.$router.go(0)
      }
    }
  }
</script>
