<template>
  <div>
    <div class="clearflex">
      <el-button class="fl-right" size="small" type="primary" @click="authApiEnter">确 定</el-button>
    </div>
    <el-tree
      ref="apiTree"
      :data="apiTreeData"
      :default-checked-keys="apiTreeIds"
      :props="apiDefaultProps"
      default-expand-all
      highlight-current
      node-key="onlyId"
      show-checkbox
      @check="nodeChange"
    />
  </div>
</template>

<script>
import { getAllApis } from '@/api/permission'
import { UpdateCasbin, getPolicyPathByRoleId } from '@/api/casbin'
export default {
  name: 'Permissions',
  props: {
    row: {
      default: function() {
        return {}
      },
      type: Object
    }
  },
  data() {
    return {
      apiTreeData: [],
      apiTreeIds: [],
      needConfirm: false,
      apiDefaultProps: {
        children: 'children',
        label: 'description'
      }
    }
  },
  async created() {
    // 获取api并整理成树结构
    const res2 = await getAllApis()
    const permissions = res2.data.permissions

    this.apiTreeData = this.buildApiTree(permissions)
    const res = await getPolicyPathByRoleId({
      roleId: this.row.roleId
    })
    this.activeUserId = this.row.roleId
    this.apiTreeIds = []
    res.data.paths && res.data.paths.map(item => {
      this.apiTreeIds.push('p:' + item.path + 'm:' + item.method)
    })
  },
  methods: {
    nodeChange() {
      this.needConfirm = true
    },
    // 暴露给外层使用的切换拦截统一方法
    enterAndNext() {
      this.authApiEnter()
    },
    // 创建api树方法
    buildApiTree(permissions) {
      const apiObj = {}
      permissions &&
        permissions.map(item => {
          item.onlyId = 'p:' + item.path + 'm:' + item.method
          if (Object.prototype.hasOwnProperty.call(apiObj, item.group)) {
            apiObj[item.group].push(item)
          } else {
            Object.assign(apiObj, { [item.group]: [item] })
          }
        })
      const apiTree = []
      for (const key in apiObj) {
        const treeNode = {
          ID: key,
          description: key + '组',
          children: apiObj[key]
        }
        apiTree.push(treeNode)
      }
      return apiTree
    },
    // 关联关系确定
    async authApiEnter() {
      const checkArr = this.$refs.apiTree.getCheckedNodes(true)
      var casbinInfos = []
      checkArr && checkArr.map(item => {
        var casbinInfo = {
          path: item.path,
          method: item.method
        }
        casbinInfos.push(casbinInfo)
      })
      const res = await UpdateCasbin({
        roleId: this.activeUserId,
        casbinInfos
      })
      if (res.code === 0) {
        this.$message({ type: 'success', message: 'api设置成功' })
      }
    }
  }
}
</script>
