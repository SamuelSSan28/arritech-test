<template>
  <div class="users-page">
    <!-- Header Section -->
    <div class="page-header">
      <div class="header-content">
        <h1 class="page-title">Users Management</h1>
        <p class="page-subtitle">{{ total }} users in total</p>
      </div>
      <div class="header-actions">
        <el-button 
          v-if="selectedUsers.length > 0"
          type="danger" 
          size="large"
          :icon="Delete"
          @click="confirmBulkDelete"
          class="bulk-delete-btn"
        >
          Delete Selected ({{ selectedUsers.length }})
        </el-button>
        <el-button 
          type="primary" 
          size="large"
          :icon="Plus" 
          @click="openCreateModal"
          class="add-user-btn"
        >
          Add User
        </el-button>
      </div>
    </div>

    <!-- Search Section -->
    <div class="search-section">
      <el-input
        v-model="searchQuery"
        placeholder="Search users by name, email, or phone..."
        :prefix-icon="Search"
        clearable
        @input="debouncedSearch"
        size="large"
        class="search-input"
      />
    </div>

    <!-- Users Table -->
    <div class="table-container" v-if="users.length > 0">
      <el-table
        :data="users"
        v-loading="loading"
        stripe
        style="width: 100%"
        @selection-change="handleSelectionChange"
        @sort-change="handleSortChange"
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
            <el-tag size="default" class="age-tag">{{ row.age }}</el-tag>
          </template>
        </el-table-column>
        
        <el-table-column prop="phone" label="Phone" min-width="140" sortable="custom">
          <template #header="{ column }">
                          <div class="column-header" :class="{ active: sortBy === 'phone' }">
                Phone
                <el-icon v-if="sortBy === 'phone'" class="sort-icon sorted">
                  <ArrowUpBold v-if="sortDir === 'asc'" />
                  <ArrowDownBold v-else />
                </el-icon>
                <el-icon v-else class="sort-icon normal">
                  <ArrowUp />
                </el-icon>
              </div>
          </template>
          <template #default="{ row }">
            <div class="phone-cell">
              {{ formatPhoneWithFlag(row.phone) }}
            </div>
          </template>
        </el-table-column>
        
        <el-table-column prop="created_at" label="Joined" width="160" sortable="custom">
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
        
        <el-table-column label="Actions" width="120" fixed="right">
          <template #default="{ row }">
            <div class="action-buttons">
              <el-button
                type="primary"
                size="small"
                :icon="Edit"
                @click="openEditModal(row)"
                circle
              />
              <el-button
                type="danger"
                size="small"
                :icon="Delete"
                @click="confirmDelete(row)"
                circle
              />
            </div>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- Empty State -->
    <div class="empty-state" v-else-if="!loading">
      <div class="empty-illustration">
        <svg width="200" height="200" viewBox="0 0 200 200" fill="none" xmlns="http://www.w3.org/2000/svg">
          <!-- Background circle -->
          <circle cx="100" cy="100" r="90" fill="#f8f9fa" stroke="#e9ecef" stroke-width="2"/>
          
          <!-- User icon -->
          <circle cx="100" cy="70" r="25" fill="#dee2e6" stroke="#adb5bd" stroke-width="2"/>
          <path d="M60 140c0-22.091 17.909-40 40-40s40 17.909 40 40" fill="#dee2e6" stroke="#adb5bd" stroke-width="2"/>
          
          <!-- Plus icon -->
          <circle cx="140" cy="60" r="15" fill="#667eea" stroke="#5a6fd8" stroke-width="2"/>
          <path d="M140 52v16M132 60h16" stroke="white" stroke-width="2" stroke-linecap="round"/>
          
          <!-- Decorative elements -->
          <circle cx="70" cy="50" r="8" fill="#e9ecef" opacity="0.6"/>
          <circle cx="130" cy="120" r="6" fill="#e9ecef" opacity="0.4"/>
          <circle cx="50" cy="130" r="10" fill="#e9ecef" opacity="0.3"/>
        </svg>
      </div>
      <h3 class="empty-title">No users found</h3>
      <p class="empty-description">
        {{ searchQuery ? 'Try adjusting your search terms' : 'Get started by adding your first user' }}
      </p>
      <el-button 
        type="primary" 
        size="large"
        :icon="Plus"
        @click="openCreateModal"
      >
        Add First User
      </el-button>
    </div>

    <!-- Pagination -->
    <div class="pagination-section" v-if="total > 0">
      <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :total="total"
        :page-sizes="[10, 20, 50]"
        layout="total, sizes, prev, pager, next"
        @current-change="handlePageChange"
        @size-change="handlePageSizeChange"
        background
      />
    </div>

    <!-- User Modal -->
    <UserModal
      v-model:visible="modalVisible"
      :user="selectedUser"
      @user-saved="handleUserSaved"
    />
  </div>
