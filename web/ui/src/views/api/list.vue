<template>
  <div class="app-container">
    <el-button type="primary" @click.native="editApi()">添加api</el-button>

    <el-table
      :data="api_list"
      style="width: 100%;margin-top: 20px;">
      <el-table-column
        prop="base_api_name"
        label="接口名"
        width="180">
      </el-table-column>
      <el-table-column
        prop="base_api_url"
        label="url"
        width="180">
      </el-table-column>
      <el-table-column
        prop="base_cluster_name"
        label="对应服务"
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
        label="其他">
        <template slot-scope="scope">
          <el-button type="success" @click.native="turnEdit(scope.row.base_api_url, scope.row.base_cluster_name)">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>

  import {getApiList} from '../../api/api'

  export default {
    data() {
      return {
        api_list: [],
      }
    },
    created() {
      this.getList()
    },
    methods: {
      getList() {
        getApiList({}).then(response => {
          if (response.ret === 1) {
            this.api_list = response.data
          }
        })
      },
      editApi() {
        this.$router.push("/api/edit")
      },
      turnEdit(url, base_cluster_name) {
        this.$router.push("/api/edit?api_url=" + url + "&base_cluster_name=" + base_cluster_name)
      }
    }
  }
</script>
