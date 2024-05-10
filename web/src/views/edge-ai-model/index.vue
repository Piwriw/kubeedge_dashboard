<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button class="filter-item" type="primary" @click="dialogVisible=true">
        上传模型 <i class="el-icon-upload" />
      </el-button>
    </div>
    <el-table :data="tableData" fit style="width: 100%;">
      <el-table-column prop="version" label="模型版本" width="200" align="center" />
      <el-table-column prop="class_names" label="分类信息" width="200" />
      <el-table-column prop="file_md5" label="文件Md5" width="400" align="center" />
      <el-table-column prop="file_path" label="文件保存路径" width="600" align="center" />
      <el-table-column label="操作">
        <template slot-scope="{row}">
          <el-dropdown style="cursor: pointer" placement="bottom-start">
            <span class="el-dropdown-link">
              <svg-icon icon-class="3points-vertical" />
            </span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item @click.native="distributeModel(row.version)">模型下发</el-dropdown-item>
              <el-dropdown-item @click.native="showDeleteModel(row.version)">删除</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog
      title="模型下发"
      :visible.sync="dialogVisible2distribute"
      width="30%"
    ></el-dialog>
    <el-dialog
      title="上传模型"
      :visible.sync="dialogVisible"
      width="30%"
    >
      <el-form ref="form" :model="form" :rules="rules" label-width="100px">
        <el-form-item
          prop="version.value"
        >
          <span slot="label">
            {{ form.version.label.title }}
            <el-tooltip placement="right">
              <div slot="content">
                {{ form.version.label.desc }}
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input v-model="form.version.value" :placeholder="form.version.placeholder" />
        </el-form-item>
        <el-form-item
          prop="class_names.value"
        >
          <span slot="label">
            {{ form.class_names.label.title }}
            <el-tooltip placement="right">
              <div slot="content">
                {{ form.class_names.label.desc }}
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input v-model="form.class_names.value" :placeholder="form.class_names.placeholder" />
        </el-form-item>
        <el-form-item prop="file.value">
          <span slot="label">
            {{ form.file.label.title }}
            <el-tooltip placement="right">
              <div slot="content">
                {{ form.file.label.desc }}
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <input type="file" @change="triggerFile($event)"></input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="onSubmit('form')">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog
      title="提示"
      :visible.sync="dialogVisible2Delete"
      width="30%"
    >
      <span>删除以后AI模型无法恢复，确定删除？</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible2Delete = false">取 消</el-button>
        <el-button type="primary" @click="deleteModel">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog
      title="模型下发"
      :visible.sync="dialogVisible2distribute"
      width="30%"
    >
      <distribute-ai-model-dialog/>
    </el-dialog>
  </div>
</template>
<script>
  import { fetchList, create } from '@/api/ai-model'
  import { Message } from 'element-ui'
  import DistributeAiModelDialog from './components/distribute_aimodel_dialog'

  export default {
    name: 'AiModel',
    components: { DistributeAiModelDialog },
    data() {
      return {
        dialogVisible2distribute: false,
        dialogVisible2Delete: false,
        dialogVisible: false,
        tableData: null,
        form: {
          class_names: {
            label: {
              title: '分类信息',
              desc: '分类神经网络分类信息'
            },
            value: '',
            placeholder: '请输入分类信息'
          },
          version: {
            label: {
              title: '模型版本',
              desc: 'AI模型版本号'
            },
            value: '',
            placeholder: '请输入AI模型版本号'
          },
          file: {
            label: {
              title: '模型文件',
              desc: 'AI模型文件'
            },
            value: ''
          }
        },
        rules: {
          'class_names.value': [
            { required: true, message: '请输入模型分类信息', trigger: 'blur' }
          ],
          'version.value': [
            { required: true, message: '请输入模型版本号', trigger: 'blur' }
          ],
          'file.value': [
            { required: true, message: '请上传模型' }
          ]
        }
      }
    },
    created() {
      this.getList()
    },
    methods: {
      showDeleteModel(version) {
        this.$store.dispatch('edge-aimodel/setAiModelVersion', version)
        this.dialogVisible2Delete = true
      },
      deleteModel() {
        this.$store.dispatch('edge-aimodel/deleteAiModel').then(
          _ => {
            Message({
              message: '删除AI模型成功',
              type: 'success',
              duration: 5 * 1000
            })
            this.dialogVisible2Delete = false
            this.getList()
          }
        )
      },
      distributeModel(version) {
        this.$store.dispatch('edge-aimodel/setAiModelVersion', version)
        this.dialogVisible2distribute = true
      },
      handleClose(done) {
        this.$confirm('确认关闭？')
          .then(_ => {
            done()
          })
          .catch(_ => {
          })
      },
      triggerFile(event) {
        this.form.file.value = event.target.files[0]
      },
      onSubmit(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            let formData = new FormData()
            formData.append('file', this.form.file.value)
            formData.append('version', this.form.version.value)
            formData.append('class_names', this.form.class_names.value)
            return new Promise((resolve, reject) => {
              create(
                formData
              ).then(response => {
                Message({
                  message: '创建成功',
                  type: 'success',
                  duration: 5 * 1000
                })
                resolve()
                //关闭窗口，刷新页面
                this.dialogVisible = false
                this.getList()
              }).catch(error => {
                reject(error)
              })
            })
          } else {
            return false
          }
        })
      },
      getList() {
        this.$store.dispatch('edge-aimodel/getAiModelList').then(response => {
          this.tableData = response
        })
      }
    }
  }
</script>
<style scoped>
</style>
