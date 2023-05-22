<template>
  <div class="app-container">
    <el-form ref="ruleForm" v-loading="loading" :model="ruleForm" :rules="rules" label-width="100px">

      <!-- <el-form-item label="客户" prop="proprietor">
        <el-select v-model="ruleForm.proprietor" placeholder="请选择客户" @change="selectGet">
          <el-option
            v-for="item in accountList"
            :key="item.accountId"
            :label="item.userName"
            :value="item.accountId"
          >
            <span style="float: left">{{ item.userName }}</span>
            <span style="float: right; color: #8492a6; font-size: 13px">{{ item.accountId }}</span>
          </el-option>
        </el-select>
      </el-form-item> -->


      <!-- <el-form-item label="总空间 ㎡" prop="totalArea">
        <el-input-number v-model="ruleForm.totalArea" :precision="2" :step="0.1" :min="0" />
      </el-form-item> -->
      <el-form-item label="烟草原料产地" prop="totalArea">
        <el-input size="medium"  v-model="ruleForm.yzcxx" placeholder="请输入内容" style="width: 300px"></el-input>
      </el-form-item>
      <el-form-item label="烟草原料品质（优中良）" prop="totalArea">
        <el-input size="medium"  v-model="ruleForm.yzcwz" placeholder="请输入内容" style="width: 300px"></el-input>
      </el-form-item>
      <el-form-item label="烟草原料生长周期" prop="totalArea">
        <el-input size="medium"  v-model="ruleForm.yzcdh" placeholder="请输入内容" style="width: 300px"></el-input>
      </el-form-item>
      <el-form-item label="烟草原料采集时间" prop="totalArea">
        <el-input size="medium"  v-model="ruleForm.yzcnzebh" placeholder="请输入内容" style="width: 300px"></el-input>
      </el-form-item>
      <el-form-item label="烤烟等级" prop="totalArea">
        <el-input size="medium"  v-model="ruleForm.yzctz" placeholder="请输入内容" style="width: 300px"></el-input>
      </el-form-item>
      <el-form-item label="烤烟温度" prop="totalArea">
        <el-input size="medium"  v-model="ruleForm.yzcrcsj" placeholder="请输入内容" style="width: 300px"></el-input>
      </el-form-item>
      <el-form-item label="烘干温湿度" prop="totalArea">
        <el-input size="medium"  v-model="ruleForm.yzcclsj" placeholder="请输入内容" style="width: 300px"></el-input>
      </el-form-item>
      <el-form-item label="烹饪状态" prop="totalArea">
        <el-input size="medium"  v-model="ruleForm.yzcymjzcs" placeholder="请输入内容" style="width: 300px"></el-input>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')">立即创建</el-button>
        <el-button @click="resetForm('ruleForm')">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryAccountList } from '@/api/account'
import { createRealEstate } from '@/api/realEstate'

export default {
  name: 'AddRealeState',
  data() {
    var checkArea = (rule, value, callback) => {
      if (value <= 0) {
        callback(new Error('必须大于0'))
      } else {
        callback()
      }
    }
    return {
      ruleForm: {
        proprietor: '4e07408562be',
        yzcxx: '',
        yzcwz: '',
        yzcdh: '',
        yzcnzebh: '',
        yzctz: '',
        yzcrcsj: '',
        yzcclsj: '',
        yzcymjzcs: '',
        pfsaccountid: '',
        tzcmc: '',
        tzcdh: '',
        tzcnzebh: '',
        tzccpbh: '',
        tzcnrbw: '',
        tzczl: '',
        tzctzsj: '',
        tzcczyxm: '',
        zhisigongyi: '',
        jiagongsj: '',
        chengpinpp: '',
        chengpingg: '',
        cysccx: '',
        ccysccph: '',
        ccyssj: '',
        sygsaccid: '',
        yuqywlgsmc: '',
        yuqycph: '',
        yuqyysxx: '',
        syysccp: '',
        syyysj: '',
        xsqymc: '',
      },
      accountList: [],
      rules: {
        proprietor: [
          { required: true, message: '请选择客户', trigger: 'change' }
        ],
        // totalArea: [
        //   { validator: checkArea, trigger: 'blur' }
        // ],
        // livingSpace: [
        //   { validator: checkArea, trigger: 'blur' }
        // ]
      },
      loading: false
    }
  },
  computed: {
    ...mapGetters([
      'accountId'
    ])
  },
  created() {
    queryAccountList().then(response => {
      if (response !== null) {
        // 过滤掉管理员
        this.accountList = response.filter(item =>
          item.userName !== '管理员'
        )
      }
    })
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm('是否立即创建?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'success'
          }).then(() => {
            this.loading = true
            var time = Date.now().toString();
            createRealEstate({
              realEstateId:time,
              accountId: this.accountId,
              proprietor: this.ruleForm.proprietor,
              yzcxx: this.ruleForm.yzcxx,
              yzcwz: this.ruleForm.yzcwz,
              yzcdh: this.ruleForm.yzcdh,
              yzcnzebh: this.ruleForm.yzcnzebh,
              yzctz: this.ruleForm.yzctz,
              yzcrcsj: this.ruleForm.yzcrcsj,
              yzcclsj: this.ruleForm.yzcclsj,
              yzcymjzcs: this.ruleForm.yzcymjzcs,
              pfsaccountid: this.ruleForm.pfsaccountid,
              tzcmc: this.ruleForm.tzcmc,
              tzcdh: this.ruleForm.tzcdh,
              tzcnzebh: this.ruleForm.tzcnzebh,
              tzccpbh: this.ruleForm.tzccpbh,
              tzcnrbw: this.ruleForm.tzcnrbw,
              tzczl: this.ruleForm.tzczl,
              tzctzsj: this.ruleForm.tzctzsj,
              tzcczyxm: this.ruleForm.tzcczyxm,
              zhisigongyi: this.ruleForm.zhisigongyi,
              jiagongsj: this.ruleForm.jiagongsj,
              chengpinpp: this.ruleForm.chengpinpp,
              chengpingg: this.ruleForm.chengpingg,
              cysccx: this.ruleForm.cysccx,
              ccysccph: this.ruleForm.ccysccph,
              ccyssj: this.ruleForm.ccyssj,
              sygsaccid: this.ruleForm.sygsaccid,
              yuqywlgsmc: this.ruleForm.yuqywlgsmc,
              yuqycph: this.ruleForm.yuqycph,
              yuqyysxx: this.ruleForm.yuqyysxx,
              syysccp: this.ruleForm.syysccp,
              syyysj: this.ruleForm.syyysj,
              xsqymc: this.ruleForm.xsqymc,
            }).then(response => {
              this.loading = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: '创建成功!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: '创建失败!'
                })
              }
            }).catch(_ => {
              this.loading = false
            })
          }).catch(() => {
            this.loading = false
            this.$message({
              type: 'info',
              message: '已取消创建'
            })
          })
        } else {
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    selectGet(accountId) {
      this.ruleForm.proprietor = accountId
    }
  }
}
</script>

<style scoped>
</style>
