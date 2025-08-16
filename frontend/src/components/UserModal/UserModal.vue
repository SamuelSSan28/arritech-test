<template>
  <el-dialog
    v-model="dialogVisible"
    width="600px"
    :before-close="handleClose"
    :close-on-click-modal="false"
    destroy-on-close
    align-center
    class="user-modal"
  >
    <template #title>
      <span class="dialog-title">{{ isEditing ? 'Edit User' : 'Add New User' }}</span>
    </template>
    <div class="modal-content">
      <UserForm
        ref="userForm"
        :form-data="formData"
        :is-editing="isEditing"
        :calculatedAge="calculatedAge"
        :loading="loading"
        :validation-rules="validationRules"
        :errors="errors"
        @submit="handleSubmit"
      />
      
      <div class="form-actions">
        <el-button 
          @click="handleClose" 
          :disabled="loading" 
          size="large"
          class="cancel-btn"
        >
          Cancel
        </el-button>
        <el-button 
          type="primary" 
          @click="validateAndSubmit"
          :loading="loading"
          size="large"
          class="submit-btn"
        >
          {{ isEditing ? 'Update User' : 'Create User' }}
        </el-button>
      </div>
    </div>
  </el-dialog>
</template>

<script>
import { ref } from 'vue'
import UserForm from './UserForm.vue'
import { useUserModal } from '../../composables/useUserModal'

export default {
  name: 'UserModal',
  components: {
    UserForm
  },
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    user: {
      type: Object,
      default: null
    }
  },
  emits: ['update:visible', 'user-saved'],
  setup(props, { emit }) {
    const userForm = ref(null)

    const validateAndSubmit = () => {
      userForm.value?.validate((valid) => {
        if (valid) {
          handleSubmit()
        }
      })
    }

    // Use the modal composable
    const {
      loading,
      formData,
      calculatedAge,
      errors,
      dialogVisible,
      isEditing,
      validationRules,
      handleSubmit,
      handleClose
    } = useUserModal(props, emit)

    return {
      loading,
      userForm,
      formData,
      calculatedAge,
      errors,
      dialogVisible,
      isEditing,
      validationRules,
      handleSubmit,
      handleClose,
      validateAndSubmit
    }
  }
}
</script>

<style scoped>
.user-modal :deep(.el-dialog) {
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.15);
  border: 1px solid #e9ecef;
}

.dialog-title {
  font-size: 24px;
  font-weight: 800;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  letter-spacing: -0.025em;
  line-height: 1.4;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.user-modal :deep(.el-dialog__header) {
  background: white;
  padding: 24px 32px 20px 32px;
  margin: 0;
  border-bottom: 1px solid #e9ecef;
  position: relative;
}

.user-modal :deep(.el-dialog__title) {
  color: white;
  font-size: 24px;
  font-weight: 700;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  letter-spacing: -0.025em;
  line-height: 1.4;
}

.user-modal :deep(.el-dialog__headerbtn) {
  top: 20px;
  right: 24px;
  width: 36px;
  height: 36px;
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.15);
  transition: all 0.2s ease;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.user-modal :deep(.el-dialog__headerbtn:hover) {
  background: rgba(255, 255, 255, 0.25);
  transform: scale(1.05);
}

.user-modal :deep(.el-dialog__close) {
  color: white;
  font-size: 20px;
  font-weight: 500;
}

.user-modal :deep(.el-dialog__body) {
  padding: 32px;
  background: white;
}

.user-modal :deep(.el-dialog__footer) {
  display: none;
}

.modal-content {
  padding: 0;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 24px;
  padding-top: 20px;
  border-top: 1px solid #f3f4f6;
}

.cancel-btn {
  border-radius: 10px !important;
  background: #f8f9fa !important;
  border: 1px solid #dee2e6 !important;
  color: #6c757d !important;
  font-weight: 500 !important;
  padding: 10px 20px !important;
  transition: all 0.2s ease !important;
  font-size: 14px !important;
}

.cancel-btn:hover {
  background: #e9ecef !important;
  border-color: #adb5bd !important;
  transform: translateY(-1px) !important;
}

.submit-btn {
  border-radius: 10px !important;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
  border: none !important;
  color: white !important;
  font-weight: 600 !important;
  padding: 10px 24px !important;
  transition: all 0.2s ease !important;
  font-size: 14px !important;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3) !important;
}

.submit-btn:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 6px 16px rgba(102, 126, 234, 0.4) !important;
}

@media (max-width: 768px) {
  .user-modal :deep(.el-dialog) {
    width: 95% !important;
    margin: 4% auto !important;
  }
  
  .modal-content {
    padding: 24px 24px 20px 24px;
  }
  
  .form-actions {
    flex-direction: column;
    gap: 12px;
  }
  
  .submit-btn, .cancel-btn {
    width: 100% !important;
  }
}
</style> 