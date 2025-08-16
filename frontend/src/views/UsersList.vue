<template>
  <div class="users-page">
    <div class="users-container">
      <!-- Header Section -->
      <UsersPageHeader
        :total="total"
        :selected-count="selectedUsers.length"
        @bulk-delete="confirmBulkDelete"
        @create-user="openCreateModal"
      />

      <!-- Search Section -->
      <UsersSearchBar
        v-model:search-query="searchQuery"
        @search="debouncedSearch"
      />

      <!-- Users Table -->
      <UsersTable
        v-if="users.length > 0"
        :users="users"
        :loading="loading"
        :sort-by="sortBy"
        :sort-dir="sortDir"
        @selection-change="handleSelectionChange"
        @sort-change="handleSortChange"
        @edit-user="openEditModal"
        @delete-user="confirmDelete"
      />

      <!-- Empty State -->
      <UsersEmptyState
        v-else
        :search-query="searchQuery"
        @create-user="openCreateModal"
      />

      <!-- Pagination -->
      <div class="pagination-section" v-if="users.length > 0">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :page-sizes="[5, 10, 20, 50]"
          :small="false"
          :disabled="loading"
          :background="true"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
          @size-change="loadUsers"
          @current-change="loadUsers"
        />
      </div>

      <!-- User Modal -->
      <UserModal
        v-model:visible="modalVisible"
        :user="selectedUser"
        @user-saved="handleUserSaved"
      />

      <!-- Delete Modal -->
      <DeleteUserModal
        v-model:visible="deleteModalVisible"
        :user="selectedUser"
        :loading="loading"
        @confirm="handleDeleteConfirm"
        @cancel="handleDeleteCancel"
      />
    </div>
  </div>
</template>

<script>
import { ref, onMounted, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { debounce } from '@/utils/helpers'
import { userService } from '@/services/userService'
import UserModal from '../components/UserModal/UserModal.vue'
import DeleteUserModal from '../components/UserModal/DeleteUserModal.vue'
import { UsersPageHeader, UsersSearchBar, UsersTable, UsersEmptyState } from '../components/UsersList'

export default {
  name: 'UsersList',
  components: {
    UserModal,
    DeleteUserModal,
    UsersPageHeader,
    UsersSearchBar,
    UsersTable,
    UsersEmptyState
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
    const deleteModalVisible = ref(false)

    // Methods
    const loadUsers = async () => {
      loading.value = true
      try {
        console.log('Loading users...')
        const response = await userService.getUsers({
          page: currentPage.value,
          perPage: pageSize.value,
          search: searchQuery.value,
          sortBy: sortBy.value,
          sortDir: sortDir.value
        })
        
        console.log('Response received:', response)
        console.log('Response data:', response.data)
        
        users.value = response.data.users || []
        total.value = response.data.total || 0
        
        console.log('Users set:', users.value)
        console.log('Total set:', total.value)
      } catch (error) {
        console.error('Error loading users:', error)
        ElMessage.error('Failed to load users')
      } finally {
        loading.value = false
      }
    }

    const openCreateModal = () => {
      selectedUser.value = null
      modalVisible.value = true
    }

    const openEditModal = (user) => {
      selectedUser.value = { ...user }
      modalVisible.value = true
    }

    const handleSelectionChange = (selection) => {
      selectedUsers.value = selection
    }

    const handleSortChange = ({ prop, order }) => {
      if (prop) {
        sortBy.value = prop
        sortDir.value = order === 'ascending' ? 'asc' : 'desc'
        loadUsers()
      }
    }

    const handleUserSaved = () => {
      loadUsers()
      modalVisible.value = false
      selectedUsers.value = []
    }

    const confirmDelete = async (user) => {
      selectedUser.value = user
      deleteModalVisible.value = true
    }

    const handleDeleteConfirm = async () => {
      try {
        await userService.deleteUser(selectedUser.value.id)
        ElMessage.success('User deleted successfully')
        loadUsers()
      } catch (error) {
        console.error('Error deleting user:', error)
        ElMessage.error('Failed to delete user')
      } finally {
        deleteModalVisible.value = false
        selectedUser.value = null
      }
    }

    const handleDeleteCancel = () => {
      deleteModalVisible.value = false
      selectedUser.value = null
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
      loadUsers,
      openCreateModal,
      openEditModal,
      handleSelectionChange,
      handleSortChange,
      handleUserSaved,
      confirmDelete,
      confirmBulkDelete,
      bulkDeleteUsers,
      debouncedSearch,
      deleteModalVisible,
      handleDeleteConfirm,
      handleDeleteCancel
    }
  }
}
</script>

<style scoped>
.users-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  padding: 32px;
}

.users-container {
  max-width: 90%;
  margin: 0 auto;
}

.pagination-section {
  display: flex;
  justify-content: center;
}

/* Custom Delete Confirmation Modal */
:deep(.delete-confirm-modal) {
  border-radius: 16px !important;
  overflow: hidden !important;
}

:deep(.delete-confirm-modal .el-message-box__header) {
  background: linear-gradient(135deg, #ff6b6b 0%, #ee5a52 100%) !important;
  padding: 24px 32px 20px 32px !important;
  border-bottom: none !important;
}

:deep(.delete-confirm-modal .el-message-box__title) {
  color: white !important;
  font-size: 20px !important;
  font-weight: 600 !important;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif !important;
  letter-spacing: -0.025em !important;
}

:deep(.delete-confirm-modal .el-message-box__close) {
  color: white !important;
  background: rgba(255, 255, 255, 0.15) !important;
  border-radius: 8px !important;
  width: 36px !important;
  height: 36px !important;
  border: 1px solid rgba(255, 255, 255, 0.2) !important;
}

:deep(.delete-confirm-modal .el-message-box__close:hover) {
  background: rgba(255, 255, 255, 0.25) !important;
  transform: scale(1.05) !important;
}

:deep(.delete-confirm-modal .el-message-box__content) {
  padding: 32px !important;
  font-size: 15px !important;
  color: #495057 !important;
  line-height: 1.5 !important;
}

:deep(.delete-confirm-modal .el-message-box__btns) {
  padding: 20px 32px 32px !important;
  gap: 12px !important;
  display: flex !important;
  justify-content: flex-end !important;
}

:deep(.delete-confirm-modal .el-button--default) {
  background: #f8f9fa !important;
  border: 1px solid #dee2e6 !important;
  color: #6c757d !important;
  border-radius: 10px !important;
  font-weight: 500 !important;
  padding: 10px 20px !important;
}

:deep(.delete-confirm-modal .el-button--default:hover) {
  background: #e9ecef !important;
  border-color: #adb5bd !important;
  transform: translateY(-1px) !important;
}

:deep(.delete-confirm-modal .el-button--primary) {
  background: linear-gradient(135deg, #ff6b6b 0%, #ee5a52 100%) !important;
  border: none !important;
  border-radius: 10px !important;
  font-weight: 600 !important;
  padding: 10px 24px !important;
  box-shadow: 0 4px 12px rgba(255, 107, 107, 0.3) !important;
}

:deep(.delete-confirm-modal .el-button--primary:hover) {
  transform: translateY(-2px) !important;
  box-shadow: 0 6px 16px rgba(255, 107, 107, 0.4) !important;
}

@media (max-width: 768px) {
  .users-page {
    padding: 16px;
  }
  
  .users-container {
    max-width: 100%;
  }
}
</style> 