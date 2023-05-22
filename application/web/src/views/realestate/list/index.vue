<template>
<div>
  <div v-show="userName=='客户模拟账户'" class="container">
    <div style="margin-top: 15px;" >
      <el-input v-model="input" placeholder="请输入农产品溯源编码" style="width: 400px" >
        <el-button slot="append" icon="el-icon-search" v-on:click="searchtext = input" >溯源</el-button>
      </el-input>
    </div>
      <el-row v-loading="loading" :gutter="20">
        <el-col v-for="(val,index) in realEstateList" :key="index" :span="6" :offset="1">
          <!-- <el-card class="realEstate-card"> -->
          <el-card class="text item" v-show="searchtext == val.realEstateId">
            <div class="item">
              <el-tag>溯源码: </el-tag>
              <span>{{ val.realEstateId }}</span>
            </div>          
            <div class="item">
              <el-tag type="success">原料供应商ID: </el-tag>
              <span>{{ val.accountId }}</span>
            </div>
            <div class="item">
              <el-tag type="success">烟草原料产地: </el-tag>
              <span>{{ val.yzcxx }}</span>
            </div>
            <div class="item">
              <el-tag type="success">烟草原料品质（优中良）: </el-tag>
              <span>{{ val.yzcwz }}</span>
            </div>
            <div class="item">
              <el-tag type="success">烟草原料生长周期: </el-tag>
              <span>{{ val.yzcdh }}</span>
            </div>
            <div class="item1">
              <el-tag type="success">烟草原料采集时间: </el-tag>
              <span>{{ val.yzcnzebh }}</span>
            </div>
            <div class="item">
              <el-tag type="success">烤烟等级: </el-tag>
              <span>{{ val.yzctz }}</span>
            </div>
            <div class="item">
              <el-tag type="success">烤烟温度: </el-tag>
              <span>{{ val.yzcrcsj }}</span>
            </div>
            <div class="item">
              <el-tag type="success">烘干温湿度: </el-tag>
              <span>{{ val.yzcclsj }}</span>
            </div>
            <div class="item">
              <el-tag type="success">烹饪状态: </el-tag>
              <span>{{ val.yzcymjzcs }}</span>
            </div>
            <div class="item">
              <el-tag type="success">工业公司ID: </el-tag>
              <span>{{ val.pfsaccountid }}</span>
            </div>
            <div class="item">
              <el-tag type="warning">仓库等级和编号（ABC）: </el-tag>
              <span>{{ val.tzcmc }}</span>
            </div>
            <div class="item">
              <el-tag type="warning">烟叶年份: </el-tag>
              <span>{{ val.tzcdh }}</span>
            </div>
            <div class="item">
              <el-tag type="warning">采购时间: </el-tag>
              <span>{{ val.tzcnzebh }}</span>
            </div>
            <div class="item">
              <el-tag type="warning">醇化后烟叶等级: </el-tag>
              <span>{{ val.tzccpbh }}</span>
            </div>   
            <div class="item">
              <el-tag type="warning">复烤温湿度: </el-tag>
              <span>{{ val.tzcnrbw }}</span>
            </div>
            <div class="item">
              <el-tag type="warning">醇化详细信息: </el-tag>
              <span>{{ val.tzczl }}</span>
            </div>
            <div class="item">
              <el-tag type="warning">醇化时间: </el-tag>
              <span>{{ val.tzctzsj }}</span>
            </div>                
            <div class="item">
              <el-tag type="warning">烟叶回潮信息: </el-tag>
              <span>{{ val.tzcczyxm }}</span>
            </div>
            <div class="item">
              <el-tag type="warning">制丝工艺: </el-tag>
              <span>{{ val.zhisigongyi }}</span>
            </div>            
            <div class="item">
              <el-tag type="warning">加工时间: </el-tag>
              <span>{{ val.jiagongsj }}</span>
            </div>
            <div class="item">
              <el-tag type="warning">成品品牌: </el-tag>
              <span>{{ val.chengpinpp }}</span>
            </div>
            <div class="item">
              <el-tag type="warning">成品规格: </el-tag>
              <span>{{ val.chengpingg }}</span>
            </div>
            <div class="item">
              <el-tag type="warning">出场运输车车型: </el-tag>
              <span>{{ val.cysccx }}</span>
            </div>
            <div class="item">
              <el-tag type="warning">出场运输车车牌号: </el-tag>
              <span>{{ val.ccysccph }}</span>
            </div>
            <div class="item">
              <el-tag type="warning">运输时间: </el-tag>
              <span>{{ val.ccyssj }}</span>
            </div>
            <div class="item">
              <el-tag type="danger">商业公司ID: </el-tag>
              <span>{{ val.sygsaccid }}</span>
            </div>       
            <div class="item">
              <el-tag type="danger">成品采购时间: </el-tag>
              <span>{{ val.yuqywlgsmc }}</span>
            </div>
            <div class="item">
              <el-tag type="danger">成品采购规格: </el-tag>
              <span>{{ val.yuqycph }}</span>
            </div>
            <div class="item">
              <el-tag type="danger">成品配送车型号: </el-tag>
              <span>{{ val.yuqyysxx }}</span>
            </div>
            <div class="item">
              <el-tag type="danger">成品配送车车牌号: </el-tag>
              <span>{{ val.syysccp }}</span>
            </div>
            <div class="item">
              <el-tag type="danger">成品配送时间: </el-tag>
              <span>{{ val.syyysj }}</span>
            </div>
            <div class="item">
              <el-tag type="success">零售店: </el-tag>
              <span>{{ val.xsqymc }}</span>
            </div>
                 
            <div v-if="!val.encumbrance&&roles[0] !== 'admin'">
            </div>
            <el-rate v-if="roles[0] === 'admin'" />
          </el-card>
        </el-col>
      </el-row>










  </div>

  <div v-show="userName!=='客户模拟账户'">
  <div class="container">
    <el-alert
      type="success"
    >
      <p>账户ID: {{ accountId }}</p>
      <p>当前角色: {{ userName }}</p>
    </el-alert>
    <div v-if="realEstateList.length==0" style="text-align: center;">
      <el-alert
        title="查询不到数据"
        type="warning"
      />
    </div>
    <el-table
      :data="tableData"
      border
      style="width: 100%">
      <div>
      <el-table-column
        fixed
        prop="realEstateId"
        label="溯源码"
        width="150">
      </el-table-column>
      <el-table-column
        prop="accountId"
        label="原料供应商ID"
        width="150">
      </el-table-column>
      <el-table-column
        prop="yzcxx"
        label="烟草原料产地"
        width="120">
      </el-table-column>
      <el-table-column
        prop="yzcwz"
        label="烟草原料品质（优中良）"
        width="120">
      </el-table-column>
      <el-table-column
        prop="yzcdh"
        label="烟草原料生长周期"
        width="140">
      </el-table-column>
      <el-table-column
        prop="yzcnzebh"
        label="烟草原料采集时间"
        width="140">
      </el-table-column>
      <el-table-column
        prop="yzctz"
        label="烤烟等级"
        width="120">
      </el-table-column>
      <el-table-column
        prop="yzcrcsj"
        label="烤烟温度"
        width="120">
      </el-table-column>
      <el-table-column
        prop="yzcclsj"
        label="烘干温湿度"
        width="120">
      </el-table-column>
      <el-table-column
        prop="yzcymjzcs"
        label="烹饪状态"
        width="120">
      </el-table-column>
      <el-table-column
        prop="pfsaccountid"
        label="工业公司id"
        width="120">
      </el-table-column>
      <el-table-column
        prop="tzcmc"
        label="仓库等级和编号（ABC）"
        width="120">
      </el-table-column>
      <el-table-column
        prop="tzcdh"
        label="烟叶年份"
        width="120">
      </el-table-column>
      <el-table-column
        prop="tzcnzebh"
        label="采购时间"
        width="120">
      </el-table-column>
      <el-table-column
        prop="tzccpbh"
        label="醇化后烟叶等级"
        width="120">
      </el-table-column>
      <el-table-column
        prop="tzcnrbw"
        label="复烤温湿度"
        width="120">
      </el-table-column>
      <el-table-column
        prop="tzczl"
        label="醇化详细信息"
        width="120">
      </el-table-column>
      <el-table-column
        prop="tzctzsj"
        label="醇化时间"
        width="120">
      </el-table-column>
      <el-table-column
        prop="tzcczyxm"
        label="烟叶回潮信息"
        width="120">
      </el-table-column>
      <el-table-column
        prop="zhisigongyi"
        label="制丝工艺"
        width="120">
      </el-table-column>
      <el-table-column
        prop="jiagongsj"
        label="加工时间"
        width="120">
      </el-table-column>
      <el-table-column
        prop="chengpinpp"
        label="成品品牌"
        width="120">
      </el-table-column>
      <el-table-column
        prop="chengpingg"
        label="成品规格"
        width="120">
      </el-table-column>
      <el-table-column
        prop="cysccx"
        label="出场运输车车型"
        width="120">
      </el-table-column>
      <el-table-column
        prop="ccysccph"
        label="出场运输车车牌号"
        width="150">
      </el-table-column>
      <el-table-column
        prop="ccyssj"
        label="运输时间"
        width="120">
      </el-table-column>
      <el-table-column
        prop="sygsaccid"
        label="商业公司ID"
        width="120">
      </el-table-column>
      <el-table-column
        prop="yuqywlgsmc"
        label="成品采购时间"
        width="130">
      </el-table-column>
      <el-table-column
        prop="yuqycph"
        label="成品采购规格"
        width="120">
      </el-table-column>
      <el-table-column
        prop="yuqyysxx"
        label="成品配送车型号"
        width="150">
      </el-table-column>
      <el-table-column
        prop="syysccp"
        label="成品配送车车牌号"
        width="150">
      </el-table-column>
      <el-table-column
        prop="syyysj"
        label="成品配送时间"
        width="150">
      </el-table-column>
      <el-table-column
        prop="xsqymc"
        label="零售店"
        width="120">
      </el-table-column>
      <el-table-column
        fixed="right"
        label="操作"
        width="100">
        <template slot-scope="scope">
          <el-button @click="handleUpdate(scope.row)" type="text" size="small" v-show="userName!=='原料供应商1' && userName!=='原料供应商2' && userName!=='原料供应商3'" >添加信息</el-button>
          <!-- <el-button @click="handleUpdate(scope.row)" type="text" size="small" v-show ="scope.row.sygsaccid=='' && userName == '烟草商业公司'" >添加信息</el-button>
          <el-button @click="handleUpdate(scope.row)" type="text" size="small" v-show ="scope.row.xsqymc=='' && userName == '销售企业模拟账户'" >添加信息</el-button> -->
          <!-- <el-button @click="handleUpdate(scope.row)" type="text" size="small" v-show ="scope.row.tzcmc!=='' && userName == '客户模拟账户'" >添加信息</el-button> -->

        </template>
      </el-table-column>
      </div>
    </el-table>


    <el-dialog title="添加信息" :visible.sync="dialogFormVisible">
      <el-form ref="dataForm" :model="temp" label-position="left" label-width="70px" style="width: 400px; margin-left:50px;">
        <div v-if="userName == '烟草工业公司' && rnstatus=='已出栏'">
          
        <el-form-item  label="仓库等级和编号（ABC）" prop="title">
          <el-input v-model="temp.tzcmc" />
        </el-form-item>
        <el-form-item  label="烟叶年份" prop="title">
          <el-input v-model="temp.tzcdh" />
        </el-form-item>
        <el-form-item  label="采购时间" prop="title">
          <el-input v-model="temp.tzccpbh" />
        </el-form-item>
        <el-form-item  label="醇化后烟叶等级" prop="title">
          <el-input v-model="temp.tzcnzebh" />
        </el-form-item>
        <el-form-item  label="复烤温湿度" prop="title" >
          <el-input v-model="temp.tzcnrbw" />
        </el-form-item>
        <el-form-item  label="醇化详细信息" prop="title">
          <el-input v-model="temp.tzczl" />
        </el-form-item>
        <el-form-item  label="醇化时间" prop="title">
          <el-input v-model="temp.tzctzsj" />
        </el-form-item>
        <el-form-item  label="烟叶回潮信息" prop="title" >
          <el-input v-model="temp.tzcczyxm" />
        </el-form-item>
        <el-form-item  label="制丝工艺" prop="title" >
          <el-input v-model="temp.zhisigongyi" />
        </el-form-item>
        <el-form-item  label="加工时间" prop="title" >
          <el-input v-model="temp.jiagongsj" />
        </el-form-item>
        <el-form-item  label="成品品牌" prop="title" >
          <el-input v-model="temp.chengpinpp" />
        </el-form-item>
        <el-form-item  label="成品规格" prop="title" >
          <el-input v-model="temp.chengpingg" />
        </el-form-item>
        <el-form-item  label="出场运输车车型" prop="title" >
          <el-input v-model="temp.cysccx" />
        </el-form-item>
        <el-form-item  label="出场运输车车牌号" prop="title" >
          <el-input v-model="temp.ccysccph" />
        </el-form-item>
        <el-form-item  label="运输时间" prop="title" >
          <el-input v-model="temp.ccyssj" />
        </el-form-item>
        </div>
        
        
        
        <div v-if="(userName == '烟草商业公司1' || userName == '烟草商业公司2' || userName == '烟草商业公司3') && rnstatus=='等待运输'">
        <el-form-item label="成品采购时间" prop="title" >
          <el-input v-model="temp.yuqywlgsmc" />
        </el-form-item>
        <el-form-item label="成品采购规格" prop="title" >
          <el-input v-model="temp.yuqycph" />
        </el-form-item>
        <el-form-item label="成品配送车型号" prop="title" >
          <el-input v-model="temp.yuqyysxx" />
        </el-form-item>
        <el-form-item label="成品配送车车牌号" prop="title" >
          <el-input v-model="temp.syysccp" />
        </el-form-item>
        <el-form-item label="成品配送时间" prop="title" >
          <el-input v-model="temp.syyysj" />
        </el-form-item>
        </div>
        
        
        <div  v-if="userName == '零售店模拟账户' && rnstatus=='等待出售' ">
        <el-form-item label="销售企业名称" prop="title">
          <el-input v-model="temp.xsqymc" />
        </el-form-item>
        </div>


      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">
          取消
        </el-button>
        <!-- <el-button type="primary" @click="submitdata()"> -->
        <!-- <el-button type="primary" @click="submitForm('temp')"> -->
        <el-button type="primary" @click="confirmnzebh()">

          
          确认
        </el-button>

      </div>
    </el-dialog>    



    





  </div>
  </div>
