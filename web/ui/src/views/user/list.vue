<template>
  <div class="app-container">
    <el-button type="primary" @click.native="dialogFormVisible = true">添加账户</el-button>

    <el-table
      :data="users"
      style="width: 100%;margin-top: 20px;">
      <el-table-column
        prop="user_name"
        label="名称"
        width="180">
      </el-table-column>
      <el-table-column
        prop="phone"
        label="手机"
        width="180">
      </el-table-column>
      <el-table-column
        prop="level"
        label="级别"
        width="180">
      </el-table-column>

      <el-table-column
        label="其他">
        <template slot-scope="scope">
          <el-button type="danger" @click.native="opStatus(scope.row.phone)">禁用</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog title="新增账户" :visible.sync="dialogFormVisible">
      <el-form :model="editForm" :rules="rules" ref="editForm">
        <el-form-item label="姓名" :label-width="formLabelWidth" prop="user_name">
          <el-input v-model="editForm.user_name"></el-input>
        </el-form-item>
        <el-form-item label="手机号" :label-width="formLabelWidth" prop="phone">
          <el-input v-model="editForm.phone"></el-input>
        </el-form-item>
        <el-form-item label="密码" :label-width="formLabelWidth" prop="pwd">
          <el-input v-model="editForm.pwd"></el-input>
        </el-form-item>
        <el-form-item label="级别" :label-width="formLabelWidth" prop="level">
          <el-select v-model.number="editForm.level" placeholder="请选择">
            <el-option
              v-for="item in levelOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>

      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click.native="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click.native="save('editForm')">确 定</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
  import {getUserList, saveUser, delUser} from '../../api/user'

  export default {
    data() {
      return {
        users: [],
        dialogFormVisible: false,
        editForm: {
          user_name: "",
          phone: "",
          pwd: "",
          level: 1,
        },
        levelOptions: [
          {label: "普通", value: 1},
          {label: "开发", value: 2},
          {label: "超级", value: 3},
        ],
        formLabelWidth: "120px",
        rules: {
          user_name: [
            {required: true, message: '姓名'},
          ],
          phone: [
            {required: true, message: '手机'},
          ],
          pwd: [
            {required: true, message: '密码'},
          ],
          level: [
            {required: true, message: '级别'},
          ],
        },
      }
    },
    created() {
      this.getList()
    },
    methods: {
      getList() {
        getUserList().then(response => {
          if (response.ret === 1) {
            this.users = response.data;
          }
        })
      },
      opStatus(phone) {
        delUser({phone: phone}).then(response => {
          if (response.ret === 1) {
            this.getList();
          } else {
            this.$message("删除失败");
          }
        })
      },
      save(formName) {
        //验证表单
        this.$refs[formName].validate((valid) => {
          if (valid) {
            saveUser(this.editForm).then(response => {
              if (response.ret === 1) {
                this.dialogFormVisible = false;
                this.getList()
              } else {
                this.$message(response.error.msg || "保存失败")
              }
            })
          } else {
            this.$message("请检查输入")
          }
        });
      }

    }
  }
</script>
