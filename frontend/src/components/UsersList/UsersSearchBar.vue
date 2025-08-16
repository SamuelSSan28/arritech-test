<template>
  <div class="search-section">
    <el-input
      :model-value="searchQuery"
      @input="handleInput"
      placeholder="Search users by name, email, or phone..."
      :prefix-icon="Search"
      clearable
      size="large"
      class="search-input"
    />
  </div>
</template>

<script>
import { Search } from '@element-plus/icons-vue'

export default {
  name: 'UsersSearchBar',
  components: {
    Search
  },
  props: {
    searchQuery: {
      type: String,
      default: ''
    }
  },
  emits: ['update:searchQuery', 'search'],
  methods: {
    handleInput(value) {
      this.$emit('update:searchQuery', value)
      this.$emit('search', value)
    }
  }
}
</script>

<style scoped>
.search-section {
  margin-bottom: 24px;
}

.search-input {
  max-width: 400px;
  width: 100%;
}

.search-input :deep(.el-input__wrapper) {
  border-radius: 12px !important;
  border: 2px solid #e9ecef !important;
  transition: all 0.3s ease !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05) !important;
}

.search-input :deep(.el-input__wrapper:hover) {
  border-color: #667eea !important;
  box-shadow: 0 4px 8px rgba(102, 126, 234, 0.1) !important;
}

.search-input :deep(.el-input__wrapper.is-focus) {
  border-color: #667eea !important;
  box-shadow: 0 0 0 4px rgba(102, 126, 234, 0.1) !important;
}

@media (max-width: 768px) {
  .search-input {
    max-width: none;
  }
}
</style> 