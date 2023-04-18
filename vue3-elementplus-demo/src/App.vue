<template>
  <div class="table-demo">
    <!-- 标题 -->
    <div class="title">
      <h2> CURD </h2>
      <p> {{ initialText }} </p>
    </div>
    
    <!-- 功能 -->
    <div class="qurey-box">
      <!-- <el-input v-model="query_input" placeholder="请输入:" @click="handelQureyName" /> -->
      <el-input class="query-input" v-model="query_input" placeholder="请输入姓名搜索🔍" @change="handleQueryName"/>
      <div class="btn-list">
        <el-button type="primary" @click="handleadd">新增</el-button>
        <el-button type="danger" @click="handledel" v-if="multipleSelection.length > 0">删除选中</el-button>
      </div>
    </div>

    <!-- table -->
    <div class="table-box">
      <el-table 
      ref="multipleTableRef"  
      :data="tableData" 
      style="width: 100%" 
      border
      @selection-change="handleSelectionChange"
    >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="name" label="名字" width="120" />
        <el-table-column prop="phone" label="电话" width="200" />
        <el-table-column prop="email" label="邮箱" width="300" />
        <el-table-column prop="status" label="状态" width="120" />
        <el-table-column prop="city" label="城市" width="120" />
        <el-table-column fixed="right" label="操作" width="120">
          <template #default="scope">
            <!-- 方式一 -->
            <!-- <el-button link type="primary" size="small" @click.prevent="deleteRow(scope.$index)" style="color: #F56C6C">删除</el-button> -->
            
            <!-- 方式二 -->
            <el-button link type="primary" size="small" @click="deleteRow(scope.row)" style="color: #F56C6C">删除</el-button>
            <el-button link type="primary" size="small" @click="handleEdit(scope.row)">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination 
          background 
          layout="prev, pager, next" 
          :total="total" 
          v-model:current-page="curPage"
          @current-change="handleChangePage"
        />
      </div>
      

      <!-- 弹窗 -->
      <el-dialog v-model="dialogFormVisible" :title="dialogFormTitle === 'add' ? '新增':'编辑'">
        <el-form :model="dialogdata">
          <el-form-item fixed label="请输入姓名: " :label-width="120">
            <el-input v-model="dialogdata.name" autocomplete="off" />
          </el-form-item>
          <el-form-item label="请输入电话: " :label-width="120">
            <el-input v-model="dialogdata.phone" autocomplete="off" />
          </el-form-item>
          <el-form-item label="请输入邮箱: " :label-width="120">
            <el-input v-model="dialogdata.email" autocomplete="off" />
          </el-form-item>
          <el-form-item label="请输入状态: " :label-width="120">
            <el-input v-model="dialogdata.status" autocomplete="off" />
          </el-form-item>
          <el-form-item label="请输入城市: " :label-width="120">
            <el-input v-model="dialogdata.city" autocomplete="off" />
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取消</el-button>
            <el-button type="primary" @click="dialogConfim">
              确认
            </el-button>
          </span>
        </template>
      </el-dialog>
    </div>
  </div>

</template>

<script setup>
/* 引入 */
import { ref } from 'vue'

import request from './utils/requests.js'; // 引入request

/* 数据 */
// 打字机效果
const initialText = ref("你好，这是vue3-element-plus的一个小小案例,只是CURD的操作")
const index = ref(0)
const text = ref("")

setInterval = (() => {
  const autoTyping = () => {
    index.value ++
    text.value = initialText.value.slice(0,index);
    if (text.value.length >= initialText.value.length) {
      index.value = 0;
    }
  }
}, 300);


// 输入框
const query_input = ref('')

// 多选
const multipleSelection = ref([])

// 弹窗
const dialogFormVisible = ref(false)
const dialogFormTitle = ref('')
const dialogdata = ref({})

