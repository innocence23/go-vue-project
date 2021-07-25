<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="路径">
          <el-input v-model="searchInfo.path" placeholder="路径" />
        </el-form-item>
        <el-form-item label="简介">
          <el-input v-model="searchInfo.description" placeholder="描述" />
        </el-form-item>
        <el-form-item label="组">
          <el-input v-model="searchInfo.group" placeholder="组" />
        </el-form-item>
        <el-form-item label="请求">
          <el-select v-model="searchInfo.method" clearable placeholder="请选择">
            <el-option
              v-for="item in methodOptions"
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" border stripe @sort-change="sortChange" @selection-change="handleSelectionChange">
      <el-table-column
        type="selection"
        width="55"
      />
      <el-table-column label="id" min-width="60" prop="ID" sortable="custom" />
      <el-table-column label="路径" min-width="150" prop="path" sortable="custom" />
      <el-table-column label="分组" min-width="150" prop="group" sortable="custom" />
      <el-table-column label="简介" min-width="150" prop="description" sortable="custom" />
      <el-table-column label="请求" min-width="150" prop="method" sortable="custom">
        <template slot-scope="scope">
          <div>
            {{ scope.row.method }}
            <el-tag
              :key="scope.row.methodFiletr"
              :type="scope.row.method|tagTypeFiletr"
              effect="dark"
              size="mini"
            >{{ scope.row.method|methodFiletr }}</el-tag>
            <!-- {{scope.row.method|methodFiletr}} -->
          </div>
        </template>
      </el-table-column>

      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button
            size="small"
            type="danger"
            icon="el-icon-delete"
            @click="deleteApi(scope.row)"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      layout="total, sizes, prev, pager, next, jumper"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
  </div>
</template>

<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索

import {
  getApiList,
} from '@/api/permission'
import infoList from '@/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'
const methodOptions = [
  {
    value: 'POST',
    label: '创建',
    type: 'success'
  },
  {
    value: 'GET',
    label: '查看',
    type: ''
  },
  {
    value: 'PUT',
    label: '更新',
    type: 'warning'
  },
  {
    value: 'DELETE',
    label: '删除',
    type: 'danger'
  }
]

export default {
  name: 'Api',
  filters: {
    methodFiletr(value) {
      const target = methodOptions.filter(item => item.value === value)[0]
      // return target && `${target.label}(${target.value})`
      return target && `${target.label}`
    },
    tagTypeFiletr(value) {
      const target = methodOptions.filter(item => item.value === value)[0]
      return target && `${target.type}`
    }
  },
  mixins: [infoList],
  data() {
    return {
      deleteVisible: false,
      listApi: getApiList,
      dialogFormVisible: false,
      dialogTitle: '新增Api',
      apis: [],
      form: {
        path: '',
        apiGroup: '',
        method: '',
        description: ''
      },
      methodOptions: methodOptions,
      type: '',
      rules: {
        path: [{ required: true, message: '请输入路径', trigger: 'blur' }],
        apiGroup: [
          { required: true, message: '请输入组名称', trigger: 'blur' }
        ],
        method: [
          { required: true, message: '请选择请求方式', trigger: 'blur' }
        ],
        description: [
          { required: true, message: '请输入介绍', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getTableData()
  },
  methods: {
    //  选中api
    handleSelectionChange(val) {
      this.apis = val
    },
    // 排序
    sortChange({ prop, order }) {
      if (prop) {
        this.searchInfo.orderKey = toSQLLine(prop)
        this.searchInfo.desc = order === 'descending'
      }
      this.getTableData()
    },
    // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    initForm() {
      this.$refs.apiForm.resetFields()
      this.form = {
        path: '',
        apiGroup: '',
        method: '',
        description: ''
      }
    },
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
  }
}
</script>

<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.el-tag--mini {
  margin-left: 5px;
}
.warning {
  color: #dc143c;
}
</style>
