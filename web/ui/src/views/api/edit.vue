<template>
  <div class="app-container">
    <el-form :model="apiForm" :rules="rules" ref="apiForm" label-width="200px" style="width: 600px;" class="demo-ruleForm">

      <el-form-item label="接口名" prop="base_api_name" required>
        <el-input v-model="apiForm.base_api_name"></el-input>
      </el-form-item>

      <el-form-item label="入口url" prop="base_api_url" required="">
        <el-input v-model="apiForm.base_api_url"></el-input>
      </el-form-item>

      <el-form-item label="转发后端url" prop="base_redirect_url" required="">
        <el-input v-model="apiForm.base_redirect_url"></el-input>
      </el-form-item>

      <el-form-item label="请求类型" prop="base_api_request_type" required>
        <el-checkbox-group v-model="apiForm.base_api_request_type">
          <el-checkbox label="POST"></el-checkbox>
          <el-checkbox label="GET"></el-checkbox>
        </el-checkbox-group>
      </el-form-item>

      <el-form-item label="集群(服务名)" prop="base_cluster_name" required>
        <el-select v-model="apiForm.base_cluster_name" placeholder="请选择集群">
          <el-option v-for="(item,i) in cluster_list" :key="i" :label="item.cluster_name" :value="item.cluster_name"></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="是否开启接口" prop="base_api_status">
        <el-switch v-model="apiForm.base_api_status"></el-switch>
      </el-form-item>

      <el-form-item label="是否降级" prop="reduce_level">
        <el-switch v-model="apiForm.reduce_level"></el-switch>
      </el-form-item>

      <el-form-item label="是否需要登录" prop="user_auth">
        <el-switch v-model="apiForm.user_auth"></el-switch>
      </el-form-item>

      <el-form-item label="缓存/秒(0就是禁用)" prop="cache_save" required="">
        <el-input v-model.number="apiForm.cache_save"></el-input>
      </el-form-item>

      <el-form-item label="接口请求数限制/秒" prop="limit_request" required="">
        <el-input v-model.number="apiForm.limit_request"></el-input>
      </el-form-item>

      <el-form-item label="接口超时/秒" prop="time_out" required="">
        <el-input v-model.number="apiForm.time_out"></el-input>
      </el-form-item>

      <el-form-item label="ip黑名单/ip用逗号隔离">
        <el-input type="textarea" v-model="apiForm.ip_black"></el-input>
      </el-form-item>

      <el-form-item label="只允许白名单访问/手机号" prop="white_list_check">
        <el-switch v-model="apiForm.white_list_check"></el-switch>
      </el-form-item>

      <el-form-item label="白名单">
        <el-input type="textarea" v-model="apiForm.white_list"></el-input>
      </el-form-item>










      <el-form-item>
        <el-button type="primary" @click="submitForm('apiForm')">保存接口</el-button>
        <el-button @click="resetForm('apiForm')">重置</el-button>
      </el-form-item>


    </el-form>
  </div>
</template>

<script>
  import {getOneApi, saveApi} from '../../api/api'
  import {getClusterList} from '../../api/cluster'
  import {AppListCurrent} from '../../enum/enum'

  export default {
    data() {
      return {
        apiForm: {
          base_api_name: '',
          base_api_url: '',
          base_redirect_url: '',
          base_api_request_type: [],
          base_api_status: true,
          base_cluster_name: '',
          cache_save: 0,
          limit_request: 10000,
          reduce_level: false,
          time_out: 3,
          ip_black: '',
          user_auth: false,
          white_list_check: false,
          white_list: "",


        },
        rules: {
          base_api_name: [
            {required: true, message: '接口名'},
          ],
          base_api_url: [
            {required: true, message: '入口url'}
          ],
          base_redirect_url: [
            {required: true, message: '转发url'}
          ],
          base_api_request_type: [
            {required: true, message: '请求类型'}
          ],
          base_api_status: [
            {required: true, message: '状态'}
          ],
          base_cluster_name: [
            {required: true, message: '集群名称'}
          ],
          cache_save: [
            {required: true, message: '缓存时间'}
          ],
          limit_request: [
            {required: true, message: '接口qps', type: 'number'}
          ],
          reduce_level: [
            {required: true, message: '降级状态'}
          ],
          user_auth: [
            {required: true, message: '登录校验'}
          ],
          white_list_check: [
            {required: true, message: '白名单状态'}
          ]

        },
        cluster_list: [],
        app_name: "",
      }
    },
    created() {
      let app_name = localStorage.getItem(AppListCurrent) || "";
      let api_url = this.$route.query.api_url || ""
      let base_cluster_name = this.$route.query.base_cluster_name;
      if (app_name === "") {
        this.$message("请选择应用");
        return false;
      }
      this.app_name = app_name;
      this.getClusters();
      if (api_url && base_cluster_name) {
        //已存在 编辑
        getOneApi({
          "base_api_url": api_url,
          "app_name": app_name,
          "base_cluster_name": base_cluster_name
        }).then(response => {
          if (response.ret === 1) {
            this.apiForm = response.data
          }
        })
      }
    },
    methods: {
      getClusters() {
        getClusterList({"app_name": this.app_name}).then(response => {
          if (response.ret === 1) {
            this.cluster_list = response.data
          }
        })
      },
      submitForm(formName) {
        //验证表单
        this.$refs[formName].validate((valid) => {
          if (valid) {
            //组装参数
            let params = {
              api_form: this.apiForm,
              app_name: this.app_name
            };

            saveApi(params).then(response => {
              if (response.ret === 1) {
                this.$router.push("/api/list")
              } else {
                this.$message(response.error.msg || "保存失败")
              }
            })
          } else {
            this.$message("请检查输入")
          }
        });
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      }
    }
  }
</script>