// 表格数据
const tableData = ref([
  {
    id: '1',
    name: 'Tom1',
    phone: "13800138000",
    email: "123@admin.com",
    status: "active",
    city: "广东"
  },
  {
    id: '2',
    name: 'Tom2',
    phone: "13800138000",
    email: "123@admin.com",
    status: "active",
    city: "广东"
  },
  {
    id: '3',
    name: 'Tom3',
    phone: "13800138000",
    email: "123@admin.com",
    status: "active",
    city: "广东"
  },
  {
    id: '4',
    name: 'Tom4',
    phone: "13800138000",
    email: "123@admin.com",
    status: "active",
    city: "广东"
  }
]
)
const tableDataCopy = Object.assign(tableData)  // 浅拷贝


// 分页: 页数, 当前页数
let total = ref(10)
let curPage = ref(1)


/* 方法 */
// 全量搜索

const getTableData = async (cur = 1) => {
  let res = await request.get("/list",{
    pageSize: 10,
    pageNum: cur
  })
  console.log(res)
  tableData.value = res.list
  
  total.value = res.total
  curPage.value = res.pageNum
  console.log(curPage)
}
getTableData(1)
 
// 分页
const handleChangePage = (val) => {
  getTableData(curPage.value)
}


// 搜索
const handleQueryName = async (val) => {
  console.log("val => ",val)

  // console.log(val)
  // if (val.length > 0) {
  //   tableData.value = tableData.value.filter(item => (item.name).toLowerCase().match(val.toLowerCase()))
  // } else {
  //   tableData.value = tableDataCopy.value
  // }

  // 搜索
  if (val.length > 0) {
    tableDataCopy.value = await request.get(`/list/${val}`)

  } else { 
    await getTableData(curPage.value)
  }

}


// 新增-弹窗
const handleadd = () => {
  dialogFormVisible.value = true; // 开启弹窗
  //dialogdata.value = {}  // 清空数据
  dialogFormTitle.value = "add"  // tatle
}

// 新增-弹窗-确认
const dialogConfim = async () => {
  dialogFormVisible.value = false; // 关闭弹窗

  // 判断tatle
  if (dialogFormTitle.value === "add") {
    // 添加数据
    // tableData.value.push({
    //   id: (tableData.length + 1).toString(),
    //   ...dialogdata.value
    // })

    // 添加数据到数据库
    await request.post("/add",{
      ...dialogdata.value
    })
    // 刷新页面
    await getTableData(curPage.value)

  } else {
    // const index = tableData.value.findIndex(item => item.id === dialogdata.value.id)
    // tableData.value[index] = dialogdata.value

    // 修改
    await request.put(`/update/${dialogdata.value.ID}`,{
      ...dialogdata.value
    })
    await getTableData(curPage.value)

  }
  
}

// 编辑
const handleEdit = (row) => {
  dialogFormVisible.value = true; // 开启弹窗
  dialogFormTitle.value = "edit"; // 编辑
  // console.log(row);

  dialogdata.value = { ...row}
    
}

// 选中删除
const handledel = () => {
  multipleSelection.value.forEach(ID => {
    deleteRow({ID})
  })
  console.log(multipleSelection.value)
}

//  表格点击删除
// 方式一
// const deleteRow = (index) => {
//   tableData.value.splice(index, 1)
//   console.log(index)
// }
// 方式二
// const deleteRow = ({row}) => {
//   console.log(row)  
// }
const deleteRow = async ({ID}) => {
  console.log("id => ",ID)
  // // 1. 通过id获取index索引的值
  // const index = tableData.value.findIndex(item => item.id == ID)
  // console.log("index => ",index)
  // // 2、根据索引进行删除
  // tableData.value.splice(index, 1)

  await request.delete(`/delete/${ID}`)
  await getTableData(curPage.value)
}


// 多选
const handleSelectionChange = (val) => {
  // multipleSelection.value = val
  // console.log(multipleSelection)

  multipleSelection.value = []
  val.forEach(item => {
    multipleSelection.value.push(item.ID)
  });
  console.log(multipleSelection.value)
}




</script>



<style scoped>
.table-demo {
  width: 800px;
  margin: 50px auto; 
  text-align: center;
}

h2 {
  text-align: center;
}

.qurey-box {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}

.el-input {
  width: 200px;
}

.btn-list {
  display: flex;
}

.pagination {
  margin: 20px 25%;
  
}
</style>
