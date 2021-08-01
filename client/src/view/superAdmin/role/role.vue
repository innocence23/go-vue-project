<template>
  <div class="role">
    <div class="button-box clearflex">
      <el-button size="mini" type="primary" icon="el-icon-plus" @click="addRole(0)">新增角色</el-button>
    </div>
    <el-table
      :data="tableData"
      :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
      border
      row-key="ID"
      stripe
      style="width: 100%"
    >
      <el-table-column label="角色id" min-width="180" prop="ID" />
      <el-table-column label="角色名称" min-width="180" prop="name" />
      <el-table-column fixed="right" label="操作" width="460">
        <template slot-scope="scope">
          <el-button size="mini" type="primary" @click="opdendrawer(scope.row)">设置权限</el-button>
          <el-button
            icon="el-icon-plus"
            size="mini"
            type="primary"
            @click="addRole(scope.row.id)"
          >新增子角色</el-button>
          <el-button
            icon="el-icon-edit"
            size="mini"
            type="primary"
            @click="editRole(scope.row)"
          >编辑</el-button>
          <el-button
            icon="el-icon-delete"
            size="mini"
            type="danger"
            @click="deleteAuth(scope.row)"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 新增角色弹窗 -->
    <el-dialog :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form ref="roleForm" :model="form" :rules="rules">
        <el-form-item label="父级角色" prop="parentId">
          <el-cascader
            v-model="form.parentId"
            :disabled="dialogType=='add'"
            :options="RoleOption"
            :props="{ checkStrictly: true,label:'name',value:'ID',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          />
        </el-form-item>
        <el-form-item label="角色姓名" prop="name">
          <el-input v-model="form.name" autocomplete="off" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>

    <el-drawer v-if="drawer" :visible.sync="drawer" :with-header="false" size="40%" title="角色配置">
      <el-tabs :before-leave="autoEnter" class="role-box" type="border-card">
        <el-tab-pane label="角色菜单">
          <Menus ref="menus" :row="activeRow" />
        </el-tab-pane>
        <el-tab-pane label="角色权限">
          <apis ref="apis" :row="activeRow" />
        </el-tab-pane>
      </el-tabs>
    </el-drawer>
  </div>
</template>

<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成

import {
  getRoleList,
  deleteRole,
  createRole,
  updateRole
} from '@/api/role'

import Menus from '@/view/superAdmin/role/components/menus'
import Apis from '@/view/superAdmin/role/components/permissions'

import infoList from '@/mixins/infoList'
export default {
  name: 'Role',
  components: {
    Menus,
    Apis
  },
  mixins: [infoList],
  data() {
    var mustUint = (rule, value, callback) => {
      if (!/^[0-9]*[1-9][0-9]*$/.test(value)) {
        return callback(new Error('请输入正整数'))
      }
      return callback()
    }

    return {
      RoleOption: [
        {
          ID: 0,
          name: '根角色'
        }
      ],
      listApi: getRoleList,
      drawer: false,
      dialogType: 'add',
      activeRow: {},
      activeUserId: 0,
      dialogTitle: '新增角色',
      dialogFormVisible: false,
      apiDialogFlag: false,
      form: {
        ID: 0,
        name: '',
        parentId: 0
      },
      rules: {
        name: [
          { required: true, message: '请输入角色名', trigger: 'blur' }
        ],
        parentId: [
          { required: true, message: '请输入父级ID', trigger: 'blur' }
        ]
      }
    }
  },
  async created() {
    this.pageSize = 999
    await this.getTableData()
  },
  methods: {
    autoEnter(activeName, oldActiveName) {
      const paneArr = ['menus', 'apis', 'datas']
      if (oldActiveName) {
        if (this.$refs[paneArr[oldActiveName]].needConfirm) {
          this.$refs[paneArr[oldActiveName]].enterAndNext()
          this.$refs[paneArr[oldActiveName]].needConfirm = false
        }
      }
    },
    opdendrawer(row) {
      this.drawer = true
      this.activeRow = row
    },
    // 删除角色
    deleteAuth(row) {
      this.$confirm('此操作将永久删除该角色, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteRole({ ID: row.ID })
          if (res.code === 0) {
            this.$message({
              type: 'success',
              message: '删除成功!'
            })
            if (this.tableData.length === 1 && this.page > 1) {
              this.page--
            }
            this.getTableData()
          }
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    },
    // 初始化表单
    initForm() {
      if (this.$refs.roleForm) {
        this.$refs.roleForm.resetFields()
      }
      this.form = {
        ID: 0,
        name: '',
        parentId: 0
      }
    },
    // 关闭窗口
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
      this.apiDialogFlag = false
    },
    // 确定弹窗

    async enterDialog() {
      if (this.form.ID === 0) {
        this.$message({
          type: 'error',
          message: '角色id不能为0'
        })
        return false
      }
      this.$refs.roleForm.validate(async valid => {
        if (valid) {
          switch (this.dialogType) {
            case 'add':
              {
                const res = await createRole(this.form)
                if (res.code === 0) {
                  this.$message({
                    type: 'success',
                    message: '添加成功!'
                  })
                  this.getTableData()
                  this.closeDialog()
                }
              }
              break
            case 'edit':
              {
                const res = await updateRole(this.form)
                if (res.code === 0) {
                  this.$message({
                    type: 'success',
                    message: '添加成功!'
                  })
                  this.getTableData()
                  this.closeDialog()
                }
              }
              break
          }

          this.initForm()
          this.dialogFormVisible = false
        }
      })
    },
    setOptions() {
      this.RoleOption = [
        {
          ID: 0,
          name: '根角色'
        }
      ]
      this.setRoleOptions(this.tableData, this.RoleOption, false)
    },
    setRoleOptions(RoleData, optionsData, disabled) {
      RoleData &&
        RoleData.map(item => {
          if (item.children && item.children.length) {
            const option = {
              ID: item.ID,
              name: item.name,
              disabled: disabled || item.ID === this.form.ID,
              children: []
            }
            this.setRoleOptions(
              item.children,
              option.children,
              disabled || item.ID === this.form.ID
            )
            optionsData.push(option)
          } else {
            const option = {
              ID: item.ID,
              name: item.name,
              disabled: disabled || item.ID === this.form.ID
            }
            optionsData.push(option)
          }
        })
    },
    // 增加角色
    addRole(parentId) {
      this.initForm()
      this.dialogTitle = '新增角色'
      this.dialogType = 'add'
      this.form.parentId = parentId
      this.setOptions()
      this.dialogFormVisible = true
    },
    // 编辑角色
    editRole(row) {
      this.setOptions()
      this.dialogTitle = '编辑角色'
      this.dialogType = 'edit'
      for (const key in this.form) {
        this.form[key] = row[key]
      }
      this.setOptions()
      this.dialogFormVisible = true
    }
  }
}
</script>

<style lang="scss">
.ro le {
  .el-input-number {
    margin-left: 15px;
    span {
      display: none;
    }
  }
  .button-box {
    padding: 10px 20px;
    .el-button {
      float: right;
    }
  }
}
.role-box {
  .el-tabs__content {
    height: calc(100vh - 150px);
    overflow: auto;
  }
}
</style>
