<template>
  <div class="card-container">
    <div class="card-header">
      <div class="title">节点标签</div>
      <div class="edit" @click="dialogVisible=true">
        <svg-icon icon-class="edit" />
      </div>
    </div>
    <div class="card-body">
      <el-tag class="my-tag" v-for="(item,key) in labelList">{{item['key']}}:{{item['value']}}</el-tag>
    </div>
    <el-dialog
      title="节点标签管理"
      :visible.sync="dialogVisible"
      width="35%"
    >
      <el-row v-for="(item,key) in labelList" :key="key">
        <el-col :span="12">
          <el-row style="text-align: center">
            <el-col :span="6" style="line-height: 30px">
              标签键：
            </el-col>
            <el-col :span="18">
              <el-input :disabled="item['disabled']" v-model="item['key']" />
            </el-col>
          </el-row>
        </el-col>
        <el-col :span="8" style="padding-left: 10px;text-align: center">
          <el-row>
            <el-col :span="4" style="line-height: 30px">
              值：
            </el-col>
            <el-col :span="20">
              <el-input v-model="item['value']" />
            </el-col>
          </el-row>
        </el-col>
        <el-col :span="4" style="text-align:center;line-height: 40px">
          <i v-if="item['status']=='add'" class="el-icon-check" @click="addLabel(key)" />
          <i v-else class="el-icon-close" @click="deleteLabel(key)" />
        </el-col>
      </el-row>
      <div style="margin: 10px 35px">
        <el-button type="primary" icon="el-icon-circle-plus" @click="addViewLabel()">添加标签</el-button>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">关 闭</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
  import { deleteLabel } from '@/api/edge-node'
  import { Message } from 'element-ui'

  export default {
    name: 'LabelsCard',
    props: ['labels'],
    data() {
      return {
        dialogVisible: false,
        labelList: []
      }
    },
    watch: {
      'labels'(val) {
        for (let key in val) {
          this.labelList.push(
            {
              'disabled': true,
              'key': key,
              'value': val[key],
              'status': 'del'
            }
          )
        }
      }
    },
    methods: {
      addLabel(index) {
        let label = this.labelList[index]
        //校验是否已经包含key
        if (this.labels.hasOwnProperty(label['key'])) {
          Message({
            message: '标签键已存在',
            type: 'error',
            duration: 5 * 1000
          })
          return
        }
        if (label['key'] === '') {
          Message({
            message: '标签键不能为空',
            type: 'error',
            duration: 5 * 1000
          })
          return
        }
        let params = {
          'label_key': label['key'],
          'label_value': label['value']
        }
        this.$store.dispatch('edge-node/createNodeLabel', params)
        //改变状态
        this.labelList[index]['disabled'] = 'true'
        this.labelList[index]['status'] = 'del'
        Message({
          message: '标签添加成功！',
          type: 'success',
          duration: 5 * 1000
        })
      },
      addViewLabel() {
        this.labelList.push(
          {
            'disabled': false,
            'key': '',
            'value': '',
            'status': 'add'
          }
        )
      },
      deleteLabel(index) {
        let label = this.labelList[index]
        //后台接口请求
        let params = {
          'label_key': label['key']
        }
        this.$store.dispatch('edge-node/deleteNodeLabel', params)
        //改变数据界面变化
        this.labelList.splice(index, 1)
        Message({
          message: '标签删除成功！',
          type: 'success',
          duration: 5 * 1000
        })
      }
    }
  }
</script>
<style scoped>
  .my-tag {
    margin: 5px;
    font-size: 16px;
  }

  .el-icon-close, .el-icon-check {
    font-size: 20px;
    cursor: pointer;
  }

  .el-icon-close:hover,.el-icon-check:hover {
    color: #409EFF;
  }

  .card-header > * {
    display: inline-block;
  }

  .edit {
    margin-left: 10px;
    cursor: pointer
  }
</style>
