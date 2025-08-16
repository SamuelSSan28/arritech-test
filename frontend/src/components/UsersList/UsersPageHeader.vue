<template>
  <div class="page-header">
    <div class="header-content">
      <h1 class="page-title">Users Management</h1>
      <p class="page-subtitle">{{ total }} users in total</p>
    </div>
    <div class="header-actions">
      <el-button 
        v-if="selectedCount > 0"
        type="danger" 
        size="large"
        :icon="Delete"
        @click="$emit('bulk-delete')"
        class="bulk-delete-btn"
      >
        Delete Selected ({{ selectedCount }})
      </el-button>
      <el-button 
        type="primary" 
        size="large"
        @click="$emit('create-user')"
        class="add-user-btn"
      >
        <el-icon><Plus /></el-icon>
        Add User
      </el-button>
    </div>
  </div>
</template>

<script>
import { Delete, Plus } from '@element-plus/icons-vue'

export default {
  name: 'UsersPageHeader',
  components: {
    Delete,
    Plus
  },
  props: {
    total: {
      type: Number,
      default: 0
    },
    selectedCount: {
      type: Number,
      default: 0
    }
  },
  emits: ['bulk-delete', 'create-user']
}
</script>

<style scoped>
.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 32px;
}

.header-content {
  flex: 1;
}

.page-title {
  margin: 0 0 4px 0;
  font-size: 28px;
  font-weight: 700;
  color: #1a1a1a;
}

.page-subtitle {
  margin: 0;
  color: #6c757d;
  font-size: 14px;
}

.header-actions {
  display: flex;
  gap: 12px;
}

.add-user-btn {
  border-radius: 12px !important;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
  border: none !important;
  font-weight: 600 !important;
  padding: 12px 24px !important;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3) !important;
  transition: all 0.3s ease !important;
  display: flex !important;
  align-items: center !important;
  gap: 8px !important;
}

.add-user-btn :deep(.el-icon) {
  font-size: 16px !important;
  margin-right: 4px !important;
}

.add-user-btn:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 6px 20px rgba(102, 126, 234, 0.4) !important;
}

.bulk-delete-btn {
  border-radius: 12px !important;
  background: linear-gradient(135deg, #ff6b6b 0%, #ee5a52 100%) !important;
  border: none !important;
  font-weight: 600 !important;
  padding: 12px 24px !important;
  box-shadow: 0 4px 12px rgba(255, 107, 107, 0.3) !important;
  transition: all 0.3s ease !important;
}

.bulk-delete-btn:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 6px 20px rgba(255, 107, 107, 0.4) !important;
}

@media (max-width: 768px) {
  .page-header {
    flex-direction: column;
    align-items: stretch;
    gap: 16px;
  }
  
  .header-actions {
    justify-content: center;
  }
}
</style> 