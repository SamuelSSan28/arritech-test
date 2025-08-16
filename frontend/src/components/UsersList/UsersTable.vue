<template>
  <div class="table-container">
    <el-table
      :data="users"
      v-loading="loading"
      stripe
      style="width: 100%"
      @selection-change="$emit('selection-change', $event)"
      @sort-change="$emit('sort-change', $event)"
      class="users-table"
    >
      <el-table-column type="selection" width="55" />
      
      <el-table-column prop="id" label="ID" width="100" sortable="custom">
        <template #header="{ column }">
          <div class="column-header" :class="{ active: sortBy === 'id' }">
            ID
            <el-icon v-if="sortBy === 'id'" class="sort-icon sorted">
              <ArrowUpBold v-if="sortDir === 'asc'" />
              <ArrowDownBold v-else />
            </el-icon>
            <el-icon v-else class="sort-icon normal">
              <ArrowUp />
            </el-icon>
          </div>
        </template>
      </el-table-column>
      
      <el-table-column prop="name" label="User" min-width="250" sortable="custom">
        <template #header="{ column }">
          <div class="column-header" :class="{ active: sortBy === 'name' }">
            User
            <el-icon v-if="sortBy === 'name'" class="sort-icon sorted">
              <ArrowUpBold v-if="sortDir === 'asc'" />
              <ArrowDownBold v-else />
            </el-icon>
            <el-icon v-else class="sort-icon normal">
              <ArrowUp />
            </el-icon>
          </div>
        </template>
        <template #default="{ row }">
          <div class="user-cell">
            <el-avatar :size="36" class="user-avatar">
              {{ row.name.charAt(0).toUpperCase() }}
            </el-avatar>
            <div class="user-info">
              <div class="user-name">{{ row.name }}</div>
              <div class="user-email">{{ row.email }}</div>
            </div>
          </div>
        </template>
      </el-table-column>
      
      <el-table-column prop="age" label="Age" width="120" sortable="custom">
        <template #header="{ column }">
          <div class="column-header" :class="{ active: sortBy === 'age' }">
            Age
            <el-icon v-if="sortBy === 'age'" class="sort-icon sorted">
              <ArrowUpBold v-if="sortDir === 'asc'" />
              <ArrowDownBold v-else />
            </el-icon>
            <el-icon v-else class="sort-icon normal">
              <ArrowUp />
            </el-icon>
          </div>
        </template>
        <template #default="{ row }">
          <el-tag class="age-tag">
            {{ row.age }} years
          </el-tag>
        </template>
      </el-table-column>
      
      <el-table-column prop="phone" label="Phone" width="200">
        <template #default="{ row }">
          <div class="phone-cell" v-if="row.phone">
            <span>{{ formatPhone(row.phone) }}</span>
            <span class="phone-flag">{{ getFlagFromPhone(row.phone) }}</span>
          </div>
          <span v-else class="no-data">—</span>
        </template>
      </el-table-column>
      
      <el-table-column prop="created_at" label="Joined" width="140" sortable="custom">
        <template #header="{ column }">
          <div class="column-header" :class="{ active: sortBy === 'created_at' }">
            Joined
            <el-icon v-if="sortBy === 'created_at'" class="sort-icon sorted">
              <ArrowUpBold v-if="sortDir === 'asc'" />
              <ArrowDownBold v-else />
            </el-icon>
            <el-icon v-else class="sort-icon normal">
              <ArrowUp />
            </el-icon>
          </div>
        </template>
        <template #default="{ row }">
          {{ formatDate(row.created_at) }}
        </template>
      </el-table-column>
      
      <el-table-column label="Actions" width="140" fixed="right">
        <template #default="{ row }">
          <div class="action-buttons">
            <el-button
              type="primary"
              size="small"
              @click="$emit('edit-user', row)"
            >
              <el-icon><Edit /></el-icon>
            </el-button>
            <el-button
              type="danger"
              size="small"
              @click="$emit('delete-user', row)"
            >
              <el-icon><Delete /></el-icon>
            </el-button>
          </div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { Edit, Delete, ArrowUp, ArrowDown, ArrowUpBold, ArrowDownBold } from '@element-plus/icons-vue'
import { calculateAge } from '@/utils/helpers'
import { formatPhone } from '@/utils/helpers'
import { formatDate } from '@/utils/helpers'
import { getFlagFromPhone } from '@/utils/countryFlags'