</template>

<script>
import { ref, onMounted, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus, Search, Edit, Delete, ArrowUp, ArrowDown, ArrowUpBold, ArrowDownBold } from '@element-plus/icons-vue'
import { userService } from '@/services/userService'
import { debounce } from '@/utils/helpers'
import { formatPhoneWithFlag } from '@/utils/countryFlags'
import UserModal from '@/components/UserModal.vue'

export default {
  name: 'UsersList',
  components: {
    UserModal,
    ArrowUp,
    ArrowDown,
    ArrowUpBold,
    ArrowDownBold
  },
  setup() {
    // Reactive data
    const users = ref([])
    const loading = ref(false)
    const searchQuery = ref('')
    const currentPage = ref(1)
    const pageSize = ref(10)
    const total = ref(0)
    const modalVisible = ref(false)
    const selectedUser = ref(null)
    const selectedUsers = ref([])
    const sortBy = ref('created_at')
    const sortDir = ref('desc')

    // Methods
    const loadUsers = async () => {
      loading.value = true
      try {
        const response = await userService.getUsers({
          page: currentPage.value,
          perPage: pageSize.value,
          search: searchQuery.value,
          sortBy: sortBy.value,
          sortDir: sortDir.value
        })

        users.value = response.data.users || []
        total.value = response.data.total || 0
      } catch (error) {
        console.error('Error loading users:', error)
        ElMessage.error('Failed to load users')
      } finally {
        loading.value = false
      }
    }

    const handlePageChange = (page) => {
      currentPage.value = page
      loadUsers()
    }

    const handlePageSizeChange = (size) => {
      pageSize.value = size
      currentPage.value = 1
      loadUsers()
    }

    const handleSelectionChange = (selection) => {
      selectedUsers.value = selection
    }

    const handleSortChange = ({ prop }) => {
      // Map frontend field names to backend field names
      const fieldMapping = {
        'id': 'id',
        'name': 'name',
        'email': 'email',
        'age': 'age',
        'phone': 'phone',
        'created_at': 'created_at',
        'updated_at': 'updated_at'
      }
      
      const field = fieldMapping[prop] || prop
      
      // If clicking the same column
      if (field === sortBy.value) {
        // Cycle through: asc -> desc -> none (default)
        if (sortDir.value === 'asc') {
          sortDir.value = 'desc'
        } else if (sortDir.value === 'desc') {
          // Reset to default sort
          sortBy.value = 'created_at'
          sortDir.value = 'desc'
        }
      } else {
        // New column, start with ascending
        sortBy.value = field
        sortDir.value = 'asc'
      }
      
      currentPage.value = 1 // Reset to first page when sorting
      loadUsers()
    }

    const openCreateModal = () => {
      selectedUser.value = null
      modalVisible.value = true
    }

    const openEditModal = (user) => {
      selectedUser.value = { ...user }
      modalVisible.value = true
    }

    const handleUserSaved = () => {
      loadUsers()
      modalVisible.value = false
      selectedUsers.value = []
    }

    const confirmDelete = async (user) => {
      try {
        await ElMessageBox.confirm(
          `Delete user "${user.name}"?`,
          'Confirm Delete',
          {
            confirmButtonText: 'Delete',
            cancelButtonText: 'Cancel',
            type: 'warning'
          }
        )

        await deleteUser(user.id)
      } catch (error) {
        // User cancelled
      }
    }

    const confirmBulkDelete = async () => {
      if (selectedUsers.value.length === 0) return
      
      try {
        const userNames = selectedUsers.value.map(u => u.name).join(', ')
        await ElMessageBox.confirm(
          `Delete ${selectedUsers.value.length} selected users?\n\n${userNames}`,
          'Confirm Bulk Delete',
          {
            confirmButtonText: 'Delete All',
            cancelButtonText: 'Cancel',
            type: 'warning'
          }
        )

        await bulkDeleteUsers()
      } catch (error) {
        // User cancelled
      }
    }

    const deleteUser = async (id) => {
      try {
        await userService.deleteUser(id)
        ElMessage.success('User deleted successfully')
        loadUsers()
      } catch (error) {
        console.error('Error deleting user:', error)
        ElMessage.error('Failed to delete user')
      }
    }

    const bulkDeleteUsers = async () => {
      try {
        const deletePromises = selectedUsers.value.map(user => 
          userService.deleteUser(user.id)
        )
        
        await Promise.all(deletePromises)
        
        ElMessage.success(`${selectedUsers.value.length} users deleted successfully`)
        selectedUsers.value = []
        loadUsers()
      } catch (error) {
        console.error('Error bulk deleting users:', error)
        ElMessage.error('Failed to delete some users')
      }
    }

    const formatDate = (dateString) => {
      if (!dateString) return '—'
      return new Date(dateString).toLocaleDateString('en-US', {
        month: 'short',
        day: 'numeric',
        year: 'numeric'
      })
    }

    // Debounced search
    const debouncedSearch = debounce(() => {
      currentPage.value = 1
      loadUsers()
    }, 500)

    // Watchers
    watch(searchQuery, () => {
      if (searchQuery.value === '') {
        currentPage.value = 1
        loadUsers()
      }
    })

    // Lifecycle
    onMounted(() => {
      console.log('UsersList component mounted')
      loadUsers()
    })

    // Debug watcher
    watch(users, (newUsers) => {
      console.log('Users updated:', newUsers.length)
    })

    watch(loading, (newLoading) => {
      console.log('Loading state:', newLoading)
    })

    return {
      users,
      loading,
      searchQuery,
      currentPage,
      pageSize,
      total,
      modalVisible,
      selectedUser,
      selectedUsers,
      sortBy,
      sortDir,
      handlePageChange,
      handlePageSizeChange,
      handleSelectionChange,
      handleSortChange,
      openCreateModal,
      openEditModal,
      handleUserSaved,
      confirmDelete,
      confirmBulkDelete,
      formatDate,
      debouncedSearch,
      formatPhoneWithFlag,
      // Icons
      Plus,
      Search,
      Edit,
      Delete,
      ArrowUp,
      ArrowDown,
      ArrowUpBold,
      ArrowDownBold
    }
  }
}
</script>

