<template>
  <el-dialog
    v-model="dialogVisible"
    width="360px"
    :before-close="handleClose"
    :close-on-click-modal="false"
    destroy-on-close
    align-center
    class="delete-modal"
  >
    <template #title>
      <div class="dialog-header">
        <el-icon class="warning-icon"><Warning /></el-icon>
        <span class="dialog-title">Delete User</span>
      </div>
    </template>

    <div class="modal-content">
      <p class="confirmation-text">
        Are you sure you want to delete "<span class="user-name">{{ user?.name }}</span>"?
      </p>
    </div>

    <template #footer>
      <div class="modal-actions">
        <el-button 
          @click="handleClose" 
          :disabled="loading" 
          size="large"
          class="cancel-btn"
        >
          Cancel
        </el-button>
        <el-button 
          type="danger" 
          @click="handleConfirm"
          :loading="loading"
          size="large"
          class="delete-btn"
        >
          Delete
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script>
import { ref, computed } from 'vue'
import { Warning } from '@element-plus/icons-vue'

export default {
  name: 'DeleteUserModal',
  components: {
    Warning
  },
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    user: {
      type: Object,
      default: null
    },
    loading: {
      type: Boolean,
      default: false
    }
  },
  emits: ['update:visible', 'confirm', 'cancel'],
  setup(props, { emit }) {
    const dialogVisible = computed({
      get: () => props.visible,
      set: (value) => emit('update:visible', value)
    })

    const handleClose = () => {
      emit('cancel')
      dialogVisible.value = false
    }

    const handleConfirm = () => {
      emit('confirm')
    }

    return {
      dialogVisible,
      handleClose,
      handleConfirm
    }
  }
}
</script>

<style scoped>
.delete-modal {
  border-radius: 16px;
  overflow: hidden;
}

.dialog-header {
  display: flex;
  align-items: center;
  gap: 12px;
}

.warning-icon {
  font-size: 24px;
  color: #ff6b6b;
}

.dialog-title {
  font-size: 20px;
  font-weight: 700;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  letter-spacing: -0.025em;
  line-height: 1.4;
  background: linear-gradient(135deg, #ff6b6b 0%, #ee5a52 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.delete-modal :deep(.el-dialog__header) {
  background: white;
  padding: 16px 24px;
  margin: 0;
  border-bottom: 1px solid #e9ecef;
  position: relative;
}

.modal-content {
  padding: 12px 16px;
  text-align: left;
}

.confirmation-text {
  font-size: 13px;
  color: #374151;
  line-height: 1.4;
  margin: 0;
}

.user-name {
  font-weight: 600;
  color: #111827;
}

.warning-text {
  font-size: 13px;
  color: #ff6b6b;
  margin: 0;
  font-weight: 500;
}

.modal-actions {
  padding: 0 16px 12px;
  display: flex;
  justify-content: stretch;
  gap: 8px;
}

.cancel-btn {
  flex: 1;
  border-radius: 6px !important;
  background: #f8f9fa !important;
  border: 1px solid #dee2e6 !important;
  color: #6c757d !important;
  font-weight: 500 !important;
  padding: 4px 12px !important;
  transition: all 0.2s ease !important;
  font-size: 13px !important;
  height: 28px !important;
}

.cancel-btn:hover {
  background: #e9ecef !important;
  border-color: #adb5bd !important;
  transform: translateY(-1px) !important;
}

.delete-btn {
  flex: 1;
  border-radius: 6px !important;
  background: linear-gradient(135deg, #ff6b6b 0%, #ee5a52 100%) !important;
  border: none !important;
  font-weight: 600 !important;
  padding: 4px 12px !important;
  transition: all 0.2s ease !important;
  font-size: 13px !important;
  box-shadow: 0 2px 8px rgba(255, 107, 107, 0.3) !important;
  height: 28px !important;
}

.delete-btn:hover {
  transform: translateY(-1px) !important;
  box-shadow: 0 4px 12px rgba(255, 107, 107, 0.4) !important;
}

@media (max-width: 768px) {
  .modal-content {
    padding: 12px;
  }

  .modal-actions {
    padding: 0 12px 12px;
  }

  .cancel-btn,
  .delete-btn {
    width: 100% !important;
  }
}
</style> 