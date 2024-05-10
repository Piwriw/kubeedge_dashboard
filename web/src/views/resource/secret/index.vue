<template>
  <div class="app-container">

    <el-dropdown>
  <span class="custom-dropdown-link">
    {{ currentNamespace }}
<i class="el-icon-arrow-down el-icon--right"></i>
  </span>
      <el-dropdown-menu slot="dropdown">
        <el-dropdown-item v-for="ns in  nsList" @click.native="getSecretList(ns.metadata.name)">
          {{ ns.metadata.name }}
        </el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>

    <el-table
      :data="secretList"
      fit
      highlight-current-row
      style="width: 100%;"
    >
      <el-table-column prop="namespace" label="namespace" width="200" align="center"/>
      <el-table-column prop="name" label="name" width="200" align="center"/>
      <el-table-column prop="type" label="type" width="200" align="center"/>

      <el-table-column prop="age" label="创建时间" width="150">
        <template slot-scope="{row}" style="cursor: pointer">
          <el-tooltip class="item" effect="dark" :content="row.create_time" placement="right-start">
            <el-button style="border: none;outline: none">{{ dateTimeFormat(row.create_time) }} <i
              class="el-icon-info"/></el-button>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="{row}">
          <el-dropdown style="cursor: pointer" placement="bottom-start">
            <span class="el-dropdown-link">
              <svg-icon icon-class="3points-vertical"/>
            </span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item @click.native="showAppDetail(row.name,row.namespace)">查看</el-dropdown-item>
              <!--            <el-dropdown-item @click.native="showDeleteApp(row.name,row.namespace)">删除</el-dropdown-item>-->
              <!--<el-dropdown-item>升级</el-dropdown-item>-->
              <!--<el-dropdown-item>日志</el-dropdown-item>-->
            </el-dropdown-menu>
          </el-dropdown>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog
      title="Secret详情"
      :visible.sync="dialogVisible2detail"
      width="50%"
      :before-close="handleBeforeClose"
      :destroy-on-close="true"
    >
      <div class="dialog-content">
        <el-table
          :data="detailContentList"
          border
          style="width: 100%"
        >
          <el-table-column
            prop="key"
            label="key"
            width="220"
          ></el-table-column>
          <el-table-column
            prop="value"
            label="value"
            width="300"
          >
            <template slot-scope="scope">
              <div>
      <span v-if="scope.row.value.length > 25 && scope.row.isCollapsed">
        {{ scope.row.value.substring(0, 25) }}...
      </span>
                <span v-else>
        {{ scope.row.value }}
      </span>
                <a href="javascript:void(0)" v-if="scope.row.value.length > 25" @click="toggleCollapse(scope.row)">
                  {{ scope.row.isCollapsed ? '展开' : '收起' }}
                </a>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-dialog>

  </div>
</template>

<script>
import {getAllNamespace} from "@/api/namespace";
import {getSecretList} from "@/api/resource"
import {getSecretDetails} from "@/api/resource"
import {dateTimeFormat} from "@/utils";

export default {
  name: 'index',
  data() {
    return {
      currentNamespace: 'default',
      dialogVisible2detail: false,
      detailContentList: [],
      nsList: [],
      secretList: [],
      list: null,
      total: 0,
      pageInfo: {
        page: 1,
        size: 10
      }
    }
  },
  created() {
    this.getList()
    this.getSecretList('default')
  },
  updated() {
    // 在创建组件时为每一行数据添加isCollapsed属性并设置默认值为true
    this.detailContentList.forEach((row) => {
      this.$set(row, 'isCollapsed', true);
    });
    // 当对话框内容发生变化时重新计算对话框高度
    this.$nextTick(() => {
      const dialogContent = this.$el.querySelector('.dialog-content');
      const dialogHeight = dialogContent.scrollHeight;
      dialogContent.style.maxHeight = `${dialogHeight}px`;
    });
  },
  methods: {
    dateTimeFormat,
    toggleCollapse(row) {
      this.$set(row, 'isCollapsed', !row.isCollapsed);
      this.handleBeforeClose()
    },

    handleBeforeClose(done) {
      // 在对话框关闭前重置对话框高度，以便下次展开时重新计算高度
      const dialogContent = this.$el.querySelector('.dialog-content');
      dialogContent.style.maxHeight = '';
      done();
    },
    showAppDetail(name, namespace) {
      getSecretDetails(name, namespace).then(res => {
        this.detailContentList = res.data.data
        this.dialogVisible2detail = true
      })
    },
    getList() {
      getAllNamespace().then(res => {
        this.nsList = res.data
      })
    },
    getSecretList(namespace) {
      this.currentNamespace = namespace
      getSecretList(namespace).then(res => {
        this.secretList = res.data
      })
    }
  }
}
</script>

<style scoped>
.dialog-content {
  max-height: 400px; /* 初始最大高度，根据需要进行调整 */
  overflow-y: auto;
}

.custom-dropdown-link {
  margin-left: 50px;
  display: inline-block;
  padding: 5px 10px;
  font-weight: bold;
  color: #333;
  cursor: pointer;
  border: 1px solid #ccc;
  background-color: #51e09f;
  border-radius: 4px;
}

.custom-dropdown-link .el-icon-arrow-down {
  margin-left: 5px;
  transition: transform 0.3s;
}

.custom-dropdown-link.is-active .el-icon-arrow-down {
  transform: rotate(90deg);
}
</style>