<style scoped>
.users-page {
  padding: 24px;
  max-width: 1200px;
  margin: 0 auto;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 32px;
}

.header-content h1 {
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

.add-user-btn {
  border-radius: 12px !important;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
  border: none !important;
  font-weight: 600 !important;
  padding: 12px 24px !important;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3) !important;
  transition: all 0.3s ease !important;
}

.add-user-btn:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 6px 20px rgba(102, 126, 234, 0.4) !important;
}

.search-section {
  margin-bottom: 24px;
}

.search-input {
  max-width: 400px;
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

.table-container {
  background: white;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  margin-bottom: 24px;
}

.users-table {
  border: none;
}

.column-header {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  cursor: pointer;
  font-weight: 600;
  font-size: 14px;
  color: #495057;
  transition: all 0.2s ease;
  padding: 8px 4px;
  
  &:hover {
    color: #667eea;
  }
  
  &.active {
    color: #667eea;
  }
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
  display: none !important; /* esconde as setas padrão do EL-Table */
}

.column-header:hover .sort-icons {
  opacity: 1;
}

.age-tag {
  font-size: 14px;
  padding: 4px 12px;
  border-radius: 8px;
  background: linear-gradient(135deg, #667eea20 0%, #764ba220 100%);
  color: #667eea;
  border: none;
  font-weight: 600;
}

.user-cell {
  display: flex;
  align-items: center;
  gap: 12px;
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
}

.action-buttons {
  display: flex;
  gap: 8px;
}

.action-buttons :deep(.el-button) {
  border-radius: 10px !important;
  border: none !important;
  transition: all 0.3s ease !important;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1) !important;
}

.action-buttons :deep(.el-button--primary) {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
}

.action-buttons :deep(.el-button--danger) {
  background: linear-gradient(135deg, #ff6b6b 0%, #ee5a52 100%) !important;
}

.action-buttons :deep(.el-button:hover) {
  transform: translateY(-2px) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15) !important;
}

.pagination-section {
  display: flex;
  justify-content: center;
}

/* Empty State Styles */
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
  padding: 0 !important;
}

:deep(.el-table th.is-sortable) {
  .cell {
    padding: 0;
  }
}

:deep(.el-table td) {
  border: none !important;
  border-bottom: 1px solid #f1f3f4 !important;
}

:deep(.el-table__row) {
  transition: background-color 0.2s ease;
  
  &:hover {
    background-color: #f8f9ff !important;
  }
}

@media (max-width: 768px) {
  .users-page {
    padding: 16px;
  }
  
  .page-header {
    flex-direction: column;
    align-items: stretch;
    gap: 16px;
  }
  
  .header-actions {
    justify-content: center;
  }
  
  .search-input {
    max-width: none;
  }
  
  .empty-state {
    padding: 40px 20px;
  }
  
  .empty-illustration svg {
    width: 150px;
    height: 150px;
  }
}
</style> 