</div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryAccountList } from '@/api/account'
import { queryRealEstateList } from '@/api/realEstate'
import { createSelling } from '@/api/selling'
import { createDonating } from '@/api/donating'
import { createRealEstate } from '@/api/realEstate'

export default {
  // username=userName,
  rnstatus:'',
  name: 'RealeState',
  data() {
    var checkArea = (rule, value, callback) => {
      if (value <= 0) {
        callback(new Error('必须大于0'))
      } else {
        callback()
      }
    }
    return {

      temp: [{
        RealEstateID: '',
        proprietor: '',
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
        tzcnzebh:''
      }],
      dialogFormVisible:false,
      tableData: [],
      searchall:'',
      searchtext:'',
      input: '',
      all:'',
      loading: true,
      loadingDialog: false,
      realEstateList: [],
      dialogCreateSelling: false,
      dialogCreateDonating: false,
      realForm: {
        price: 0,
        salePeriod: 0
      },
      rules: {
        price: [
          { validator: checkArea, trigger: 'blur' }
        ],
        salePeriod: [
          { validator: checkArea, trigger: 'blur' }
        ]
      },
      DonatingForm: {
        proprietor: ''
      },
      rulesDonating: {
        proprietor: [
          { required: true, message: '请选择客户', trigger: 'change' }
        ]
      },
      accountList: [],
      valItem: {}
    }
  },
  computed: {
    ...mapGetters([
      'accountId',
      'roles',
      'userName',
      'balance'
    ])
  },
  created() {
    if (true) {
      queryRealEstateList().then(response => {
        if (response !== null) {
          this.realEstateList = response
          this.tableData = response
          console.log(response,'response')
          console.log(this.tableData)
          console.log(this.realEstateList)
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    } else {
      queryRealEstateList({ proprietor: this.accountId }).then(response => {
        if (response !== null) {
          this.realEstateList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    }
  },
  methods: {
    created2() {
      if (true) {
        queryRealEstateList().then(response => {
  	  if (response !== null) {
	    this.tableData = response
	    console.log(response,'response')
	  }
	  this.loading = false
        }).catch(_ => {
	  this.loading = false
        })
      } 
    },

    handleUpdate(row) {

      this.temp = Object.assign({}, row) // copy obj
      console.log(this.temp)
      if(this.temp.xsqymc!==''){
        this.rnstatus = "已售出";
        console.log(this.rnstatus)

      }else if(this.temp.yuqyysxx!==''){
        this.rnstatus = "等待出售";
        console.log(this.rnstatus)

      }else if(this.temp.tzcczyxm!==''){
        this.rnstatus = "等待运输";
        console.log(this.rnstatus)

      }else{
        this.rnstatus = "已出栏";
        console.log(this.rnstatus)

      }
      
      this.dialogFormVisible = true



    },
    openDialog(item) {
      this.dialogCreateSelling = true
      this.valItem = item
    },
    openDonatingDialog(item) {
      this.dialogCreateDonating = true
      this.valItem = item
      queryAccountList().then(response => {
        if (response !== null) {
          // 过滤掉管理员和当前用户
          this.accountList = response.filter(item =>
            item.userName !== '管理员' && item.accountId !== this.accountId
          )
        }
      })
    },
    createSelling(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm('是否立即出售?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'success'
          }).then(() => {
            this.loadingDialog = true
            createSelling({
              objectOfSale: this.valItem.realEstateId,
              seller: this.valItem.proprietor,
              price: this.realForm.price,
              salePeriod: this.realForm.salePeriod
            }).then(response => {
              this.loadingDialog = false
              this.dialogCreateSelling = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: '出售成功!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: '出售失败!'
                })
              }
              setTimeout(() => {
                window.location.reload()
              }, 1000)
            }).catch(_ => {
              this.loadingDialog = false
              this.dialogCreateSelling = false
            })
          }).catch(() => {
            this.loadingDialog = false
            this.dialogCreateSelling = false
            this.$message({
              type: 'info',
              message: '已取消出售'
            })
          })
        } else {
          return false
        }
      })
    },
    onload(){
    

    },
 
  

    

    
    confirmnzebh(){
      this.submitForm('temp')

    },

    submitForm(formName) {
      this.$confirm('是否立即添加信息?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'success'
      }).then(() => {
      this.loading = true
      if (this.userName == '烟草工业公司') {
      	this.temp.pfsaccountid = this.accountId
      } else if (this.userName == '烟草商业公司1' || this.userName == '烟草商业公司2' || this.userName == '烟草商业公司3') {
      	this.temp.sygsaccid = this.accountId
      }
      createRealEstate({
        realEstateId:this.temp.realEstateId,
        accountId: this.temp.accountId,
        proprietor: this.temp.proprietor,
        yzcxx: this.temp.yzcxx,
        yzcwz: this.temp.yzcwz,
        yzcdh: this.temp.yzcdh,
        yzcnzebh: this.temp.yzcnzebh,
        yzctz: this.temp.yzctz,
        yzcrcsj: this.temp.yzcrcsj,
        yzcclsj: this.temp.yzcclsj,
        yzcymjzcs: this.temp.yzcymjzcs,
        pfsaccountid: this.temp.pfsaccountid,
        tzcmc: this.temp.tzcmc,
        tzcdh: this.temp.tzcdh,
        tzccpbh: this.temp.tzccpbh,
        tzcnrbw: this.temp.tzcnrbw,
        tzczl: this.temp.tzczl,
        tzctzsj: this.temp.tzctzsj,
        tzcczyxm: this.temp.tzcczyxm,
        zhisigongyi: this.temp.zhisigongyi,
        jiagongsj: this.temp.jiagongsj,
        chengpinpp: this.temp.chengpinpp,
        chengpingg: this.temp.chengpingg,
        cysccx: this.temp.cysccx,
        ccysccph: this.temp.ccysccph,
        ccyssj: this.temp.ccyssj,
        sygsaccid: this.temp.sygsaccid,
        yuqywlgsmc: this.temp.yuqywlgsmc,
        yuqycph: this.temp.yuqycph,
        yuqyysxx: this.temp.yuqyysxx,
        syysccp: this.temp.syysccp,
        syyysj: this.temp.syyysj,
        xsqymc: this.temp.xsqymc,
        tzcnzebh:this.temp.tzcnzebh,
      }).then(response => {
        this.loading = false

        if (response !== null) {
          this.sleep(2000)
          queryRealEstateList().then(response => {
            if (response !== null) {
              this.tableData = response
              console.log(response,'responsenew')
            }
            this.loading = false
          }).catch(_ => {
            this.loading = false
          })
          this.$message({
            type: 'success',
            message: '添加信息成功!'
          })
        } else {
          this.$message({
            type: 'error',
            message: '添加信息失败!'
          })
        }
      }).catch(_ => {
        this.loading = false
      })
    }).catch(() => {
      this.loading = false
      this.$message({
        type: 'info',
        message: '已取消添加信息'
      })
    })
    this.dialogFormVisible = false
    },
    
    sleep(ms) { //sleep延迟方法
            var unixtime_ms = new Date().getTime();
            while(new Date().getTime() < unixtime_ms + ms) {}
    },


    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    createDonating(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm('是否立即捐赠?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'success'
          }).then(() => {
            this.loadingDialog = true
            createDonating({
              objectOfDonating: this.valItem.realEstateId,
              donor: this.valItem.proprietor,
              grantee: this.DonatingForm.proprietor
            }).then(response => {
              this.loadingDialog = false
              this.dialogCreateDonating = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: '捐赠成功!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: '捐赠失败!'
                })
              }
              setTimeout(() => {
                window.location.reload()
              }, 1000)
            }).catch(_ => {
              this.loadingDialog = false
              this.dialogCreateDonating = false
            })
          }).catch(() => {
            this.loadingDialog = false
            this.dialogCreateDonating = false
            this.$message({
              type: 'info',
              message: '已取消捐赠'
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
      this.DonatingForm.proprietor = accountId
    }
  }

  }


</script>

<style>
  .container{
    width: 100%;
    text-align: center;
    min-height: 100%;
    overflow: hidden;
  }
  .tag {
    float: left;
  }

  .item {
    font-size: 14px;
    margin-bottom: 18px;
    color: #999;
  }

  .item1 {
    font-size: 14px;
    margin-bottom: 22px;
    color: #999;
  }

  .clearfix:before,
  .clearfix:after {
    display: table;
  }
  .clearfix:after {
    clear: both
  }

  .realEstate-card {
    width: 280px;
    height: 600px;
    margin: 18px;
    text-align: center;
  }
</style>
