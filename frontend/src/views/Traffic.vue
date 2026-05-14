<template>
  <div class="traffic-page">
    <el-row :gutter="20">
      <!-- 添加里程表单 -->
      <el-col :span="8">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>添加里程记录</span>
            </div>
          </template>
          
          <el-form :model="form" label-width="100px">
            <el-form-item label="用户ID">
              <el-input-number v-model="form.userId" :min="1" style="width: 100%" />
            </el-form-item>
            
            <el-form-item label="车牌号">
              <el-input v-model="form.plateNumber" placeholder="例如：粤A00001" />
            </el-form-item>
            
            <el-form-item label="行驶里程">
              <el-input-number v-model="form.mileage" :min="0" :precision="2" style="width: 100%" />
              <span style="margin-left: 10px; color: #909399">公里</span>
            </el-form-item>
            
            <el-form-item label="通行时间">
              <el-date-picker
                v-model="form.trafficTime"
                type="datetime"
                placeholder="选择日期时间"
                format="YYYY-MM-DD HH:mm:ss"
                value-format="YYYY-MM-DD HH:mm:ss"
                style="width: 100%"
              />
            </el-form-item>
            
            <el-form-item>
              <el-button type="primary" @click="handleSubmit" :loading="submitting">
                提交
              </el-button>
              <el-button @click="handleReset">重置</el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
      
      <!-- 里程列表 -->
      <el-col :span="16">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>里程记录列表</span>
              <el-button type="success" @click="fetchTrafficList">
                <el-icon><Refresh /></el-icon>
                刷新
              </el-button>
            </div>
          </template>
          
          <el-table :data="trafficList" border stripe v-loading="loading">
            <el-table-column prop="id" label="ID" width="80" />
            <el-table-column prop="userId" label="用户ID" width="100" />
            <el-table-column prop="plateNumber" label="车牌号" width="120" />
            <el-table-column prop="mileage" label="里程(公里)" width="120" />
            <el-table-column prop="trafficTime" label="通行时间" width="180" />
            <el-table-column prop="isCalculate" label="核算状态" width="100">
              <template #default="{ row }">
                <el-tag :type="row.isCalculate === 1 ? 'success' : 'info'">
                  {{ row.isCalculate === 1 ? '已核算' : '未核算' }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="createTime" label="创建时间" />
          </el-table>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { trafficApi } from '@/api'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const loading = ref(false)
const submitting = ref(false)
const trafficList = ref([])

const form = reactive({
  userId: userStore.currentUserId,
  plateNumber: '',
  mileage: 0,
  trafficTime: ''
})

// 获取里程列表
const fetchTrafficList = async () => {
  loading.value = true
  try {
    const res = await trafficApi.getTrafficList(userStore.currentUserId)
    trafficList.value = res.list || []
  } catch (error) {
    ElMessage.error('获取里程列表失败')
  } finally {
    loading.value = false
  }
}

// 提交表单
const handleSubmit = async () => {
  if (!form.plateNumber || !form.mileage || !form.trafficTime) {
    ElMessage.warning('请填写完整信息')
    return
  }
  
  submitting.value = true
  try {
    await trafficApi.addTraffic(form)
    ElMessage.success('添加成功')
    handleReset()
    fetchTrafficList()
  } catch (error) {
    ElMessage.error('添加失败')
  } finally {
    submitting.value = false
  }
}

// 重置表单
const handleReset = () => {
  form.userId = userStore.currentUserId
  form.plateNumber = ''
  form.mileage = 0
  form.trafficTime = ''
}

onMounted(() => {
  fetchTrafficList()
})
</script>

<style scoped>
.traffic-page {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
