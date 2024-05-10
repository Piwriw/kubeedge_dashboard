<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button class="filter-item" type="primary" @click="dialogVisible = true">
        新增密钥 <i class="el-icon-upload" />
      </el-button>
    </div>
    <el-table :data="tableData" fit style="width: 100%;">
      <el-table-column prop="dockerRegistry" label="DockerRegistry" width="200">
        <template slot-scope="{row}">
          {{ row.dockerRegistry }}
        </template>
      </el-table-column>
      <el-table-column prop="dockerServer" label="DockerServer" width="200">
        <template slot-scope="{row}">
          {{ row.dockerServer }}
        </template>
      </el-table-column>
      <el-table-column prop="createtime" label="CreateTime" width="150">
        <template slot-scope="{row}" style="cursor: pointer">
          <el-tooltip class="item" effect="dark" :content="row.createTime" placement="right-start">
            <el-button style="border: none;outline: none">{{ dateTimeFormat(row.createTime) }} <i class="el-icon-info" /></el-button>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column label="Operation">
        <template slot-scope="{row}">
          <el-dropdown style="cursor: pointer" placement="bottom-start">
            <span class="el-dropdown-link">
              <svg-icon icon-class="3points-vertical" />
            </span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item>
                <router-link :to="'/edge-app/detail/'+row.name">
                  查看
                </router-link>
              </el-dropdown-item>
              <el-dropdown-item>删除</el-dropdown-item>
              <el-dropdown-item>更新</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog
      title="创建密钥"
      :visible.sync="dialogVisible"
      width="30%"
    >
      <el-form ref="form" :model="form" :rules="rules" label-width="160px">
        <el-form-item prop="dockerRegistry.value">
          <span slot="label">
            {{ form.dockerRegistry.label.title }}
            <el-tooltip placement="right">
              <div slot="content">
                {{ form.dockerRegistry.label.desc }}
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input v-model="form.dockerRegistry.value" :placeholder="form.dockerRegistry.placeholder" />
        </el-form-item>
        <el-form-item
          prop="dockerServer.value"
        >
          <span slot="label">
            {{ form.dockerServer.label.title }}
            <el-tooltip placement="right">
              <div slot="content">
                {{ form.dockerServer.label.desc }}
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input v-model="form.dockerServer.value" :placeholder="form.dockerServer.placeholder" />
        </el-form-item>
        <el-form-item
          prop="dockerUsername.value"
        >
          <span slot="label">
            {{ form.dockerUsername.label.title }}
            <el-tooltip placement="right">
              <div slot="content">
                {{ form.dockerUsername.label.desc }}
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input v-model="form.dockerUsername.value" :placeholder="form.dockerUsername.placeholder" />
        </el-form-item>
        <el-form-item
          prop="dockerPassword.value"
        >
          <span slot="label">
            {{ form.dockerPassword.label.title }}
            <el-tooltip placement="right">
              <div slot="content">
                {{ form.dockerPassword.label.desc }}
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input v-model="form.dockerPassword.value" :placeholder="form.dockerPassword.placeholder" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="onSubmit('form')">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import { fetchList } from '@/api/secret'
import { getDateDiff } from '../../utils'
export default {
  name: 'Secret',
  data() {
    return {
      dialogVisible: false,
      tableData: null,
      form: {
        dockerRegistry: {
          label: {
            title: 'DockerRegistry',
            desc: 'Docker镜像仓库名称'
          },
          value: '',
          placeholder: 'Please input docker registry'
        },
        dockerServer: {
          label: {
            title: 'DockerServer',
            desc: 'Docker镜像仓库地址'
          },
          value: '',
          placeholder: 'Please input model docker server'
        },
        dockerUsername: {
          label: {
            title: 'DockerUsername',
            desc: 'Docker镜像仓库用户名'
          },
          value: '',
          placeholder: 'Please input model version'
        },
        dockerPassword: {
          label: {
            title: 'DockerPassword',
            desc: 'Docker镜像仓库密码'
          },
          value: '',
          placeholder: 'Please input model version'
        }
      },
      rules: {
        'dockerRegistry.value': [
          { required: true, message: '请输入Docker镜像名称', trigger: 'blur' }
        ],
        'dockerServer.value': [
          { required: true, message: '请输入Docker镜像地址', trigger: 'blur' }
        ],
        'dockerUsername.value': [
          { required: true, message: '请输入Docker镜像地址用户名', trigger: 'blur' }
        ],
        'dockerPassword.value': [
          { required: true, message: '请输入Docker镜像地址密码', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    // 时间格式化为xx时间之前
    dateTimeFormat(createTime) {
      // return getDateDiff(Date.parse(new Date(createTime)))
      return getDateDiff(Date.parse(new Date(createTime)))
    },
    getList() {
      fetchList().then(response => {
        console.log(response.data)
        this.tableData = response.data
      })
    },
    onSubmit(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          alert('submit!')
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>
<style scoped>
  .el-dropdown {
    vertical-align: top;
  }

  .el-dropdown + .el-dropdown {
    margin-left: 15px;
  }

  .el-icon-arrow-down {
    font-size: 12px;
  }
</style>