export default {
  name: 'UsersTable',
  components: {
    Edit,
    Delete,
    ArrowUp,
    ArrowDown,
    ArrowUpBold,
    ArrowDownBold
  },
  props: {
    users: {
      type: Array,
      default: () => []
    },
    loading: {
      type: Boolean,
      default: false
    },
    sortBy: {
      type: String,
      default: 'created_at'
    },
    sortDir: {
      type: String,
      default: 'desc'
    }
  },
  emits: ['selection-change', 'sort-change', 'edit-user', 'delete-user'],
  methods: {
    calculateAge,
    formatPhone,
    formatDate,
    getFlagFromPhone
  }
}
</script>

<style scoped>
.table-container {
  background: white;
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  margin-bottom: 32px;
}

.column-header {
  display: flex;
  align-items: center;
  gap: 6px;
  cursor: pointer;
  font-weight: 600;
  font-size: 14px;
  color: #495057;
  transition: all 0.2s ease;
  padding: 8px 4px;
}

.column-header:hover {
  color: #667eea;
}

.column-header.active {
  color: #667eea;
}

.sort-icon {
  font-size: 14px;
}

.sort-icon.normal {
  color: #c0c4cc;
  opacity: 0.6;
}

.sort-icon.sorted {
  color: #667eea;
  font-weight: bold;
}

:deep(.users-table .caret-wrapper) {
  display: none !important;
}

.column-header:hover .sort-icons {
  opacity: 1;
}

.age-tag {
  font-size: 14px;
  padding: 6px 12px;
  border-radius: 8px;
  background: linear-gradient(135deg, #667eea20 0%, #764ba220 100%);
  color: #667eea;
  border: none;
  font-weight: 600;
  margin: 2px 0;
}

.user-cell {
  display: flex;
  align-items: center;
  gap: 12px;
  height: 100%;
  padding: 4px 0;
}

.user-avatar {
  font-weight: 600;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.user-info {
  flex: 1;
  min-width: 0;
}

.user-name {
  font-weight: 600;
  color: #1a1a1a;
  margin-bottom: 2px;
}

.user-email {
  font-size: 13px;
  color: #6c757d;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.phone-cell {
  display: flex;
  align-items: center;
  gap: 8px;
  height: 100%;
}

.phone-flag {
  font-size: 16px;
}

.no-data {
  color: #9ca3af;
  font-style: italic;
}

.action-buttons {
  display: flex;
  gap: 8px;
  height: 100%;
  align-items: center;
  justify-content: center;
}

.action-buttons :deep(.el-button) {
  border-radius: 6px !important;
  border: none !important;
  transition: all 0.2s ease !important;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1) !important;
  transform: none !important;
  min-width: 28px !important;
  height: 28px !important;
  padding: 4px 6px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

.action-buttons :deep(.el-button--primary) {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
}

.action-buttons :deep(.el-button--danger) {
  background: linear-gradient(135deg, #ff6b6b 0%, #ee5a52 100%) !important;
}

.action-buttons :deep(.el-button:hover) {
  transform: scale(1.1) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
}

.action-buttons :deep(.el-button .el-icon) {
  font-size: 14px !important;
  color: white !important;
  margin: 0 !important;
}

:deep(.el-table) {
  border: none !important;
}

:deep(.el-table__header) {
  background: #f8f9fa !important;
  border-bottom: 2px solid #e9ecef !important;
}

:deep(.el-table th) {
  background: transparent !important;
  border: none !important;
  color: #495057 !important;
  font-weight: 600 !important;
  font-size: 14px !important;
  padding: 12px !important;
  height: 48px !important;
}

:deep(.el-table th.is-sortable .cell) {
  padding: 0;
}

:deep(.el-table td) {
  border: none !important;
  border-bottom: 1px solid #f1f3f4 !important;
  padding: 12px !important;
  height: 50px !important;
  vertical-align: middle !important;
}

:deep(.el-table__row) {
  transition: background-color 0.2s ease;
  height: 50px !important;
}

:deep(.el-table__row:hover) {
  background-color: #f8f9ff !important;
}

@media (max-width: 768px) {
  .table-container {
    border-radius: 16px;
    margin-bottom: 24px;
  }
  
  .action-buttons {
    flex-direction: column;
    gap: 4px;
  }
}
</style> 