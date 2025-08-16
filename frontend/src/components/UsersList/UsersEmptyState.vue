<template>
  <div class="empty-state">
    <div class="empty-illustration">
      <svg width="200" height="200" viewBox="0 0 200 200" fill="none">
        <circle cx="100" cy="100" r="80" fill="#f3f4f6" stroke="#e5e7eb" stroke-width="2"/>
        <circle cx="85" cy="85" r="8" fill="#9ca3af"/>
        <circle cx="115" cy="85" r="8" fill="#9ca3af"/>
        <path d="M75 125 Q100 145 125 125" stroke="#9ca3af" stroke-width="3" fill="none" stroke-linecap="round"/>
      </svg>
    </div>
    <h3 class="empty-title">No users found</h3>
    <p class="empty-description">
      {{ searchQuery ? 'Try adjusting your search terms or filters.' : 'Start by adding your first user to the system.' }}
    </p>
    <el-button 
      v-if="!searchQuery"
      type="primary" 
      size="large"
      :icon="Plus"
      @click="$emit('create-user')"
      class="empty-action-btn"
    >
      Add First User
    </el-button>
  </div>
</template>

<script>
import { Plus } from '@element-plus/icons-vue'

export default {
  name: 'UsersEmptyState',
  components: {
    Plus
  },
  props: {
    searchQuery: {
      type: String,
      default: ''
    }
  },
  emits: ['create-user']
}
</script>

<style scoped>
.empty-state {
  text-align: center;
  padding: 80px 20px;
  background: white;
  border-radius: 16px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  margin: 24px 0;
}

.empty-illustration {
  margin-bottom: 24px;
  opacity: 0.7;
}

.empty-title {
  font-size: 24px;
  font-weight: 600;
  color: #495057;
  margin-bottom: 12px;
}

.empty-description {
  font-size: 16px;
  color: #6c757d;
  margin-bottom: 24px;
  max-width: 400px;
  margin-left: auto;
  margin-right: auto;
}

.empty-action-btn {
  border-radius: 12px !important;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
  border: none !important;
  font-weight: 600 !important;
  padding: 12px 24px !important;
  transition: all 0.3s ease !important;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3) !important;
}

.empty-action-btn:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 6px 20px rgba(102, 126, 234, 0.4) !important;
}

@media (max-width: 768px) {
  .empty-state {
    padding: 40px 20px;
  }
  
  .empty-illustration svg {
    width: 150px;
    height: 150px;
  }
}
</style> 