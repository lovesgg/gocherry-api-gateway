<template>
  <div class="app-container">
    <el-button type="primary" @click.native="dialogFormVisible = true">添加服务</el-button>

    <el-table
      :data="cluster_list"
      style="width: 100%;margin-top: 20px;">
      <el-table-column
        prop="title"
        label="服务名"
        width="180">
      </el-table-column>
      <el-table-column
        prop="cluster_name"
        label="服务id"
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
        label="操作">
        <template slot-scope="scope">
          <el-button @click.native="" type="text" size="small">删除</el-button>
        </template>
      </el-table-column>
      <el-table-column
        label="其他">
        <template slot-scope="scope">
          <el-button type="success" @click.native="delCluster(scope.row.cluster_name)">测试</el-button>
        </template>
      </el-table-column>
    </el-table>


    <el-dialog title="新增服务" :visible.sync="dialogFormVisible">
      <el-form :model="clusterEditForm">
        <el-form-item label="服务中文名称" :label-width="formLabelWidth">
          <el-input v-model="clusterEditForm.title" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="服务英文名称" :label-width="formLabelWidth">
          <el-input v-model="clusterEditForm.cluster_name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="服务简要描述" :label-width="formLabelWidth">
          <el-input v-model="clusterEditForm.detail" autocomplete="off"></el-input>
        </el-form-item>

      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click.native="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click.native="saveCluster()">确 定</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
  import {getClusterList, saveCluster} from '../../api/cluster'
  import {AppListCurrent} from '../../enum/enum'

  export default {
    data() {
      return {
        cluster_list: [],
        dialogFormVisible: false,
        clusterEditForm: {
          cluster_name: "",
          title: "",
          detail: "",
        },
        formLabelWidth: "120px",
      }
    },
    created() {
      this.getList()
    },
    methods: {
      saveCluster() {
        let app_name = localStorage.getItem(AppListCurrent) || ""
        if (this.clusterEditForm.cluster_name === "" || this.clusterEditForm.title === "" || this.clusterEditForm.detail === "" || app_name === "") {
          this.$message("服务内容|应用 不能为空")
          return false;
        }
        let params = this.clusterEditForm
        params.app_name = app_name
        saveCluster(params).then(response => {
          if (response.ret === 1) {
            this.dialogFormVisible = false
            this.getList()
          } else {
            this.$message(response.error.msg)
          }
        })
      },
      getList() {
        let app_name = localStorage.getItem(AppListCurrent) || ""
        if (app_name === "") {
          this.$message("请选择应用")
          return false;
        }
        getClusterList({"app_name": app_name}).then(response => {
          if (response.ret === 1) {
            this.cluster_list = response.data
          }
        })
      }
    }
  }
</script>
