<template>
  <div class="dashboard-container">
    <el-button type="primary" @click.native="dialogFormVisible = true">添加节点</el-button>

    <el-dropdown style="width: 400px;margin-top: 20px;">
      <el-button type="primary" style="width: 400px;">
        选择服务 [ {{cluster_name}} ]<i class="el-icon-arrow-down el-icon--right"></i>
      </el-button>
      <el-dropdown-menu slot="dropdown" style="width: 400px;">
        <el-dropdown-item v-for="(item, i) in cluster_list"
                          style="width: 400px;" @click.native="searchList(item.cluster_name)">
          {{item.cluster_name}}
        </el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>


    <el-table
      :data="server_list"
      style="width: 100%;margin-top: 20px;">
      <el-table-column
        prop="server_name"
        label="节点名"
        width="180">
      </el-table-column>
      <el-table-column
        prop="ip"
        label="节点id"
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
          <el-button type="danger" @click.native="delServer(scope.row.ip)">删除</el-button>
        </template>
      </el-table-column>
      <el-table-column
        label="其他">
        <template slot-scope="scope">
          <el-button type="success" @click.native="editServer(scope.row.server_name, scope.row.ip)">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>


    <el-dialog title="新增节点" :visible.sync="dialogFormVisible">
      <el-form :model="serverEditForm">
        <el-form-item label="完整ip/域名" :label-width="formLabelWidth">
          <el-input v-model="serverEditForm.ip" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="节点名称" :label-width="formLabelWidth">
          <el-input v-model="serverEditForm.server_name" autocomplete="off"></el-input>
        </el-form-item>


      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click.native="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click.native="saveServer()">确 定</el-button>
      </div>
    </el-dialog>


  </div>
</template>

<script>
  import {getClusterList} from '../../api/cluster'
  import {getServerList, saveServer, delServer} from '../../api/server'
  import {AppListCurrent} from '../../enum/enum'

  export default {
    name: 'ServerList',
    components: {},
    data() {
      return {
        cluster_list: [],
        cluster_name: "",
        app_name: "",
        server_list: [],
        serverEditForm: {
          ip: "",
          server_name: "",//默认是ip
        },
        dialogFormVisible: false,
        formLabelWidth: "120px",
      }
    },
    created() {
      this.app_name = localStorage.getItem(AppListCurrent);
      this.getCluster()
    },
    methods: {
      getCluster() {
        let app_name = localStorage.getItem(AppListCurrent) || "";
        if (app_name === "") {
          this.$message("请选择应用");
          return false;
        }
        getClusterList({"app_name": app_name}).then(response => {
          this.cluster_list = response.data
        })
      },
      searchList(cluster_name) {
        this.cluster_name = cluster_name;
        this.getServerLists()
      },
      getServerLists() {
        //统一请求服务节点的方法
        let params = {
          app_name: this.app_name,
          cluster_name: this.cluster_name
        }
        if (this.app_name === "" || this.cluster_name === "") {
          this.$message("服务应用 不能为空");
          return false;
        }
        getServerList(params).then(response => {
          if (response.ret === 1) {
            this.server_list = response.data
          } else {
            this.server_list = []
          }
        })
      },
      saveServer() {
        let app_name = localStorage.getItem(AppListCurrent) || ""
        if (this.serverEditForm.ip === "" || this.serverEditForm.server_name === "" || this.cluster_name === "" || app_name === "") {
          this.$message("服务 | ip | 节点名 都不能为空")
          return false;
        }
        let params = this.serverEditForm
        params.app_name = app_name
        params.cluster_name = this.cluster_name
        saveServer(params).then(response => {
          if (response.ret === 1) {
            this.dialogFormVisible = false
            this.getServerLists()
          }
        })
      },
      editServer(server_name, ip) {
       this.serverEditForm.server_name = server_name;
       this.serverEditForm.ip = ip;
        this.dialogFormVisible = true;
      },
      delServer(ip) {
        let params = {
          app_name: this.app_name,
          cluster_name: this.cluster_name,
          ip: ip,
        }
        delServer(params).then(response => {
          if (response.ret === 1) {
            this.getServerLists()
          } else {
            this.$message("删除失败")
          }
        })
      }
    }
  }
</script>

<style lang="scss" scoped>
@import '~@/styles/defines.scss';
.dashboard {
  &-container {
    padding: 30px;
    width: 100%;
    height: 100%;
    background-color: $--my-background-color1;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
